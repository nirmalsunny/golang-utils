/*
 * Copyright (C) 2020-2022 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */
// Concurrency module
package parallelisation

import (
	"context"
	"reflect"
	"time"

	"go.uber.org/atomic"

	"github.com/ARM-software/golang-utils/utils/commonerrors"
)

// DetermineContextError determines what the context error is if any.
func DetermineContextError(ctx context.Context) error {
	return commonerrors.ConvertContextError(ctx.Err())
}

type result struct {
	Item interface{}
	err  error
}

// Parallelise parallelises an action over as many goroutines as specified by the argList and retrieves all the results when all the goroutines are done.
func Parallelise(argList interface{}, action func(arg interface{}) (interface{}, error), resultType reflect.Type) (results interface{}, err error) {
	keepReturn := resultType != nil
	argListValue := reflect.ValueOf(argList)
	length := argListValue.Len()
	channel := make(chan result, length)
	for i := 0; i < length; i++ {
		go func(args reflect.Value, actionFunc func(arg interface{}) (interface{}, error)) {
			var r result
			r.Item, r.err = func(v reflect.Value) (interface{}, error) {
				return actionFunc(v.Interface())
			}(args)
			channel <- r
		}(argListValue.Index(i), action)
	}
	var v reflect.Value
	if keepReturn {
		v = reflect.MakeSlice(resultType, 0, length)
	}
	for i := 0; i < length; i++ {
		r := <-channel
		err = r.err
		if err != nil {
			return
		}
		if keepReturn {
			v = reflect.Append(v, reflect.ValueOf(r.Item))
		}
	}
	if keepReturn {
		results = v.Interface()
	}
	return
}

// SleepWithContext performs an interruptable sleep
// Similar to time.Sleep() but also responding to context cancellation instead of blocking for the whole length of time.
func SleepWithContext(ctx context.Context, delay time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(delay):
	}
}

// SleepWithInterruption performs an interruptable sleep
// Similar to time.Sleep() but also interrupting when requested instead of blocking for the whole length of time.
func SleepWithInterruption(stop chan bool, delay time.Duration) {
	select {
	case <-stop:
	case <-time.After(delay):
	}
}

// ScheduleAfter calls once function `f` after `offset`
func ScheduleAfter(ctx context.Context, offset time.Duration, f func(time.Time)) {
	err := DetermineContextError(ctx)
	if err != nil {
		return
	}
	timer := time.NewTimer(offset)
	go func(ctx context.Context, function func(time.Time)) {
		select {
		case v := <-timer.C:
			function(v)
		case <-ctx.Done():
			timer.Stop()
			return
		}
	}(ctx, f)
}

// Schedule calls function `f` regularly with a `period` and an `offset`.
func Schedule(ctx context.Context, period time.Duration, offset time.Duration, f func(time.Time)) {
	err := DetermineContextError(ctx)
	if err != nil {
		return
	}
	go func(ctx context.Context, period time.Duration, offset time.Duration, function func(time.Time)) {
		// Position the first execution
		first := time.Now().Truncate(period).Add(offset)
		if first.Before(time.Now()) {
			first = first.Add(period)
		}
		firstC := time.After(time.Until(first))

		// Receiving from a nil channel blocks forever
		t := &time.Ticker{C: nil}

		for {
			select {
			case v := <-firstC:
				// The ticker has to be started before f as it can take some time to finish
				t = time.NewTicker(period)
				function(v)
			case v := <-t.C:
				function(v)
			case <-ctx.Done():
				t.Stop()
				return
			}
		}
	}(ctx, period, offset, f)
}

// RunActionWithTimeout runs an action with timeout
func RunActionWithTimeout(blockingAction func(stop chan bool) error, timeout time.Duration) (err error) {
	channel := make(chan error, 1)
	stop := make(chan bool)
	completed := atomic.NewBool(false)

	go func(action func(stop chan bool) error) {
		channel <- action(stop)
	}(blockingAction)

	select {
	case err = <-channel:
		completed.Store(true)
	case <-time.After(timeout):
		stop <- true
		err = commonerrors.ErrTimeout
	}
	if !completed.Load() {
		<-channel
	}
	return
}

// RunActionWithTimeoutAndContext runs an action with timeout
// blockingAction's context will be cancelled on exit.
func RunActionWithTimeoutAndContext(ctx context.Context, timeout time.Duration, blockingAction func(context.Context) error) error {
	store := NewCancelFunctionsStore()
	defer store.Cancel()
	return RunActionWithTimeoutAndCancelStore(ctx, timeout, store, blockingAction)
}

// RunActionWithTimeoutAndCancelStore runs an action with timeout
// The cancel store is used just to register the cancel function so that it can be called on Cancel.
func RunActionWithTimeoutAndCancelStore(ctx context.Context, timeout time.Duration, store *CancelFunctionStore, blockingAction func(context.Context) error) error {
	err := DetermineContextError(ctx)
	if err != nil {
		return err
	}
	timeoutContext, timeoutCancel := context.WithTimeout(ctx, timeout)
	store.RegisterCancelFunction(timeoutCancel)
	defer timeoutCancel()
	cancelCtx, actionCancel := context.WithCancel(ctx)
	store.RegisterCancelFunction(actionCancel)
	channel := make(chan error, 1)
	go func(actionCtx context.Context, action func(context.Context) error) {
		channel <- action(actionCtx)
	}(cancelCtx, blockingAction)

	select {
	case err = <-channel:
		if err != nil {
			actionCancel()
			<-cancelCtx.Done()
		}
		err2 := DetermineContextError(timeoutContext)
		if err2 != nil {
			return err2
		}
		timeoutCancel()
		return err
	case <-timeoutContext.Done():
		actionCancel()
		timeoutCancel()
		<-cancelCtx.Done()
		<-channel
		return DetermineContextError(timeoutContext)
	}
}

// RunActionWithParallelCheck runs an action with a check in parallel
// The function performing the check should return true if the check was favorable; false otherwise. If the check did not have the expected result and the whole function would be cancelled.
func RunActionWithParallelCheck(ctx context.Context, action func(ctx context.Context) error, checkAction func(ctx context.Context) bool, checkPeriod time.Duration) error {
	err := DetermineContextError(ctx)
	if err != nil {
		return err
	}
	cancelStore := NewCancelFunctionsStore()
	defer cancelStore.Cancel()
	cancellableCtx, cancelFunc := context.WithCancel(ctx)
	cancelStore.RegisterCancelFunction(cancelFunc)
	go func(ctx context.Context, store *CancelFunctionStore) {
		for {
			select {
			case <-ctx.Done():
				store.Cancel()
				return
			default:
				if !checkAction(ctx) {
					store.Cancel()
					return
				}
				SleepWithContext(ctx, checkPeriod)
			}
		}
	}(cancellableCtx, cancelStore)
	err = action(cancellableCtx)
	err2 := DetermineContextError(cancellableCtx)
	if err2 != nil {
		return err2
	}
	return err
}

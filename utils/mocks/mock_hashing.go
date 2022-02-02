/*
 * Copyright (C) 2020-2022 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ARM-software/golang-utils/utils/hashing (interfaces: IHash)

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIHash is a mock of IHash interface
type MockIHash struct {
	ctrl     *gomock.Controller
	recorder *MockIHashMockRecorder
}

// MockIHashMockRecorder is the mock recorder for MockIHash
type MockIHashMockRecorder struct {
	mock *MockIHash
}

// NewMockIHash creates a new mock instance
func NewMockIHash(ctrl *gomock.Controller) *MockIHash {
	mock := &MockIHash{ctrl: ctrl}
	mock.recorder = &MockIHashMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIHash) EXPECT() *MockIHashMockRecorder {
	return m.recorder
}

// Calculate mocks base method
func (m *MockIHash) Calculate(arg0 io.Reader) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Calculate", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Calculate indicates an expected call of Calculate
func (mr *MockIHashMockRecorder) Calculate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Calculate", reflect.TypeOf((*MockIHash)(nil).Calculate), arg0)
}

// GetType mocks base method
func (m *MockIHash) GetType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetType indicates an expected call of GetType
func (mr *MockIHashMockRecorder) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockIHash)(nil).GetType))
}

/*
 * Copyright (C) 2020-2022 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ARM-software/golang-utils/utils/logs (interfaces: Loggers,WriterWithSource)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLoggers is a mock of Loggers interface
type MockLoggers struct {
	ctrl     *gomock.Controller
	recorder *MockLoggersMockRecorder
}

// MockLoggersMockRecorder is the mock recorder for MockLoggers
type MockLoggersMockRecorder struct {
	mock *MockLoggers
}

// NewMockLoggers creates a new mock instance
func NewMockLoggers(ctrl *gomock.Controller) *MockLoggers {
	mock := &MockLoggers{ctrl: ctrl}
	mock.recorder = &MockLoggersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLoggers) EXPECT() *MockLoggersMockRecorder {
	return m.recorder
}

// Check mocks base method
func (m *MockLoggers) Check() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check")
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check
func (mr *MockLoggersMockRecorder) Check() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockLoggers)(nil).Check))
}

// Close mocks base method
func (m *MockLoggers) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockLoggersMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLoggers)(nil).Close))
}

// Log mocks base method
func (m *MockLoggers) Log(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Log", varargs...)
}

// Log indicates an expected call of Log
func (mr *MockLoggersMockRecorder) Log(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLoggers)(nil).Log), arg0...)
}

// LogError mocks base method
func (m *MockLoggers) LogError(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogError", varargs...)
}

// LogError indicates an expected call of LogError
func (mr *MockLoggersMockRecorder) LogError(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogError", reflect.TypeOf((*MockLoggers)(nil).LogError), arg0...)
}

// SetLogSource mocks base method
func (m *MockLoggers) SetLogSource(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLogSource", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLogSource indicates an expected call of SetLogSource
func (mr *MockLoggersMockRecorder) SetLogSource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogSource", reflect.TypeOf((*MockLoggers)(nil).SetLogSource), arg0)
}

// SetLoggerSource mocks base method
func (m *MockLoggers) SetLoggerSource(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLoggerSource", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLoggerSource indicates an expected call of SetLoggerSource
func (mr *MockLoggersMockRecorder) SetLoggerSource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLoggerSource", reflect.TypeOf((*MockLoggers)(nil).SetLoggerSource), arg0)
}

// MockWriterWithSource is a mock of WriterWithSource interface
type MockWriterWithSource struct {
	ctrl     *gomock.Controller
	recorder *MockWriterWithSourceMockRecorder
}

// MockWriterWithSourceMockRecorder is the mock recorder for MockWriterWithSource
type MockWriterWithSourceMockRecorder struct {
	mock *MockWriterWithSource
}

// NewMockWriterWithSource creates a new mock instance
func NewMockWriterWithSource(ctrl *gomock.Controller) *MockWriterWithSource {
	mock := &MockWriterWithSource{ctrl: ctrl}
	mock.recorder = &MockWriterWithSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriterWithSource) EXPECT() *MockWriterWithSourceMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockWriterWithSource) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockWriterWithSourceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWriterWithSource)(nil).Close))
}

// SetSource mocks base method
func (m *MockWriterWithSource) SetSource(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSource", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSource indicates an expected call of SetSource
func (mr *MockWriterWithSourceMockRecorder) SetSource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSource", reflect.TypeOf((*MockWriterWithSource)(nil).SetSource), arg0)
}

// Write mocks base method
func (m *MockWriterWithSource) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockWriterWithSourceMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockWriterWithSource)(nil).Write), arg0)
}

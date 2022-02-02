/*
 * Copyright (C) 2020-2022 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ARM-software/golang-utils/utils/config (interfaces: IServiceConfiguration)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIServiceConfiguration is a mock of IServiceConfiguration interface
type MockIServiceConfiguration struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceConfigurationMockRecorder
}

// MockIServiceConfigurationMockRecorder is the mock recorder for MockIServiceConfiguration
type MockIServiceConfigurationMockRecorder struct {
	mock *MockIServiceConfiguration
}

// NewMockIServiceConfiguration creates a new mock instance
func NewMockIServiceConfiguration(ctrl *gomock.Controller) *MockIServiceConfiguration {
	mock := &MockIServiceConfiguration{ctrl: ctrl}
	mock.recorder = &MockIServiceConfigurationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIServiceConfiguration) EXPECT() *MockIServiceConfigurationMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockIServiceConfiguration) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockIServiceConfigurationMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockIServiceConfiguration)(nil).Validate))
}

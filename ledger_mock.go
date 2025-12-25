//go:build ledger_mock
// +build ledger_mock

// Copyright (C) 2019-2025, Lux Industries Inc. All rights reserved.
// Forked from github.com/zondax/ledger-go
// Licensed under the Apache License, Version 2.0

package ledger_go

import "errors"

type LedgerAdminMock struct{}

type LedgerDeviceMock struct{}

func NewLedgerAdmin() LedgerAdmin {
	return &LedgerAdminMock{}
}

func (admin *LedgerAdminMock) CountDevices() int {
	return 1
}

func (admin *LedgerAdminMock) ListDevices() ([]string, error) {
	return []string{"mock"}, nil
}

func (admin *LedgerAdminMock) Connect(deviceIndex int) (LedgerDevice, error) {
	if deviceIndex != 0 {
		return nil, errors.New("device not found")
	}
	return &LedgerDeviceMock{}, nil
}

func (ledger *LedgerDeviceMock) Exchange(command []byte) ([]byte, error) {
	// Return a mock response with success status code
	return []byte{0x90, 0x00}, nil
}

func (ledger *LedgerDeviceMock) Close() error {
	return nil
}

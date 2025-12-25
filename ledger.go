// Copyright (C) 2019-2025, Lux Industries Inc. All rights reserved.
// Forked from github.com/zondax/ledger-go
// Licensed under the Apache License, Version 2.0

package ledger_go

// LedgerAdmin defines the interface for managing Ledger devices.
type LedgerAdmin interface {
	CountDevices() int
	ListDevices() ([]string, error)
	Connect(deviceIndex int) (LedgerDevice, error)
}

// LedgerDevice defines the interface for interacting with a Ledger device.
type LedgerDevice interface {
	Exchange(command []byte) ([]byte, error)
	Close() error
}

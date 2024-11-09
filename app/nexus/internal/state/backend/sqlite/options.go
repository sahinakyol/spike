//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package sqlite

import (
	"fmt"
	"time"
)

// Options defines SQLite-specific configuration options
type Options struct {
	// DataDir specifies the directory where the SQLite database file will be stored
	DataDir string

	// DatabaseFile specifies the name of the SQLite database file
	DatabaseFile string

	// JournalMode specifies the SQLite journal mode (DELETE, WAL, MEMORY, etc.)
	JournalMode string

	// BusyTimeoutMs specifies the busy timeout in milliseconds
	BusyTimeoutMs int

	// MaxOpenConns specifies the maximum number of open connections
	MaxOpenConns int

	// MaxIdleConns specifies the maximum number of idle connections
	MaxIdleConns int

	// ConnMaxLifetime specifies the maximum amount of time a connection may be reused
	ConnMaxLifetime time.Duration
}

// DefaultOptions returns the default SQLite options
func DefaultOptions() *Options {
	return &Options{
		DataDir:         "data",
		DatabaseFile:    "state.db",
		JournalMode:     "WAL",
		BusyTimeoutMs:   5000,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
	}
}

// parseOptions parses and validates the provided backend options
func parseOptions(opts map[string]any) (*Options, error) {
	if opts == nil {
		return DefaultOptions(), nil
	}

	sqliteOpts := &Options{}

	// Parse each field from the map
	if dataDir, ok := opts["DataDir"].(string); ok {
		sqliteOpts.DataDir = dataDir
	}
	if dbFile, ok := opts["DatabaseFile"].(string); ok {
		sqliteOpts.DatabaseFile = dbFile
	}
	if journalMode, ok := opts["JournalMode"].(string); ok {
		sqliteOpts.JournalMode = journalMode
	}
	if busyTimeout, ok := opts["BusyTimeoutMs"].(int); ok {
		sqliteOpts.BusyTimeoutMs = busyTimeout
	}
	if maxOpen, ok := opts["MaxOpenConns"].(int); ok {
		sqliteOpts.MaxOpenConns = maxOpen
	}
	if maxIdle, ok := opts["MaxIdleConns"].(int); ok {
		sqliteOpts.MaxIdleConns = maxIdle
	}
	if lifetime, ok := opts["ConnMaxLifetime"].(time.Duration); ok {
		sqliteOpts.ConnMaxLifetime = lifetime
	}

	// Apply defaults for zero values
	if sqliteOpts.DataDir == "" {
		sqliteOpts.DataDir = DefaultOptions().DataDir
	}
	if sqliteOpts.DatabaseFile == "" {
		sqliteOpts.DatabaseFile = DefaultOptions().DatabaseFile
	}
	if sqliteOpts.JournalMode == "" {
		sqliteOpts.JournalMode = DefaultOptions().JournalMode
	}
	if sqliteOpts.BusyTimeoutMs == 0 {
		sqliteOpts.BusyTimeoutMs = DefaultOptions().BusyTimeoutMs
	}
	if sqliteOpts.MaxOpenConns == 0 {
		sqliteOpts.MaxOpenConns = DefaultOptions().MaxOpenConns
	}
	if sqliteOpts.MaxIdleConns == 0 {
		sqliteOpts.MaxIdleConns = DefaultOptions().MaxIdleConns
	}
	if sqliteOpts.ConnMaxLifetime == 0 {
		sqliteOpts.ConnMaxLifetime = DefaultOptions().ConnMaxLifetime
	}

	// Validate options
	if sqliteOpts.MaxIdleConns > sqliteOpts.MaxOpenConns {
		return nil, fmt.Errorf("MaxIdleConns (%d) cannot be greater than MaxOpenConns (%d)",
			sqliteOpts.MaxIdleConns, sqliteOpts.MaxOpenConns)
	}

	return sqliteOpts, nil
}

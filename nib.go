// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package nib provides log-level based line printer for human-facing Go CLI applications.
// This package represents the currently latest API v0.
package nib

// Nib is a log-level based line printer for human-facing messages.
type Nib interface {
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Successf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
}

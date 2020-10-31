// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package nib provides log-level based line printer for human-facing Go CLI applications.
// This package represents the currently latest API v0.
package nib

import "io"

// Nib is a log-level based line printer for human-facing messages.
type Nib interface {
	// Compile compiles a message for the verbosity level using the given format and arguments.
	Compile(v Verbosity, format string, args ...interface{}) string
	// Debugf writes  with nib.DebugVerbosity level for the given format and arguments.
	Debugf(format string, args ...interface{})
	// Errorf writes a message with nib.ErrorVerbosity level for the given format and arguments.
	Errorf(format string, args ...interface{})
	// Fatalf writes a message with nib.FatalVerbosity level for the given format and arguments.
	Fatalf(format string, args ...interface{})
	// Infof writes a message with nib.InfoVerbosity level for the given format and arguments.
	Infof(format string, args ...interface{})
	// Successf writes a message with nib.SuccessVerbosity level for the given format and arguments.
	Successf(format string, args ...interface{})
	// Warnf writes a message with nib.WarnVerbosity level for the given format and arguments.
	Warnf(format string, args ...interface{})
	// Writer returns the underlying io.Writer.
	Writer() io.Writer
}

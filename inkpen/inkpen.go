// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package inkpen composes pencil.Pencil to support colored output including automatic TTY and terminal color detection.
package inkpen

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/svengreb/nib"
	"github.com/svengreb/nib/pencil"
)

// Inkpen composes pencil.Pencil to support colored output including automatic TTY and terminal color detection.
type Inkpen struct {
	*pencil.Pencil
	opts Options
}

// New creates and returns a new inkpen with default Options and pencil.Options.
func New(opts ...Option) *Inkpen {
	opt := NewOptions(opts...)
	p := pencil.NewFrom(opt.pencilOpts)
	return &Inkpen{
		Pencil: p,
		opts:   opt,
	}
}

// Compile compiles the format and arguments, ensuring a trailing newline, when the given verbosity level is enabled.
func (i *Inkpen) Compile(v nib.Verbosity, format string, args ...interface{}) string {
	if i.Enabled(v) {
		return i.formatMsg(v, format, args...)
	}
	return ""
}

// Debugf writes a debug message and ensures a trailing newline for the given format.
func (i *Inkpen) Debugf(msg string, args ...interface{}) {
	i.writef(nib.DebugVerbosity, msg, args...)
}

// Enabled checks if the verbosity is greater than the given verbosity.
func (i *Inkpen) Enabled(v nib.Verbosity) bool { return i.Pencil.GetVerbosity() >= v }

// Errorf writes a error message and ensures a trailing newline for the given format.
func (i *Inkpen) Errorf(msg string, args ...interface{}) {
	i.writef(nib.ErrorVerbosity, msg, args...)
}

// Fatalf writes a fatal message and ensures a trailing newline for the given format.
func (i *Inkpen) Fatalf(msg string, args ...interface{}) {
	i.writef(nib.FatalVerbosity, msg, args...)
}

// Infof writes a information message and ensures a trailing newline for the given format.
func (i *Inkpen) Infof(msg string, args ...interface{}) {
	i.writef(nib.InfoVerbosity, msg, args...)
}

// Successf writes a success message and ensures a trailing newline for the given format.
func (i *Inkpen) Successf(msg string, args ...interface{}) {
	i.writef(nib.SuccessVerbosity, msg, args...)
}

// Warnf writes a warning message and ensures a trailing newline for the given format.
func (i *Inkpen) Warnf(msg string, args ...interface{}) {
	i.writef(nib.WarnVerbosity, msg, args...)
}

// formatMsg formats the message and ensures a trailing newline for the given format.
func (i *Inkpen) formatMsg(v nib.Verbosity, format string, args ...interface{}) string {
	msg := i.Pencil.Compile(v, format, args...)

	if msg != "" && i.Pencil.IconsEnabled() && i.opts.coloredIcons {
		icf := i.opts.iconColorFuncs[v]
		icon, _ := utf8.DecodeRuneInString(msg)
		if icon != utf8.RuneError {
			msg = strings.Replace(msg, string(icon), icf(string(icon)), 1)
		}
	}

	return msg
}

// writef writes to the underlying writer and ensures a trailing newline for the given format.
func (i *Inkpen) writef(v nib.Verbosity, format string, args ...interface{}) {
	if i.Enabled(v) {
		msg := i.formatMsg(v, format, args...)

		if _, err := fmt.Fprint(i.opts.pencilOpts.Writer, msg); err != nil {
			_, _ = fmt.Fprint(os.Stderr, msg)
		}
	}
}

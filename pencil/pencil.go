// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package pencil implements nib.Nib to provide a writer for human-facing CLI messages with support for custom prefixes
// and verbosity level icons.
package pencil

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"github.com/svengreb/nib"
)

// Pencil is a writer for human-facing CLI messages with support for custom prefixes and verbosity level icons.
type Pencil struct {
	opts *Options
}

// New creates new pencil with default Options and merges them with the given list of Option.
func New(opts ...Option) *Pencil { return &Pencil{opts: NewOptions(opts...)} }

// Compile compiles the format and arguments, ensuring a trailing newline, when the given verbosity level is enabled.
func (p *Pencil) Compile(v nib.Verbosity, format string, args ...interface{}) string {
	if p.Enabled(v) {
		return p.formatMsg(v, format, args...)
	}
	return ""
}

// Debugf writes a debug message and ensures a trailing newline for the given format.
func (p *Pencil) Debugf(format string, args ...interface{}) {
	p.writef(nib.DebugVerbosity, format, args...)
}

// Enabled checks if the verbosity is greater than the given verbosity.
func (p *Pencil) Enabled(v nib.Verbosity) bool { return p.opts.Verbosity >= v }

// Errorf writes a error message and ensures a trailing newline for the given format.
func (p *Pencil) Errorf(format string, args ...interface{}) {
	p.writef(nib.ErrorVerbosity, format, args...)
}

// Fatalf writes a fatal message and ensures a trailing newline for the given format.
func (p *Pencil) Fatalf(format string, args ...interface{}) {
	p.writef(nib.FatalVerbosity, format, args...)
}

// GetVerbosity returns the nib.Verbosity level.
func (p *Pencil) GetVerbosity() nib.Verbosity {
	return p.opts.Verbosity
}

// IconsEnabled indicates whether verbosity level icons are enabled.
func (p *Pencil) IconsEnabled() bool {
	return p.opts.UseIcons && unicode.IsPrint(p.opts.Icons[p.opts.Verbosity])
}

// Infof writes a information message and ensures a trailing newline for the given format.
func (p *Pencil) Infof(format string, args ...interface{}) {
	p.writef(nib.InfoVerbosity, format, args...)
}

// Successf writes a success message and ensures a trailing newline for the given format.
func (p *Pencil) Successf(format string, args ...interface{}) {
	p.writef(nib.SuccessVerbosity, format, args...)
}

// Warnf writes a warning message and ensures a trailing newline for the given format.
func (p *Pencil) Warnf(format string, args ...interface{}) {
	p.writef(nib.WarnVerbosity, format, args...)
}

// Writer returns the underlying io.Writer.
func (p *Pencil) Writer() io.Writer {
	return p.opts.Writer
}

// formatMsg formats the message and ensures a trailing newline for the given format.
func (p *Pencil) formatMsg(v nib.Verbosity, format string, args ...interface{}) string {
	if len(args) > 0 || !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	msg := append(p.opts.Prefixes, format)
	if p.opts.UseIcons {
		icon, hasIcon := p.opts.Icons[v]
		if hasIcon {
			msg = append([]string{string(icon)}, msg...)
		}
	}

	return fmt.Sprintf(strings.Join(msg, " "), args...)
}

// writef writes to the underlying writer.
// If an error occurs while writing to the underlying io.Writer the message is printed to os.Stdout instead.
// When this also returns an error the error is written to os.Stderr instead.
func (p *Pencil) writef(v nib.Verbosity, format string, args ...interface{}) {
	if p.Enabled(v) {
		msg := p.formatMsg(v, format, args...)
		if _, err := fmt.Fprint(p.opts.Writer, msg); err != nil {
			if _, err = fmt.Fprint(os.Stdout, msg); err != nil {
				_, _ = fmt.Fprint(os.Stderr, err)
			}
		}
	}
}

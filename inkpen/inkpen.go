// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package inkpen implements nib.Nib to provide a colorized writer for human-facing CLI messages with support for custom
// prefixes and verbosity level icons.
package inkpen

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/svengreb/nib"
)

// The default inkpen color functions.
var (
	DefaultDebugColorFunc   = color.New(color.FgMagenta).Sprintf
	DefaultErrorColorFunc   = color.New(color.FgRed).Sprintf
	DefaultFatalColorFunc   = color.New(color.FgRed, color.Bold).Sprintf
	DefaultInfoColorFunc    = color.New(color.FgBlue).Sprintf
	DefaultSuccessColorFunc = color.New(color.FgGreen).Sprintf
	DefaultWarnColorFunc    = color.New(color.FgYellow).Sprintf
)

// Default inkpen verbosity level icons.
var (
	DefaultDebugIcon   = '⦿'
	DefaultErrorIcon   = '✕'
	DefaultFatalIcon   = '⭍'
	DefaultInfoIcon    = '➜'
	DefaultSuccessIcon = '✓'
	DefaultWarnIcon    = '!'
)

var (
	// DefaultVerbosity is the default inkpen verbosity level.
	DefaultVerbosity = nib.InfoVerbosity
	// DefaultWriter is the default inkpen writer.
	// It automatically detects TTY and terminal color support.
	DefaultWriter = color.Output
)

// Inkpen is a colorized writer for human-facing CLI messages with support for custom prefixes and verbosity level
// icons.
type Inkpen struct {
	WithIcon bool

	icons     map[nib.Verbosity]Icon
	prefixes  []string
	verbosity nib.Verbosity
	writer    io.Writer
}

// Icon is a verbosity level icon.
type Icon struct {
	ColorFunc func(format string, args ...interface{}) string
	Value     rune
}

// NewDefault creates and returns a new inkpen with default configurations.
func NewDefault() *Inkpen {
	return &Inkpen{
		WithIcon:  true,
		icons:     getDefaultIcons(),
		verbosity: DefaultVerbosity,
		writer:    DefaultWriter,
	}
}

// New creates and returns a new inkpen.
func New(v nib.Verbosity, w io.Writer) *Inkpen {
	return &Inkpen{
		icons:     make(map[nib.Verbosity]Icon),
		verbosity: v,
		writer:    w,
	}
}

// Debugf writes a debug message and ensures a trailing newline for the given format.
func (i *Inkpen) Debugf(msg string, args ...interface{}) {
	i.writef(nib.DebugVerbosity, msg, args...)
}

// Enabled checks if the verbosity is greater than the given verbosity.
func (i *Inkpen) Enabled(v nib.Verbosity) bool { return i.verbosity >= v }

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

// SetIcons sets the inkpen verbosity level icons.
func (i *Inkpen) SetIcons(icons map[nib.Verbosity]Icon) {
	i.WithIcon = true
	for k, v := range icons {
		i.icons[k] = v
	}
}

// SetPrefixes sets the inkpen message prefixes.
func (i *Inkpen) SetPrefixes(p ...string) {
	i.prefixes = p
}

// SetVerbosity sets the inkpen verbosity level.
func (i *Inkpen) SetVerbosity(v nib.Verbosity) {
	i.verbosity = v
}

// Successf writes a success message and ensures a trailing newline for the given format.
func (i *Inkpen) Successf(msg string, args ...interface{}) {
	i.writef(nib.SuccessVerbosity, msg, args...)
}

// Warnf writes a warning message and ensures a trailing newline for the given format.
func (i *Inkpen) Warnf(msg string, args ...interface{}) {
	i.writef(nib.WarnVerbosity, msg, args...)
}

// getDefaultIcons returns a map with default verbosity level icons.
func getDefaultIcons() map[nib.Verbosity]Icon {
	return map[nib.Verbosity]Icon{
		nib.DebugVerbosity: {
			ColorFunc: DefaultDebugColorFunc,
			Value:     DefaultDebugIcon,
		},
		nib.ErrorVerbosity: {
			ColorFunc: DefaultErrorColorFunc,
			Value:     DefaultErrorIcon,
		},
		nib.FatalVerbosity: {
			ColorFunc: DefaultFatalColorFunc,
			Value:     DefaultFatalIcon,
		},
		nib.InfoVerbosity: {
			ColorFunc: DefaultInfoColorFunc,
			Value:     DefaultInfoIcon,
		},
		nib.SuccessVerbosity: {
			ColorFunc: DefaultSuccessColorFunc,
			Value:     DefaultSuccessIcon,
		},
		nib.WarnVerbosity: {
			ColorFunc: DefaultWarnColorFunc,
			Value:     DefaultWarnIcon,
		},
	}
}

// writef writes to the underlying writer and ensures a trailing newline for the given format.
func (i *Inkpen) writef(v nib.Verbosity, format string, args ...interface{}) {
	if i.Enabled(v) {
		if len(args) > 0 || !strings.HasSuffix(format, "\n") {
			format += "\n"
		}

		msg := append(i.prefixes, format)
		if i.WithIcon {
			icon, hasIcon := i.icons[v]
			if hasIcon {
				msg = append([]string{icon.ColorFunc(string(icon.Value))}, msg...)
			}
		}

		if _, err := fmt.Fprintf(i.writer, strings.Join(msg, " "), args...); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, strings.Join(msg, " "), args...)
		}
	}
}

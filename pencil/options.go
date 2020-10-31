// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package pencil

import (
	"io"
	"os"
	"unicode"

	"github.com/svengreb/nib"
)

var (
	// DefaultVerbosity is the default pencil verbosity level.
	DefaultVerbosity = nib.InfoVerbosity
	// DefaultWriter is the default pencil writer.
	DefaultWriter = os.Stdout
)

// Default pencil verbosity level icons.
var (
	DefaultDebugIcon   = '⦿'
	DefaultErrorIcon   = '✕'
	DefaultFatalIcon   = '⭍'
	DefaultInfoIcon    = '➜'
	DefaultSuccessIcon = '✓'
	DefaultWarnIcon    = '!'
)

// Options are all pencil options.
type Options struct {
	Icons     map[nib.Verbosity]rune
	Prefixes  []string
	UseIcons  bool
	Verbosity nib.Verbosity
	Writer    io.Writer
}

// Option is a pencil option.
type Option func(*Options)

// NewOptions creates new default Options and merges them with the given options.
// By default the verbosity level is DefaultVerbosity and icons are enabled with output to DefaultWriter.
func NewOptions(opts ...Option) *Options {
	opt := &Options{
		Icons:     GetDefaultIcons(),
		Prefixes:  []string{},
		UseIcons:  true,
		Verbosity: DefaultVerbosity,
		Writer:    DefaultWriter,
	}
	for _, o := range opts {
		o(opt)
	}
	return opt
}

// UseIcons sets whether the configured icons should be printed.
func UseIcons(useIcons bool) Option {
	return func(o *Options) {
		o.UseIcons = useIcons
	}
}

// WithIcons sets the pencil verbosity level icons.
// Note that the provided map gets merged with the default map in order to ensure there are no missing icons.
// The default icon map uses Unicode characters which are supported by almost all terminals nowadays,
// but there might be cases where simple ASCII characters work better instead.
func WithIcons(icons map[nib.Verbosity]rune) Option {
	return func(o *Options) {
		i := make(map[nib.Verbosity]rune)
		for verb, icon := range icons {
			var v rune
			if unicode.IsPrint(icon) {
				v = icon
				i[verb] = v
			}
			i[verb] = v
		}
		o.Icons = i
	}
}

// WithPrefixes sets the pencil message prefixes.
// By default only the verbosity level icons are printed as prefix before the given message format and arguments.
func WithPrefixes(prefixes ...string) Option {
	return func(o *Options) {
		o.Prefixes = prefixes
	}
}

// WithVerbosity sets the verbosity level.
func WithVerbosity(verbosity nib.Verbosity) Option {
	return func(o *Options) {
		o.Verbosity = verbosity
	}
}

// WithWriter sets the output writer.
func WithWriter(writer io.Writer) Option {
	return func(o *Options) {
		o.Writer = writer
	}
}

// GetDefaultIcons returns a map with default verbosity level icons.
func GetDefaultIcons() map[nib.Verbosity]rune {
	return map[nib.Verbosity]rune{
		nib.DebugVerbosity:   DefaultDebugIcon,
		nib.ErrorVerbosity:   DefaultErrorIcon,
		nib.FatalVerbosity:   DefaultFatalIcon,
		nib.InfoVerbosity:    DefaultInfoIcon,
		nib.SuccessVerbosity: DefaultSuccessIcon,
		nib.WarnVerbosity:    DefaultWarnIcon,
	}
}

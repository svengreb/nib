// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package inkpen

import (
	"github.com/fatih/color"

	"github.com/svengreb/nib"
	"github.com/svengreb/nib/pencil"
)

// DefaultWriter is the default inkpen writer.
// It automatically detects TTY and terminal color support.
var DefaultWriter = color.Output

// The default inkpen color functions.
var (
	DefaultDebugColorFunc   = color.New(color.FgMagenta).Sprintf
	DefaultErrorColorFunc   = color.New(color.FgRed).Sprintf
	DefaultFatalColorFunc   = color.New(color.FgRed, color.Bold).Sprintf
	DefaultInfoColorFunc    = color.New(color.FgBlue).Sprintf
	DefaultSuccessColorFunc = color.New(color.FgGreen).Sprintf
	DefaultWarnColorFunc    = color.New(color.FgYellow).Sprintf
)

// IconColorFunc is a function to color verbosity level icons.
type IconColorFunc func(format string, args ...interface{}) string

// Options are all inkpen options.
type Options struct {
	coloredIcons   bool
	iconColorFuncs map[nib.Verbosity]IconColorFunc
	pencilOpts     []pencil.Option
}

// Option is a inkpen option.
type Option func(*Options)

// NewOptions creates new default inkpen Options and pencil.Options and merges them with the given list of Option.
// By default the output writer is DefaultWriter and colorization for verbosity level icons is enabled.
func NewOptions(opts ...Option) *Options {
	opt := &Options{
		iconColorFuncs: getDefaultIconColorFuncs(),
		coloredIcons:   true,
		pencilOpts: []pencil.Option{
			pencil.WithWriter(DefaultWriter),
		},
	}
	for _, o := range opts {
		o(opt)
	}
	return opt
}

// UseColoredIcons sets whether icons should be colored.
func UseColoredIcons(colorIcons bool) Option {
	return func(o *Options) {
		o.coloredIcons = colorIcons
	}
}

// WithIconColorFuncs sets the given ist of inkpen.IconColorFunc.
func WithIconColorFuncs(iconColorFuncs map[nib.Verbosity]IconColorFunc) Option {
	return func(o *Options) {
		icfs := getDefaultIconColorFuncs()
		for verb, cf := range iconColorFuncs {
			icfs[verb] = cf
		}
		o.iconColorFuncs = icfs
	}
}

// WithPencilOptions sets the given list of pencil.Option.
func WithPencilOptions(opts ...pencil.Option) Option {
	return func(o *Options) {
		o.pencilOpts = append(o.pencilOpts, opts...)
	}
}

// getDefaultIconColorFuncs returns a map with default verbosity level icon color functions.
func getDefaultIconColorFuncs() map[nib.Verbosity]IconColorFunc {
	return map[nib.Verbosity]IconColorFunc{
		nib.DebugVerbosity:   DefaultDebugColorFunc,
		nib.ErrorVerbosity:   DefaultErrorColorFunc,
		nib.FatalVerbosity:   DefaultFatalColorFunc,
		nib.InfoVerbosity:    DefaultInfoColorFunc,
		nib.SuccessVerbosity: DefaultSuccessColorFunc,
		nib.WarnVerbosity:    DefaultWarnColorFunc,
	}
}

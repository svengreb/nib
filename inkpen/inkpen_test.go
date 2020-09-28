// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package inkpen_test

import (
	"bytes"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"

	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"
	"github.com/svengreb/nib/pencil"
)

func TestCompileMessage(t *testing.T) {
	buf := &bytes.Buffer{}
	ink := inkpen.New(inkpen.WithPencilOptions(pencil.WithWriter(buf)))

	fruit := "Raspberries"
	favCount := "second"
	msg := "%s are my %s favorite fruit"
	compiled := ink.Compile(nib.InfoVerbosity, msg, fruit, favCount)

	ink.Infof(msg, fruit, favCount)
	assert.Equalf(t, buf.String(), compiled, "%q != %q", compiled, buf.String())
}

func TestCompileMessageVerbosityDisabled(t *testing.T) {
	ink := inkpen.New()
	compiled := ink.Compile(nib.DebugVerbosity, "%s are my %s favorite fruit", "Raspberries", "second")

	assert.Emptyf(t, compiled, "must be empty, but was %q", compiled)
}

func TestDefaultVerbosity(t *testing.T) {
	buf := &bytes.Buffer{}
	ink := inkpen.New(inkpen.WithPencilOptions(pencil.WithWriter(buf)))

	ink.Debugf("debug")
	assert.Empty(t, buf.String())
	buf.Reset()

	ink.Infof("info")
	assert.NotEmpty(t, buf.String())
}

func TestVisualCustomIconColorFuncs(t *testing.T) {
	buf := &bytes.Buffer{}
	ink := inkpen.New(inkpen.WithPencilOptions(pencil.WithWriter(buf)))
	inkCustomIconColor := inkpen.New(
		inkpen.WithPencilOptions(pencil.WithWriter(buf)),
		inkpen.WithIconColorFuncs(map[nib.Verbosity]inkpen.IconColorFunc{
			nib.InfoVerbosity: color.New(color.FgGreen).Sprintf,
		}),
	)

	ink.Infof("With default icon color")
	t.Logf(buf.String())
	buf.Reset()

	inkCustomIconColor.Infof("With custom icon color")
	t.Logf(buf.String())
}

func TestVisualDisableColoredIcons(t *testing.T) {
	buf := &bytes.Buffer{}
	ink := inkpen.New(inkpen.WithPencilOptions(pencil.WithWriter(buf)))
	inkNoColoredIcons := inkpen.New(
		inkpen.WithPencilOptions(pencil.WithWriter(buf)),
		inkpen.UseColoredIcons(false),
	)

	ink.Infof("Colored icons enabled")
	t.Logf(buf.String())
	buf.Reset()

	inkNoColoredIcons.Infof("Colored icons disabled")
	t.Logf(buf.String())
}

func TestWriteVerbosityLevels(t *testing.T) {
	buf := &bytes.Buffer{}
	ink := inkpen.New(inkpen.WithPencilOptions(
		pencil.WithWriter(buf),
		pencil.WithVerbosity(nib.DebugVerbosity),
	))

	testCases := []struct {
		f   func(format string, args ...interface{})
		msg string
		v   nib.Verbosity
	}{
		{ink.Debugf, nib.VerbosityNameDebug, nib.DebugVerbosity},
		{ink.Infof, nib.VerbosityNameInfo, nib.InfoVerbosity},
		{ink.Successf, nib.VerbosityNameSuccess, nib.SuccessVerbosity},
		{ink.Errorf, nib.VerbosityNameError, nib.ErrorVerbosity},
		{ink.Fatalf, nib.VerbosityNameFatal, nib.FatalVerbosity},
		{ink.Warnf, nib.VerbosityNameWarn, nib.WarnVerbosity},
	}

	for _, tc := range testCases {
		compiled := ink.Compile(tc.v, tc.msg)
		tc.f(tc.msg)
		assert.Equalf(t, compiled, buf.String(), "%q != %q", buf.String(), compiled)
		buf.Reset()
	}
}

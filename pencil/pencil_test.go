// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package pencil_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"

	"github.com/svengreb/nib"
	"github.com/svengreb/nib/pencil"
)

func TestCompileMessage(t *testing.T) {
	buf := &bytes.Buffer{}
	pen := pencil.New(pencil.WithWriter(buf))

	fruit := "Raspberries"
	favCount := "second"
	msg := "%s are my %s favorite fruit"
	compiled := pen.Compile(nib.InfoVerbosity, msg, fruit, favCount)

	pen.Infof(msg, fruit, favCount)
	assert.Equalf(t, buf.String(), compiled, "%q != %q", compiled, buf.String())
}

func TestCompileMessageVerbosityDisabled(t *testing.T) {
	pen := pencil.New()
	compiled := pen.Compile(nib.DebugVerbosity, "%s are my %s favorite fruit", "Raspberries", "second")

	assert.Emptyf(t, compiled, "must be empty, but was %q", compiled)
}

func TestCustomIconValid(t *testing.T) {
	buf := &bytes.Buffer{}
	icons := pencil.GetDefaultIcons()
	icons[nib.InfoVerbosity] = 'i'
	icons[nib.SuccessVerbosity] = ' '
	pen := pencil.New(pencil.WithWriter(buf), pencil.WithIcons(icons))

	pen.Infof("")
	checkIconRune(t, buf, icons, nib.InfoVerbosity)
}

func TestCustomIconInvalid(t *testing.T) {
	buf := &bytes.Buffer{}
	icons := map[nib.Verbosity]rune{nib.InfoVerbosity: unicode.ReplacementChar}
	pen := pencil.New(pencil.WithWriter(buf), pencil.WithIcons(icons))

	pen.Infof("info")
	icon, _ := utf8.DecodeRuneInString(buf.String())
	assert.Equalf(t, utf8.RuneError, icon, "%c != %c", utf8.RuneError, icon)
}

func TestCustomPrefixes(t *testing.T) {
	buf := &bytes.Buffer{}
	prefix := "[prefix]"
	pen := pencil.New(pencil.WithWriter(buf), pencil.WithPrefixes(prefix))
	penNoIcons := pencil.New(pencil.WithWriter(buf), pencil.WithPrefixes(prefix), pencil.UseIcons(false))

	pen.Infof("")
	assert.Truef(t, strings.Contains(buf.String(), prefix), "%q not contains %q", buf.String(), prefix)
	buf.Reset()

	penNoIcons.Infof("")
	msgStart := strings.Split(buf.String(), " ")
	assert.Truef(t, strings.HasPrefix(buf.String(), prefix), "must start with %q, but was %q", prefix, msgStart[0])
}

func TestCustomVerbosity(t *testing.T) {
	buf := &bytes.Buffer{}
	pen := pencil.New(pencil.WithWriter(buf), pencil.WithVerbosity(nib.ErrorVerbosity))

	pen.Infof("info")
	assert.Empty(t, buf)
	buf.Reset()

	pen.Errorf("error")
	assert.NotEmpty(t, buf)
}

func TestCustomWriter(t *testing.T) {
	buf := &bytes.Buffer{}
	pen := pencil.New(pencil.WithWriter(buf))

	pen.Infof("")
	assert.NotEmpty(t, buf)
}

func TestDefaultIcons(t *testing.T) {
	buf := &bytes.Buffer{}
	pen := pencil.New(pencil.WithWriter(buf), pencil.WithVerbosity(nib.DebugVerbosity))
	defaultIcons := pencil.GetDefaultIcons()

	pen.Debugf("")
	checkIconRune(t, buf, defaultIcons, nib.DebugVerbosity)
	buf.Reset()

	pen.Infof("")
	checkIconRune(t, buf, defaultIcons, nib.InfoVerbosity)
	buf.Reset()

	pen.Successf("")
	checkIconRune(t, buf, defaultIcons, nib.SuccessVerbosity)
	buf.Reset()

	pen.Errorf("")
	checkIconRune(t, buf, defaultIcons, nib.ErrorVerbosity)
	buf.Reset()

	pen.Fatalf("")
	checkIconRune(t, buf, defaultIcons, nib.FatalVerbosity)
	buf.Reset()

	pen.Warnf("")
	checkIconRune(t, buf, defaultIcons, nib.WarnVerbosity)
}

func TestDefaultVerbosity(t *testing.T) {
	buf := &bytes.Buffer{}
	pen := pencil.New(pencil.WithWriter(buf))

	pen.Debugf("debug")
	assert.Empty(t, buf.String())
	buf.Reset()

	pen.Infof("info")
	assert.NotEmpty(t, buf.String())
}

func TestDisabledIcons(t *testing.T) {
	buf := &bytes.Buffer{}
	pen := pencil.New(pencil.WithWriter(buf), pencil.UseIcons(false))

	msg := "info"
	pen.Infof(msg)
	assert.Truef(t, strings.HasPrefix(buf.String(), msg), "%q != %q", buf.String(), msg)
	assert.False(t, pen.IconsEnabled())
}

func TestEnsureTrailingNewline(t *testing.T) {
	buf := &bytes.Buffer{}
	pen := pencil.New(pencil.WithWriter(buf))
	penNoIcons := pencil.New(pencil.WithWriter(buf), pencil.UseIcons(false))

	pen.Infof("")
	assert.NotEmpty(t, buf)
	assert.True(t, strings.HasSuffix(buf.String(), "\n"))
	buf.Reset()

	penNoIcons.Infof("\n\n")
	indexLastNewline := strings.LastIndex(buf.String(), "\n")
	indexFirstNewline := strings.IndexAny(buf.String(), "\n")
	assert.Truef(
		t,
		indexFirstNewline != indexLastNewline,
		"multiple newlines not kept: index %d of first newline not equals index %d",
		indexFirstNewline,
		indexLastNewline,
	)
	assert.True(t, strings.HasSuffix(buf.String(), "\n"))
	buf.Reset()

	penNoIcons.Infof("")
	assert.NotEmpty(t, buf)
	assert.True(t, strings.HasSuffix(buf.String(), "\n"))
}

func TestGetVerbosity(t *testing.T) {
	buf := &bytes.Buffer{}
	vExpected := nib.ErrorVerbosity
	pen := pencil.New(pencil.WithWriter(buf), pencil.WithVerbosity(vExpected))

	vGot := pen.GetVerbosity()
	assert.Equalf(t, vGot, vExpected, "expected %q, but got %q", vExpected, vGot)
}

func TestWriteVerbosityLevels(t *testing.T) {
	buf := &bytes.Buffer{}
	pen := pencil.New(pencil.WithWriter(buf), pencil.WithVerbosity(nib.DebugVerbosity))

	testCases := []struct {
		f   func(format string, args ...interface{})
		msg string
		v   nib.Verbosity
	}{
		{pen.Debugf, nib.VerbosityNameDebug, nib.DebugVerbosity},
		{pen.Infof, nib.VerbosityNameInfo, nib.InfoVerbosity},
		{pen.Successf, nib.VerbosityNameSuccess, nib.SuccessVerbosity},
		{pen.Errorf, nib.VerbosityNameError, nib.ErrorVerbosity},
		{pen.Fatalf, nib.VerbosityNameFatal, nib.FatalVerbosity},
		{pen.Warnf, nib.VerbosityNameWarn, nib.WarnVerbosity},
	}

	for _, tc := range testCases {
		compiled := pen.Compile(tc.v, tc.msg)
		tc.f(tc.msg)
		assert.Equalf(t, compiled, buf.String(), "%q != %q", buf.String(), compiled)
		buf.Reset()
	}
}

func TestWriteVerbosityLevelDisabled(t *testing.T) {
	buf := &bytes.Buffer{}
	pen := pencil.New(pencil.WithWriter(buf))

	pen.Debugf(nib.VerbosityNameDebug)
	assert.Empty(t, buf.String(), "must be empty, but was %q", buf.String())
}

func checkIconRune(t *testing.T, buf fmt.Stringer, icons map[nib.Verbosity]rune, v nib.Verbosity) {
	icon, _ := utf8.DecodeRuneInString(buf.String())
	if icon == utf8.RuneError {
		assert.Fail(t, "Decoding '%s' verbosity level icon from message (utf8.RuneError): %q", v, utf8.RuneError)
	}
	assert.Equalf(t, icons[v], icon, "%c != %c", icon, icons[v])
}

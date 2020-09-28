// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package nib_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/svengreb/nib"
)

func TestMarshalVerbosity(t *testing.T) {
	testCases := []struct {
		v    nib.Verbosity
		name string
	}{
		{nib.DebugVerbosity, nib.VerbosityNameDebug},
		{nib.InfoVerbosity, nib.VerbosityNameInfo},
		{nib.SuccessVerbosity, nib.VerbosityNameSuccess},
		{nib.ErrorVerbosity, nib.VerbosityNameError},
		{nib.FatalVerbosity, nib.VerbosityNameFatal},
		{nib.SuppressVerbosity, nib.VerbosityNameSuppress},
		{nib.WarnVerbosity, nib.VerbosityNameWarn},
	}

	for _, tc := range testCases {
		n, err := tc.v.MarshalText()
		if err != nil {
			assert.Fail(t, "marshaling verbosity %q name: %v", tc.name, err)
		}
		assert.Equalf(t, string(n), tc.name, "expected %q, but got %q", tc.name, n)
	}
}

func TestMarshalVerbosityInvalid(t *testing.T) {
	vInvalid := nib.Verbosity(math.MaxInt32)
	_, err := vInvalid.MarshalText()
	assert.NotNilf(t, err, "must fail to marshal invalid verbosity level: %v", err)
}

func TestParseVerbosity(t *testing.T) {
	testCases := []struct {
		v    nib.Verbosity
		name string
	}{
		{nib.DebugVerbosity, nib.VerbosityNameDebug},
		{nib.InfoVerbosity, nib.VerbosityNameInfo},
		{nib.SuccessVerbosity, nib.VerbosityNameSuccess},
		{nib.ErrorVerbosity, nib.VerbosityNameError},
		{nib.FatalVerbosity, nib.VerbosityNameFatal},
		{nib.SuppressVerbosity, nib.VerbosityNameSuppress},
		{nib.WarnVerbosity, nib.VerbosityNameWarn},
	}

	for _, tc := range testCases {
		v, err := nib.ParseVerbosity(tc.name)
		if err != nil {
			assert.Fail(t, "parsing %q verbosity name: %v", tc.name, err)
		}
		assert.Equalf(t, v, tc.v, "expected %q, but got %q", tc.v, v)
	}
}

func TestParseVerbosityInvalid(t *testing.T) {
	vName := "invalid"
	_, err := nib.ParseVerbosity(vName)
	assert.NotNilf(t, err, "must fail to parse invalid verbosity level name %q: %v", vName, err)
}

func TestUnmarshalVerbosity(t *testing.T) {
	vOld := nib.DebugVerbosity
	vNew := nib.ErrorVerbosity
	err := vOld.UnmarshalText([]byte(vNew.String()))
	assert.Nilf(t, err, "unmarshalling verbosity name %q", nib.ErrorVerbosity.String())
	assert.Equalf(t, vOld, vNew, "expected %q, but got %q", vNew, vOld)
}

func TestUnmarshalVerbosityInvalid(t *testing.T) {
	v := nib.DebugVerbosity
	invalidName := "invalid"
	err := v.UnmarshalText([]byte(invalidName))
	assert.NotNilf(t, err, "unmarshal error: %v", err)
}

func TestVerbosityUnknownName(t *testing.T) {
	assert.Equal(t, nib.Verbosity(math.MaxUint32).String(), nib.VerbosityNameUnknown)
}

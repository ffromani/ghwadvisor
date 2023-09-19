//go:build !linux && !windows && !wasm
// +build !linux,!windows,!wasm

// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package net

import (
	"runtime"

	"github.com/pkg/errors"
)

func (i *Info) load() error {
	return errors.New("netFillInfo not implemented on " + runtime.GOOS)
}

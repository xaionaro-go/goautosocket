//go:build linux || darwin
// +build linux darwin

package gas

import (
	"syscall"
)

func isConnResetErrorNix(err error) bool {
	if se, ok := err.(syscall.Errno); ok {
		return se == syscall.ECONNRESET || se == syscall.EPIPE
	}
	return false
}

func isConnRefusedErrorNix(err error) bool {
	if se, ok := err.(syscall.Errno); ok {
		return se == syscall.ECONNREFUSED
	}
	return false
}

func isConnResetError(err error) bool {
	return isConnResetErrorNix(err)
}

func isConnRefusedError(err error) bool {
	return isConnRefusedErrorNix(err)
}

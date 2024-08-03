//go:build windows
// +build windows

package gas

import (
	"syscall"
)

const (
	WSAECONNREFUSED syscall.Errno = 10061
)

func isConnResetErrorWin(err error) bool {
	if se, ok := err.(syscall.Errno); ok {
		return se == syscall.WSAECONNRESET || se == syscall.WSAECONNABORTED
	}
	return false
}

func isConnRefusedErrorWin(err error) bool {
	if se, ok := err.(syscall.Errno); ok {
		return se == WSAECONNREFUSED
	}
	return false
}

func isConnResetError(err error) bool {
	return isConnResetErrorWin(err)
}

func isConnRefusedError(err error) bool {
	return isConnRefusedErrorWin(err)
}

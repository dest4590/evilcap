//go:build windows
// +build windows

package zeroconf

import "syscall"

func setReuseAddr(fd uintptr) error {
	return syscall.SetsockoptInt(syscall.Handle(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
}

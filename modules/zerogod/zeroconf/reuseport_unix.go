//go:build !windows
// +build !windows

package zeroconf

import "syscall"

func setReuseAddr(fd uintptr) error {
	err := syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	if err != nil {
		return err
	}
	// Also try to set SO_REUSEPORT if available
	_ = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, 0xF, 1) // 0xF is SO_REUSEPORT on most unix
	return nil
}

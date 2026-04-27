//go:build !windows
// +build !windows

package network

import (
	"net"
)

func getInterfaceName(iface net.Interface) string {
	return iface.Name
}

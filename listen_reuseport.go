//go:build go1.11 && (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd)
// +build go1.11
// +build aix darwin dragonfly freebsd linux netbsd openbsd

package dns

import (
	"context"
	"net"
	"strings"
	"syscall"

	"golang.org/x/sys/unix"
)

const supportsReusePort = true

func reuseportControl(network, address string, c syscall.RawConn) error {
	var opErr error
	err := c.Control(func(fd uintptr) {
		opErr = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
	})
	if err != nil {
		return err
	}

	return opErr
}

func ListenTCP(network, addr string, reuseport bool) (net.Listener, error) {
	var lc net.ListenConfig
	if reuseport {
		lc.Control = reuseportControl
	}

	network = strings.TrimSuffix(network, "-tls")

	return lc.Listen(context.Background(), network, addr)
}

func ListenUDP(network, addr string, reuseport bool) (net.PacketConn, error) {
	var lc net.ListenConfig
	if reuseport {
		lc.Control = reuseportControl
	}

	network = strings.TrimSuffix(network, "-tls")

	return lc.ListenPacket(context.Background(), network, addr)
}

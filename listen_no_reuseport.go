//go:build !go1.11 || (!aix && !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd)
// +build !go1.11 !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd

package dns

import (
	"net"
	"strings"
)

const supportsReusePort = false

func ListenTCP(network, addr string, reuseport bool) (net.Listener, error) {
	network = strings.TrimSuffix(network, "-tls")

	if reuseport {
		// TODO(tmthrgd): return an error?
	}

	return net.Listen(network, addr)
}

func ListenUDP(network, addr string, reuseport bool) (net.PacketConn, error) {
	network = strings.TrimSuffix(network, "-tls")

	if reuseport {
		// TODO(tmthrgd): return an error?
	}

	return net.ListenPacket(network, addr)
}

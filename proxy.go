package smtp

import (
	"net"
)

func CreateProxiedDialer(socks_addr string) (net.Dialer, error) {
	return net.Dialer{}, nil
}

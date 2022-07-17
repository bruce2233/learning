package netpoll

import (
	"net"
	"testing"

	"github.com/panjf2000/gnet/v2/pkg/errors"
	"golang.org/x/sys/unix"
)

func TestEpoll(t *testing.T) {
	// type eventpoll = unix.EpollEvent
	poller, err := OpenPoller()
	t.Log(poller, err)
}

func TestSocket(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8086")
	sa6 := &unix.SockaddrInet6{Port: tcpAddr.Port}
	if tcpAddr.IP != nil {
		copy(sa6.Addr[:], tcpAddr.IP) // copy all bytes of slice to array
	}
	if tcpAddr.Zone != "" {
		var iface *net.Interface
		iface, err = net.InterfaceByName(tcpAddr.Zone)
		if err != nil {
			return
		}

		sa6.ZoneId = uint32(iface.Index)
	}
	t.Log(sa6)
	if err != nil {
		t.Log("error")
		return
	}
	t.Log(tcpAddr)
	net.Listen()
}

func GetTCPSockAddr(proto, addr string) (sa unix.Sockaddr, family int, tcpAddr *net.TCPAddr, ipv6only bool, err error) {
	var tcpVersion string

	tcpAddr, err = net.ResolveTCPAddr(proto, addr)
	if err != nil {
		return
	}

	tcpVersion, err = determineTCPProto(proto, tcpAddr)
	if err != nil {
		return
	}

	switch tcpVersion {
	case "tcp4":
		sa4 := &unix.SockaddrInet4{Port: tcpAddr.Port}

		if tcpAddr.IP != nil {
			if len(tcpAddr.IP) == 16 {
				copy(sa4.Addr[:], tcpAddr.IP[12:16]) // copy last 4 bytes of slice to array
			} else {
				copy(sa4.Addr[:], tcpAddr.IP) // copy all bytes of slice to array
			}
		}

		sa, family = sa4, unix.AF_INET
	case "tcp6":
		ipv6only = true
		fallthrough
	case "tcp":
		sa6 := &unix.SockaddrInet6{Port: tcpAddr.Port}

		if tcpAddr.IP != nil {
			copy(sa6.Addr[:], tcpAddr.IP) // copy all bytes of slice to array
		}

		if tcpAddr.Zone != "" {
			var iface *net.Interface
			iface, err = net.InterfaceByName(tcpAddr.Zone)
			if err != nil {
				return
			}

			sa6.ZoneId = uint32(iface.Index)
		}

		sa, family = sa6, unix.AF_INET6
	default:
		err = errors.ErrUnsupportedProtocol
	}

	return
}

func determineTCPProto(proto string, addr *net.TCPAddr) (string, error) {
	// If the protocol is set to "tcp", we try to determine the actual protocol
	// version from the size of the resolved IP address. Otherwise, we simple use
	// the protocol given to us by the caller.

	if addr.IP.To4() != nil {
		return "tcp4", nil
	}

	if addr.IP.To16() != nil {
		return "tcp6", nil
	}

	switch proto {
	case "tcp", "tcp4", "tcp6":
		return proto, nil
	}

	return "", errors.ErrUnsupportedTCPProtocol
}

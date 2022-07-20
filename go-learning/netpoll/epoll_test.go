package netpoll

import (
	"net"
	"testing"
	"time"

	"golang.org/x/sys/unix"
)

func TestSocket(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8866")
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
}

func TestTcpSocket(t *testing.T) {
	fd, netAddr, err := tcpSocket("tcp", "127.0.0.1:8866", true)
	if err != nil {
		t.Log("error")
	}
	t.Log(fd)
	t.Log(netAddr)
	// net.Listener.Accept()
	time.Sleep(1e12)

}

func TestEpollListener(t *testing.T) {
	socketFd, netAddr, err := tcpSocket("tcp", "127.0.0.1:8866", false)
	if err != nil {
		t.Log("socket error")
	}
	t.Log(socketFd)
	t.Log(netAddr)

	epollFd, err := unix.EpollCreate1(unix.EPOLL_CLOEXEC)
	if err != nil {
		t.Log(err)
	}
	t.Log(epollFd)

	unix.EpollCtl(epollFd, unix.EPOLL_CTL_ADD, socketFd, &unix.EpollEvent{
		Fd: int32(socketFd), Events: unix.EPOLLIN})
	el := make([]unix.EpollEvent, 128)
	n, err := unix.EpollWait(epollFd, el, -1)
	t.Log(el)
	t.Log(err)
	t.Log(n)
	bytes := make([]byte, 256)
	bytesNum, err := unix.Read(epollFd, bytes)
	if err != nil {
		t.Log(err)
	}
	t.Log(bytesNum)
}

func TestEpollConn(t *testing.T) {
	socketFd, err := unix.Open("./tmp.txt", unix.O_RDWR|unix.O_CREAT, 777)
	if err != nil {
		t.Log(err)
	}
	t.Log(socketFd)
	// t.Log(netAddr)

	epollFd, err := unix.EpollCreate1(unix.EPOLL_CLOEXEC)
	if err != nil {
	}
	t.Log(epollFd)

	unix.EpollCtl(epollFd, unix.EPOLL_CTL_ADD, socketFd, &unix.EpollEvent{
		Fd: int32(socketFd), Events: unix.EPOLLIN})
	el := make([]unix.EpollEvent, 128)
	// for {
	n, err := unix.EpollWait(epollFd, el, -1)
	t.Log(el)
	t.Log(err)
	t.Log(n)
	bytes := make([]byte, 256)
	bytesNum, err := unix.Read(epollFd, bytes)
	if err != nil {
		t.Log(err)
	}
	t.Log(bytesNum)
	// }

}

func TestWrite(t *testing.T) {
	fd, err := unix.Open("./tmp.txt", unix.O_RDWR, 0)
	if err != nil {
		t.Log(err)
	}
	epfd, err :=unix.EpollCreate1(1)
	if err!=nil{
		t.Log(err)
	}
	unix.
	unix.Write(fd, []byte("Hello IO"))
}

func TestWriteSocket(t *testing.T) {
	conn, err := net.Dial("tcp", "192.168.0.1:50341")
	if err != nil {
		t.Log(err)
	}
	conn.Write([]byte("Hello Socket"))
}

func TestSocketCall(t *testing.T) {
	sa, family, tcpAddr, ipv6only, err := GetTCPSockAddr("tcp", "127.0.0.1:8866")
	if err != nil {
		t.Log(err)
	}

	t.Log(sa)
	t.Log(family)
	t.Log(tcpAddr)
	t.Log(ipv6only)

	fd, err := unix.Socket(family, unix.SOCK_STREAM|unix.SOCK_NONBLOCK|unix.SOCK_CLOEXEC, unix.IPPROTO_TCP)

	t.Log(fd)
	if err != nil {
		t.Log(err)
	}
	unix.Bind(fd, sa)
	unix.Listen(fd, 1024)
	t.Log("after listen")
	time.Sleep(1e12)
}

func TestGetTCPAddr(t *testing.T) {
	sa, family, tcpAddr, ipv6only, err := GetTCPSockAddr("tcp", "127.0.0.1:8866")
	if err != nil {
		t.Log(err)
	}
	t.Log(sa)
	t.Log(family)
	t.Log(tcpAddr)
	t.Log(ipv6only)
}

func TestNet(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:8866")
	time.Sleep(1e12)
	if err != nil {
		t.Log(err)
	}
	ln.Accept()
	t.Log(ln)
}

package netpoll

import (
	"golang.org/x/sys/unix"
)

func epollWrapped() {
	unix.EpollCreate1()
}

func main() {
	println("main func")
}

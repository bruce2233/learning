package mynetpoll

import (
	"net"
)

func main() {
	net.Listen("tcp", "10086")
}

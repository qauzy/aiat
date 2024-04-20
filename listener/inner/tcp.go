package inner

import (
	"errors"
	"net"

	"github.com/qauzy/aiat/adapter/inbound"
	C "github.com/qauzy/aiat/constant"
)

var tcpIn chan<- C.ConnContext

func New(in chan<- C.ConnContext) {
	tcpIn = in
}

func HandleTcp(address string) (conn net.Conn, err error) {
	if tcpIn == nil {
		return nil, errors.New("tcp uninitialized")
	}
	// executor Parsed
	conn1, conn2 := net.Pipe()
	context := inbound.NewInner(conn2, address)
	tcpIn <- context
	return conn1, nil
}

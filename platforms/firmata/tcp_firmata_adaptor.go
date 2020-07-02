package firmata

import (
	"io"
	"net"
	"time"

	"gobot.io/x/gobot"
)

// TCPAdaptor represents a TCP based connection to a microcontroller running
// WiFiFirmata
type TCPAdaptor struct {
	*Adaptor
}

func connect(address string, timeout time.Duration) (io.ReadWriteCloser, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(time.Second * 30)

	return conn, err
}

// NewTCPAdaptor opens and uses a TCP connection to a microcontroller running
// WiFiFirmata
func NewTCPAdaptor(args ...interface{}) *TCPAdaptor {
	address := args[0].(string)
	timeout := args[1].(time.Duration)

	a := NewAdaptor(address)
	a.SetName(gobot.DefaultName("TCPFirmata"))
	a.PortOpener = func(port string) (io.ReadWriteCloser, error) {
		return connect(port, timeout)
	}

	return &TCPAdaptor{
		Adaptor: a,
	}
}

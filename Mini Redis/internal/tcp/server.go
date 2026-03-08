package tcp

import (
	"fmt"
	"net"
	"redis-clone/internal/handler"
)


type Tcp struct {
	handler handler.Handler
}

func NewHandler(hand handler.Handler) *Tcp{
	return &Tcp{handler: hand}
}



func (t *Tcp) Start(){
	listener, err := net.Listen("tcp", ":8084")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

		fmt.Println("TCP server 8084 portda ishlayapti...")


	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connection error:", err)
			continue
		}

		go t.handleConnection(conn)
	}
}


func (t *Tcp) handleConnection(conn net.Conn){
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
				fmt.Println("read error:", err)
			return
		}

		input := string(buf[:n])
		res := t.handler.Handle(input)
		conn.Write([]byte(res + "\n"))
	}
}

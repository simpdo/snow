package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

type Server struct {
	Net       string
	Port      int
	Timeout   int
	listener  *net.TCPListener
	WaitGroup sync.WaitGroup
	ExitChan  chan struct{}
}

func (this *Server) Start() {
	addr, err := net.ResolveTCPAddr(this.Net, fmt.Sprintln(":%d", this.Port))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	this.listener, err = net.ListenTCP(this.Net, addr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	for {
		this.listener.SetDeadline(time.Now().Add(time.Duration(this.Timeout) * time.Second))
		conn, err := this.listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		conn.Close()
	}
}

func (this *Server) ProcessConn(conn *net.Conn) {

}

func (this *Server) Close() {
	close(this.ExitChan)
	this.WaitGroup.Wait()
}

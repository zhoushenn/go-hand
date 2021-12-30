package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {

	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for true {
		accept, err := listen.Accept()
		log.Println("接受连接:", accept)
		if err != nil {
			log.Print(err)
			continue
		}
		//加go关键字, 让他再自己的goroutine里面运行.
		go handConn(accept)
	}

}

func handConn(c net.Conn) {
	defer c.Close()
	for true {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Println("客户端断开连接:", c)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

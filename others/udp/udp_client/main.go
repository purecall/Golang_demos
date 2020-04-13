package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("connect error", err)
		return
	}
	defer socket.Close()

	// send data
	senddata := []byte("hello server!")
	_, err = socket.Write(senddata)
	if err != nil {
		fmt.Println("send data error", err)
		return
	}

	// receive data
	data := make([]byte, 4096)
	read, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("read data error", err)
		return
	}
	fmt.Println(read, remoteAddr)
	fmt.Printf("%s\n", data)

}

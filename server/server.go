package main

import (
	"fmt"
	"log"
	"net"
)

func handleClient(listener *net.UDPConn){
	buf := make([]byte, 1024)
	n, addr, err := listener.ReadFromUDP(buf)
	if err != nil{
		fmt.Println(err)
	}
	listener.WriteToUDP(append([]byte("Hello, you said "), buf[:n]...), addr)
}

func main(){
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("localhost"), Port: 8080})
	if err != nil{
		log.Fatal(err)
	}
	for{
		handleClient(listener)
	}
}
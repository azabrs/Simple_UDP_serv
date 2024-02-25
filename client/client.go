package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)
func copyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
	   log.Fatal(err)
	}
 }

func main(){
	con, err := net.Dial("udp", "localhost:8080")
	if err != nil{
		fmt.Println(err)
	}
	for{
		go copyTo(con, os.Stdout)
		copyTo(os.Stdin, con)
	}
	 

}
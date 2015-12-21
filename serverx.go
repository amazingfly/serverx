package main

import (
	"fmt"
//	"io"
	"bufio"
//	"os"
	"os/exec"
	"time"
	"net"
)

func Outlet(s string){

	exec.Command(s).Start()

}

func main() {
	fmt.Println("Starting Server")
	ln, err := net.Listen("tcp", ":8484")
	if err != nil{
		fmt.Println("Error while starting to listen")
		fmt.Println(err)
	}
	fmt.Println("Starting infinite loop")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error while accepting connection")
			fmt.Println(err)
		}
		fmt.Println("connection acceepted")

		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("error while reading from buffer")
			fmt.Println(err)
		}
		fmt.Println("message recieved")
		fmt.Println(msg)
		
		fmt.Println("attempting to send message to outlet")
		go Outlet(msg)
		
		/*go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			fmt.Println(c)
			b := make([]byte, 0)
			fmt.Println(c.Read(b))
			fmt.Println(b)
			// Shut down the connection.
			c.Close()
		}(conn)*/
		conn.Close()
	}
	fmt.Println("Attempting outlet 1 off")
	exec.Command("outlet1off").Start()
	time.Sleep(500)
	//exec.Command("outlet1on")
}

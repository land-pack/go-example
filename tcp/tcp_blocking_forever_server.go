package main

import "net"
import "fmt"
import "bufio"
import "strings"

func main() {

	fmt.Println("Launching server ...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// run loop forever (or until Ctrl-C)

	for {
		// accept connection on port
		conn, _ := ln.Accept()
		for {
			message, _ := bufio.NewReader(conn).ReadString('\n')
			// output message received
			fmt.Print("Message Received:", string(message))
			// Sample process string received
			newmessage := strings.ToUpper(message)
			//send new string back to client
			conn.Write([]byte(newmessage + "\n"))
			conn.Close()
			break
		}

	}

}

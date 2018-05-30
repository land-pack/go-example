package main

import "net"
import "fmt"
import "bufio"
import "strings" // Only needed below for sample processing

func main() {
	fmt.Println("Launching server ...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until Ctrl-C)
	for {
		// will listen for messages to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// Sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
        conn.Close()
        break
	}

}

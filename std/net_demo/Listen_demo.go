package main

import "os"
import "net"


func handleConnection(conn net.Conn){
    buffer := make([] byte, 2048)
    for {
        n, err := conn.Read(buffer)
        if err != nil {
            println(err)
            return
        }
        println(string(buffer[:n]))
    }
}


func main() {
    ln, err := net.Listen("tcp", ":8990")
    if err != nil {
        println("Listen error:", err)
        os.Exit(1)
        }

   for {
            conn, err := ln.Accept()
            if err != nil{
                println("Accept error:", err)
                continue
            }
            go handleConnection(conn)

      }
}

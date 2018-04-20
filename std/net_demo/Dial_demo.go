package main

import "net"
import "os"
import "fmt"
import "bufio"

func main(){
    conn, err := net.Dial("tcp", "golang.org:80")

    if err != nil {
        println(err)
        os.Exit(1)
    }

    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    status, err := bufio.NewReader(conn).ReadString('\n')
    println(status)
}

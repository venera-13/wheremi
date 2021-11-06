package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

const (
    connHost = "localhost"
    connPort = "1917"
    connType = "tcp"
)

func main() {
    fmt.Println(timeString() + "Starting " + connType + " server on " + connHost + ":" + connPort)
    l, err := net.Listen(connType, connHost+":"+connPort)
    if err != nil {
        fmt.Println(timeString() + "Error listening:", err.Error())
        os.Exit(1)
    }
    defer l.Close()

    for {
        c, err := l.Accept()
        if err != nil {
            fmt.Println(timeString() + "Error connecting:", err.Error())
            return
        }
	
        fmt.Println(timeString() + "Client " + c.RemoteAddr().String() + " connected.")

        go handleConnection(c)
    }
}

func handleConnection(conn net.Conn) {

    conn.Write([]byte("HTTP/1.0 200 OK\n\n"))
    conn.Write([]byte("Thanks for using the okroshka.net WhereMI service!\n"))
    conn.Write([]byte("Your TCP source address is: " + conn.RemoteAddr().String() + "\n"))
    conn.Write([]byte("Goodbye!\n"))

    conn.Close()
}

func timeString() (timeString string) {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05\t")
}

// ░█████╗░░█████╗░███╗░░██╗████████╗███████╗███╗░░░███╗██████╗░░█████╗░  ███╗░░██╗░█████╗░████████╗███████╗
// ██╔══██╗██╔══██╗████╗░██║╚══██╔══╝██╔════╝████╗░████║██╔══██╗██╔══██╗  ████╗░██║██╔══██╗╚══██╔══╝██╔════╝
// ██║░░╚═╝██║░░██║██╔██╗██║░░░██║░░░█████╗░░██╔████╔██║██████╔╝██║░░██║  ██╔██╗██║██║░░██║░░░██║░░░█████╗░░
// ██║░░██╗██║░░██║██║╚████║░░░██║░░░██╔══╝░░██║╚██╔╝██║██╔═══╝░██║░░██║  ██║╚████║██║░░██║░░░██║░░░██╔══╝░░
// ╚█████╔╝╚█████╔╝██║░╚███║░░░██║░░░███████╗██║░╚═╝░██║██║░░░░░╚█████╔╝  ██║░╚███║╚█████╔╝░░░██║░░░███████╗
// ░╚════╝░░╚════╝░╚═╝░░╚══╝░░░╚═╝░░░╚══════╝╚═╝░░░░░╚═╝╚═╝░░░░░░╚════╝░  ╚═╝░░╚══╝░╚════╝░░░░╚═╝░░░╚══════╝
//
// ██████╗░░█████╗░░██████╗░██████╗██╗███╗░░██╗░██████╗░
// ██╔══██╗██╔══██╗██╔════╝██╔════╝██║████╗░██║██╔════╝░
// ██████╔╝███████║╚█████╗░╚█████╗░██║██╔██╗██║██║░░██╗░
// ██╔═══╝░██╔══██║░╚═══██╗░╚═══██╗██║██║╚████║██║░░╚██╗
// ██║░░░░░██║░░██║██████╔╝██████╔╝██║██║░╚███║╚██████╔╝
// ╚═╝░░░░░╚═╝░░╚═╝╚═════╝░╚═════╝░╚═╝╚═╝░░╚══╝░╚═════╝░
// this logo is redundant, but still necessary at the same time.
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var (
	CONN_HOST   string     //hostname
	CONN_PORT   string     //port
	CONN_TYPE   string     //connection type; idk why i did this choose tcp always
	connections []net.Conn //a list of clients connected to the server; used for broadcasting one message to all clients
)

func init() { // i didn't want to use fmt.Scanln, but i had to because when i did io.Reader, it would add a redundant new line. why did it do that?
	fmt.Print("Enter hostname:\n>>>")
	fmt.Scanln(&CONN_HOST)

	fmt.Print("Enter port:\n>>>")
	fmt.Scanln(&CONN_PORT)

	fmt.Print("Enter connection type:\n>>>")
	fmt.Scanln(&CONN_TYPE)
}

func main() {
	fmt.Println("Starting " + CONN_TYPE + " server on " + CONN_HOST + ":" + CONN_PORT)
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		connections = append(connections, c)

		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	broadcast(conn, buffer)

	if err != nil {
		fmt.Println("Client left.")
		for i := range connections {
			if connections[i] == conn {
				break
			}
		}
		conn.Close()
		return
	}

	handleConnection(conn)
}

func broadcast(conn net.Conn, msg []byte) {
	for i := range connections {
		if connections[i] != conn {
			connections[i].Write(msg)
			fmt.Println(connections)
		}
	}
}

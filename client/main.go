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
	CONN_HOST string //hostname
	CONN_PORT string //port
	CONN_TYPE string //connection type; idk why i did this choose tcp always
)

func init() { // i didn't want to use fmt.Scanln, but i had to because when i did io.Reader, it would add a redundant new line. why did it do that?
	fmt.Print("Enter hostname:\n>>>")
	fmt.Scanln(&CONN_HOST)

	fmt.Print("Enter port:\n>>>")
	fmt.Scanln(&CONN_PORT)

	fmt.Print("Enter connection type:\n>>>")
	fmt.Scanln(&CONN_TYPE)
}

func main() { //calls useful functions
	fmt.Println("Connecting to", CONN_TYPE, "server", CONN_HOST+":"+CONN_PORT)
	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	go readMessage(conn)
	writeMessage(conn)
}

func readMessage(conn net.Conn) { //reads text from server
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print("Server relay: " + string(message))
	}
}

func writeMessage(conn net.Conn) { //writes text and sends to server
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Text to send: ")

		input, _ := reader.ReadString('\n')

		conn.Write([]byte(input))
	}
}

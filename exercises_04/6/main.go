// # Building upon the code from the previous exercise:
//
// We are now going to get "I see you connected" to be written.
//
// When we used bufio.NewScanner(), our code was reading from an io.Reader that never ended.
//
// We will now break out of the reading.
//
// Package bufio has the Scanner type. The Scanner type "provides a convenient interface for reading data". When you have a Scanner type, you can call the SCAN method on it. Successive calls to the Scan method will step through the tokens (piece of data). The default token is a line. The Scanner type also has a TEXT method. When you call this method, you will be given the text from the current token. Here is how you will use it:
//
// scanner := bufio.NewScanner(conn)
// for scanner.Scan() {
// 	ln := scanner.Text()
// 	fmt.Println(ln)
// }
//
// Use this code to READ from an incoming connection and print the incoming text to standard out (the terminal).
//
// When your "ln" line of text is equal to an empty string, break out of the loop.
//
// Run your code and go to localhost:8080 in **your browser.**
//
// What do you find?
// _______________________________________________________________
// # Building upon the code from the previous problem:
//
// Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".
//
// Pass the connection of type net.Conn as an argument into this function.
//
// Add "go" in front of the call to "serve" to enable concurrency and multiple connections.
// _______________________________________________________________
// Add code to WRITE to the connection.
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}

func handle(c net.Conn) {
	defer c.Close()
	read(c)
	write(c)
}

func read(c net.Conn) {
	s := bufio.NewScanner(c)
	for s.Scan() {
		t := s.Text()
		fmt.Println(t)
		if t == "" {
			break
		}
	}
}

func write(c net.Conn) {
	io.WriteString(c, "I see you connected")
}

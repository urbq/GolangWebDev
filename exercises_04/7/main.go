// # Building upon the code from the previous problem:
//
// Before we WRITE our RESPONSE , let's WRITE to our RESPONSE the STATUS LINE and some REPONSE HEADERS. Remember the request line and status line:
//
// REQUEST LINE
// GET / HTTP/1.1
// method SP request-target SP HTTP-version CRLF
// https://tools.ietf.org/html/rfc7230#section-3.1.1
//
// RESPONSE (STATUS) LINE
// HTTP/1.1 302 Found
// HTTP-version SP status-code SP reason-phrase CRLF
// https://tools.ietf.org/html/rfc7230#section-3.1.2
//
// Write the following strings to the response - use io.WriteString for all of the following except the second and third:
//
// "HTTP/1.1 200 OK\r\n"
//
// fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
//
// fmt.Fprint(c, "Content-Type: text/plain\r\n")
//
// "\r\n"
//
// Look in your browser "developer tools" under the network tab. Compare the RESPONSE HEADERS from the previous file with the RESPONSE HEADERS in your new solution.
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
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len("I see you connected"))
	fmt.Fprint(c, "Content-Tyoe: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, "I see you connected")
}

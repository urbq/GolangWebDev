// # Building upon the code from the previous problem:
//
// Add code to respond to the following METHODS & ROUTES:
// 	GET /
// 	GET /apply
// 	POST /apply

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
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

	m, u := read(c)
	write(c, m, u)
}

func read(c net.Conn) (string, string) {
	var method, uri string
	var i int
	s := bufio.NewScanner(c)
	for s.Scan() {
		t := s.Text()
		fmt.Println(t)
		if t == "" {
			break
		}
		if i == 0 {
			firstLine := strings.Fields(t)
			method = firstLine[0]
			uri = firstLine[1]
		}
		i++
	}
	return method, uri
}

func write(c net.Conn, method, uri string) {
	//body := fmt.Sprint(method, " ", uri)
	var body string
	switch uri {
	case "/":
		body += defaultGet()
	case "/apply":
		switch method {
		case "GET":
			body += applyGet()
		case "POST":
			body += applyPost()
		}
	}
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Tyoe: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func defaultGet() string {
	return `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
		</head>
		<body>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			<form action="/apply" method="POST">
			<input type="hidden" value="In my good death">
			<input type="submit" value="submit">
			</form>
		</body>
		</html>
	`
}

func applyGet() string {
	return "applyGet"
}

func applyPost() string {
	return "applyPost"
}

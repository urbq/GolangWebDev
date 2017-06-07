package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

//create a multiplexer (mux, servemux, router, server, http mux, ...) so that your server responds to different URIs and METHODS.
func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	url := request(conn)

	// write response
	respond(conn, url)
}

func request(conn net.Conn) string {
	var url string
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD", m)
			url = strings.Fields(ln)[1]
			fmt.Println("***URL", url)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	return url
}

func respond(conn net.Conn, url string) {

	body := createBody(url)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func createBody(url string) string {
	var b string
	if url == "/" {
		b = `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
	} else if url == "/test" {
		b = `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Test</strong></body></html>`
	} else {
		b = `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Default</strong></body></html>`
	}
	return b
}

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	method := strings.Fields(ln)[0]
	url := strings.Fields(ln)[1]
	fmt.Println("***METHOD", method)
	fmt.Println("***URI", url)

	if method == "GET" && url == "/" {
		index(conn)
	}
	if method == "GET" && url == "/about" {
		about(conn)
	}
	if method == "GET" && url == "/contact" {
		contact(conn)
	}
	if method == "GET" && url == "/apply" {
		apply(conn)
	}
	if method == "POST" && url == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {
	html := readHTML("index.gohtml")
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, html)
}

func about(conn net.Conn) {
	html := readHTML("about.gohtml")
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, html)
}

func contact(conn net.Conn) {
	html := readHTML("contact.gohtml")
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, html)
}

func apply(conn net.Conn) {
	html := readHTML("apply.gohtml")
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, html)
}

func applyProcess(conn net.Conn) {
	html := readHTML("apply_proccess.gohtml")
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, html)
}

func readHTML(file string) string {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err.Error())
	}
	html := string(body)
	return html
}

// telnet localhost 8080

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	fmt.Println("Listening on port :8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handleRequest(conn)
		time.Sleep(time.Second)
	}
}

func handleRequest(conn net.Conn) {
	fmt.Println(" --- Beggining of request ---")

	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	urlResult := ""
	temp := scanner.Scan()
	if temp {
		line := scanner.Text()
		/*fmt.Println("line: ", line)*/
		f := strings.Fields(line)
		if len(f) > 0 {
			fmt.Println(f)
			switch f[0] {
			case "GET":
				urlResult = "http://localhost:8080" + f[1]
			default:
				urlResult = "This method is not suported. "
			}
		}
	} else {
		fmt.Println("The request couldn't be read")
	}

	fmt.Println(urlResult)
	bodyResponse := fmt.Sprintf(`<html>
			<head>
				<title>Get the URL</title>
			</head>
			<body>
				<h3>Hello user:</h3>
				<p>					
					this is your URL: <b>%v</b>
				</p>
			</body>
		</html>`, urlResult)
	fmt.Println(bodyResponse)
	fmt.Println(" --- End of request ---")

	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintln(conn, "Content-Length: ", len(bodyResponse))
	fmt.Fprintln(conn, "Content-Type: text/html")
	fmt.Fprintln(conn, "")
	fmt.Fprintln(conn, bodyResponse)
	/*
		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	*/
}

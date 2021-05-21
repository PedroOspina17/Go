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
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("-------- Listening on port 8080  ")
	fmt.Println(" ")

	for {

		conn, err := li.Accept()
		fmt.Println("------- A new connection was accepted")
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handleRequest(conn)
		time.Sleep(time.Second)
	}

}

func handleRequest(conn net.Conn) {
	fmt.Println("---- Starting to process the request")

	defer conn.Close()
	defer fmt.Print("------- The connection was closed. \n\n")

	scanner := bufio.NewScanner(conn)

	response := ""
	if s := scanner.Scan(); s {
		lineWords := strings.Fields(scanner.Text())
		if len(lineWords) >= 2 {
			httpVerb := lineWords[0]
			urlRequest := "http://localhost:8080" + lineWords[1]
			responseMessage := ""
			fmt.Println("--- Executing the method: ", httpVerb)
			fmt.Println("--- The request url is: ", urlRequest)
			fmt.Println("--- Header information: ", lineWords)
			switch httpVerb {
			case "GET":
				responseMessage = "Congrats, Your GET was processed successfully !! <br> <br> ##Returning some data here##"
			case "POST":
				responseMessage = "Congrats, Your POST was processed successfully !! <br> <br>##Saving information...##"
			default:
				responseMessage = "Warning, Your request method was not recognized nor supported. "
			}
			fmt.Println("--- The response message was: ", responseMessage)

			response = fmt.Sprintf(`	
			<html>
				<head>
					<title>My first MUX</title>
				</head>
				<body>
					<h3>Hello user:</h3>
					Your request was a %v, with the following URL: %v. 
					Response
					<p>					
						<b>%v</b>
					</p>
					
					<a href="file:///Users/pedro.ospina/Documents/Repos/Go/golang-web-dev/016_building-a-tcp-server-for-http/04_hands-on/index.html">Go back</a>
				</body>
			</html>
			`, httpVerb, urlRequest, responseMessage)
		} else {
			response = fmt.Sprintf(`	
			<html>
				<head>
					<title>My first MUX</title>
				</head>
				<body>
					<h3>Hello user:</h3>
					The request couldn't be processed.
					<a href="http://Users/pedro.ospina/Documents/Repos/Go/golang-web-dev/016_building-a-tcp-server-for-http/04_hands-on/index.html">Go back</a>
				</body>
			</html>
			`)
		}

	}

	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintln(conn, "Content-Length: ", len(response))
	fmt.Fprintln(conn, "Content-Type: text/html")
	fmt.Fprintln(conn, "")
	fmt.Fprintln(conn, response)

	fmt.Println("---- Ending the request processing.  ")

}

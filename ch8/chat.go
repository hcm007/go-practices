package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
)

func main(){

	listener, err := net.Listen("tcp","localhost:8000")

	if err != nil{
		log.Fatal(err)
	}

	for{
		 conn, err := listener.Accept()
		 if err != nil{
		 	log.Print(err)
		 	continue
		 }
		 go handleConn(conn)
	}
}

type client chan <- string

var (
	entering = make(chan client)
	leaving =make(chan client)
	messages=make(chan string)
)

func broadcaster(){
	clients := make(map[client]bool) // all connetcted clients
	for {
		select {
		case msg := <- messages:
			fmt.Println("this is arrived")
			for cli := range clients{
				cli <- msg
			}

		case cli := <-entering:
			fmt.Println("this is started")
			clients[cli] = true

		case cli := <- leaving:
			fmt.Println("this is leaving")
			delete(clients,cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn){

	ch := make(chan string)
	go clientWriter(conn,ch)

	who := conn.RemoteAddr().String()
	ch <- "you are " + who
	messages <- who + "has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan(){
		messages <- who + ":" + input.Text()
	}
	leaving <- ch
	messages <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string){

	for msg := range ch{
		fmt.Fprintln(conn,msg)
	}

}








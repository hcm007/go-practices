package main

import (
	"io"
	"log"
	"net"
	"os"
	"fmt"
)


func main(){
	conn, err := net.Dial("tcp", "localhost:8000")
	//defer conn.Close()
	if err != nil{
		log.Fatal(err)
	}
	done := make(chan int)
	go func(){
		_,err:=io.Copy(os.Stdout, conn)
		if err!= nil{
			log.Println(err)
		}
		log.Println("done")
		done <- 10
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	intdown := <- done
	fmt.Print(intdown)

}

func mustCopy(dis io.Writer, src io.Reader){

	if _, err := io.Copy(dis,src); err != nil{
		log.Fatal(err)
	}

}
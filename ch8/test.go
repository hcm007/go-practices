package main

import (
	//"fmt"
	//"golang.org/x/net/html/atom"
	"fmt"
	"log"
	"./links"
)



func crawl(url string) []string {
fmt.Println(url)
list, err := links.Extract(url)
if err != nil {
log.Print(err)
}
return list
}


func pinger(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	 c := make(chan string)
	//go func() {c <- 1}()

	<- c


	}
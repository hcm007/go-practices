package main

import (
	"./links"
	"fmt"
	"log"
	//"os"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main(){

	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func(){worklist <- os.Args[1:]}()

	//<- unseenLinks

	for i :=0; i<20;i++{
		go func() {
			for link := range unseenLinks{
				foundLinks := crawl(link)
				go func() {worklist <- foundLinks}()
			}
		}()
	}

	seen := make(map[string] bool)

	for list := range worklist{
		for _,link := range list{
			if !seen[link]{
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

}
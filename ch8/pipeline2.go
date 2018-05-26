package main

import "fmt"

func main(){

	natuals := make(chan int)
	squares := make(chan int)

	go func(){
		for x :=0; x < 100; x++{
			natuals <- x
		}
		close(natuals)
	}()

	go func(){
		for x := range natuals{
			squares <- x*x
		}
		close(squares)
	}()


	for x := range squares{
		fmt.Println(x)
	}

}

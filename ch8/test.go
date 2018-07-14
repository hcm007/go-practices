package main

import "fmt"

func main(){
	a := make(map[string]int)
	for x := range a{
		x = "woaini"
		fmt.Println(x)
	}

}

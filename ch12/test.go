package main

import (
	"reflect"
	"fmt"
)

func main() {
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%d \n",v)
	fmt.Println(3)

	slice := []int{3,8}
	fmt.Println(len(slice))
}

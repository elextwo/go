package main

import (
	"fmt"
)

func main(){
	//var a [10]int
	for i := 0;i< 1000;i++ {
		go func(i  int) {
			fmt.Println(i)
		}(i)
	}
}

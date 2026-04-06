package main

import (
	"fmt"
	"net/http"
	"time"

	_ "net/http/pprof"
)


func slowFunction() {
	sum := 0 
	for i := 0; i < 1e7; i++ {
		sum += i
	}
  fmt.Println(sum)
}

func fastFunction() {
	sum := 0 
	for i := 0; i < 1e5; i++ {
		sum += i
	}
  fmt.Println(sum)
}

func main(){

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	for {
		slowFunction()
		fastFunction()
		time.Sleep( 1* time.Second)
	}
}
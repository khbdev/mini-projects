package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)



func main(){
	mux :=  http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server ishlayabdi")
	})
	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request Keldi")
		time.Sleep(8 * time.Second)
		log.Println("Request tugadi")
		fmt.Fprintln(w, "Request tugadi")
	})

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr: ":8085",
		Handler: mux,
	}

	go func() {
		fmt.Println("Server Ishga tushdi: 8085")
		server.ListenAndServe()
	}()

	sig := <-sigCh
    fmt.Printf("habar keldi: %d", sig)

	fmt.Println("ShutDown ishga tushdi")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    
	defer cancel()

	server.Shutdown(ctx)
	server.Close()

	fmt.Println("Server ochdi")

}
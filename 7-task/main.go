package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)



func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application-json")
		json.NewEncoder(w).Encode("Salom")
	})

	mux.HandleFunc("/aziz", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request Keldi")
		time.Sleep(8 * time.Second)
		log.Println("Request tugadi")
		fmt.Fprintln(w, "Request tugadi")
	})

	

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)


server := &http.Server{
		Addr: ":8084",
		Handler: mux,
	}

	go func() {
			fmt.Println("Server Ishga tushdi: 8084")
			server.ListenAndServe()
	}()


	sig := <-sigCh
	fmt.Println("Signal keldi", sig)

	fmt.Println("ShutDown ishga tushdi")
  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  server.Shutdown(ctx)
  server.Close()

  fmt.Println("Server tugadi")
}
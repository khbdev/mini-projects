package main

import (
	"fmt"

	"net/http"
	"time"
)


func loggerMiddlewar(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		method := r.Method
		path  := r.URL.Path

		next.ServeHTTP(w, r)

		latency := time.Since(start)

		fmt.Printf("Method %s, Path %s, latency %s\n", method, path, latency)
	})
}


func PanicMiddleware(next http.Handler) http.Handler {
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func ()  {
		if err := recover(); err != nil {
			fmt.Println("Panic  Ushlab qoldindi ")
			http.Error(w, "Invalid server error", http.StatusInternalServerError)	
		}	
		}()
		next.ServeHTTP(w, r)
	})
}


func ApiKeyMiddleware(next http.Handler) http.Handler {
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Form.Get("X-API-Key")
		if apiKey != "secret123" {
			http.Error(w, "token yoq", http.StatusUnauthorized)
			return 
		}
		next.ServeHTTP(w, r)
	})
}



func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello!")
}



func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("Xatolik")
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", helloHandler)
  mux.HandleFunc("/panic", PanicHandler)


  log := loggerMiddlewar(PanicMiddleware(mux))
    http.ListenAndServe(":8081", log)
}
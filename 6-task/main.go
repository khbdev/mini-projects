package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Data struct {
	count int
	time time.Time
	
}


var datas sync.Map


func RateLimitingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return  func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		val, ok := datas.Load(ip)
		if !ok {
			datas.Store(ip, &Data{count: 1, time: time.Now()})
			next(w, r)
			return 
		}

		v := val.(*Data)


		if time.Since(v.time) > time.Minute {
			v.time = time.Now()
			v.count = 1
		} else {
			v.count++
		}
		if v.count > 10 {
			http.Error(w, "Too Many Request", http.StatusTooManyRequests)
			return 
		}

		next(w, r)
	}
}
	



func HelloHandler(w http.ResponseWriter, r *http.Request) {

	value := "Hello World"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(value)

}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)

	log := RateLimitingMiddleware(mux.ServeHTTP)

	fmt.Println("Server Running Port :8083")
	http.ListenAndServe(":8083", log)
}

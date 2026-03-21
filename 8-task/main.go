package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)


type  CreateUserRequest struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

type FailedError struct {
	Failed string `json:"failed"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Errors []FailedError `json:"errors"`
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/users", CreateUserHandler)
	
	fmt.Println("Server Running, Port :8082")
	http.ListenAndServe(":8082", mux)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}


	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Errors: []FailedError{
				{Failed: "Json", Message: "Invalid json"},
			},
		})
	}

	var errors []FailedError

   if strings.TrimSpace(req.Name) == "" {
	errors = append(errors, FailedError{
		Failed: "Name",
	   Message: "Name is reqired",
	})
   } else if len(req.Name) > 50 {
		errors = append(errors, FailedError{
		Failed: "Name",
	   Message: "Name 50 max",
	})
   }

   if req.Age < 18 {
	errors = append(errors, FailedError{
		Failed: "Age",
	   Message: "Age 18+",
	})

}
	if !strings.Contains(req.Email, "@") || !strings.Contains(req.Email, ".") {
		errors = append(errors, FailedError{
		Failed: "Email",
	   Message: "Email is Not Format",
	})
	}

	if len(errors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Errors: errors,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string {
		"Message": "Create User Succesfull",
	})
}
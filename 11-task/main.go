package main

import (
 _ "net/http/pprof"
	"net/http"
)

type User struct {
	ID int
	Name string
	Email string
}

	
var users []*User





func main() {
    go func() {
     http.ListenAndServe("localhost:6060", nil)
    }()

    // Sample actions
    createUser(1, "Ali", "ali@example.com")
    createUser(2, "Vali", "vali@example.com")
    updateUser(1, "AliUpdated", "ali2@example.com")
    deleteUser(2)

    select {} // programni o‘chirmay turish
}
func createUser(id int, name, email string) {
	user := &User{ID: id, Name: name, Email: email}
	users  = append(users, user)
}

func updateUser(id int, name, email string) {
	for _, u := range users {
		if u.ID == id {
			u.Name = name
			u.Email = email
		}
	}
}


func deleteUser(id int) {
    for i, u := range users {
        if u.ID == id {
            users = append(users[:i], users[i+1:]...)
            break
        }
    }
}




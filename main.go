package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Class struct {
	ID      string  `json:"id"`
	Student string  `json:"student"`
	Mark    float64 `json:"mark"`
	Teacher string  `json:"teacher"`
}

var class = []Class{
	{ID: "1", Student: "Nikita", Mark: 3.4, Teacher: "Elena Gavlitskaya"},
	{ID: "2", Student: "Artem", Mark: 2.1, Teacher: "Oleg Slushniy"},
	{ID: "3", Student: "Diana", Mark: 4.4, Teacher: "Elena Gavlitskaya"},
	{ID: "4", Student: "Egor", Mark: 4.9, Teacher: "Elena Gavlitskaya"},
	{ID: "5", Student: "Dasha", Mark: 1.9, Teacher: "Oleg Slushniy"},
	{ID: "6", Student: "Igor", Mark: 4.2, Teacher: "Elena Gavlitskaya"},
	{ID: "7", Student: "Damir", Mark: 3.0, Teacher: "Oleg Slushniy"},
	{ID: "8", Student: "Andrey", Mark: 3.3, Teacher: "Elena Gavlitskaya"},
	{ID: "9", Student: "Lina", Mark: 4.5, Teacher: "Oleg Slushniy"},
	{ID: "10", Student: "Sasha", Mark: 3.9, Teacher: "Oleg Slushniy"},
}

func main() {
	r := mux.NewRouter()

	r.Handle("/student/{id}/{teacher}", auth(http.HandlerFunc(GetStudents)))

	fmt.Println("Server is starting at port 8080")
	http.ListenAndServe(":8080", r)
}

type User struct {
	UserName     string
	UserPassword string
}

var AdminUser1 = User{
	UserName:     "ElenaGavlitskaya",
	UserPassword: "admin1",
}

var AdminUser2 = User{
	UserName:     "OlegSlushniy",
	UserPassword: "admin2",
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if ok && (username == AdminUser2.UserName && password == AdminUser2.UserPassword) {
			// Oleg Slushniy
			next.ServeHTTP(w, r)
			return
		}

		if ok && (username == AdminUser1.UserName && password == AdminUser1.UserPassword) {
			// Elena Gavlitskaya
			params := mux.Vars(r)
			teacher := params["teacher"]

			if teacher == "Elena Gavlitskaya" {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusForbidden)
	})
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	teacher := params["teacher"]

	for _, student := range class {
		if student.ID == id && student.Teacher == teacher {
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
}


package main

import (
	controllers "dbtest/db_controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//route := mux.NewRouter()
	//r := route.PathPrefix("/user").Subrouter() //Base Path
	//Routes
	r := mux.NewRouter()
	r.HandleFunc("/delete", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/update/{id}", controllers.UpdateDetails).Methods("PUT")
	r.HandleFunc("/add", controllers.AddUser).Methods("POST")
	r.HandleFunc("/", controllers.GetAllUser).Methods("GET")
	r.HandleFunc("/{id}", controllers.GetParticularUser).Methods("GET")
	r.Use(mux.CORSMethodMiddleware(r))
	http.ListenAndServe(":8000", r)
}

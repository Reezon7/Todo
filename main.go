package main

import (

	"net/http"
	"todo/database"
)

func main(){
	var db database.MyDb = database.MyDb{}
	defer db.CloseDb()
	db.ConnectDb()

	http.HandleFunc("/", db.CreateTask)
	http.HandleFunc("/allTasks", db.ViewAllTasks)

	http.ListenAndServe(":8080", nil)
}
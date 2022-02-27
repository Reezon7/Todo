package database

import (
	"encoding/json"
	"net/http"
	"todo/entity"
	"io/ioutil"

)


func (db *MyDb) CreateTask(w http.ResponseWriter, r *http.Request){
	var task entity.Task
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &task)
	if err !=nil{
		http.Error(w, err.Error(), 500)
	}
	db.AddTask(task)
	w.WriteHeader(http.StatusOK)
}

func (db *MyDb) ViewAllTasks(w http.ResponseWriter, r *http.Request){
	output, err := json.Marshal(db.GetAllTasks())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
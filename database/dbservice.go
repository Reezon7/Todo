package database

import (
	"database/sql"
	"fmt"
	"time"
	"log"
	"todo/entity"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type MyDb sql.DB

const FormatTime = "2006-01-02"

func (db *MyDb) ConnectDb () bool{
	connectionOptions := "user=postgres password=alonesomewhale dbname=apptasks sslmode=disable"
	newDb, err := sql.Open("pgx", connectionOptions)
	
	if err != nil{
		return false
	}
	
	*db = (MyDb)(*newDb)
	newDb.Close()
	return true
}

func (db *MyDb) AddTask(task entity.Task) bool{

	end := "'"+task.EndTask.Format(FormatTime)+"'"
	if task.EndTask.IsZero(){
		end = "-infinity"
	}

	query := fmt.Sprintf("INSERT INTO taskrepo (name, description, \"start\", \"end\", active) VALUES('%s', '%s', '%s', '%s', %t);",
																							task.Name,
																							task.Description, 
																							task.StartTask.Format(FormatTime),
																							end, 
																							task.Active)
	fmt.Println(query)
	_, err := (*sql.DB)(db).Exec(query)
	if err!=nil{
		fmt.Println(err)
		return false
	}else{
		return true
	}
}

func (db *MyDb) GetAllTasks() ([]entity.Task){
	rows, err := (*sql.DB)(db).Query("SELECT * FROM taskrepo")
	if err != nil{
		return nil
	}

	var id int
	var resultListOfTasks []entity.Task = []entity.Task{}
	var name, description, startTask, endTask string
	var active bool

	for rows.Next(){
		t := entity.Task{}
		localError := rows.Scan(&id, &name,  &description, &startTask, &endTask, &active)
		startTask = startTask[:10]
		if localError == nil{
			t.Id = id
			t.Name = name
			t.Description = description
			t.StartTask, _ = time.Parse(FormatTime, startTask)
			if endTask == "-infinity"{
				t.EndTask, _ = time.Parse(FormatTime, "01/01/0001")
			}else{
				t.EndTask, _ = time.Parse(FormatTime, endTask)
			}
			t.Active = active
			resultListOfTasks = append(resultListOfTasks, t)
		}else{
			log.Fatal(err)
		}
	}
	return resultListOfTasks
}

func (db *MyDb) DeleteTask (task entity.Task) bool{
	queryDelete := fmt.Sprintf("DELETE FROM taskrepo WHERE id = %d ", task.Id)
	_, err := (*sql.DB)(db).Exec(queryDelete)
	return err == nil
}

func (db *MyDb) DeleteDoneTasks() bool{
	query := "DELETE FROM taskrepo WHERE active = false"
	_, err := (*sql.DB)(db).Exec(query)
	return err == nil
}

func (db *MyDb) CloseDb(){
	(*sql.DB)(db).Close();
}

package entity
import(
	"time"
)

type Task struct{
	Id int `json:"id,string"`
	Name string `json:"name"`
	Description string `json:"description"`
	StartTask time.Time `json:"starttask"`
	EndTask time.Time `json:"endtask"`
	Active bool `json:"active,string"`
}
package tasks

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

var Tasks = []Task{}
var NextID = 1



//http://localhost:8080/tasks
//http://localhost:8080/
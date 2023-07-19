package request

import (
	"studylesson/postgres"
)

type GetTask struct {
	ID string `param:"id" validated:"required"`
}

type ListTask struct {
}

type CreateTask struct {
	Title string `json:"title validated:required`
}

func ToCreateTask() *postgres.Task {
	createTask := &postgres.Task{}
	return createTask
}

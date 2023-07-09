package response

import (
	"studylesson/postgres"
)

type TaskList struct {
	TaskList []*Task `json:"task_list"`
}
type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToTaskList(slice postgres.TaskList) []*Task {
	const layout2 = "2006/01/02 15:04:05"
	res := make([]*Task, len(slice))
	for x, v := range slice {
		responseGet := &Task{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt.Format(layout2),
			UpdatedAt: v.UpdatedAt.Format(layout2),
		}
		res[x] = responseGet
	}
	return res
}

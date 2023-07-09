package server

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"studylesson/postgres"
	"studylesson/request"
	"studylesson/response"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Get(c echo.Context) error {
	//バインドする
	var req request.GetTask
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	//バリデートする
	err = c.Validate(&req)
	if err != nil {
		return err
	}
	//db, err := sql.Open("postgres", "user=yamada-hi password='' host=localhost port=5432 dbname=lesson sslmode=disable")
	//if err != nil {
	//	return err
	//}
	//defer db.Close()
	//if err := db.Ping(); err != nil {
	//	return err
	//}
	var i postgres.Task
	db := postgres.ConnectDB()
	defer db.Close()
	//err = db.QueryRow("SELECT * FROM task WHERE id = $1", req.ID).
	//	Scan(&i.ID, &i.Title, &i.CreatedAt, &i.UpdatedAt)
	err = db.QueryRow("SELECT * FROM task WHERE id = $1", req.ID).
		Scan(&i.ID, &i.Title, &i.CreatedAt, &i.UpdatedAt)
	if err != nil {
		return err
	}

	//responseの形式に変える
	//createdAt,UpdatedAtを変える
	const layout2 = "2006/01/02 15:04:05"
	responseGet := &response.Task{
		ID:        i.ID,
		Title:     i.Title,
		CreatedAt: i.CreatedAt.Format(layout2),
		UpdatedAt: i.UpdatedAt.Format(layout2),
	}
	return c.JSON(http.StatusOK, responseGet)
}

func List(c echo.Context) error {
	//ctx := c.Request().Context()

	//var i postgres.Task
	//var j postgres.TaskList
	//db := postgres.ConnectDB()
	//defer db.Close()
	//
	//rows, err := db.Query("SELECT * FROM task")
	//if err != nil {
	//	return err
	//}
	//for rows.Next() {
	//	rows.Scan(&i.ID, &i.Title, &i.UpdatedAt, &i.CreatedAt)
	//	j = append(j, i)
	//}

	//res := make([]*response.Task, len(j))
	//const layout2 = "2006/01/02 15:04:05"
	//for x, v := range j {
	//	responseGet := &response.Task{
	//		ID:        v.ID,
	//		Title:     v.Title,
	//		CreatedAt: v.CreatedAt.Format(layout2),
	//		UpdatedAt: v.UpdatedAt.Format(layout2),
	//	}
	//	res[x] = responseGet
	//}

	listTask, err := postgres.ListTask()
	if err != nil {
		return err
	}
	var list response.TaskList
	list.TaskList = response.ToTaskList(listTask)
	return c.JSON(http.StatusOK, list)
}

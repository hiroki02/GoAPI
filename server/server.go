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
	var i postgres.Task
	db := postgres.ConnectDB()
	defer db.Close()
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

	listTask, err := postgres.ListTask()
	if err != nil {
		return err
	}
	var list response.TaskList
	list.TaskList = response.ToTaskList(listTask)
	return c.JSON(http.StatusOK, list)
}

func Update(c echo.Context) error {
	var req request.CreateTask
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	err = c.Validate(&req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, req)
}

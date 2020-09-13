package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

/* Models */

// User ユーザ
type User struct {
	gorm.Model
	Name string
	Email string
	Tasks []Task
}
// Project プロジェクト
type Project struct {
	gorm.Model
	Name string
	Description string
	StartDate string
	EndDate string
	Tasks []Task
}
// Task タスク
type Task struct {
	gorm.Model
	Title string
	Description string
	Status string
	StartDate string
	EndDate string
	ProjectID uint
	UserID uint
}

// HealthResponse Response - ヘルスチェック
type HealthResponse struct {
	Status int `json:"status" xml:"status"`
}

func health(c echo.Context) error {
	res := &HealthResponse{
		Status: http.StatusOK,
	}
	return c.JSON(http.StatusOK, res)
}

func main() {
	e := echo.New()
	e.GET("/health", health)
	e.Logger.Fatal(e.Start(":5000"))
}
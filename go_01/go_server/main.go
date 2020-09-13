package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

// Constant
const (
	dbName = "datas.db"
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


func notImplemented(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, map[string]string{"message": "NotImplemented"})
}

func health(c echo.Context) error {
	res := &HealthResponse{
		Status: http.StatusOK,
	}
	return c.JSON(http.StatusOK, res)
}

func findUsers(c echo.Context) error {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	var users []User
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}
func findUser(c echo.Context) error {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	var user User
	db.First(&user, id)
	return c.JSON(http.StatusOK, user)
}
func createUser(c echo.Context) error {
	type Request struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}
	request := new(Request)
	if err := c.Bind(request); err != nil {
	   return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Validate(request); err != nil {
	   return err
	}
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	user := User{Name: request.Name, Email: request.Email}
	db.Create(&user)
	return c.JSON(http.StatusOK, user)
}

func initialize() {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database (Intialize)")
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Task{})
}
func main() {
	initialize()
	e := echo.New()
	e.GET("/health", health)
	e.GET("/users", findUsers)
	e.GET("/user/:id", findUser)
	e.POST("/user/create", createUser)
	e.POST("/user/update", notImplemented)
	e.POST("/user/delete", notImplemented)
	e.Logger.Fatal(e.Start(":5000"))
}
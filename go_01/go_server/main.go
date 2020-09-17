package main

import (
	"net/http"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

const (
	dbUser="user"
	dbPassword="user"
	dbHost="db"
	dbPort=3306
	dbName="master"
)

var db *gorm.DB

/* Models */

// Validator リクエストバリデーター
type Validator struct {
    validator *validator.Validate
}
// Validate バリデート
func (v *Validator) Validate(i interface{}) error {
    return v.validator.Struct(i)
}

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

// API
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
	var users []User
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}
func findUser(c echo.Context) error {
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
	user := User{Name: request.Name, Email: request.Email}
	db.Create(&user)
	return c.JSON(http.StatusOK, user)
}
func updateUser(c echo.Context) error {
	type Request struct {
		ID int `json:"id" validate:"required"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	request := new(Request)
	if err := c.Bind(request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	var user User
	result := db.First(&user, request.ID)
	if result.Error != nil {
		return result.Error
	}
	if len(request.Name) > 0 {
		user.Name = request.Name
	}
	if len(request.Email) > 0 {
		// Temp
		if err := validator.New().Var(request.Email, "email"); err != nil {
			return err
		}
		user.Email = request.Email
	}
	db.Save(&user)
	return c.JSON(http.StatusOK, user)
}
func deleteUser(c echo.Context) error {
	type Request struct {
		ID int `json:"id" validate:"required"`
	}

	request := new(Request)
	if err := c.Bind(request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if result := db.Delete(&User{}, request.ID); result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, request)
}
func findProjects(c echo.Context) error {
	var projects []Project
	db.Find(&projects)
	return c.JSON(http.StatusOK, projects)
}
func findProject(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	var project Project
	db.First(&project, id)
	return c.JSON(http.StatusOK, project)
}
func createProject(c echo.Context) error {
	type Request struct {
		Name string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		StartDate string `json:"start_date" validate:"required"` 
		EndDate string `json:"end_date" validate:"required"`
	}
	request := new(Request)
	if err := c.Bind(request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	project := Project{
		Name: request.Name,
		Description: request.Description,
		StartDate: request.StartDate,
		EndDate: request.EndDate,
	}
	db.Create(&project)
	return c.JSON(http.StatusOK, project)
}
func updateProject(c echo.Context) error {
	type Request struct {
		ID string `json:"id" validate:"required"`
		Name string `json:"name"`
		Description string `json:"description"`
		StartDate string `json:"start_date"` 
		EndDate string `json:"end_date"`
	}
	request := new(Request)
	if err := c.Bind(request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	var project Project
	result := db.First(&project, request.ID)
	if result.Error != nil {
		return result.Error
	}
	if len(request.Name) > 0 {
		project.Name = request.Name
	}
	if len(request.Description) > 0 {
		project.Description = request.Description
	}
	if len(request.StartDate) > 0 {
		project.StartDate = request.StartDate
	}
	if len(request.EndDate) > 0 {
		project.EndDate = request.EndDate
	}
	db.Save(&project)
	return c.JSON(http.StatusOK, project)
}
func deleteProject(c echo.Context) error {
	type Request struct {
		ID int `json:"id" validate:"required"`
	}

	request := new(Request)
	if err := c.Bind(request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if result := db.Delete(&Project{}, request.ID); result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, request)
}
func findTasks(c echo.Context) error {
	var tasks []Task
	db.Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}
func findTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	var task Task
	db.First(&task, id)
	return c.JSON(http.StatusOK, task)
}
func createTask(c echo.Context) error {
	type Request struct {
		Title string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
		Status string `json:"status" validate:"required"`
		StartDate string `json:"start_date" validate:"required"`
		EndDate string `json:"end_date" validate:"required"`
		ProjectID uint `json:"project_id"`
		UserID uint `json:"user_id"`
	}
	request := new(Request)
	if err := c.Bind(request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	task := Task{
		Title: request.Title,
		Description: request.Description,
		Status: request.Status,
		StartDate: request.StartDate,
		EndDate: request.EndDate,
	}
	if request.ProjectID != 0 {
		task.ProjectID = request.ProjectID
	}
	if request.UserID != 0 {
		task.UserID = request.UserID
	}
	db.Create(&task)
	return c.JSON(http.StatusOK, task)
}
func updateTask(c echo.Context) error {
	type Request struct {
		ID string `json:"id" validate:"required"`
		Title string `json:"title"`
		Description string `json:"description"`
		Status string `json:"status"`
		StartDate string `json:"start_date"`
		EndDate string `json:"end_date"`
		ProjectID uint `json:"project_id"`
		UserID uint `json:"user_id"`
	}
	request := new(Request)
	if err := c.Bind(request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	var task Task
	result := db.First(&task, request.ID)
	if result.Error != nil {
		return result.Error
	}
	if len(request.Title) > 0 {
		task.Title = request.Title
	}
	if len(request.Description) > 0 {
		task.Description = request.Description
	}
	if len(request.Status) > 0 {
		task.Status = request.Status
	}
	if len(request.StartDate) > 0 {
		task.StartDate = request.StartDate
	}
	if len(request.EndDate) > 0 {
		task.EndDate = request.EndDate
	}
	if request.ProjectID != 0 {
		task.ProjectID = request.ProjectID
	}
	if request.UserID != 0 {
		task.UserID = request.UserID
	}
	db.Save(&task)
	return c.JSON(http.StatusOK, task)
}
func deleteTask(c echo.Context) error {
	type Request struct {
		ID int `json:"id" validate:"required"`
	}

	request := new(Request)
	if err := c.Bind(request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if result := db.Delete(&Task{}, request.ID); result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, request)
}

// Main Stream
func connectDb() {
	var err error
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database (Initialize)")
	}
}
func initializeDb() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Task{})
}
func main() {
	e := echo.New()
	e.Debug = true
	e.Validator = &Validator{validator: validator.New()}
	e.GET("/health", health)
	e.GET("/users", findUsers)
	e.GET("/user/:id", findUser)
	e.GET("/user/:id/tasks", notImplemented) // findTasksSpecifiedUser
	e.POST("/user/create", createUser)
	e.POST("/user/update", updateUser)
	e.POST("/user/delete", deleteUser)
	e.GET("/projects", findProjects)
	e.GET("/project/:id", findProject)
	e.GET("/project/:id/tasks", notImplemented) // findTasksSpecifiedProject
	e.POST("/project/create", createProject)
	e.POST("/project/update", updateProject)
	e.POST("/project/delete", deleteProject)
	e.GET("/tasks", findTasks)
	e.GET("/task/:id", findTask)
	e.POST("/task/create", createTask)
	e.POST("/task/update", updateTask)
	e.POST("/task/delete", deleteTask)

	connectDb()
	initializeDb()
	e.Logger.Fatal(e.Start(":5000"))
}
package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

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
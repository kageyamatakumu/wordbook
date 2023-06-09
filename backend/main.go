package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// ルートエンドポイントのハンドラ
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!")
	})

	// サーバーを開始
	e.Start(":8080")
}

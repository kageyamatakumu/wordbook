package main

import (
	"fmt"
	"net/http"
	"wordbook_app/backend/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// dbの接続を確認
	db := db.NewDB()
	fmt.Println(db)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	// ルートエンドポイントのハンドラ
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// API通信が正常に行えるか確認
	e.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "backend to fromtend")
	})

	// サーバーを開始
	e.Start(":8080")
}

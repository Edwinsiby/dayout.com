package main

import (
	"fmt"
	"main/Handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Can't load env")
	}

	router := gin.New()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	router.GET("/", Handlers.IndexHandler)
	router.GET("/signup", Handlers.SignupHandler)
	router.POST("/signuppost", Handlers.SignupPost)
	router.GET("/login", Handlers.LoginHandler)
	router.POST("/loginpost", Handlers.LoginPost)
	router.GET("/home", Handlers.HomeHandler)
	router.GET("/logout", Handlers.LogoutHandler)

	router.GET("/admin", Handlers.AdminHandler)
	router.GET("/admin/edit", Handlers.EditHandler)
	router.GET("/admin/delete", Handlers.DeleteHandler)
	router.POST("/update", Handlers.UpdateHandler)
	router.GET("/loadcreate", Handlers.LoadcreateHandler)
	router.POST("/create", Handlers.CreateHandler)
	router.POST("/search", Handlers.SearchHandler)

	router.Run(":8081")
}

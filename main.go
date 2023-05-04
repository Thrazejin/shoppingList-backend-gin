// loading configuration from config.yaml
/*
app:
  name: "shoppingList-backend-gin"
*/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"shoppingList-backend-gin.com/m/config"
	"shoppingList-backend-gin.com/m/controller"
	"shoppingList-backend-gin.com/m/middleware"
	"shoppingList-backend-gin.com/m/model"
)

func Engine() *gin.Engine {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)

	protectedRoutes.POST("/unit", controller.CreateUnit)
	protectedRoutes.GET("/unit", controller.GetUnities)
	protectedRoutes.GET("/unit/:id", controller.GetUnitById)
	protectedRoutes.PUT("/unit/:id", controller.UpdateUnit)
	protectedRoutes.PATCH("/unit/:id", controller.UpdateUnit)
	protectedRoutes.DELETE("/unit/:id", controller.DeleteUnitById)

	return router
}

func main() {

	if err := config.ReadConfig("DEV"); err != nil {
		return
	}

	if err := model.ConnectDatabase(); err != nil {
		fmt.Println("ERROR connecting to database: ", err.Error())
		return
	}

	if err := Engine().Run(); err != nil {
		fmt.Println("ERROR running server: ", err.Error())
		return
	}
}

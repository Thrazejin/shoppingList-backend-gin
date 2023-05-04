package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"shoppingList-backend-gin.com/m/config"
	"shoppingList-backend-gin.com/m/controller"
	"shoppingList-backend-gin.com/m/middleware"
	"shoppingList-backend-gin.com/m/model"
)

func router() *gin.Engine {
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

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()
	teardown()

	os.Exit(exitCode)
}

func setup() error {
	if err := config.ReadConfig("QA"); err != nil {
		return err
	}

	if err := model.ConnectDatabase(); err != nil {
		return err
	}

	return nil
}

func teardown() {
	migrator := model.DB.Migrator()

	migrator.DropTable(&model.AppUser{})
	migrator.DropTable(&model.Entry{})
	migrator.DropTable(&model.Unit{})
}

func makeRequest(method, url string, body interface{}, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if isAuthenticatedRequest {
		request.Header.Add("Authorization", "Bearer "+bearerToken())
	}
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}

func bearerToken() string {
	user := model.AuthenticationInput{
		Username: "thraze",
		Password: "test",
	}

	writer := makeRequest("POST", "/auth/login", user, false)
	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	return response["jwt"]
}

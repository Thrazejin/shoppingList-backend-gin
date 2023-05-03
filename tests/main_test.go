package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
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

	return router
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()
	teardown()

	os.Exit(exitCode)
}

func setup() {
	model.ConnectDatabase()
	model.DB.AutoMigrate(&model.User{})
	model.DB.AutoMigrate(&model.Entry{})
}

func teardown() {
	migrator := model.DB.Migrator()
	migrator.DropTable(&model.User{})
	migrator.DropTable(&model.Entry{})
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
	user := controller.AuthenticationInput{
		Username: "yemiwebby",
		Password: "test",
	}

	writer := makeRequest("POST", "/auth/login", user, false)
	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	return response["jwt"]
}

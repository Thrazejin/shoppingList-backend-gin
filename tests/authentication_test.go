package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"shoppingList-backend-gin.com/m/controller"
)

func TestRegister(t *testing.T) {
	newUser := controller.AuthenticationInput{
		Username: "yemiwebby",
		Password: "test",
	}
	writer := makeRequest("POST", "/auth/register", newUser, false)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestLogin(t *testing.T) {
	user := controller.AuthenticationInput{
		Username: "yemiwebby",
		Password: "test",
	}

	writer := makeRequest("POST", "/auth/login", user, false)

	assert.Equal(t, http.StatusOK, writer.Code)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["jwt"]

	assert.Equal(t, true, exists)
}

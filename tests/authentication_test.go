package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/magiconair/properties/assert"
	"shoppingList-backend-gin.com/m/model"
)

func TestRegister(t *testing.T) {
	newUser := model.UserInput{
		Name:     "Guilherme de Souza Pereira",
		Username: "thraze",
		Password: "test",
	}
	writer := makeRequest("POST", "/auth/register", newUser, false)
	assert.Equal(t, writer.Code, http.StatusCreated)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["jwt"]

	if writer.Code == 400 {
		fmt.Println(response["error"])
		return
	}

	name, exists := response["name"]
	assert.Equal(t, true, exists)
	assert.Equal(t, name, newUser.Name)

	shortName, exists := response["username"]
	assert.Equal(t, true, exists)
	assert.Equal(t, shortName, newUser.Username)
}

func TestLogin(t *testing.T) {
	user := model.AuthenticationInput{
		Username: "thraze",
		Password: "test",
	}

	writer := makeRequest("POST", "/auth/login", user, false)

	assert.Equal(t, writer.Code, http.StatusOK)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["jwt"]

	assert.Equal(t, exists, true)
}

package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/magiconair/properties/assert"
	"shoppingList-backend-gin.com/m/model"
)

func TestCreateUnit(t *testing.T) {
	unit := model.UnitInput{
		Name:      "Unidade",
		ShortName: "un",
	}

	writer := makeRequest("POST", "/api/unit", unit, true)
	assert.Equal(t, writer.Code, http.StatusCreated)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	name, exists := response["name"]
	assert.Equal(t, exists, true)
	assert.Equal(t, name, unit.Name)

	shortName, exists := response["shortName"]
	assert.Equal(t, exists, true)
	assert.Equal(t, shortName, unit.ShortName)

}

func TestGetUnits(t *testing.T) {
	writer := makeRequest("GET", "/api/unit", nil, true)
	assert.Equal(t, writer.Code, http.StatusOK)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	fmt.Println(response)
}

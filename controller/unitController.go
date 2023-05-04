package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"shoppingList-backend-gin.com/m/helper"
	"shoppingList-backend-gin.com/m/model"
)

// -------------------------
// POST /unit
// Create an unit
// -------------------------
func CreateUnit(context *gin.Context) {
	// Get Corrent User
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Put the request body into the input varriable
	var input model.UnitInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validates if it already exists
	if _, err := input.FindIfDoesntExist(user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Try to save
	unit, err := input.Save(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the UnitOutput representation of the unit saved
	context.JSON(http.StatusCreated, unit.ParseOutput())
}

// -------------------------
// PATCH /unit/:id
// PUT /unit/:id
// Update an unit
// -------------------------
func UpdateUnit(context *gin.Context) {
	// Get Corrent User
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate if the id in query params is an interger
	id, err := strconv.ParseInt(context.Param("id"), 10, 0)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Query Param 'id' should be integer"})
	}

	// Put the request body into the input varriable
	var input model.UnitUpdateInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate if the id in query params exists
	var unit model.Unit

	if err := model.GetUnitById(user, &unit, int(id)); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validates if it already exists
	if _, err := input.FindIfDoesntExist(user, unit); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate if can be updated
	if _, err := input.Update(unit); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the UnitOutput representation of the unit saved
	context.JSON(http.StatusCreated, unit.ParseOutput())
}

// -------------------------
// Get /unit
// get all user unities
// -------------------------
func GetUnities(context *gin.Context) {
	// Get Corrent User
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get a array of all user unities
	unities, err := model.FindAllUnities(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the a array of UnitOutput representation of the unities
	context.JSON(http.StatusOK, model.ParseUnitArrayOutput(unities))
}

// -------------------------
// Get /unit/:id
// get an user unities
// -------------------------
func GetUnitById(context *gin.Context) {
	// Get Corrent User
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate if the id in query params is an interger
	id, err := strconv.ParseInt(context.Param("id"), 10, 0)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Query Param 'id' should be integer"})
		return
	}

	// Get the Unit with this id if exists and belong to the user
	var unit model.Unit

	if err := model.GetUnitById(user, &unit, int(id)); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the UnitOutput representation of the unit
	context.JSON(http.StatusOK, unit.ParseOutput())
}

// -------------------------
// Get /unit/:id
// delete an unities
// -------------------------
func DeleteUnitById(context *gin.Context) {
	// Get Corrent User
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate if the id in query params is an interger
	id, err := strconv.ParseInt(context.Param("id"), 10, 0)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Query Param 'id' should be integer"})
		return
	}

	// Get the Unit with this id if exists and belong to the user
	var unit model.Unit

	if err := model.GetUnitById(user, &unit, int(id)); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delete the Unit
	if err := unit.Delete(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return a no Content response
	context.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	context.Writer.WriteHeader(http.StatusNoContent)
}

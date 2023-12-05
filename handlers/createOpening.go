package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/AndreSDS/goApi/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opnening := schemas.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	}

	if err := db.Create(&opnening).Error; err != nil {
		logger.Errorf("failed to create opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening")
		return
	}

	sendSuccess(ctx, "create-opening", opnening)
}
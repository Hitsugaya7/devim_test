package controllers

import (
	"github.com/Hitsugaya/rest-api-project/cmd/internal/models"
	"github.com/Hitsugaya/rest-api-project/cmd/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	services *services.AllServices
}

func NewController(services *services.AllServices) *Controller {
	return &Controller{services: services}
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, models.ErrorResponse{Message: message})
}

func (c *Controller) Register(router *gin.Engine) {
	router.GET("/exchange-rate", c.exchangeRate)
}

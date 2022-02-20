package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) exchangeRate(context *gin.Context) {
	coordX := context.Query("coord_x")
	coordY := context.Query("coord_y")

	fmt.Println(coordX)
	fmt.Println(coordY)

	exchangeCurrency, err := c.services.ExchangeService.GetRate(coordX, coordY)
	if err != nil {
		ErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"exchangeRate": exchangeCurrency,
	})
}

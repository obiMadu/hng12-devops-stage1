package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/number-classifier/internal/calculator"
	"github.com/number-classifier/internal/client"
	"github.com/number-classifier/pkg/models"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func ClassifyNumber(c *gin.Context) {
	// Get number from query parameter
	numStr := c.Query("number")
	if numStr == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Number: "",
			Error:  true,
		})
		return
	}

	// Convert to integer
	num, err := strconv.Atoi(numStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Number: numStr,
			Error:  true,
		})
		return
	}

	// Calculate properties
	properties := []string{}
	if calculator.IsArmstrong(num) {
		if calculator.IsEven(num) {
			properties = append(properties, "armstrong", "even")
		} else {
			properties = append(properties, "armstrong", "odd")
		}
	} else {
		if calculator.IsEven(num) {
			properties = append(properties, "even")
		} else {
			properties = append(properties, "odd")
		}
	}

	// Get fun fact from Numbers API
	funFact, err := client.GetNumberFact(num)
	if err != nil {
		funFact = fmt.Sprintf("%d is an interesting number", num)
	}

	response := models.NumberResponse{
		Number:     num,
		IsPrime:    calculator.IsPrime(num),
		IsPerfect:  calculator.IsPerfect(num),
		Properties: properties,
		DigitSum:   calculator.GetDigitSum(num),
		FunFact:    funFact,
	}

	c.JSON(http.StatusOK, response)
}

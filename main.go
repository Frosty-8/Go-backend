package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Numbers struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

func main() {
	r := gin.Default()

	// Optional: avoid proxy warning
	r.SetTrustedProxies(nil)

	// Root
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Go backend running 🚀",
			"status":  "healthy",
		})
	})

	// GET add (query params)
	r.GET("/add", func(c *gin.Context) {
		a, _ := strconv.ParseFloat(c.Query("a"), 64)
		b, _ := strconv.ParseFloat(c.Query("b"), 64)

		c.JSON(http.StatusOK, gin.H{
			"operation": "add",
			"result":    a + b,
		})
	})

	// GET multiply
	r.GET("/multiply", func(c *gin.Context) {
		a, _ := strconv.ParseFloat(c.Query("a"), 64)
		b, _ := strconv.ParseFloat(c.Query("b"), 64)

		c.JSON(http.StatusOK, gin.H{
			"operation": "multiply",
			"result":    a * b,
		})
	})

	// POST JSON add (FastAPI style)
	r.POST("/add", func(c *gin.Context) {
		var nums Numbers

		if err := c.ShouldBindJSON(&nums); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid input",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"operation": "add",
			"result":    nums.A + nums.B,
		})
	})

	// Start server
	r.Run(":8080")
}
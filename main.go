package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rurumimic/blitzkv/internal/kvstore"
)

func main() {
	store := kvstore.NewMemStore()
	store.Set("hello", "world")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/store", func(c *gin.Context) {
		key := c.Query("key")
		if key == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Key is required",
			})
			return
		}

		value, err := store.Get(key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get value",
			})
			return
		}
		if value == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Key not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"key":   key,
			"value": value,
		})
	})

	r.POST("/store", func(c *gin.Context) {
		var json struct {
			Key   string `json:"key" binding:"required"`
			Value string `json:"value" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input",
			})
			return
		}
		if err := store.Set(json.Key, json.Value); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to set value",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Value set successfully",
			"key":     json.Key,
			"value":   json.Value,
		})
	})

	r.DELETE("/store", func(c *gin.Context) {
		key := c.Query("key")
		if key == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Key is required",
			})
			return
		}
		if err := store.Delete(key); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to delete key",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Key deleted successfully",
			"key":     key,
		})
	})

	r.Run()
}

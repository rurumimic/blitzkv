package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rurumimic/blitzkv/internal/kvstore"
)

type Handler struct {
	store *kvstore.MemStore
}

func NewStoreHandler(store *kvstore.MemStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) GetValue(c *gin.Context) {
	key := c.Query("key")

	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Key is required",
		})
		return
	}

	value, err := h.store.Get(key)
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
}

func (h *Handler) SetValue(c *gin.Context) {
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
	if err := h.store.Set(json.Key, json.Value); err != nil {
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
}

func (h *Handler) DeleteValue(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Key is required",
		})
		return
	}
	if err := h.store.Delete(key); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete key",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Key deleted successfully",
		"key":     key,
	})

}

func (h *Handler) ListKeys(c *gin.Context) {
	keys, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to list keys",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"keys": keys,
	})
}

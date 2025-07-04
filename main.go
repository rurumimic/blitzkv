package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rurumimic/blitzkv/internal/handler"
	"github.com/rurumimic/blitzkv/internal/kvstore"
	"github.com/rurumimic/blitzkv/pkg/config"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}
	config.Display(conf)

	store := kvstore.NewMemStore()
	store.Set("hello", "world")

	storeHandler := handler.NewStoreHandler(store)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/store", storeHandler.GetValue)
	r.POST("/store", storeHandler.SetValue)
	r.DELETE("/store", storeHandler.DeleteValue)
	r.GET("/store/keys", storeHandler.ListKeys)

	r.Run(":" + strconv.Itoa(conf.Server.Port))
}

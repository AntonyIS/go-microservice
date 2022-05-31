package main

import (
	h "example.com/app/handler"
	db "example.com/app/model"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	handler := h.NewHandler()
	// Routes
	router.GET("/", handler.Home)
	router.GET("/movies", handler.GetMovies)
	router.GET("/movies/:id", handler.GetMovie)
	router.POST("/movies", handler.CreateMovie)
	router.DELETE("/movies/:id", handler.DeleteMovie)

	defer db.DB.Close()
	router.Run(":5000")

}

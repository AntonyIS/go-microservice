package handler

import (
	"fmt"
	"net/http"

	"example.com/app/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Handler interface {
	Home(c *gin.Context)
	GetMovies(c *gin.Context)
	GetMovie(c *gin.Context)
	CreateMovie(c *gin.Context)
	DeleteMovie(c *gin.Context)
}

type handler struct {
	db *gorm.DB
}

func NewHandler() Handler {
	h := &handler{}
	h.db = model.InitDB()
	return h
}
func (h *handler) Home(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "welcome",
	})

}

func (h *handler) GetMovies(c *gin.Context) {
	movies := []model.Movie{}

	h.db.Find(&movies)
	c.JSON(http.StatusOK, gin.H{
		"movies": movies,
	})

}

func (h *handler) GetMovie(c *gin.Context) {
	id := c.Param("id")
	movie := model.Movie{}

	h.db.First(&movie, id)

	if movie.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"movie": movie,
	})
}

func (h *handler) CreateMovie(c *gin.Context) {
	movie := model.Movie{}

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.db.Create(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"movie": movie,
	})
}

func (h *handler) DeleteMovie(c *gin.Context) {
	var movie model.Movie
	h.db.First(&movie, c.Param("id"))
	tmp_name := movie.Name
	h.db.Delete(&movie)
	if err := h.db.Create(&movie).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s deleted successfuly", tmp_name)})
		return
	}
}

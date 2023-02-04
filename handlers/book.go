package handler

import (
	"log"
	"net/http"

	"github.com/DigantaChauduri06/week2-GL1-CipherSchools/database"
	"github.com/DigantaChauduri06/week2-GL1-CipherSchools/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
type Handler struct{}
func GetBooks(in *gin.Context) {
	DB *gorm.DB
}

func (h *Handler) GetBooks(in *gin.Context) {
	books, err := database.GetBooks(h.DB) // h.DB is fully configured and can give access of book table
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, books)

}
func (h *Handler) SaveBook(in *gin.Context) {
	book := models.Book{}
	err := in.BindJSOn(&book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}

	err = database.SaveBook(h.DB, &book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, book)
}

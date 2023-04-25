package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"halyklife-lib/config"
	"halyklife-lib/models"
	"net/http"
	"strconv"
)

func respondWithError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	config.DB.Preload("Author").Preload("Readers").Find(&books)
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	var book models.Book
	result := config.DB.Preload("Author").Preload("Readers").First(&book, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		respondWithError(w, http.StatusNotFound, result.Error)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	result := config.DB.Create(&book)
	if result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, result.Error)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	var book models.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	result := config.DB.First(&models.Book{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		respondWithError(w, http.StatusNotFound, result.Error)
		return
	}

	book.ID = uint(id)
	result = config.DB.Save(&book)
	if result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, result.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	var book models.Book
	result := config.DB.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		respondWithError(w, http.StatusNotFound, result.Error)
		return
	}

	config.DB.Delete(&book)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted successfully"})
}

func GetReaders(w http.ResponseWriter, r *http.Request) {
	var readers []models.Reader
	config.DB.Preload("BorrowedBooks").Find(&readers)
	json.NewEncoder(w).Encode(readers)
}

func GetReader(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	var reader models.Reader
	result := config.DB.Preload("BorrowedBooks").First(&reader, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		respondWithError(w, http.StatusNotFound, result.Error)
		return
	}

	json.NewEncoder(w).Encode(reader)
}
func CreateReader(w http.ResponseWriter, r *http.Request) {
	var reader models.Reader
	err := json.NewDecoder(r.Body).Decode(&reader)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	result := config.DB.Create(&reader)
	if result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, result.Error)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reader)
}

func UpdateReader(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	var reader models.Reader
	err = json.NewDecoder(r.Body).Decode(&reader)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	result := config.DB.First(&models.Reader{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		respondWithError(w, http.StatusNotFound, result.Error)
		return
	}

	reader.ID = uint(id)
	result = config.DB.Save(&reader)
	if result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, result.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reader)
}

func DeleteReader(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	var reader models.Reader
	result := config.DB.First(&reader, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		respondWithError(w, http.StatusNotFound, result.Error)
		return
	}

	config.DB.Delete(&reader)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Reader deleted successfully"})
}

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	config.DB.Preload("Books").Find(&authors).Preload("Books")
	json.NewEncoder(w).Encode(authors)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	var author models.Author
	result := config.DB.Preload("Books").First(&author, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		respondWithError(w, http.StatusNotFound, result.Error)
		return
	}

	json.NewEncoder(w).Encode(author)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	result := config.DB.Create(&author)
	if result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, result.Error)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	var author models.Author
	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	result := config.DB.First(&models.Author{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		respondWithError(w, http.StatusNotFound, result.Error)
		return
	}

	author.ID = uint(id)
	result = config.DB.Save(&author)
	if result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, result.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	var author models.Author
	result := config.DB.First(&author, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		respondWithError(w, http.StatusNotFound, result.Error)
		return
	}

	config.DB.Delete(&author)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Author deleted successfully"})
}

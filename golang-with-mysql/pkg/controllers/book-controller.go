package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/atharva/golang-with-mysql/pkg/models"
	"github.com/atharva/golang-with-mysql/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookbyId(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// bookId := vars["bookId"]
	w.Header().Set("Content-Type", "pkglication/json")
	bookId := strings.TrimPrefix(r.URL.Path, "/book/")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing", err, ID, bookId)
	}
	bookDetails, _ := models.GetBookById(ID)
	if bookDetails.Name == "" {
		msg := "{\"error\":\"No Book with given ID found\"}"
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		return
	}
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// bookId := vars["bookId"]
	bookId := strings.TrimPrefix(r.URL.Path, "/book/")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing", err, ID, bookId)
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, updateBook)

	bookId := strings.TrimPrefix(r.URL.Path, "/book/")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing", err, ID, bookId)
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

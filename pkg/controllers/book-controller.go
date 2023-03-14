package controllers

import (
	"book-management-go/pkg/models"
	"book-management-go/pkg/utils"
	"encoding/json"
	"net/http"
)

var NewBook models.Book

func GetBook(res http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	result, _ := json.Marshal(newBooks)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(req, createBook)
	b := createBook.CreateBook()
	result, _ := json.Marshal(b)
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

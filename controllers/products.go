package controllers

import (
	"goWebApp/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))


func Index(writer http.ResponseWriter, request *http.Request) {
	products := models.GetAllProducts()
	temp.ExecuteTemplate(writer,"Index",products)
}

func New(writer http.ResponseWriter, request *http.Request) {
	temp.ExecuteTemplate(writer,"New",nil)
}
func Delete(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	models.DeleteProduct(id)
	products := models.GetAllProducts()
	temp.ExecuteTemplate(writer,"Index",products)
}

func Edit(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	product := models.GetProduct(id)
	temp.ExecuteTemplate(writer,"Edit",product)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		name := request.FormValue("name")
		description := request.FormValue("description")
		price := request.FormValue("price")
		priceConverted, err := strconv.ParseFloat(price,32)
		if err != nil {
			log.Println("Error getting price")
		}
		amount := request.FormValue("amount")
		amountConverted, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error getting amount")
		}
		product := models.Product{Name: name, Description: description, Price: float32(priceConverted), Amount: amountConverted}
		models.UpdateProduct(product)
		http.Redirect(writer,request,"/",301)
	}
}

func Insert(writer http.ResponseWriter, request *http.Request)  {
	if request.Method == "POST"{
		name := request.FormValue("name")
		description := request.FormValue("description")
		price := request.FormValue("price")
		priceConverted, err := strconv.ParseFloat(price,32)
		if err != nil {
			log.Println("Error getting price")
		}
		amount := request.FormValue("amount")
		amountConverted, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error getting amount")
		}
		product := models.Product{Name: name, Description: description, Price: float32(priceConverted), Amount: amountConverted}
		models.CreateNewProduct(product)
		http.Redirect(writer,request,"/",301)
	}
}

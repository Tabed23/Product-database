package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"./database"
	"./models"
	"github.com/gorilla/mux"
)

var products []models.Product

func GetProducts(res http.ResponseWriter, req *http.Request) {

	db, err := database.GetDatabase()
	if err != nil {
		panic("Database not connected")
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		panic("Some Database error")
	}

	for rows.Next() {
		var product models.Product
		rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
		products = append(products, product)
	}
	json.NewEncoder(res).Encode(&products)
}

func GetProduct(res http.ResponseWriter, req *http.Request) {
	db, err := database.GetDatabase()
	if err != nil {
		panic("Database not connected")
	}
	params := mux.Vars(req)
	row, err := db.Query("SELECT * FROM product WHERE id =?", params["id"])
	if err != nil {
		panic("Query Error")
	}
	defer db.Close()
	var product models.Product
	for row.Next() {
		row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
	}

	json.NewEncoder(res).Encode(&product)
}

func CreateProduct(res http.ResponseWriter, req *http.Request) {
	db, err := database.GetDatabase()
	if err != nil {
		panic("Database not connected")
	}
	stat, err := db.Prepare("INSERT INTO product(id,name, price,quantity) VALUES(?,?,?,?)")
	if err != nil {
		panic("Query Error")
	}
	params := mux.Vars(req)
	stat.Exec(params["id"], params["name"], params["price"], params["quantity"])
	row, err := db.Query("SELECT * FROM product WHERE id=?", params["id"])
	var product models.Product
	for row.Next() {
		row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
	}
	products = append(products, product)
	json.NewEncoder(res).Encode(&products)

}
func DeleteProduct(res http.ResponseWriter, req *http.Request) {
	db, err := database.GetDatabase()
	if err != nil {
		fmt.Println(err)
		panic("Database not connected")
	}
	defer db.Close()
	params := mux.Vars(req)
	_, err2 := db.Query("DELETE FROM product WHERE id = ?", params["id"])
	if err2 != nil {
		panic("Query Error")
	}
	for index, item :=  range products{
		if string(item.ID) ==  params["id"]{
			products =  append(products[:index],products[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(&products)
}

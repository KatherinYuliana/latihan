package controllers

import (
	"encoding/json"
	m "latihan/models"
	"net/http"
	"strconv"
)

func InsertTransactions(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		return
	}
	userId, _ := strconv.Atoi(r.Form.Get("userId"))
	productId, _ := strconv.Atoi(r.Form.Get("productId"))
	quantity, _ := strconv.Atoi(r.Form.Get("quantity"))

	_, errQuery := db.Exec("INSERT INTO users(userId, productId, quantity) values (?,?,?)",
		userId,
		productId,
		quantity,
	)

	var response m.TransactionResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 400
		response.Message = "Insert Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

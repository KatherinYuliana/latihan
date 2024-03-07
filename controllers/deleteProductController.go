package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteSingleProduct(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	productId := vars["product_id"]

	_, errQuery := db.Exec("DELETE FROM products WHERE id=?",
		productId,
	)

	if errQuery == nil {
		sendSuccessResponse(w)
	} else {
		sendErrorResponse(w)
	}
}

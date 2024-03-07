package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	m "latihan/models"

	"github.com/gorilla/mux"
)

// GetAllTransactions
func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM transactions"

	// Read from Query Param
	userId := r.URL.Query()["userId"]
	productId := r.URL.Query()["productId"]
	if userId != nil {
		fmt.Println(userId[0])
		query += "WHERE name='" + userId[0] + "'"
	}

	if productId != nil {
		if userId[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " price='" + productId[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	var transaction m.Transaction
	var transactions []m.Transaction
	for rows.Next() {
		if err := rows.Scan(&transaction.ID, &transaction.UserId, &transaction.ProductId, &transaction.Quantity); err != nil {
			log.Println(err)
			return
		} else {
			transactions = append(transactions, transaction)
		}
	}
	w.Header().Set("Content-Type", "application/json")

	var response m.TransactionsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = transactions
	json.NewEncoder(w).Encode(response)
}

// InsertTransaction
func InsertTransaction(w http.ResponseWriter, r *http.Request) {
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

// DeleteTransaction
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	transactionId := vars["transaction_id"]

	_, errQuery := db.Exec("DELETE FROM transactions WHERE id=?",
		transactionId,
	)

	if errQuery == nil {
		sendSuccessResponse(w)
	} else {
		sendErrorResponse(w)
	}
}

// DeleteTransaction (GORM)
func DeleteTransaction2(w http.ResponseWriter, r *http.Request) {
	db := conn()

	// Parse parameters
	vars := mux.Vars(r)
	transactionID := vars["transaction_id"]

	// Hapus transaksi dari database
	result := db.Delete(&m.Transaction{}, transactionID)
	if result.Error != nil {
		http.Error(w, "Failed to delete transaction", http.StatusInternalServerError)
		return
	}

	// Berikan respons
	response := m.TransactionResponse{Status: http.StatusOK, Message: "Transaction deleted successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

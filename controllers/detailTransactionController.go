package controllers

import (
	"encoding/json"
	"net/http"

	m "latihan/models"
)

func GetDetailUserTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// 2 queries
	userRows, err := db.Query("SELECT * FROM users")
	if err != nil {
		sendErrorResponse(w)
		return
	}
	var user m.User
	var users []m.User
	for userRows.Next() {
		if err := userRows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.UserType); err != nil {
			sendErrorResponse(w)
			return
		} else {
			users = append(users, user)
		}
	}

	transactionRows, err := db.Query("SELECT * FROM transactions")
	if err != nil {
		sendErrorResponse(w)
		return
	}
	var transaction m.Transaction
	var transactions []m.Transaction
	for transactionRows.Next() {
		if err := transactionRows.Scan(&transaction.ID, &transaction.UserId, &transaction.ProductId, &transaction.Quantity); err != nil {
			sendErrorResponse(w)
			return
		} else {
			transactions = append(transactions, transaction)
		}
	}

	// 1 query with join
	query := `SELECT u.id, u.name, u.age, u.address, u.type, p.id, p.name, p.price, t.id, t.userId, t.productId, t.quantity FROM users u, products p JOIN transactions t ON u.ID = t.userId AND p.ID = t.productId`
	userTransactionRow, err := db.Query(query)
	if err != nil {
		print(err.Error())
		sendErrorResponse(w)
		return
	}
	var userTransaction m.UserTransaction
	var userTransactions []m.UserTransaction
	for userTransactionRow.Next() {
		if err := userTransactionRow.Scan(
			&userTransaction.User.ID, &userTransaction.User.Name, &userTransaction.User.Age,
			&userTransaction.User.Address, &userTransaction.User.UserType,
			&userTransaction.Transaction.ID, &userTransaction.Transaction.UserId, &userTransaction.Transaction.ProductId, &userTransaction.Transaction.Quantity); err != nil {
			print(err.Error())
			sendErrorResponse(w)
			return
		} else {
			userTransactions = append(userTransactions, userTransaction)
		}
	}

	var response m.UserTransactionsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = userTransactions
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

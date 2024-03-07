package controllers

import (
	"encoding/json"
	"net/http"

	m "latihan/models"
)

func GetUserAddresses(w http.ResponseWriter, r *http.Request) {
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

	addressRows, err := db.Query("SELECT * FROM detail_address")
	if err != nil {
		sendErrorResponse(w)
		return
	}
	var address m.Address
	var addresses []m.Address
	for addressRows.Next() {
		if err := addressRows.Scan(&address.ID, &address.Street, &address.UserID); err != nil {
			sendErrorResponse(w)
			return
		} else {
			addresses = append(addresses, address)
		}
	}

	// 1 query with join
	query := `SELECT u.id, u.name, u.age, u.address, u.type, da.id, da.street, da.user_id FROM users u JOIN detail_address da ON u.ID = da.user_id`
	userAddressRow, err := db.Query(query)
	if err != nil {
		print(err.Error())
		sendErrorResponse(w)
		return
	}
	var userAddress m.UserAddress
	var userAddresses []m.UserAddress
	for userAddressRow.Next() {
		if err := userAddressRow.Scan(
			&userAddress.User.ID, &userAddress.User.Name, &userAddress.User.Age,
			&userAddress.User.Address, &userAddress.User.UserType,
			&userAddress.Address.ID, &userAddress.Address.Street, &userAddress.Address.UserID); err != nil {
			print(err.Error())
			sendErrorResponse(w)
			return
		} else {
			userAddresses = append(userAddresses, userAddress)
		}
	}

	var response m.UserAddressesResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = userAddresses
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetUserAddresses2(w http.ResponseWriter, r *http.Request) {
	db := conn()

	var usersWithAddress []m.UserAddress

	// Query menggunakan GORM dengan Preload untuk memuat relasi detail alamat
	result := db.Preload("DetailAddress").Find(&usersWithAddress)
	if result.Error != nil {
		http.Error(w, "Failed to fetch user addresses", http.StatusInternalServerError)
		return
	}

	// Berikan respons
	response := m.UserAddressesResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    usersWithAddress,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

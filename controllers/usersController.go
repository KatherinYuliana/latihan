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

// GetAllUsers
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users"

	// Read from Query Param
	name := r.URL.Query()["name"]
	age := r.URL.Query()["age"]
	if name != nil {
		fmt.Println(name[0])
		query += "WHERE name='" + name[0] + "'"
	}

	if age != nil {
		if name[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " age='" + age[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	var user m.User
	var users []m.User
	for rows.Next() {
		// if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.UserType); err != nil {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address); err != nil {
			log.Println(err)
			return
		} else {
			users = append(users, user)
		}
	}
	w.Header().Set("Content-Type", "application/json")

	var response m.UsersResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = users
	json.NewEncoder(w).Encode(response)
}

// InsertUser
func InsertUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		return
	}
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")

	_, errQuery := db.Exec("INSERT INTO users(name, age, address) values (?,?,?)",
		name,
		age,
		address,
	)

	var response m.UserResponse
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

// DeleteUser
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	userId := vars["user_id"]

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?",
		userId,
	)

	if errQuery == nil {
		sendSuccessResponse(w)
	} else {
		sendErrorResponse(w)
	}
}

// InsertUser (GORM)
func InsertUser2(w http.ResponseWriter, r *http.Request) {
	db := conn()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")

	// Insert user ke dalam database
	err = db.Create(&m.User{Name: name, Age: age, Address: address}).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Berikan respons
	response := m.UserResponse{Status: http.StatusOK, Message: "Success"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetAllUsers (GORM)
func GetAllUsers2(w http.ResponseWriter, r *http.Request) {
	db := conn()

	var users []m.User

	// Query menggunakan GORM
	query := db.Model(&m.User{})

	// Read from Query Param
	if name := r.URL.Query().Get("name"); name != "" {
		query = query.Where("name = ?", name)
	}
	if ageStr := r.URL.Query().Get("age"); ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}
		query = query.Where("age = ?", age)
	}

	// Eksekusi query
	result := query.Find(&users)
	if result.Error != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	// Berikan respons
	response := m.UsersResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    users,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(w http.ResponseWriter) {
	var response m.UserResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter) {
	var response m.UserResponse
	response.Status = 400
	response.Message = "Failed"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

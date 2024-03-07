package controllers

import (
	"encoding/json"
	"fmt"
	m "latihan/models"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM login"

	// Mendekode payload JSON ke struct User
	var user m.Login
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request payload")
		return
	}

	// Memeriksa apakah email dan password sesuai dengan yang ada di database
	rows, err := db.Query(query)
	var login m.Login
	var login2 []m.Login
	for rows.Next() {
		if err := rows.Scan(&login.Email, &login.Password); err != nil {
			log.Println(err)
			return
		} else {
			login2 = append(login2, login)
		}
	}

	// Jika email dan password cocok, kirimkan respons sukses
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Login successful")
}

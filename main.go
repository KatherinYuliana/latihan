package main

import (
	"fmt"
	"latihan/controllers"
	"log"
	"net/http"

	//"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	//"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()

	// users
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", controllers.InsertUser).Methods("POST")
	router.HandleFunc("/users", controllers.DeleteUser).Methods("PUT")
	router.HandleFunc("/users/{user_id}", controllers.DeleteUser).Methods("PUT")

	// products
	router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/products", controllers.InsertProduct).Methods("POST")
	router.HandleFunc("/products", controllers.DeleteProduct).Methods("PUT")
	router.HandleFunc("/products/{product_id}", controllers.DeleteProduct).Methods("PUT")

	// transactions
	router.HandleFunc("/transactions", controllers.GetAllTransactions).Methods("GET")
	router.HandleFunc("/transactions", controllers.InsertTransaction).Methods("POST")
	router.HandleFunc("/transactions", controllers.DeleteTransaction).Methods("PUT")
	router.HandleFunc("/transactions/{transaction_id}", controllers.DeleteTransaction).Methods("PUT")

	// detail transaction
	router.HandleFunc("/transactions", controllers.GetDetailUserTransaction).Methods("GET")
	// delete single product
	router.HandleFunc("/products", controllers.DeleteSingleProduct).Methods("PUT")
	// insert transaction
	router.HandleFunc("/transactions", controllers.InsertTransactions).Methods("POST")
	// login
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	// GORM
	router.HandleFunc("/users", controllers.InsertUser2).Methods("POST")                 // insert user
	router.HandleFunc("/products", controllers.UpdateProduct).Methods("PUT")             // update product
	router.HandleFunc("/transactions", controllers.DeleteTransaction2).Methods("DELETE") // delete transaction
	router.HandleFunc("/users", controllers.GetAllUsers2).Methods("GET")                 // select user
	router.HandleFunc("/users", controllers.GetUserAddresses2).Methods("GET")            // select user address

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}

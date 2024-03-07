package models

// user
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	UserType int    `json:"type"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// product
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ProductResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Product `json:"data"`
}

type ProductsResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}

// transaction
type Transaction struct {
	ID        int `json:"id"`
	UserId    int `json:"userId"`
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type TransactionResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Transaction `json:"data"`
}

type TransactionsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Transaction `json:"data"`
}

// address
type Address struct {
	ID     int    `json:"id"`
	Street string `json:"street"`
	UserID int    `json:"user_id"`
}

type UserAddresses struct {
	User      User      `json:"user"`
	Addresses []Address `json:"addresses"`
}

type UserAddress struct {
	User    User    `json:"user"`
	Address Address `json:"addresses"`
}

type UserAddressesResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []UserAddress `json:"data"`
}

// user transaction
type UserTransaction struct {
	User        User        `json:"user"`
	Transaction Transaction `json:"transactions"`
}

type UserTransactions struct {
	User         User          `json:"user"`
	Transactions []Transaction `json:"transactions"`
}

type UserTransactionsResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []UserTransaction `json:"data"`
}

// login
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Login  `json:"data"`
}

type Login2Response struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Login `json:"data"`
}

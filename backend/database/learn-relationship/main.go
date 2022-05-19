package main

// many to many
// User table, name, age
// product table, name, price, quantity
// conjuction userID, ProductID

type UserRow struct {
	ID   int // primary Key
	Name string
	Age  int
}

type ProductRow struct {
	ID       int // primary Key
	Name     string
	Price    int
	Quantity int
}

type UserProductRow struct {
	UserID    int // Foreign Key
	ProductID int // Foreign Key
}

type UserTable []UserRow
type ProductTable []ProductRow
type UserProductTable []UserProductRow // conjuction table

type Database struct {
	Users        UserTable
	Products     ProductTable
	UserProducts UserProductTable
}

func main() {
	var (
		user1 = UserRow{1, "John", 23}
		user2 = UserRow{2, "Jane", 24}
		user3 = UserRow{3, "John", 24}
		user4 = UserRow{4, "Bernard", 30}
	)

	var (
		product1 = ProductRow{1, "Laptop", 100000, 10}
		product2 = ProductRow{2, "Mouse", 10000, 20}
		product3 = ProductRow{3, "Keyboard", 20000, 30}
		product4 = ProductRow{4, "Monitor", 300000, 40}
	)

	userTable := UserTable{user1, user2, user3, user4}
	productTable := ProductTable{product1, product2, product3, product4}

	var database Database

	database.Users = userTable
	database.Products = productTable

	var (
		userProduct1 = UserProductRow{UserID: 1, ProductID: 1}
		userProduct2 = UserProductRow{UserID: 1, ProductID: 3}
		userProduct3 = UserProductRow{UserID: 2, ProductID: 1}
		userProduct4 = UserProductRow{UserID: 2, ProductID: 3}
		userProduct5 = UserProductRow{UserID: 3, ProductID: 1}
		userProduct6 = UserProductRow{UserID: 3, ProductID: 4}
	)

	userProductTable := UserProductTable{userProduct1, userProduct2, userProduct3, userProduct4, userProduct5, userProduct6}

	database.UserProducts = userProductTable
}

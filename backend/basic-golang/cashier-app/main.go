package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/api"
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/repository"
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/terminal"
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/ui"
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/web"
)

// tambahan materi sedikit
// penjelasan API , http memberikan output json

// API : application program interface

// 1 folder , mono repo, single repo (html, css, js, php) -> complexity, data storage, scalability, SECURITY
// server & client
// backend & frontend
// (business logic, database) <-> (tampilan, responsive, efek2) (interface / GUI)

// API ()
// BE : golang - FE: web - javascript
// JSON (kirim2 data antar bahasa pemrograman)

// routingan / path

// authorization ?

// data => db => repository -> terminal
//							-> ui
//							-> api

func main() {
	// factory pattern

	// salah satu pilar Object Oriented programming => encapsulation
	// db := &db.CsvDB{}
	db := db.NewCsvDB()
	usersRepo := repository.NewUserRepository(db)
	productsRepo := repository.NewProductRepository(db)
	cartItemRepo := repository.NewCartItemRepository(db)
	transactionRepo := repository.NewTransactionRepository(cartItemRepo)

	fmt.Println("List Cashier App mode:")
	fmt.Println("1. Terminal")
	fmt.Println("2. Tview")
	fmt.Println("3. API")
	fmt.Println("4. Web")
	fmt.Printf("Please choose mode: ")
	var choice int
	fmt.Scan(&choice)
	fmt.Printf("\x1bc")

	switch choice {
	case 1:
		mainTerminal := terminal.NewTerminal(usersRepo, productsRepo, cartItemRepo, transactionRepo)
		mainTerminal.Start()
	case 2:
		mainUI := ui.NewUI(usersRepo, productsRepo, cartItemRepo, transactionRepo)
		mainUI.Start()
	case 3:
		mainAPI := api.NewAPI(usersRepo, productsRepo, cartItemRepo, transactionRepo)
		mainAPI.Start()
	case 4:
		mainAPI := api.NewAPI(usersRepo, productsRepo, cartItemRepo, transactionRepo)
		go mainAPI.Start()

		mainWEB := web.NewWEB(usersRepo, productsRepo, cartItemRepo, transactionRepo)
		mainWEB.Start()
	default:
		fmt.Println("Invalid choice")
	}

}

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var (
	sqlAccount = `CREATE TABLE accounts ( 
		account_no INTEGER NOT NULL, 
		balance DECIMAL NOT NULL DEFAULT 0,
		PRIMARY KEY(account_no),
			CHECK(balance >= 0)
	);`

	sqlAccountChange = `CREATE TABLE account_changes (
		change_no INT NOT NULL PRIMARY KEY,
		account_no INTEGER NOT NULL, 
		flag TEXT NOT NULL, 
		amount DECIMAL NOT NULL, 
		changed_at TEXT NOT NULL 
	);`
)

type Account struct {
	AccountNo int `db:"account_no"`
	Balance   int `db:"balance"`
}

type AccountChange struct {
	ChangeNo  int    `db:"change_no"`
	AccountNo int    `db:"account_no"`
	Flag      string `db:"flag"`
	Amount    int    `db:"amount"`
	ChangedAt string `db:"changed_at"`
}

func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./example_transaction.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(sqlAccount)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(sqlAccountChange)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	//db, err := Migrate()
	//if err != nil {
	//	panic(err)
	//}

	db, err := sql.Open("sqlite3", "./example_transaction.db")
	if err != nil {
		panic(err)
	}

	// BEGIN TRANSCTION;
	// sqlTx, err := db.BeginTx(context.Background(), nil)

	// _, err = sqlTx.Exec(`INSERT INTO accounts VALUES (100, 20100);`)
	// if err != nil {
	// 	sqlTx.Rollback()
	// 	panic(err)
	// }

	// _, err = sqlTx.Exec(`INSERT INTO account_changes (change_no,account_no,flag,amount,changed_at) VALUES (1, 100, '+', 20100, '2022-05-01');`)
	// if err != nil {
	// 	sqlTx.Rollback()
	// 	panic(err)
	// }

	// _, err = sqlTx.Exec(`INSERT INTO accounts VALUES (200, 10100);`)
	// if err != nil {
	// 	sqlTx.Rollback()
	// 	panic(err)
	// }

	// _, err = sqlTx.Exec(`INSERT INTO account_changes (change_no, account_no,flag,amount,changed_at) VALUES (2, 200, '+', 10100, '2022-05-01');`)
	// if err != nil {
	// 	sqlTx.Rollback()
	// 	panic(err)
	// }

	// err = sqlTx.Commit()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Transaction insert data SUCCESS")

	rows, err := db.Query(`SELECT * FROM accounts;`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var account Account
		err = rows.Scan(&account.AccountNo, &account.Balance)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", account)
	}
}

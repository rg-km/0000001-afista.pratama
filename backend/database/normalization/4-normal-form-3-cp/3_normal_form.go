package main

// Normalisasi database dalam bentuk 3NF bertujuan untuk menghilangkan seluruh atribut atau field yang tidak berhubungan dengan primary key.
// Dengan demikian tidak ada ketergantungan transitif pada setiap kandidat key
// fungsi normalisasi 3NF antara lain :
// 1. Memenuhi semua persyaratan dari bentuk normal kedua.
// 2. Menghapus kolom yang tidak tergantung pada primary key.

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//the struct are irrelevant for the code, but hint for column

type Rekap struct {
	NoBon      string
	KodeBarang string
	Harga      int
	Jumlah     int
	Biaya      int
	SubTotal   int
	Discount   int
	Total      int
	Bayar      int
	Kembalian  int
	KodeKasir  string
	Tanggal    string
	Waktu      string
}

type Rekap2 struct {
	NoBon     int
	Discount  int
	Total     int
	Bayar     int
	Kembalian int
	KodeKasir string
	Tanggal   string
	Waktu     string
}

type DetailRekap struct {
	NoBon      int
	KodeBarang string
	Harga      int
	Jumlah     int
}

type Barang struct {
	KodeBarang string
	NamaBarang string
	Harga      int
}

type Kasir struct {
	KodeKasir string
	NamaKasir string
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi Migrate:
// Buatlah tabel dengan nama rekap, rekap_detail, barang, dan kasir
// Lalu insert data ke masing-masing tabel seperti pada contoh di bagian bawah file ini
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE rekap (
		no_bon VARCHAR,
		discount INT,
		total INT,
		bayar INT,
		kembalian INT,
		kode_kasir VARCHAR,
		tanggal VARCHAR,
		waktu VARCHAR
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`INSERT INTO rekap VALUES 
	("00001", 77000, 0, 100000, 23000, "K01", "04-05-2022", "12:00:00"),
	("00002", 117500, 0, 117500, 0, "K02", "04-05-2022", "12:00:00")
	;`)

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE rekap_detail (
		no_bon VARCHAR,
		kode_barang VARCHAR,
		harga INT,
		jumlah INT
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`INSERT INTO rekap_detail 
	VALUES
	("00001", "B001", 4500, 3),
	("00001", "B002", 22500, 1),
	("00001", "B003", 1500, 4),
	("00001", "B004", 17500, 2),
	("00002", "B001", 4500, 1),
	("00002", "B004", 17400, 1),
	("00002", "BOO5", 100000, 1)
	;`)

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE barang (
		kode_barang VARCHAR PRIMARY KEY,
		nama_barang VARCHAR,
		harga INT
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`INSERT INTO barang VALUES 
	("B001", "Disket", 4500),
	("B002", "Refil Tinta", 22500),
	("B003", "CD Blank", 1500),
	("B004", "Mouse", 17500),
	("B005", "Flash Disk", 100000)
	;`)

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE kasir (
		kode_kasir VARCHAR PRIMARY KEY,
		nama_kasir VARCHAR
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`INSERT INTO kasir VALUES 
	("K01", "Rosi"),
 	("K02", "Dewi")
	;`)

	if err != nil {
		panic(err)
	}

	return db, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkNoBonExists:
// checkNoBonExists digunakan untuk menghitung jumlah data yang ada berdasarkan no_bon
func checkNoBonExists(noBon string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT COUNT(1) FROM rekap WHERE no_bon = ?;`

	row := db.QueryRow(sqlStmt, noBon)
	var countBon int
	err = row.Scan(&countBon)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi countRekapDetailByNoBon:
// countRekapDetailByNoBon digunakan untuk menghitung jumlah rekap detail yang ada berdasarkan no_bon
func countRekapDetailByNoBon(noBon string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT COUNT(1) FROM rekap_detail WHERE no_bon = ?;`

	row := db.QueryRow(sqlStmt, noBon)
	var countBon int
	err = row.Scan(&countBon)
	if err != nil {
		return 0, err
	}
	return countBon, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkBarangExists:
func checkBarangExists(kodeBarang string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT 1 FROM barang WHERE kode_barang = ?;`

	row := db.QueryRow(sqlStmt, kodeBarang)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkKasirExists:
func checkKasirExists(kodeKasir string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	// datanya ada 3
	// kode kasir ada diurutan ke 2

	// latest id 2

	sqlStmt := `SELECT 1 FROM kasir WHERE kode_kasir = ?;`

	row := db.QueryRow(sqlStmt, kodeKasir)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return false, err
	}
	return true, nil
}

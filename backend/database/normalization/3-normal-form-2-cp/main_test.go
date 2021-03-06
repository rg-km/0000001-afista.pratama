package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SQL Create", func() {
	BeforeEach(func() {
		_, err := sql.Open("sqlite3", "./normalize-cp.db")
		if err != nil {
			panic(err)
		}
		_, err = Migrate()
		if err != nil {
			panic(err)
		}
	})

	AfterEach(func() {
		db, err := sql.Open("sqlite3", "./normalize-cp.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
		DROP TABLE rekap;
		DROP TABLE barang;
		DROP TABLE kasir;`)

		if err != nil {
			panic(err)
		}

	})

	Describe("Check Migration", func() {
		It("should return number of column from rekap table", func() {
			db, err := sql.Open("sqlite3", "./normalize-cp.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('rekap');`)
			Expect(err).To(BeNil())
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(13))
			}
		})
	})

	Describe("Check Migration", func() {
		It("should return number of column from barang table", func() {
			db, err := sql.Open("sqlite3", "./normalize-cp.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('barang');`)
			Expect(err).To(BeNil())
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(3))
			}
		})
	})

	Describe("Check Migration", func() {
		It("should return number of column from kasir table", func() {
			db, err := sql.Open("sqlite3", "./normalize-cp.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('kasir');`)
			Expect(err).To(BeNil())
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(2))
			}
		})
	})

	Describe("Check Insert Latest data 1", func() {
		It("should return latest no_bon", func() {
			lastInsertData, err := countByNoBon("00002")
			Expect(err).To(BeNil())
			Expect(lastInsertData).To(Equal(3))

		})
	})

	Describe("Check Insert Latest data 2", func() {
		It("should return latest no_barang", func() {
			lastInsertData, err := checkBarangExists("B005")
			Expect(err).To(BeNil())
			Expect(lastInsertData).To(Equal(true))

		})
	})

	Describe("Check Insert Latest data 3", func() {
		It("should return latest no_kasir", func() {
			lastInsertData, err := checkKasirExists("K02")
			Expect(err).To(BeNil())
			Expect(lastInsertData).To(Equal(true))

		})
	})

})

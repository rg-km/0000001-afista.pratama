package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// clean code DRY (dont repeat yourself)
func handleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func WriteCSV(newData [][]string) {
	path := "../../golang-io/lecture/csv/data/siswaBaru.csv"
	dirFile, err := filepath.Abs(path)
	handleErr(err)

	file, err := os.OpenFile(dirFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	handleErr(err)

	defer file.Close()

	write := csv.NewWriter(file)

	// Write
	err = write.WriteAll(newData)
	handleErr(err)

	write.Flush()

	fmt.Println("berhasil membuat file csv")
}

func ReadCSV() [][]string {
	path := "../../golang-io/lecture/csv/data/siswaBaru.csv"

	dirFile, err := filepath.Abs(path)
	handleErr(err)

	file, err := os.OpenFile(dirFile, os.O_RDWR, os.ModePerm)
	handleErr(err)

	defer file.Close()

	csvReader := csv.NewReader(file)

	csvData, err := csvReader.ReadAll()
	handleErr(err)

	return csvData

	// [][]string
	/* [[nama, nilai, kelas],[v, v, v], [v2,v2,v2], dst]
	 */

	// for {
	// 	func generator / yield
	// 	record, err := csvReader.Read() //readline si txt (\n)
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	fmt.Println(record)
	// }

	// r1, _ := csvReader.Read()
	// r2, _ := csvReader.Read()
	// r3, _ := csvReader.Read()

	// fmt.Println(r1, r2, r3)

	// record, err := csvReader.Read()
	// handleErr(err)

	// fmt.Println(record)

	// readData, err := csvReader.ReadAll()
	// handleErr(err)

	// fmt.Println(readData)
}

// implement read csv
func main() {
	// readFile() => dapat data
	// kita olah => nambah orang lagi
	// writeFile() => output
	data := ReadCSV()

	data = append(data, []string{"Raam", "90", "BE01"})
	WriteCSV(data)
}

package db

import "os"

// OOP => encapsulation
type dataCSV struct{}

func NewDataCSV() dataCSV {
	return dataCSV{}
}

// open data, sebatas buka koneksi
func (d dataCSV) ReadCSV() (*os.File, error) {
	file, err := os.OpenFile("./data/data.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return nil, err
	}

	return file, nil
}

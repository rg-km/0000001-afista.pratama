package repository

import (
	"encoding/csv"
	"log"
	"os"
)

// oop => encapsulation

type dataRepository struct {
	F *os.File
}

func NewDataRepository(f *os.File) dataRepository {
	return dataRepository{
		F: f,
	}
}

func (d dataRepository) ReadData() [][]string {
	csvReader := csv.NewReader(d.F)

	readerData, err := csvReader.ReadAll()
	if err != nil {
		log.Panic(err)
	}
	return readerData
}

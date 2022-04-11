package db

type DBName = string
type TableName = string
type Rows = []Row  // [][]string
type Row = []Cell  // []string
type Cell = string // string => Cell

type DB interface {
	Load(dbName DBName) (Rows, error) //[][]string
	Save(dbName DBName, rows Rows) error
	Delete(dbName DBName) error
}

package main

type EmployeeRow struct {
	ID        int
	Name      string
	Position  string
	Salary    int
	ManagerID int
}
type EmployeeDB []EmployeeRow

func NewEmployeeDB() *EmployeeDB {
	return &EmployeeDB{}
}

func (db *EmployeeDB) Where(id int) *EmployeeRow {
	for i := 0; i < len(*db); i++ {
		if (*db)[i].ID == id {
			return &(*db)[i]
		}
	}
	return nil
}

func (db *EmployeeDB) Insert(name string, position string, salary int, managerID int) {
	(*db) = append(*db, EmployeeRow{
		ID:        len(*db) + 1,
		Name:      name,
		Position:  position,
		Salary:    salary,
		ManagerID: managerID,
	})
}

func (db *EmployeeDB) Update(id int, name string, position string, salary int, managerID int) {
	// lakukan pencarian dari array, menggunakan looping
	// kalau ketemu dengan parameter ID, maka kita ubah datanya

	for i := 0; i < len(*db); i++ {
		if (*db)[i].ID == id {
			(*db)[i].Name = name
			(*db)[i].Position = position
			(*db)[i].Salary = salary
			(*db)[i].ManagerID = managerID
			return
		}
	}
}

func (db *EmployeeDB) Delete(id int) {
	// lakukan pencarian dari array data, menggunakan looping
	// cocokkan id dengan parameter id yang ada di function
	// kalau sama, maka kita hapus datanya

	for i := 0; i < len(*db); i++ {
		if (*db)[i].ID == id {
			// 1,2,3,4,5,6
			// menghapus id 3
			// [1,2] + [4,5,6]
			// 1,2,4,5,6
			(*db) = append((*db)[:i], (*db)[i+1:]...)
			return
		}
	}
}

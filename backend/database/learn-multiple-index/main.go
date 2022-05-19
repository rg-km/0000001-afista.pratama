package main

// one to one
// one to many
// many to many

type UserRow struct {
	ID            int // primary Key
	Name          string
	Address       string
	Age           int
	PhoneNumberID []int
}

type PhoneNumberRow struct {
	ID          int // primary Key
	CountryCode int
	Number      int
}

type UserTable map[int]UserRow
type PhoneNumberTable map[int]PhoneNumberRow

func (d *UserTable) Search(id int) *UserRow {
	if userRow, ok := (*d)[id]; ok {
		return &userRow
	}
	return nil
}

func (p *PhoneNumberTable) Search(id int) *PhoneNumberRow {
	if phoneNumberRow, ok := (*p)[id]; ok {
		return &phoneNumberRow
	}
	return nil
}

func (d *UserTable) SearchPhoneNumber(userId int) []int {
	if userRow, ok := (*d)[userId]; ok {
		return userRow.PhoneNumberID
	}
	return nil
}

// multiple index

type PrimaryKey int
type SecondaryKey string
type ThirdKey string

type UserRowMap struct {
	ID            PrimaryKey   // primary Key
	Name          SecondaryKey // bisa duplicate
	Address       ThirdKey
	Age           int
	PhoneNumberID []int
}

type IndexById map[PrimaryKey]UserRowMap

type IndexByName map[SecondaryKey][]PrimaryKey

type IndexByAddress map[ThirdKey][]PrimaryKey

type UserDB struct {
	ById   IndexById
	ByName IndexByName
}

// method search by PrimaryKey
func (u *UserDB) SearchById(id PrimaryKey) *UserRowMap {
	if userRowMap, ok := u.ById[id]; ok {
		return &userRowMap
	}
	return nil
}

func (u *UserDB) SearchByName(name SecondaryKey) []*UserRowMap {
	ids := u.ByName[name]
	rows := make([]*UserRowMap, len(ids))
	//O(m), m == len(ids)
	for i, id := range ids {
		//O(1)
		rows[i] = u.SearchById(id)
	}
	return rows
}

func main() {

}

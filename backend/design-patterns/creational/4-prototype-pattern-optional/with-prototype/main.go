package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/design-patterns/creational/4-prototype-pattern-optional/with-prototype/house"
)

func main() {
	houseRumah := house.NewHouse(
		1,
		2,
		"bardi",
		"bodi",
	)

	houseRumah2 := house.CloneHouse(houseRumah)
	houseRumah3 := house.CloneHouse(houseRumah2).(*house.House)

	houseRumah3.NumOfDoors = 4
	houseRumah3.NumOfWindows = 4

	//houseRumah4 := houseRumah2.(*house.House)
	//houseRumah4.NumOfWindows = 3

	fmt.Println(houseRumah)
	fmt.Println(houseRumah2)
	fmt.Println(houseRumah3)
}

package main

import "fmt"

// custom otak atik dari si object / struct
type House struct {
	NumOfWindows    int
	NumOfDoors      int
	NumOfRooms      int
	HasKitchen      bool
	HasGarage       bool
	HasSwimmingPool bool
}

type HouseBuilder interface {
	buildWindow(numOfWindow int)
	buildDoor()
	buildRoom()
	buildKitchen()
	buildGarage()
	buildSwimmingPool()
	getHouse() House
}

// 2 house
type indonesiaHouseBuilder struct {
	house House
}

func (i *indonesiaHouseBuilder) buildWindow(numOfWindow int) {
	i.house.NumOfWindows = numOfWindow
}

func (i *indonesiaHouseBuilder) buildDoor() {
	i.house.NumOfDoors++
}

func (i *indonesiaHouseBuilder) buildRoom() {
	i.house.NumOfRooms++
}

func (i *indonesiaHouseBuilder) buildKitchen() {
	i.house.HasKitchen = true
}

func (i *indonesiaHouseBuilder) buildGarage() {
	i.house.HasGarage = true
}

func (i *indonesiaHouseBuilder) buildSwimmingPool() {
	i.house.HasSwimmingPool = true
}

func (i *indonesiaHouseBuilder) getHouse() House {
	return i.house
}

type zimbabweHouseBuilder struct {
	house House
}

func (i *zimbabweHouseBuilder) buildWindow(numOfWindow int) {
	if numOfWindow > 2 {
		i.house.NumOfWindows = 2
	} else {
		i.house.NumOfWindows = numOfWindow
	}
}

func (i *zimbabweHouseBuilder) buildDoor() {
	i.house.NumOfDoors++
}

func (i *zimbabweHouseBuilder) buildRoom() {
	i.house.NumOfRooms++
}

func (i *zimbabweHouseBuilder) buildKitchen() {
	i.house.HasKitchen = true
}

func (i *zimbabweHouseBuilder) buildGarage() {
	i.house.HasGarage = true
}

func (i *zimbabweHouseBuilder) buildSwimmingPool() {
	if i.house.NumOfRooms > 2 {
		i.house.HasSwimmingPool = true
	}
}

func (i *zimbabweHouseBuilder) getHouse() House {
	return i.house
}

type Kontraktor struct {
	builder HouseBuilder
}

func NewKontraktor(builder HouseBuilder) Kontraktor {
	// if builderType == "zimbabwe" {
	// 	return Kontraktor{builder: &zimbabweHouseBuilder{}}
	// }
	// if builderType == "indonesia" {
	// 	return Kontraktor{builder: &indonesiaHouseBuilder{}}
	// }
	// return Kontraktor{}

	return Kontraktor{builder: builder}
}

func (k *Kontraktor) BuildHouse(numOfWindow int) House {
	k.builder.buildWindow(numOfWindow)
	k.builder.buildDoor()
	k.builder.buildRoom()
	k.builder.buildKitchen()
	k.builder.buildGarage()
	k.builder.buildSwimmingPool()
	return k.builder.getHouse()
}

func main() {

	indoBuilder := &indonesiaHouseBuilder{}
	kontraktor := NewKontraktor(indoBuilder)
	building := kontraktor.BuildHouse(5)

	fmt.Println(building)

	zimbBuilder := &zimbabweHouseBuilder{}
	kontraktor = NewKontraktor(zimbBuilder)
	building = kontraktor.BuildHouse(5)

	fmt.Println(building)
}

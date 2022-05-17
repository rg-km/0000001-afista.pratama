package main

import "fmt"

// agar mudah dijalankan, implementasi interface akan ditaruh di file yang sama dengan context

type eyePowerStrategy interface { //ini adalah strategy
	useEyePower()
}

// core data, data penting, codingan penting
// hard wiring
//
type cat struct {
	eyePower eyePowerStrategy
}

func createCat(eyePower eyePowerStrategy) *cat {
	return &cat{
		eyePower: eyePower,
	}
}

// method general
func (c *cat) walk() {
	fmt.Println("cat walk")
}
func (c *cat) talk() {
	fmt.Println("cat meow")
}
func (c *cat) sleep() {
	fmt.Println("cat sleep")
}
func (c *cat) sit() {
	fmt.Println("cat sit")
}
func (c *cat) eat() {
	fmt.Println("cat eat")
}

// method specific
func (c *cat) useEyePower() {
	c.eyePower.useEyePower()
} // use eyePower implementation of useEyePower

// concreatePattern1
type laserEyePower struct{}

func (e *laserEyePower) useEyePower() {
	fmt.Println("shoot laser")
}

// concreatePattern2
type noEyePower struct{}

func (e *noEyePower) useEyePower() {
	fmt.Println("do nothing")
}

func main() {
	noLaser := noEyePower{}
	withLaser := laserEyePower{}

	cat1 := createCat(&noLaser)
	cat2 := createCat(&withLaser)

	cat3 := createCat(&withLaser)

	fmt.Println("=============== cat 1")
	cat1.eyePower.useEyePower()
	cat1.walk()
	cat1.eat()

	fmt.Println("=============== cat 2")
	cat2.eyePower.useEyePower()
	cat2.walk()
	cat2.eat()

	fmt.Println("=============== cat 3")
	cat3.eyePower.useEyePower()
	cat3.walk()
	cat3.eat()
}

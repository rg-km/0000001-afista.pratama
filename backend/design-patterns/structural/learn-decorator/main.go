package main

import "fmt"

// signature
type Opener interface {
	Open()
}

// base struct
type Door struct{}

func (d *Door) Open() {
	fmt.Println("Open door")
}

// concreate , implementasi si decorator
type SmartGagang struct {
	opener Opener
}

func NewSmartGagang(opener Opener) Opener {
	return &SmartGagang{opener: opener}
}

func (s *SmartGagang) Open() {
	fmt.Println("Open smart gagang")

	ok := s.ConnectToWIFI()
	if ok {
		s.opener.Open()
	}
}

func (s *SmartGagang) ConnectToWIFI() bool {
	fmt.Println("Connect to wifi")
	return true
}

type OldGagang struct {
	opener Opener
}

func NewOldGagang(opener Opener) Opener {
	return &OldGagang{opener: opener}
}

func (o *OldGagang) Open() {
	fmt.Println("Open old gagang")
	o.opener.Open()
}

// decoration to decoration
type SmartScreen struct {
	opener Opener
}

func NewSmartScreen(opener Opener) Opener {
	return &SmartScreen{opener: opener}
}

func (s *SmartScreen) Open() {
	fmt.Println("Open smart screen")
	s.opener.Open()
}

// client
func main() {
	d := Door{}
	d.Open()

	fmt.Println("=============")

	doorWithSmartGagang := NewSmartGagang(&d)
	withSmartScreen := NewSmartScreen(doorWithSmartGagang)
	withSmartScreen.Open()

	fmt.Println("=============")

	doorWithOldGagang := NewOldGagang(&d)
	oldWithSmartScreen := NewSmartScreen(doorWithOldGagang)
	oldWithSmartScreen.Open()
}

/*
			door
			  |
			  |
		smart gagang
			  |
			  |
		smart screen

contoh ke 2
			door
			  |
			  |
			  smart screen
			  |
			  |
		smart gagang

*/

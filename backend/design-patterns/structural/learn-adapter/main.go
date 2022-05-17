package main

import "fmt"

// legacy code
type Walkman struct{}

type Kaset struct {
	PitaMusik string
}

func (w *Walkman) PlayMusic(k *Kaset) {
	fmt.Println("Play from walkman :", k.PitaMusik)
}

// new code

type Mp3 struct {
	Data []byte
}

func (m *Mp3) PlayAudio() {
	fmt.Println("Play from mp3 :", string(m.Data))
}

// adapter
type AudioPlayer interface {
	PlayMusic(m Mp3)
}

type Mp3toWalkmanAdapter struct {
	Adaptee Walkman
}

func (a *Mp3toWalkmanAdapter) PlayMusic(m Mp3) {
	k := Kaset{}

	fmt.Println("using adapter : Play from mp3 to walkman")

	k.PitaMusik = string(m.Data)
	a.Adaptee.PlayMusic(&k)
}

// client
func main() {
	mp3 := Mp3{Data: []byte("Hello World")}

	walkman := Walkman{}
	adapt := Mp3toWalkmanAdapter{Adaptee: walkman}
	adapt.PlayMusic(mp3)

	// legacy
	kaset := Kaset{PitaMusik: "Hello World"}
	walkman.PlayMusic(&kaset)
}

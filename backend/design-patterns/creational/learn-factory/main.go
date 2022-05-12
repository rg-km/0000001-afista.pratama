package main

import "fmt"

// superclass
type Content interface {
	Play()
}

type AnimeContent struct {
	name string
}

func (a AnimeContent) Play() {
	fmt.Println(a.name)
	fmt.Println("Anime content is playing")
}

type KDramaContent struct {
	name string
}

func (k KDramaContent) Play() {
	fmt.Println(k.name)
	fmt.Println("Korean Drama content is playing")
}

// studio
type ContentCreator interface {
	Produce(name string) Content
}

type Kappa struct{}

func (k *Kappa) Produce(name string) Content {
	return &AnimeContent{
		name: name,
	}
}

type NetflixKorea struct{}

func (n *NetflixKorea) Produce(name string) Content {
	return &KDramaContent{
		name: name,
	}
}

func main() {
	var contentCreator ContentCreator

	contentCreator = &Kappa{}
	content := contentCreator.Produce("attack on titan")
	content.Play()

	contentCreator = &NetflixKorea{}
	content = contentCreator.Produce("k-drama")
	content.Play()
}

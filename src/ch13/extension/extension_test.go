package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct { // 不支持 LSP, 不支持重载
	// Extension
	// p *Pet
	Pet
}

func (d *Dog) Speak() {
	fmt.Print("Wang !")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Yang")
}

package polymorphism_test

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct{}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World! GoProgrammer\")"
}

type JavaProgrammer struct{}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World! JavaProgrammer\")"
}

func writeFirstProgram(p Programmer) { // Programmer 是 interface，只能够接受指针
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld()) // %T 显示传入类型
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}

package empty_interface_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// if value, ok := p.(int); ok {
	// 	fmt.Println("Integer", value)
	// 	return
	// }
	// if value, ok := p.(string); ok {
	// 	fmt.Println("string", value)
	// 	return
	// }
	// fmt.Println("Unknow Type")

	switch value := p.(type) {
	case int:
		fmt.Println("Integer", value)
	case string:
		fmt.Println("string", value)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}

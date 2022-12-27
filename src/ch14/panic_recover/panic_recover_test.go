package panicrecover_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVsExit(t *testing.T) {
	defer func() {
		fmt.Println("Finally!")
	}()
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("recoverd from ", err)
	// 	}
	// }()

	fmt.Println("Start")
	panic(errors.New("something wrong !"))
	// os.Exit(-1) // show 'exit status 255'
	// fmt.Println("End")
}

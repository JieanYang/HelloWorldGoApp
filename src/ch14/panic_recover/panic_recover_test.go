package panicrecover_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVsExit(t *testing.T) {
	// defer func() {
	// 	fmt.Println("Finally!")
	// }()
	defer func() {
		// 恢复错误
		if err := recover(); err != nil {
			fmt.Println("recoverd from ", err) // 只是打印和记录并非正确的错误恢复方式
		}
	}()

	fmt.Println("Start")
	panic(errors.New("something wrong !"))
	// os.Exit(-1) // show 'exit status 255'
	// fmt.Println("End")
}

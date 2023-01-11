package client_test

import (
	"testing"

	"github.com/JieanYang/HelloWorldGoApp/src/ch15/series"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))
	// series.square(5) // Not support
}

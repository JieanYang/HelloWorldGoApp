package client_test

import (
	"src/freelancer_work/ansys/HelloWorldGoApp/src/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))
	// series.square(5) // Not support
}

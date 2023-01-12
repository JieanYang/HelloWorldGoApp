package flexible_reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}

	// fmt.Println(a == b) // Error: map can only be compared to nil
	// Correct writing
	// fmt.Println(a == nil)
	// fmt.Println(a != nil)
	fmt.Println(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	// fmt.Println(s1 == s2) // Error: slice can only be compared to nil
	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println("c1 == c2?", c1 == c2)
	fmt.Println("c1 == c2 DeepEqual ?", reflect.DeepEqual(c1, c2))
}

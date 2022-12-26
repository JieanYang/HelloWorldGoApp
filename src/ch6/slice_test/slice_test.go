package slicetest_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0)) // length, capacity
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	// t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
	t.Log(s2[0], s2[1], s2[2])

	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{
		"Januaray",
		"Feburary",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	Q2 := year[3:6]
	t.Log(len(Q2), cap(Q2))
	t.Log(Q2)

	summer := year[5:8]
	t.Log(len(summer), cap(summer))
	t.Log(summer)

	summer[0] = "Unknow"
	t.Log(Q2)   // Modified
	t.Log(year) // Modified
}

func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	t.Log(a, b)
	// if a == b { // a == b (slice can only be compared to nil)
	// 	t.Log("equal")
	// }
}

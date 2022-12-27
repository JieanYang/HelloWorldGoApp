package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Logf("len m1=%d", len(m1))
	t.Log(m1)
	t.Log(m1[2])

	m2 := map[int]int{}
	t.Logf("len m2=%d", len(m2))
	t.Log((m2))
	m2[4] = 16
	t.Log("len m2", len(m2))
	t.Logf("len m2=%d", len(m2))
	t.Log((m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
	// t.Logf("cap m3=%d", cap(m3)) // Not support
}

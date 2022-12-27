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

func TestAccessNotExistingKey(t *testing.T) { // How to know a key exists or not ?
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	// m1[3] = 0
	if value, isExist := m1[3]; isExist {
		t.Logf("Key 3's value is %d", value)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for key, value := range m1 {
		t.Logf("key is %d, value is %d", key, value)
	}
}

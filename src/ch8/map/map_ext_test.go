package map_ext

import "testing"

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	num := 1
	if mySet[num] { // Check if exist
		t.Logf("%d is existing", num)
	} else {
		t.Logf("%d is not existing", num)
	}
	mySet[2] = true
	t.Logf("The lenght of mySet is %d", len(mySet)) // Get Set length
	delete(mySet, 1)                                // Delete element of Set
	t.Logf("The lenght of mySet is %d", len(mySet))
	if mySet[num] { // Check if exist
		t.Logf("%d is existing", num)
	} else {
		t.Logf("%d is not existing", num)
	}
}

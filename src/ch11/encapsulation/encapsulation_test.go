package encapsulation_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e *Employee) StringPointer() string {
	fmt.Printf("StringPointer Addresse is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("StringPointer -> ID:%s,Name:%s,Age:%d\n", e.Id, e.Name, e.Age)
}
func (e Employee) StringObj() string {
	fmt.Printf("StringObj Addresse is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("StringObj -> ID:%s,Name:%s,Age:%d\n", e.Id, e.Name, e.Age)
}
func (e Employee) String() string {
	return fmt.Sprintf("String -> ID:%s,Name:%s,Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"

	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e is %T", &e)
	t.Logf("e2 is %T", e2)
}

func TestStructOperations(t *testing.T) {
	eObj := Employee{"0", "Bob", 20}
	ePointer := &Employee{"0", "Bob", 20}
	fmt.Printf("TestStructOperations Addresse is %x\n", unsafe.Pointer(&ePointer.Name))
	t.Log(eObj.StringObj())
	t.Log(ePointer.StringPointer()) // 避免内存拷贝，更少的空间开销
	t.Log(ePointer.StringObj())
}

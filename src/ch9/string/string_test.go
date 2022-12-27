package string_test

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化位默认零值 ""
	s = "hello"
	t.Log(len(s))
	t.Log("=============")

	// s[1] = '3'         // string 是不可变的 byte slice
	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	// s = "\xE4\xBA\xB5\xFF" // 任意二进制数据
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(s)
	t.Log(len(s))
	t.Log("=============")

	c := []rune(s) // rune 能够取出字符串里unicode
	t.Log(len(c))
	// t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
	t.Logf("中 UTF8 %s", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s { // c is rune type
		t.Logf("%[1]c %[1]x %[1]d", c)
	}
	for a, c := range s {
		t.Log(a, c)
	}
}

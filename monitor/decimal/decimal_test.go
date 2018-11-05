package decimal

import (
	"testing"
	"fmt"
)

func TestCompare(t *testing.T) {
	type compare_case struct{
		a string
		b string
		expected bool
	}

	var cases = []*compare_case{
		{"1", "2", true},
		{"5", "0", true},
		{"12", "12", true},
	}

	for _, c := range cases{
		res := CompareStr(c.a, c.b)
		fmt.Println(res)
	}
}

func TestSum(t *testing.T) {
	type compare_case []string

	var cases = []*compare_case{
		{"1","2","3"},
		{"0","0","0"},
	}
	for _, c := range cases{
		res := SumStr("0", *c...)
		fmt.Println(res)
	}
}
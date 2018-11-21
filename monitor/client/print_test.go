package main

import (
	"fmt"
	"testing"
)

func TestRedPrint(t *testing.T){
	fmt.Println("测试")
	fmt.Println("test")
	redPrint("test")
	redPrint("测试")
}
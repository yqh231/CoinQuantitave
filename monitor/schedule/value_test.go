package main

import(
	"testing"
	"fmt"
)

func TestTop(t *testing.T){
	fmt.Println(topPrice("bch_usdt"))
}

func TestCur(t *testing.T){
	fmt.Println(curPrice("bch_usdt"))
}


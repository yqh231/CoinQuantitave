package main

import (
	"fmt"
	"github.com/fatih/color"
)

const (
	line = 20 
)

func redPrint(content string){
	color.Red(content)
}

func err(){
	color.Red("error try again")
}

func empty(){
	color.Red("empty")
}

func blank(s string) (r string){
	l := len(s)
	for l < line{
		r += " "
		l += 1
	}
	return
}

func Fprint(args ...string){
	var output string
	for _, s := range args{
		output += s
		output += blank(s)
	}
	fmt.Println(output)
}

func RedFprint(args ...string){
	var output string
	for _, s := range args{
		output += s
		output += blank(s)
	}
	color.Red(output)
}
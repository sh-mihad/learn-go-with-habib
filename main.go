package main

// for using another variable and function you need to create package scop
// for package scop you need to write command: go mod init example.com //একানে example.com হচ্চছে ডোমেইন নেইম

import (
	"fmt"

	"example.com/mathlib"
)

func main(){
	fmt.Println("showing custom package")
	mathlib.Add(3,4)
}
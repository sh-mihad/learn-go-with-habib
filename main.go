package main

import "fmt"

var (
	a = 10
	b =20
)

func printNumber (num int){
	fmt.Println(num)
}

func add (x int,y int){
	res:=x+y
	printNumber(res)
}

func main(){
	add(a,b)
}

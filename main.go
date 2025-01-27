package main

import "fmt"

var (
	a = 10
	b =20
)

func printNumber (num int){
	fmt.Println(num)
}



func main(){
	add(a,b)
}

func add (x int,y int){
	res:=x+y
	printNumber(res)
}

// স্কোপ এর ক্ষেত্রে আমরা ফাংশন ডেফিনেশন এর অর্ডার  উলটাপালটা হলেও ফাংশন এক্সিকিউশন এর সময় ঠিক ঠাক ভাবে কাজ করবে

package main

import "fmt"

var a = 10

func main (){
	age:=30
	if age >=18 {
		a:=40 // এখানে a ভেরিয়েবল তার ব্লোক স্কোপে আলাদা ভাবে ডিফাইন হয়েছে যদিও উপরে a নামে একটা ভ্যারিয়েবল রয়েছে কিন্তূ সেটা এখানে রিডিক্লার হবে না । আর উপরে a থাকার কারনে  ব্লোক স্কোপের a কে "shadowing variable" বলা হয়
		fmt.Println(a)
	}
	fmt.Println(a)
} 
package variable

import "fmt"

var pkgTestVariable int = 150
var GlobalTestVar int = 90

// Define and Declare two integer variables in 3 different ways as discussed in the class!
// Define and Declare two  string variables in 3 different ways as discussed in the class!
func VariableFunc() {
	fmt.Println("-----------\n\nVariable declaration :\n")
	//Method 1- Declaring int
	l := 25
	m := 20

	fmt.Println("Method 1 of int declaration", l+m)

	//Method 2- Declaring int
	var n int
	var o int
	n = 10
	o = 12
	fmt.Println("Method 2 of int declaration", n+o)

	//Method 3- Declaring int
	var p int = 20
	var q int = 22
	fmt.Println("Method 3 of int declaration", p+q)

	//Method 1- Declaring string
	str1 := "Method 1 "
	str2 := "String variable declaration"
	fmt.Println("\nThis is ", str1+str2)

	//Method 2- Declaring string
	var str3, str4 string
	str3 = "Method 2 "
	str4 = "String variable declaration"
	fmt.Println("This is ", str3+" "+str4)

	//Method 3- Declaring string
	var str5 string = "Method 3 "
	var str6 string = "String variable declaration"
	fmt.Println("This is ", str5+str6)
}

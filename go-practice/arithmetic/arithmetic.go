package arithmetic

import (
	"fmt"
)

//Perform all possible arithmetic operations on integer, float variables
//calculate min value and maximum value that can be saved in float64 variable
// Try type conversions and observe behavior of those conversations(compiler error)

func ArithmeticOperationFunc() {
	//arithmetic operation on integer variables
	fmt.Println("-----------\n\nArithmetic operations :\n")
	var num int8
	num = 3
	fmt.Println("int8 variable ", num)
	var num1 int16
	num1 = 1523
	fmt.Println("int16 variable ", num1)
	var num2 int16
	num2 = 44
	var num5 int32
	num5 = 984654

	fmt.Println("\narithmetic operations on integer variables-")
	res1 := num1 + num2 // integer addition operation
	fmt.Println("Addition operation ", res1)
	res2 := num1 - num2 // integer subtraction operation
	fmt.Println("Subtraction operation ", res2)
	res3 := num1 * num2 // integer multiplication operation
	fmt.Println("Multiplication operation ", res3)
	res4 := num1 / num2 // integer division operation
	fmt.Println("Division operation ", res4)
	//

	//resx:=num+ num4 --> Throws an error because of size mismatch
	//fmt.Println(resx)

	//arithmetic operations on float variables
	var num3 float64
	num3 = 12.25
	fmt.Println("\nfloat32 value ", num3)

	var num4 float64
	num4 = 4848.148
	fmt.Println("float64 value ", num4)
	fmt.Println("\narithmetic operations on float variables-")
	res5 := num3 + num4 // float addition operation
	fmt.Println("Addition operation ", res5)
	res6 := num3 - num4 // float subtraction operation
	fmt.Println("Subtraction operation ", res6)
	res7 := num3 * num4 // float multiplication operation
	fmt.Println("Multiplication operation ", res7)
	res8 := num3 * num4 // float division operation
	fmt.Println("Division operation ", res8)

	res9 := float64(num5) + num4 // Addition by converting integer to float
	fmt.Println("\nfloat and integer addition ", res9)

}

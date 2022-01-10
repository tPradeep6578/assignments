package datatype

import "fmt"

/* Create variables of different data-types we discussed(bool, string, int8, int16, int32, int64).
Find out the number ranges for [int8, int16, int32, int64]
Demonstrate that you cannot assign, higher integers/lower integers than data type can support
*/
func DataTypeFunc() {
	fmt.Println("-----------\n\ndatatypes :\n")
	var a bool // Boolean variable datatype
	a = true
	//a="modify" -> cannot be assigned as it is of different datatype
	fmt.Println("Boolean variable:", a)

	var b string //  String datatype
	b = "string datatype"
	fmt.Println(b)

	var num int
	num = 27
	fmt.Println("Integer datatype:", num)

	//1 Byte = 8 bits
	var num1 int8 //2^8=256 combinations
	num1 = 127    //-(2^7)to 2^7-1 -->range is -128 to 127
	fmt.Println("This is int8 datatype:", num1)
	//num1 =128 --> higher integer not supported

	var num2 int16 //2^16 combinations
	num2 = 6454    //-(2^15) to 2^15-1 -> range = -32,768 to 32,767
	fmt.Println("This is int16 datatype:", num2)

	var num3 int32   //2^32 combinations
	num3 = 849812316 //-(2^31) to 2^31-1 -->range is -2,147,483,648 to 2,147,483,647
	fmt.Println("This is int32 datatype:", num3)

	var num4 int64        // 2^64 combinations
	num4 = 64654684986465 //-(2^63) to 2^63-1 --> range is -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
	fmt.Println("This is int64 datatype :", num4)

}

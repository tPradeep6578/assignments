package scope

import "fmt"

var pkgVariable int = 150 // package level variable starts with small letter
var GlobalVar int = 90    //global variable starts with capital letter

func ScopeDeclaration() {
	fmt.Println("-----------\n\nScope Variables :\n")
	var funcVariable int
	funcVariable = 5                                          //functional level variable - declared within the function and used within the function
	fmt.Println("functional level variable : ", funcVariable) //block scope variable
	fmt.Println("Package level variable", pkgVariable)        //using package level variable within the package
}
func ScopeTestFunc() {
	var funcVariable2 int
	funcVariable2 = 25
	//fmt.Println("local variable or functional level scope value: ", funcVariable)
	// the above line throws an error as local variable cannot be accessed outside of the function
	//fmt.Println("functional level scope : ", pkgTestVariable)
	// the above line throws an error as package level variable of variable module cannot be accessed outside of the package
	fmt.Println("functional level variable : ", funcVariable2)
	fmt.Println("global level variable", GlobalVar) //Using Global level variable anywhere
	//fmt.Println("global level variable", variable.GlobalTestVar)

}

/* Assignment-Create Variables in different scopes
Functional
Package
Global
Your code should contain examples that variables created in one function is NOT accessible in another function
Your code should contain examples that package variables are accessible ONLY inside package, NOT from another package
Your code should contain examples that global variables are accessible everywhere using <package_name>.<variable_name>
*/

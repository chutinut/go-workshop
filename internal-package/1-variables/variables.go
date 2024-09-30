package variable

import (
	"fmt"
)

func ShowVariables() {
	var name string                                // declare without initial value
	var fullName string = "SHEDKAE, THE DEVELOPER" // declare with initial value (string)
	var age int = 20                               // declare with initial value (number : integer)
	var email = "THE DEVELOPER"                    // declare without variable type (string)
	male := true                                   // shorthand declaration variable (boolean)
	const pi = 3.414                               // constant declaration variable (number : float)
	name = "SHEDKAE"                               // initial value or re-assignment value [line:8]
	fmt.Println("name is : ", name)
	fmt.Println("full name is : ", fullName)
	fmt.Println("email is : ", email)
	fmt.Println("age is : ", age)
	fmt.Println("My gender is male : ", male)
	fmt.Println("Pi value is : ", pi)
}

/*
Variable type (name : param)
1. string : string // text variable  (Eg. "ABC", "1", "")
2. boolean : bool // boolean variable (Eg. true, false)
3. number : int(8, 16, 32, 64) // number integer variable (Eg. 1, 50, 100, 999, -1, -5) *int(value bit size)
4. number : uint // number integer variable greater than or equal 0 (Eg. 1, 50, 100, 999)
5. number : float // number integer variable greater than or equal 0 (Eg. 1.2, 3.414, -0.25) *float(value bit size)
*/

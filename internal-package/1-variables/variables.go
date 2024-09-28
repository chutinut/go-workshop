package variable

import (
	"fmt"
)

func ShowVariables() {
	var name string
	var fullName string = "SHEDKAE, THE DEVELOPER"
	var email = "THE DEVELOPER"
	var age int = 20
	male := true
	pi := 3.414
	name = "SHEDKAE"
	fmt.Println("name is : ", name)
	fmt.Println("full name is : ", fullName)
	fmt.Println("email is : ", email)
	fmt.Println("age is : ", age)
	fmt.Println("My gender is male : ", male)
	fmt.Println("Pi value is : ", pi)
}

package controlStructure

import (
	"fmt"
)

/*
Condition
1. equal : == | // if a = b is true
2. not equal : != // if a = b is false
3. greater than : > // if a > b is true, if a = b is false
4. greater than or equal : >= // if a > b is true, if a = b is true
5. less than : < // if a > b is false, if a = b is false
6. less than or equal : < =// if a > b is false, if a = b is true
*/

// example of if else condition
func gradeCalculator(score float32) {
	var grade string
	if score == 100 { // with equal condition
		grade = "Top Score !!"
	} else if score >= 80 { // with greater than or equal condition
		grade = "Your grade is : A"
	} else if score >= 70 {
		grade = "Your grade is : B"
	} else if score >= 60 {
		grade = "Your grade is : C"
	} else if score >= 50 {
		grade = "Your grade is : D"
	} else if (score < 50) && (score > 30) { // multiple condition with less than and greater than condition
		grade = "Your grade is : F"
	} else { // else case (if non all above)
		grade = "You are fool !!!!"
	}
	fmt.Printf("Your score is %.2f | %s \n", score, grade)
}

// example of pre process in if else condition
// process value in if statement
func moreThanTarget(a int, b int) {
	target := 10
	// this statement will declare sum and assign value from a + b and use sum to compare with target
	if sum := a + b; sum > target {
		fmt.Printf("a = %d and b = %d | a + b is %d is more than target \n", a, b, a+b)
	} else {
		fmt.Printf("a = %d and b = %d | a + b is %d is less than target \n", a, b, a+b)
	}
}

// example of switch case
func getDayOfWeek(dayOfWeek int) {
	switch dayOfWeek {
	case 1:
		fmt.Printf("Your input day is Monday \n")
	case 2:
		fmt.Printf("Your input day is Tuesday \n")
	case 3:
		fmt.Printf("Your input day is Wednesday \n")
	case 4:
		fmt.Printf("Your input day is Thursday \n")
	case 5:
		fmt.Printf("Your input day is Friday \n")
	case 6:
		fmt.Printf("Your input day is Saturday \n")
	case 7:
		fmt.Printf("Your input day is Sunday \n")
	default:
		fmt.Printf("Your input day is Non all above \n")
	}
}

func ShowControlStructure() {
	gradeCalculator(67.5)
	moreThanTarget(10, 15)
	getDayOfWeek(2)
}

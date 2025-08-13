package main
import ("myproject/simplecalc"
"fmt"
)
func main() {
	
	fmt.Println("Hello World")
a,b := 6,4
fmt.Println(simplecalc.Add(a,b))
c, d := 10, 3
	fmt.Println(simplecalc.Sub(c, d))

	e, f := 20, 4
	fmt.Println(simplecalc.Div(e, f))

	x, y := 30, 3
	fmt.Println(simplecalc.Mul(x, y))
num := 6
	if num > 6 {
		fmt.Println("Number is greater than 6")
	} else if num == 6 {
		fmt.Println("Number is exactly 6")
	} else {
		fmt.Println("Number is less than 6")
}
day := "monday"
	switch day {
	case "monday":
		fmt.Println("Weekday")
	case "Saturday":
		fmt.Println("Weekend")
	default:
		fmt.Println("It's another day")
	}
	fmt.Println("Counting from 1 to 6:")
	for i := 1; i <= 6; i++ {
		fmt.Println(i)
}
}

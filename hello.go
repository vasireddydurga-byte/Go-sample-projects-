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

}

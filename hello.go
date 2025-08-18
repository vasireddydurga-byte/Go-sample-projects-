package main

import (
	"fmt"
	"myproject/simplecalc"
)

type address struct {
	address1 string
	address2 string
	state    string
	pin      int
}

func main() {{


	L:= make([]int,3 )
	L = append(L,3)
	L= append(L,4)
	L = append(L,5)
	L= append(L,6)
	fmt.Println(len(L))
	fmt.Println(cap(L))
K :=[5] int{20,30,40,50, 60}
slice := K[:4]
fmt.Println(slice)
}
 
	D := address{
		address1: "paseo laguna ",
		address2: "livermore",
		state:    "california",
		pin:      94451,
	}
	fmt.Println(D.address1)
	fmt.Println(D.address2)
	fmt.Println(D.state)
	fmt.Println(D.pin)
	fmt.Println(D)

fmt.Println("Hello World")
	a, b := 6, 4
	fmt.Println(simplecalc.Add(a, b))
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
	day := "wednesday"
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



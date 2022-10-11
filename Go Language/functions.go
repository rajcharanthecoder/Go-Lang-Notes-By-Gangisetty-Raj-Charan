package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good Morning %v \n", n)
}

func SayBye(n string) {
	fmt.Printf("Good Bye %v \n", n)
}

func cycleName(n []string, f func(string)) {
	for _, v := range n {
		f(v) //v is value and f is function
	}
}

func CircleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("RaJeSh")
	SayBye("ILoVeGo")
	cycleName([]string{"Rajesh", "Kumar", "Tapadia"}, sayGreeting)
	cycleName([]string{"Rajesh", "Kumar", "Tapadia"}, SayBye)

	area := CircleArea(5.5)
	fmt.Println(area)

	// math. functions

	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Abs(-2.7))
	fmt.Println(math.Max(2, 3))
	fmt.Println(math.Min(2, 3))
	fmt.Println(math.Mod(10, 3))
	fmt.Println(math.Remainder(10, 3))
	fmt.Println(math.Round(2.7))
	fmt.Println(math.RoundToEven(2.5))
	fmt.Println(math.RoundToEven(3.5))
	fmt.Println(math.Trunc(2.7))
	fmt.Println(math.Trunc(2.5))
	fmt.Println(math.Trunc(3.5))
	fmt.Println(math.Hypot(3, 4))
	fmt.Println(math.Hypot(5, 12))
	fmt.Println(math.Hypot(8, 15))
	fmt.Println(math.Pi)
	fmt.Println(math.E)
	fmt.Println(math.Phi)
	fmt.Println(math.NaN())
	fmt.Println(math.Inf(1))
	fmt.Println(math.Inf(-1))
	fmt.Println(math.IsNaN(math.NaN()))
	fmt.Println(math.IsInf(math.Inf(1), 1))
	fmt.Println(math.IsInf(math.Inf(-1), -1))

	fmt.Println(math.Cbrt(8))
	fmt.Println(math.Cbrt(27))
	fmt.Println(math.Cbrt(64))
	fmt.Println(math.Cbrt(125))
	fmt.Println(math.Cbrt(216))
	fmt.Println(math.Cbrt(343))

	fmt.Println(math.Acos(1))         //inverse of cos
	fmt.Println(math.Asin(1))         //inverse of sin
	fmt.Println(math.Atan(1))         //inverse of tan
	fmt.Println(math.Atan2(1, 1))     //inverse of tan2
	fmt.Println(math.Cos(1))          //cos
	fmt.Println(math.Sin(1))          //sin
	fmt.Println(math.Tan(1))          //tan
	fmt.Println(math.Copysign(1, -1)) //Copysign
	fmt.Println(math.Dim(1, 2))
	fmt.Println(math.Exp(1))
	fmt.Println(math.Expm1(1))
	fmt.Println(math.Frexp(1))
	fmt.Println(math.Frexp(2))
	fmt.Println(math.Frexp(3))
	fmt.Println(math.Frexp(4))
	fmt.Println(math.Frexp(5))
	fmt.Println(math.Frexp(6))
	fmt.Println(math.Frexp(7))
	fmt.Println(math.Gamma(1))

	fmt.Println(math.Copysign(1, -1))

	//Multiple Return Values//

}

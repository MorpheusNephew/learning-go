package main

import "fmt"

// Define a generic interface called Printable that matches a type that implements fmt.Stringer and has an underlying type of int or float64.
// Define types that meet this interface. Write a function that takes in a Printable and prints its value to the screen using fmt.Println.
func main() {
	myPrintableInt := PrintInt(13)
	myPrintableFloat := PrintFloat(5.03843)

	Print(myPrintableInt)
	Print(myPrintableFloat)
}

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintInt int

func (pi PrintInt) String() string {
	return fmt.Sprintf("%d", pi)
}

type PrintFloat float64

func (pf PrintFloat) String() string {
	return fmt.Sprintf("%f", pf)
}

func Print[T Printable](printable T) {
	fmt.Println(printable)
}

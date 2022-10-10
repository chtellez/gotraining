//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//

package main

import "fmt"

//--Requirements:
//* Mathematical operations must be defined as constants using iota

const (
	Add = iota
	Sub
    Mul
    Div
)

//* Write a receiver function that performs the mathematical operation
//  on two operands

type Operacion int

func (op Operacion) calculate(a, b int)int{
	switch op {
    case Add:
        return a + b
	case Sub:
		return a - b
	case Mul:
		return a * b
	case Div:
		return a / b
	default:
		panic("operacion not implemented")
	}
}

//* Operations required:
//  - Add, Subtract, Multiply, Divide

//* The existing function calls in main() represent the API and cannot be changed



func main() {
	add := Operacion(Add)
	fmt.Println(add.calculate(2, 2)) // = 4
	
	sub := Operacion(Sub)
	fmt.Println(sub.calculate(10, 3)) // = 7
	
	mul := Operacion(Mul)
	fmt.Println(mul.calculate(3, 3)) // = 9
	
	div := Operacion(Div)
	fmt.Println(div.calculate(100, 2)) // = 50
}

//
//--Notes:
//* Your program is complete when it compiles and prints the correct results
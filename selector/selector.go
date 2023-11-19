package selector

import (
	"bufio"
	"github.com/AndresMGdev/calculator/operators"
	"fmt"
	"os"
	"strings"
)

//Selector clase que maneja las operaciones
func Selector() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operation := scanner.Text()
	result := 0
	
	if strings.Contains(operation, "+") {
		result = operators.Addition(operation)
	} else if strings.Contains(operation, "-") {
		result = operators.Subtract(operation)
	} else if strings.Contains(operation, "*") {
		result = operators.Multiplication(operation)
	} else if strings.Contains(operation, "/") {
		result = operators.Division(operation)
	} else {
		fmt.Println("Error: esta ingresando un operador incorrecto, operadores permitidos (+, -, *, /)")
	}
	fmt.Println(result)
}
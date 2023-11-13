package operators

import (
	"fmt"
	"strconv"
	"strings"
)

//Subtract realiza la resta
func Subtract(operation string) int {
	values := strings.Split(operation, "-")
	result := 0

	for i := 0; i < len(values); i++ {
		num, err:= strconv.Atoi(values[i])
		if err != nil {
			fmt.Printf("Error %s ingresa un numero entero \n", err)
			fmt.Println("Solo se permite usar un operador en cada calculo ✔")
		} else {
			if result == 0 {
				result = num
			} else {
				result -= num
			}
		}
	}
	return result
}
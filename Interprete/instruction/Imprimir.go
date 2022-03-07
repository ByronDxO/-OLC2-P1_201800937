package instruction

import (
	// "fmt"
	"strconv"
	// "reflect"
	"OLC2/Interprete/interfaces"
)

type Println struct {
	Expresion interfaces.Expresion
}

func PRINTLN(val interfaces.Expresion) Println {
	exp := Println{val}
	return exp
}

func (p Println) Interpretar(env interface{}) interface{} {

	var result interfaces.Symbol
	result = p.Expresion.Interpretar(env)
	// var value string
	if result.Tipo == 0 {
		// entero
		return strconv.Itoa(result.Valor.(int)) + "\n"
	}
	

	return result.Valor.(string)
}

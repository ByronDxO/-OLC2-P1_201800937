package instruction

import (
	// "fmt"
	"strconv"
	// "reflect"
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/ast"
)

type Println struct {
	Expresion interfaces.Expresion
}

func PRINTLN(val interfaces.Expresion) Println {
	exp := Println{val}
	return exp
}

func (p Println) Interpretar(env interface{}, tree *ast.Arbol) interface{} {
	var value string
	value = ""
	var result interfaces.Symbol
	result = p.Expresion.Interpretar(env, tree)
	
	if result.Tipo == interfaces.INTEGER { // INTEGER	
		value := strconv.Itoa(result.Valor.(int)) + "\n"
		tree.AddCode(value)
	
	}else if result.Tipo == interfaces.FLOAT {		
		value := strconv.FormatFloat(result.Valor.(float64), 'f', 6, 64) + "\n"
		tree.AddCode(value)

	}else{
		value := result.Valor.(string)
		tree.AddCode(value)
	}

	return value
}

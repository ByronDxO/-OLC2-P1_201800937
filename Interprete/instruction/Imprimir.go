package instruction

import (
	"fmt"
	"strconv"
	// "reflect"
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/ast"
)

type Println struct {
	Expresion interfaces.Expresion
	Formato string
}

func PRINTLN(val interfaces.Expresion, formato string) Println {
	exp := Println{val, formato}
	return exp
}

func (p Println) Interpretar(env interface{}, tree *ast.Arbol) interface{} {
	var value string
	value = ""
	var result interfaces.Symbol
	result = p.Expresion.Interpretar(env, tree)
	getFormato(p.Formato)
	if(result.Tipo == interfaces.EXCEPTION){
		eTipo := result.Valor.(*ast.Exception).Tipo
		eDesc := result.Valor.(*ast.Exception).Descripcion
		
		value += fmt.Sprintf("%v", eTipo)
		value += " - "
		value += fmt.Sprintf("%v", eDesc)
		value += "\n"
		tree.AddCode(value)
		return value
	}

	if result.Tipo == interfaces.INTEGER { // INTEGER	
		value := strconv.Itoa(result.Valor.(int)) + "\n"
		tree.AddCode(value)
	
	}else if result.Tipo == interfaces.FLOAT {		
		value := strconv.FormatFloat(result.Valor.(float64), 'f', 6, 64) + "\n"
		tree.AddCode(value)

	}else if result.Tipo == interfaces.BOOLEAN {
		
		value += fmt.Sprintf("%v", result.Valor)
		value += "\n"
		tree.AddCode(value)
	}else{
		value := result.Valor.(string)
		value += "\n"
		tree.AddCode(value)
	}

	return value
}


func getFormato(formato string) {
	
	fmt.Println(len(formato))
}
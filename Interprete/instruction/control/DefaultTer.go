package control

import (

	"OLC2/Interprete/interfaces"
	// "OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
)

type DefaultTer struct {

	Instrucciones	interfaces.Expresion
}


func NewDefaultTer(instrucciones interfaces.Expresion) DefaultTer {
	instr := DefaultTer{instrucciones}
	return instr
}	



func (p DefaultTer) Interpretar(env interface{}, tree *ast.Arbol) interfaces.Symbol {
	
	var result interfaces.Symbol
	result = p.Instrucciones.Interpretar(env, tree)
		

 	return result
}
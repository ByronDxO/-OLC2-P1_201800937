package transferencia

import (
	"OLC2/Interprete/ast"
	"OLC2/Interprete/interfaces"
)


type Return struct {
	Expresion 	interfaces.Expresion
	Row 		int
	Column 		int
}

func NewReturn(espresion interfaces.Expresion, row int, column int) Return {
	instr := Return{espresion, row, column}
	return instr
}


func (p Return) Interpretar(env interface{}, tree *ast.Arbol) interface{} {
	var result interfaces.Symbol
	result = p.Expresion.Interpretar(env, tree)
	return result
}
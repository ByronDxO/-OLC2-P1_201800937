package transferencia

import (
	"OLC2/Interprete/ast"
	"OLC2/Interprete/interfaces"
)


type Break struct {
	Expresion 	interfaces.Expresion
	Row 		int
	Column 		int
}

func NewBreak(expresion interfaces.Expresion, row int, column int) Break {
	instr := Break{expresion, row, column}
	return instr
}


func (p Break) Interpretar(env interface{}, tree *ast.Arbol) interface{} {
	if p.Expresion != nil {
		return p
	}
	return p
}
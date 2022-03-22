package transferencia

import (
	"OLC2/Interprete/ast"
)


type Break struct {
	Row 	int
	Column 	int
}

func NewBreak(row int, column int) Break {
	instr := Break{row, column}
	return instr
}


func (p Break) Interpretar(env interface{}, tree *ast.Arbol) interface{} {

	return p
}
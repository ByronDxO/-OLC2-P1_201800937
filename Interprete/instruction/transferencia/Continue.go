package transferencia

import (
	"OLC2/Interprete/ast"
)


type Continue struct {
	Row 	int
	Column 	int
}

func NewContinue(row int, column int) Continue {
	instr := Continue{row, column}
	return instr
}


func (p Continue) Interpretar(env interface{}, tree *ast.Arbol) interface{} {

	return p
}
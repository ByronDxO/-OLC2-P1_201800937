package instruction

import (
	// "fmt"
	// "strconv"
	// "reflect"
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/ast"
)

type ListExpre struct {
	Expresion interfaces.Expresion

}

func NewListExpre(val interfaces.Expresion) ListExpre {
	exp := ListExpre{val}
	return exp
}

func (p ListExpre) Interpretar(env interface{}, tree *ast.Arbol) interfaces.Symbol {


	var newInstr interfaces.Symbol

	newInstr = p.Expresion.Interpretar(env, tree)
	
	return newInstr

}
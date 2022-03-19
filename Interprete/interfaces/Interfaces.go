package interfaces

import (
	"OLC2/Interprete/ast"
)

type Symbol struct {
	Id    string
	Tipo  TipoExpresion
	Valor interface{}
	IsMut bool
}

type Exception struct {
	Tipo 		string
	Descripcion string
}


type Expresion interface {
	Interpretar(env interface{}, tree *ast.Arbol) Symbol
}

type Instruction interface {
	Interpretar(env interface{},  tree *ast.Arbol) interface{}
}

package interfaces

import (
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
)

type Symbol struct {
	Id    string
	Tipo  TipoExpresion
	Valor interface{}
	IsMut bool
}

type SymbolFunction struct {
	Id    			string
	Tipo  			TipoExpresion
	Instrucciones	*arrayList.List
	Parametro		*arrayList.List
	IsMut	 		bool
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

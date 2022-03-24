package function

import (

	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/ast"
)

type ListParam struct {
	Id 	string
	Tipo 		interfaces.TipoExpresion

}

func NewListParam(val string, tipo interfaces.TipoExpresion) ListParam {
	exp := ListParam{val, tipo}
	return exp
}

func (p ListParam) Interpretar(env interface{}, tree *ast.Arbol) interface{} {

	return p
}
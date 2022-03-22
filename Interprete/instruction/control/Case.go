package control

import (

	"OLC2/Interprete/interfaces"
	// "OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
	
)

type Case struct {

	Condicion 		interfaces.Expresion
	ListaExpresion	*arrayList.List
	Instrucciones	*arrayList.List

}


func NewCase(cond interfaces.Expresion, listaExpresion *arrayList.List, instrucciones *arrayList.List) Case {
	instr := Case{cond, listaExpresion, instrucciones}
	return instr
}	



func (p Case) Interpretar(env interface{}, tree *ast.Arbol) interfaces.Symbol {
	
	var result interfaces.Symbol
	if p.Condicion != nil {
		result = p.Condicion.Interpretar(env, tree)
		return result
	
	}
	return result
}
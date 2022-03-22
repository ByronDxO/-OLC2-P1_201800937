package control

import (

	"OLC2/Interprete/interfaces"
	// "OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
	
)

type CaseTer struct {

	Condicion 		interfaces.Expresion
	ListaExpresion	*arrayList.List
	Instrucciones	interfaces.Expresion

}


func NewCaseTer(cond interfaces.Expresion, listaExpresion *arrayList.List, instrucciones interfaces.Expresion) CaseTer {
	instr := CaseTer{cond, listaExpresion, instrucciones}
	return instr
}	



func (p CaseTer) Interpretar(env interface{}, tree *ast.Arbol) interfaces.Symbol {
	
	var result interfaces.Symbol
	if p.Condicion != nil {
		result = p.Condicion.Interpretar(env, tree)
		return result
	
	}
	return result
}
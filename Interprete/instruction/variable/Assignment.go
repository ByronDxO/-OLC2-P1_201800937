package variable

import (
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
)


type Assignment struct {
	Id 			string
	Expresion	interfaces.Expresion
}


func NewAssignment(id string, val interfaces.Expresion) Assignment {
	instr := Assignment{id, val}
	return instr
}


func (p Assignment) Interpretar(env interface{}, tree *ast.Arbol) interface{} {

	/* Buscar si el id ya existe */
	symbol := env.(environment.Environment).GetSymbol(p.Id)

	if symbol.Tipo == interfaces.NULL {
		excep := ast.NewException("Semantico","No Existe ese Id "+p.Id)
		tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No Existe ese Id "+ p.Id})
		return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

	}
	
	var result interfaces.Symbol
	result = p.Expresion.Interpretar(env, tree)
	env.(environment.Environment).SetSymbol(p.Id, result)
	
	return result.Valor

}
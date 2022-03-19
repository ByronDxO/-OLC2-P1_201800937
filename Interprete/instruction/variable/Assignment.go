package variable

import (
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
	"fmt"
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
	/* ERROR */
	var value string = ""

	/* Buscar si el id ya existe */
	symbol := env.(environment.Environment).GetSymbol(p.Id)

	if symbol.Tipo == interfaces.NULL {
		excep := ast.NewException("Semantico","No Existe ese Id "+p.Id)
		tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No Existe ese Id "+ p.Id})

		eTipo := excep.Tipo
		eDesc := excep.Descripcion
		
		value += fmt.Sprintf("%v", eTipo)
		value += " - "
		value += fmt.Sprintf("%v", eDesc)
		value += "\n"
		tree.AddCode(value)
		return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

	}
	
	var result interfaces.Symbol
	result = p.Expresion.Interpretar(env, tree)

	if symbol.IsMut {
		env.(environment.Environment).SetSymbol(p.Id, result)
	}else {
		excep := ast.NewException("Semantico","No se puede asignar a " + p.Id + ", no es mutable.")
		tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No se puede asignar a " + p.Id + ", no es mutable."})
		eTipo := excep.Tipo
		eDesc := excep.Descripcion
		value += fmt.Sprintf("%v", eTipo)
		value += " - "
		value += fmt.Sprintf("%v", eDesc)
		value += "\n"
		tree.AddCode(value)
		return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}
	}
	
	
	return result.Valor

}
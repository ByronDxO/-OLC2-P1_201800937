package variable

import (
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
)


type Identifier struct {
	Id string
}

func NewIdentifier(id string) Identifier {
	instr := Identifier{id}
	return instr
}


func (p Identifier) Interpretar(env interface{}, tree *ast.Arbol) interfaces.Symbol {

	symbol := env.(environment.Environment).GetSymbol(p.Id)

	if symbol.Tipo == interfaces.NULL {
		excep := ast.NewException("Semantico","No Existe ese Id "+p.Id)
		tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No Existe ese Id "+ p.Id})
		return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

	}

	return interfaces.Symbol{
		Id: symbol.Id,
		Tipo: symbol.Valor.(interfaces.Symbol).Tipo,
		Valor: symbol.Valor.(interfaces.Symbol).Valor,
	}

}
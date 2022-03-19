package variable



import (

	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
)

type Declaration struct {
	Id 			string
	Tipo 		interfaces.TipoExpresion
	Expresion 	interfaces.Expresion
	IsArray 	bool
	IsStruct 	bool
}

func NewDeclaration(id string, tipo interfaces.TipoExpresion, val interfaces.Expresion, isArray bool, isStruct bool) Declaration {
	instr := Declaration{id, tipo, val, isArray, isStruct}
	return instr
}


func (p Declaration) Interpretar(env interface{}, tree *ast.Arbol) interface{} {

	/* Buscar si el id ya existe */
	symbol := env.(environment.Environment).GetSymbol(p.Id)

	if symbol.Tipo != interfaces.NULL {
		excep := ast.NewException("Semantico","Ya Existe ese Id "+p.Id)
		tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Ya Existe ese Id "+ p.Id})
		return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

	}

	var result interfaces.Symbol
	result = p.Expresion.Interpretar(env, tree)

	
	if (result.Tipo == p.Tipo || p.Tipo == interfaces.NULL) {
		env.(environment.Environment).AddSymbol(p.Id, result, result.Tipo)
	
	/*}else if p.IsArray {
		env.(environment.Environment).AddSymbol(p.Id, result, interfaces.ARRAY)

	} else if p.IsStruct {
		env.(environment.Environment).AddSymbol(p.Id, result, interfaces.STRUCT)*/
	}else {
		excep := ast.NewException("Semantico","Tipo incorrecto en Declaracion.")
		tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Tipo incorrecto en Declaracion."})
		return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}
	}


	return result.Valor

}
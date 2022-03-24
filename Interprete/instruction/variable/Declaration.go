package variable



import (

	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
	"fmt"
)

type Declaration struct {
	Id 			string
	Tipo 		interfaces.TipoExpresion
	Expresion 	interfaces.Expresion
	IsMut		bool
	IsArray 	bool
	IsStruct 	bool
	Row			int
	Column		int
}

func NewDeclaration(id string, tipo interfaces.TipoExpresion, val interfaces.Expresion, isMut bool, isArray bool, isStruct bool, row int, column int) Declaration {
	instr := Declaration{id, tipo, val, isMut, isArray, isStruct, row, column}
	return instr
}


func (p Declaration) Interpretar(env interface{}, tree *ast.Arbol) interface{} {

	/* ERROR */
	var value string = ""
	fmt.Println("estamos aca perro")
	/* Buscar si el id ya existe */
	symbol := env.(environment.Environment).GetSymbol(p.Id)

	if symbol.Tipo != interfaces.NULL {
		excep := ast.NewException("Semantico","Ya Existe ese Id "+p.Id, p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Ya Existe ese Id "+ p.Id, Row: p.Row, Column: p.Column})
		return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

	}

	var result interfaces.Symbol
	if p.Expresion != nil {
		result = p.Expresion.Interpretar(env, tree)
		p.Tipo = result.Tipo
	}else{
		fmt.Println("oye es declaracion vacia")
		result.Tipo = p.Tipo
	}

	
	if (result.Tipo == p.Tipo || p.Tipo == interfaces.NULL) {
		env.(environment.Environment).AddSymbol(p.Id, result, result.Tipo, p.IsMut)
	
	/*}else if p.IsArray {
		env.(environment.Environment).AddSymbol(p.Id, result, interfaces.ARRAY)

	} else if p.IsStruct {
		env.(environment.Environment).AddSymbol(p.Id, result, interfaces.STRUCT)*/
	}else {
		excep := ast.NewException("Semantico","Tipo incorrecto en Declaracion.", p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Tipo incorrecto en Declaracion.", Row: p.Row, Column: p.Column})
		
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
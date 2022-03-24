package function

import (
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/ast"
	"OLC2/Interprete/instruction"
	"OLC2/Interprete/instruction/transferencia"
	arrayList "github.com/colegno/arraylist"
	"fmt"
	"reflect"	
	"OLC2/Interprete/environment"
)

type LlamadaExpre struct {

	Id 				string
	Parametro 	    *arrayList.List
	Row				int
	Column			int

}

func NewLlamadaExpre(Id string, Parametro *arrayList.List, Row int, Column int) LlamadaExpre {
	instr := LlamadaExpre{Id, Parametro, Row, Column}
	return instr
}

func (p LlamadaExpre) Interpretar(env interface{}, tree *ast.Arbol) interfaces.Symbol {

	

	var newTable environment.Environment
	newTable = environment.NewEnvironment(env.(environment.Environment))
	
	fnction := env.(environment.Environment).GetFunction(p.Id) 

	instr := fnction.Valor.(interfaces.Symbol).Valor.(interfaces.SymbolFunction)

	if fnction.Id != p.Id {
		excep := ast.NewException("Semantico","No Existe esa Funcion "+p.Id, p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

	}


	if p.Parametro.Len() == instr.Parametro.Len() {
		// contador := 0
		

		for s := 0; s < instr.Parametro.Len(); s++ {
			instrFun := instr.Parametro.GetValue(s).(ListParam)
			fmt.Println(instrFun.Id)
			fmt.Print("-")
			callFunc := p.Parametro.GetValue(s).(instruction.ListExpre).Interpretar(newTable, tree)

			if callFunc.Tipo == instrFun.Tipo {
				fmt.Println("SEMANTICO - TIPOS CORRECTOS")
				newTable.AddSymbol(instrFun.Id, callFunc, callFunc.Tipo, true)

			} else {
				excep := ast.NewException("Semantico","Tipo de Parametros incorrectos en llamada "+p.Id, p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}
			}
			
			
		}

		

	}else {
		excep := ast.NewException("Semantico","Parametros incorrectos en Funcion "+p.Id, p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}
	}

	var symbol interfaces.Symbol
	
	fmt.Println("$$$$$$$$$")
	for s := 0; s < instr.Instrucciones.Len(); s++ {

		fmt.Println(reflect.TypeOf(instr.Instrucciones.GetValue(s)).String())
		if reflect.TypeOf(instr.Instrucciones.GetValue(s)).String() == "transferencia.Return" {
			
			symbol = instr.Instrucciones.GetValue(s).(transferencia.Return).Expresion.Interpretar(newTable, tree )

			
			return symbol
		}

		instr.Instrucciones.GetValue(s).(interfaces.Instruction).Interpretar(newTable, tree)


	}

	fmt.Println("---TIPO---")
	fmt.Println(instr.Tipo)
	

	return interfaces.Symbol{
		Id: "",
		Tipo: interfaces.NULL,
		Valor: nil,
	}
	
}
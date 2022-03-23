package control

import (

	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
	"reflect"
	// "fmt"
)


type If struct {

	Condicion	interfaces.Expresion
	InstrIf		*arrayList.List
	InstrElse	*arrayList.List
	InstrElseIf *arrayList.List
	Row 		int

}


func NewIf(condicion interfaces.Expresion, instrIf *arrayList.List, instrElse *arrayList.List, instrElseIf *arrayList.List, row int) If {
	instr := If{condicion, instrIf, instrElse, instrElseIf, row}
	return instr
}


func (p If) Interpretar(env interface{}, tree *ast.Arbol) interface{} {

	var cond interfaces.Symbol
	cond = p.Condicion.Interpretar(env, tree)

	if cond.Tipo == interfaces.EXCEPTION {
		return cond
	} 

	if cond.Tipo == interfaces.BOOLEAN {

		if cond.Valor.(bool) {

			var newTable environment.Environment
			newTable = environment.NewEnvironment(env.(environment.Environment))

			for _, s := range p.InstrIf.ToArray() {
				newInstr := s.(interfaces.Instruction).Interpretar(newTable, tree)
				
				if reflect.TypeOf(newInstr).String() == "transferencia.Break"	 { return newInstr }
				if reflect.TypeOf(newInstr).String() == "transferencia.Continue" { return newInstr }
				if reflect.TypeOf(newInstr).String() == "transferencia.Return"   { return newInstr }

			}
	

		}else {
	
			if p.InstrElse != nil {

				var newTable environment.Environment
				newTable = environment.NewEnvironment(env.(environment.Environment))

				for _, s := range p.InstrElse.ToArray() {
					newInstr := s.(interfaces.Instruction).Interpretar(newTable, tree)

					if reflect.TypeOf(newInstr).String() == "transferencia.Break"	 { return newInstr }
					if reflect.TypeOf(newInstr).String() == "transferencia.Continue" { return newInstr }
					if reflect.TypeOf(newInstr).String() == "transferencia.Return"   { return newInstr }
					
				}
				
			}

			if p.InstrElseIf != nil {

				for _, s := range p.InstrElseIf.ToArray() {
					newInstr := s.(interfaces.Instruction).Interpretar(env, tree)
				
					if reflect.TypeOf(newInstr).String() == "transferencia.Break"	 { return newInstr }
					if reflect.TypeOf(newInstr).String() == "transferencia.Continue" { return newInstr }
					if reflect.TypeOf(newInstr).String() == "transferencia.Return"   { return newInstr }

				}
			}

		}
	} else {

		excep := ast.NewException("Semantico","Tipo de Dato no Booleano en If.", p.Row, p.Row)
		tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
		return excep

	}


	return cond.Valor
}	
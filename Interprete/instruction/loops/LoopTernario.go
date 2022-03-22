package loops

import (

	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/instruction/transferencia"
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
	"reflect"
	// "fmt"
)

type LoopTernario struct {
	Instrucciones	*arrayList.List
}

func NewLoopTernario(instrucciones	*arrayList.List) LoopTernario {
	instr := LoopTernario{instrucciones}
	return instr
}

func (p LoopTernario) Interpretar(env interface{}, tree *ast.Arbol) interfaces.Symbol {
	
	var result interfaces.Symbol
	for true {

	
		var newTable environment.Environment
		newTable = environment.NewEnvironment(env.(environment.Environment))

		for _, s := range p.Instrucciones.ToArray() {
			newInstr := s.(interfaces.Instruction).Interpretar(newTable, tree)
			
			if reflect.TypeOf(newInstr).String() == "transferencia.Break" 	 { 
				result = newInstr.(transferencia.Break).Expresion.Interpretar(newTable, tree)
				return result 
			}
			if reflect.TypeOf(newInstr).String() == "transferencia.Continue" { break }
			// if reflect.TypeOf(newInstr).String() == "transferencia.Return"   { return newInstr }

		}

	

		

	}

	return result
}
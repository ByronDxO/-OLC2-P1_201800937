package loops

import (

	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	// "OLC2/Interprete/instruction/transferencia"
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
	"reflect"
	"fmt"
)

type Loop struct {
	Instrucciones	*arrayList.List
}

func NewLoop(instrucciones	*arrayList.List) Loop {
	instr := Loop{instrucciones}
	return instr
}

func (p Loop) Interpretar(env interface{}, tree *ast.Arbol) interface{} {
	fmt.Println("entro al loop")
	var result interfaces.Symbol
	for true {

	
		var newTable environment.Environment
		newTable = environment.NewEnvironment(env.(environment.Environment))

		for _, s := range p.Instrucciones.ToArray() {
			newInstr := s.(interfaces.Instruction).Interpretar(newTable, tree)
			fmt.Println(reflect.TypeOf(newInstr))
			if reflect.TypeOf(newInstr).String() == "transferencia.Break" 	 { return newInstr }
			if reflect.TypeOf(newInstr).String() == "transferencia.Continue" { break }
			if reflect.TypeOf(newInstr).String() == "transferencia.Return"   { return newInstr }

		}

	

		

	}

	return result
}
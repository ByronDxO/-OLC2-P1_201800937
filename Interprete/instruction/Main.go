package instruction


import (
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)


type Main struct {
	Instrucciones   *arrayList.List
	Row 			int
	Column			int
}


func NewMain(instrucciones *arrayList.List, row int, column int) Main {
	instr := Main{instrucciones, row, column}
	return instr
}


func (p Main) Interpretar(env interface{}, tree *ast.Arbol) interface{} {

	var newTable environment.Environment
	newTable = environment.NewEnvironment(env.(environment.Environment))

	if p.Instrucciones != nil {
		for _, s := range p.Instrucciones.ToArray() {
			newInstr := s.(interfaces.Instruction).Interpretar(newTable, tree)

			if reflect.TypeOf(newInstr).String() == "transferencia.Break" 	 { 
				excep := ast.NewException("Semantico","Sentencia Break fuera de Ciclo.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				return excep
			}
			if reflect.TypeOf(newInstr).String() == "transferencia.Continue" { 
				excep := ast.NewException("Semantico","Sentencia Continue fuera de Ciclo.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				return excep
			}
			if reflect.TypeOf(newInstr).String() == "transferencia.Return"   { 
				excep := ast.NewException("Semantico","Sentencia Return fuera de Ciclo.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				return excep 
			}

		}
	}

	return nil
}
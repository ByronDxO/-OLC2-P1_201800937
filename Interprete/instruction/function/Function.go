package function

import (
	"OLC2/Interprete/environment"
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)

type Function struct {

	Id 				string
	Parametro 	    *arrayList.List
	Instrucciones 	*arrayList.List
	Tipo			interfaces.TipoExpresion
	Row				int
	Column			int

}

func NewFunction(Id string, Parametro *arrayList.List, Instrucciones *arrayList.List, tipo interfaces.TipoExpresion, Row int, Column int) Function {
	instr := Function{Id, Parametro, Instrucciones, tipo, Row, Column}
	return instr
}

func (p Function) Interpretar(env interface{}, tree *ast.Arbol) interface{} {

	var newTable environment.Environment
	newTable = environment.NewEnvironment(env.(environment.Environment))

	if p.Instrucciones != nil {
		for _, s := range p.Instrucciones.ToArray() {
			
			if reflect.TypeOf(s).String() == "transferencia.Break" 	 { 
				excep := ast.NewException("Semantico","Sentencia Break fuera de Ciclo.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				return excep
			}
			if reflect.TypeOf(s).String() == "transferencia.Continue" { 
				excep := ast.NewException("Semantico","Sentencia Continue fuera de Ciclo.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				return excep
			}
			if reflect.TypeOf(s).String() == "transferencia.Return"   { 
				excep := ast.NewException("Semantico","Sentencia Return fuera de Ciclo.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				return excep 
			}
			
			s.(interfaces.Instruction).Interpretar(newTable, tree)
			
			
			

		}
	}
	
	return nil
}
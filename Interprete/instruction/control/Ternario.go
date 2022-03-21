package control

import (

	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
	
)


type Ternario struct {

	Condicion	interfaces.Expresion
	InstrIf		interfaces.Expresion
	InstrElse	interfaces.Expresion
	InstrElseIf *arrayList.List
	Row 		int

}


func NewTernario(condicion interfaces.Expresion, instrIf interfaces.Expresion, instrElse interfaces.Expresion, instrElseIf *arrayList.List, row int) Ternario {
	instr := Ternario{condicion, instrIf, instrElse, instrElseIf, row}
	return instr
}


func (p Ternario) Interpretar(env interface{}, tree *ast.Arbol) interfaces.Symbol {

	var cond interfaces.Symbol
	cond = p.Condicion.Interpretar(env, tree)

	if cond.Tipo == interfaces.EXCEPTION {
		return cond
	} 

	if cond.Tipo == interfaces.BOOLEAN {

		if cond.Valor.(bool) {

			var newTable environment.Environment
			newTable = environment.NewEnvironment(env.(environment.Environment))

			var result interfaces.Symbol
			result = p.InstrIf.Interpretar(newTable, tree)

			return result
			
	

		}else {
	
			if p.InstrElse != nil {

				var newTable environment.Environment
				newTable = environment.NewEnvironment(env.(environment.Environment))

				var result interfaces.Symbol
				result = p.InstrElse.Interpretar(newTable, tree)
				return result
				
				
			}

			if p.InstrElseIf != nil {

				var result interfaces.Symbol
				for _, s := range p.InstrElseIf.ToArray() {
					
					result = s.(Ternario).Interpretar(env, tree)

				}

				return result
			}

		}
	} else {

		excep := ast.NewException("Semantico","Tipo de Dato no Booleano en If.", p.Row, p.Row)
		tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Tipo de Dato no Booleano en If.", Row: p.Row, Column: p.Row})
		return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

	}


	return cond
}	
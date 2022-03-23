package loops

import (

	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
	"reflect"
	"fmt"
)


type For struct {
	Id 	 	  		string
	Tipo 	  		interfaces.TipoExpresion
	Left	  		interfaces.Expresion
	Right     		interfaces.Expresion
	Instrucciones   *arrayList.List
	ListArray 		*arrayList.List
	Row	      		int
	Column    		int
}


func NewFor(id string, tipo interfaces.TipoExpresion, left interfaces.Expresion, right interfaces.Expresion, instrucciones *arrayList.List, listArray *arrayList.List, row int, column int) For {
	instr := For{id, tipo, left, right, instrucciones, listArray, row, column}
	return instr
}


func (p For) Interpretar(env interface{}, tree *ast.Arbol) interface{} {

	var newTable, newTable2 environment.Environment
	newTable = environment.NewEnvironment(env.(environment.Environment))
	var cond bool = false

	if p.Tipo == interfaces.INTEGER {
		var left, right, temp interfaces.Symbol
		left = p.Left.Interpretar(newTable, tree)
		if left.Tipo == interfaces.EXCEPTION { return left }

		right = p.Right.Interpretar(newTable, tree)
		if right.Tipo == interfaces.EXCEPTION { return right }
		
		left.IsMut = true
		right.IsMut = true
		
		if left.Tipo == interfaces.INTEGER && right.Tipo == interfaces.INTEGER{
			
			
			symbol := env.(environment.Environment).GetSymbol(p.Id)

			if symbol.Tipo == interfaces.NULL {
				cond = true
				env.(environment.Environment).AddSymbol(p.Id, left, left.Tipo, true)
				symbol = env.(environment.Environment).GetSymbol(p.Id)
				fmt.Println(symbol)
			}else {
				temp = symbol.Valor.(interfaces.Symbol)
			}


			for {

				newTable2 = environment.NewEnvironment(newTable)

				if left.Valor.(int) < right.Valor.(int){
					env.(environment.Environment).SetSymbol(p.Id, left, true)
					for _, s := range p.Instrucciones.ToArray() {
						newInstr := s.(interfaces.Instruction).Interpretar(newTable2, tree)
	
						if reflect.TypeOf(newInstr).String() == "transferencia.Break" 	 { return nil }
						if reflect.TypeOf(newInstr).String() == "transferencia.Continue" { break }
						if reflect.TypeOf(newInstr).String() == "transferencia.Return"   { return newInstr }
	
					}
					left.Valor = left.Valor.(int) + 1
					env.(environment.Environment).SetSymbol(p.Id, left, true)
				}else {
					break
				}

			}
			
			if !cond {
				env.(environment.Environment).SetSymbol(p.Id, temp, true)
			}

		}else {
			excep := ast.NewException("Semantico","Tipo de Dato incorrecto en For (int).", p.Row, p.Column)
			tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
			return excep
		}



	} else if p.Tipo == interfaces.STRING {
		
		var left, temp interfaces.Symbol
		var aux string = ""
		left = p.Left.Interpretar(newTable, tree)
		if left.Tipo == interfaces.EXCEPTION { return left }
		
		left.IsMut = true

		if left.Tipo == interfaces.STRING {

			symbol := env.(environment.Environment).GetSymbol(p.Id)

			if symbol.Tipo == interfaces.NULL {
				cond = true
				env.(environment.Environment).AddSymbol(p.Id, left, left.Tipo, true)
				symbol = env.(environment.Environment).GetSymbol(p.Id)
				fmt.Println(symbol)
			}else {
				temp = symbol.Valor.(interfaces.Symbol)
			}


			for i := 0; i < len(left.Valor.(string)); i++ {
				
				newTable2 = environment.NewEnvironment(newTable)
				aux = fmt.Sprintf("%v", left.Valor)
				
				env.(environment.Environment).SetSymbol(p.Id, interfaces.Symbol{p.Id,left.Tipo,string(aux[i]),true}, true)

				for _, s := range p.Instrucciones.ToArray() {
					newInstr := s.(interfaces.Instruction).Interpretar(newTable2, tree)

					if reflect.TypeOf(newInstr).String() == "transferencia.Break" 	 { return nil }
					if reflect.TypeOf(newInstr).String() == "transferencia.Continue" { break }
					if reflect.TypeOf(newInstr).String() == "transferencia.Return"   { return newInstr }

				}

			}

			if !cond {
				env.(environment.Environment).SetSymbol(p.Id, temp, true)
			}

		}else {
			excep := ast.NewException("Semantico","Tipo de Dato incorrecto en For (string).", p.Row, p.Column)
			tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
			return excep
		}

	}

	return nil
}
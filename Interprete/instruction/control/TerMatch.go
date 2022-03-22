package control

import (

	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
	"fmt"
	"reflect"
)


type TerMatch struct {

	Condicion	 interfaces.Expresion
	InstrCase	 *arrayList.List
	InstrDefault interfaces.Expresion
	Row 		 int

}

func NewTerMatch(cond interfaces.Expresion, instrCase *arrayList.List, instrDefault interfaces.Expresion, row int) TerMatch {
	instr := TerMatch{cond, instrCase, instrDefault, row}
	return instr
}

func (p TerMatch) Interpretar(env interface{}, tree *ast.Arbol) interfaces.Symbol {
	
	
	var result interfaces.Symbol
	var boolCond bool = false
	var cond, newCond interfaces.Symbol
	cond = p.Condicion.Interpretar(env, tree)
	
	if p.InstrCase != nil && p.InstrDefault != nil { // [<CASES_LIST>] [<DEFAULT>]
		fmt.Println("entro")
		for _, s := range p.InstrCase.ToArray() {
			
			if s.(CaseTer).Condicion != nil { // SOLO 1 CONDICION DE |
				fmt.Println("conde 1")
				newCond = s.(CaseTer).Condicion.Interpretar(env, tree)
				fmt.Println("conde 1")
				if newCond.Valor == cond.Valor {
					if !boolCond {
						
						boolCond = true
						 
						result = s.(CaseTer).Instrucciones.Interpretar(env, tree)
						return result

					} else{
						ast.NewException("Semantico","Expresion repetida en Match.", p.Row, p.Row)
						tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Expresion repetida en Matc.", Row: p.Row, Column: p.Row})
						
					}
					
				}

			}else{
			
				fmt.Println("list 1")
				for _, i := range s.(CaseTer).ListaExpresion.ToArray() {
						fmt.Println(reflect.TypeOf(i))
						newCond = i.(CaseTer).Condicion.Interpretar(env, tree)
						fmt.Println("list -")
						if newCond.Valor == cond.Valor {
							if !boolCond {

								boolCond = true
								 
								result = s.(CaseTer).Instrucciones.Interpretar(env, tree)
								return result

							} else{
								ast.NewException("Semantico","Expresion repetida en Match.", p.Row, p.Row)
								tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Expresion repetida en Matc.", Row: p.Row, Column: p.Row})
						
							}
						}		
				}
			}
		}

		if !boolCond {
			
			result =  p.InstrDefault.Interpretar(env, tree) 
			return result
			
			
		}

	}else if p.InstrCase != nil && p.InstrDefault == nil { // [<CASES_LIST>] 

		for _, s := range p.InstrCase.ToArray() {
			
			if s.(CaseTer).Condicion != nil { // SOLO 1 CONDICION DE |
				newCond = s.(CaseTer).Condicion.Interpretar(env, tree)

				if newCond.Valor == cond.Valor {
					if !boolCond {
						
						boolCond = true
						 
						result = s.(CaseTer).Instrucciones.Interpretar(env, tree)
						return result

					} else{
						ast.NewException("Semantico","Expresion repetida en Match.", p.Row, p.Row)
						tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Expresion repetida en Matc.", Row: p.Row, Column: p.Row})
						
					}
				}

			}else{

				for _, i := range s.(CaseTer).ListaExpresion.ToArray() {
						
						newCond = i.(CaseTer).Condicion.Interpretar(env, tree)

						if newCond.Valor == cond.Valor {
							if !boolCond {

								boolCond = true
								 
								result = s.(CaseTer).Instrucciones.Interpretar(env, tree)
								return result

							} else{
								ast.NewException("Semantico","Expresion repetida en Match.", p.Row, p.Row)
								tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Expresion repetida en Matc.", Row: p.Row, Column: p.Row})
						
							}			
						}
				}
			}
		}


	}else if p.InstrCase == nil && p.InstrDefault != nil { // [<DEFAULT>]
	
		
			result =  p.InstrDefault.Interpretar(env, tree) 
			return result
	}

	return result
}



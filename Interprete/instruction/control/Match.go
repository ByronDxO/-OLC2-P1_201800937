package control

import (

	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
	"fmt"
)


type Match struct {

	Condicion	 interfaces.Expresion
	InstrCase	 *arrayList.List
	InstrDefault *arrayList.List
	Row 		 int

}

func NewMatch(cond interfaces.Expresion, instrCase *arrayList.List, instrDefault *arrayList.List, row int) Match {
	instr := Match{cond, instrCase, instrDefault, row}
	return instr
}

func (p Match) Interpretar(env interface{}, tree *ast.Arbol) interface{} {
	fmt.Println("entro a match")
	var newTable environment.Environment
	newTable = environment.NewEnvironment(env.(environment.Environment))

	var boolCond bool = false
	var cond, newCond interfaces.Symbol
	cond = p.Condicion.Interpretar(env, tree)
	
	if p.InstrCase != nil && p.InstrDefault != nil { // [<CASES_LIST>] [<DEFAULT>]

		for _, s := range p.InstrCase.ToArray() {
			
			if s.(Case).Condicion != nil { // SOLO 1 CONDICION DE |
				newCond = s.(Case).Condicion.Interpretar(env, tree)

				if newCond.Valor == cond.Valor {
					if !boolCond {
						
						boolCond = true
						for _, j := range s.(Case).Instrucciones.ToArray() {
							j.(interfaces.Instruction).Interpretar(newTable, tree)
							
						} 
					} else{
						ast.NewException("Semantico","Expresion repetida en Match.", p.Row, p.Row)
						tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Expresion repetida en Matc.", Row: p.Row, Column: p.Row})
						
					}
					
				}

			}else{
			
				
				for _, i := range s.(Case).ListaExpresion.ToArray() {
						
						newCond = i.(Case).Condicion.Interpretar(env, tree)

						if newCond.Valor == cond.Valor {
							if !boolCond {

								boolCond = true
								for _, j := range s.(Case).Instrucciones.ToArray() {
									j.(interfaces.Instruction).Interpretar(newTable, tree)
									
								} 

							} else{
								ast.NewException("Semantico","Expresion repetida en Match.", p.Row, p.Row)
								tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Expresion repetida en Matc.", Row: p.Row, Column: p.Row})
						
							}
						}		
				}
			}
		}

		if !boolCond {
			
			for s := 0; s < p.InstrDefault.Len(); s++ {
				instr := p.InstrDefault.GetValue(s).(Default)
				instr.Interpretar(env, tree)
			}
		}

	}else if p.InstrCase != nil && p.InstrDefault == nil { // [<CASES_LIST>] 

		for _, s := range p.InstrCase.ToArray() {
			
			if s.(Case).Condicion != nil { // SOLO 1 CONDICION DE |
				newCond = s.(Case).Condicion.Interpretar(env, tree)

				if newCond.Valor == cond.Valor {
					if !boolCond {
						
						boolCond = true
						for _, j := range s.(Case).Instrucciones.ToArray() {
							j.(interfaces.Instruction).Interpretar(newTable, tree)
							
						} 
					} else{
						ast.NewException("Semantico","Expresion repetida en Match.", p.Row, p.Row)
						tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Expresion repetida en Matc.", Row: p.Row, Column: p.Row})
						
					}
				}

			}else{

				for _, i := range s.(Case).ListaExpresion.ToArray() {
						
						newCond = i.(Case).Condicion.Interpretar(env, tree)

						if newCond.Valor == cond.Valor {
							if !boolCond {

								boolCond = true
								for _, j := range s.(Case).Instrucciones.ToArray() {
									j.(interfaces.Instruction).Interpretar(newTable, tree)
									
								} 

							} else{
								ast.NewException("Semantico","Expresion repetida en Match.", p.Row, p.Row)
								tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Expresion repetida en Matc.", Row: p.Row, Column: p.Row})
						
							}			
						}
				}
			}
		}


	}else if p.InstrCase == nil && p.InstrDefault != nil { // [<DEFAULT>]
	
		for s := 0; s < p.InstrDefault.Len(); s++ {
			instr := p.InstrDefault.GetValue(s).(Default)
			instr.Interpretar(env, tree)

		}
	}

	return nil
}



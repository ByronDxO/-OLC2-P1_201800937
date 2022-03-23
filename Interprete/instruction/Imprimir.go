package instruction

import (
	"fmt"
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
)

type Println struct {
	Expresion 	*arrayList.List
	Formato 	string
	Condicion 	interfaces.Expresion
	Salto		string
}

func PRINTLN(val *arrayList.List, formato string, cond interfaces.Expresion, salto string) Println {
	exp := Println{val, formato, cond, salto}
	return exp
}

func (p Println) Interpretar(env interface{}, tree *ast.Arbol) interface{} {
	fmt.Println("llego")
	var value string
	var condBool bool = false
	var instr interfaces.Symbol
	value = ""
	

	if string(p.Formato) != "-1" && p.Condicion == nil {
		newPos := -1;
		contExpre := 0;
		
		for i := 0; i < len(p.Formato); i++ {

			if string(p.Formato[i]) == "{" {
				
				newPos = devolverPos(p.Formato, i)
				if newPos == -2 {
					// error 
					break
				} 

				condBool = true
			}
			
			if newPos == i && condBool{
				if contExpre < p.Expresion.Len() {
					auxCont := 0

					for _, s := range p.Expresion.ToArray() {
						
						if contExpre == auxCont {
							instr = s.(ListExpre).Interpretar(env, tree)
						}

						auxCont++;
					}
				
					value += fmt.Sprintf("%v", instr.Valor)
					contExpre++;
					condBool = false

				}else {
					break;
					// error en el formato y cantidad
				}
			}

			if !condBool && string(p.Formato[i]) != "}" {
				fmt.Println(string(p.Formato[i]))
				value += fmt.Sprintf("%v", string(p.Formato[i]))
			}
			
			
		}
		
	}else{
		instr = p.Condicion.Interpretar(env, tree)
		if instr.Tipo == interfaces.STRING {
			value += fmt.Sprintf("%v", instr.Valor)
		}else {
			
			// error
		}
		
	
	}
	if p.Salto == "println!"{
		value += "\n"
	}
	
	fmt.Println("---- RE ---- ")
	fmt.Println(value)
	fmt.Println("---- RE ---- ")
	tree.AddCode(value)
	return value
}


func devolverPos(text string, pos int) int {

	var newPos int = -2

	for i := pos; pos < len(text); i++{ 

		if string(text[pos]) == "}" {
			return pos
		}
		pos++;
	}
	
	return newPos
}

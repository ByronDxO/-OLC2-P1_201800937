package expresion

import (
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/ast"
	"fmt"
)

type Aritmetica struct {
	left      	interfaces.Expresion
	Operator 	string
	right      	interfaces.Expresion
	Unario   	bool
}

func NewOperacion(left interfaces.Expresion, Operator string, right interfaces.Expresion, unario bool) Aritmetica {

	exp := Aritmetica{left, Operator, right, unario}
	return exp
}

func (p Aritmetica) Interpretar(env interface{}, tree *ast.Arbol) interfaces.Symbol {
	
	SUMA_RESTA_DOMINANTE := [5][5]interfaces.TipoExpresion{
		//INTEGER			//FLOAT			   //STRING			  //BOOLEAN		   //NULL

		//INTEGER
		{interfaces.INTEGER, interfaces.FLOAT, interfaces.STRING, interfaces.NULL, interfaces.NULL},
		//FLOAT
		{interfaces.FLOAT, interfaces.FLOAT, interfaces.STRING, interfaces.NULL, interfaces.NULL},
		//STRING
		{interfaces.STRING, interfaces.STRING, interfaces.STRING, interfaces.STRING, interfaces.NULL},
		//BOOLEAN
		{interfaces.NULL, interfaces.NULL, interfaces.STRING, interfaces.NULL, interfaces.NULL},
		//NULL
		{interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL},
	}

	MULT_DIV_DOMINANTE := [5][5]interfaces.TipoExpresion{
		{interfaces.INTEGER, interfaces.FLOAT, interfaces.NULL, interfaces.NULL, interfaces.NULL},
		{interfaces.FLOAT, interfaces.FLOAT, interfaces.NULL, interfaces.NULL, interfaces.NULL},
		{interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL},
		{interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL},
		{interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL},
	}

	RELACIONAL_DOMINANTE := [5][5]interfaces.TipoExpresion{
		{interfaces.INTEGER, interfaces.FLOAT, interfaces.NULL, interfaces.NULL, interfaces.NULL},
		{interfaces.FLOAT, interfaces.FLOAT, interfaces.NULL, interfaces.NULL, interfaces.NULL},
		{interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL},
		{interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL},
		{interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL},
	}

	var exp_left interfaces.Symbol
	var exp_right interfaces.Symbol

	if p.Unario == true {
		exp_left = p.left.Interpretar(env, tree)
	} else {
		exp_left = p.left.Interpretar(env, tree)
		exp_right = p.right.Interpretar(env, tree)
	}

	var resultado interface{}

	var dominante interfaces.TipoExpresion

	switch p.Operator {
	case "+":
		{
			
			
			dominante = SUMA_RESTA_DOMINANTE[exp_left.Tipo][exp_right.Tipo]

			if dominante == interfaces.INTEGER {

				return interfaces.Symbol{Id: "", Tipo: dominante, Valor: exp_left.Valor.(int) + exp_right.Valor.(int)}

			} else if dominante == interfaces.FLOAT {
				if exp_left.Tipo == interfaces.INTEGER{
					return interfaces.Symbol{Id: "", Tipo: dominante, Valor: float64(exp_left.Valor.(int)) + exp_right.Valor.(float64)}

				}else if exp_right.Tipo == interfaces.INTEGER {
					return interfaces.Symbol{Id: "", Tipo: dominante, Valor: exp_left.Valor.(float64) + float64(exp_right.Valor.(int))}

				}

				return interfaces.Symbol{Id: "", Tipo: dominante, Valor: exp_left.Valor.(float64) + exp_right.Valor.(float64)}

			} else if dominante == interfaces.STRING {
				r1 := fmt.Sprintf("%v", exp_left.Valor)
				r2 := fmt.Sprintf("%v", exp_right.Valor)

				return interfaces.Symbol{Id: "", Tipo: dominante, Valor: r1 + r2}
			} else {
				fmt.Print("ERROR: No es posible sumar")
			}

		}

	case "-":
		{
			dominante = SUMA_RESTA_DOMINANTE[exp_left.Tipo][exp_right.Tipo]

			if dominante == interfaces.INTEGER {

				return interfaces.Symbol{Id: "", Tipo: dominante, Valor: exp_left.Valor.(int) - exp_right.Valor.(int)}

			} else if dominante == interfaces.FLOAT {
				return interfaces.Symbol{Id: "", Tipo: dominante, Valor: exp_left.Valor.(float64) - exp_right.Valor.(float64)}

			} else {
				fmt.Print("ERROR: No es posible restar")
			}
		}

	case "*":
		{
			dominante = MULT_DIV_DOMINANTE[exp_left.Tipo][exp_right.Tipo]

			if dominante == interfaces.INTEGER {
				return interfaces.Symbol{Id: "", Tipo: dominante, Valor: exp_left.Valor.(int) * exp_right.Valor.(int)}

			} else if dominante == interfaces.FLOAT {
				return interfaces.Symbol{Id: "", Tipo: dominante, Valor: exp_left.Valor.(float64) * exp_right.Valor.(float64)}

			} else {
				fmt.Print("ERROR: No es posible Multiplicar")
			}

		}

	case "/":
		{
			dominante = MULT_DIV_DOMINANTE[exp_left.Tipo][exp_right.Tipo]

			if dominante == interfaces.INTEGER {
				return interfaces.Symbol{Id: "", Tipo: dominante, Valor: exp_left.Valor.(int) / exp_right.Valor.(int)}

			} else if dominante == interfaces.FLOAT {
				return interfaces.Symbol{Id: "", Tipo: dominante, Valor: exp_left.Valor.(float64) / exp_right.Valor.(float64)}

			} else {
				fmt.Print("ERROR: No es posible Dividir")
			}

		}

	case "<":
		{
			dominante = RELACIONAL_DOMINANTE[exp_left.Tipo][exp_right.Tipo]

			if dominante == interfaces.INTEGER {

				return interfaces.Symbol{Id: "", Tipo: interfaces.BOOLEAN, Valor: exp_left.Valor.(int) < exp_right.Valor.(int)}

			} else if dominante == interfaces.FLOAT {
				return interfaces.Symbol{Id: "", Tipo: interfaces.BOOLEAN, Valor: exp_left.Valor.(float64) < exp_right.Valor.(float64)}

			} else {
				fmt.Print("ERROR: No es posible comparar <")
			}
		}

	case ">":
		{
			dominante = RELACIONAL_DOMINANTE[exp_left.Tipo][exp_right.Tipo]

			if dominante == interfaces.INTEGER {

				return interfaces.Symbol{Id: "", Tipo: interfaces.BOOLEAN, Valor: exp_left.Valor.(int) > exp_right.Valor.(int)}

			} else if dominante == interfaces.FLOAT {
				return interfaces.Symbol{Id: "", Tipo: interfaces.BOOLEAN, Valor: exp_left.Valor.(float64) > exp_right.Valor.(float64)}

			} else {
				fmt.Print("ERROR: No es posible comparar <")
			}
		}

	case "<=":
		{
			dominante = RELACIONAL_DOMINANTE[exp_left.Tipo][exp_right.Tipo]

			if dominante == interfaces.INTEGER {

				return interfaces.Symbol{Id: "", Tipo: interfaces.BOOLEAN, Valor: exp_left.Valor.(int) <= exp_right.Valor.(int)}

			} else if dominante == interfaces.FLOAT {
				return interfaces.Symbol{Id: "", Tipo: interfaces.BOOLEAN, Valor: exp_left.Valor.(float64) <= exp_right.Valor.(float64)}

			} else {
				fmt.Print("ERROR: No es posible comparar <")
			}
		}

	case ">=":
		{
			dominante = RELACIONAL_DOMINANTE[exp_left.Tipo][exp_right.Tipo]

			if dominante == interfaces.INTEGER {

				return interfaces.Symbol{Id: "", Tipo: interfaces.BOOLEAN, Valor: exp_left.Valor.(int) >= exp_right.Valor.(int)}

			} else if dominante == interfaces.FLOAT {
				return interfaces.Symbol{Id: "", Tipo: interfaces.BOOLEAN, Valor: exp_left.Valor.(float64) >= exp_right.Valor.(float64)}

			} else {
				fmt.Print("ERROR: No es posible comparar <")
			}
		}
	}

	return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: resultado}
}

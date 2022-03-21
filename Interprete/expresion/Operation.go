package expresion

import (
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/ast"
	"fmt"	
	"math"
	"strconv"
)

type Aritmetica struct {
	left      	interfaces.Expresion
	Operator 	string
	right      	interfaces.Expresion
	Unario   	bool
	type_left   string
	type_right  string
	Row 		int
	Column 		int
}

func NewOperacion(left interfaces.Expresion, Operator string, right interfaces.Expresion, unario bool, type_left string, type_right string, row int, column int) Aritmetica {

	exp := Aritmetica{left, Operator, right, unario, type_left, type_right, row, column}
	return exp
}

func (p Aritmetica) Interpretar(env interface{}, tree *ast.Arbol) interfaces.Symbol {
	

	var exp_left interfaces.Symbol
	var exp_right interfaces.Symbol

	if p.Unario == true {
		exp_left = p.left.Interpretar(env, tree)
	} else {
		exp_left = p.left.Interpretar(env, tree)
		exp_right = p.right.Interpretar(env, tree)
	}

	var resultado interface{}


	switch p.Operator {
	case "+":
		{
			

			/* ************************************************************** INTEGER ************************************************************** */
			if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as i64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(exp_left.Valor.(int) + int(exp_right.Valor.(float64)))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as i64"){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(int(exp_left.Valor.(float64)) + exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(exp_left.Valor.(int) + exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(int(exp_left.Valor.(float64)) + int(exp_right.Valor.(float64)))}
			
			/* ************************************************************** FLOAT ************************************************************** */
			}else if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as f64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(float64(exp_left.Valor.(int)) + exp_right.Valor.(float64))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as f64"){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(exp_left.Valor.(float64) + float64(exp_right.Valor.(int)))}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(float64(exp_left.Valor.(int)) + float64(exp_right.Valor.(int)))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(exp_left.Valor.(float64) + exp_right.Valor.(float64))}

			
			/* ************************************************************** STRING ************************************************************** */
			} else if (exp_left.Tipo == interfaces.STRING && exp_right.Tipo == interfaces.STRING) {
				r1 := fmt.Sprintf("%v", exp_left.Valor)
				r2 := fmt.Sprintf("%v", exp_right.Valor)

				return interfaces.Symbol{Id: "", Tipo: interfaces.STRING, Valor: r1 + r2}

			}else {
				excep := ast.NewException("Semantico","No es posible Sumar.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible Sumar.", Row: p.Row, Column: p.Column})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}

		}

	case "-":
		{

			
			if p.Unario {
				
				if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as i64" || p.type_left == "-1")) { 
					return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int((-1)*(exp_left.Valor.(int)))}

				}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64")) {
					return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int((-1)*(exp_left.Valor.(int)))}

				}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) {
					return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64((-1.0)*(exp_left.Valor.(float64)))}
					
				}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as f64")) {
					return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64((-1.0)*(exp_left.Valor.(float64)))}
				
				}else {
					
					excep := ast.NewException("Semantico","No es posible Unario.", p.Row, p.Column)
					tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible Unario.", Row: p.Row, Column: p.Column})
					return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

				}

				

			}else{
						/* ************************************************************** INTEGER ************************************************************** */
				if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as i64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
					return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(exp_left.Valor.(int) - int(exp_right.Valor.(float64)))}

				}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as i64"){
					return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(int(exp_left.Valor.(float64)) - exp_right.Valor.(int))}

				}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as i64" || p.type_right == "-1")){
					return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(exp_left.Valor.(int) - exp_right.Valor.(int))}

				}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
					return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(int(exp_left.Valor.(float64)) - int(exp_right.Valor.(float64)))}
				
				/* ************************************************************** FLOAT ************************************************************** */
				}else if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as f64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
					return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(float64(exp_left.Valor.(int)) - exp_right.Valor.(float64))}

				}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as f64"){
					return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(exp_left.Valor.(float64) - float64(exp_right.Valor.(int)))}

				}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as f64" || p.type_right == "-1")){
					return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(float64(exp_left.Valor.(int)) - float64(exp_right.Valor.(int)))}

				}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
					return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(exp_left.Valor.(float64) - exp_right.Valor.(float64))}

				}else {
					
					excep := ast.NewException("Semantico","No es posible Restar.", p.Row, p.Column)
					tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible Restar.", Row: p.Row, Column: p.Column})
					return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

				}

			}
					
		}
	case "*":
		{
			/* ************************************************************** INTEGER ************************************************************** */
			if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as i64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(exp_left.Valor.(int) * int(exp_right.Valor.(float64)))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as i64"){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(int(exp_left.Valor.(float64)) * exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(exp_left.Valor.(int) * exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(int(exp_left.Valor.(float64)) * int(exp_right.Valor.(float64)))}
			
			/* ************************************************************** FLOAT ************************************************************** */
			}else if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as f64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(float64(exp_left.Valor.(int)) * exp_right.Valor.(float64))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as f64"){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(exp_left.Valor.(float64) * float64(exp_right.Valor.(int)))}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(float64(exp_left.Valor.(int)) * float64(exp_right.Valor.(int)))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(exp_left.Valor.(float64) * exp_right.Valor.(float64))}

			}else {
				
				excep := ast.NewException("Semantico","No es posible Multiplicar.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible Multiplicar.", Row: p.Row, Column: p.Column})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}

		}

	case "/":
		{
			/* ************************************************************** INTEGER ************************************************************** */
			if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as i64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(exp_left.Valor.(int) / int(exp_right.Valor.(float64)))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as i64"){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(int(exp_left.Valor.(float64)) / exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(exp_left.Valor.(int) / exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(int(exp_left.Valor.(float64)) / int(exp_right.Valor.(float64)))}
			
			/* ************************************************************** FLOAT ************************************************************** */
			}else if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as f64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(float64(exp_left.Valor.(int)) / exp_right.Valor.(float64))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as f64"){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(exp_left.Valor.(float64) / float64(exp_right.Valor.(int)))}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(float64(exp_left.Valor.(int)) / float64(exp_right.Valor.(int)))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(exp_left.Valor.(float64) / exp_right.Valor.(float64))}

			}else {
				
				excep := ast.NewException("Semantico","No es posible Dividir.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible Dividir.", Row: p.Row, Column: p.Column})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}

		}

	case "%":
		{
		/* ************************************************************** INTEGER ************************************************************** */
			if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as i64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(exp_left.Valor.(int) % int(exp_right.Valor.(float64)))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as i64"){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(int(exp_left.Valor.(float64)) % exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(exp_left.Valor.(int) % exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: int(int(exp_left.Valor.(float64)) % int(exp_right.Valor.(float64)))}
			
			/* ************************************************************** FLOAT ************************************************************** */
			}else if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as f64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(math.Mod(float64(exp_left.Valor.(int)) , exp_right.Valor.(float64)))  }

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as f64"){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(math.Mod(exp_left.Valor.(float64), float64(exp_right.Valor.(int)))) }

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(math.Mod(float64(exp_left.Valor.(int)), float64(exp_right.Valor.(int))))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: interfaces.FLOAT, Valor: float64(math.Mod(exp_left.Valor.(float64), exp_right.Valor.(float64)))}

			}else {
				
				excep := ast.NewException("Semantico","No es posible el Modulo %.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible el Modulo %.", Row: p.Row, Column: p.Column})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}
		}

	case "<":
		{
			auxType := interfaces.BOOLEAN
			/* ************************************************************** INTEGER ************************************************************** */
			if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as i64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: bool(exp_left.Valor.(int) < int(exp_right.Valor.(float64)))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as i64"){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: int(exp_left.Valor.(float64)) < exp_right.Valor.(int)}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(int) < exp_right.Valor.(int)}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: int(exp_left.Valor.(float64)) < int(exp_right.Valor.(float64))}
			
			/* ************************************************************** FLOAT ************************************************************** */
			}else if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as f64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: float64(exp_left.Valor.(int)) < exp_right.Valor.(float64)}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as f64"){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(float64) < float64(exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: float64(exp_left.Valor.(int)) < float64(exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(float64) < exp_right.Valor.(float64)}

			/* ************************************************************** BOOLEAN ************************************************************** */
			}else if (exp_left.Tipo == interfaces.BOOLEAN) && (exp_right.Tipo == interfaces.BOOLEAN) {
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: casteoBool(exp_left.Valor.(string)) < casteoBool(exp_right.Valor.(string))}

			}else {
				
				excep := ast.NewException("Semantico","No es posible comparar <.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar <.", Row: p.Row, Column: p.Column})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}
		}

	case ">":
		{
			auxType := interfaces.BOOLEAN
			/* ************************************************************** INTEGER ************************************************************** */
			if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as i64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(int) > int(exp_right.Valor.(float64))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as i64"){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: int(exp_left.Valor.(float64)) > exp_right.Valor.(int)}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(int) > exp_right.Valor.(int)}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: int(exp_left.Valor.(float64)) > int(exp_right.Valor.(float64))}
			
			/* ************************************************************** FLOAT ************************************************************** */
			}else if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as f64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: float64(exp_left.Valor.(int)) > exp_right.Valor.(float64)}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as f64"){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(float64) > float64(exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: float64(exp_left.Valor.(int)) > float64(exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(float64) > exp_right.Valor.(float64)}

			/* ************************************************************** BOOLEAN ************************************************************** */
			}else if (exp_left.Tipo == interfaces.BOOLEAN) && (exp_right.Tipo == interfaces.BOOLEAN) {
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: casteoBool(exp_left.Valor.(string)) > casteoBool(exp_right.Valor.(string))}

			}else {
				
				excep := ast.NewException("Semantico","No es posible comparar >.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar >.", Row: p.Row, Column: p.Column})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}
		}

	case "<=":
		{
			auxType := interfaces.BOOLEAN
			/* ************************************************************** INTEGER ************************************************************** */
			if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as i64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(int) <= int(exp_right.Valor.(float64))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as i64"){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: int(exp_left.Valor.(float64)) <= exp_right.Valor.(int)}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(int) <= exp_right.Valor.(int)}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: int(exp_left.Valor.(float64)) <= int(exp_right.Valor.(float64))}
			
			/* ************************************************************** FLOAT ************************************************************** */
			}else if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as f64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: float64(exp_left.Valor.(int)) <= exp_right.Valor.(float64)}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as f64"){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(float64) <= float64(exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: float64(exp_left.Valor.(int)) <= float64(exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(float64) <= exp_right.Valor.(float64)}

			/* ************************************************************** BOOLEAN ************************************************************** */
			}else if (exp_left.Tipo == interfaces.BOOLEAN) && (exp_right.Tipo == interfaces.BOOLEAN) {
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: casteoBool(exp_left.Valor.(string)) <= casteoBool(exp_right.Valor.(string))}

			}else {
				
				excep := ast.NewException("Semantico","No es posible comparar <=.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar <=.", Row: p.Row, Column: p.Column})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}
		}

	case ">=":
		{
			auxType := interfaces.BOOLEAN
			/* ************************************************************** INTEGER ************************************************************** */
			if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as i64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(int) >= int(exp_right.Valor.(float64))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as i64"){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: int(exp_left.Valor.(float64)) >= exp_right.Valor.(int)}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(int) >= exp_right.Valor.(int)}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: int(exp_left.Valor.(float64)) >= int(exp_right.Valor.(float64))}
			
			/* ************************************************************** FLOAT ************************************************************** */
			}else if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as f64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: float64(exp_left.Valor.(int)) >= exp_right.Valor.(float64)}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as f64"){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(float64) >= float64(exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: float64(exp_left.Valor.(int)) >= float64(exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(float64) >= exp_right.Valor.(float64)}

			/* ************************************************************** BOOLEAN ************************************************************** */
			}else if (exp_left.Tipo == interfaces.BOOLEAN) && (exp_right.Tipo == interfaces.BOOLEAN) {
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: casteoBool(exp_left.Valor.(string)) >= casteoBool(exp_right.Valor.(string))}

			}else {
				
				excep := ast.NewException("Semantico","No es posible comparar >=.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar >=.", Row: p.Row, Column: p.Column})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}
		}
	
	case "!=":
		{
			auxType := interfaces.BOOLEAN
			/* ************************************************************** INTEGER ************************************************************** */
			if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as i64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(int) != int(exp_right.Valor.(float64))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as i64"){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: int(exp_left.Valor.(float64)) != exp_right.Valor.(int)}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(int) != exp_right.Valor.(int)}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: int(exp_left.Valor.(float64)) != int(exp_right.Valor.(float64))}
			
			/* ************************************************************** FLOAT ************************************************************** */
			}else if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as f64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: float64(exp_left.Valor.(int)) != exp_right.Valor.(float64)}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as f64"){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(float64) != float64(exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: float64(exp_left.Valor.(int)) != float64(exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(float64) != exp_right.Valor.(float64)}

			/* ************************************************************** STRING ************************************************************** */
			}else if (exp_left.Tipo == interfaces.STRING) && (exp_right.Tipo == interfaces.STRING) {
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(string) != exp_right.Valor.(string)}

			/* ************************************************************** BOOLEAN ************************************************************** */
			}else if (exp_left.Tipo == interfaces.BOOLEAN) && (exp_right.Tipo == interfaces.BOOLEAN) {
				exp1,_ := strconv.ParseBool(fmt.Sprintf("%v", exp_left.Valor))
				exp2,_ := strconv.ParseBool(fmt.Sprintf("%v", exp_right.Valor))
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp1 != exp2}

			}else {
				
				excep := ast.NewException("Semantico","No es posible comparar !=.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar !=.", Row: p.Row, Column: p.Column})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}
		}
	
	case "==":
		{
			auxType := interfaces.BOOLEAN
			/* ************************************************************** INTEGER ************************************************************** */
			if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as i64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(int) == int(exp_right.Valor.(float64))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as i64"){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: int(exp_left.Valor.(float64)) == exp_right.Valor.(int)}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(int) == exp_right.Valor.(int)}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as i64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: int(exp_left.Valor.(float64)) == int(exp_right.Valor.(float64))}
			
			/* ************************************************************** FLOAT ************************************************************** */
			}else if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as f64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: float64(exp_left.Valor.(int)) == exp_right.Valor.(float64)}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && p.type_right == "as f64"){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(float64) == float64(exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.INTEGER && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.INTEGER && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: float64(exp_left.Valor.(int)) == float64(exp_right.Valor.(int))}

			}else if (exp_left.Tipo == interfaces.FLOAT && (p.type_left == "as f64" || p.type_left == "-1")) && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as f64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(float64) == exp_right.Valor.(float64)}

			/* ************************************************************** STRING ************************************************************** */
			}else if (exp_left.Tipo == interfaces.STRING) && (exp_right.Tipo == interfaces.STRING) {
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(string) == exp_right.Valor.(string)}

			/* ************************************************************** BOOLEAN ************************************************************** */
			}else if (exp_left.Tipo == interfaces.BOOLEAN) && (exp_right.Tipo == interfaces.BOOLEAN) {
				exp1,_ := strconv.ParseBool(fmt.Sprintf("%v", exp_left.Valor))
				exp2,_ := strconv.ParseBool(fmt.Sprintf("%v", exp_right.Valor))
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp1 == exp2}

			}else {
				
				excep := ast.NewException("Semantico","No es posible comparar ==.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar ==.", Row: p.Row, Column: p.Column})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}
		}


	case "&&":
		{
			// auxType := interfaces.BOOLEAN
			/* ************************************************************** INTEGER ************************************************************** */
			if exp_left.Tipo == interfaces.BOOLEAN && exp_right.Tipo == interfaces.BOOLEAN {
				exp1,_ := strconv.ParseBool(fmt.Sprintf("%v", exp_left.Valor))
				exp2,_ := strconv.ParseBool(fmt.Sprintf("%v", exp_right.Valor))
				
				return interfaces.Symbol{Id: "", Tipo: interfaces.BOOLEAN, Valor: exp1 && exp2 }

		
			}else {
				
				excep := ast.NewException("Semantico","No es posible comparar &&.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar &&.", Row: p.Row, Column: p.Column})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}
		}

	case "||":
		{
			// auxType := interfaces.BOOLEAN
			/* ************************************************************** INTEGER ************************************************************** */
			if exp_left.Tipo == interfaces.BOOLEAN && exp_right.Tipo == interfaces.BOOLEAN {
				exp1,_ := strconv.ParseBool(fmt.Sprintf("%v", exp_left.Valor))
				exp2,_ := strconv.ParseBool(fmt.Sprintf("%v", exp_right.Valor))
				
				return interfaces.Symbol{Id: "", Tipo: interfaces.BOOLEAN, Valor: exp1 || exp2 }

		
			}else {
				
				excep := ast.NewException("Semantico","No es posible comparar ||.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar ||.", Row: p.Row, Column: p.Column})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}
		}

	case "!":
		{
			if p.Unario {
				if exp_left.Tipo == interfaces.BOOLEAN {
					exp1,_ := strconv.ParseBool(fmt.Sprintf("%v", exp_left.Valor))
					return interfaces.Symbol{Id: "", Tipo: interfaces.BOOLEAN, Valor: !exp1 }
	
			
				}else {
					
					excep := ast.NewException("Semantico","No es posible comparar !.", p.Row, p.Column)
					tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar !.", Row: p.Row, Column: p.Column})
					return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}
	
				}
			}

			excep := ast.NewException("Semantico","No es posible, unario incorrecto en !.", p.Row, p.Column)
			tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible, unario incorrecto en !", Row: p.Row, Column: p.Column})
			return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			
		}

	}

	

	return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: resultado}
}



func casteoBool(exp string) int {

	exp1,_ := strconv.ParseBool(exp)
	bitSetVar1 := int(0)

	if exp1 {
		bitSetVar1 = 1
	}
	
	return bitSetVar1
}
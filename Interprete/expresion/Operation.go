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
	type_left   string
	type_right  string
}

func NewOperacion(left interfaces.Expresion, Operator string, right interfaces.Expresion, unario bool, type_left string, type_right string) Aritmetica {

	exp := Aritmetica{left, Operator, right, unario, type_left, type_right}
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
				excep := ast.NewException("Semantico","No es posible Sumar.")
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible Sumar."})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}

		}

	case "-":
		{
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
				
				excep := ast.NewException("Semantico","No es posible Restar.")
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible Restar."})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

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
				
				excep := ast.NewException("Semantico","No es posible Multiplicar.")
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible Multiplicar."})
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
				
				excep := ast.NewException("Semantico","No es posible Dividir.")
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible Dividir."})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}

		}

	case "<":
		{
			auxType := interfaces.BOOLEAN
			/* ************************************************************** INTEGER ************************************************************** */
			if (exp_left.Tipo == interfaces.INTEGER && p.type_left == "as i64") && (exp_right.Tipo == interfaces.FLOAT && (p.type_right == "as i64" || p.type_right == "-1")){
				return interfaces.Symbol{Id: "", Tipo: auxType, Valor: exp_left.Valor.(int) < int(exp_right.Valor.(float64))}

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

			}else {
				
				excep := ast.NewException("Semantico","No es posible comparar <.")
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar <."})
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

			}else {
				
				excep := ast.NewException("Semantico","No es posible comparar >.")
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar >."})
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

			}else {
				
				excep := ast.NewException("Semantico","No es posible comparar <=.")
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar <=."})
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

			}else {
				
				excep := ast.NewException("Semantico","No es posible comparar >=.")
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar >=."})
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

			}else {
				
				excep := ast.NewException("Semantico","No es posible comparar >=.")
				tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "No es posible comparar >=."})
				return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}

			}
		}
	}

	return interfaces.Symbol{Id: "", Tipo: interfaces.INTEGER, Valor: resultado}
}

package expresion

import (
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/ast"
)


type Primitivo struct {
	Valor interface{}
	Tipo interfaces.TipoExpresion
	Row 	int
	Column 	int
}

func (p Primitivo) Interpretar(env interface{}, tree *ast.Arbol) interfaces.Symbol {
	
	return interfaces.Symbol{
		Id: "",
		Tipo: p.Tipo,
		Valor: p.Valor,
	}

}


func PRIMITIVO(val interface{}, tipo interfaces.TipoExpresion, row int, column int) Primitivo {
	expression := Primitivo{val, tipo, row, column}
	return expression

}
package expresion

import (

	"OLC2/Interprete/interfaces"
)


type Primitivo struct {
	Valor interface{}
	Tipo interfaces.TipoExpresion
}

func (p Primitivo) Interpretar(env interface{}) interfaces.Symbol {
	
	return interfaces.Symbol{
		Id: "",
		Tipo: p.Tipo,
		Valor: p.Valor,
	}

}


func PRIMITIVO(val interface{}, tipo interfaces.TipoExpresion) Primitivo {

	expression := Primitivo{val, tipo}
	return expression

}
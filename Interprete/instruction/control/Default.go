package control

import (

	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
	arrayList "github.com/colegno/arraylist"
	
)

type Default struct {

	Instrucciones	*arrayList.List
}


func NewDefault(instrucciones *arrayList.List) Default {
	instr := Default{instrucciones}
	return instr
}	



func (p Default) Interpretar(env interface{}, tree *ast.Arbol) interface{} {
	
	var newTable environment.Environment
	newTable = environment.NewEnvironment(env.(environment.Environment))
	for _, i := range p.Instrucciones.ToArray() {
		i.(interfaces.Instruction).Interpretar(newTable, tree)
	}


	return nil
}
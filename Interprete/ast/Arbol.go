package ast

import (
	// "OLC2/Interprete/environment"
	// "fmt"
	arrayList "github.com/colegno/arraylist"
)

type Arbol struct {
	Code     	*arrayList.List
	_Exception  *arrayList.List
	// Tabla_Global environment.Environment
}

type Exception struct {
	Tipo 		string
	Descripcion string
}


func NewArbol() *Arbol {
	tree := Arbol{Code: arrayList.New(), _Exception: arrayList.New()} 	
	return &tree
}

func NewException(tipo string, descripcion string) *Exception {
	e := Exception{Tipo: tipo, Descripcion: descripcion}
	return &e
}

/* Add text to Code */
func (a *Arbol) AddCode(input string) {
	a.Code.Add(input) 
}

/* Get text to Code */
func (a Arbol) GetCode() *arrayList.List {
	return a.Code
}

/* Add Exception */
func (a *Arbol) AddException(e Exception) {
	a._Exception.Add(e) 
}


/* Get Exception */
func (a Arbol) GetException() *arrayList.List {
	return a._Exception
}

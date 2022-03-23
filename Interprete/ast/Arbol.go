package ast

import (
	// "OLC2/Interprete/environment"
	// "fmt"
	arrayList "github.com/colegno/arraylist"
	"time"
	"fmt"
)

type Arbol struct {
	Code     	*arrayList.List
	_Exception  *arrayList.List
	// Tabla_Global environment.Environment
}

type Exception struct {
	Tipo 		string
	Descripcion string
	Row			int
	Column 		int
	Time 		string
}


func NewArbol() *Arbol {
	tree := Arbol{Code: arrayList.New(), _Exception: arrayList.New()} 	
	return &tree
}

func NewException(tipo string, descripcion string, row int, column int) *Exception {
	t := time.Now()
	hora := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	
	e := Exception{Tipo: tipo, Descripcion: descripcion, Row: row, Column: column, Time: hora}
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
	t := time.Now()
	hora := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d",
		t.Day(), t.Month(), t.Year(), 
		t.Hour(), t.Minute(), t.Second())
	e.Time = hora
	a._Exception.Add(e) 
}


/* Get Exception */
func (a Arbol) GetException() *arrayList.List {
	return a._Exception
}

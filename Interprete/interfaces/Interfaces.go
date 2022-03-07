package interfaces

type Symbol struct {
	Id    string
	Tipo  TipoExpresion
	Valor interface{}
}

type Expresion interface {
	Interpretar(env interface{}) Symbol
}

type Instruction interface {
	Interpretar(env interface{}) interface{}
}

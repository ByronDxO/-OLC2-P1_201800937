package environment

import (
	"OLC2/Interprete/interfaces"
	"fmt"
)

type Environment struct {
	father   interface{}
	variable map[string]interfaces.Symbol
}

func NewEnvironment(father interface{}) Environment {
	env := Environment{father, make(map[string]interfaces.Symbol)}
	return env
}

func (env Environment) SaveVariable(id string, value interfaces.Symbol, tipo interfaces.TipoExpresion) {
	if variable, ok := env.variable[id]; ok {
		fmt.Println("La variable " + variable.Id + " ya existe")
		return
	}
	env.variable[id] = interfaces.Symbol{Id: id, Tipo: tipo, Valor: value}
}

func (env Environment) GetVariable(id string) interfaces.Symbol {

	var tmpEnv Environment
	tmpEnv = env

	for {
		if variable, ok := tmpEnv.variable[id]; ok {
			return variable
		}

		if tmpEnv.father == nil {
			break
		} else {
			tmpEnv = tmpEnv.father.(Environment)
		}
	}

	fmt.Println("La variable no existe")
	return interfaces.Symbol{Id: "", Tipo: interfaces.NULL, Valor: interfaces.Symbol{Id: "", Tipo: interfaces.NULL, Valor: 0}}
}

func (env Environment) AlterVariable(id string, value interfaces.Symbol) interfaces.Symbol {

	var tmpEnv Environment
	tmpEnv = env

	for {
		if variable, ok := tmpEnv.variable[id]; ok {
			tmpEnv.variable[id] = interfaces.Symbol{Id: id, Tipo: variable.Tipo, Valor: value}
			return variable
		}

		if tmpEnv.father == nil {
			break
		} else {
			tmpEnv = tmpEnv.father.(Environment)
		}
	}

	fmt.Println("La variable no existe")
	return interfaces.Symbol{Id: "", Tipo: interfaces.NULL, Valor: interfaces.Symbol{Id: "", Tipo: interfaces.NULL, Valor: 0}}
}

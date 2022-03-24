package environment

import (
	"OLC2/Interprete/interfaces"
	"fmt"
)

type Environment struct {
	anterior   interface{}
	variable map[string]interfaces.Symbol
	Function map[string]interfaces.Symbol
	// structs  map[string]interfaces.Symbol
}

func NewEnvironment(anterior interface{}) Environment {
	
	env := Environment{anterior, make(map[string]interfaces.Symbol), make(map[string]interfaces.Symbol)}
	return env
}

func (env Environment) AddSymbol(id string, value interfaces.Symbol, tipo interfaces.TipoExpresion, isMut bool) {
	if variable, ok := env.variable[id]; ok {
		fmt.Println("La variable " + variable.Id + " ya existe")
		return
	}
	env.variable[id] = interfaces.Symbol{Id: id, Tipo: tipo, Valor: value, IsMut: isMut}
}

func (env Environment) GetSymbol(id string) interfaces.Symbol {

	var tmpEnv Environment
	tmpEnv = env

	for {
		if variable, ok := tmpEnv.variable[id]; ok {
			return variable
		}

		if tmpEnv.anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.anterior.(Environment)
		}
	}

	fmt.Println("La variable no existe")
	return interfaces.Symbol{Id: "", Tipo: interfaces.NULL, Valor: interfaces.Symbol{Id: "", Tipo: interfaces.NULL, Valor: 0}}
}

func (env Environment) SetSymbol(id string, value interfaces.Symbol, mut bool) interfaces.Symbol {

	var tmpEnv Environment
	tmpEnv = env

	for {
		if variable, ok := tmpEnv.variable[id]; ok {
			tmpEnv.variable[id] = interfaces.Symbol{Id: id, Tipo: variable.Tipo, Valor: value, IsMut:mut}
			return variable
		}

		if tmpEnv.anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.anterior.(Environment)
		}
	}

	fmt.Println("La variable no existe")
	return interfaces.Symbol{Id: "", Tipo: interfaces.NULL, Valor: interfaces.Symbol{Id: "", Tipo: interfaces.NULL, Valor: 0}}
}





func (env Environment) AddFunction(id string, value interfaces.Symbol, tipo interfaces.TipoExpresion) {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if tmpEnv.anterior == nil {
			tmpEnv.Function[id] = interfaces.Symbol{Id: id, Tipo: tipo, Valor: value, IsMut: true}
			fmt.Println("--- func ")
			fmt.Println(env.Function[id])
			fmt.Println("--- *** ---- ")
			break;
		} else {
			tmpEnv = tmpEnv.anterior.(Environment)
		}
	}
	

	
}


func (env Environment) GetFunction(id string) interfaces.Symbol {

	var tmpEnv Environment
	tmpEnv = env
	fmt.Println(env.Function)

	for {
		if tmpEnv.anterior == nil {
			break
			
		} else {
			tmpEnv = tmpEnv.anterior.(Environment)
		}
	}

	
	
	if function, ok := tmpEnv.Function[id]; ok {
		return function
	}else {
		tmpEnv = tmpEnv.anterior.(Environment)
	}


	fmt.Println("La function no existe")
	fmt.Println("---- fin func --")
	return interfaces.Symbol{Id: "", Tipo: interfaces.NULL, Valor: interfaces.Symbol{Id: "", Tipo: interfaces.NULL, Valor: 0}}


	
}




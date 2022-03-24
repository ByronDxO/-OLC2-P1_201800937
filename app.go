package main

import (
	"fmt"
	// "log"
	"reflect"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"

	/*INTEPRETE*/
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
	"OLC2/Interprete/instruction"
	"OLC2/Interprete/instruction/function"
	"OLC2/Interprete/ANTLR/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)
var CODE_OUT_ string = ""
var tablaSimboloP []interface{}

func main() {
	// Declarate
	

	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"CODE_INPUT": "",
			"CODE_OUT": "",
			"Tabla_Error" : nil,
		})
	})

	app.Post("/interpretar", Execute)

	/*
		Se ejecuta el servidor y en caso de fallar, muetra log.Fatal con el error.
	*/
	_ = app.Listen(":3000")
}

type getInput struct {
	Input string `form:"editor"`
}

func Execute(c *fiber.Ctx) error {
	data := new(getInput)
	// fmt.Pritln(data)
	if err := c.BodyParser(data); err != nil {
		return err
	}

	is := antlr.NewInputStream(data.Input)

	lexer := parser.NewChemsLexer(is) // aqui van tokens
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewChems(stream) // Aqui va esquema

	p.BuildParseTrees = true
	tree := p.Start()
	
	result := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(result, tree)

	fmt.Println("--")
	// fmt.Println(tablaSimbolo)

	return c.Render("index", fiber.Map{
		"CODE_INPUT"  : data.Input,
		"CODE_OUT"	  : CODE_OUT_,
		"Tabla_Error" : tablaSimboloP,
	})
}



/* ANTLR*/

type TreeShapeListener struct {
	*parser.BaseChemsListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	result := ctx.GetLista()

	var tree *ast.Arbol
	tree = ast.NewArbol()

	var _salida string
	_salida = ""


	var globalEnv environment.Environment
	globalEnv = environment.NewEnvironment(nil)
	CODE_OUT_ ="";
	var contMain int = 0

	for _ , s := range result.ToArray() {
		newInstr := s.(interfaces.Instruction)
		if reflect.TypeOf(newInstr).String() != "instruction.Main" && reflect.TypeOf(newInstr).String() != "function.Function" { 
			excep := ast.NewException("Semantico","Solo puede ir Main, Func, Array y Mod.", -1, -1)
			tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
			break
		}

		if reflect.TypeOf(newInstr).String() == "function.Function" {
			// fmt.Println("--- func ---")
			// fmt.Println()
			// fmt.Println(newInstr.(function.Function).Parametro)
			// fmt.Println()
			// fmt.Println()
			// fmt.Println(newInstr.(function.Function).Row)
			// fmt.Println(newInstr.(function.Function).Column)
			// fmt.Println("--- *** ---")

			value := interfaces.Symbol {
				Id     : newInstr.(function.Function).Id,
				Tipo   : newInstr.(function.Function).Tipo,
				Valor  : interfaces.SymbolFunction {
					
					Id			    : newInstr.(function.Function).Id,
					Tipo			: newInstr.(function.Function).Tipo,
					Instrucciones	: newInstr.(function.Function).Instrucciones,
					Parametro		: newInstr.(function.Function).Parametro,
					IsMut			: true,
						
				},
				IsMut  : true,
			}



			globalEnv.AddFunction(newInstr.(function.Function).Id, value, newInstr.(function.Function).Tipo)
			// AddSymbol(id, value interfaces.Symbol, tipo interfaces.TipoExpresion) {
		}
	}
	for _ , s := range result.ToArray() {

		newInstr := s.(interfaces.Instruction)
		
		if reflect.TypeOf(newInstr).String() == "instruction.Main" {
			if contMain > 0 {
				excep := ast.NewException("Semantico","Existen dos funciones Main.", newInstr.(instruction.Main).Row, newInstr.(instruction.Main).Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				break
			}
			newInstr.Interpretar(globalEnv, tree)
			
			if reflect.TypeOf(newInstr).String() == "transferencia.Break" 	 { 
				excep := ast.NewException("Semantico","Sentencia Break fuera de Ciclo.", newInstr.(instruction.Main).Row, newInstr.(instruction.Main).Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				break
			}
			if reflect.TypeOf(newInstr).String() == "transferencia.Continue" { 
				excep := ast.NewException("Semantico","Sentencia Continue fuera de Ciclo.", newInstr.(instruction.Main).Row, newInstr.(instruction.Main).Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				break
			}
			if reflect.TypeOf(newInstr).String() == "transferencia.Return"   { 
				excep := ast.NewException("Semantico","Sentencia Return fuera de Ciclo.", newInstr.(instruction.Main).Row, newInstr.(instruction.Main).Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				break
			}
			contMain++;
		}
		
	}


	fmt.Println(" ************* RESULTADO ************* ")
	for _, s := range tree.GetCode().ToArray() {
		_salida += fmt.Sprintf("%v", s)
	}
	fmt.Println(_salida)
	fmt.Println("----------")

	
	var OutException string
	OutException = ""

	
	fmt.Println(tablaSimboloP)
	for _, s := range tree.GetException().ToArray() {
		OutException += fmt.Sprintf("%v", s)
		m := make(map[string]string)
		m["Id"] 		 = fmt.Sprintf("%v", s.(ast.Exception).Tipo)
		m["Descripcion"] = fmt.Sprintf("%v", s.(ast.Exception).Descripcion)
		m["Row"] 		 = fmt.Sprintf("%v", s.(ast.Exception).Row)
		m["Column"]      = fmt.Sprintf("%v", s.(ast.Exception).Column)
		m["Time"] 		 = fmt.Sprintf("%v", s.(ast.Exception).Time)
		
		tablaSimboloP = append(tablaSimboloP, m) 
	}

	fmt.Println(OutException)
	fmt.Println("----------")


	CODE_OUT_ = _salida
}


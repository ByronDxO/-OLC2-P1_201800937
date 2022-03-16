package main

import (
	"fmt"
	// "log"
	// "reflect"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"

	/*INTEPRETE*/
	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
	// "OLC2/Interprete/instruction"
	"OLC2/Interprete/ANTLR/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)
var CODE_OUT_ string = ""
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

	// fmt.Println(CODE_OUT_)

	return c.Render("index", fiber.Map{
		"CODE_INPUT": data.Input,
		"CODE_OUT": CODE_OUT_,
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
	for _ , s := range result.ToArray() {
		s.(interfaces.Instruction).Interpretar(globalEnv, tree)
		
	}


	fmt.Println("SALIDA AS")
	for _, s := range tree.GetCode().ToArray() {
		_salida += fmt.Sprintf("%v", s)
	}
	fmt.Println(_salida)
	fmt.Println("----------")
	CODE_OUT_ = _salida
}




// Code generated from Interprete/ANTLR/Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChemsListener is a complete listener for a parse tree produced by Chems.
type ChemsListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExp_arit is called when entering the exp_arit production.
	EnterExp_arit(c *Exp_aritContext)

	// EnterPrimitivo is called when entering the primitivo production.
	EnterPrimitivo(c *PrimitivoContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExp_arit is called when exiting the exp_arit production.
	ExitExp_arit(c *Exp_aritContext)

	// ExitPrimitivo is called when exiting the primitivo production.
	ExitPrimitivo(c *PrimitivoContext)
}

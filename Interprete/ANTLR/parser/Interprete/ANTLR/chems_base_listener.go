// Code generated from Interprete/ANTLR/Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseChemsListener is a complete listener for a parse tree produced by Chems.
type BaseChemsListener struct{}

var _ ChemsListener = &BaseChemsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChemsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChemsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChemsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChemsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseChemsListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseChemsListener) ExitStart(ctx *StartContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseChemsListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseChemsListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseChemsListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseChemsListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseChemsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseChemsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExp_arit is called when production exp_arit is entered.
func (s *BaseChemsListener) EnterExp_arit(ctx *Exp_aritContext) {}

// ExitExp_arit is called when production exp_arit is exited.
func (s *BaseChemsListener) ExitExp_arit(ctx *Exp_aritContext) {}

// EnterPrimitivo is called when production primitivo is entered.
func (s *BaseChemsListener) EnterPrimitivo(ctx *PrimitivoContext) {}

// ExitPrimitivo is called when production primitivo is exited.
func (s *BaseChemsListener) ExitPrimitivo(ctx *PrimitivoContext) {}

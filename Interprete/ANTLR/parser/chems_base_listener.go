// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

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

// EnterEnd_instr is called when production end_instr is entered.
func (s *BaseChemsListener) EnterEnd_instr(ctx *End_instrContext) {}

// ExitEnd_instr is called when production end_instr is exited.
func (s *BaseChemsListener) ExitEnd_instr(ctx *End_instrContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseChemsListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseChemsListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterInstr_println is called when production instr_println is entered.
func (s *BaseChemsListener) EnterInstr_println(ctx *Instr_printlnContext) {}

// ExitInstr_println is called when production instr_println is exited.
func (s *BaseChemsListener) ExitInstr_println(ctx *Instr_printlnContext) {}

// EnterInstr_declaracion is called when production instr_declaracion is entered.
func (s *BaseChemsListener) EnterInstr_declaracion(ctx *Instr_declaracionContext) {}

// ExitInstr_declaracion is called when production instr_declaracion is exited.
func (s *BaseChemsListener) ExitInstr_declaracion(ctx *Instr_declaracionContext) {}

// EnterInstr_asignacion is called when production instr_asignacion is entered.
func (s *BaseChemsListener) EnterInstr_asignacion(ctx *Instr_asignacionContext) {}

// ExitInstr_asignacion is called when production instr_asignacion is exited.
func (s *BaseChemsListener) ExitInstr_asignacion(ctx *Instr_asignacionContext) {}

// EnterInstr_if is called when production instr_if is entered.
func (s *BaseChemsListener) EnterInstr_if(ctx *Instr_ifContext) {}

// ExitInstr_if is called when production instr_if is exited.
func (s *BaseChemsListener) ExitInstr_if(ctx *Instr_ifContext) {}

// EnterInstr_else_if is called when production instr_else_if is entered.
func (s *BaseChemsListener) EnterInstr_else_if(ctx *Instr_else_ifContext) {}

// ExitInstr_else_if is called when production instr_else_if is exited.
func (s *BaseChemsListener) ExitInstr_else_if(ctx *Instr_else_ifContext) {}

// EnterInstr_ternario is called when production instr_ternario is entered.
func (s *BaseChemsListener) EnterInstr_ternario(ctx *Instr_ternarioContext) {}

// ExitInstr_ternario is called when production instr_ternario is exited.
func (s *BaseChemsListener) ExitInstr_ternario(ctx *Instr_ternarioContext) {}

// EnterInstr_else_if_ternario is called when production instr_else_if_ternario is entered.
func (s *BaseChemsListener) EnterInstr_else_if_ternario(ctx *Instr_else_if_ternarioContext) {}

// ExitInstr_else_if_ternario is called when production instr_else_if_ternario is exited.
func (s *BaseChemsListener) ExitInstr_else_if_ternario(ctx *Instr_else_if_ternarioContext) {}

// EnterInstr_match is called when production instr_match is entered.
func (s *BaseChemsListener) EnterInstr_match(ctx *Instr_matchContext) {}

// ExitInstr_match is called when production instr_match is exited.
func (s *BaseChemsListener) ExitInstr_match(ctx *Instr_matchContext) {}

// EnterInstr_case is called when production instr_case is entered.
func (s *BaseChemsListener) EnterInstr_case(ctx *Instr_caseContext) {}

// ExitInstr_case is called when production instr_case is exited.
func (s *BaseChemsListener) ExitInstr_case(ctx *Instr_caseContext) {}

// EnterBlock_instr_match is called when production block_instr_match is entered.
func (s *BaseChemsListener) EnterBlock_instr_match(ctx *Block_instr_matchContext) {}

// ExitBlock_instr_match is called when production block_instr_match is exited.
func (s *BaseChemsListener) ExitBlock_instr_match(ctx *Block_instr_matchContext) {}

// EnterList_case is called when production list_case is entered.
func (s *BaseChemsListener) EnterList_case(ctx *List_caseContext) {}

// ExitList_case is called when production list_case is exited.
func (s *BaseChemsListener) ExitList_case(ctx *List_caseContext) {}

// EnterList_expre_case is called when production list_expre_case is entered.
func (s *BaseChemsListener) EnterList_expre_case(ctx *List_expre_caseContext) {}

// ExitList_expre_case is called when production list_expre_case is exited.
func (s *BaseChemsListener) ExitList_expre_case(ctx *List_expre_caseContext) {}

// EnterBlock_case is called when production block_case is entered.
func (s *BaseChemsListener) EnterBlock_case(ctx *Block_caseContext) {}

// ExitBlock_case is called when production block_case is exited.
func (s *BaseChemsListener) ExitBlock_case(ctx *Block_caseContext) {}

// EnterInstr_default is called when production instr_default is entered.
func (s *BaseChemsListener) EnterInstr_default(ctx *Instr_defaultContext) {}

// ExitInstr_default is called when production instr_default is exited.
func (s *BaseChemsListener) ExitInstr_default(ctx *Instr_defaultContext) {}

// EnterBlock_default is called when production block_default is entered.
func (s *BaseChemsListener) EnterBlock_default(ctx *Block_defaultContext) {}

// ExitBlock_default is called when production block_default is exited.
func (s *BaseChemsListener) ExitBlock_default(ctx *Block_defaultContext) {}

// EnterInstr_match_ter is called when production instr_match_ter is entered.
func (s *BaseChemsListener) EnterInstr_match_ter(ctx *Instr_match_terContext) {}

// ExitInstr_match_ter is called when production instr_match_ter is exited.
func (s *BaseChemsListener) ExitInstr_match_ter(ctx *Instr_match_terContext) {}

// EnterInstr_case_ter is called when production instr_case_ter is entered.
func (s *BaseChemsListener) EnterInstr_case_ter(ctx *Instr_case_terContext) {}

// ExitInstr_case_ter is called when production instr_case_ter is exited.
func (s *BaseChemsListener) ExitInstr_case_ter(ctx *Instr_case_terContext) {}

// EnterList_case_ternario is called when production list_case_ternario is entered.
func (s *BaseChemsListener) EnterList_case_ternario(ctx *List_case_ternarioContext) {}

// ExitList_case_ternario is called when production list_case_ternario is exited.
func (s *BaseChemsListener) ExitList_case_ternario(ctx *List_case_ternarioContext) {}

// EnterList_expre_case_ter is called when production list_expre_case_ter is entered.
func (s *BaseChemsListener) EnterList_expre_case_ter(ctx *List_expre_case_terContext) {}

// ExitList_expre_case_ter is called when production list_expre_case_ter is exited.
func (s *BaseChemsListener) ExitList_expre_case_ter(ctx *List_expre_case_terContext) {}

// EnterBlock_case_ter is called when production block_case_ter is entered.
func (s *BaseChemsListener) EnterBlock_case_ter(ctx *Block_case_terContext) {}

// ExitBlock_case_ter is called when production block_case_ter is exited.
func (s *BaseChemsListener) ExitBlock_case_ter(ctx *Block_case_terContext) {}

// EnterInstr_default_ter is called when production instr_default_ter is entered.
func (s *BaseChemsListener) EnterInstr_default_ter(ctx *Instr_default_terContext) {}

// ExitInstr_default_ter is called when production instr_default_ter is exited.
func (s *BaseChemsListener) ExitInstr_default_ter(ctx *Instr_default_terContext) {}

// EnterInstr_while is called when production instr_while is entered.
func (s *BaseChemsListener) EnterInstr_while(ctx *Instr_whileContext) {}

// ExitInstr_while is called when production instr_while is exited.
func (s *BaseChemsListener) ExitInstr_while(ctx *Instr_whileContext) {}

// EnterInstr_break is called when production instr_break is entered.
func (s *BaseChemsListener) EnterInstr_break(ctx *Instr_breakContext) {}

// ExitInstr_break is called when production instr_break is exited.
func (s *BaseChemsListener) ExitInstr_break(ctx *Instr_breakContext) {}

// EnterInstr_tipo is called when production instr_tipo is entered.
func (s *BaseChemsListener) EnterInstr_tipo(ctx *Instr_tipoContext) {}

// ExitInstr_tipo is called when production instr_tipo is exited.
func (s *BaseChemsListener) ExitInstr_tipo(ctx *Instr_tipoContext) {}

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

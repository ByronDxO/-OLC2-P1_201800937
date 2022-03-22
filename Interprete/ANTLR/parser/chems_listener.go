// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChemsListener is a complete listener for a parse tree produced by Chems.
type ChemsListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterEnd_instr is called when entering the end_instr production.
	EnterEnd_instr(c *End_instrContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterInstr_println is called when entering the instr_println production.
	EnterInstr_println(c *Instr_printlnContext)

	// EnterInstr_declaracion is called when entering the instr_declaracion production.
	EnterInstr_declaracion(c *Instr_declaracionContext)

	// EnterInstr_asignacion is called when entering the instr_asignacion production.
	EnterInstr_asignacion(c *Instr_asignacionContext)

	// EnterInstr_if is called when entering the instr_if production.
	EnterInstr_if(c *Instr_ifContext)

	// EnterInstr_else_if is called when entering the instr_else_if production.
	EnterInstr_else_if(c *Instr_else_ifContext)

	// EnterInstr_ternario is called when entering the instr_ternario production.
	EnterInstr_ternario(c *Instr_ternarioContext)

	// EnterInstr_else_if_ternario is called when entering the instr_else_if_ternario production.
	EnterInstr_else_if_ternario(c *Instr_else_if_ternarioContext)

	// EnterInstr_match is called when entering the instr_match production.
	EnterInstr_match(c *Instr_matchContext)

	// EnterInstr_case is called when entering the instr_case production.
	EnterInstr_case(c *Instr_caseContext)

	// EnterBlock_instr_match is called when entering the block_instr_match production.
	EnterBlock_instr_match(c *Block_instr_matchContext)

	// EnterList_case is called when entering the list_case production.
	EnterList_case(c *List_caseContext)

	// EnterList_expre_case is called when entering the list_expre_case production.
	EnterList_expre_case(c *List_expre_caseContext)

	// EnterBlock_case is called when entering the block_case production.
	EnterBlock_case(c *Block_caseContext)

	// EnterInstr_default is called when entering the instr_default production.
	EnterInstr_default(c *Instr_defaultContext)

	// EnterBlock_default is called when entering the block_default production.
	EnterBlock_default(c *Block_defaultContext)

	// EnterInstr_match_ter is called when entering the instr_match_ter production.
	EnterInstr_match_ter(c *Instr_match_terContext)

	// EnterInstr_case_ter is called when entering the instr_case_ter production.
	EnterInstr_case_ter(c *Instr_case_terContext)

	// EnterList_case_ternario is called when entering the list_case_ternario production.
	EnterList_case_ternario(c *List_case_ternarioContext)

	// EnterList_expre_case_ter is called when entering the list_expre_case_ter production.
	EnterList_expre_case_ter(c *List_expre_case_terContext)

	// EnterBlock_case_ter is called when entering the block_case_ter production.
	EnterBlock_case_ter(c *Block_case_terContext)

	// EnterInstr_default_ter is called when entering the instr_default_ter production.
	EnterInstr_default_ter(c *Instr_default_terContext)

	// EnterInstr_while is called when entering the instr_while production.
	EnterInstr_while(c *Instr_whileContext)

	// EnterInstr_break is called when entering the instr_break production.
	EnterInstr_break(c *Instr_breakContext)

	// EnterInstr_tipo is called when entering the instr_tipo production.
	EnterInstr_tipo(c *Instr_tipoContext)

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

	// ExitEnd_instr is called when exiting the end_instr production.
	ExitEnd_instr(c *End_instrContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitInstr_println is called when exiting the instr_println production.
	ExitInstr_println(c *Instr_printlnContext)

	// ExitInstr_declaracion is called when exiting the instr_declaracion production.
	ExitInstr_declaracion(c *Instr_declaracionContext)

	// ExitInstr_asignacion is called when exiting the instr_asignacion production.
	ExitInstr_asignacion(c *Instr_asignacionContext)

	// ExitInstr_if is called when exiting the instr_if production.
	ExitInstr_if(c *Instr_ifContext)

	// ExitInstr_else_if is called when exiting the instr_else_if production.
	ExitInstr_else_if(c *Instr_else_ifContext)

	// ExitInstr_ternario is called when exiting the instr_ternario production.
	ExitInstr_ternario(c *Instr_ternarioContext)

	// ExitInstr_else_if_ternario is called when exiting the instr_else_if_ternario production.
	ExitInstr_else_if_ternario(c *Instr_else_if_ternarioContext)

	// ExitInstr_match is called when exiting the instr_match production.
	ExitInstr_match(c *Instr_matchContext)

	// ExitInstr_case is called when exiting the instr_case production.
	ExitInstr_case(c *Instr_caseContext)

	// ExitBlock_instr_match is called when exiting the block_instr_match production.
	ExitBlock_instr_match(c *Block_instr_matchContext)

	// ExitList_case is called when exiting the list_case production.
	ExitList_case(c *List_caseContext)

	// ExitList_expre_case is called when exiting the list_expre_case production.
	ExitList_expre_case(c *List_expre_caseContext)

	// ExitBlock_case is called when exiting the block_case production.
	ExitBlock_case(c *Block_caseContext)

	// ExitInstr_default is called when exiting the instr_default production.
	ExitInstr_default(c *Instr_defaultContext)

	// ExitBlock_default is called when exiting the block_default production.
	ExitBlock_default(c *Block_defaultContext)

	// ExitInstr_match_ter is called when exiting the instr_match_ter production.
	ExitInstr_match_ter(c *Instr_match_terContext)

	// ExitInstr_case_ter is called when exiting the instr_case_ter production.
	ExitInstr_case_ter(c *Instr_case_terContext)

	// ExitList_case_ternario is called when exiting the list_case_ternario production.
	ExitList_case_ternario(c *List_case_ternarioContext)

	// ExitList_expre_case_ter is called when exiting the list_expre_case_ter production.
	ExitList_expre_case_ter(c *List_expre_case_terContext)

	// ExitBlock_case_ter is called when exiting the block_case_ter production.
	ExitBlock_case_ter(c *Block_case_terContext)

	// ExitInstr_default_ter is called when exiting the instr_default_ter production.
	ExitInstr_default_ter(c *Instr_default_terContext)

	// ExitInstr_while is called when exiting the instr_while production.
	ExitInstr_while(c *Instr_whileContext)

	// ExitInstr_break is called when exiting the instr_break production.
	ExitInstr_break(c *Instr_breakContext)

	// ExitInstr_tipo is called when exiting the instr_tipo production.
	ExitInstr_tipo(c *Instr_tipoContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExp_arit is called when exiting the exp_arit production.
	ExitExp_arit(c *Exp_aritContext)

	// ExitPrimitivo is called when exiting the primitivo production.
	ExitPrimitivo(c *PrimitivoContext)
}

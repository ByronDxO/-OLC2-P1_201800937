// Code generated from Interprete/ANTLR/Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "OLC2/Interprete/interfaces"
import "OLC2/Interprete/expresion"
import "OLC2/Interprete/instruction"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 73, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 2, 3, 3, 7, 3, 19, 10, 3, 12, 3, 14, 3, 22, 11, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 45, 10, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 7, 6, 62, 10, 6, 12, 6, 14, 6, 65, 11, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 5, 7, 71, 10, 7, 3, 7, 2, 3, 10, 8, 2, 4, 6, 8, 10, 12, 2, 5, 3, 2,
	18, 19, 3, 2, 20, 21, 3, 2, 14, 17, 2, 72, 2, 14, 3, 2, 2, 2, 4, 20, 3,
	2, 2, 2, 6, 25, 3, 2, 2, 2, 8, 32, 3, 2, 2, 2, 10, 44, 3, 2, 2, 2, 12,
	70, 3, 2, 2, 2, 14, 15, 5, 4, 3, 2, 15, 16, 8, 2, 1, 2, 16, 3, 3, 2, 2,
	2, 17, 19, 5, 6, 4, 2, 18, 17, 3, 2, 2, 2, 19, 22, 3, 2, 2, 2, 20, 18,
	3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 23, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2,
	23, 24, 8, 3, 1, 2, 24, 5, 3, 2, 2, 2, 25, 26, 7, 3, 2, 2, 26, 27, 7, 22,
	2, 2, 27, 28, 5, 8, 5, 2, 28, 29, 7, 23, 2, 2, 29, 30, 7, 10, 2, 2, 30,
	31, 8, 4, 1, 2, 31, 7, 3, 2, 2, 2, 32, 33, 5, 10, 6, 2, 33, 34, 8, 5, 1,
	2, 34, 9, 3, 2, 2, 2, 35, 36, 8, 6, 1, 2, 36, 37, 5, 12, 7, 2, 37, 38,
	8, 6, 1, 2, 38, 45, 3, 2, 2, 2, 39, 40, 7, 22, 2, 2, 40, 41, 5, 8, 5, 2,
	41, 42, 7, 23, 2, 2, 42, 43, 8, 6, 1, 2, 43, 45, 3, 2, 2, 2, 44, 35, 3,
	2, 2, 2, 44, 39, 3, 2, 2, 2, 45, 63, 3, 2, 2, 2, 46, 47, 12, 7, 2, 2, 47,
	48, 9, 2, 2, 2, 48, 49, 5, 10, 6, 8, 49, 50, 8, 6, 1, 2, 50, 62, 3, 2,
	2, 2, 51, 52, 12, 6, 2, 2, 52, 53, 9, 3, 2, 2, 53, 54, 5, 10, 6, 7, 54,
	55, 8, 6, 1, 2, 55, 62, 3, 2, 2, 2, 56, 57, 12, 5, 2, 2, 57, 58, 9, 4,
	2, 2, 58, 59, 5, 10, 6, 6, 59, 60, 8, 6, 1, 2, 60, 62, 3, 2, 2, 2, 61,
	46, 3, 2, 2, 2, 61, 51, 3, 2, 2, 2, 61, 56, 3, 2, 2, 2, 62, 65, 3, 2, 2,
	2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 11, 3, 2, 2, 2, 65, 63,
	3, 2, 2, 2, 66, 67, 7, 6, 2, 2, 67, 71, 8, 7, 1, 2, 68, 69, 7, 7, 2, 2,
	69, 71, 8, 7, 1, 2, 70, 66, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 13, 3,
	2, 2, 2, 7, 20, 44, 61, 63, 70,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'println!'", "'number'", "'string'", "", "", "", "'.'", "';'", "','",
	"'!'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'",
	"'('", "')'", "'{'", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "R_PRINTLN", "P_NUMBER", "P_STRING", "NUMBER", "STRING", "ID", "TK_PUNTO",
	"TK_PUNTOCOMA", "TK_COMA", "TK_DIFERENTE", "TK_IGUAL", "TK_MAYORIGUAL",
	"TK_MENORIGUAL", "TK_MAYOR", "TK_MENOR", "TK_MULT", "TK_DIV", "TK_MAS",
	"TK_MENOS", "TK_PARA", "TK_PARC", "TK_LLAVEA", "TK_LLAVEC", "TK_CORA",
	"TK_CORC", "WHITESPACE",
}

var ruleNames = []string{
	"start", "instrucciones", "instruccion", "expression", "exp_arit", "primitivo",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Chems struct {
	*antlr.BaseParser
}

func NewChems(input antlr.TokenStream) *Chems {
	this := new(Chems)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Chems.g4"

	return this
}

// Chems tokens.
const (
	ChemsEOF           = antlr.TokenEOF
	ChemsR_PRINTLN     = 1
	ChemsP_NUMBER      = 2
	ChemsP_STRING      = 3
	ChemsNUMBER        = 4
	ChemsSTRING        = 5
	ChemsID            = 6
	ChemsTK_PUNTO      = 7
	ChemsTK_PUNTOCOMA  = 8
	ChemsTK_COMA       = 9
	ChemsTK_DIFERENTE  = 10
	ChemsTK_IGUAL      = 11
	ChemsTK_MAYORIGUAL = 12
	ChemsTK_MENORIGUAL = 13
	ChemsTK_MAYOR      = 14
	ChemsTK_MENOR      = 15
	ChemsTK_MULT       = 16
	ChemsTK_DIV        = 17
	ChemsTK_MAS        = 18
	ChemsTK_MENOS      = 19
	ChemsTK_PARA       = 20
	ChemsTK_PARC       = 21
	ChemsTK_LLAVEA     = 22
	ChemsTK_LLAVEC     = 23
	ChemsTK_CORA       = 24
	ChemsTK_CORC       = 25
	ChemsWHITESPACE    = 26
)

// Chems rules.
const (
	ChemsRULE_start         = 0
	ChemsRULE_instrucciones = 1
	ChemsRULE_instruccion   = 2
	ChemsRULE_expression    = 3
	ChemsRULE_exp_arit      = 4
	ChemsRULE_primitivo     = 5
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayList.List
	_instrucciones IInstruccionesContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *StartContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *StartContext) GetLista() *arrayList.List { return s.lista }

func (s *StartContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *StartContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Chems) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChemsRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)

		var _x = p.Instrucciones()

		localctx.(*StartContext)._instrucciones = _x
	}
	localctx.(*StartContext).lista = localctx.(*StartContext).Get_instrucciones().GetL()

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccionContext

	// SetE sets the e rule context list.
	SetE([]IInstruccionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	_instruccion IInstruccionContext
	e            []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetE() []IInstruccionContext { return s.e }

func (s *InstruccionesContext) SetE(v []IInstruccionContext) { s.e = v }

func (s *InstruccionesContext) GetL() *arrayList.List { return s.l }

func (s *InstruccionesContext) SetL(v *arrayList.List) { s.l = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *Chems) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChemsRULE_instrucciones)

	localctx.(*InstruccionesContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ChemsR_PRINTLN {
		{
			p.SetState(15)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetE()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruction
	_expression IExpressionContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_expression() IExpressionContext { return s._expression }

func (s *InstruccionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *InstruccionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *InstruccionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *InstruccionContext) R_PRINTLN() antlr.TerminalNode {
	return s.GetToken(ChemsR_PRINTLN, 0)
}

func (s *InstruccionContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *InstruccionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InstruccionContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *InstruccionContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *Chems) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ChemsRULE_instruccion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.Match(ChemsR_PRINTLN)
	}
	{
		p.SetState(24)
		p.Match(ChemsTK_PARA)
	}
	{
		p.SetState(25)

		var _x = p.Expression()

		localctx.(*InstruccionContext)._expression = _x
	}
	{
		p.SetState(26)
		p.Match(ChemsTK_PARC)
	}
	{
		p.SetState(27)
		p.Match(ChemsTK_PUNTOCOMA)
	}
	localctx.(*InstruccionContext).instr = instruction.PRINTLN(localctx.(*InstruccionContext).Get_expression().GetP())

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_exp_arit returns the _exp_arit rule contexts.
	Get_exp_arit() IExp_aritContext

	// Set_exp_arit sets the _exp_arit rule contexts.
	Set_exp_arit(IExp_aritContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	p         interfaces.Expresion
	_exp_arit IExp_aritContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Get_exp_arit() IExp_aritContext { return s._exp_arit }

func (s *ExpressionContext) Set_exp_arit(v IExp_aritContext) { s._exp_arit = v }

func (s *ExpressionContext) GetP() interfaces.Expresion { return s.p }

func (s *ExpressionContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *ExpressionContext) Exp_arit() IExp_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp_aritContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Chems) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChemsRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)

		var _x = p.exp_arit(0)

		localctx.(*ExpressionContext)._exp_arit = _x
	}
	localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_exp_arit().GetP()

	return localctx
}

// IExp_aritContext is an interface to support dynamic dispatch.
type IExp_aritContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpIz returns the opIz rule contexts.
	GetOpIz() IExp_aritContext

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IExp_aritContext

	// SetOpIz sets the opIz rule contexts.
	SetOpIz(IExp_aritContext)

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IExp_aritContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsExp_aritContext differentiates from other interfaces.
	IsExp_aritContext()
}

type Exp_aritContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Expresion
	opIz        IExp_aritContext
	_primitivo  IPrimitivoContext
	_expression IExpressionContext
	op          antlr.Token
	opDe        IExp_aritContext
}

func NewEmptyExp_aritContext() *Exp_aritContext {
	var p = new(Exp_aritContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_exp_arit
	return p
}

func (*Exp_aritContext) IsExp_aritContext() {}

func NewExp_aritContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp_aritContext {
	var p = new(Exp_aritContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_exp_arit

	return p
}

func (s *Exp_aritContext) GetParser() antlr.Parser { return s.parser }

func (s *Exp_aritContext) GetOp() antlr.Token { return s.op }

func (s *Exp_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Exp_aritContext) GetOpIz() IExp_aritContext { return s.opIz }

func (s *Exp_aritContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Exp_aritContext) Get_expression() IExpressionContext { return s._expression }

func (s *Exp_aritContext) GetOpDe() IExp_aritContext { return s.opDe }

func (s *Exp_aritContext) SetOpIz(v IExp_aritContext) { s.opIz = v }

func (s *Exp_aritContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Exp_aritContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Exp_aritContext) SetOpDe(v IExp_aritContext) { s.opDe = v }

func (s *Exp_aritContext) GetP() interfaces.Expresion { return s.p }

func (s *Exp_aritContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *Exp_aritContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Exp_aritContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Exp_aritContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Exp_aritContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Exp_aritContext) AllExp_arit() []IExp_aritContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExp_aritContext)(nil)).Elem())
	var tst = make([]IExp_aritContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExp_aritContext)
		}
	}

	return tst
}

func (s *Exp_aritContext) Exp_arit(i int) IExp_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp_aritContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExp_aritContext)
}

func (s *Exp_aritContext) TK_MULT() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MULT, 0)
}

func (s *Exp_aritContext) TK_DIV() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DIV, 0)
}

func (s *Exp_aritContext) TK_MAS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MAS, 0)
}

func (s *Exp_aritContext) TK_MENOS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENOS, 0)
}

func (s *Exp_aritContext) TK_MENOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENOR, 0)
}

func (s *Exp_aritContext) TK_MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENORIGUAL, 0)
}

func (s *Exp_aritContext) TK_MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MAYORIGUAL, 0)
}

func (s *Exp_aritContext) TK_MAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MAYOR, 0)
}

func (s *Exp_aritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp_aritContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp_aritContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExp_arit(s)
	}
}

func (s *Exp_aritContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExp_arit(s)
	}
}

func (p *Chems) Exp_arit() (localctx IExp_aritContext) {
	return p.exp_arit(0)
}

func (p *Chems) exp_arit(_p int) (localctx IExp_aritContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExp_aritContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExp_aritContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, ChemsRULE_exp_arit, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(42)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsNUMBER, ChemsSTRING:
		{
			p.SetState(34)

			var _x = p.Primitivo()

			localctx.(*Exp_aritContext)._primitivo = _x
		}
		localctx.(*Exp_aritContext).p = localctx.(*Exp_aritContext).Get_primitivo().GetP()

	case ChemsTK_PARA:
		{
			p.SetState(37)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(38)

			var _x = p.Expression()

			localctx.(*Exp_aritContext)._expression = _x
		}
		{
			p.SetState(39)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Exp_aritContext).p = localctx.(*Exp_aritContext).Get_expression().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(44)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(45)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ChemsTK_MULT || _la == ChemsTK_DIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(46)

					var _x = p.exp_arit(6)

					localctx.(*Exp_aritContext).opDe = _x
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetOpIz().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetOpDe().GetP(), false)

			case 2:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(49)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(50)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ChemsTK_MAS || _la == ChemsTK_MENOS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(51)

					var _x = p.exp_arit(5)

					localctx.(*Exp_aritContext).opDe = _x
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetOpIz().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetOpDe().GetP(), false)

			case 3:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(54)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(55)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsTK_MAYORIGUAL)|(1<<ChemsTK_MENORIGUAL)|(1<<ChemsTK_MAYOR)|(1<<ChemsTK_MENOR))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(56)

					var _x = p.exp_arit(4)

					localctx.(*Exp_aritContext).opDe = _x
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetOpIz().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetOpDe().GetP(), false)

			}

		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimitivoContext is an interface to support dynamic dispatch.
type IPrimitivoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	p       interfaces.Expresion
	_NUMBER antlr.Token
	_STRING antlr.Token
}

func NewEmptyPrimitivoContext() *PrimitivoContext {
	var p = new(PrimitivoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_primitivo
	return p
}

func (*PrimitivoContext) IsPrimitivoContext() {}

func NewPrimitivoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivoContext {
	var p = new(PrimitivoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_primitivo

	return p
}

func (s *PrimitivoContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivoContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *PrimitivoContext) Get_STRING() antlr.Token { return s._STRING }

func (s *PrimitivoContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *PrimitivoContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *PrimitivoContext) GetP() interfaces.Expresion { return s.p }

func (s *PrimitivoContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *PrimitivoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChemsNUMBER, 0)
}

func (s *PrimitivoContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChemsSTRING, 0)
}

func (s *PrimitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPrimitivo(s)
	}
}

func (s *PrimitivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPrimitivo(s)
	}
}

func (p *Chems) Primitivo() (localctx IPrimitivoContext) {
	localctx = NewPrimitivoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ChemsRULE_primitivo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)

			var _m = p.Match(ChemsNUMBER)

			localctx.(*PrimitivoContext)._NUMBER = _m
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = expresion.PRIMITIVO(num, interfaces.INTEGER)

	case ChemsSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)

			var _m = p.Match(ChemsSTRING)

			localctx.(*PrimitivoContext)._STRING = _m
		}

		str := (func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}()))-1]
		localctx.(*PrimitivoContext).p = expresion.PRIMITIVO(str, interfaces.STRING)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *Exp_aritContext = nil
		if localctx != nil {
			t = localctx.(*Exp_aritContext)
		}
		return p.Exp_arit_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Chems) Exp_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

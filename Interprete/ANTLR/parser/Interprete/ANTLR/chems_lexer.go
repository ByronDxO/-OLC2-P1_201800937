// Code generated from Interprete/ANTLR/ChemsLexer.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 27, 147,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 6, 5, 80,
	10, 5, 13, 5, 14, 5, 81, 3, 6, 3, 6, 7, 6, 86, 10, 6, 12, 6, 14, 6, 89,
	11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 7, 7, 95, 10, 7, 12, 7, 14, 7, 98, 11, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 6, 26, 139, 10, 26, 13,
	26, 14, 26, 140, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 2, 2, 28, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43,
	23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 2, 3, 2, 8, 3, 2, 50, 59, 3, 2,
	36, 36, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99,
	124, 6, 2, 11, 12, 15, 15, 34, 34, 94, 94, 9, 2, 34, 35, 37, 37, 45, 45,
	47, 48, 60, 60, 66, 66, 93, 95, 2, 149, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2,
	37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2,
	2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2,
	2, 3, 55, 3, 2, 2, 2, 5, 64, 3, 2, 2, 2, 7, 71, 3, 2, 2, 2, 9, 79, 3, 2,
	2, 2, 11, 83, 3, 2, 2, 2, 13, 92, 3, 2, 2, 2, 15, 99, 3, 2, 2, 2, 17, 101,
	3, 2, 2, 2, 19, 103, 3, 2, 2, 2, 21, 105, 3, 2, 2, 2, 23, 107, 3, 2, 2,
	2, 25, 110, 3, 2, 2, 2, 27, 113, 3, 2, 2, 2, 29, 115, 3, 2, 2, 2, 31, 117,
	3, 2, 2, 2, 33, 119, 3, 2, 2, 2, 35, 121, 3, 2, 2, 2, 37, 123, 3, 2, 2,
	2, 39, 125, 3, 2, 2, 2, 41, 127, 3, 2, 2, 2, 43, 129, 3, 2, 2, 2, 45, 131,
	3, 2, 2, 2, 47, 133, 3, 2, 2, 2, 49, 135, 3, 2, 2, 2, 51, 138, 3, 2, 2,
	2, 53, 144, 3, 2, 2, 2, 55, 56, 7, 114, 2, 2, 56, 57, 7, 116, 2, 2, 57,
	58, 7, 107, 2, 2, 58, 59, 7, 112, 2, 2, 59, 60, 7, 118, 2, 2, 60, 61, 7,
	110, 2, 2, 61, 62, 7, 112, 2, 2, 62, 63, 7, 35, 2, 2, 63, 4, 3, 2, 2, 2,
	64, 65, 7, 112, 2, 2, 65, 66, 7, 119, 2, 2, 66, 67, 7, 111, 2, 2, 67, 68,
	7, 100, 2, 2, 68, 69, 7, 103, 2, 2, 69, 70, 7, 116, 2, 2, 70, 6, 3, 2,
	2, 2, 71, 72, 7, 117, 2, 2, 72, 73, 7, 118, 2, 2, 73, 74, 7, 116, 2, 2,
	74, 75, 7, 107, 2, 2, 75, 76, 7, 112, 2, 2, 76, 77, 7, 105, 2, 2, 77, 8,
	3, 2, 2, 2, 78, 80, 9, 2, 2, 2, 79, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2,
	81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 10, 3, 2, 2, 2, 83, 87, 7,
	36, 2, 2, 84, 86, 10, 3, 2, 2, 85, 84, 3, 2, 2, 2, 86, 89, 3, 2, 2, 2,
	87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 90, 3, 2, 2, 2, 89, 87, 3,
	2, 2, 2, 90, 91, 7, 36, 2, 2, 91, 12, 3, 2, 2, 2, 92, 96, 9, 4, 2, 2, 93,
	95, 9, 5, 2, 2, 94, 93, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94, 3, 2, 2,
	2, 96, 97, 3, 2, 2, 2, 97, 14, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99, 100,
	7, 48, 2, 2, 100, 16, 3, 2, 2, 2, 101, 102, 7, 61, 2, 2, 102, 18, 3, 2,
	2, 2, 103, 104, 7, 46, 2, 2, 104, 20, 3, 2, 2, 2, 105, 106, 7, 63, 2, 2,
	106, 22, 3, 2, 2, 2, 107, 108, 7, 64, 2, 2, 108, 109, 7, 63, 2, 2, 109,
	24, 3, 2, 2, 2, 110, 111, 7, 62, 2, 2, 111, 112, 7, 63, 2, 2, 112, 26,
	3, 2, 2, 2, 113, 114, 7, 64, 2, 2, 114, 28, 3, 2, 2, 2, 115, 116, 7, 62,
	2, 2, 116, 30, 3, 2, 2, 2, 117, 118, 7, 44, 2, 2, 118, 32, 3, 2, 2, 2,
	119, 120, 7, 49, 2, 2, 120, 34, 3, 2, 2, 2, 121, 122, 7, 45, 2, 2, 122,
	36, 3, 2, 2, 2, 123, 124, 7, 47, 2, 2, 124, 38, 3, 2, 2, 2, 125, 126, 7,
	42, 2, 2, 126, 40, 3, 2, 2, 2, 127, 128, 7, 43, 2, 2, 128, 42, 3, 2, 2,
	2, 129, 130, 7, 125, 2, 2, 130, 44, 3, 2, 2, 2, 131, 132, 7, 127, 2, 2,
	132, 46, 3, 2, 2, 2, 133, 134, 7, 93, 2, 2, 134, 48, 3, 2, 2, 2, 135, 136,
	7, 95, 2, 2, 136, 50, 3, 2, 2, 2, 137, 139, 9, 6, 2, 2, 138, 137, 3, 2,
	2, 2, 139, 140, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2,
	141, 142, 3, 2, 2, 2, 142, 143, 8, 26, 2, 2, 143, 52, 3, 2, 2, 2, 144,
	145, 7, 94, 2, 2, 145, 146, 9, 7, 2, 2, 146, 54, 3, 2, 2, 2, 7, 2, 81,
	87, 96, 140, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'println!'", "'number'", "'string'", "", "", "", "'.'", "';'", "','",
	"'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'('",
	"')'", "'{'", "'}'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "R_PRINTLN", "P_NUMBER", "P_STRING", "NUMBER", "STRING", "ID", "TK_PUNTO",
	"TK_PUNTOCOMA", "TK_COMA", "TK_IGUAL", "TK_MAYORIGUAL", "TK_MENORIGUAL",
	"TK_MAYOR", "TK_MENOR", "TK_MULT", "TK_DIV", "TK_MAS", "TK_MENOS", "TK_PARA",
	"TK_PARC", "TK_LLAVEA", "TK_LLAVEC", "TK_CORA", "TK_CORC", "WHITESPACE",
}

var lexerRuleNames = []string{
	"R_PRINTLN", "P_NUMBER", "P_STRING", "NUMBER", "STRING", "ID", "TK_PUNTO",
	"TK_PUNTOCOMA", "TK_COMA", "TK_IGUAL", "TK_MAYORIGUAL", "TK_MENORIGUAL",
	"TK_MAYOR", "TK_MENOR", "TK_MULT", "TK_DIV", "TK_MAS", "TK_MENOS", "TK_PARA",
	"TK_PARC", "TK_LLAVEA", "TK_LLAVEC", "TK_CORA", "TK_CORC", "WHITESPACE",
	"ESC_SEQ",
}

type ChemsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewChemsLexer(input antlr.CharStream) *ChemsLexer {

	l := new(ChemsLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ChemsLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ChemsLexer tokens.
const (
	ChemsLexerR_PRINTLN     = 1
	ChemsLexerP_NUMBER      = 2
	ChemsLexerP_STRING      = 3
	ChemsLexerNUMBER        = 4
	ChemsLexerSTRING        = 5
	ChemsLexerID            = 6
	ChemsLexerTK_PUNTO      = 7
	ChemsLexerTK_PUNTOCOMA  = 8
	ChemsLexerTK_COMA       = 9
	ChemsLexerTK_IGUAL      = 10
	ChemsLexerTK_MAYORIGUAL = 11
	ChemsLexerTK_MENORIGUAL = 12
	ChemsLexerTK_MAYOR      = 13
	ChemsLexerTK_MENOR      = 14
	ChemsLexerTK_MULT       = 15
	ChemsLexerTK_DIV        = 16
	ChemsLexerTK_MAS        = 17
	ChemsLexerTK_MENOS      = 18
	ChemsLexerTK_PARA       = 19
	ChemsLexerTK_PARC       = 20
	ChemsLexerTK_LLAVEA     = 21
	ChemsLexerTK_LLAVEC     = 22
	ChemsLexerTK_CORA       = 23
	ChemsLexerTK_CORC       = 24
	ChemsLexerWHITESPACE    = 25
)

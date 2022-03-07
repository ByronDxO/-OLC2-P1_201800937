// Code generated from ChemsLexer.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 28, 151,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 6, 5, 82, 10, 5, 13, 5, 14, 5, 83, 3, 6, 3, 6, 7, 6, 88, 10, 6, 12,
	6, 14, 6, 91, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 7, 7, 97, 10, 7, 12, 7, 14,
	7, 100, 11, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	27, 6, 27, 143, 10, 27, 13, 27, 14, 27, 144, 3, 27, 3, 27, 3, 28, 3, 28,
	3, 28, 2, 2, 29, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55,
	2, 3, 2, 8, 3, 2, 50, 59, 3, 2, 36, 36, 5, 2, 67, 92, 97, 97, 99, 124,
	6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 6, 2, 11, 12, 15, 15, 34, 34, 94,
	94, 9, 2, 34, 35, 37, 37, 45, 45, 47, 48, 60, 60, 66, 66, 93, 95, 2, 153,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3,
	2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41,
	3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2,
	49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 3, 57, 3, 2, 2, 2,
	5, 66, 3, 2, 2, 2, 7, 73, 3, 2, 2, 2, 9, 81, 3, 2, 2, 2, 11, 85, 3, 2,
	2, 2, 13, 94, 3, 2, 2, 2, 15, 101, 3, 2, 2, 2, 17, 103, 3, 2, 2, 2, 19,
	105, 3, 2, 2, 2, 21, 107, 3, 2, 2, 2, 23, 109, 3, 2, 2, 2, 25, 111, 3,
	2, 2, 2, 27, 114, 3, 2, 2, 2, 29, 117, 3, 2, 2, 2, 31, 119, 3, 2, 2, 2,
	33, 121, 3, 2, 2, 2, 35, 123, 3, 2, 2, 2, 37, 125, 3, 2, 2, 2, 39, 127,
	3, 2, 2, 2, 41, 129, 3, 2, 2, 2, 43, 131, 3, 2, 2, 2, 45, 133, 3, 2, 2,
	2, 47, 135, 3, 2, 2, 2, 49, 137, 3, 2, 2, 2, 51, 139, 3, 2, 2, 2, 53, 142,
	3, 2, 2, 2, 55, 148, 3, 2, 2, 2, 57, 58, 7, 114, 2, 2, 58, 59, 7, 116,
	2, 2, 59, 60, 7, 107, 2, 2, 60, 61, 7, 112, 2, 2, 61, 62, 7, 118, 2, 2,
	62, 63, 7, 110, 2, 2, 63, 64, 7, 112, 2, 2, 64, 65, 7, 35, 2, 2, 65, 4,
	3, 2, 2, 2, 66, 67, 7, 112, 2, 2, 67, 68, 7, 119, 2, 2, 68, 69, 7, 111,
	2, 2, 69, 70, 7, 100, 2, 2, 70, 71, 7, 103, 2, 2, 71, 72, 7, 116, 2, 2,
	72, 6, 3, 2, 2, 2, 73, 74, 7, 117, 2, 2, 74, 75, 7, 118, 2, 2, 75, 76,
	7, 116, 2, 2, 76, 77, 7, 107, 2, 2, 77, 78, 7, 112, 2, 2, 78, 79, 7, 105,
	2, 2, 79, 8, 3, 2, 2, 2, 80, 82, 9, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82, 83,
	3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 10, 3, 2, 2, 2,
	85, 89, 7, 36, 2, 2, 86, 88, 10, 3, 2, 2, 87, 86, 3, 2, 2, 2, 88, 91, 3,
	2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 92, 3, 2, 2, 2, 91,
	89, 3, 2, 2, 2, 92, 93, 7, 36, 2, 2, 93, 12, 3, 2, 2, 2, 94, 98, 9, 4,
	2, 2, 95, 97, 9, 5, 2, 2, 96, 95, 3, 2, 2, 2, 97, 100, 3, 2, 2, 2, 98,
	96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 14, 3, 2, 2, 2, 100, 98, 3, 2,
	2, 2, 101, 102, 7, 48, 2, 2, 102, 16, 3, 2, 2, 2, 103, 104, 7, 61, 2, 2,
	104, 18, 3, 2, 2, 2, 105, 106, 7, 46, 2, 2, 106, 20, 3, 2, 2, 2, 107, 108,
	7, 35, 2, 2, 108, 22, 3, 2, 2, 2, 109, 110, 7, 63, 2, 2, 110, 24, 3, 2,
	2, 2, 111, 112, 7, 64, 2, 2, 112, 113, 7, 63, 2, 2, 113, 26, 3, 2, 2, 2,
	114, 115, 7, 62, 2, 2, 115, 116, 7, 63, 2, 2, 116, 28, 3, 2, 2, 2, 117,
	118, 7, 64, 2, 2, 118, 30, 3, 2, 2, 2, 119, 120, 7, 62, 2, 2, 120, 32,
	3, 2, 2, 2, 121, 122, 7, 44, 2, 2, 122, 34, 3, 2, 2, 2, 123, 124, 7, 49,
	2, 2, 124, 36, 3, 2, 2, 2, 125, 126, 7, 45, 2, 2, 126, 38, 3, 2, 2, 2,
	127, 128, 7, 47, 2, 2, 128, 40, 3, 2, 2, 2, 129, 130, 7, 42, 2, 2, 130,
	42, 3, 2, 2, 2, 131, 132, 7, 43, 2, 2, 132, 44, 3, 2, 2, 2, 133, 134, 7,
	125, 2, 2, 134, 46, 3, 2, 2, 2, 135, 136, 7, 127, 2, 2, 136, 48, 3, 2,
	2, 2, 137, 138, 7, 93, 2, 2, 138, 50, 3, 2, 2, 2, 139, 140, 7, 95, 2, 2,
	140, 52, 3, 2, 2, 2, 141, 143, 9, 6, 2, 2, 142, 141, 3, 2, 2, 2, 143, 144,
	3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 146, 3, 2,
	2, 2, 146, 147, 8, 27, 2, 2, 147, 54, 3, 2, 2, 2, 148, 149, 7, 94, 2, 2,
	149, 150, 9, 7, 2, 2, 150, 56, 3, 2, 2, 2, 7, 2, 83, 89, 98, 144, 3, 8,
	2, 2,
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
	"'!'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'",
	"'('", "')'", "'{'", "'}'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "R_PRINTLN", "P_NUMBER", "P_STRING", "NUMBER", "STRING", "ID", "TK_PUNTO",
	"TK_PUNTOCOMA", "TK_COMA", "TK_DIFERENTE", "TK_IGUAL", "TK_MAYORIGUAL",
	"TK_MENORIGUAL", "TK_MAYOR", "TK_MENOR", "TK_MULT", "TK_DIV", "TK_MAS",
	"TK_MENOS", "TK_PARA", "TK_PARC", "TK_LLAVEA", "TK_LLAVEC", "TK_CORA",
	"TK_CORC", "WHITESPACE",
}

var lexerRuleNames = []string{
	"R_PRINTLN", "P_NUMBER", "P_STRING", "NUMBER", "STRING", "ID", "TK_PUNTO",
	"TK_PUNTOCOMA", "TK_COMA", "TK_DIFERENTE", "TK_IGUAL", "TK_MAYORIGUAL",
	"TK_MENORIGUAL", "TK_MAYOR", "TK_MENOR", "TK_MULT", "TK_DIV", "TK_MAS",
	"TK_MENOS", "TK_PARA", "TK_PARC", "TK_LLAVEA", "TK_LLAVEC", "TK_CORA",
	"TK_CORC", "WHITESPACE", "ESC_SEQ",
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
	ChemsLexerTK_DIFERENTE  = 10
	ChemsLexerTK_IGUAL      = 11
	ChemsLexerTK_MAYORIGUAL = 12
	ChemsLexerTK_MENORIGUAL = 13
	ChemsLexerTK_MAYOR      = 14
	ChemsLexerTK_MENOR      = 15
	ChemsLexerTK_MULT       = 16
	ChemsLexerTK_DIV        = 17
	ChemsLexerTK_MAS        = 18
	ChemsLexerTK_MENOS      = 19
	ChemsLexerTK_PARA       = 20
	ChemsLexerTK_PARC       = 21
	ChemsLexerTK_LLAVEA     = 22
	ChemsLexerTK_LLAVEC     = 23
	ChemsLexerTK_CORA       = 24
	ChemsLexerTK_CORC       = 25
	ChemsLexerWHITESPACE    = 26
)

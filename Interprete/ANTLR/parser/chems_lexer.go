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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 30, 179,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 6, 5, 86, 10, 5, 13, 5, 14, 5, 87, 3,
	6, 6, 6, 91, 10, 6, 13, 6, 14, 6, 92, 3, 6, 3, 6, 6, 6, 97, 10, 6, 13,
	6, 14, 6, 98, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 7, 8, 107, 10, 8, 12,
	8, 14, 8, 110, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 5, 9, 123, 10, 9, 3, 10, 3, 10, 7, 10, 127, 10, 10, 12,
	10, 14, 10, 130, 11, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3,
	29, 6, 29, 171, 10, 29, 13, 29, 14, 29, 172, 3, 29, 3, 29, 3, 30, 3, 30,
	3, 30, 2, 2, 31, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55,
	29, 57, 30, 59, 2, 3, 2, 8, 3, 2, 50, 59, 3, 2, 36, 36, 5, 2, 67, 92, 97,
	97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15,
	34, 34, 9, 2, 34, 35, 37, 37, 45, 45, 47, 48, 60, 60, 66, 66, 93, 95, 2,
	184, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2,
	33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2,
	2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2,
	2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2,
	2, 2, 2, 57, 3, 2, 2, 2, 3, 61, 3, 2, 2, 2, 5, 70, 3, 2, 2, 2, 7, 77, 3,
	2, 2, 2, 9, 85, 3, 2, 2, 2, 11, 90, 3, 2, 2, 2, 13, 100, 3, 2, 2, 2, 15,
	104, 3, 2, 2, 2, 17, 122, 3, 2, 2, 2, 19, 124, 3, 2, 2, 2, 21, 131, 3,
	2, 2, 2, 23, 133, 3, 2, 2, 2, 25, 135, 3, 2, 2, 2, 27, 137, 3, 2, 2, 2,
	29, 139, 3, 2, 2, 2, 31, 142, 3, 2, 2, 2, 33, 145, 3, 2, 2, 2, 35, 147,
	3, 2, 2, 2, 37, 149, 3, 2, 2, 2, 39, 151, 3, 2, 2, 2, 41, 153, 3, 2, 2,
	2, 43, 155, 3, 2, 2, 2, 45, 157, 3, 2, 2, 2, 47, 159, 3, 2, 2, 2, 49, 161,
	3, 2, 2, 2, 51, 163, 3, 2, 2, 2, 53, 165, 3, 2, 2, 2, 55, 167, 3, 2, 2,
	2, 57, 170, 3, 2, 2, 2, 59, 176, 3, 2, 2, 2, 61, 62, 7, 114, 2, 2, 62,
	63, 7, 116, 2, 2, 63, 64, 7, 107, 2, 2, 64, 65, 7, 112, 2, 2, 65, 66, 7,
	118, 2, 2, 66, 67, 7, 110, 2, 2, 67, 68, 7, 112, 2, 2, 68, 69, 7, 35, 2,
	2, 69, 4, 3, 2, 2, 2, 70, 71, 7, 112, 2, 2, 71, 72, 7, 119, 2, 2, 72, 73,
	7, 111, 2, 2, 73, 74, 7, 100, 2, 2, 74, 75, 7, 103, 2, 2, 75, 76, 7, 116,
	2, 2, 76, 6, 3, 2, 2, 2, 77, 78, 7, 117, 2, 2, 78, 79, 7, 118, 2, 2, 79,
	80, 7, 116, 2, 2, 80, 81, 7, 107, 2, 2, 81, 82, 7, 112, 2, 2, 82, 83, 7,
	105, 2, 2, 83, 8, 3, 2, 2, 2, 84, 86, 9, 2, 2, 2, 85, 84, 3, 2, 2, 2, 86,
	87, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 10, 3, 2, 2,
	2, 89, 91, 9, 2, 2, 2, 90, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 90,
	3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 96, 7, 48, 2, 2,
	95, 97, 9, 2, 2, 2, 96, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 96, 3,
	2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 12, 3, 2, 2, 2, 100, 101, 7, 41, 2, 2,
	101, 102, 10, 3, 2, 2, 102, 103, 7, 41, 2, 2, 103, 14, 3, 2, 2, 2, 104,
	108, 7, 36, 2, 2, 105, 107, 10, 3, 2, 2, 106, 105, 3, 2, 2, 2, 107, 110,
	3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 111, 3, 2,
	2, 2, 110, 108, 3, 2, 2, 2, 111, 112, 7, 36, 2, 2, 112, 16, 3, 2, 2, 2,
	113, 114, 7, 118, 2, 2, 114, 115, 7, 116, 2, 2, 115, 116, 7, 119, 2, 2,
	116, 123, 7, 103, 2, 2, 117, 118, 7, 104, 2, 2, 118, 119, 7, 99, 2, 2,
	119, 120, 7, 110, 2, 2, 120, 121, 7, 117, 2, 2, 121, 123, 7, 103, 2, 2,
	122, 113, 3, 2, 2, 2, 122, 117, 3, 2, 2, 2, 123, 18, 3, 2, 2, 2, 124, 128,
	9, 4, 2, 2, 125, 127, 9, 5, 2, 2, 126, 125, 3, 2, 2, 2, 127, 130, 3, 2,
	2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 20, 3, 2, 2, 2,
	130, 128, 3, 2, 2, 2, 131, 132, 7, 48, 2, 2, 132, 22, 3, 2, 2, 2, 133,
	134, 7, 61, 2, 2, 134, 24, 3, 2, 2, 2, 135, 136, 7, 46, 2, 2, 136, 26,
	3, 2, 2, 2, 137, 138, 7, 63, 2, 2, 138, 28, 3, 2, 2, 2, 139, 140, 7, 64,
	2, 2, 140, 141, 7, 63, 2, 2, 141, 30, 3, 2, 2, 2, 142, 143, 7, 62, 2, 2,
	143, 144, 7, 63, 2, 2, 144, 32, 3, 2, 2, 2, 145, 146, 7, 64, 2, 2, 146,
	34, 3, 2, 2, 2, 147, 148, 7, 62, 2, 2, 148, 36, 3, 2, 2, 2, 149, 150, 7,
	44, 2, 2, 150, 38, 3, 2, 2, 2, 151, 152, 7, 49, 2, 2, 152, 40, 3, 2, 2,
	2, 153, 154, 7, 45, 2, 2, 154, 42, 3, 2, 2, 2, 155, 156, 7, 47, 2, 2, 156,
	44, 3, 2, 2, 2, 157, 158, 7, 42, 2, 2, 158, 46, 3, 2, 2, 2, 159, 160, 7,
	43, 2, 2, 160, 48, 3, 2, 2, 2, 161, 162, 7, 125, 2, 2, 162, 50, 3, 2, 2,
	2, 163, 164, 7, 127, 2, 2, 164, 52, 3, 2, 2, 2, 165, 166, 7, 93, 2, 2,
	166, 54, 3, 2, 2, 2, 167, 168, 7, 95, 2, 2, 168, 56, 3, 2, 2, 2, 169, 171,
	9, 6, 2, 2, 170, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 170, 3, 2,
	2, 2, 172, 173, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 175, 8, 29, 2, 2,
	175, 58, 3, 2, 2, 2, 176, 177, 7, 94, 2, 2, 177, 178, 9, 7, 2, 2, 178,
	60, 3, 2, 2, 2, 10, 2, 87, 92, 98, 108, 122, 128, 172, 3, 8, 2, 2,
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
	"", "'println!'", "'number'", "'string'", "", "", "", "", "", "", "'.'",
	"';'", "','", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'",
	"'-'", "'('", "')'", "'{'", "'}'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "R_PRINTLN", "P_NUMBER", "P_STRING", "NUMBER", "DOUBLE", "CHAR", "STRING",
	"BOOLEAN", "ID", "TK_PUNTO", "TK_PUNTOCOMA", "TK_COMA", "TK_IGUAL", "TK_MAYORIGUAL",
	"TK_MENORIGUAL", "TK_MAYOR", "TK_MENOR", "TK_MULT", "TK_DIV", "TK_MAS",
	"TK_MENOS", "TK_PARA", "TK_PARC", "TK_LLAVEA", "TK_LLAVEC", "TK_CORA",
	"TK_CORC", "WHITESPACE",
}

var lexerRuleNames = []string{
	"R_PRINTLN", "P_NUMBER", "P_STRING", "NUMBER", "DOUBLE", "CHAR", "STRING",
	"BOOLEAN", "ID", "TK_PUNTO", "TK_PUNTOCOMA", "TK_COMA", "TK_IGUAL", "TK_MAYORIGUAL",
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
	ChemsLexerDOUBLE        = 5
	ChemsLexerCHAR          = 6
	ChemsLexerSTRING        = 7
	ChemsLexerBOOLEAN       = 8
	ChemsLexerID            = 9
	ChemsLexerTK_PUNTO      = 10
	ChemsLexerTK_PUNTOCOMA  = 11
	ChemsLexerTK_COMA       = 12
	ChemsLexerTK_IGUAL      = 13
	ChemsLexerTK_MAYORIGUAL = 14
	ChemsLexerTK_MENORIGUAL = 15
	ChemsLexerTK_MAYOR      = 16
	ChemsLexerTK_MENOR      = 17
	ChemsLexerTK_MULT       = 18
	ChemsLexerTK_DIV        = 19
	ChemsLexerTK_MAS        = 20
	ChemsLexerTK_MENOS      = 21
	ChemsLexerTK_PARA       = 22
	ChemsLexerTK_PARC       = 23
	ChemsLexerTK_LLAVEA     = 24
	ChemsLexerTK_LLAVEC     = 25
	ChemsLexerTK_CORA       = 26
	ChemsLexerTK_CORC       = 27
	ChemsLexerWHITESPACE    = 28
)

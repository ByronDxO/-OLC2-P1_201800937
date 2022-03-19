// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
import "OLC2/Interprete/instruction/variable"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 44, 249,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 25,
	10, 3, 12, 3, 14, 3, 28, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 54, 10, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 90, 10, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 104, 10,
	7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 186, 10, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 232, 10, 9, 12, 9, 14,
	9, 235, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 5, 10, 247, 10, 10, 3, 10, 2, 3, 16, 11, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 2, 8, 3, 2, 5, 6, 3, 2, 28, 30, 3, 2, 31, 32, 3, 2, 23,
	27, 4, 2, 32, 32, 41, 41, 3, 2, 39, 40, 2, 268, 2, 20, 3, 2, 2, 2, 4, 26,
	3, 2, 2, 2, 6, 53, 3, 2, 2, 2, 8, 89, 3, 2, 2, 2, 10, 91, 3, 2, 2, 2, 12,
	103, 3, 2, 2, 2, 14, 105, 3, 2, 2, 2, 16, 185, 3, 2, 2, 2, 18, 246, 3,
	2, 2, 2, 20, 21, 5, 4, 3, 2, 21, 22, 8, 2, 1, 2, 22, 3, 3, 2, 2, 2, 23,
	25, 5, 6, 4, 2, 24, 23, 3, 2, 2, 2, 25, 28, 3, 2, 2, 2, 26, 24, 3, 2, 2,
	2, 26, 27, 3, 2, 2, 2, 27, 29, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 29, 30,
	8, 3, 1, 2, 30, 5, 3, 2, 2, 2, 31, 32, 7, 3, 2, 2, 32, 33, 7, 33, 2, 2,
	33, 34, 5, 14, 8, 2, 34, 35, 7, 34, 2, 2, 35, 36, 7, 19, 2, 2, 36, 37,
	8, 4, 1, 2, 37, 54, 3, 2, 2, 2, 38, 39, 7, 3, 2, 2, 39, 40, 7, 33, 2, 2,
	40, 41, 7, 15, 2, 2, 41, 42, 7, 20, 2, 2, 42, 43, 5, 14, 8, 2, 43, 44,
	7, 34, 2, 2, 44, 45, 7, 19, 2, 2, 45, 46, 8, 4, 1, 2, 46, 54, 3, 2, 2,
	2, 47, 48, 5, 8, 5, 2, 48, 49, 8, 4, 1, 2, 49, 54, 3, 2, 2, 2, 50, 51,
	5, 10, 6, 2, 51, 52, 8, 4, 1, 2, 52, 54, 3, 2, 2, 2, 53, 31, 3, 2, 2, 2,
	53, 38, 3, 2, 2, 2, 53, 47, 3, 2, 2, 2, 53, 50, 3, 2, 2, 2, 54, 7, 3, 2,
	2, 2, 55, 56, 7, 7, 2, 2, 56, 57, 7, 8, 2, 2, 57, 58, 7, 17, 2, 2, 58,
	59, 7, 22, 2, 2, 59, 60, 5, 14, 8, 2, 60, 61, 7, 19, 2, 2, 61, 62, 8, 5,
	1, 2, 62, 90, 3, 2, 2, 2, 63, 64, 7, 7, 2, 2, 64, 65, 7, 8, 2, 2, 65, 66,
	7, 17, 2, 2, 66, 67, 7, 21, 2, 2, 67, 68, 5, 12, 7, 2, 68, 69, 7, 22, 2,
	2, 69, 70, 5, 14, 8, 2, 70, 71, 7, 19, 2, 2, 71, 72, 8, 5, 1, 2, 72, 90,
	3, 2, 2, 2, 73, 74, 7, 7, 2, 2, 74, 75, 7, 17, 2, 2, 75, 76, 7, 22, 2,
	2, 76, 77, 5, 14, 8, 2, 77, 78, 7, 19, 2, 2, 78, 79, 8, 5, 1, 2, 79, 90,
	3, 2, 2, 2, 80, 81, 7, 7, 2, 2, 81, 82, 7, 17, 2, 2, 82, 83, 7, 21, 2,
	2, 83, 84, 5, 12, 7, 2, 84, 85, 7, 22, 2, 2, 85, 86, 5, 14, 8, 2, 86, 87,
	7, 19, 2, 2, 87, 88, 8, 5, 1, 2, 88, 90, 3, 2, 2, 2, 89, 55, 3, 2, 2, 2,
	89, 63, 3, 2, 2, 2, 89, 73, 3, 2, 2, 2, 89, 80, 3, 2, 2, 2, 90, 9, 3, 2,
	2, 2, 91, 92, 7, 17, 2, 2, 92, 93, 7, 22, 2, 2, 93, 94, 5, 14, 8, 2, 94,
	95, 7, 19, 2, 2, 95, 96, 8, 6, 1, 2, 96, 11, 3, 2, 2, 2, 97, 98, 7, 9,
	2, 2, 98, 104, 8, 7, 1, 2, 99, 100, 7, 10, 2, 2, 100, 104, 8, 7, 1, 2,
	101, 102, 7, 11, 2, 2, 102, 104, 8, 7, 1, 2, 103, 97, 3, 2, 2, 2, 103,
	99, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 104, 13, 3, 2, 2, 2, 105, 106, 5,
	16, 9, 2, 106, 107, 8, 8, 1, 2, 107, 15, 3, 2, 2, 2, 108, 109, 8, 9, 1,
	2, 109, 110, 7, 33, 2, 2, 110, 111, 5, 16, 9, 2, 111, 112, 9, 2, 2, 2,
	112, 113, 7, 34, 2, 2, 113, 114, 9, 3, 2, 2, 114, 115, 5, 16, 9, 18, 115,
	116, 8, 9, 1, 2, 116, 186, 3, 2, 2, 2, 117, 118, 7, 33, 2, 2, 118, 119,
	5, 16, 9, 2, 119, 120, 9, 2, 2, 2, 120, 121, 7, 34, 2, 2, 121, 122, 9,
	3, 2, 2, 122, 123, 7, 33, 2, 2, 123, 124, 5, 16, 9, 2, 124, 125, 9, 2,
	2, 2, 125, 126, 7, 34, 2, 2, 126, 127, 8, 9, 1, 2, 127, 186, 3, 2, 2, 2,
	128, 129, 7, 33, 2, 2, 129, 130, 5, 16, 9, 2, 130, 131, 9, 2, 2, 2, 131,
	132, 7, 34, 2, 2, 132, 133, 9, 4, 2, 2, 133, 134, 5, 16, 9, 14, 134, 135,
	8, 9, 1, 2, 135, 186, 3, 2, 2, 2, 136, 137, 7, 33, 2, 2, 137, 138, 5, 16,
	9, 2, 138, 139, 9, 2, 2, 2, 139, 140, 7, 34, 2, 2, 140, 141, 9, 4, 2, 2,
	141, 142, 7, 33, 2, 2, 142, 143, 5, 16, 9, 2, 143, 144, 9, 2, 2, 2, 144,
	145, 7, 34, 2, 2, 145, 146, 8, 9, 1, 2, 146, 186, 3, 2, 2, 2, 147, 148,
	7, 33, 2, 2, 148, 149, 5, 16, 9, 2, 149, 150, 9, 2, 2, 2, 150, 151, 7,
	34, 2, 2, 151, 152, 9, 5, 2, 2, 152, 153, 5, 16, 9, 10, 153, 154, 8, 9,
	1, 2, 154, 186, 3, 2, 2, 2, 155, 156, 7, 33, 2, 2, 156, 157, 5, 16, 9,
	2, 157, 158, 9, 2, 2, 2, 158, 159, 7, 34, 2, 2, 159, 160, 9, 5, 2, 2, 160,
	161, 7, 33, 2, 2, 161, 162, 5, 16, 9, 2, 162, 163, 9, 2, 2, 2, 163, 164,
	7, 34, 2, 2, 164, 165, 8, 9, 1, 2, 165, 186, 3, 2, 2, 2, 166, 167, 9, 6,
	2, 2, 167, 168, 5, 14, 8, 2, 168, 169, 8, 9, 1, 2, 169, 186, 3, 2, 2, 2,
	170, 171, 9, 6, 2, 2, 171, 172, 7, 33, 2, 2, 172, 173, 5, 16, 9, 2, 173,
	174, 9, 2, 2, 2, 174, 175, 7, 34, 2, 2, 175, 176, 8, 9, 1, 2, 176, 186,
	3, 2, 2, 2, 177, 178, 5, 18, 10, 2, 178, 179, 8, 9, 1, 2, 179, 186, 3,
	2, 2, 2, 180, 181, 7, 33, 2, 2, 181, 182, 5, 14, 8, 2, 182, 183, 7, 34,
	2, 2, 183, 184, 8, 9, 1, 2, 184, 186, 3, 2, 2, 2, 185, 108, 3, 2, 2, 2,
	185, 117, 3, 2, 2, 2, 185, 128, 3, 2, 2, 2, 185, 136, 3, 2, 2, 2, 185,
	147, 3, 2, 2, 2, 185, 155, 3, 2, 2, 2, 185, 166, 3, 2, 2, 2, 185, 170,
	3, 2, 2, 2, 185, 177, 3, 2, 2, 2, 185, 180, 3, 2, 2, 2, 186, 233, 3, 2,
	2, 2, 187, 188, 12, 19, 2, 2, 188, 189, 9, 3, 2, 2, 189, 190, 5, 16, 9,
	20, 190, 191, 8, 9, 1, 2, 191, 232, 3, 2, 2, 2, 192, 193, 12, 15, 2, 2,
	193, 194, 9, 4, 2, 2, 194, 195, 5, 16, 9, 16, 195, 196, 8, 9, 1, 2, 196,
	232, 3, 2, 2, 2, 197, 198, 12, 11, 2, 2, 198, 199, 9, 5, 2, 2, 199, 200,
	5, 16, 9, 12, 200, 201, 8, 9, 1, 2, 201, 232, 3, 2, 2, 2, 202, 203, 12,
	7, 2, 2, 203, 204, 9, 7, 2, 2, 204, 205, 5, 16, 9, 8, 205, 206, 8, 9, 1,
	2, 206, 232, 3, 2, 2, 2, 207, 208, 12, 17, 2, 2, 208, 209, 9, 3, 2, 2,
	209, 210, 7, 33, 2, 2, 210, 211, 5, 16, 9, 2, 211, 212, 9, 2, 2, 2, 212,
	213, 7, 34, 2, 2, 213, 214, 8, 9, 1, 2, 214, 232, 3, 2, 2, 2, 215, 216,
	12, 13, 2, 2, 216, 217, 9, 4, 2, 2, 217, 218, 7, 33, 2, 2, 218, 219, 5,
	16, 9, 2, 219, 220, 9, 2, 2, 2, 220, 221, 7, 34, 2, 2, 221, 222, 8, 9,
	1, 2, 222, 232, 3, 2, 2, 2, 223, 224, 12, 9, 2, 2, 224, 225, 9, 5, 2, 2,
	225, 226, 7, 33, 2, 2, 226, 227, 5, 16, 9, 2, 227, 228, 9, 2, 2, 2, 228,
	229, 7, 34, 2, 2, 229, 230, 8, 9, 1, 2, 230, 232, 3, 2, 2, 2, 231, 187,
	3, 2, 2, 2, 231, 192, 3, 2, 2, 2, 231, 197, 3, 2, 2, 2, 231, 202, 3, 2,
	2, 2, 231, 207, 3, 2, 2, 2, 231, 215, 3, 2, 2, 2, 231, 223, 3, 2, 2, 2,
	232, 235, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234,
	17, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 236, 237, 7, 12, 2, 2, 237, 247,
	8, 10, 1, 2, 238, 239, 7, 13, 2, 2, 239, 247, 8, 10, 1, 2, 240, 241, 7,
	15, 2, 2, 241, 247, 8, 10, 1, 2, 242, 243, 7, 16, 2, 2, 243, 247, 8, 10,
	1, 2, 244, 245, 7, 17, 2, 2, 245, 247, 8, 10, 1, 2, 246, 236, 3, 2, 2,
	2, 246, 238, 3, 2, 2, 2, 246, 240, 3, 2, 2, 2, 246, 242, 3, 2, 2, 2, 246,
	244, 3, 2, 2, 2, 247, 19, 3, 2, 2, 2, 10, 26, 53, 89, 103, 185, 231, 233,
	246,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'println!'", "'number'", "'as f64'", "'as i64'", "'let'", "'mut'",
	"'i64'", "'f64'", "'String'", "", "", "", "", "", "", "'.'", "';'", "','",
	"':'", "'='", "'>='", "'<='", "'!='", "'>'", "'<'", "'*'", "'/'", "'%'",
	"'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'&&'", "'||'",
	"'!'",
}
var symbolicNames = []string{
	"", "R_PRINTLN", "P_NUMBER", "R_AS_DOUBLE", "R_AS_INTEGER", "R_LET", "R_MUT",
	"R_INT", "R_FLOAT", "R_STRING", "NUMBER", "DOUBLE", "CHAR", "STRING", "BOOLEAN",
	"ID", "TK_PUNTO", "TK_PUNTOCOMA", "TK_COMA", "TK_DOSPUNTOS", "TK_IGUAL",
	"TK_MAYORIGUAL", "TK_MENORIGUAL", "TK_DIFIGUAL", "TK_MAYOR", "TK_MENOR",
	"TK_MULT", "TK_DIV", "TK_MODULO", "TK_MAS", "TK_MENOS", "TK_PARA", "TK_PARC",
	"TK_LLAVEA", "TK_LLAVEC", "TK_CORA", "TK_CORC", "TK_AND", "TK_OR", "TK_NOT",
	"WHITESPACE", "TK_MULTI", "TK_LINE",
}

var ruleNames = []string{
	"start", "instrucciones", "instruccion", "instr_declaracion", "instr_asignacion",
	"instr_tipo", "expression", "exp_arit", "primitivo",
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
	ChemsR_AS_DOUBLE   = 3
	ChemsR_AS_INTEGER  = 4
	ChemsR_LET         = 5
	ChemsR_MUT         = 6
	ChemsR_INT         = 7
	ChemsR_FLOAT       = 8
	ChemsR_STRING      = 9
	ChemsNUMBER        = 10
	ChemsDOUBLE        = 11
	ChemsCHAR          = 12
	ChemsSTRING        = 13
	ChemsBOOLEAN       = 14
	ChemsID            = 15
	ChemsTK_PUNTO      = 16
	ChemsTK_PUNTOCOMA  = 17
	ChemsTK_COMA       = 18
	ChemsTK_DOSPUNTOS  = 19
	ChemsTK_IGUAL      = 20
	ChemsTK_MAYORIGUAL = 21
	ChemsTK_MENORIGUAL = 22
	ChemsTK_DIFIGUAL   = 23
	ChemsTK_MAYOR      = 24
	ChemsTK_MENOR      = 25
	ChemsTK_MULT       = 26
	ChemsTK_DIV        = 27
	ChemsTK_MODULO     = 28
	ChemsTK_MAS        = 29
	ChemsTK_MENOS      = 30
	ChemsTK_PARA       = 31
	ChemsTK_PARC       = 32
	ChemsTK_LLAVEA     = 33
	ChemsTK_LLAVEC     = 34
	ChemsTK_CORA       = 35
	ChemsTK_CORC       = 36
	ChemsTK_AND        = 37
	ChemsTK_OR         = 38
	ChemsTK_NOT        = 39
	ChemsWHITESPACE    = 40
	ChemsTK_MULTI      = 41
	ChemsTK_LINE       = 42
)

// Chems rules.
const (
	ChemsRULE_start             = 0
	ChemsRULE_instrucciones     = 1
	ChemsRULE_instruccion       = 2
	ChemsRULE_instr_declaracion = 3
	ChemsRULE_instr_asignacion  = 4
	ChemsRULE_instr_tipo        = 5
	ChemsRULE_expression        = 6
	ChemsRULE_exp_arit          = 7
	ChemsRULE_primitivo         = 8
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
		p.SetState(18)

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
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_PRINTLN)|(1<<ChemsR_LET)|(1<<ChemsID))) != 0 {
		{
			p.SetState(21)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(26)
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

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_instr_declaracion returns the _instr_declaracion rule contexts.
	Get_instr_declaracion() IInstr_declaracionContext

	// Get_instr_asignacion returns the _instr_asignacion rule contexts.
	Get_instr_asignacion() IInstr_asignacionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_instr_declaracion sets the _instr_declaracion rule contexts.
	Set_instr_declaracion(IInstr_declaracionContext)

	// Set_instr_asignacion sets the _instr_asignacion rule contexts.
	Set_instr_asignacion(IInstr_asignacionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	instr              interfaces.Instruction
	_expression        IExpressionContext
	_STRING            antlr.Token
	_instr_declaracion IInstr_declaracionContext
	_instr_asignacion  IInstr_asignacionContext
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

func (s *InstruccionContext) Get_STRING() antlr.Token { return s._STRING }

func (s *InstruccionContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *InstruccionContext) Get_expression() IExpressionContext { return s._expression }

func (s *InstruccionContext) Get_instr_declaracion() IInstr_declaracionContext {
	return s._instr_declaracion
}

func (s *InstruccionContext) Get_instr_asignacion() IInstr_asignacionContext {
	return s._instr_asignacion
}

func (s *InstruccionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *InstruccionContext) Set_instr_declaracion(v IInstr_declaracionContext) {
	s._instr_declaracion = v
}

func (s *InstruccionContext) Set_instr_asignacion(v IInstr_asignacionContext) { s._instr_asignacion = v }

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

func (s *InstruccionContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChemsSTRING, 0)
}

func (s *InstruccionContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *InstruccionContext) Instr_declaracion() IInstr_declaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_declaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_declaracionContext)
}

func (s *InstruccionContext) Instr_asignacion() IInstr_asignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_asignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_asignacionContext)
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

	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Match(ChemsR_PRINTLN)
		}
		{
			p.SetState(30)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(31)

			var _x = p.Expression()

			localctx.(*InstruccionContext)._expression = _x
		}
		{
			p.SetState(32)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(33)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*InstruccionContext).instr = instruction.PRINTLN(localctx.(*InstruccionContext).Get_expression().GetP(), "-1")

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)
			p.Match(ChemsR_PRINTLN)
		}
		{
			p.SetState(37)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(38)

			var _m = p.Match(ChemsSTRING)

			localctx.(*InstruccionContext)._STRING = _m
		}
		{
			p.SetState(39)
			p.Match(ChemsTK_COMA)
		}
		{
			p.SetState(40)

			var _x = p.Expression()

			localctx.(*InstruccionContext)._expression = _x
		}
		{
			p.SetState(41)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(42)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*InstruccionContext).instr = instruction.PRINTLN(localctx.(*InstruccionContext).Get_expression().GetP(), (func() string {
			if localctx.(*InstruccionContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*InstruccionContext).Get_STRING().GetText()
			}
		}())[1:len((func() string {
			if localctx.(*InstruccionContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*InstruccionContext).Get_STRING().GetText()
			}
		}()))-1])

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(45)

			var _x = p.Instr_declaracion()

			localctx.(*InstruccionContext)._instr_declaracion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_declaracion().GetInstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(48)

			var _x = p.Instr_asignacion()

			localctx.(*InstruccionContext)._instr_asignacion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_asignacion().GetInstr()

	}

	return localctx
}

// IInstr_declaracionContext is an interface to support dynamic dispatch.
type IInstr_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_instr_tipo returns the _instr_tipo rule contexts.
	Get_instr_tipo() IInstr_tipoContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_instr_tipo sets the _instr_tipo rule contexts.
	Set_instr_tipo(IInstr_tipoContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_declaracionContext differentiates from other interfaces.
	IsInstr_declaracionContext()
}

type Instr_declaracionContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruction
	_ID         antlr.Token
	_expression IExpressionContext
	_instr_tipo IInstr_tipoContext
}

func NewEmptyInstr_declaracionContext() *Instr_declaracionContext {
	var p = new(Instr_declaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_declaracion
	return p
}

func (*Instr_declaracionContext) IsInstr_declaracionContext() {}

func NewInstr_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_declaracionContext {
	var p = new(Instr_declaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_declaracion

	return p
}

func (s *Instr_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_declaracionContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_declaracionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_declaracionContext) Get_expression() IExpressionContext { return s._expression }

func (s *Instr_declaracionContext) Get_instr_tipo() IInstr_tipoContext { return s._instr_tipo }

func (s *Instr_declaracionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Instr_declaracionContext) Set_instr_tipo(v IInstr_tipoContext) { s._instr_tipo = v }

func (s *Instr_declaracionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_declaracionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_declaracionContext) R_LET() antlr.TerminalNode {
	return s.GetToken(ChemsR_LET, 0)
}

func (s *Instr_declaracionContext) R_MUT() antlr.TerminalNode {
	return s.GetToken(ChemsR_MUT, 0)
}

func (s *Instr_declaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_declaracionContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUAL, 0)
}

func (s *Instr_declaracionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instr_declaracionContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *Instr_declaracionContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DOSPUNTOS, 0)
}

func (s *Instr_declaracionContext) Instr_tipo() IInstr_tipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_tipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_tipoContext)
}

func (s *Instr_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_declaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_declaracion(s)
	}
}

func (s *Instr_declaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_declaracion(s)
	}
}

func (p *Chems) Instr_declaracion() (localctx IInstr_declaracionContext) {
	localctx = NewInstr_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChemsRULE_instr_declaracion)

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

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Match(ChemsR_LET)
		}
		{
			p.SetState(54)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(55)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(56)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(57)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(58)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), interfaces.NULL, localctx.(*Instr_declaracionContext).Get_expression().GetP(), true, false, false)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Match(ChemsR_LET)
		}
		{
			p.SetState(62)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(63)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(64)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(65)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(66)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(67)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(68)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_declaracionContext).Get_instr_tipo().GetTipo_exp(), localctx.(*Instr_declaracionContext).Get_expression().GetP(), true, false, false)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(71)
			p.Match(ChemsR_LET)
		}
		{
			p.SetState(72)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(73)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(74)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(75)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), interfaces.NULL, localctx.(*Instr_declaracionContext).Get_expression().GetP(), false, false, false)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(78)
			p.Match(ChemsR_LET)
		}
		{
			p.SetState(79)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(80)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(81)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(82)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(83)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(84)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_declaracionContext).Get_instr_tipo().GetTipo_exp(), localctx.(*Instr_declaracionContext).Get_expression().GetP(), false, false, false)

	}

	return localctx
}

// IInstr_asignacionContext is an interface to support dynamic dispatch.
type IInstr_asignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_asignacionContext differentiates from other interfaces.
	IsInstr_asignacionContext()
}

type Instr_asignacionContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruction
	_ID         antlr.Token
	_expression IExpressionContext
}

func NewEmptyInstr_asignacionContext() *Instr_asignacionContext {
	var p = new(Instr_asignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_asignacion
	return p
}

func (*Instr_asignacionContext) IsInstr_asignacionContext() {}

func NewInstr_asignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_asignacionContext {
	var p = new(Instr_asignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_asignacion

	return p
}

func (s *Instr_asignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_asignacionContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_asignacionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_asignacionContext) Get_expression() IExpressionContext { return s._expression }

func (s *Instr_asignacionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Instr_asignacionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_asignacionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_asignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_asignacionContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUAL, 0)
}

func (s *Instr_asignacionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instr_asignacionContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *Instr_asignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_asignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_asignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_asignacion(s)
	}
}

func (s *Instr_asignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_asignacion(s)
	}
}

func (p *Chems) Instr_asignacion() (localctx IInstr_asignacionContext) {
	localctx = NewInstr_asignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ChemsRULE_instr_asignacion)

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
		p.SetState(89)

		var _m = p.Match(ChemsID)

		localctx.(*Instr_asignacionContext)._ID = _m
	}
	{
		p.SetState(90)
		p.Match(ChemsTK_IGUAL)
	}
	{
		p.SetState(91)

		var _x = p.Expression()

		localctx.(*Instr_asignacionContext)._expression = _x
	}
	{
		p.SetState(92)
		p.Match(ChemsTK_PUNTOCOMA)
	}
	localctx.(*Instr_asignacionContext).instr = variable.NewAssignment((func() string {
		if localctx.(*Instr_asignacionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Instr_asignacionContext).Get_ID().GetText()
		}
	}()), localctx.(*Instr_asignacionContext).Get_expression().GetP())

	return localctx
}

// IInstr_tipoContext is an interface to support dynamic dispatch.
type IInstr_tipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipo_exp returns the tipo_exp attribute.
	GetTipo_exp() interfaces.TipoExpresion

	// SetTipo_exp sets the tipo_exp attribute.
	SetTipo_exp(interfaces.TipoExpresion)

	// IsInstr_tipoContext differentiates from other interfaces.
	IsInstr_tipoContext()
}

type Instr_tipoContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	tipo_exp interfaces.TipoExpresion
}

func NewEmptyInstr_tipoContext() *Instr_tipoContext {
	var p = new(Instr_tipoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_tipo
	return p
}

func (*Instr_tipoContext) IsInstr_tipoContext() {}

func NewInstr_tipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_tipoContext {
	var p = new(Instr_tipoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_tipo

	return p
}

func (s *Instr_tipoContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_tipoContext) GetTipo_exp() interfaces.TipoExpresion { return s.tipo_exp }

func (s *Instr_tipoContext) SetTipo_exp(v interfaces.TipoExpresion) { s.tipo_exp = v }

func (s *Instr_tipoContext) R_INT() antlr.TerminalNode {
	return s.GetToken(ChemsR_INT, 0)
}

func (s *Instr_tipoContext) R_FLOAT() antlr.TerminalNode {
	return s.GetToken(ChemsR_FLOAT, 0)
}

func (s *Instr_tipoContext) R_STRING() antlr.TerminalNode {
	return s.GetToken(ChemsR_STRING, 0)
}

func (s *Instr_tipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_tipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_tipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_tipo(s)
	}
}

func (s *Instr_tipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_tipo(s)
	}
}

func (p *Chems) Instr_tipo() (localctx IInstr_tipoContext) {
	localctx = NewInstr_tipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ChemsRULE_instr_tipo)

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

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Match(ChemsR_INT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.INTEGER

	case ChemsR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.Match(ChemsR_FLOAT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.FLOAT

	case ChemsR_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(99)
			p.Match(ChemsR_STRING)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STRING

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

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
	p.EnterRule(localctx, 12, ChemsRULE_expression)

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
		p.SetState(103)

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

	// GetTipo_left returns the tipo_left token.
	GetTipo_left() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetTipo_right returns the tipo_right token.
	GetTipo_right() antlr.Token

	// SetTipo_left sets the tipo_left token.
	SetTipo_left(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetTipo_right sets the tipo_right token.
	SetTipo_right(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExp_aritContext

	// GetRight returns the right rule contexts.
	GetRight() IExp_aritContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExp_aritContext)

	// SetRight sets the right rule contexts.
	SetRight(IExp_aritContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

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
	left        IExp_aritContext
	tipo_left   antlr.Token
	op          antlr.Token
	right       IExp_aritContext
	tipo_right  antlr.Token
	_expression IExpressionContext
	_primitivo  IPrimitivoContext
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

func (s *Exp_aritContext) GetTipo_left() antlr.Token { return s.tipo_left }

func (s *Exp_aritContext) GetOp() antlr.Token { return s.op }

func (s *Exp_aritContext) GetTipo_right() antlr.Token { return s.tipo_right }

func (s *Exp_aritContext) SetTipo_left(v antlr.Token) { s.tipo_left = v }

func (s *Exp_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Exp_aritContext) SetTipo_right(v antlr.Token) { s.tipo_right = v }

func (s *Exp_aritContext) GetLeft() IExp_aritContext { return s.left }

func (s *Exp_aritContext) GetRight() IExp_aritContext { return s.right }

func (s *Exp_aritContext) Get_expression() IExpressionContext { return s._expression }

func (s *Exp_aritContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Exp_aritContext) SetLeft(v IExp_aritContext) { s.left = v }

func (s *Exp_aritContext) SetRight(v IExp_aritContext) { s.right = v }

func (s *Exp_aritContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Exp_aritContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Exp_aritContext) GetP() interfaces.Expresion { return s.p }

func (s *Exp_aritContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *Exp_aritContext) AllTK_PARA() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_PARA)
}

func (s *Exp_aritContext) TK_PARA(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, i)
}

func (s *Exp_aritContext) AllTK_PARC() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_PARC)
}

func (s *Exp_aritContext) TK_PARC(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, i)
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

func (s *Exp_aritContext) AllR_AS_DOUBLE() []antlr.TerminalNode {
	return s.GetTokens(ChemsR_AS_DOUBLE)
}

func (s *Exp_aritContext) R_AS_DOUBLE(i int) antlr.TerminalNode {
	return s.GetToken(ChemsR_AS_DOUBLE, i)
}

func (s *Exp_aritContext) AllR_AS_INTEGER() []antlr.TerminalNode {
	return s.GetTokens(ChemsR_AS_INTEGER)
}

func (s *Exp_aritContext) R_AS_INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(ChemsR_AS_INTEGER, i)
}

func (s *Exp_aritContext) TK_MULT() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MULT, 0)
}

func (s *Exp_aritContext) TK_DIV() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DIV, 0)
}

func (s *Exp_aritContext) TK_MODULO() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MODULO, 0)
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

func (s *Exp_aritContext) TK_DIFIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DIFIGUAL, 0)
}

func (s *Exp_aritContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Exp_aritContext) TK_NOT() antlr.TerminalNode {
	return s.GetToken(ChemsTK_NOT, 0)
}

func (s *Exp_aritContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Exp_aritContext) TK_AND() antlr.TerminalNode {
	return s.GetToken(ChemsTK_AND, 0)
}

func (s *Exp_aritContext) TK_OR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_OR, 0)
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
	_startState := 14
	p.EnterRecursionRule(localctx, 14, ChemsRULE_exp_arit, _p)
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
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(107)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(108)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(109)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).tipo_left = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsR_AS_DOUBLE || _la == ChemsR_AS_INTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).tipo_left = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(110)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(111)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsTK_MULT)|(1<<ChemsTK_DIV)|(1<<ChemsTK_MODULO))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(112)

			var _x = p.exp_arit(16)

			localctx.(*Exp_aritContext).right = _x
		}
		localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetText()
			}
		}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, (func() string {
			if localctx.(*Exp_aritContext).GetTipo_left() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetTipo_left().GetText()
			}
		}()), "-1")

	case 2:
		{
			p.SetState(115)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(116)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(117)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).tipo_left = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsR_AS_DOUBLE || _la == ChemsR_AS_INTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).tipo_left = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(118)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(119)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsTK_MULT)|(1<<ChemsTK_DIV)|(1<<ChemsTK_MODULO))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(120)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(121)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(122)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).tipo_right = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsR_AS_DOUBLE || _la == ChemsR_AS_INTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).tipo_right = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(123)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetText()
			}
		}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, (func() string {
			if localctx.(*Exp_aritContext).GetTipo_left() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetTipo_left().GetText()
			}
		}()), (func() string {
			if localctx.(*Exp_aritContext).GetTipo_right() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetTipo_right().GetText()
			}
		}()))

	case 3:
		{
			p.SetState(126)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(127)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(128)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).tipo_left = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsR_AS_DOUBLE || _la == ChemsR_AS_INTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).tipo_left = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(129)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(130)

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
			p.SetState(131)

			var _x = p.exp_arit(12)

			localctx.(*Exp_aritContext).right = _x
		}
		localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetText()
			}
		}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, (func() string {
			if localctx.(*Exp_aritContext).GetTipo_left() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetTipo_left().GetText()
			}
		}()), "-1")

	case 4:
		{
			p.SetState(134)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(135)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(136)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).tipo_left = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsR_AS_DOUBLE || _la == ChemsR_AS_INTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).tipo_left = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(137)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(138)

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
			p.SetState(139)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(140)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(141)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).tipo_right = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsR_AS_DOUBLE || _la == ChemsR_AS_INTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).tipo_right = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(142)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetText()
			}
		}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, (func() string {
			if localctx.(*Exp_aritContext).GetTipo_left() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetTipo_left().GetText()
			}
		}()), (func() string {
			if localctx.(*Exp_aritContext).GetTipo_right() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetTipo_right().GetText()
			}
		}()))

	case 5:
		{
			p.SetState(145)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(146)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(147)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).tipo_left = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsR_AS_DOUBLE || _la == ChemsR_AS_INTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).tipo_left = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(148)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(149)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsTK_MAYORIGUAL)|(1<<ChemsTK_MENORIGUAL)|(1<<ChemsTK_DIFIGUAL)|(1<<ChemsTK_MAYOR)|(1<<ChemsTK_MENOR))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(150)

			var _x = p.exp_arit(8)

			localctx.(*Exp_aritContext).right = _x
		}
		localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetText()
			}
		}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, (func() string {
			if localctx.(*Exp_aritContext).GetTipo_left() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetTipo_left().GetText()
			}
		}()), "-1")

	case 6:
		{
			p.SetState(153)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(154)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(155)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).tipo_left = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsR_AS_DOUBLE || _la == ChemsR_AS_INTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).tipo_left = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(156)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(157)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsTK_MAYORIGUAL)|(1<<ChemsTK_MENORIGUAL)|(1<<ChemsTK_DIFIGUAL)|(1<<ChemsTK_MAYOR)|(1<<ChemsTK_MENOR))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(158)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(159)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(160)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).tipo_right = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsR_AS_DOUBLE || _la == ChemsR_AS_INTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).tipo_right = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(161)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetText()
			}
		}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, (func() string {
			if localctx.(*Exp_aritContext).GetTipo_left() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetTipo_left().GetText()
			}
		}()), (func() string {
			if localctx.(*Exp_aritContext).GetTipo_right() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetTipo_right().GetText()
			}
		}()))

	case 7:
		{
			p.SetState(164)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsTK_MENOS || _la == ChemsTK_NOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(165)

			var _x = p.Expression()

			localctx.(*Exp_aritContext)._expression = _x
		}
		localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).Get_expression().GetP(), (func() string {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetText()
			}
		}()), nil, true, "-1", "-1")

	case 8:
		{
			p.SetState(168)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsTK_MENOS || _la == ChemsTK_NOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(169)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(170)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(171)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).tipo_left = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsR_AS_DOUBLE || _la == ChemsR_AS_INTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).tipo_left = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(172)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetText()
			}
		}()), nil, true, (func() string {
			if localctx.(*Exp_aritContext).GetTipo_left() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetTipo_left().GetText()
			}
		}()), "-1")

	case 9:
		{
			p.SetState(175)

			var _x = p.Primitivo()

			localctx.(*Exp_aritContext)._primitivo = _x
		}
		localctx.(*Exp_aritContext).p = localctx.(*Exp_aritContext).Get_primitivo().GetP()

	case 10:
		{
			p.SetState(178)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(179)

			var _x = p.Expression()

			localctx.(*Exp_aritContext)._expression = _x
		}
		{
			p.SetState(180)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Exp_aritContext).p = localctx.(*Exp_aritContext).Get_expression().GetP()

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(229)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(185)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(186)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsTK_MULT)|(1<<ChemsTK_DIV)|(1<<ChemsTK_MODULO))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(187)

					var _x = p.exp_arit(18)

					localctx.(*Exp_aritContext).right = _x
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, "-1", "-1")

			case 2:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(190)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(191)

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
					p.SetState(192)

					var _x = p.exp_arit(14)

					localctx.(*Exp_aritContext).right = _x
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, "-1", "-1")

			case 3:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(196)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsTK_MAYORIGUAL)|(1<<ChemsTK_MENORIGUAL)|(1<<ChemsTK_DIFIGUAL)|(1<<ChemsTK_MAYOR)|(1<<ChemsTK_MENOR))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(197)

					var _x = p.exp_arit(10)

					localctx.(*Exp_aritContext).right = _x
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, "-1", "-1")

			case 4:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(200)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(201)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ChemsTK_AND || _la == ChemsTK_OR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(202)

					var _x = p.exp_arit(6)

					localctx.(*Exp_aritContext).right = _x
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, "-1", "-1")

			case 5:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(205)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(206)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsTK_MULT)|(1<<ChemsTK_DIV)|(1<<ChemsTK_MODULO))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(207)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(208)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(209)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).tipo_right = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ChemsR_AS_DOUBLE || _la == ChemsR_AS_INTEGER) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).tipo_right = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(210)
					p.Match(ChemsTK_PARC)
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, "-1", (func() string {
					if localctx.(*Exp_aritContext).GetTipo_right() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetTipo_right().GetText()
					}
				}()))

			case 6:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(213)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(214)

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
					p.SetState(215)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(216)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(217)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).tipo_right = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ChemsR_AS_DOUBLE || _la == ChemsR_AS_INTEGER) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).tipo_right = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(218)
					p.Match(ChemsTK_PARC)
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, "-1", (func() string {
					if localctx.(*Exp_aritContext).GetTipo_right() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetTipo_right().GetText()
					}
				}()))

			case 7:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(221)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(222)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsTK_MAYORIGUAL)|(1<<ChemsTK_MENORIGUAL)|(1<<ChemsTK_DIFIGUAL)|(1<<ChemsTK_MAYOR)|(1<<ChemsTK_MENOR))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(223)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(224)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(225)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).tipo_right = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ChemsR_AS_DOUBLE || _la == ChemsR_AS_INTEGER) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).tipo_right = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(226)
					p.Match(ChemsTK_PARC)
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, "-1", (func() string {
					if localctx.(*Exp_aritContext).GetTipo_right() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetTipo_right().GetText()
					}
				}()))

			}

		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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

	// Get_DOUBLE returns the _DOUBLE token.
	Get_DOUBLE() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_BOOLEAN returns the _BOOLEAN token.
	Get_BOOLEAN() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_DOUBLE sets the _DOUBLE token.
	Set_DOUBLE(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_BOOLEAN sets the _BOOLEAN token.
	Set_BOOLEAN(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	p        interfaces.Expresion
	_NUMBER  antlr.Token
	_DOUBLE  antlr.Token
	_STRING  antlr.Token
	_BOOLEAN antlr.Token
	_ID      antlr.Token
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

func (s *PrimitivoContext) Get_DOUBLE() antlr.Token { return s._DOUBLE }

func (s *PrimitivoContext) Get_STRING() antlr.Token { return s._STRING }

func (s *PrimitivoContext) Get_BOOLEAN() antlr.Token { return s._BOOLEAN }

func (s *PrimitivoContext) Get_ID() antlr.Token { return s._ID }

func (s *PrimitivoContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *PrimitivoContext) Set_DOUBLE(v antlr.Token) { s._DOUBLE = v }

func (s *PrimitivoContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *PrimitivoContext) Set_BOOLEAN(v antlr.Token) { s._BOOLEAN = v }

func (s *PrimitivoContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *PrimitivoContext) GetP() interfaces.Expresion { return s.p }

func (s *PrimitivoContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *PrimitivoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChemsNUMBER, 0)
}

func (s *PrimitivoContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ChemsDOUBLE, 0)
}

func (s *PrimitivoContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChemsSTRING, 0)
}

func (s *PrimitivoContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ChemsBOOLEAN, 0)
}

func (s *PrimitivoContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
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
	p.EnterRule(localctx, 16, ChemsRULE_primitivo)

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

	p.SetState(244)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)

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

	case ChemsDOUBLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)

			var _m = p.Match(ChemsDOUBLE)

			localctx.(*PrimitivoContext)._DOUBLE = _m
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = expresion.PRIMITIVO(num, interfaces.FLOAT)

	case ChemsSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(238)

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

	case ChemsBOOLEAN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(240)

			var _m = p.Match(ChemsBOOLEAN)

			localctx.(*PrimitivoContext)._BOOLEAN = _m
		}

		// str:= (func() string { if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText() }}())[1:len((func() string { if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText() }}()))-1]
		localctx.(*PrimitivoContext).p = expresion.PRIMITIVO((func() string {
			if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText()
			}
		}()), interfaces.BOOLEAN)

	case ChemsID:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(242)

			var _m = p.Match(ChemsID)

			localctx.(*PrimitivoContext)._ID = _m
		}
		localctx.(*PrimitivoContext).p = variable.NewIdentifier((func() string {
			if localctx.(*PrimitivoContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_ID().GetText()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
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
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

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
import "OLC2/Interprete/instruction/control"
import "OLC2/Interprete/instruction/loops"
import "OLC2/Interprete/instruction/transferencia"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 62, 628,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3, 2, 3,
	2, 3, 2, 3, 3, 6, 3, 81, 10, 3, 13, 3, 14, 3, 82, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 5, 4, 90, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 5, 5, 126, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 153, 10, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 189, 10, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5,
	10, 224, 10, 10, 3, 11, 7, 11, 227, 10, 11, 12, 11, 14, 11, 230, 11, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 261, 10,
	12, 3, 13, 7, 13, 264, 10, 13, 12, 13, 14, 13, 267, 11, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 286, 10, 14, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5,
	15, 301, 10, 15, 3, 16, 3, 16, 3, 16, 3, 17, 6, 17, 307, 10, 17, 13, 17,
	14, 17, 308, 3, 17, 3, 17, 3, 18, 6, 18, 314, 10, 18, 13, 18, 14, 18, 315,
	3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 327,
	10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 342, 10, 20, 3, 21, 6, 21, 345, 10,
	21, 13, 21, 14, 21, 346, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5,
	22, 366, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 6, 24,
	375, 10, 24, 13, 24, 14, 24, 376, 3, 24, 3, 24, 3, 25, 6, 25, 382, 10,
	25, 13, 25, 14, 25, 383, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 5, 26, 395, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 442, 10,
	31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32,
	453, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	5, 35, 474, 10, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37,
	556, 10, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 7, 37, 602, 10, 37, 12, 37, 14, 37, 605, 11,
	37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 626,
	10, 38, 3, 38, 2, 3, 72, 39, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	62, 64, 66, 68, 70, 72, 74, 2, 8, 3, 2, 5, 6, 3, 2, 45, 47, 3, 2, 48, 49,
	4, 2, 38, 39, 41, 44, 4, 2, 49, 49, 59, 59, 3, 2, 56, 57, 2, 652, 2, 76,
	3, 2, 2, 2, 4, 80, 3, 2, 2, 2, 6, 89, 3, 2, 2, 2, 8, 125, 3, 2, 2, 2, 10,
	127, 3, 2, 2, 2, 12, 152, 3, 2, 2, 2, 14, 188, 3, 2, 2, 2, 16, 190, 3,
	2, 2, 2, 18, 223, 3, 2, 2, 2, 20, 228, 3, 2, 2, 2, 22, 260, 3, 2, 2, 2,
	24, 265, 3, 2, 2, 2, 26, 285, 3, 2, 2, 2, 28, 300, 3, 2, 2, 2, 30, 302,
	3, 2, 2, 2, 32, 306, 3, 2, 2, 2, 34, 313, 3, 2, 2, 2, 36, 326, 3, 2, 2,
	2, 38, 341, 3, 2, 2, 2, 40, 344, 3, 2, 2, 2, 42, 365, 3, 2, 2, 2, 44, 367,
	3, 2, 2, 2, 46, 374, 3, 2, 2, 2, 48, 381, 3, 2, 2, 2, 50, 394, 3, 2, 2,
	2, 52, 396, 3, 2, 2, 2, 54, 402, 3, 2, 2, 2, 56, 409, 3, 2, 2, 2, 58, 415,
	3, 2, 2, 2, 60, 441, 3, 2, 2, 2, 62, 452, 3, 2, 2, 2, 64, 454, 3, 2, 2,
	2, 66, 458, 3, 2, 2, 2, 68, 473, 3, 2, 2, 2, 70, 475, 3, 2, 2, 2, 72, 555,
	3, 2, 2, 2, 74, 625, 3, 2, 2, 2, 76, 77, 5, 4, 3, 2, 77, 78, 8, 2, 1, 2,
	78, 3, 3, 2, 2, 2, 79, 81, 5, 8, 5, 2, 80, 79, 3, 2, 2, 2, 81, 82, 3, 2,
	2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85,
	8, 3, 1, 2, 85, 5, 3, 2, 2, 2, 86, 87, 7, 34, 2, 2, 87, 90, 8, 4, 1, 2,
	88, 90, 8, 4, 1, 2, 89, 86, 3, 2, 2, 2, 89, 88, 3, 2, 2, 2, 90, 7, 3, 2,
	2, 2, 91, 92, 5, 12, 7, 2, 92, 93, 5, 6, 4, 2, 93, 94, 8, 5, 1, 2, 94,
	126, 3, 2, 2, 2, 95, 96, 5, 14, 8, 2, 96, 97, 8, 5, 1, 2, 97, 126, 3, 2,
	2, 2, 98, 99, 5, 16, 9, 2, 99, 100, 8, 5, 1, 2, 100, 126, 3, 2, 2, 2, 101,
	102, 5, 18, 10, 2, 102, 103, 8, 5, 1, 2, 103, 126, 3, 2, 2, 2, 104, 105,
	5, 26, 14, 2, 105, 106, 8, 5, 1, 2, 106, 126, 3, 2, 2, 2, 107, 108, 5,
	54, 28, 2, 108, 109, 8, 5, 1, 2, 109, 126, 3, 2, 2, 2, 110, 111, 5, 62,
	32, 2, 111, 112, 8, 5, 1, 2, 112, 126, 3, 2, 2, 2, 113, 114, 5, 64, 33,
	2, 114, 115, 8, 5, 1, 2, 115, 126, 3, 2, 2, 2, 116, 117, 5, 56, 29, 2,
	117, 118, 8, 5, 1, 2, 118, 126, 3, 2, 2, 2, 119, 120, 5, 60, 31, 2, 120,
	121, 8, 5, 1, 2, 121, 126, 3, 2, 2, 2, 122, 123, 5, 10, 6, 2, 123, 124,
	8, 5, 1, 2, 124, 126, 3, 2, 2, 2, 125, 91, 3, 2, 2, 2, 125, 95, 3, 2, 2,
	2, 125, 98, 3, 2, 2, 2, 125, 101, 3, 2, 2, 2, 125, 104, 3, 2, 2, 2, 125,
	107, 3, 2, 2, 2, 125, 110, 3, 2, 2, 2, 125, 113, 3, 2, 2, 2, 125, 116,
	3, 2, 2, 2, 125, 119, 3, 2, 2, 2, 125, 122, 3, 2, 2, 2, 126, 9, 3, 2, 2,
	2, 127, 128, 7, 20, 2, 2, 128, 129, 7, 19, 2, 2, 129, 130, 7, 50, 2, 2,
	130, 131, 7, 51, 2, 2, 131, 132, 7, 52, 2, 2, 132, 133, 5, 4, 3, 2, 133,
	134, 7, 53, 2, 2, 134, 135, 8, 6, 1, 2, 135, 11, 3, 2, 2, 2, 136, 137,
	7, 3, 2, 2, 137, 138, 7, 50, 2, 2, 138, 139, 5, 70, 36, 2, 139, 140, 7,
	51, 2, 2, 140, 141, 7, 34, 2, 2, 141, 142, 8, 7, 1, 2, 142, 153, 3, 2,
	2, 2, 143, 144, 7, 3, 2, 2, 144, 145, 7, 50, 2, 2, 145, 146, 7, 29, 2,
	2, 146, 147, 7, 35, 2, 2, 147, 148, 5, 70, 36, 2, 148, 149, 7, 51, 2, 2,
	149, 150, 7, 34, 2, 2, 150, 151, 8, 7, 1, 2, 151, 153, 3, 2, 2, 2, 152,
	136, 3, 2, 2, 2, 152, 143, 3, 2, 2, 2, 153, 13, 3, 2, 2, 2, 154, 155, 7,
	7, 2, 2, 155, 156, 7, 8, 2, 2, 156, 157, 7, 31, 2, 2, 157, 158, 7, 37,
	2, 2, 158, 159, 5, 70, 36, 2, 159, 160, 7, 34, 2, 2, 160, 161, 8, 8, 1,
	2, 161, 189, 3, 2, 2, 2, 162, 163, 7, 7, 2, 2, 163, 164, 7, 8, 2, 2, 164,
	165, 7, 31, 2, 2, 165, 166, 7, 36, 2, 2, 166, 167, 5, 68, 35, 2, 167, 168,
	7, 37, 2, 2, 168, 169, 5, 70, 36, 2, 169, 170, 7, 34, 2, 2, 170, 171, 8,
	8, 1, 2, 171, 189, 3, 2, 2, 2, 172, 173, 7, 7, 2, 2, 173, 174, 7, 31, 2,
	2, 174, 175, 7, 37, 2, 2, 175, 176, 5, 70, 36, 2, 176, 177, 7, 34, 2, 2,
	177, 178, 8, 8, 1, 2, 178, 189, 3, 2, 2, 2, 179, 180, 7, 7, 2, 2, 180,
	181, 7, 31, 2, 2, 181, 182, 7, 36, 2, 2, 182, 183, 5, 68, 35, 2, 183, 184,
	7, 37, 2, 2, 184, 185, 5, 70, 36, 2, 185, 186, 7, 34, 2, 2, 186, 187, 8,
	8, 1, 2, 187, 189, 3, 2, 2, 2, 188, 154, 3, 2, 2, 2, 188, 162, 3, 2, 2,
	2, 188, 172, 3, 2, 2, 2, 188, 179, 3, 2, 2, 2, 189, 15, 3, 2, 2, 2, 190,
	191, 7, 31, 2, 2, 191, 192, 7, 37, 2, 2, 192, 193, 5, 70, 36, 2, 193, 194,
	7, 34, 2, 2, 194, 195, 8, 9, 1, 2, 195, 17, 3, 2, 2, 2, 196, 197, 7, 9,
	2, 2, 197, 198, 5, 70, 36, 2, 198, 199, 7, 52, 2, 2, 199, 200, 5, 4, 3,
	2, 200, 201, 7, 53, 2, 2, 201, 202, 8, 10, 1, 2, 202, 224, 3, 2, 2, 2,
	203, 204, 7, 9, 2, 2, 204, 205, 5, 70, 36, 2, 205, 206, 7, 52, 2, 2, 206,
	207, 5, 4, 3, 2, 207, 208, 7, 53, 2, 2, 208, 209, 7, 10, 2, 2, 209, 210,
	7, 52, 2, 2, 210, 211, 5, 4, 3, 2, 211, 212, 7, 53, 2, 2, 212, 213, 8,
	10, 1, 2, 213, 224, 3, 2, 2, 2, 214, 215, 7, 9, 2, 2, 215, 216, 5, 70,
	36, 2, 216, 217, 7, 52, 2, 2, 217, 218, 5, 4, 3, 2, 218, 219, 7, 53, 2,
	2, 219, 220, 7, 10, 2, 2, 220, 221, 5, 20, 11, 2, 221, 222, 8, 10, 1, 2,
	222, 224, 3, 2, 2, 2, 223, 196, 3, 2, 2, 2, 223, 203, 3, 2, 2, 2, 223,
	214, 3, 2, 2, 2, 224, 19, 3, 2, 2, 2, 225, 227, 5, 18, 10, 2, 226, 225,
	3, 2, 2, 2, 227, 230, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 228, 229, 3, 2,
	2, 2, 229, 231, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 231, 232, 8, 11, 1, 2,
	232, 21, 3, 2, 2, 2, 233, 234, 7, 9, 2, 2, 234, 235, 5, 70, 36, 2, 235,
	236, 7, 52, 2, 2, 236, 237, 5, 70, 36, 2, 237, 238, 7, 53, 2, 2, 238, 239,
	8, 12, 1, 2, 239, 261, 3, 2, 2, 2, 240, 241, 7, 9, 2, 2, 241, 242, 5, 70,
	36, 2, 242, 243, 7, 52, 2, 2, 243, 244, 5, 70, 36, 2, 244, 245, 7, 53,
	2, 2, 245, 246, 7, 10, 2, 2, 246, 247, 7, 52, 2, 2, 247, 248, 5, 70, 36,
	2, 248, 249, 7, 53, 2, 2, 249, 250, 8, 12, 1, 2, 250, 261, 3, 2, 2, 2,
	251, 252, 7, 9, 2, 2, 252, 253, 5, 70, 36, 2, 253, 254, 7, 52, 2, 2, 254,
	255, 5, 70, 36, 2, 255, 256, 7, 53, 2, 2, 256, 257, 7, 10, 2, 2, 257, 258,
	5, 24, 13, 2, 258, 259, 8, 12, 1, 2, 259, 261, 3, 2, 2, 2, 260, 233, 3,
	2, 2, 2, 260, 240, 3, 2, 2, 2, 260, 251, 3, 2, 2, 2, 261, 23, 3, 2, 2,
	2, 262, 264, 5, 22, 12, 2, 263, 262, 3, 2, 2, 2, 264, 267, 3, 2, 2, 2,
	265, 263, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 268, 3, 2, 2, 2, 267,
	265, 3, 2, 2, 2, 268, 269, 8, 13, 1, 2, 269, 25, 3, 2, 2, 2, 270, 271,
	7, 11, 2, 2, 271, 272, 5, 70, 36, 2, 272, 273, 7, 52, 2, 2, 273, 274, 5,
	32, 17, 2, 274, 275, 5, 40, 21, 2, 275, 276, 7, 53, 2, 2, 276, 277, 8,
	14, 1, 2, 277, 286, 3, 2, 2, 2, 278, 279, 7, 11, 2, 2, 279, 280, 5, 70,
	36, 2, 280, 281, 7, 52, 2, 2, 281, 282, 5, 40, 21, 2, 282, 283, 7, 53,
	2, 2, 283, 284, 8, 14, 1, 2, 284, 286, 3, 2, 2, 2, 285, 270, 3, 2, 2, 2,
	285, 278, 3, 2, 2, 2, 286, 27, 3, 2, 2, 2, 287, 288, 5, 34, 18, 2, 288,
	289, 7, 40, 2, 2, 289, 290, 7, 52, 2, 2, 290, 291, 5, 4, 3, 2, 291, 292,
	7, 53, 2, 2, 292, 293, 8, 15, 1, 2, 293, 301, 3, 2, 2, 2, 294, 295, 5,
	34, 18, 2, 295, 296, 7, 40, 2, 2, 296, 297, 5, 30, 16, 2, 297, 298, 7,
	35, 2, 2, 298, 299, 8, 15, 1, 2, 299, 301, 3, 2, 2, 2, 300, 287, 3, 2,
	2, 2, 300, 294, 3, 2, 2, 2, 301, 29, 3, 2, 2, 2, 302, 303, 5, 8, 5, 2,
	303, 304, 8, 16, 1, 2, 304, 31, 3, 2, 2, 2, 305, 307, 5, 28, 15, 2, 306,
	305, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 308, 309,
	3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 311, 8, 17, 1, 2, 311, 33, 3, 2,
	2, 2, 312, 314, 5, 36, 19, 2, 313, 312, 3, 2, 2, 2, 314, 315, 3, 2, 2,
	2, 315, 313, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317,
	318, 8, 18, 1, 2, 318, 35, 3, 2, 2, 2, 319, 320, 5, 70, 36, 2, 320, 321,
	7, 58, 2, 2, 321, 322, 8, 19, 1, 2, 322, 327, 3, 2, 2, 2, 323, 324, 5,
	70, 36, 2, 324, 325, 8, 19, 1, 2, 325, 327, 3, 2, 2, 2, 326, 319, 3, 2,
	2, 2, 326, 323, 3, 2, 2, 2, 327, 37, 3, 2, 2, 2, 328, 329, 7, 31, 2, 2,
	329, 330, 7, 40, 2, 2, 330, 331, 7, 52, 2, 2, 331, 332, 5, 4, 3, 2, 332,
	333, 7, 53, 2, 2, 333, 334, 8, 20, 1, 2, 334, 342, 3, 2, 2, 2, 335, 336,
	7, 31, 2, 2, 336, 337, 7, 40, 2, 2, 337, 338, 5, 30, 16, 2, 338, 339, 7,
	35, 2, 2, 339, 340, 8, 20, 1, 2, 340, 342, 3, 2, 2, 2, 341, 328, 3, 2,
	2, 2, 341, 335, 3, 2, 2, 2, 342, 39, 3, 2, 2, 2, 343, 345, 5, 38, 20, 2,
	344, 343, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 346,
	347, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 349, 8, 21, 1, 2, 349, 41,
	3, 2, 2, 2, 350, 351, 7, 11, 2, 2, 351, 352, 5, 70, 36, 2, 352, 353, 7,
	52, 2, 2, 353, 354, 5, 46, 24, 2, 354, 355, 5, 52, 27, 2, 355, 356, 7,
	53, 2, 2, 356, 357, 8, 22, 1, 2, 357, 366, 3, 2, 2, 2, 358, 359, 7, 11,
	2, 2, 359, 360, 5, 70, 36, 2, 360, 361, 7, 52, 2, 2, 361, 362, 5, 52, 27,
	2, 362, 363, 7, 53, 2, 2, 363, 364, 8, 22, 1, 2, 364, 366, 3, 2, 2, 2,
	365, 350, 3, 2, 2, 2, 365, 358, 3, 2, 2, 2, 366, 43, 3, 2, 2, 2, 367, 368,
	5, 48, 25, 2, 368, 369, 7, 40, 2, 2, 369, 370, 5, 70, 36, 2, 370, 371,
	7, 35, 2, 2, 371, 372, 8, 23, 1, 2, 372, 45, 3, 2, 2, 2, 373, 375, 5, 44,
	23, 2, 374, 373, 3, 2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 374, 3, 2, 2, 2,
	376, 377, 3, 2, 2, 2, 377, 378, 3, 2, 2, 2, 378, 379, 8, 24, 1, 2, 379,
	47, 3, 2, 2, 2, 380, 382, 5, 50, 26, 2, 381, 380, 3, 2, 2, 2, 382, 383,
	3, 2, 2, 2, 383, 381, 3, 2, 2, 2, 383, 384, 3, 2, 2, 2, 384, 385, 3, 2,
	2, 2, 385, 386, 8, 25, 1, 2, 386, 49, 3, 2, 2, 2, 387, 388, 5, 70, 36,
	2, 388, 389, 7, 58, 2, 2, 389, 390, 8, 26, 1, 2, 390, 395, 3, 2, 2, 2,
	391, 392, 5, 70, 36, 2, 392, 393, 8, 26, 1, 2, 393, 395, 3, 2, 2, 2, 394,
	387, 3, 2, 2, 2, 394, 391, 3, 2, 2, 2, 395, 51, 3, 2, 2, 2, 396, 397, 7,
	31, 2, 2, 397, 398, 7, 40, 2, 2, 398, 399, 5, 70, 36, 2, 399, 400, 7, 35,
	2, 2, 400, 401, 8, 27, 1, 2, 401, 53, 3, 2, 2, 2, 402, 403, 7, 12, 2, 2,
	403, 404, 5, 70, 36, 2, 404, 405, 7, 52, 2, 2, 405, 406, 5, 4, 3, 2, 406,
	407, 7, 53, 2, 2, 407, 408, 8, 28, 1, 2, 408, 55, 3, 2, 2, 2, 409, 410,
	7, 16, 2, 2, 410, 411, 7, 52, 2, 2, 411, 412, 5, 4, 3, 2, 412, 413, 7,
	53, 2, 2, 413, 414, 8, 29, 1, 2, 414, 57, 3, 2, 2, 2, 415, 416, 7, 16,
	2, 2, 416, 417, 7, 52, 2, 2, 417, 418, 5, 4, 3, 2, 418, 419, 7, 53, 2,
	2, 419, 420, 8, 30, 1, 2, 420, 59, 3, 2, 2, 2, 421, 422, 7, 17, 2, 2, 422,
	423, 7, 31, 2, 2, 423, 424, 7, 18, 2, 2, 424, 425, 5, 70, 36, 2, 425, 426,
	7, 32, 2, 2, 426, 427, 5, 70, 36, 2, 427, 428, 7, 52, 2, 2, 428, 429, 5,
	4, 3, 2, 429, 430, 7, 53, 2, 2, 430, 431, 8, 31, 1, 2, 431, 442, 3, 2,
	2, 2, 432, 433, 7, 17, 2, 2, 433, 434, 7, 31, 2, 2, 434, 435, 7, 18, 2,
	2, 435, 436, 5, 70, 36, 2, 436, 437, 7, 52, 2, 2, 437, 438, 5, 4, 3, 2,
	438, 439, 7, 53, 2, 2, 439, 440, 8, 31, 1, 2, 440, 442, 3, 2, 2, 2, 441,
	421, 3, 2, 2, 2, 441, 432, 3, 2, 2, 2, 442, 61, 3, 2, 2, 2, 443, 444, 7,
	13, 2, 2, 444, 445, 5, 6, 4, 2, 445, 446, 8, 32, 1, 2, 446, 453, 3, 2,
	2, 2, 447, 448, 7, 13, 2, 2, 448, 449, 5, 70, 36, 2, 449, 450, 5, 6, 4,
	2, 450, 451, 8, 32, 1, 2, 451, 453, 3, 2, 2, 2, 452, 443, 3, 2, 2, 2, 452,
	447, 3, 2, 2, 2, 453, 63, 3, 2, 2, 2, 454, 455, 7, 14, 2, 2, 455, 456,
	5, 6, 4, 2, 456, 457, 8, 33, 1, 2, 457, 65, 3, 2, 2, 2, 458, 459, 7, 15,
	2, 2, 459, 460, 5, 70, 36, 2, 460, 461, 5, 6, 4, 2, 461, 462, 8, 34, 1,
	2, 462, 67, 3, 2, 2, 2, 463, 464, 7, 21, 2, 2, 464, 474, 8, 35, 1, 2, 465,
	466, 7, 22, 2, 2, 466, 474, 8, 35, 1, 2, 467, 468, 7, 23, 2, 2, 468, 474,
	8, 35, 1, 2, 469, 470, 7, 25, 2, 2, 470, 474, 8, 35, 1, 2, 471, 472, 7,
	24, 2, 2, 472, 474, 8, 35, 1, 2, 473, 463, 3, 2, 2, 2, 473, 465, 3, 2,
	2, 2, 473, 467, 3, 2, 2, 2, 473, 469, 3, 2, 2, 2, 473, 471, 3, 2, 2, 2,
	474, 69, 3, 2, 2, 2, 475, 476, 5, 72, 37, 2, 476, 477, 8, 36, 1, 2, 477,
	71, 3, 2, 2, 2, 478, 479, 8, 37, 1, 2, 479, 480, 7, 50, 2, 2, 480, 481,
	5, 72, 37, 2, 481, 482, 9, 2, 2, 2, 482, 483, 7, 51, 2, 2, 483, 484, 9,
	3, 2, 2, 484, 485, 5, 72, 37, 18, 485, 486, 8, 37, 1, 2, 486, 556, 3, 2,
	2, 2, 487, 488, 7, 50, 2, 2, 488, 489, 5, 72, 37, 2, 489, 490, 9, 2, 2,
	2, 490, 491, 7, 51, 2, 2, 491, 492, 9, 3, 2, 2, 492, 493, 7, 50, 2, 2,
	493, 494, 5, 72, 37, 2, 494, 495, 9, 2, 2, 2, 495, 496, 7, 51, 2, 2, 496,
	497, 8, 37, 1, 2, 497, 556, 3, 2, 2, 2, 498, 499, 7, 50, 2, 2, 499, 500,
	5, 72, 37, 2, 500, 501, 9, 2, 2, 2, 501, 502, 7, 51, 2, 2, 502, 503, 9,
	4, 2, 2, 503, 504, 5, 72, 37, 14, 504, 505, 8, 37, 1, 2, 505, 556, 3, 2,
	2, 2, 506, 507, 7, 50, 2, 2, 507, 508, 5, 72, 37, 2, 508, 509, 9, 2, 2,
	2, 509, 510, 7, 51, 2, 2, 510, 511, 9, 4, 2, 2, 511, 512, 7, 50, 2, 2,
	512, 513, 5, 72, 37, 2, 513, 514, 9, 2, 2, 2, 514, 515, 7, 51, 2, 2, 515,
	516, 8, 37, 1, 2, 516, 556, 3, 2, 2, 2, 517, 518, 7, 50, 2, 2, 518, 519,
	5, 72, 37, 2, 519, 520, 9, 2, 2, 2, 520, 521, 7, 51, 2, 2, 521, 522, 9,
	5, 2, 2, 522, 523, 5, 72, 37, 10, 523, 524, 8, 37, 1, 2, 524, 556, 3, 2,
	2, 2, 525, 526, 7, 50, 2, 2, 526, 527, 5, 72, 37, 2, 527, 528, 9, 2, 2,
	2, 528, 529, 7, 51, 2, 2, 529, 530, 9, 5, 2, 2, 530, 531, 7, 50, 2, 2,
	531, 532, 5, 72, 37, 2, 532, 533, 9, 2, 2, 2, 533, 534, 7, 51, 2, 2, 534,
	535, 8, 37, 1, 2, 535, 556, 3, 2, 2, 2, 536, 537, 9, 6, 2, 2, 537, 538,
	5, 70, 36, 2, 538, 539, 8, 37, 1, 2, 539, 556, 3, 2, 2, 2, 540, 541, 9,
	6, 2, 2, 541, 542, 7, 50, 2, 2, 542, 543, 5, 72, 37, 2, 543, 544, 9, 2,
	2, 2, 544, 545, 7, 51, 2, 2, 545, 546, 8, 37, 1, 2, 546, 556, 3, 2, 2,
	2, 547, 548, 5, 74, 38, 2, 548, 549, 8, 37, 1, 2, 549, 556, 3, 2, 2, 2,
	550, 551, 7, 50, 2, 2, 551, 552, 5, 70, 36, 2, 552, 553, 7, 51, 2, 2, 553,
	554, 8, 37, 1, 2, 554, 556, 3, 2, 2, 2, 555, 478, 3, 2, 2, 2, 555, 487,
	3, 2, 2, 2, 555, 498, 3, 2, 2, 2, 555, 506, 3, 2, 2, 2, 555, 517, 3, 2,
	2, 2, 555, 525, 3, 2, 2, 2, 555, 536, 3, 2, 2, 2, 555, 540, 3, 2, 2, 2,
	555, 547, 3, 2, 2, 2, 555, 550, 3, 2, 2, 2, 556, 603, 3, 2, 2, 2, 557,
	558, 12, 19, 2, 2, 558, 559, 9, 3, 2, 2, 559, 560, 5, 72, 37, 20, 560,
	561, 8, 37, 1, 2, 561, 602, 3, 2, 2, 2, 562, 563, 12, 15, 2, 2, 563, 564,
	9, 4, 2, 2, 564, 565, 5, 72, 37, 16, 565, 566, 8, 37, 1, 2, 566, 602, 3,
	2, 2, 2, 567, 568, 12, 11, 2, 2, 568, 569, 9, 5, 2, 2, 569, 570, 5, 72,
	37, 12, 570, 571, 8, 37, 1, 2, 571, 602, 3, 2, 2, 2, 572, 573, 12, 7, 2,
	2, 573, 574, 9, 7, 2, 2, 574, 575, 5, 72, 37, 8, 575, 576, 8, 37, 1, 2,
	576, 602, 3, 2, 2, 2, 577, 578, 12, 17, 2, 2, 578, 579, 9, 3, 2, 2, 579,
	580, 7, 50, 2, 2, 580, 581, 5, 72, 37, 2, 581, 582, 9, 2, 2, 2, 582, 583,
	7, 51, 2, 2, 583, 584, 8, 37, 1, 2, 584, 602, 3, 2, 2, 2, 585, 586, 12,
	13, 2, 2, 586, 587, 9, 4, 2, 2, 587, 588, 7, 50, 2, 2, 588, 589, 5, 72,
	37, 2, 589, 590, 9, 2, 2, 2, 590, 591, 7, 51, 2, 2, 591, 592, 8, 37, 1,
	2, 592, 602, 3, 2, 2, 2, 593, 594, 12, 9, 2, 2, 594, 595, 9, 5, 2, 2, 595,
	596, 7, 50, 2, 2, 596, 597, 5, 72, 37, 2, 597, 598, 9, 2, 2, 2, 598, 599,
	7, 51, 2, 2, 599, 600, 8, 37, 1, 2, 600, 602, 3, 2, 2, 2, 601, 557, 3,
	2, 2, 2, 601, 562, 3, 2, 2, 2, 601, 567, 3, 2, 2, 2, 601, 572, 3, 2, 2,
	2, 601, 577, 3, 2, 2, 2, 601, 585, 3, 2, 2, 2, 601, 593, 3, 2, 2, 2, 602,
	605, 3, 2, 2, 2, 603, 601, 3, 2, 2, 2, 603, 604, 3, 2, 2, 2, 604, 73, 3,
	2, 2, 2, 605, 603, 3, 2, 2, 2, 606, 607, 7, 26, 2, 2, 607, 626, 8, 38,
	1, 2, 608, 609, 7, 27, 2, 2, 609, 626, 8, 38, 1, 2, 610, 611, 7, 29, 2,
	2, 611, 626, 8, 38, 1, 2, 612, 613, 7, 30, 2, 2, 613, 626, 8, 38, 1, 2,
	614, 615, 7, 31, 2, 2, 615, 626, 8, 38, 1, 2, 616, 617, 5, 22, 12, 2, 617,
	618, 8, 38, 1, 2, 618, 626, 3, 2, 2, 2, 619, 620, 5, 42, 22, 2, 620, 621,
	8, 38, 1, 2, 621, 626, 3, 2, 2, 2, 622, 623, 5, 58, 30, 2, 623, 624, 8,
	38, 1, 2, 624, 626, 3, 2, 2, 2, 625, 606, 3, 2, 2, 2, 625, 608, 3, 2, 2,
	2, 625, 610, 3, 2, 2, 2, 625, 612, 3, 2, 2, 2, 625, 614, 3, 2, 2, 2, 625,
	616, 3, 2, 2, 2, 625, 619, 3, 2, 2, 2, 625, 622, 3, 2, 2, 2, 626, 75, 3,
	2, 2, 2, 29, 82, 89, 125, 152, 188, 223, 228, 260, 265, 285, 300, 308,
	315, 326, 341, 346, 365, 376, 383, 394, 441, 452, 473, 555, 601, 603, 625,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'println!'", "'number'", "'as f64'", "'as i64'", "'let'", "'mut'",
	"'if'", "'else'", "'match'", "'while'", "'break'", "'continue'", "'return'",
	"'loop'", "'for'", "'in'", "'main'", "'fn'", "'i64'", "'f64'", "'String'",
	"'bool'", "'&str'", "", "", "", "", "", "", "'..'", "'.'", "';'", "','",
	"':'", "'='", "'=='", "'>='", "'=>'", "'<='", "'!='", "'>'", "'<'", "'*'",
	"'/'", "'%'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'&&'",
	"'||'", "'|'", "'!'",
}
var symbolicNames = []string{
	"", "R_PRINTLN", "P_NUMBER", "R_AS_DOUBLE", "R_AS_INTEGER", "R_LET", "R_MUT",
	"R_IF", "R_ELSE", "R_MATCH", "R_WHILE", "R_BREAK", "R_CONTINUE", "R_RETURN",
	"R_LOOP", "R_FOR", "R_IN", "R_MAIN", "R_FUNCTION", "R_INT", "R_FLOAT",
	"R_STRING", "R_BOOL", "R_STR", "NUMBER", "DOUBLE", "CHAR", "STRING", "BOOLEAN",
	"ID", "TK_DOBLEPUNTO", "TK_PUNTO", "TK_PUNTOCOMA", "TK_COMA", "TK_DOSPUNTOS",
	"TK_IGUAL", "TK_IGUALIGUAL", "TK_MAYORIGUAL", "TK_IGUALMAYOR", "TK_MENORIGUAL",
	"TK_DIFIGUAL", "TK_MAYOR", "TK_MENOR", "TK_MULT", "TK_DIV", "TK_MODULO",
	"TK_MAS", "TK_MENOS", "TK_PARA", "TK_PARC", "TK_LLAVEA", "TK_LLAVEC", "TK_CORA",
	"TK_CORC", "TK_AND", "TK_OR", "TK_BARRA", "TK_NOT", "WHITESPACE", "TK_MULTI",
	"TK_LINE",
}

var ruleNames = []string{
	"start", "instrucciones", "end_instr", "instruccion", "instr_main", "instr_println",
	"instr_declaracion", "instr_asignacion", "instr_if", "instr_else_if", "instr_ternario",
	"instr_else_if_ternario", "instr_match", "instr_case", "block_instr_match",
	"list_case", "list_expre_case", "block_case", "instr_default", "block_default",
	"instr_match_ter", "instr_case_ter", "list_case_ternario", "list_expre_case_ter",
	"block_case_ter", "instr_default_ter", "instr_while", "instr_loop", "instr_loop_ternario",
	"instr_for_in", "instr_break", "instr_continue", "instr_return", "instr_tipo",
	"expression", "exp_arit", "primitivo",
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
	ChemsR_IF          = 7
	ChemsR_ELSE        = 8
	ChemsR_MATCH       = 9
	ChemsR_WHILE       = 10
	ChemsR_BREAK       = 11
	ChemsR_CONTINUE    = 12
	ChemsR_RETURN      = 13
	ChemsR_LOOP        = 14
	ChemsR_FOR         = 15
	ChemsR_IN          = 16
	ChemsR_MAIN        = 17
	ChemsR_FUNCTION    = 18
	ChemsR_INT         = 19
	ChemsR_FLOAT       = 20
	ChemsR_STRING      = 21
	ChemsR_BOOL        = 22
	ChemsR_STR         = 23
	ChemsNUMBER        = 24
	ChemsDOUBLE        = 25
	ChemsCHAR          = 26
	ChemsSTRING        = 27
	ChemsBOOLEAN       = 28
	ChemsID            = 29
	ChemsTK_DOBLEPUNTO = 30
	ChemsTK_PUNTO      = 31
	ChemsTK_PUNTOCOMA  = 32
	ChemsTK_COMA       = 33
	ChemsTK_DOSPUNTOS  = 34
	ChemsTK_IGUAL      = 35
	ChemsTK_IGUALIGUAL = 36
	ChemsTK_MAYORIGUAL = 37
	ChemsTK_IGUALMAYOR = 38
	ChemsTK_MENORIGUAL = 39
	ChemsTK_DIFIGUAL   = 40
	ChemsTK_MAYOR      = 41
	ChemsTK_MENOR      = 42
	ChemsTK_MULT       = 43
	ChemsTK_DIV        = 44
	ChemsTK_MODULO     = 45
	ChemsTK_MAS        = 46
	ChemsTK_MENOS      = 47
	ChemsTK_PARA       = 48
	ChemsTK_PARC       = 49
	ChemsTK_LLAVEA     = 50
	ChemsTK_LLAVEC     = 51
	ChemsTK_CORA       = 52
	ChemsTK_CORC       = 53
	ChemsTK_AND        = 54
	ChemsTK_OR         = 55
	ChemsTK_BARRA      = 56
	ChemsTK_NOT        = 57
	ChemsWHITESPACE    = 58
	ChemsTK_MULTI      = 59
	ChemsTK_LINE       = 60
)

// Chems rules.
const (
	ChemsRULE_start                  = 0
	ChemsRULE_instrucciones          = 1
	ChemsRULE_end_instr              = 2
	ChemsRULE_instruccion            = 3
	ChemsRULE_instr_main             = 4
	ChemsRULE_instr_println          = 5
	ChemsRULE_instr_declaracion      = 6
	ChemsRULE_instr_asignacion       = 7
	ChemsRULE_instr_if               = 8
	ChemsRULE_instr_else_if          = 9
	ChemsRULE_instr_ternario         = 10
	ChemsRULE_instr_else_if_ternario = 11
	ChemsRULE_instr_match            = 12
	ChemsRULE_instr_case             = 13
	ChemsRULE_block_instr_match      = 14
	ChemsRULE_list_case              = 15
	ChemsRULE_list_expre_case        = 16
	ChemsRULE_block_case             = 17
	ChemsRULE_instr_default          = 18
	ChemsRULE_block_default          = 19
	ChemsRULE_instr_match_ter        = 20
	ChemsRULE_instr_case_ter         = 21
	ChemsRULE_list_case_ternario     = 22
	ChemsRULE_list_expre_case_ter    = 23
	ChemsRULE_block_case_ter         = 24
	ChemsRULE_instr_default_ter      = 25
	ChemsRULE_instr_while            = 26
	ChemsRULE_instr_loop             = 27
	ChemsRULE_instr_loop_ternario    = 28
	ChemsRULE_instr_for_in           = 29
	ChemsRULE_instr_break            = 30
	ChemsRULE_instr_continue         = 31
	ChemsRULE_instr_return           = 32
	ChemsRULE_instr_tipo             = 33
	ChemsRULE_expression             = 34
	ChemsRULE_exp_arit               = 35
	ChemsRULE_primitivo              = 36
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
		p.SetState(74)

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

	// GetList returns the list rule context list.
	GetList() []IInstruccionContext

	// SetList sets the list rule context list.
	SetList([]IInstruccionContext)

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
	list         []IInstruccionContext
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

func (s *InstruccionesContext) GetList() []IInstruccionContext { return s.list }

func (s *InstruccionesContext) SetList(v []IInstruccionContext) { s.list = v }

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
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_PRINTLN)|(1<<ChemsR_LET)|(1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_WHILE)|(1<<ChemsR_BREAK)|(1<<ChemsR_CONTINUE)|(1<<ChemsR_LOOP)|(1<<ChemsR_FOR)|(1<<ChemsR_FUNCTION)|(1<<ChemsID))) != 0) {
		{
			p.SetState(77)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).list = append(localctx.(*InstruccionesContext).list, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetList()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IEnd_instrContext is an interface to support dynamic dispatch.
type IEnd_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v attribute.
	GetV() int

	// SetV sets the v attribute.
	SetV(int)

	// IsEnd_instrContext differentiates from other interfaces.
	IsEnd_instrContext()
}

type End_instrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      int
}

func NewEmptyEnd_instrContext() *End_instrContext {
	var p = new(End_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_end_instr
	return p
}

func (*End_instrContext) IsEnd_instrContext() {}

func NewEnd_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *End_instrContext {
	var p = new(End_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_end_instr

	return p
}

func (s *End_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *End_instrContext) GetV() int { return s.v }

func (s *End_instrContext) SetV(v int) { s.v = v }

func (s *End_instrContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *End_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *End_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *End_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterEnd_instr(s)
	}
}

func (s *End_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitEnd_instr(s)
	}
}

func (p *Chems) End_instr() (localctx IEnd_instrContext) {
	localctx = NewEnd_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ChemsRULE_end_instr)

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

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_PUNTOCOMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*End_instrContext).v = 1

	case ChemsEOF, ChemsR_PRINTLN, ChemsR_LET, ChemsR_IF, ChemsR_MATCH, ChemsR_WHILE, ChemsR_BREAK, ChemsR_CONTINUE, ChemsR_LOOP, ChemsR_FOR, ChemsR_FUNCTION, ChemsID, ChemsTK_COMA, ChemsTK_LLAVEC:
		p.EnterOuterAlt(localctx, 2)
		localctx.(*End_instrContext).v = 0

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_println returns the _instr_println rule contexts.
	Get_instr_println() IInstr_printlnContext

	// Get_instr_declaracion returns the _instr_declaracion rule contexts.
	Get_instr_declaracion() IInstr_declaracionContext

	// Get_instr_asignacion returns the _instr_asignacion rule contexts.
	Get_instr_asignacion() IInstr_asignacionContext

	// Get_instr_if returns the _instr_if rule contexts.
	Get_instr_if() IInstr_ifContext

	// Get_instr_match returns the _instr_match rule contexts.
	Get_instr_match() IInstr_matchContext

	// Get_instr_while returns the _instr_while rule contexts.
	Get_instr_while() IInstr_whileContext

	// Get_instr_break returns the _instr_break rule contexts.
	Get_instr_break() IInstr_breakContext

	// Get_instr_continue returns the _instr_continue rule contexts.
	Get_instr_continue() IInstr_continueContext

	// Get_instr_loop returns the _instr_loop rule contexts.
	Get_instr_loop() IInstr_loopContext

	// Get_instr_for_in returns the _instr_for_in rule contexts.
	Get_instr_for_in() IInstr_for_inContext

	// Get_instr_main returns the _instr_main rule contexts.
	Get_instr_main() IInstr_mainContext

	// Set_instr_println sets the _instr_println rule contexts.
	Set_instr_println(IInstr_printlnContext)

	// Set_instr_declaracion sets the _instr_declaracion rule contexts.
	Set_instr_declaracion(IInstr_declaracionContext)

	// Set_instr_asignacion sets the _instr_asignacion rule contexts.
	Set_instr_asignacion(IInstr_asignacionContext)

	// Set_instr_if sets the _instr_if rule contexts.
	Set_instr_if(IInstr_ifContext)

	// Set_instr_match sets the _instr_match rule contexts.
	Set_instr_match(IInstr_matchContext)

	// Set_instr_while sets the _instr_while rule contexts.
	Set_instr_while(IInstr_whileContext)

	// Set_instr_break sets the _instr_break rule contexts.
	Set_instr_break(IInstr_breakContext)

	// Set_instr_continue sets the _instr_continue rule contexts.
	Set_instr_continue(IInstr_continueContext)

	// Set_instr_loop sets the _instr_loop rule contexts.
	Set_instr_loop(IInstr_loopContext)

	// Set_instr_for_in sets the _instr_for_in rule contexts.
	Set_instr_for_in(IInstr_for_inContext)

	// Set_instr_main sets the _instr_main rule contexts.
	Set_instr_main(IInstr_mainContext)

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
	_instr_println     IInstr_printlnContext
	_instr_declaracion IInstr_declaracionContext
	_instr_asignacion  IInstr_asignacionContext
	_instr_if          IInstr_ifContext
	_instr_match       IInstr_matchContext
	_instr_while       IInstr_whileContext
	_instr_break       IInstr_breakContext
	_instr_continue    IInstr_continueContext
	_instr_loop        IInstr_loopContext
	_instr_for_in      IInstr_for_inContext
	_instr_main        IInstr_mainContext
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

func (s *InstruccionContext) Get_instr_println() IInstr_printlnContext { return s._instr_println }

func (s *InstruccionContext) Get_instr_declaracion() IInstr_declaracionContext {
	return s._instr_declaracion
}

func (s *InstruccionContext) Get_instr_asignacion() IInstr_asignacionContext {
	return s._instr_asignacion
}

func (s *InstruccionContext) Get_instr_if() IInstr_ifContext { return s._instr_if }

func (s *InstruccionContext) Get_instr_match() IInstr_matchContext { return s._instr_match }

func (s *InstruccionContext) Get_instr_while() IInstr_whileContext { return s._instr_while }

func (s *InstruccionContext) Get_instr_break() IInstr_breakContext { return s._instr_break }

func (s *InstruccionContext) Get_instr_continue() IInstr_continueContext { return s._instr_continue }

func (s *InstruccionContext) Get_instr_loop() IInstr_loopContext { return s._instr_loop }

func (s *InstruccionContext) Get_instr_for_in() IInstr_for_inContext { return s._instr_for_in }

func (s *InstruccionContext) Get_instr_main() IInstr_mainContext { return s._instr_main }

func (s *InstruccionContext) Set_instr_println(v IInstr_printlnContext) { s._instr_println = v }

func (s *InstruccionContext) Set_instr_declaracion(v IInstr_declaracionContext) {
	s._instr_declaracion = v
}

func (s *InstruccionContext) Set_instr_asignacion(v IInstr_asignacionContext) { s._instr_asignacion = v }

func (s *InstruccionContext) Set_instr_if(v IInstr_ifContext) { s._instr_if = v }

func (s *InstruccionContext) Set_instr_match(v IInstr_matchContext) { s._instr_match = v }

func (s *InstruccionContext) Set_instr_while(v IInstr_whileContext) { s._instr_while = v }

func (s *InstruccionContext) Set_instr_break(v IInstr_breakContext) { s._instr_break = v }

func (s *InstruccionContext) Set_instr_continue(v IInstr_continueContext) { s._instr_continue = v }

func (s *InstruccionContext) Set_instr_loop(v IInstr_loopContext) { s._instr_loop = v }

func (s *InstruccionContext) Set_instr_for_in(v IInstr_for_inContext) { s._instr_for_in = v }

func (s *InstruccionContext) Set_instr_main(v IInstr_mainContext) { s._instr_main = v }

func (s *InstruccionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *InstruccionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *InstruccionContext) Instr_println() IInstr_printlnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_printlnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_printlnContext)
}

func (s *InstruccionContext) End_instr() IEnd_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnd_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnd_instrContext)
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

func (s *InstruccionContext) Instr_if() IInstr_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_ifContext)
}

func (s *InstruccionContext) Instr_match() IInstr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_matchContext)
}

func (s *InstruccionContext) Instr_while() IInstr_whileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_whileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_whileContext)
}

func (s *InstruccionContext) Instr_break() IInstr_breakContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_breakContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_breakContext)
}

func (s *InstruccionContext) Instr_continue() IInstr_continueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_continueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_continueContext)
}

func (s *InstruccionContext) Instr_loop() IInstr_loopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_loopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_loopContext)
}

func (s *InstruccionContext) Instr_for_in() IInstr_for_inContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_for_inContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_for_inContext)
}

func (s *InstruccionContext) Instr_main() IInstr_mainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_mainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_mainContext)
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
	p.EnterRule(localctx, 6, ChemsRULE_instruccion)

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

	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_PRINTLN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)

			var _x = p.Instr_println()

			localctx.(*InstruccionContext)._instr_println = _x
		}
		{
			p.SetState(90)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_println().GetInstr()

	case ChemsR_LET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)

			var _x = p.Instr_declaracion()

			localctx.(*InstruccionContext)._instr_declaracion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_declaracion().GetInstr()

	case ChemsID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)

			var _x = p.Instr_asignacion()

			localctx.(*InstruccionContext)._instr_asignacion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_asignacion().GetInstr()

	case ChemsR_IF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(99)

			var _x = p.Instr_if()

			localctx.(*InstruccionContext)._instr_if = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_if().GetInstr()

	case ChemsR_MATCH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(102)

			var _x = p.Instr_match()

			localctx.(*InstruccionContext)._instr_match = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_match().GetInstr()

	case ChemsR_WHILE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(105)

			var _x = p.Instr_while()

			localctx.(*InstruccionContext)._instr_while = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_while().GetInstr()

	case ChemsR_BREAK:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(108)

			var _x = p.Instr_break()

			localctx.(*InstruccionContext)._instr_break = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_break().GetInstr()

	case ChemsR_CONTINUE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(111)

			var _x = p.Instr_continue()

			localctx.(*InstruccionContext)._instr_continue = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_continue().GetInstr()

	case ChemsR_LOOP:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(114)

			var _x = p.Instr_loop()

			localctx.(*InstruccionContext)._instr_loop = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_loop().GetInstr()

	case ChemsR_FOR:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(117)

			var _x = p.Instr_for_in()

			localctx.(*InstruccionContext)._instr_for_in = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_for_in().GetInstr()

	case ChemsR_FUNCTION:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(120)

			var _x = p.Instr_main()

			localctx.(*InstruccionContext)._instr_main = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_main().GetInstr()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInstr_mainContext is an interface to support dynamic dispatch.
type IInstr_mainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_MAIN returns the _R_MAIN token.
	Get_R_MAIN() antlr.Token

	// Set_R_MAIN sets the _R_MAIN token.
	Set_R_MAIN(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_mainContext differentiates from other interfaces.
	IsInstr_mainContext()
}

type Instr_mainContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_MAIN        antlr.Token
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_mainContext() *Instr_mainContext {
	var p = new(Instr_mainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_main
	return p
}

func (*Instr_mainContext) IsInstr_mainContext() {}

func NewInstr_mainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_mainContext {
	var p = new(Instr_mainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_main

	return p
}

func (s *Instr_mainContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_mainContext) Get_R_MAIN() antlr.Token { return s._R_MAIN }

func (s *Instr_mainContext) Set_R_MAIN(v antlr.Token) { s._R_MAIN = v }

func (s *Instr_mainContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_mainContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_mainContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_mainContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_mainContext) R_FUNCTION() antlr.TerminalNode {
	return s.GetToken(ChemsR_FUNCTION, 0)
}

func (s *Instr_mainContext) R_MAIN() antlr.TerminalNode {
	return s.GetToken(ChemsR_MAIN, 0)
}

func (s *Instr_mainContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_mainContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_mainContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_mainContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_mainContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_mainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_mainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_mainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_main(s)
	}
}

func (s *Instr_mainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_main(s)
	}
}

func (p *Chems) Instr_main() (localctx IInstr_mainContext) {
	localctx = NewInstr_mainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ChemsRULE_instr_main)

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
		p.SetState(125)
		p.Match(ChemsR_FUNCTION)
	}
	{
		p.SetState(126)

		var _m = p.Match(ChemsR_MAIN)

		localctx.(*Instr_mainContext)._R_MAIN = _m
	}
	{
		p.SetState(127)
		p.Match(ChemsTK_PARA)
	}
	{
		p.SetState(128)
		p.Match(ChemsTK_PARC)
	}
	{
		p.SetState(129)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(130)

		var _x = p.Instrucciones()

		localctx.(*Instr_mainContext)._instrucciones = _x
	}
	{
		p.SetState(131)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_mainContext).instr = instruction.NewMain(localctx.(*Instr_mainContext).Get_instrucciones().GetL(), (func() int {
		if localctx.(*Instr_mainContext).Get_R_MAIN() == nil {
			return 0
		} else {
			return localctx.(*Instr_mainContext).Get_R_MAIN().GetLine()
		}
	}()), localctx.(*Instr_mainContext).Get_R_MAIN().GetColumn())

	return localctx
}

// IInstr_printlnContext is an interface to support dynamic dispatch.
type IInstr_printlnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_printlnContext differentiates from other interfaces.
	IsInstr_printlnContext()
}

type Instr_printlnContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruction
	_expression IExpressionContext
	_STRING     antlr.Token
}

func NewEmptyInstr_printlnContext() *Instr_printlnContext {
	var p = new(Instr_printlnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_println
	return p
}

func (*Instr_printlnContext) IsInstr_printlnContext() {}

func NewInstr_printlnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_printlnContext {
	var p = new(Instr_printlnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_println

	return p
}

func (s *Instr_printlnContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_printlnContext) Get_STRING() antlr.Token { return s._STRING }

func (s *Instr_printlnContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *Instr_printlnContext) Get_expression() IExpressionContext { return s._expression }

func (s *Instr_printlnContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Instr_printlnContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_printlnContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_printlnContext) R_PRINTLN() antlr.TerminalNode {
	return s.GetToken(ChemsR_PRINTLN, 0)
}

func (s *Instr_printlnContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_printlnContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instr_printlnContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_printlnContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *Instr_printlnContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChemsSTRING, 0)
}

func (s *Instr_printlnContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_printlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_printlnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_printlnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_println(s)
	}
}

func (s *Instr_printlnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_println(s)
	}
}

func (p *Chems) Instr_println() (localctx IInstr_printlnContext) {
	localctx = NewInstr_printlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ChemsRULE_instr_println)

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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Match(ChemsR_PRINTLN)
		}
		{
			p.SetState(135)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(136)

			var _x = p.Expression()

			localctx.(*Instr_printlnContext)._expression = _x
		}
		{
			p.SetState(137)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(138)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_printlnContext).instr = instruction.PRINTLN(localctx.(*Instr_printlnContext).Get_expression().GetP(), "-1")

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Match(ChemsR_PRINTLN)
		}
		{
			p.SetState(142)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(143)

			var _m = p.Match(ChemsSTRING)

			localctx.(*Instr_printlnContext)._STRING = _m
		}
		{
			p.SetState(144)
			p.Match(ChemsTK_COMA)
		}
		{
			p.SetState(145)

			var _x = p.Expression()

			localctx.(*Instr_printlnContext)._expression = _x
		}
		{
			p.SetState(146)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(147)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_printlnContext).instr = instruction.PRINTLN(localctx.(*Instr_printlnContext).Get_expression().GetP(), (func() string {
			if localctx.(*Instr_printlnContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*Instr_printlnContext).Get_STRING().GetText()
			}
		}())[1:len((func() string {
			if localctx.(*Instr_printlnContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*Instr_printlnContext).Get_STRING().GetText()
			}
		}()))-1])

	}

	return localctx
}

// IInstr_declaracionContext is an interface to support dynamic dispatch.
type IInstr_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_LET returns the _R_LET token.
	Get_R_LET() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_R_LET sets the _R_LET token.
	Set_R_LET(antlr.Token)

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
	_R_LET      antlr.Token
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

func (s *Instr_declaracionContext) Get_R_LET() antlr.Token { return s._R_LET }

func (s *Instr_declaracionContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_declaracionContext) Set_R_LET(v antlr.Token) { s._R_LET = v }

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
	p.EnterRule(localctx, 12, ChemsRULE_instr_declaracion)

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

	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(153)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(154)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(155)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(156)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(157)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), interfaces.NULL, localctx.(*Instr_declaracionContext).Get_expression().GetP(), true, false, false, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(161)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(162)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(163)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(164)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(165)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(166)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(167)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_declaracionContext).Get_instr_tipo().GetTipo_exp(), localctx.(*Instr_declaracionContext).Get_expression().GetP(), true, false, false, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(170)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(171)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(172)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(173)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(174)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), interfaces.NULL, localctx.(*Instr_declaracionContext).Get_expression().GetP(), false, false, false, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(177)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(178)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(179)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(180)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(181)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(182)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(183)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_declaracionContext).Get_instr_tipo().GetTipo_exp(), localctx.(*Instr_declaracionContext).Get_expression().GetP(), false, false, false, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

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
	p.EnterRule(localctx, 14, ChemsRULE_instr_asignacion)

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
		p.SetState(188)

		var _m = p.Match(ChemsID)

		localctx.(*Instr_asignacionContext)._ID = _m
	}
	{
		p.SetState(189)
		p.Match(ChemsTK_IGUAL)
	}
	{
		p.SetState(190)

		var _x = p.Expression()

		localctx.(*Instr_asignacionContext)._expression = _x
	}
	{
		p.SetState(191)
		p.Match(ChemsTK_PUNTOCOMA)
	}
	localctx.(*Instr_asignacionContext).instr = variable.NewAssignment((func() string {
		if localctx.(*Instr_asignacionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Instr_asignacionContext).Get_ID().GetText()
		}
	}()), localctx.(*Instr_asignacionContext).Get_expression().GetP(), (func() int {
		if localctx.(*Instr_asignacionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Instr_asignacionContext).Get_ID().GetLine()
		}
	}()), localctx.(*Instr_asignacionContext).Get_ID().GetColumn())

	return localctx
}

// IInstr_ifContext is an interface to support dynamic dispatch.
type IInstr_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_IF returns the _R_IF token.
	Get_R_IF() antlr.Token

	// Set_R_IF sets the _R_IF token.
	Set_R_IF(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// GetLeft_instr returns the left_instr rule contexts.
	GetLeft_instr() IInstruccionesContext

	// GetRight_intr returns the right_intr rule contexts.
	GetRight_intr() IInstruccionesContext

	// Get_instr_else_if returns the _instr_else_if rule contexts.
	Get_instr_else_if() IInstr_else_ifContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// SetLeft_instr sets the left_instr rule contexts.
	SetLeft_instr(IInstruccionesContext)

	// SetRight_intr sets the right_intr rule contexts.
	SetRight_intr(IInstruccionesContext)

	// Set_instr_else_if sets the _instr_else_if rule contexts.
	Set_instr_else_if(IInstr_else_ifContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_ifContext differentiates from other interfaces.
	IsInstr_ifContext()
}

type Instr_ifContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_IF          antlr.Token
	_expression    IExpressionContext
	left_instr     IInstruccionesContext
	right_intr     IInstruccionesContext
	_instr_else_if IInstr_else_ifContext
}

func NewEmptyInstr_ifContext() *Instr_ifContext {
	var p = new(Instr_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_if
	return p
}

func (*Instr_ifContext) IsInstr_ifContext() {}

func NewInstr_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_ifContext {
	var p = new(Instr_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_if

	return p
}

func (s *Instr_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_ifContext) Get_R_IF() antlr.Token { return s._R_IF }

func (s *Instr_ifContext) Set_R_IF(v antlr.Token) { s._R_IF = v }

func (s *Instr_ifContext) Get_expression() IExpressionContext { return s._expression }

func (s *Instr_ifContext) GetLeft_instr() IInstruccionesContext { return s.left_instr }

func (s *Instr_ifContext) GetRight_intr() IInstruccionesContext { return s.right_intr }

func (s *Instr_ifContext) Get_instr_else_if() IInstr_else_ifContext { return s._instr_else_if }

func (s *Instr_ifContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Instr_ifContext) SetLeft_instr(v IInstruccionesContext) { s.left_instr = v }

func (s *Instr_ifContext) SetRight_intr(v IInstruccionesContext) { s.right_intr = v }

func (s *Instr_ifContext) Set_instr_else_if(v IInstr_else_ifContext) { s._instr_else_if = v }

func (s *Instr_ifContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_ifContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_ifContext) R_IF() antlr.TerminalNode {
	return s.GetToken(ChemsR_IF, 0)
}

func (s *Instr_ifContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instr_ifContext) AllTK_LLAVEA() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_LLAVEA)
}

func (s *Instr_ifContext) TK_LLAVEA(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, i)
}

func (s *Instr_ifContext) AllTK_LLAVEC() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_LLAVEC)
}

func (s *Instr_ifContext) TK_LLAVEC(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, i)
}

func (s *Instr_ifContext) AllInstrucciones() []IInstruccionesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem())
	var tst = make([]IInstruccionesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionesContext)
		}
	}

	return tst
}

func (s *Instr_ifContext) Instrucciones(i int) IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_ifContext) R_ELSE() antlr.TerminalNode {
	return s.GetToken(ChemsR_ELSE, 0)
}

func (s *Instr_ifContext) Instr_else_if() IInstr_else_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_else_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_else_ifContext)
}

func (s *Instr_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_if(s)
	}
}

func (s *Instr_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_if(s)
	}
}

func (p *Chems) Instr_if() (localctx IInstr_ifContext) {
	localctx = NewInstr_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ChemsRULE_instr_if)

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

	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(194)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(195)

			var _x = p.Expression()

			localctx.(*Instr_ifContext)._expression = _x
		}
		{
			p.SetState(196)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(197)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(198)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_ifContext).instr = control.NewIf(localctx.(*Instr_ifContext).Get_expression().GetP(), localctx.(*Instr_ifContext).GetLeft_instr().GetL(), nil, nil, (func() int {
			if localctx.(*Instr_ifContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ifContext).Get_R_IF().GetLine()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(201)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(202)

			var _x = p.Expression()

			localctx.(*Instr_ifContext)._expression = _x
		}
		{
			p.SetState(203)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(204)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(205)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(206)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(207)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(208)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).right_intr = _x
		}
		{
			p.SetState(209)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_ifContext).instr = control.NewIf(localctx.(*Instr_ifContext).Get_expression().GetP(), localctx.(*Instr_ifContext).GetLeft_instr().GetL(), localctx.(*Instr_ifContext).GetRight_intr().GetL(), nil, (func() int {
			if localctx.(*Instr_ifContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ifContext).Get_R_IF().GetLine()
			}
		}()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(212)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(213)

			var _x = p.Expression()

			localctx.(*Instr_ifContext)._expression = _x
		}
		{
			p.SetState(214)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(215)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(216)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(217)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(218)

			var _x = p.Instr_else_if()

			localctx.(*Instr_ifContext)._instr_else_if = _x
		}
		localctx.(*Instr_ifContext).instr = control.NewIf(localctx.(*Instr_ifContext).Get_expression().GetP(), localctx.(*Instr_ifContext).GetLeft_instr().GetL(), nil, localctx.(*Instr_ifContext).Get_instr_else_if().GetL(), (func() int {
			if localctx.(*Instr_ifContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ifContext).Get_R_IF().GetLine()
			}
		}()))

	}

	return localctx
}

// IInstr_else_ifContext is an interface to support dynamic dispatch.
type IInstr_else_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_if returns the _instr_if rule contexts.
	Get_instr_if() IInstr_ifContext

	// Set_instr_if sets the _instr_if rule contexts.
	Set_instr_if(IInstr_ifContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_ifContext

	// SetE sets the e rule context list.
	SetE([]IInstr_ifContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstr_else_ifContext differentiates from other interfaces.
	IsInstr_else_ifContext()
}

type Instr_else_ifContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	l         *arrayList.List
	_instr_if IInstr_ifContext
	e         []IInstr_ifContext
}

func NewEmptyInstr_else_ifContext() *Instr_else_ifContext {
	var p = new(Instr_else_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_else_if
	return p
}

func (*Instr_else_ifContext) IsInstr_else_ifContext() {}

func NewInstr_else_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_else_ifContext {
	var p = new(Instr_else_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_else_if

	return p
}

func (s *Instr_else_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_else_ifContext) Get_instr_if() IInstr_ifContext { return s._instr_if }

func (s *Instr_else_ifContext) Set_instr_if(v IInstr_ifContext) { s._instr_if = v }

func (s *Instr_else_ifContext) GetE() []IInstr_ifContext { return s.e }

func (s *Instr_else_ifContext) SetE(v []IInstr_ifContext) { s.e = v }

func (s *Instr_else_ifContext) GetL() *arrayList.List { return s.l }

func (s *Instr_else_ifContext) SetL(v *arrayList.List) { s.l = v }

func (s *Instr_else_ifContext) AllInstr_if() []IInstr_ifContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_ifContext)(nil)).Elem())
	var tst = make([]IInstr_ifContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_ifContext)
		}
	}

	return tst
}

func (s *Instr_else_ifContext) Instr_if(i int) IInstr_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_ifContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_ifContext)
}

func (s *Instr_else_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_else_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_else_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_else_if(s)
	}
}

func (s *Instr_else_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_else_if(s)
	}
}

func (p *Chems) Instr_else_if() (localctx IInstr_else_ifContext) {
	localctx = NewInstr_else_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ChemsRULE_instr_else_if)

	localctx.(*Instr_else_ifContext).l = arrayList.New()

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(223)

				var _x = p.Instr_if()

				localctx.(*Instr_else_ifContext)._instr_if = _x
			}
			localctx.(*Instr_else_ifContext).e = append(localctx.(*Instr_else_ifContext).e, localctx.(*Instr_else_ifContext)._instr_if)

		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	listInt := localctx.(*Instr_else_ifContext).GetE()
	for _, e := range listInt {
		localctx.(*Instr_else_ifContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IInstr_ternarioContext is an interface to support dynamic dispatch.
type IInstr_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_IF returns the _R_IF token.
	Get_R_IF() antlr.Token

	// Set_R_IF sets the _R_IF token.
	Set_R_IF(antlr.Token)

	// GetBlock returns the block rule contexts.
	GetBlock() IExpressionContext

	// GetLeft_instr returns the left_instr rule contexts.
	GetLeft_instr() IExpressionContext

	// GetRight_intr returns the right_intr rule contexts.
	GetRight_intr() IExpressionContext

	// Get_instr_else_if_ternario returns the _instr_else_if_ternario rule contexts.
	Get_instr_else_if_ternario() IInstr_else_if_ternarioContext

	// SetBlock sets the block rule contexts.
	SetBlock(IExpressionContext)

	// SetLeft_instr sets the left_instr rule contexts.
	SetLeft_instr(IExpressionContext)

	// SetRight_intr sets the right_intr rule contexts.
	SetRight_intr(IExpressionContext)

	// Set_instr_else_if_ternario sets the _instr_else_if_ternario rule contexts.
	Set_instr_else_if_ternario(IInstr_else_if_ternarioContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsInstr_ternarioContext differentiates from other interfaces.
	IsInstr_ternarioContext()
}

type Instr_ternarioContext struct {
	*antlr.BaseParserRuleContext
	parser                  antlr.Parser
	p                       interfaces.Expresion
	_R_IF                   antlr.Token
	block                   IExpressionContext
	left_instr              IExpressionContext
	right_intr              IExpressionContext
	_instr_else_if_ternario IInstr_else_if_ternarioContext
}

func NewEmptyInstr_ternarioContext() *Instr_ternarioContext {
	var p = new(Instr_ternarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_ternario
	return p
}

func (*Instr_ternarioContext) IsInstr_ternarioContext() {}

func NewInstr_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_ternarioContext {
	var p = new(Instr_ternarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_ternario

	return p
}

func (s *Instr_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_ternarioContext) Get_R_IF() antlr.Token { return s._R_IF }

func (s *Instr_ternarioContext) Set_R_IF(v antlr.Token) { s._R_IF = v }

func (s *Instr_ternarioContext) GetBlock() IExpressionContext { return s.block }

func (s *Instr_ternarioContext) GetLeft_instr() IExpressionContext { return s.left_instr }

func (s *Instr_ternarioContext) GetRight_intr() IExpressionContext { return s.right_intr }

func (s *Instr_ternarioContext) Get_instr_else_if_ternario() IInstr_else_if_ternarioContext {
	return s._instr_else_if_ternario
}

func (s *Instr_ternarioContext) SetBlock(v IExpressionContext) { s.block = v }

func (s *Instr_ternarioContext) SetLeft_instr(v IExpressionContext) { s.left_instr = v }

func (s *Instr_ternarioContext) SetRight_intr(v IExpressionContext) { s.right_intr = v }

func (s *Instr_ternarioContext) Set_instr_else_if_ternario(v IInstr_else_if_ternarioContext) {
	s._instr_else_if_ternario = v
}

func (s *Instr_ternarioContext) GetP() interfaces.Expresion { return s.p }

func (s *Instr_ternarioContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *Instr_ternarioContext) R_IF() antlr.TerminalNode {
	return s.GetToken(ChemsR_IF, 0)
}

func (s *Instr_ternarioContext) AllTK_LLAVEA() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_LLAVEA)
}

func (s *Instr_ternarioContext) TK_LLAVEA(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, i)
}

func (s *Instr_ternarioContext) AllTK_LLAVEC() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_LLAVEC)
}

func (s *Instr_ternarioContext) TK_LLAVEC(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, i)
}

func (s *Instr_ternarioContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Instr_ternarioContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instr_ternarioContext) R_ELSE() antlr.TerminalNode {
	return s.GetToken(ChemsR_ELSE, 0)
}

func (s *Instr_ternarioContext) Instr_else_if_ternario() IInstr_else_if_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_else_if_ternarioContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_else_if_ternarioContext)
}

func (s *Instr_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_ternario(s)
	}
}

func (s *Instr_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_ternario(s)
	}
}

func (p *Chems) Instr_ternario() (localctx IInstr_ternarioContext) {
	localctx = NewInstr_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ChemsRULE_instr_ternario)

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

	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(231)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(232)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).block = _x
		}
		{
			p.SetState(233)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(234)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(235)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_ternarioContext).p = control.NewTernario(localctx.(*Instr_ternarioContext).GetBlock().GetP(), localctx.(*Instr_ternarioContext).GetLeft_instr().GetP(), nil, nil, (func() int {
			if localctx.(*Instr_ternarioContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ternarioContext).Get_R_IF().GetLine()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(239)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).block = _x
		}
		{
			p.SetState(240)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(241)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(242)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(243)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(244)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(245)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).right_intr = _x
		}
		{
			p.SetState(246)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_ternarioContext).p = control.NewTernario(localctx.(*Instr_ternarioContext).GetBlock().GetP(), localctx.(*Instr_ternarioContext).GetLeft_instr().GetP(), localctx.(*Instr_ternarioContext).GetRight_intr().GetP(), nil, (func() int {
			if localctx.(*Instr_ternarioContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ternarioContext).Get_R_IF().GetLine()
			}
		}()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(249)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(250)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).block = _x
		}
		{
			p.SetState(251)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(252)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(253)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(254)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(255)

			var _x = p.Instr_else_if_ternario()

			localctx.(*Instr_ternarioContext)._instr_else_if_ternario = _x
		}
		localctx.(*Instr_ternarioContext).p = control.NewTernario(localctx.(*Instr_ternarioContext).GetBlock().GetP(), localctx.(*Instr_ternarioContext).GetLeft_instr().GetP(), nil, localctx.(*Instr_ternarioContext).Get_instr_else_if_ternario().GetL(), (func() int {
			if localctx.(*Instr_ternarioContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ternarioContext).Get_R_IF().GetLine()
			}
		}()))

	}

	return localctx
}

// IInstr_else_if_ternarioContext is an interface to support dynamic dispatch.
type IInstr_else_if_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_ternario returns the _instr_ternario rule contexts.
	Get_instr_ternario() IInstr_ternarioContext

	// Set_instr_ternario sets the _instr_ternario rule contexts.
	Set_instr_ternario(IInstr_ternarioContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_ternarioContext

	// SetE sets the e rule context list.
	SetE([]IInstr_ternarioContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstr_else_if_ternarioContext differentiates from other interfaces.
	IsInstr_else_if_ternarioContext()
}

type Instr_else_if_ternarioContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	l               *arrayList.List
	_instr_ternario IInstr_ternarioContext
	e               []IInstr_ternarioContext
}

func NewEmptyInstr_else_if_ternarioContext() *Instr_else_if_ternarioContext {
	var p = new(Instr_else_if_ternarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_else_if_ternario
	return p
}

func (*Instr_else_if_ternarioContext) IsInstr_else_if_ternarioContext() {}

func NewInstr_else_if_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_else_if_ternarioContext {
	var p = new(Instr_else_if_ternarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_else_if_ternario

	return p
}

func (s *Instr_else_if_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_else_if_ternarioContext) Get_instr_ternario() IInstr_ternarioContext {
	return s._instr_ternario
}

func (s *Instr_else_if_ternarioContext) Set_instr_ternario(v IInstr_ternarioContext) {
	s._instr_ternario = v
}

func (s *Instr_else_if_ternarioContext) GetE() []IInstr_ternarioContext { return s.e }

func (s *Instr_else_if_ternarioContext) SetE(v []IInstr_ternarioContext) { s.e = v }

func (s *Instr_else_if_ternarioContext) GetL() *arrayList.List { return s.l }

func (s *Instr_else_if_ternarioContext) SetL(v *arrayList.List) { s.l = v }

func (s *Instr_else_if_ternarioContext) AllInstr_ternario() []IInstr_ternarioContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_ternarioContext)(nil)).Elem())
	var tst = make([]IInstr_ternarioContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_ternarioContext)
		}
	}

	return tst
}

func (s *Instr_else_if_ternarioContext) Instr_ternario(i int) IInstr_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_ternarioContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_ternarioContext)
}

func (s *Instr_else_if_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_else_if_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_else_if_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_else_if_ternario(s)
	}
}

func (s *Instr_else_if_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_else_if_ternario(s)
	}
}

func (p *Chems) Instr_else_if_ternario() (localctx IInstr_else_if_ternarioContext) {
	localctx = NewInstr_else_if_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ChemsRULE_instr_else_if_ternario)

	localctx.(*Instr_else_if_ternarioContext).l = arrayList.New()

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(260)

				var _x = p.Instr_ternario()

				localctx.(*Instr_else_if_ternarioContext)._instr_ternario = _x
			}
			localctx.(*Instr_else_if_ternarioContext).e = append(localctx.(*Instr_else_if_ternarioContext).e, localctx.(*Instr_else_if_ternarioContext)._instr_ternario)

		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	listInt := localctx.(*Instr_else_if_ternarioContext).GetE()
	for _, e := range listInt {
		localctx.(*Instr_else_if_ternarioContext).l.Add(e.GetP())
	}

	return localctx
}

// IInstr_matchContext is an interface to support dynamic dispatch.
type IInstr_matchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_MATCH returns the _R_MATCH token.
	Get_R_MATCH() antlr.Token

	// Set_R_MATCH sets the _R_MATCH token.
	Set_R_MATCH(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_list_case returns the _list_case rule contexts.
	Get_list_case() IList_caseContext

	// Get_block_default returns the _block_default rule contexts.
	Get_block_default() IBlock_defaultContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_list_case sets the _list_case rule contexts.
	Set_list_case(IList_caseContext)

	// Set_block_default sets the _block_default rule contexts.
	Set_block_default(IBlock_defaultContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_matchContext differentiates from other interfaces.
	IsInstr_matchContext()
}

type Instr_matchContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_MATCH       antlr.Token
	_expression    IExpressionContext
	_list_case     IList_caseContext
	_block_default IBlock_defaultContext
}

func NewEmptyInstr_matchContext() *Instr_matchContext {
	var p = new(Instr_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_match
	return p
}

func (*Instr_matchContext) IsInstr_matchContext() {}

func NewInstr_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_matchContext {
	var p = new(Instr_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_match

	return p
}

func (s *Instr_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_matchContext) Get_R_MATCH() antlr.Token { return s._R_MATCH }

func (s *Instr_matchContext) Set_R_MATCH(v antlr.Token) { s._R_MATCH = v }

func (s *Instr_matchContext) Get_expression() IExpressionContext { return s._expression }

func (s *Instr_matchContext) Get_list_case() IList_caseContext { return s._list_case }

func (s *Instr_matchContext) Get_block_default() IBlock_defaultContext { return s._block_default }

func (s *Instr_matchContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Instr_matchContext) Set_list_case(v IList_caseContext) { s._list_case = v }

func (s *Instr_matchContext) Set_block_default(v IBlock_defaultContext) { s._block_default = v }

func (s *Instr_matchContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_matchContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_matchContext) R_MATCH() antlr.TerminalNode {
	return s.GetToken(ChemsR_MATCH, 0)
}

func (s *Instr_matchContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instr_matchContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_matchContext) List_case() IList_caseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_caseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_caseContext)
}

func (s *Instr_matchContext) Block_default() IBlock_defaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_defaultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_defaultContext)
}

func (s *Instr_matchContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_match(s)
	}
}

func (s *Instr_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_match(s)
	}
}

func (p *Chems) Instr_match() (localctx IInstr_matchContext) {
	localctx = NewInstr_matchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ChemsRULE_instr_match)

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

	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_matchContext)._R_MATCH = _m
		}
		{
			p.SetState(269)

			var _x = p.Expression()

			localctx.(*Instr_matchContext)._expression = _x
		}
		{
			p.SetState(270)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(271)

			var _x = p.List_case()

			localctx.(*Instr_matchContext)._list_case = _x
		}
		{
			p.SetState(272)

			var _x = p.Block_default()

			localctx.(*Instr_matchContext)._block_default = _x
		}
		{
			p.SetState(273)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_matchContext).instr = control.NewMatch(localctx.(*Instr_matchContext).Get_expression().GetP(), localctx.(*Instr_matchContext).Get_list_case().GetL(), localctx.(*Instr_matchContext).Get_block_default().GetL(), (func() int {
			if localctx.(*Instr_matchContext).Get_R_MATCH() == nil {
				return 0
			} else {
				return localctx.(*Instr_matchContext).Get_R_MATCH().GetLine()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_matchContext)._R_MATCH = _m
		}
		{
			p.SetState(277)

			var _x = p.Expression()

			localctx.(*Instr_matchContext)._expression = _x
		}
		{
			p.SetState(278)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(279)

			var _x = p.Block_default()

			localctx.(*Instr_matchContext)._block_default = _x
		}
		{
			p.SetState(280)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_matchContext).instr = control.NewMatch(localctx.(*Instr_matchContext).Get_expression().GetP(), nil, localctx.(*Instr_matchContext).Get_block_default().GetL(), (func() int {
			if localctx.(*Instr_matchContext).Get_R_MATCH() == nil {
				return 0
			} else {
				return localctx.(*Instr_matchContext).Get_R_MATCH().GetLine()
			}
		}()))

	}

	return localctx
}

// IInstr_caseContext is an interface to support dynamic dispatch.
type IInstr_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_list_expre_case returns the _list_expre_case rule contexts.
	Get_list_expre_case() IList_expre_caseContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_block_instr_match returns the _block_instr_match rule contexts.
	Get_block_instr_match() IBlock_instr_matchContext

	// Set_list_expre_case sets the _list_expre_case rule contexts.
	Set_list_expre_case(IList_expre_caseContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_block_instr_match sets the _block_instr_match rule contexts.
	Set_block_instr_match(IBlock_instr_matchContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Expresion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Expresion)

	// IsInstr_caseContext differentiates from other interfaces.
	IsInstr_caseContext()
}

type Instr_caseContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	instr              interfaces.Expresion
	_list_expre_case   IList_expre_caseContext
	_instrucciones     IInstruccionesContext
	_block_instr_match IBlock_instr_matchContext
}

func NewEmptyInstr_caseContext() *Instr_caseContext {
	var p = new(Instr_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_case
	return p
}

func (*Instr_caseContext) IsInstr_caseContext() {}

func NewInstr_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_caseContext {
	var p = new(Instr_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_case

	return p
}

func (s *Instr_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_caseContext) Get_list_expre_case() IList_expre_caseContext { return s._list_expre_case }

func (s *Instr_caseContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_caseContext) Get_block_instr_match() IBlock_instr_matchContext {
	return s._block_instr_match
}

func (s *Instr_caseContext) Set_list_expre_case(v IList_expre_caseContext) { s._list_expre_case = v }

func (s *Instr_caseContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_caseContext) Set_block_instr_match(v IBlock_instr_matchContext) {
	s._block_instr_match = v
}

func (s *Instr_caseContext) GetInstr() interfaces.Expresion { return s.instr }

func (s *Instr_caseContext) SetInstr(v interfaces.Expresion) { s.instr = v }

func (s *Instr_caseContext) List_expre_case() IList_expre_caseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expre_caseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expre_caseContext)
}

func (s *Instr_caseContext) TK_IGUALMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALMAYOR, 0)
}

func (s *Instr_caseContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_caseContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_caseContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_caseContext) Block_instr_match() IBlock_instr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_instr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_instr_matchContext)
}

func (s *Instr_caseContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_case(s)
	}
}

func (s *Instr_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_case(s)
	}
}

func (p *Chems) Instr_case() (localctx IInstr_caseContext) {
	localctx = NewInstr_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ChemsRULE_instr_case)

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

	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)

			var _x = p.List_expre_case()

			localctx.(*Instr_caseContext)._list_expre_case = _x
		}
		{
			p.SetState(286)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(287)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(288)

			var _x = p.Instrucciones()

			localctx.(*Instr_caseContext)._instrucciones = _x
		}
		{
			p.SetState(289)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_caseContext).instr = control.NewCase(nil, localctx.(*Instr_caseContext).Get_list_expre_case().GetL(), localctx.(*Instr_caseContext).Get_instrucciones().GetL())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)

			var _x = p.List_expre_case()

			localctx.(*Instr_caseContext)._list_expre_case = _x
		}
		{
			p.SetState(293)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(294)

			var _x = p.Block_instr_match()

			localctx.(*Instr_caseContext)._block_instr_match = _x
		}
		{
			p.SetState(295)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Instr_caseContext).instr = control.NewCase(nil, localctx.(*Instr_caseContext).Get_list_expre_case().GetL(), localctx.(*Instr_caseContext).Get_block_instr_match().GetL())

	}

	return localctx
}

// IBlock_instr_matchContext is an interface to support dynamic dispatch.
type IBlock_instr_matchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetList returns the list rule context list.
	GetList() []IInstruccionContext

	// SetList sets the list rule context list.
	SetList([]IInstruccionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsBlock_instr_matchContext differentiates from other interfaces.
	IsBlock_instr_matchContext()
}

type Block_instr_matchContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	_instruccion IInstruccionContext
	list         []IInstruccionContext
}

func NewEmptyBlock_instr_matchContext() *Block_instr_matchContext {
	var p = new(Block_instr_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_instr_match
	return p
}

func (*Block_instr_matchContext) IsBlock_instr_matchContext() {}

func NewBlock_instr_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_instr_matchContext {
	var p = new(Block_instr_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_instr_match

	return p
}

func (s *Block_instr_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_instr_matchContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *Block_instr_matchContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *Block_instr_matchContext) GetList() []IInstruccionContext { return s.list }

func (s *Block_instr_matchContext) SetList(v []IInstruccionContext) { s.list = v }

func (s *Block_instr_matchContext) GetL() *arrayList.List { return s.l }

func (s *Block_instr_matchContext) SetL(v *arrayList.List) { s.l = v }

func (s *Block_instr_matchContext) Instruccion() IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *Block_instr_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_instr_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_instr_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_instr_match(s)
	}
}

func (s *Block_instr_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_instr_match(s)
	}
}

func (p *Chems) Block_instr_match() (localctx IBlock_instr_matchContext) {
	localctx = NewBlock_instr_matchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ChemsRULE_block_instr_match)

	localctx.(*Block_instr_matchContext).l = arrayList.New()

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
		p.SetState(300)

		var _x = p.Instruccion()

		localctx.(*Block_instr_matchContext)._instruccion = _x
	}
	localctx.(*Block_instr_matchContext).list = append(localctx.(*Block_instr_matchContext).list, localctx.(*Block_instr_matchContext)._instruccion)

	listInt := localctx.(*Block_instr_matchContext).GetList()
	for _, e := range listInt {
		localctx.(*Block_instr_matchContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IList_caseContext is an interface to support dynamic dispatch.
type IList_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_case returns the _instr_case rule contexts.
	Get_instr_case() IInstr_caseContext

	// Set_instr_case sets the _instr_case rule contexts.
	Set_instr_case(IInstr_caseContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_caseContext

	// SetE sets the e rule context list.
	SetE([]IInstr_caseContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_caseContext differentiates from other interfaces.
	IsList_caseContext()
}

type List_caseContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	_instr_case IInstr_caseContext
	e           []IInstr_caseContext
}

func NewEmptyList_caseContext() *List_caseContext {
	var p = new(List_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_case
	return p
}

func (*List_caseContext) IsList_caseContext() {}

func NewList_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_caseContext {
	var p = new(List_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_case

	return p
}

func (s *List_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *List_caseContext) Get_instr_case() IInstr_caseContext { return s._instr_case }

func (s *List_caseContext) Set_instr_case(v IInstr_caseContext) { s._instr_case = v }

func (s *List_caseContext) GetE() []IInstr_caseContext { return s.e }

func (s *List_caseContext) SetE(v []IInstr_caseContext) { s.e = v }

func (s *List_caseContext) GetL() *arrayList.List { return s.l }

func (s *List_caseContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_caseContext) AllInstr_case() []IInstr_caseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_caseContext)(nil)).Elem())
	var tst = make([]IInstr_caseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_caseContext)
		}
	}

	return tst
}

func (s *List_caseContext) Instr_case(i int) IInstr_caseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_caseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_caseContext)
}

func (s *List_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_case(s)
	}
}

func (s *List_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_case(s)
	}
}

func (p *Chems) List_case() (localctx IList_caseContext) {
	localctx = NewList_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ChemsRULE_list_case)

	localctx.(*List_caseContext).l = arrayList.New()

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(303)

				var _x = p.Instr_case()

				localctx.(*List_caseContext)._instr_case = _x
			}
			localctx.(*List_caseContext).e = append(localctx.(*List_caseContext).e, localctx.(*List_caseContext)._instr_case)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	listInt := localctx.(*List_caseContext).GetE()
	for _, e := range listInt {
		localctx.(*List_caseContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IList_expre_caseContext is an interface to support dynamic dispatch.
type IList_expre_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_case returns the _block_case rule contexts.
	Get_block_case() IBlock_caseContext

	// Set_block_case sets the _block_case rule contexts.
	Set_block_case(IBlock_caseContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_caseContext

	// SetE sets the e rule context list.
	SetE([]IBlock_caseContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_expre_caseContext differentiates from other interfaces.
	IsList_expre_caseContext()
}

type List_expre_caseContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	_block_case IBlock_caseContext
	e           []IBlock_caseContext
}

func NewEmptyList_expre_caseContext() *List_expre_caseContext {
	var p = new(List_expre_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_expre_case
	return p
}

func (*List_expre_caseContext) IsList_expre_caseContext() {}

func NewList_expre_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expre_caseContext {
	var p = new(List_expre_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_expre_case

	return p
}

func (s *List_expre_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expre_caseContext) Get_block_case() IBlock_caseContext { return s._block_case }

func (s *List_expre_caseContext) Set_block_case(v IBlock_caseContext) { s._block_case = v }

func (s *List_expre_caseContext) GetE() []IBlock_caseContext { return s.e }

func (s *List_expre_caseContext) SetE(v []IBlock_caseContext) { s.e = v }

func (s *List_expre_caseContext) GetL() *arrayList.List { return s.l }

func (s *List_expre_caseContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_expre_caseContext) AllBlock_case() []IBlock_caseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlock_caseContext)(nil)).Elem())
	var tst = make([]IBlock_caseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlock_caseContext)
		}
	}

	return tst
}

func (s *List_expre_caseContext) Block_case(i int) IBlock_caseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_caseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlock_caseContext)
}

func (s *List_expre_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expre_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expre_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_expre_case(s)
	}
}

func (s *List_expre_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_expre_case(s)
	}
}

func (p *Chems) List_expre_case() (localctx IList_expre_caseContext) {
	localctx = NewList_expre_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ChemsRULE_list_expre_case)

	localctx.(*List_expre_caseContext).l = arrayList.New()

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
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(ChemsTK_MENOS-47))|(1<<(ChemsTK_PARA-47))|(1<<(ChemsTK_NOT-47)))) != 0) {
		{
			p.SetState(310)

			var _x = p.Block_case()

			localctx.(*List_expre_caseContext)._block_case = _x
		}
		localctx.(*List_expre_caseContext).e = append(localctx.(*List_expre_caseContext).e, localctx.(*List_expre_caseContext)._block_case)

		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_expre_caseContext).GetE()
	for _, e := range listInt {
		localctx.(*List_expre_caseContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IBlock_caseContext is an interface to support dynamic dispatch.
type IBlock_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Expresion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Expresion)

	// IsBlock_caseContext differentiates from other interfaces.
	IsBlock_caseContext()
}

type Block_caseContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Expresion
	_expression IExpressionContext
}

func NewEmptyBlock_caseContext() *Block_caseContext {
	var p = new(Block_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_case
	return p
}

func (*Block_caseContext) IsBlock_caseContext() {}

func NewBlock_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_caseContext {
	var p = new(Block_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_case

	return p
}

func (s *Block_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_caseContext) Get_expression() IExpressionContext { return s._expression }

func (s *Block_caseContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Block_caseContext) GetInstr() interfaces.Expresion { return s.instr }

func (s *Block_caseContext) SetInstr(v interfaces.Expresion) { s.instr = v }

func (s *Block_caseContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Block_caseContext) TK_BARRA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_BARRA, 0)
}

func (s *Block_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_case(s)
	}
}

func (s *Block_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_case(s)
	}
}

func (p *Chems) Block_case() (localctx IBlock_caseContext) {
	localctx = NewBlock_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ChemsRULE_block_case)

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

	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)

			var _x = p.Expression()

			localctx.(*Block_caseContext)._expression = _x
		}
		{
			p.SetState(318)
			p.Match(ChemsTK_BARRA)
		}
		localctx.(*Block_caseContext).instr = control.NewCase(localctx.(*Block_caseContext).Get_expression().GetP(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(321)

			var _x = p.Expression()

			localctx.(*Block_caseContext)._expression = _x
		}
		localctx.(*Block_caseContext).instr = control.NewCase(localctx.(*Block_caseContext).Get_expression().GetP(), nil, nil)

	}

	return localctx
}

// IInstr_defaultContext is an interface to support dynamic dispatch.
type IInstr_defaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_block_instr_match returns the _block_instr_match rule contexts.
	Get_block_instr_match() IBlock_instr_matchContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_block_instr_match sets the _block_instr_match rule contexts.
	Set_block_instr_match(IBlock_instr_matchContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_defaultContext differentiates from other interfaces.
	IsInstr_defaultContext()
}

type Instr_defaultContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	instr              interfaces.Instruction
	_instrucciones     IInstruccionesContext
	_block_instr_match IBlock_instr_matchContext
}

func NewEmptyInstr_defaultContext() *Instr_defaultContext {
	var p = new(Instr_defaultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_default
	return p
}

func (*Instr_defaultContext) IsInstr_defaultContext() {}

func NewInstr_defaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_defaultContext {
	var p = new(Instr_defaultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_default

	return p
}

func (s *Instr_defaultContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_defaultContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_defaultContext) Get_block_instr_match() IBlock_instr_matchContext {
	return s._block_instr_match
}

func (s *Instr_defaultContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_defaultContext) Set_block_instr_match(v IBlock_instr_matchContext) {
	s._block_instr_match = v
}

func (s *Instr_defaultContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_defaultContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_defaultContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_defaultContext) TK_IGUALMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALMAYOR, 0)
}

func (s *Instr_defaultContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_defaultContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_defaultContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_defaultContext) Block_instr_match() IBlock_instr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_instr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_instr_matchContext)
}

func (s *Instr_defaultContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_defaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_defaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_defaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_default(s)
	}
}

func (s *Instr_defaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_default(s)
	}
}

func (p *Chems) Instr_default() (localctx IInstr_defaultContext) {
	localctx = NewInstr_defaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ChemsRULE_instr_default)

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

	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(326)
			p.Match(ChemsID)
		}
		{
			p.SetState(327)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(328)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(329)

			var _x = p.Instrucciones()

			localctx.(*Instr_defaultContext)._instrucciones = _x
		}
		{
			p.SetState(330)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_defaultContext).instr = control.NewDefault(localctx.(*Instr_defaultContext).Get_instrucciones().GetL())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(333)
			p.Match(ChemsID)
		}
		{
			p.SetState(334)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(335)

			var _x = p.Block_instr_match()

			localctx.(*Instr_defaultContext)._block_instr_match = _x
		}
		{
			p.SetState(336)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Instr_defaultContext).instr = control.NewDefault(localctx.(*Instr_defaultContext).Get_block_instr_match().GetL())

	}

	return localctx
}

// IBlock_defaultContext is an interface to support dynamic dispatch.
type IBlock_defaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_default returns the _instr_default rule contexts.
	Get_instr_default() IInstr_defaultContext

	// Set_instr_default sets the _instr_default rule contexts.
	Set_instr_default(IInstr_defaultContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_defaultContext

	// SetE sets the e rule context list.
	SetE([]IInstr_defaultContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsBlock_defaultContext differentiates from other interfaces.
	IsBlock_defaultContext()
}

type Block_defaultContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	l              *arrayList.List
	_instr_default IInstr_defaultContext
	e              []IInstr_defaultContext
}

func NewEmptyBlock_defaultContext() *Block_defaultContext {
	var p = new(Block_defaultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_default
	return p
}

func (*Block_defaultContext) IsBlock_defaultContext() {}

func NewBlock_defaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_defaultContext {
	var p = new(Block_defaultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_default

	return p
}

func (s *Block_defaultContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_defaultContext) Get_instr_default() IInstr_defaultContext { return s._instr_default }

func (s *Block_defaultContext) Set_instr_default(v IInstr_defaultContext) { s._instr_default = v }

func (s *Block_defaultContext) GetE() []IInstr_defaultContext { return s.e }

func (s *Block_defaultContext) SetE(v []IInstr_defaultContext) { s.e = v }

func (s *Block_defaultContext) GetL() *arrayList.List { return s.l }

func (s *Block_defaultContext) SetL(v *arrayList.List) { s.l = v }

func (s *Block_defaultContext) AllInstr_default() []IInstr_defaultContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_defaultContext)(nil)).Elem())
	var tst = make([]IInstr_defaultContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_defaultContext)
		}
	}

	return tst
}

func (s *Block_defaultContext) Instr_default(i int) IInstr_defaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_defaultContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_defaultContext)
}

func (s *Block_defaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_defaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_defaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_default(s)
	}
}

func (s *Block_defaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_default(s)
	}
}

func (p *Chems) Block_default() (localctx IBlock_defaultContext) {
	localctx = NewBlock_defaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ChemsRULE_block_default)

	localctx.(*Block_defaultContext).l = arrayList.New()

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
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ChemsID {
		{
			p.SetState(341)

			var _x = p.Instr_default()

			localctx.(*Block_defaultContext)._instr_default = _x
		}
		localctx.(*Block_defaultContext).e = append(localctx.(*Block_defaultContext).e, localctx.(*Block_defaultContext)._instr_default)

		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*Block_defaultContext).GetE()
	for _, e := range listInt {
		localctx.(*Block_defaultContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IInstr_match_terContext is an interface to support dynamic dispatch.
type IInstr_match_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_MATCH returns the _R_MATCH token.
	Get_R_MATCH() antlr.Token

	// Set_R_MATCH sets the _R_MATCH token.
	Set_R_MATCH(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_list_case_ternario returns the _list_case_ternario rule contexts.
	Get_list_case_ternario() IList_case_ternarioContext

	// Get_instr_default_ter returns the _instr_default_ter rule contexts.
	Get_instr_default_ter() IInstr_default_terContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_list_case_ternario sets the _list_case_ternario rule contexts.
	Set_list_case_ternario(IList_case_ternarioContext)

	// Set_instr_default_ter sets the _instr_default_ter rule contexts.
	Set_instr_default_ter(IInstr_default_terContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Expresion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Expresion)

	// IsInstr_match_terContext differentiates from other interfaces.
	IsInstr_match_terContext()
}

type Instr_match_terContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	instr               interfaces.Expresion
	_R_MATCH            antlr.Token
	_expression         IExpressionContext
	_list_case_ternario IList_case_ternarioContext
	_instr_default_ter  IInstr_default_terContext
}

func NewEmptyInstr_match_terContext() *Instr_match_terContext {
	var p = new(Instr_match_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_match_ter
	return p
}

func (*Instr_match_terContext) IsInstr_match_terContext() {}

func NewInstr_match_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_match_terContext {
	var p = new(Instr_match_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_match_ter

	return p
}

func (s *Instr_match_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_match_terContext) Get_R_MATCH() antlr.Token { return s._R_MATCH }

func (s *Instr_match_terContext) Set_R_MATCH(v antlr.Token) { s._R_MATCH = v }

func (s *Instr_match_terContext) Get_expression() IExpressionContext { return s._expression }

func (s *Instr_match_terContext) Get_list_case_ternario() IList_case_ternarioContext {
	return s._list_case_ternario
}

func (s *Instr_match_terContext) Get_instr_default_ter() IInstr_default_terContext {
	return s._instr_default_ter
}

func (s *Instr_match_terContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Instr_match_terContext) Set_list_case_ternario(v IList_case_ternarioContext) {
	s._list_case_ternario = v
}

func (s *Instr_match_terContext) Set_instr_default_ter(v IInstr_default_terContext) {
	s._instr_default_ter = v
}

func (s *Instr_match_terContext) GetInstr() interfaces.Expresion { return s.instr }

func (s *Instr_match_terContext) SetInstr(v interfaces.Expresion) { s.instr = v }

func (s *Instr_match_terContext) R_MATCH() antlr.TerminalNode {
	return s.GetToken(ChemsR_MATCH, 0)
}

func (s *Instr_match_terContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instr_match_terContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_match_terContext) List_case_ternario() IList_case_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_case_ternarioContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_case_ternarioContext)
}

func (s *Instr_match_terContext) Instr_default_ter() IInstr_default_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_default_terContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_default_terContext)
}

func (s *Instr_match_terContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_match_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_match_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_match_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_match_ter(s)
	}
}

func (s *Instr_match_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_match_ter(s)
	}
}

func (p *Chems) Instr_match_ter() (localctx IInstr_match_terContext) {
	localctx = NewInstr_match_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ChemsRULE_instr_match_ter)

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

	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_match_terContext)._R_MATCH = _m
		}
		{
			p.SetState(349)

			var _x = p.Expression()

			localctx.(*Instr_match_terContext)._expression = _x
		}
		{
			p.SetState(350)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(351)

			var _x = p.List_case_ternario()

			localctx.(*Instr_match_terContext)._list_case_ternario = _x
		}
		{
			p.SetState(352)

			var _x = p.Instr_default_ter()

			localctx.(*Instr_match_terContext)._instr_default_ter = _x
		}
		{
			p.SetState(353)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_match_terContext).instr = control.NewTerMatch(localctx.(*Instr_match_terContext).Get_expression().GetP(), localctx.(*Instr_match_terContext).Get_list_case_ternario().GetL(), localctx.(*Instr_match_terContext).Get_instr_default_ter().GetInstr(), (func() int {
			if localctx.(*Instr_match_terContext).Get_R_MATCH() == nil {
				return 0
			} else {
				return localctx.(*Instr_match_terContext).Get_R_MATCH().GetLine()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(356)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_match_terContext)._R_MATCH = _m
		}
		{
			p.SetState(357)

			var _x = p.Expression()

			localctx.(*Instr_match_terContext)._expression = _x
		}
		{
			p.SetState(358)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(359)

			var _x = p.Instr_default_ter()

			localctx.(*Instr_match_terContext)._instr_default_ter = _x
		}
		{
			p.SetState(360)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_match_terContext).instr = control.NewTerMatch(localctx.(*Instr_match_terContext).Get_expression().GetP(), nil, localctx.(*Instr_match_terContext).Get_instr_default_ter().GetInstr(), (func() int {
			if localctx.(*Instr_match_terContext).Get_R_MATCH() == nil {
				return 0
			} else {
				return localctx.(*Instr_match_terContext).Get_R_MATCH().GetLine()
			}
		}()))

	}

	return localctx
}

// IInstr_case_terContext is an interface to support dynamic dispatch.
type IInstr_case_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_list_expre_case_ter returns the _list_expre_case_ter rule contexts.
	Get_list_expre_case_ter() IList_expre_case_terContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_list_expre_case_ter sets the _list_expre_case_ter rule contexts.
	Set_list_expre_case_ter(IList_expre_case_terContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Expresion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Expresion)

	// IsInstr_case_terContext differentiates from other interfaces.
	IsInstr_case_terContext()
}

type Instr_case_terContext struct {
	*antlr.BaseParserRuleContext
	parser               antlr.Parser
	instr                interfaces.Expresion
	_list_expre_case_ter IList_expre_case_terContext
	_expression          IExpressionContext
}

func NewEmptyInstr_case_terContext() *Instr_case_terContext {
	var p = new(Instr_case_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_case_ter
	return p
}

func (*Instr_case_terContext) IsInstr_case_terContext() {}

func NewInstr_case_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_case_terContext {
	var p = new(Instr_case_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_case_ter

	return p
}

func (s *Instr_case_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_case_terContext) Get_list_expre_case_ter() IList_expre_case_terContext {
	return s._list_expre_case_ter
}

func (s *Instr_case_terContext) Get_expression() IExpressionContext { return s._expression }

func (s *Instr_case_terContext) Set_list_expre_case_ter(v IList_expre_case_terContext) {
	s._list_expre_case_ter = v
}

func (s *Instr_case_terContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Instr_case_terContext) GetInstr() interfaces.Expresion { return s.instr }

func (s *Instr_case_terContext) SetInstr(v interfaces.Expresion) { s.instr = v }

func (s *Instr_case_terContext) List_expre_case_ter() IList_expre_case_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expre_case_terContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expre_case_terContext)
}

func (s *Instr_case_terContext) TK_IGUALMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALMAYOR, 0)
}

func (s *Instr_case_terContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instr_case_terContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_case_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_case_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_case_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_case_ter(s)
	}
}

func (s *Instr_case_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_case_ter(s)
	}
}

func (p *Chems) Instr_case_ter() (localctx IInstr_case_terContext) {
	localctx = NewInstr_case_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ChemsRULE_instr_case_ter)

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
		p.SetState(365)

		var _x = p.List_expre_case_ter()

		localctx.(*Instr_case_terContext)._list_expre_case_ter = _x
	}
	{
		p.SetState(366)
		p.Match(ChemsTK_IGUALMAYOR)
	}
	{
		p.SetState(367)

		var _x = p.Expression()

		localctx.(*Instr_case_terContext)._expression = _x
	}
	{
		p.SetState(368)
		p.Match(ChemsTK_COMA)
	}
	localctx.(*Instr_case_terContext).instr = control.NewCaseTer(nil, localctx.(*Instr_case_terContext).Get_list_expre_case_ter().GetL(), localctx.(*Instr_case_terContext).Get_expression().GetP())

	return localctx
}

// IList_case_ternarioContext is an interface to support dynamic dispatch.
type IList_case_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_case_ter returns the _instr_case_ter rule contexts.
	Get_instr_case_ter() IInstr_case_terContext

	// Set_instr_case_ter sets the _instr_case_ter rule contexts.
	Set_instr_case_ter(IInstr_case_terContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_case_terContext

	// SetE sets the e rule context list.
	SetE([]IInstr_case_terContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_case_ternarioContext differentiates from other interfaces.
	IsList_case_ternarioContext()
}

type List_case_ternarioContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	l               *arrayList.List
	_instr_case_ter IInstr_case_terContext
	e               []IInstr_case_terContext
}

func NewEmptyList_case_ternarioContext() *List_case_ternarioContext {
	var p = new(List_case_ternarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_case_ternario
	return p
}

func (*List_case_ternarioContext) IsList_case_ternarioContext() {}

func NewList_case_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_case_ternarioContext {
	var p = new(List_case_ternarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_case_ternario

	return p
}

func (s *List_case_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *List_case_ternarioContext) Get_instr_case_ter() IInstr_case_terContext {
	return s._instr_case_ter
}

func (s *List_case_ternarioContext) Set_instr_case_ter(v IInstr_case_terContext) {
	s._instr_case_ter = v
}

func (s *List_case_ternarioContext) GetE() []IInstr_case_terContext { return s.e }

func (s *List_case_ternarioContext) SetE(v []IInstr_case_terContext) { s.e = v }

func (s *List_case_ternarioContext) GetL() *arrayList.List { return s.l }

func (s *List_case_ternarioContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_case_ternarioContext) AllInstr_case_ter() []IInstr_case_terContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_case_terContext)(nil)).Elem())
	var tst = make([]IInstr_case_terContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_case_terContext)
		}
	}

	return tst
}

func (s *List_case_ternarioContext) Instr_case_ter(i int) IInstr_case_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_case_terContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_case_terContext)
}

func (s *List_case_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_case_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_case_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_case_ternario(s)
	}
}

func (s *List_case_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_case_ternario(s)
	}
}

func (p *Chems) List_case_ternario() (localctx IList_case_ternarioContext) {
	localctx = NewList_case_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ChemsRULE_list_case_ternario)

	localctx.(*List_case_ternarioContext).l = arrayList.New()

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(371)

				var _x = p.Instr_case_ter()

				localctx.(*List_case_ternarioContext)._instr_case_ter = _x
			}
			localctx.(*List_case_ternarioContext).e = append(localctx.(*List_case_ternarioContext).e, localctx.(*List_case_ternarioContext)._instr_case_ter)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(374)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	listInt := localctx.(*List_case_ternarioContext).GetE()
	for _, e := range listInt {
		localctx.(*List_case_ternarioContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IList_expre_case_terContext is an interface to support dynamic dispatch.
type IList_expre_case_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_case_ter returns the _block_case_ter rule contexts.
	Get_block_case_ter() IBlock_case_terContext

	// Set_block_case_ter sets the _block_case_ter rule contexts.
	Set_block_case_ter(IBlock_case_terContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_case_terContext

	// SetE sets the e rule context list.
	SetE([]IBlock_case_terContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_expre_case_terContext differentiates from other interfaces.
	IsList_expre_case_terContext()
}

type List_expre_case_terContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	l               *arrayList.List
	_block_case_ter IBlock_case_terContext
	e               []IBlock_case_terContext
}

func NewEmptyList_expre_case_terContext() *List_expre_case_terContext {
	var p = new(List_expre_case_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_expre_case_ter
	return p
}

func (*List_expre_case_terContext) IsList_expre_case_terContext() {}

func NewList_expre_case_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expre_case_terContext {
	var p = new(List_expre_case_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_expre_case_ter

	return p
}

func (s *List_expre_case_terContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expre_case_terContext) Get_block_case_ter() IBlock_case_terContext {
	return s._block_case_ter
}

func (s *List_expre_case_terContext) Set_block_case_ter(v IBlock_case_terContext) {
	s._block_case_ter = v
}

func (s *List_expre_case_terContext) GetE() []IBlock_case_terContext { return s.e }

func (s *List_expre_case_terContext) SetE(v []IBlock_case_terContext) { s.e = v }

func (s *List_expre_case_terContext) GetL() *arrayList.List { return s.l }

func (s *List_expre_case_terContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_expre_case_terContext) AllBlock_case_ter() []IBlock_case_terContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlock_case_terContext)(nil)).Elem())
	var tst = make([]IBlock_case_terContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlock_case_terContext)
		}
	}

	return tst
}

func (s *List_expre_case_terContext) Block_case_ter(i int) IBlock_case_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_case_terContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlock_case_terContext)
}

func (s *List_expre_case_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expre_case_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expre_case_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_expre_case_ter(s)
	}
}

func (s *List_expre_case_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_expre_case_ter(s)
	}
}

func (p *Chems) List_expre_case_ter() (localctx IList_expre_case_terContext) {
	localctx = NewList_expre_case_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ChemsRULE_list_expre_case_ter)

	localctx.(*List_expre_case_terContext).l = arrayList.New()

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
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(ChemsTK_MENOS-47))|(1<<(ChemsTK_PARA-47))|(1<<(ChemsTK_NOT-47)))) != 0) {
		{
			p.SetState(378)

			var _x = p.Block_case_ter()

			localctx.(*List_expre_case_terContext)._block_case_ter = _x
		}
		localctx.(*List_expre_case_terContext).e = append(localctx.(*List_expre_case_terContext).e, localctx.(*List_expre_case_terContext)._block_case_ter)

		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_expre_case_terContext).GetE()
	for _, e := range listInt {
		localctx.(*List_expre_case_terContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IBlock_case_terContext is an interface to support dynamic dispatch.
type IBlock_case_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Expresion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Expresion)

	// IsBlock_case_terContext differentiates from other interfaces.
	IsBlock_case_terContext()
}

type Block_case_terContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Expresion
	_expression IExpressionContext
}

func NewEmptyBlock_case_terContext() *Block_case_terContext {
	var p = new(Block_case_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_case_ter
	return p
}

func (*Block_case_terContext) IsBlock_case_terContext() {}

func NewBlock_case_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_case_terContext {
	var p = new(Block_case_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_case_ter

	return p
}

func (s *Block_case_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_case_terContext) Get_expression() IExpressionContext { return s._expression }

func (s *Block_case_terContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Block_case_terContext) GetInstr() interfaces.Expresion { return s.instr }

func (s *Block_case_terContext) SetInstr(v interfaces.Expresion) { s.instr = v }

func (s *Block_case_terContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Block_case_terContext) TK_BARRA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_BARRA, 0)
}

func (s *Block_case_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_case_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_case_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_case_ter(s)
	}
}

func (s *Block_case_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_case_ter(s)
	}
}

func (p *Chems) Block_case_ter() (localctx IBlock_case_terContext) {
	localctx = NewBlock_case_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ChemsRULE_block_case_ter)

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

	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(385)

			var _x = p.Expression()

			localctx.(*Block_case_terContext)._expression = _x
		}
		{
			p.SetState(386)
			p.Match(ChemsTK_BARRA)
		}
		localctx.(*Block_case_terContext).instr = control.NewCaseTer(localctx.(*Block_case_terContext).Get_expression().GetP(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(389)

			var _x = p.Expression()

			localctx.(*Block_case_terContext)._expression = _x
		}
		localctx.(*Block_case_terContext).instr = control.NewCaseTer(localctx.(*Block_case_terContext).Get_expression().GetP(), nil, nil)

	}

	return localctx
}

// IInstr_default_terContext is an interface to support dynamic dispatch.
type IInstr_default_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Expresion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Expresion)

	// IsInstr_default_terContext differentiates from other interfaces.
	IsInstr_default_terContext()
}

type Instr_default_terContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Expresion
	_expression IExpressionContext
}

func NewEmptyInstr_default_terContext() *Instr_default_terContext {
	var p = new(Instr_default_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_default_ter
	return p
}

func (*Instr_default_terContext) IsInstr_default_terContext() {}

func NewInstr_default_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_default_terContext {
	var p = new(Instr_default_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_default_ter

	return p
}

func (s *Instr_default_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_default_terContext) Get_expression() IExpressionContext { return s._expression }

func (s *Instr_default_terContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Instr_default_terContext) GetInstr() interfaces.Expresion { return s.instr }

func (s *Instr_default_terContext) SetInstr(v interfaces.Expresion) { s.instr = v }

func (s *Instr_default_terContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_default_terContext) TK_IGUALMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALMAYOR, 0)
}

func (s *Instr_default_terContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instr_default_terContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_default_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_default_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_default_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_default_ter(s)
	}
}

func (s *Instr_default_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_default_ter(s)
	}
}

func (p *Chems) Instr_default_ter() (localctx IInstr_default_terContext) {
	localctx = NewInstr_default_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ChemsRULE_instr_default_ter)

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
		p.SetState(394)
		p.Match(ChemsID)
	}
	{
		p.SetState(395)
		p.Match(ChemsTK_IGUALMAYOR)
	}
	{
		p.SetState(396)

		var _x = p.Expression()

		localctx.(*Instr_default_terContext)._expression = _x
	}
	{
		p.SetState(397)
		p.Match(ChemsTK_COMA)
	}
	localctx.(*Instr_default_terContext).instr = control.NewDefaultTer(localctx.(*Instr_default_terContext).Get_expression().GetP())

	return localctx
}

// IInstr_whileContext is an interface to support dynamic dispatch.
type IInstr_whileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_whileContext differentiates from other interfaces.
	IsInstr_whileContext()
}

type Instr_whileContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_expression    IExpressionContext
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_whileContext() *Instr_whileContext {
	var p = new(Instr_whileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_while
	return p
}

func (*Instr_whileContext) IsInstr_whileContext() {}

func NewInstr_whileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_whileContext {
	var p = new(Instr_whileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_while

	return p
}

func (s *Instr_whileContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_whileContext) Get_expression() IExpressionContext { return s._expression }

func (s *Instr_whileContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_whileContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Instr_whileContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_whileContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_whileContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_whileContext) R_WHILE() antlr.TerminalNode {
	return s.GetToken(ChemsR_WHILE, 0)
}

func (s *Instr_whileContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instr_whileContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_whileContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_whileContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_whileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_whileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_whileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_while(s)
	}
}

func (s *Instr_whileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_while(s)
	}
}

func (p *Chems) Instr_while() (localctx IInstr_whileContext) {
	localctx = NewInstr_whileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ChemsRULE_instr_while)

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
		p.SetState(400)
		p.Match(ChemsR_WHILE)
	}
	{
		p.SetState(401)

		var _x = p.Expression()

		localctx.(*Instr_whileContext)._expression = _x
	}
	{
		p.SetState(402)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(403)

		var _x = p.Instrucciones()

		localctx.(*Instr_whileContext)._instrucciones = _x
	}
	{
		p.SetState(404)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_whileContext).instr = loops.NewWhile(localctx.(*Instr_whileContext).Get_expression().GetP(), localctx.(*Instr_whileContext).Get_instrucciones().GetL())

	return localctx
}

// IInstr_loopContext is an interface to support dynamic dispatch.
type IInstr_loopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_loopContext differentiates from other interfaces.
	IsInstr_loopContext()
}

type Instr_loopContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_loopContext() *Instr_loopContext {
	var p = new(Instr_loopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_loop
	return p
}

func (*Instr_loopContext) IsInstr_loopContext() {}

func NewInstr_loopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_loopContext {
	var p = new(Instr_loopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_loop

	return p
}

func (s *Instr_loopContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_loopContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_loopContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_loopContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_loopContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_loopContext) R_LOOP() antlr.TerminalNode {
	return s.GetToken(ChemsR_LOOP, 0)
}

func (s *Instr_loopContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_loopContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_loopContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_loopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_loopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_loopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_loop(s)
	}
}

func (s *Instr_loopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_loop(s)
	}
}

func (p *Chems) Instr_loop() (localctx IInstr_loopContext) {
	localctx = NewInstr_loopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ChemsRULE_instr_loop)

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
		p.SetState(407)
		p.Match(ChemsR_LOOP)
	}
	{
		p.SetState(408)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(409)

		var _x = p.Instrucciones()

		localctx.(*Instr_loopContext)._instrucciones = _x
	}
	{
		p.SetState(410)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_loopContext).instr = loops.NewLoop(localctx.(*Instr_loopContext).Get_instrucciones().GetL())

	return localctx
}

// IInstr_loop_ternarioContext is an interface to support dynamic dispatch.
type IInstr_loop_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Expresion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Expresion)

	// IsInstr_loop_ternarioContext differentiates from other interfaces.
	IsInstr_loop_ternarioContext()
}

type Instr_loop_ternarioContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Expresion
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_loop_ternarioContext() *Instr_loop_ternarioContext {
	var p = new(Instr_loop_ternarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_loop_ternario
	return p
}

func (*Instr_loop_ternarioContext) IsInstr_loop_ternarioContext() {}

func NewInstr_loop_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_loop_ternarioContext {
	var p = new(Instr_loop_ternarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_loop_ternario

	return p
}

func (s *Instr_loop_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_loop_ternarioContext) Get_instrucciones() IInstruccionesContext {
	return s._instrucciones
}

func (s *Instr_loop_ternarioContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_loop_ternarioContext) GetInstr() interfaces.Expresion { return s.instr }

func (s *Instr_loop_ternarioContext) SetInstr(v interfaces.Expresion) { s.instr = v }

func (s *Instr_loop_ternarioContext) R_LOOP() antlr.TerminalNode {
	return s.GetToken(ChemsR_LOOP, 0)
}

func (s *Instr_loop_ternarioContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_loop_ternarioContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_loop_ternarioContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_loop_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_loop_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_loop_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_loop_ternario(s)
	}
}

func (s *Instr_loop_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_loop_ternario(s)
	}
}

func (p *Chems) Instr_loop_ternario() (localctx IInstr_loop_ternarioContext) {
	localctx = NewInstr_loop_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ChemsRULE_instr_loop_ternario)

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
		p.SetState(413)
		p.Match(ChemsR_LOOP)
	}
	{
		p.SetState(414)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(415)

		var _x = p.Instrucciones()

		localctx.(*Instr_loop_ternarioContext)._instrucciones = _x
	}
	{
		p.SetState(416)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_loop_ternarioContext).instr = loops.NewLoopTernario(localctx.(*Instr_loop_ternarioContext).Get_instrucciones().GetL())

	return localctx
}

// IInstr_for_inContext is an interface to support dynamic dispatch.
type IInstr_for_inContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_FOR returns the _R_FOR token.
	Get_R_FOR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_R_FOR sets the _R_FOR token.
	Set_R_FOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpressionContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_for_inContext differentiates from other interfaces.
	IsInstr_for_inContext()
}

type Instr_for_inContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_FOR         antlr.Token
	_ID            antlr.Token
	left           IExpressionContext
	right          IExpressionContext
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_for_inContext() *Instr_for_inContext {
	var p = new(Instr_for_inContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_for_in
	return p
}

func (*Instr_for_inContext) IsInstr_for_inContext() {}

func NewInstr_for_inContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_for_inContext {
	var p = new(Instr_for_inContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_for_in

	return p
}

func (s *Instr_for_inContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_for_inContext) Get_R_FOR() antlr.Token { return s._R_FOR }

func (s *Instr_for_inContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_for_inContext) Set_R_FOR(v antlr.Token) { s._R_FOR = v }

func (s *Instr_for_inContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_for_inContext) GetLeft() IExpressionContext { return s.left }

func (s *Instr_for_inContext) GetRight() IExpressionContext { return s.right }

func (s *Instr_for_inContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_for_inContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *Instr_for_inContext) SetRight(v IExpressionContext) { s.right = v }

func (s *Instr_for_inContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_for_inContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_for_inContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_for_inContext) R_FOR() antlr.TerminalNode {
	return s.GetToken(ChemsR_FOR, 0)
}

func (s *Instr_for_inContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_for_inContext) R_IN() antlr.TerminalNode {
	return s.GetToken(ChemsR_IN, 0)
}

func (s *Instr_for_inContext) TK_DOBLEPUNTO() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DOBLEPUNTO, 0)
}

func (s *Instr_for_inContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_for_inContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_for_inContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_for_inContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Instr_for_inContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instr_for_inContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_for_inContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_for_inContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_for_in(s)
	}
}

func (s *Instr_for_inContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_for_in(s)
	}
}

func (p *Chems) Instr_for_in() (localctx IInstr_for_inContext) {
	localctx = NewInstr_for_inContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ChemsRULE_instr_for_in)

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

	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(419)

			var _m = p.Match(ChemsR_FOR)

			localctx.(*Instr_for_inContext)._R_FOR = _m
		}
		{
			p.SetState(420)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_for_inContext)._ID = _m
		}
		{
			p.SetState(421)
			p.Match(ChemsR_IN)
		}
		{
			p.SetState(422)

			var _x = p.Expression()

			localctx.(*Instr_for_inContext).left = _x
		}
		{
			p.SetState(423)
			p.Match(ChemsTK_DOBLEPUNTO)
		}
		{
			p.SetState(424)

			var _x = p.Expression()

			localctx.(*Instr_for_inContext).right = _x
		}
		{
			p.SetState(425)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(426)

			var _x = p.Instrucciones()

			localctx.(*Instr_for_inContext)._instrucciones = _x
		}
		{
			p.SetState(427)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_for_inContext).instr = loops.NewFor((func() string {
			if localctx.(*Instr_for_inContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_for_inContext).Get_ID().GetText()
			}
		}()), interfaces.INTEGER, localctx.(*Instr_for_inContext).GetLeft().GetP(), localctx.(*Instr_for_inContext).GetRight().GetP(), localctx.(*Instr_for_inContext).Get_instrucciones().GetL(), nil, (func() int {
			if localctx.(*Instr_for_inContext).Get_R_FOR() == nil {
				return 0
			} else {
				return localctx.(*Instr_for_inContext).Get_R_FOR().GetLine()
			}
		}()), localctx.(*Instr_for_inContext).Get_R_FOR().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(430)

			var _m = p.Match(ChemsR_FOR)

			localctx.(*Instr_for_inContext)._R_FOR = _m
		}
		{
			p.SetState(431)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_for_inContext)._ID = _m
		}
		{
			p.SetState(432)
			p.Match(ChemsR_IN)
		}
		{
			p.SetState(433)

			var _x = p.Expression()

			localctx.(*Instr_for_inContext).left = _x
		}
		{
			p.SetState(434)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(435)

			var _x = p.Instrucciones()

			localctx.(*Instr_for_inContext)._instrucciones = _x
		}
		{
			p.SetState(436)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_for_inContext).instr = loops.NewFor((func() string {
			if localctx.(*Instr_for_inContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_for_inContext).Get_ID().GetText()
			}
		}()), interfaces.STRING, localctx.(*Instr_for_inContext).GetLeft().GetP(), nil, localctx.(*Instr_for_inContext).Get_instrucciones().GetL(), nil, (func() int {
			if localctx.(*Instr_for_inContext).Get_R_FOR() == nil {
				return 0
			} else {
				return localctx.(*Instr_for_inContext).Get_R_FOR().GetLine()
			}
		}()), localctx.(*Instr_for_inContext).Get_R_FOR().GetColumn())

	}

	return localctx
}

// IInstr_breakContext is an interface to support dynamic dispatch.
type IInstr_breakContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_BREAK returns the _R_BREAK token.
	Get_R_BREAK() antlr.Token

	// Set_R_BREAK sets the _R_BREAK token.
	Set_R_BREAK(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_breakContext differentiates from other interfaces.
	IsInstr_breakContext()
}

type Instr_breakContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruction
	_R_BREAK    antlr.Token
	_expression IExpressionContext
}

func NewEmptyInstr_breakContext() *Instr_breakContext {
	var p = new(Instr_breakContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_break
	return p
}

func (*Instr_breakContext) IsInstr_breakContext() {}

func NewInstr_breakContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_breakContext {
	var p = new(Instr_breakContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_break

	return p
}

func (s *Instr_breakContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_breakContext) Get_R_BREAK() antlr.Token { return s._R_BREAK }

func (s *Instr_breakContext) Set_R_BREAK(v antlr.Token) { s._R_BREAK = v }

func (s *Instr_breakContext) Get_expression() IExpressionContext { return s._expression }

func (s *Instr_breakContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Instr_breakContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_breakContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_breakContext) R_BREAK() antlr.TerminalNode {
	return s.GetToken(ChemsR_BREAK, 0)
}

func (s *Instr_breakContext) End_instr() IEnd_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnd_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnd_instrContext)
}

func (s *Instr_breakContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instr_breakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_breakContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_breakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_break(s)
	}
}

func (s *Instr_breakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_break(s)
	}
}

func (p *Chems) Instr_break() (localctx IInstr_breakContext) {
	localctx = NewInstr_breakContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ChemsRULE_instr_break)

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

	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(441)

			var _m = p.Match(ChemsR_BREAK)

			localctx.(*Instr_breakContext)._R_BREAK = _m
		}
		{
			p.SetState(442)
			p.End_instr()
		}
		localctx.(*Instr_breakContext).instr = transferencia.NewBreak(nil, (func() int {
			if localctx.(*Instr_breakContext).Get_R_BREAK() == nil {
				return 0
			} else {
				return localctx.(*Instr_breakContext).Get_R_BREAK().GetLine()
			}
		}()), localctx.(*Instr_breakContext).Get_R_BREAK().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(445)

			var _m = p.Match(ChemsR_BREAK)

			localctx.(*Instr_breakContext)._R_BREAK = _m
		}
		{
			p.SetState(446)

			var _x = p.Expression()

			localctx.(*Instr_breakContext)._expression = _x
		}
		{
			p.SetState(447)
			p.End_instr()
		}
		localctx.(*Instr_breakContext).instr = transferencia.NewBreak(localctx.(*Instr_breakContext).Get_expression().GetP(), (func() int {
			if localctx.(*Instr_breakContext).Get_R_BREAK() == nil {
				return 0
			} else {
				return localctx.(*Instr_breakContext).Get_R_BREAK().GetLine()
			}
		}()), localctx.(*Instr_breakContext).Get_R_BREAK().GetColumn())

	}

	return localctx
}

// IInstr_continueContext is an interface to support dynamic dispatch.
type IInstr_continueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_CONTINUE returns the _R_CONTINUE token.
	Get_R_CONTINUE() antlr.Token

	// Set_R_CONTINUE sets the _R_CONTINUE token.
	Set_R_CONTINUE(antlr.Token)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_continueContext differentiates from other interfaces.
	IsInstr_continueContext()
}

type Instr_continueContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruction
	_R_CONTINUE antlr.Token
}

func NewEmptyInstr_continueContext() *Instr_continueContext {
	var p = new(Instr_continueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_continue
	return p
}

func (*Instr_continueContext) IsInstr_continueContext() {}

func NewInstr_continueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_continueContext {
	var p = new(Instr_continueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_continue

	return p
}

func (s *Instr_continueContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_continueContext) Get_R_CONTINUE() antlr.Token { return s._R_CONTINUE }

func (s *Instr_continueContext) Set_R_CONTINUE(v antlr.Token) { s._R_CONTINUE = v }

func (s *Instr_continueContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_continueContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_continueContext) R_CONTINUE() antlr.TerminalNode {
	return s.GetToken(ChemsR_CONTINUE, 0)
}

func (s *Instr_continueContext) End_instr() IEnd_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnd_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnd_instrContext)
}

func (s *Instr_continueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_continueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_continueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_continue(s)
	}
}

func (s *Instr_continueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_continue(s)
	}
}

func (p *Chems) Instr_continue() (localctx IInstr_continueContext) {
	localctx = NewInstr_continueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ChemsRULE_instr_continue)

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
		p.SetState(452)

		var _m = p.Match(ChemsR_CONTINUE)

		localctx.(*Instr_continueContext)._R_CONTINUE = _m
	}
	{
		p.SetState(453)
		p.End_instr()
	}
	localctx.(*Instr_continueContext).instr = transferencia.NewContinue((func() int {
		if localctx.(*Instr_continueContext).Get_R_CONTINUE() == nil {
			return 0
		} else {
			return localctx.(*Instr_continueContext).Get_R_CONTINUE().GetLine()
		}
	}()), localctx.(*Instr_continueContext).Get_R_CONTINUE().GetColumn())

	return localctx
}

// IInstr_returnContext is an interface to support dynamic dispatch.
type IInstr_returnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_RETURN returns the _R_RETURN token.
	Get_R_RETURN() antlr.Token

	// Set_R_RETURN sets the _R_RETURN token.
	Set_R_RETURN(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_returnContext differentiates from other interfaces.
	IsInstr_returnContext()
}

type Instr_returnContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruction
	_R_RETURN   antlr.Token
	_expression IExpressionContext
}

func NewEmptyInstr_returnContext() *Instr_returnContext {
	var p = new(Instr_returnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_return
	return p
}

func (*Instr_returnContext) IsInstr_returnContext() {}

func NewInstr_returnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_returnContext {
	var p = new(Instr_returnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_return

	return p
}

func (s *Instr_returnContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_returnContext) Get_R_RETURN() antlr.Token { return s._R_RETURN }

func (s *Instr_returnContext) Set_R_RETURN(v antlr.Token) { s._R_RETURN = v }

func (s *Instr_returnContext) Get_expression() IExpressionContext { return s._expression }

func (s *Instr_returnContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Instr_returnContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_returnContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_returnContext) R_RETURN() antlr.TerminalNode {
	return s.GetToken(ChemsR_RETURN, 0)
}

func (s *Instr_returnContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instr_returnContext) End_instr() IEnd_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnd_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnd_instrContext)
}

func (s *Instr_returnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_returnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_returnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_return(s)
	}
}

func (s *Instr_returnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_return(s)
	}
}

func (p *Chems) Instr_return() (localctx IInstr_returnContext) {
	localctx = NewInstr_returnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ChemsRULE_instr_return)

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
		p.SetState(456)

		var _m = p.Match(ChemsR_RETURN)

		localctx.(*Instr_returnContext)._R_RETURN = _m
	}
	{
		p.SetState(457)

		var _x = p.Expression()

		localctx.(*Instr_returnContext)._expression = _x
	}
	{
		p.SetState(458)
		p.End_instr()
	}
	localctx.(*Instr_returnContext).instr = transferencia.NewReturn(localctx.(*Instr_returnContext).Get_expression().GetP(), (func() int {
		if localctx.(*Instr_returnContext).Get_R_RETURN() == nil {
			return 0
		} else {
			return localctx.(*Instr_returnContext).Get_R_RETURN().GetLine()
		}
	}()), localctx.(*Instr_returnContext).Get_R_RETURN().GetColumn())

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

func (s *Instr_tipoContext) R_STR() antlr.TerminalNode {
	return s.GetToken(ChemsR_STR, 0)
}

func (s *Instr_tipoContext) R_BOOL() antlr.TerminalNode {
	return s.GetToken(ChemsR_BOOL, 0)
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
	p.EnterRule(localctx, 66, ChemsRULE_instr_tipo)

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

	p.SetState(471)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(461)
			p.Match(ChemsR_INT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.INTEGER

	case ChemsR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(463)
			p.Match(ChemsR_FLOAT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.FLOAT

	case ChemsR_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(465)
			p.Match(ChemsR_STRING)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STRING

	case ChemsR_STR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(467)
			p.Match(ChemsR_STR)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STRING

	case ChemsR_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(469)
			p.Match(ChemsR_BOOL)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.BOOLEAN

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
	p.EnterRule(localctx, 68, ChemsRULE_expression)

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
		p.SetState(473)

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

func (s *Exp_aritContext) TK_IGUALIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALIGUAL, 0)
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
	_startState := 70
	p.EnterRecursionRule(localctx, 70, ChemsRULE_exp_arit, _p)
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
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(477)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(478)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(479)

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
			p.SetState(480)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(481)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(ChemsTK_MULT-43))|(1<<(ChemsTK_DIV-43))|(1<<(ChemsTK_MODULO-43)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(482)

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
		}()), "-1", (func() int {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetLine()
			}
		}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

	case 2:
		{
			p.SetState(485)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(486)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(487)

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
			p.SetState(488)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(489)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(ChemsTK_MULT-43))|(1<<(ChemsTK_DIV-43))|(1<<(ChemsTK_MODULO-43)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(490)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(491)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(492)

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
			p.SetState(493)
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
		}()), (func() int {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetLine()
			}
		}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

	case 3:
		{
			p.SetState(496)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(497)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(498)

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
			p.SetState(499)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(500)

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
			p.SetState(501)

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
		}()), "-1", (func() int {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetLine()
			}
		}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

	case 4:
		{
			p.SetState(504)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(505)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(506)

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
			p.SetState(507)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(508)

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
			p.SetState(509)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(510)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(511)

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
			p.SetState(512)
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
		}()), (func() int {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetLine()
			}
		}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

	case 5:
		{
			p.SetState(515)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(516)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(517)

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
			p.SetState(518)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(519)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(ChemsTK_IGUALIGUAL-36))|(1<<(ChemsTK_MAYORIGUAL-36))|(1<<(ChemsTK_MENORIGUAL-36))|(1<<(ChemsTK_DIFIGUAL-36))|(1<<(ChemsTK_MAYOR-36))|(1<<(ChemsTK_MENOR-36)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(520)

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
		}()), "-1", (func() int {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetLine()
			}
		}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

	case 6:
		{
			p.SetState(523)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(524)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(525)

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
			p.SetState(526)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(527)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(ChemsTK_IGUALIGUAL-36))|(1<<(ChemsTK_MAYORIGUAL-36))|(1<<(ChemsTK_MENORIGUAL-36))|(1<<(ChemsTK_DIFIGUAL-36))|(1<<(ChemsTK_MAYOR-36))|(1<<(ChemsTK_MENOR-36)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(528)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(529)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(530)

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
			p.SetState(531)
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
		}()), (func() int {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetLine()
			}
		}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

	case 7:
		{
			p.SetState(534)

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
			p.SetState(535)

			var _x = p.Expression()

			localctx.(*Exp_aritContext)._expression = _x
		}
		localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).Get_expression().GetP(), (func() string {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetText()
			}
		}()), nil, true, "-1", "-1", (func() int {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetLine()
			}
		}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

	case 8:
		{
			p.SetState(538)

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
			p.SetState(539)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(540)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(541)

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
			p.SetState(542)
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
		}()), "-1", (func() int {
			if localctx.(*Exp_aritContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Exp_aritContext).GetOp().GetLine()
			}
		}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

	case 9:
		{
			p.SetState(545)

			var _x = p.Primitivo()

			localctx.(*Exp_aritContext)._primitivo = _x
		}
		localctx.(*Exp_aritContext).p = localctx.(*Exp_aritContext).Get_primitivo().GetP()

	case 10:
		{
			p.SetState(548)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(549)

			var _x = p.Expression()

			localctx.(*Exp_aritContext)._expression = _x
		}
		{
			p.SetState(550)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Exp_aritContext).p = localctx.(*Exp_aritContext).Get_expression().GetP()

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(601)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(599)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(555)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(556)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(ChemsTK_MULT-43))|(1<<(ChemsTK_DIV-43))|(1<<(ChemsTK_MODULO-43)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(557)

					var _x = p.exp_arit(18)

					localctx.(*Exp_aritContext).right = _x
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, "-1", "-1", (func() int {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

			case 2:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(560)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(561)

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
					p.SetState(562)

					var _x = p.exp_arit(14)

					localctx.(*Exp_aritContext).right = _x
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, "-1", "-1", (func() int {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

			case 3:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(565)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(566)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(ChemsTK_IGUALIGUAL-36))|(1<<(ChemsTK_MAYORIGUAL-36))|(1<<(ChemsTK_MENORIGUAL-36))|(1<<(ChemsTK_DIFIGUAL-36))|(1<<(ChemsTK_MAYOR-36))|(1<<(ChemsTK_MENOR-36)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(567)

					var _x = p.exp_arit(10)

					localctx.(*Exp_aritContext).right = _x
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, "-1", "-1", (func() int {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

			case 4:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(570)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(571)

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
					p.SetState(572)

					var _x = p.exp_arit(6)

					localctx.(*Exp_aritContext).right = _x
				}
				localctx.(*Exp_aritContext).p = expresion.NewOperacion(localctx.(*Exp_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Exp_aritContext).GetRight().GetP(), false, "-1", "-1", (func() int {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

			case 5:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(575)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(576)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(ChemsTK_MULT-43))|(1<<(ChemsTK_DIV-43))|(1<<(ChemsTK_MODULO-43)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(577)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(578)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(579)

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
					p.SetState(580)
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
				}()), (func() int {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

			case 6:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(583)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(584)

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
					p.SetState(585)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(586)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(587)

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
					p.SetState(588)
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
				}()), (func() int {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

			case 7:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(591)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(592)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(ChemsTK_IGUALIGUAL-36))|(1<<(ChemsTK_MAYORIGUAL-36))|(1<<(ChemsTK_MENORIGUAL-36))|(1<<(ChemsTK_DIFIGUAL-36))|(1<<(ChemsTK_MAYOR-36))|(1<<(ChemsTK_MENOR-36)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(593)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(594)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(595)

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
					p.SetState(596)
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
				}()), (func() int {
					if localctx.(*Exp_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Exp_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Exp_aritContext).GetOp().GetColumn())

			}

		}
		p.SetState(603)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
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

	// Get_instr_ternario returns the _instr_ternario rule contexts.
	Get_instr_ternario() IInstr_ternarioContext

	// Get_instr_match_ter returns the _instr_match_ter rule contexts.
	Get_instr_match_ter() IInstr_match_terContext

	// Get_instr_loop_ternario returns the _instr_loop_ternario rule contexts.
	Get_instr_loop_ternario() IInstr_loop_ternarioContext

	// Set_instr_ternario sets the _instr_ternario rule contexts.
	Set_instr_ternario(IInstr_ternarioContext)

	// Set_instr_match_ter sets the _instr_match_ter rule contexts.
	Set_instr_match_ter(IInstr_match_terContext)

	// Set_instr_loop_ternario sets the _instr_loop_ternario rule contexts.
	Set_instr_loop_ternario(IInstr_loop_ternarioContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser               antlr.Parser
	p                    interfaces.Expresion
	_NUMBER              antlr.Token
	_DOUBLE              antlr.Token
	_STRING              antlr.Token
	_BOOLEAN             antlr.Token
	_ID                  antlr.Token
	_instr_ternario      IInstr_ternarioContext
	_instr_match_ter     IInstr_match_terContext
	_instr_loop_ternario IInstr_loop_ternarioContext
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

func (s *PrimitivoContext) Get_instr_ternario() IInstr_ternarioContext { return s._instr_ternario }

func (s *PrimitivoContext) Get_instr_match_ter() IInstr_match_terContext { return s._instr_match_ter }

func (s *PrimitivoContext) Get_instr_loop_ternario() IInstr_loop_ternarioContext {
	return s._instr_loop_ternario
}

func (s *PrimitivoContext) Set_instr_ternario(v IInstr_ternarioContext) { s._instr_ternario = v }

func (s *PrimitivoContext) Set_instr_match_ter(v IInstr_match_terContext) { s._instr_match_ter = v }

func (s *PrimitivoContext) Set_instr_loop_ternario(v IInstr_loop_ternarioContext) {
	s._instr_loop_ternario = v
}

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

func (s *PrimitivoContext) Instr_ternario() IInstr_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_ternarioContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_ternarioContext)
}

func (s *PrimitivoContext) Instr_match_ter() IInstr_match_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_match_terContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_match_terContext)
}

func (s *PrimitivoContext) Instr_loop_ternario() IInstr_loop_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_loop_ternarioContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_loop_ternarioContext)
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
	p.EnterRule(localctx, 72, ChemsRULE_primitivo)

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

	p.SetState(623)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(604)

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

		localctx.(*PrimitivoContext).p = expresion.PRIMITIVO(num, interfaces.INTEGER, (func() int {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())

	case ChemsDOUBLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(606)

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
		localctx.(*PrimitivoContext).p = expresion.PRIMITIVO(num, interfaces.FLOAT, (func() int {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())

	case ChemsSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(608)

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
		localctx.(*PrimitivoContext).p = expresion.PRIMITIVO(str, interfaces.STRING, (func() int {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_STRING().GetColumn())

	case ChemsBOOLEAN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(610)

			var _m = p.Match(ChemsBOOLEAN)

			localctx.(*PrimitivoContext)._BOOLEAN = _m
		}

		// str:= (func() string { if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText() }}())[1:len((func() string { if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText() }}()))-1]
		exp, _ := strconv.ParseBool((func() string {
			if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText()
			}
		}()))
		localctx.(*PrimitivoContext).p = expresion.PRIMITIVO(exp, interfaces.BOOLEAN, (func() int {
			if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEAN().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_BOOLEAN().GetColumn())

	case ChemsID:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(612)

			var _m = p.Match(ChemsID)

			localctx.(*PrimitivoContext)._ID = _m
		}
		localctx.(*PrimitivoContext).p = variable.NewIdentifier((func() string {
			if localctx.(*PrimitivoContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_ID().GetText()
			}
		}()), (func() int {
			if localctx.(*PrimitivoContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_ID().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_ID().GetColumn())

	case ChemsR_IF:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(614)

			var _x = p.Instr_ternario()

			localctx.(*PrimitivoContext)._instr_ternario = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_ternario().GetP()

	case ChemsR_MATCH:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(617)

			var _x = p.Instr_match_ter()

			localctx.(*PrimitivoContext)._instr_match_ter = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_match_ter().GetInstr()

	case ChemsR_LOOP:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(620)

			var _x = p.Instr_loop_ternario()

			localctx.(*PrimitivoContext)._instr_loop_ternario = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_loop_ternario().GetInstr()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 35:
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

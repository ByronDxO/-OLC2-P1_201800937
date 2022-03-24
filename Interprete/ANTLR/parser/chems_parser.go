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
import "OLC2/Interprete/instruction/function"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 64, 727,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 3, 2, 3, 2, 3, 2, 3, 3, 6, 3, 95, 10, 3, 13, 3, 14, 3, 96,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 104, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 150, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 177, 10, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 213, 10, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 5, 10, 248, 10, 10, 3, 11, 7, 11, 251, 10, 11, 12, 11, 14, 11, 254,
	11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12,
	285, 10, 12, 3, 13, 7, 13, 288, 10, 13, 12, 13, 14, 13, 291, 11, 13, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 310, 10, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 5, 15, 325, 10, 15, 3, 16, 3, 16, 3, 16, 3, 17, 6, 17, 331, 10,
	17, 13, 17, 14, 17, 332, 3, 17, 3, 17, 3, 18, 6, 18, 338, 10, 18, 13, 18,
	14, 18, 339, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 5, 19, 351, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 366, 10, 20, 3, 21, 6,
	21, 369, 10, 21, 13, 21, 14, 21, 370, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 5, 22, 390, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 24, 6, 24, 399, 10, 24, 13, 24, 14, 24, 400, 3, 24, 3, 24, 3, 25, 6,
	25, 406, 10, 25, 13, 25, 14, 25, 407, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 419, 10, 26, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5,
	31, 466, 10, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 5, 32, 477, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 5, 35, 510, 10, 35, 3, 36, 6, 36, 513, 10, 36,
	13, 36, 14, 36, 514, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 530, 10, 37, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40,
	554, 10, 40, 3, 41, 6, 41, 557, 10, 41, 13, 41, 14, 41, 558, 3, 41, 3,
	41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 570, 10, 42,
	3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 652, 10, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44,
	3, 44, 7, 44, 698, 10, 44, 12, 44, 14, 44, 701, 11, 44, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 5, 45, 725,
	10, 45, 3, 45, 2, 3, 86, 46, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 2, 9, 3, 2, 3,
	4, 3, 2, 6, 7, 3, 2, 47, 49, 3, 2, 50, 51, 4, 2, 39, 40, 43, 46, 4, 2,
	51, 51, 61, 61, 3, 2, 58, 59, 2, 753, 2, 90, 3, 2, 2, 2, 4, 94, 3, 2, 2,
	2, 6, 103, 3, 2, 2, 2, 8, 149, 3, 2, 2, 2, 10, 151, 3, 2, 2, 2, 12, 176,
	3, 2, 2, 2, 14, 212, 3, 2, 2, 2, 16, 214, 3, 2, 2, 2, 18, 247, 3, 2, 2,
	2, 20, 252, 3, 2, 2, 2, 22, 284, 3, 2, 2, 2, 24, 289, 3, 2, 2, 2, 26, 309,
	3, 2, 2, 2, 28, 324, 3, 2, 2, 2, 30, 326, 3, 2, 2, 2, 32, 330, 3, 2, 2,
	2, 34, 337, 3, 2, 2, 2, 36, 350, 3, 2, 2, 2, 38, 365, 3, 2, 2, 2, 40, 368,
	3, 2, 2, 2, 42, 389, 3, 2, 2, 2, 44, 391, 3, 2, 2, 2, 46, 398, 3, 2, 2,
	2, 48, 405, 3, 2, 2, 2, 50, 418, 3, 2, 2, 2, 52, 420, 3, 2, 2, 2, 54, 426,
	3, 2, 2, 2, 56, 433, 3, 2, 2, 2, 58, 439, 3, 2, 2, 2, 60, 465, 3, 2, 2,
	2, 62, 476, 3, 2, 2, 2, 64, 478, 3, 2, 2, 2, 66, 482, 3, 2, 2, 2, 68, 509,
	3, 2, 2, 2, 70, 512, 3, 2, 2, 2, 72, 529, 3, 2, 2, 2, 74, 531, 3, 2, 2,
	2, 76, 537, 3, 2, 2, 2, 78, 553, 3, 2, 2, 2, 80, 556, 3, 2, 2, 2, 82, 569,
	3, 2, 2, 2, 84, 571, 3, 2, 2, 2, 86, 651, 3, 2, 2, 2, 88, 724, 3, 2, 2,
	2, 90, 91, 5, 4, 3, 2, 91, 92, 8, 2, 1, 2, 92, 3, 3, 2, 2, 2, 93, 95, 5,
	8, 5, 2, 94, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96,
	97, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 8, 3, 1, 2, 99, 5, 3, 2, 2,
	2, 100, 101, 7, 35, 2, 2, 101, 104, 8, 4, 1, 2, 102, 104, 8, 4, 1, 2, 103,
	100, 3, 2, 2, 2, 103, 102, 3, 2, 2, 2, 104, 7, 3, 2, 2, 2, 105, 106, 5,
	12, 7, 2, 106, 107, 5, 6, 4, 2, 107, 108, 8, 5, 1, 2, 108, 150, 3, 2, 2,
	2, 109, 110, 5, 14, 8, 2, 110, 111, 8, 5, 1, 2, 111, 150, 3, 2, 2, 2, 112,
	113, 5, 16, 9, 2, 113, 114, 8, 5, 1, 2, 114, 150, 3, 2, 2, 2, 115, 116,
	5, 18, 10, 2, 116, 117, 8, 5, 1, 2, 117, 150, 3, 2, 2, 2, 118, 119, 5,
	26, 14, 2, 119, 120, 8, 5, 1, 2, 120, 150, 3, 2, 2, 2, 121, 122, 5, 54,
	28, 2, 122, 123, 8, 5, 1, 2, 123, 150, 3, 2, 2, 2, 124, 125, 5, 62, 32,
	2, 125, 126, 8, 5, 1, 2, 126, 150, 3, 2, 2, 2, 127, 128, 5, 64, 33, 2,
	128, 129, 8, 5, 1, 2, 129, 150, 3, 2, 2, 2, 130, 131, 5, 66, 34, 2, 131,
	132, 8, 5, 1, 2, 132, 150, 3, 2, 2, 2, 133, 134, 5, 56, 29, 2, 134, 135,
	8, 5, 1, 2, 135, 150, 3, 2, 2, 2, 136, 137, 5, 60, 31, 2, 137, 138, 8,
	5, 1, 2, 138, 150, 3, 2, 2, 2, 139, 140, 5, 10, 6, 2, 140, 141, 8, 5, 1,
	2, 141, 150, 3, 2, 2, 2, 142, 143, 5, 68, 35, 2, 143, 144, 8, 5, 1, 2,
	144, 150, 3, 2, 2, 2, 145, 146, 5, 74, 38, 2, 146, 147, 5, 6, 4, 2, 147,
	148, 8, 5, 1, 2, 148, 150, 3, 2, 2, 2, 149, 105, 3, 2, 2, 2, 149, 109,
	3, 2, 2, 2, 149, 112, 3, 2, 2, 2, 149, 115, 3, 2, 2, 2, 149, 118, 3, 2,
	2, 2, 149, 121, 3, 2, 2, 2, 149, 124, 3, 2, 2, 2, 149, 127, 3, 2, 2, 2,
	149, 130, 3, 2, 2, 2, 149, 133, 3, 2, 2, 2, 149, 136, 3, 2, 2, 2, 149,
	139, 3, 2, 2, 2, 149, 142, 3, 2, 2, 2, 149, 145, 3, 2, 2, 2, 150, 9, 3,
	2, 2, 2, 151, 152, 7, 21, 2, 2, 152, 153, 7, 20, 2, 2, 153, 154, 7, 52,
	2, 2, 154, 155, 7, 53, 2, 2, 155, 156, 7, 54, 2, 2, 156, 157, 5, 4, 3,
	2, 157, 158, 7, 55, 2, 2, 158, 159, 8, 6, 1, 2, 159, 11, 3, 2, 2, 2, 160,
	161, 9, 2, 2, 2, 161, 162, 7, 52, 2, 2, 162, 163, 5, 88, 45, 2, 163, 164,
	7, 53, 2, 2, 164, 165, 7, 35, 2, 2, 165, 166, 8, 7, 1, 2, 166, 177, 3,
	2, 2, 2, 167, 168, 9, 2, 2, 2, 168, 169, 7, 52, 2, 2, 169, 170, 7, 30,
	2, 2, 170, 171, 7, 36, 2, 2, 171, 172, 5, 80, 41, 2, 172, 173, 7, 53, 2,
	2, 173, 174, 7, 35, 2, 2, 174, 175, 8, 7, 1, 2, 175, 177, 3, 2, 2, 2, 176,
	160, 3, 2, 2, 2, 176, 167, 3, 2, 2, 2, 177, 13, 3, 2, 2, 2, 178, 179, 7,
	8, 2, 2, 179, 180, 7, 9, 2, 2, 180, 181, 7, 32, 2, 2, 181, 182, 7, 38,
	2, 2, 182, 183, 5, 84, 43, 2, 183, 184, 7, 35, 2, 2, 184, 185, 8, 8, 1,
	2, 185, 213, 3, 2, 2, 2, 186, 187, 7, 8, 2, 2, 187, 188, 7, 9, 2, 2, 188,
	189, 7, 32, 2, 2, 189, 190, 7, 37, 2, 2, 190, 191, 5, 78, 40, 2, 191, 192,
	7, 38, 2, 2, 192, 193, 5, 84, 43, 2, 193, 194, 7, 35, 2, 2, 194, 195, 8,
	8, 1, 2, 195, 213, 3, 2, 2, 2, 196, 197, 7, 8, 2, 2, 197, 198, 7, 32, 2,
	2, 198, 199, 7, 38, 2, 2, 199, 200, 5, 84, 43, 2, 200, 201, 7, 35, 2, 2,
	201, 202, 8, 8, 1, 2, 202, 213, 3, 2, 2, 2, 203, 204, 7, 8, 2, 2, 204,
	205, 7, 32, 2, 2, 205, 206, 7, 37, 2, 2, 206, 207, 5, 78, 40, 2, 207, 208,
	7, 38, 2, 2, 208, 209, 5, 84, 43, 2, 209, 210, 7, 35, 2, 2, 210, 211, 8,
	8, 1, 2, 211, 213, 3, 2, 2, 2, 212, 178, 3, 2, 2, 2, 212, 186, 3, 2, 2,
	2, 212, 196, 3, 2, 2, 2, 212, 203, 3, 2, 2, 2, 213, 15, 3, 2, 2, 2, 214,
	215, 7, 32, 2, 2, 215, 216, 7, 38, 2, 2, 216, 217, 5, 84, 43, 2, 217, 218,
	7, 35, 2, 2, 218, 219, 8, 9, 1, 2, 219, 17, 3, 2, 2, 2, 220, 221, 7, 10,
	2, 2, 221, 222, 5, 84, 43, 2, 222, 223, 7, 54, 2, 2, 223, 224, 5, 4, 3,
	2, 224, 225, 7, 55, 2, 2, 225, 226, 8, 10, 1, 2, 226, 248, 3, 2, 2, 2,
	227, 228, 7, 10, 2, 2, 228, 229, 5, 84, 43, 2, 229, 230, 7, 54, 2, 2, 230,
	231, 5, 4, 3, 2, 231, 232, 7, 55, 2, 2, 232, 233, 7, 11, 2, 2, 233, 234,
	7, 54, 2, 2, 234, 235, 5, 4, 3, 2, 235, 236, 7, 55, 2, 2, 236, 237, 8,
	10, 1, 2, 237, 248, 3, 2, 2, 2, 238, 239, 7, 10, 2, 2, 239, 240, 5, 84,
	43, 2, 240, 241, 7, 54, 2, 2, 241, 242, 5, 4, 3, 2, 242, 243, 7, 55, 2,
	2, 243, 244, 7, 11, 2, 2, 244, 245, 5, 20, 11, 2, 245, 246, 8, 10, 1, 2,
	246, 248, 3, 2, 2, 2, 247, 220, 3, 2, 2, 2, 247, 227, 3, 2, 2, 2, 247,
	238, 3, 2, 2, 2, 248, 19, 3, 2, 2, 2, 249, 251, 5, 18, 10, 2, 250, 249,
	3, 2, 2, 2, 251, 254, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2, 252, 253, 3, 2,
	2, 2, 253, 255, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 255, 256, 8, 11, 1, 2,
	256, 21, 3, 2, 2, 2, 257, 258, 7, 10, 2, 2, 258, 259, 5, 84, 43, 2, 259,
	260, 7, 54, 2, 2, 260, 261, 5, 84, 43, 2, 261, 262, 7, 55, 2, 2, 262, 263,
	8, 12, 1, 2, 263, 285, 3, 2, 2, 2, 264, 265, 7, 10, 2, 2, 265, 266, 5,
	84, 43, 2, 266, 267, 7, 54, 2, 2, 267, 268, 5, 84, 43, 2, 268, 269, 7,
	55, 2, 2, 269, 270, 7, 11, 2, 2, 270, 271, 7, 54, 2, 2, 271, 272, 5, 84,
	43, 2, 272, 273, 7, 55, 2, 2, 273, 274, 8, 12, 1, 2, 274, 285, 3, 2, 2,
	2, 275, 276, 7, 10, 2, 2, 276, 277, 5, 84, 43, 2, 277, 278, 7, 54, 2, 2,
	278, 279, 5, 84, 43, 2, 279, 280, 7, 55, 2, 2, 280, 281, 7, 11, 2, 2, 281,
	282, 5, 24, 13, 2, 282, 283, 8, 12, 1, 2, 283, 285, 3, 2, 2, 2, 284, 257,
	3, 2, 2, 2, 284, 264, 3, 2, 2, 2, 284, 275, 3, 2, 2, 2, 285, 23, 3, 2,
	2, 2, 286, 288, 5, 22, 12, 2, 287, 286, 3, 2, 2, 2, 288, 291, 3, 2, 2,
	2, 289, 287, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 292, 3, 2, 2, 2, 291,
	289, 3, 2, 2, 2, 292, 293, 8, 13, 1, 2, 293, 25, 3, 2, 2, 2, 294, 295,
	7, 12, 2, 2, 295, 296, 5, 84, 43, 2, 296, 297, 7, 54, 2, 2, 297, 298, 5,
	32, 17, 2, 298, 299, 5, 40, 21, 2, 299, 300, 7, 55, 2, 2, 300, 301, 8,
	14, 1, 2, 301, 310, 3, 2, 2, 2, 302, 303, 7, 12, 2, 2, 303, 304, 5, 84,
	43, 2, 304, 305, 7, 54, 2, 2, 305, 306, 5, 40, 21, 2, 306, 307, 7, 55,
	2, 2, 307, 308, 8, 14, 1, 2, 308, 310, 3, 2, 2, 2, 309, 294, 3, 2, 2, 2,
	309, 302, 3, 2, 2, 2, 310, 27, 3, 2, 2, 2, 311, 312, 5, 34, 18, 2, 312,
	313, 7, 41, 2, 2, 313, 314, 7, 54, 2, 2, 314, 315, 5, 4, 3, 2, 315, 316,
	7, 55, 2, 2, 316, 317, 8, 15, 1, 2, 317, 325, 3, 2, 2, 2, 318, 319, 5,
	34, 18, 2, 319, 320, 7, 41, 2, 2, 320, 321, 5, 30, 16, 2, 321, 322, 7,
	36, 2, 2, 322, 323, 8, 15, 1, 2, 323, 325, 3, 2, 2, 2, 324, 311, 3, 2,
	2, 2, 324, 318, 3, 2, 2, 2, 325, 29, 3, 2, 2, 2, 326, 327, 5, 8, 5, 2,
	327, 328, 8, 16, 1, 2, 328, 31, 3, 2, 2, 2, 329, 331, 5, 28, 15, 2, 330,
	329, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2, 332, 333,
	3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 335, 8, 17, 1, 2, 335, 33, 3, 2,
	2, 2, 336, 338, 5, 36, 19, 2, 337, 336, 3, 2, 2, 2, 338, 339, 3, 2, 2,
	2, 339, 337, 3, 2, 2, 2, 339, 340, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2, 341,
	342, 8, 18, 1, 2, 342, 35, 3, 2, 2, 2, 343, 344, 5, 84, 43, 2, 344, 345,
	7, 60, 2, 2, 345, 346, 8, 19, 1, 2, 346, 351, 3, 2, 2, 2, 347, 348, 5,
	84, 43, 2, 348, 349, 8, 19, 1, 2, 349, 351, 3, 2, 2, 2, 350, 343, 3, 2,
	2, 2, 350, 347, 3, 2, 2, 2, 351, 37, 3, 2, 2, 2, 352, 353, 7, 32, 2, 2,
	353, 354, 7, 41, 2, 2, 354, 355, 7, 54, 2, 2, 355, 356, 5, 4, 3, 2, 356,
	357, 7, 55, 2, 2, 357, 358, 8, 20, 1, 2, 358, 366, 3, 2, 2, 2, 359, 360,
	7, 32, 2, 2, 360, 361, 7, 41, 2, 2, 361, 362, 5, 30, 16, 2, 362, 363, 7,
	36, 2, 2, 363, 364, 8, 20, 1, 2, 364, 366, 3, 2, 2, 2, 365, 352, 3, 2,
	2, 2, 365, 359, 3, 2, 2, 2, 366, 39, 3, 2, 2, 2, 367, 369, 5, 38, 20, 2,
	368, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 368, 3, 2, 2, 2, 370,
	371, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 373, 8, 21, 1, 2, 373, 41,
	3, 2, 2, 2, 374, 375, 7, 12, 2, 2, 375, 376, 5, 84, 43, 2, 376, 377, 7,
	54, 2, 2, 377, 378, 5, 46, 24, 2, 378, 379, 5, 52, 27, 2, 379, 380, 7,
	55, 2, 2, 380, 381, 8, 22, 1, 2, 381, 390, 3, 2, 2, 2, 382, 383, 7, 12,
	2, 2, 383, 384, 5, 84, 43, 2, 384, 385, 7, 54, 2, 2, 385, 386, 5, 52, 27,
	2, 386, 387, 7, 55, 2, 2, 387, 388, 8, 22, 1, 2, 388, 390, 3, 2, 2, 2,
	389, 374, 3, 2, 2, 2, 389, 382, 3, 2, 2, 2, 390, 43, 3, 2, 2, 2, 391, 392,
	5, 48, 25, 2, 392, 393, 7, 41, 2, 2, 393, 394, 5, 84, 43, 2, 394, 395,
	7, 36, 2, 2, 395, 396, 8, 23, 1, 2, 396, 45, 3, 2, 2, 2, 397, 399, 5, 44,
	23, 2, 398, 397, 3, 2, 2, 2, 399, 400, 3, 2, 2, 2, 400, 398, 3, 2, 2, 2,
	400, 401, 3, 2, 2, 2, 401, 402, 3, 2, 2, 2, 402, 403, 8, 24, 1, 2, 403,
	47, 3, 2, 2, 2, 404, 406, 5, 50, 26, 2, 405, 404, 3, 2, 2, 2, 406, 407,
	3, 2, 2, 2, 407, 405, 3, 2, 2, 2, 407, 408, 3, 2, 2, 2, 408, 409, 3, 2,
	2, 2, 409, 410, 8, 25, 1, 2, 410, 49, 3, 2, 2, 2, 411, 412, 5, 84, 43,
	2, 412, 413, 7, 60, 2, 2, 413, 414, 8, 26, 1, 2, 414, 419, 3, 2, 2, 2,
	415, 416, 5, 84, 43, 2, 416, 417, 8, 26, 1, 2, 417, 419, 3, 2, 2, 2, 418,
	411, 3, 2, 2, 2, 418, 415, 3, 2, 2, 2, 419, 51, 3, 2, 2, 2, 420, 421, 7,
	32, 2, 2, 421, 422, 7, 41, 2, 2, 422, 423, 5, 84, 43, 2, 423, 424, 7, 36,
	2, 2, 424, 425, 8, 27, 1, 2, 425, 53, 3, 2, 2, 2, 426, 427, 7, 13, 2, 2,
	427, 428, 5, 84, 43, 2, 428, 429, 7, 54, 2, 2, 429, 430, 5, 4, 3, 2, 430,
	431, 7, 55, 2, 2, 431, 432, 8, 28, 1, 2, 432, 55, 3, 2, 2, 2, 433, 434,
	7, 17, 2, 2, 434, 435, 7, 54, 2, 2, 435, 436, 5, 4, 3, 2, 436, 437, 7,
	55, 2, 2, 437, 438, 8, 29, 1, 2, 438, 57, 3, 2, 2, 2, 439, 440, 7, 17,
	2, 2, 440, 441, 7, 54, 2, 2, 441, 442, 5, 4, 3, 2, 442, 443, 7, 55, 2,
	2, 443, 444, 8, 30, 1, 2, 444, 59, 3, 2, 2, 2, 445, 446, 7, 18, 2, 2, 446,
	447, 7, 32, 2, 2, 447, 448, 7, 19, 2, 2, 448, 449, 5, 84, 43, 2, 449, 450,
	7, 33, 2, 2, 450, 451, 5, 84, 43, 2, 451, 452, 7, 54, 2, 2, 452, 453, 5,
	4, 3, 2, 453, 454, 7, 55, 2, 2, 454, 455, 8, 31, 1, 2, 455, 466, 3, 2,
	2, 2, 456, 457, 7, 18, 2, 2, 457, 458, 7, 32, 2, 2, 458, 459, 7, 19, 2,
	2, 459, 460, 5, 84, 43, 2, 460, 461, 7, 54, 2, 2, 461, 462, 5, 4, 3, 2,
	462, 463, 7, 55, 2, 2, 463, 464, 8, 31, 1, 2, 464, 466, 3, 2, 2, 2, 465,
	445, 3, 2, 2, 2, 465, 456, 3, 2, 2, 2, 466, 61, 3, 2, 2, 2, 467, 468, 7,
	14, 2, 2, 468, 469, 5, 6, 4, 2, 469, 470, 8, 32, 1, 2, 470, 477, 3, 2,
	2, 2, 471, 472, 7, 14, 2, 2, 472, 473, 5, 84, 43, 2, 473, 474, 5, 6, 4,
	2, 474, 475, 8, 32, 1, 2, 475, 477, 3, 2, 2, 2, 476, 467, 3, 2, 2, 2, 476,
	471, 3, 2, 2, 2, 477, 63, 3, 2, 2, 2, 478, 479, 7, 15, 2, 2, 479, 480,
	5, 6, 4, 2, 480, 481, 8, 33, 1, 2, 481, 65, 3, 2, 2, 2, 482, 483, 7, 16,
	2, 2, 483, 484, 5, 84, 43, 2, 484, 485, 5, 6, 4, 2, 485, 486, 8, 34, 1,
	2, 486, 67, 3, 2, 2, 2, 487, 488, 7, 21, 2, 2, 488, 489, 7, 32, 2, 2, 489,
	490, 7, 52, 2, 2, 490, 491, 5, 70, 36, 2, 491, 492, 7, 53, 2, 2, 492, 493,
	7, 54, 2, 2, 493, 494, 5, 4, 3, 2, 494, 495, 7, 55, 2, 2, 495, 496, 8,
	35, 1, 2, 496, 510, 3, 2, 2, 2, 497, 498, 7, 21, 2, 2, 498, 499, 7, 32,
	2, 2, 499, 500, 7, 52, 2, 2, 500, 501, 5, 70, 36, 2, 501, 502, 7, 53, 2,
	2, 502, 503, 7, 42, 2, 2, 503, 504, 5, 78, 40, 2, 504, 505, 7, 54, 2, 2,
	505, 506, 5, 4, 3, 2, 506, 507, 7, 55, 2, 2, 507, 508, 8, 35, 1, 2, 508,
	510, 3, 2, 2, 2, 509, 487, 3, 2, 2, 2, 509, 497, 3, 2, 2, 2, 510, 69, 3,
	2, 2, 2, 511, 513, 5, 72, 37, 2, 512, 511, 3, 2, 2, 2, 513, 514, 3, 2,
	2, 2, 514, 512, 3, 2, 2, 2, 514, 515, 3, 2, 2, 2, 515, 516, 3, 2, 2, 2,
	516, 517, 8, 36, 1, 2, 517, 71, 3, 2, 2, 2, 518, 519, 7, 32, 2, 2, 519,
	520, 7, 37, 2, 2, 520, 521, 5, 78, 40, 2, 521, 522, 7, 36, 2, 2, 522, 523,
	8, 37, 1, 2, 523, 530, 3, 2, 2, 2, 524, 525, 7, 32, 2, 2, 525, 526, 7,
	37, 2, 2, 526, 527, 5, 78, 40, 2, 527, 528, 8, 37, 1, 2, 528, 530, 3, 2,
	2, 2, 529, 518, 3, 2, 2, 2, 529, 524, 3, 2, 2, 2, 530, 73, 3, 2, 2, 2,
	531, 532, 7, 32, 2, 2, 532, 533, 7, 52, 2, 2, 533, 534, 5, 80, 41, 2, 534,
	535, 7, 53, 2, 2, 535, 536, 8, 38, 1, 2, 536, 75, 3, 2, 2, 2, 537, 538,
	7, 32, 2, 2, 538, 539, 7, 52, 2, 2, 539, 540, 5, 80, 41, 2, 540, 541, 7,
	53, 2, 2, 541, 542, 8, 39, 1, 2, 542, 77, 3, 2, 2, 2, 543, 544, 7, 22,
	2, 2, 544, 554, 8, 40, 1, 2, 545, 546, 7, 23, 2, 2, 546, 554, 8, 40, 1,
	2, 547, 548, 7, 24, 2, 2, 548, 554, 8, 40, 1, 2, 549, 550, 7, 26, 2, 2,
	550, 554, 8, 40, 1, 2, 551, 552, 7, 25, 2, 2, 552, 554, 8, 40, 1, 2, 553,
	543, 3, 2, 2, 2, 553, 545, 3, 2, 2, 2, 553, 547, 3, 2, 2, 2, 553, 549,
	3, 2, 2, 2, 553, 551, 3, 2, 2, 2, 554, 79, 3, 2, 2, 2, 555, 557, 5, 82,
	42, 2, 556, 555, 3, 2, 2, 2, 557, 558, 3, 2, 2, 2, 558, 556, 3, 2, 2, 2,
	558, 559, 3, 2, 2, 2, 559, 560, 3, 2, 2, 2, 560, 561, 8, 41, 1, 2, 561,
	81, 3, 2, 2, 2, 562, 563, 5, 84, 43, 2, 563, 564, 7, 36, 2, 2, 564, 565,
	8, 42, 1, 2, 565, 570, 3, 2, 2, 2, 566, 567, 5, 84, 43, 2, 567, 568, 8,
	42, 1, 2, 568, 570, 3, 2, 2, 2, 569, 562, 3, 2, 2, 2, 569, 566, 3, 2, 2,
	2, 570, 83, 3, 2, 2, 2, 571, 572, 5, 86, 44, 2, 572, 573, 8, 43, 1, 2,
	573, 85, 3, 2, 2, 2, 574, 575, 8, 44, 1, 2, 575, 576, 7, 52, 2, 2, 576,
	577, 5, 86, 44, 2, 577, 578, 9, 3, 2, 2, 578, 579, 7, 53, 2, 2, 579, 580,
	9, 4, 2, 2, 580, 581, 5, 86, 44, 18, 581, 582, 8, 44, 1, 2, 582, 652, 3,
	2, 2, 2, 583, 584, 7, 52, 2, 2, 584, 585, 5, 86, 44, 2, 585, 586, 9, 3,
	2, 2, 586, 587, 7, 53, 2, 2, 587, 588, 9, 4, 2, 2, 588, 589, 7, 52, 2,
	2, 589, 590, 5, 86, 44, 2, 590, 591, 9, 3, 2, 2, 591, 592, 7, 53, 2, 2,
	592, 593, 8, 44, 1, 2, 593, 652, 3, 2, 2, 2, 594, 595, 7, 52, 2, 2, 595,
	596, 5, 86, 44, 2, 596, 597, 9, 3, 2, 2, 597, 598, 7, 53, 2, 2, 598, 599,
	9, 5, 2, 2, 599, 600, 5, 86, 44, 14, 600, 601, 8, 44, 1, 2, 601, 652, 3,
	2, 2, 2, 602, 603, 7, 52, 2, 2, 603, 604, 5, 86, 44, 2, 604, 605, 9, 3,
	2, 2, 605, 606, 7, 53, 2, 2, 606, 607, 9, 5, 2, 2, 607, 608, 7, 52, 2,
	2, 608, 609, 5, 86, 44, 2, 609, 610, 9, 3, 2, 2, 610, 611, 7, 53, 2, 2,
	611, 612, 8, 44, 1, 2, 612, 652, 3, 2, 2, 2, 613, 614, 7, 52, 2, 2, 614,
	615, 5, 86, 44, 2, 615, 616, 9, 3, 2, 2, 616, 617, 7, 53, 2, 2, 617, 618,
	9, 6, 2, 2, 618, 619, 5, 86, 44, 10, 619, 620, 8, 44, 1, 2, 620, 652, 3,
	2, 2, 2, 621, 622, 7, 52, 2, 2, 622, 623, 5, 86, 44, 2, 623, 624, 9, 3,
	2, 2, 624, 625, 7, 53, 2, 2, 625, 626, 9, 6, 2, 2, 626, 627, 7, 52, 2,
	2, 627, 628, 5, 86, 44, 2, 628, 629, 9, 3, 2, 2, 629, 630, 7, 53, 2, 2,
	630, 631, 8, 44, 1, 2, 631, 652, 3, 2, 2, 2, 632, 633, 9, 7, 2, 2, 633,
	634, 5, 84, 43, 2, 634, 635, 8, 44, 1, 2, 635, 652, 3, 2, 2, 2, 636, 637,
	9, 7, 2, 2, 637, 638, 7, 52, 2, 2, 638, 639, 5, 86, 44, 2, 639, 640, 9,
	3, 2, 2, 640, 641, 7, 53, 2, 2, 641, 642, 8, 44, 1, 2, 642, 652, 3, 2,
	2, 2, 643, 644, 5, 88, 45, 2, 644, 645, 8, 44, 1, 2, 645, 652, 3, 2, 2,
	2, 646, 647, 7, 52, 2, 2, 647, 648, 5, 84, 43, 2, 648, 649, 7, 53, 2, 2,
	649, 650, 8, 44, 1, 2, 650, 652, 3, 2, 2, 2, 651, 574, 3, 2, 2, 2, 651,
	583, 3, 2, 2, 2, 651, 594, 3, 2, 2, 2, 651, 602, 3, 2, 2, 2, 651, 613,
	3, 2, 2, 2, 651, 621, 3, 2, 2, 2, 651, 632, 3, 2, 2, 2, 651, 636, 3, 2,
	2, 2, 651, 643, 3, 2, 2, 2, 651, 646, 3, 2, 2, 2, 652, 699, 3, 2, 2, 2,
	653, 654, 12, 19, 2, 2, 654, 655, 9, 4, 2, 2, 655, 656, 5, 86, 44, 20,
	656, 657, 8, 44, 1, 2, 657, 698, 3, 2, 2, 2, 658, 659, 12, 15, 2, 2, 659,
	660, 9, 5, 2, 2, 660, 661, 5, 86, 44, 16, 661, 662, 8, 44, 1, 2, 662, 698,
	3, 2, 2, 2, 663, 664, 12, 11, 2, 2, 664, 665, 9, 6, 2, 2, 665, 666, 5,
	86, 44, 12, 666, 667, 8, 44, 1, 2, 667, 698, 3, 2, 2, 2, 668, 669, 12,
	7, 2, 2, 669, 670, 9, 8, 2, 2, 670, 671, 5, 86, 44, 8, 671, 672, 8, 44,
	1, 2, 672, 698, 3, 2, 2, 2, 673, 674, 12, 17, 2, 2, 674, 675, 9, 4, 2,
	2, 675, 676, 7, 52, 2, 2, 676, 677, 5, 86, 44, 2, 677, 678, 9, 3, 2, 2,
	678, 679, 7, 53, 2, 2, 679, 680, 8, 44, 1, 2, 680, 698, 3, 2, 2, 2, 681,
	682, 12, 13, 2, 2, 682, 683, 9, 5, 2, 2, 683, 684, 7, 52, 2, 2, 684, 685,
	5, 86, 44, 2, 685, 686, 9, 3, 2, 2, 686, 687, 7, 53, 2, 2, 687, 688, 8,
	44, 1, 2, 688, 698, 3, 2, 2, 2, 689, 690, 12, 9, 2, 2, 690, 691, 9, 6,
	2, 2, 691, 692, 7, 52, 2, 2, 692, 693, 5, 86, 44, 2, 693, 694, 9, 3, 2,
	2, 694, 695, 7, 53, 2, 2, 695, 696, 8, 44, 1, 2, 696, 698, 3, 2, 2, 2,
	697, 653, 3, 2, 2, 2, 697, 658, 3, 2, 2, 2, 697, 663, 3, 2, 2, 2, 697,
	668, 3, 2, 2, 2, 697, 673, 3, 2, 2, 2, 697, 681, 3, 2, 2, 2, 697, 689,
	3, 2, 2, 2, 698, 701, 3, 2, 2, 2, 699, 697, 3, 2, 2, 2, 699, 700, 3, 2,
	2, 2, 700, 87, 3, 2, 2, 2, 701, 699, 3, 2, 2, 2, 702, 703, 7, 27, 2, 2,
	703, 725, 8, 45, 1, 2, 704, 705, 7, 28, 2, 2, 705, 725, 8, 45, 1, 2, 706,
	707, 7, 30, 2, 2, 707, 725, 8, 45, 1, 2, 708, 709, 7, 31, 2, 2, 709, 725,
	8, 45, 1, 2, 710, 711, 7, 32, 2, 2, 711, 725, 8, 45, 1, 2, 712, 713, 5,
	22, 12, 2, 713, 714, 8, 45, 1, 2, 714, 725, 3, 2, 2, 2, 715, 716, 5, 42,
	22, 2, 716, 717, 8, 45, 1, 2, 717, 725, 3, 2, 2, 2, 718, 719, 5, 58, 30,
	2, 719, 720, 8, 45, 1, 2, 720, 725, 3, 2, 2, 2, 721, 722, 5, 76, 39, 2,
	722, 723, 8, 45, 1, 2, 723, 725, 3, 2, 2, 2, 724, 702, 3, 2, 2, 2, 724,
	704, 3, 2, 2, 2, 724, 706, 3, 2, 2, 2, 724, 708, 3, 2, 2, 2, 724, 710,
	3, 2, 2, 2, 724, 712, 3, 2, 2, 2, 724, 715, 3, 2, 2, 2, 724, 718, 3, 2,
	2, 2, 724, 721, 3, 2, 2, 2, 725, 89, 3, 2, 2, 2, 34, 96, 103, 149, 176,
	212, 247, 252, 284, 289, 309, 324, 332, 339, 350, 365, 370, 389, 400, 407,
	418, 465, 476, 509, 514, 529, 553, 558, 569, 651, 697, 699, 724,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'print!'", "'println!'", "'number'", "'as f64'", "'as i64'", "'let'",
	"'mut'", "'if'", "'else'", "'match'", "'while'", "'break'", "'continue'",
	"'return'", "'loop'", "'for'", "'in'", "'main'", "'fn'", "'i64'", "'f64'",
	"'String'", "'bool'", "'&str'", "", "", "", "", "", "", "'..'", "'.'",
	"';'", "','", "':'", "'='", "'=='", "'>='", "'=>'", "'->'", "'<='", "'!='",
	"'>'", "'<'", "'*'", "'/'", "'%'", "'+'", "'-'", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "'&&'", "'||'", "'|'", "'!'",
}
var symbolicNames = []string{
	"", "R_PRINT", "R_PRINTLN", "P_NUMBER", "R_AS_DOUBLE", "R_AS_INTEGER",
	"R_LET", "R_MUT", "R_IF", "R_ELSE", "R_MATCH", "R_WHILE", "R_BREAK", "R_CONTINUE",
	"R_RETURN", "R_LOOP", "R_FOR", "R_IN", "R_MAIN", "R_FUNCTION", "R_INT",
	"R_FLOAT", "R_STRING", "R_BOOL", "R_STR", "NUMBER", "DOUBLE", "CHAR", "STRING",
	"BOOLEAN", "ID", "TK_DOBLEPUNTO", "TK_PUNTO", "TK_PUNTOCOMA", "TK_COMA",
	"TK_DOSPUNTOS", "TK_IGUAL", "TK_IGUALIGUAL", "TK_MAYORIGUAL", "TK_IGUALMAYOR",
	"TK_MENOSMAYOR", "TK_MENORIGUAL", "TK_DIFIGUAL", "TK_MAYOR", "TK_MENOR",
	"TK_MULT", "TK_DIV", "TK_MODULO", "TK_MAS", "TK_MENOS", "TK_PARA", "TK_PARC",
	"TK_LLAVEA", "TK_LLAVEC", "TK_CORA", "TK_CORC", "TK_AND", "TK_OR", "TK_BARRA",
	"TK_NOT", "WHITESPACE", "TK_MULTI", "TK_LINE",
}

var ruleNames = []string{
	"start", "instrucciones", "end_instr", "instruccion", "instr_main", "instr_println",
	"instr_declaracion", "instr_asignacion", "instr_if", "instr_else_if", "instr_ternario",
	"instr_else_if_ternario", "instr_match", "instr_case", "block_instr_match",
	"list_case", "list_expre_case", "block_case", "instr_default", "block_default",
	"instr_match_ter", "instr_case_ter", "list_case_ternario", "list_expre_case_ter",
	"block_case_ter", "instr_default_ter", "instr_while", "instr_loop", "instr_loop_ternario",
	"instr_for_in", "instr_break", "instr_continue", "instr_return", "instr_func",
	"list_function_parameters", "block_parameters_fn", "instr_llamada", "instr_llamada_expre",
	"instr_tipo", "list_expression", "block_list_expression", "expression",
	"exp_arit", "primitivo",
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
	ChemsR_PRINT       = 1
	ChemsR_PRINTLN     = 2
	ChemsP_NUMBER      = 3
	ChemsR_AS_DOUBLE   = 4
	ChemsR_AS_INTEGER  = 5
	ChemsR_LET         = 6
	ChemsR_MUT         = 7
	ChemsR_IF          = 8
	ChemsR_ELSE        = 9
	ChemsR_MATCH       = 10
	ChemsR_WHILE       = 11
	ChemsR_BREAK       = 12
	ChemsR_CONTINUE    = 13
	ChemsR_RETURN      = 14
	ChemsR_LOOP        = 15
	ChemsR_FOR         = 16
	ChemsR_IN          = 17
	ChemsR_MAIN        = 18
	ChemsR_FUNCTION    = 19
	ChemsR_INT         = 20
	ChemsR_FLOAT       = 21
	ChemsR_STRING      = 22
	ChemsR_BOOL        = 23
	ChemsR_STR         = 24
	ChemsNUMBER        = 25
	ChemsDOUBLE        = 26
	ChemsCHAR          = 27
	ChemsSTRING        = 28
	ChemsBOOLEAN       = 29
	ChemsID            = 30
	ChemsTK_DOBLEPUNTO = 31
	ChemsTK_PUNTO      = 32
	ChemsTK_PUNTOCOMA  = 33
	ChemsTK_COMA       = 34
	ChemsTK_DOSPUNTOS  = 35
	ChemsTK_IGUAL      = 36
	ChemsTK_IGUALIGUAL = 37
	ChemsTK_MAYORIGUAL = 38
	ChemsTK_IGUALMAYOR = 39
	ChemsTK_MENOSMAYOR = 40
	ChemsTK_MENORIGUAL = 41
	ChemsTK_DIFIGUAL   = 42
	ChemsTK_MAYOR      = 43
	ChemsTK_MENOR      = 44
	ChemsTK_MULT       = 45
	ChemsTK_DIV        = 46
	ChemsTK_MODULO     = 47
	ChemsTK_MAS        = 48
	ChemsTK_MENOS      = 49
	ChemsTK_PARA       = 50
	ChemsTK_PARC       = 51
	ChemsTK_LLAVEA     = 52
	ChemsTK_LLAVEC     = 53
	ChemsTK_CORA       = 54
	ChemsTK_CORC       = 55
	ChemsTK_AND        = 56
	ChemsTK_OR         = 57
	ChemsTK_BARRA      = 58
	ChemsTK_NOT        = 59
	ChemsWHITESPACE    = 60
	ChemsTK_MULTI      = 61
	ChemsTK_LINE       = 62
)

// Chems rules.
const (
	ChemsRULE_start                    = 0
	ChemsRULE_instrucciones            = 1
	ChemsRULE_end_instr                = 2
	ChemsRULE_instruccion              = 3
	ChemsRULE_instr_main               = 4
	ChemsRULE_instr_println            = 5
	ChemsRULE_instr_declaracion        = 6
	ChemsRULE_instr_asignacion         = 7
	ChemsRULE_instr_if                 = 8
	ChemsRULE_instr_else_if            = 9
	ChemsRULE_instr_ternario           = 10
	ChemsRULE_instr_else_if_ternario   = 11
	ChemsRULE_instr_match              = 12
	ChemsRULE_instr_case               = 13
	ChemsRULE_block_instr_match        = 14
	ChemsRULE_list_case                = 15
	ChemsRULE_list_expre_case          = 16
	ChemsRULE_block_case               = 17
	ChemsRULE_instr_default            = 18
	ChemsRULE_block_default            = 19
	ChemsRULE_instr_match_ter          = 20
	ChemsRULE_instr_case_ter           = 21
	ChemsRULE_list_case_ternario       = 22
	ChemsRULE_list_expre_case_ter      = 23
	ChemsRULE_block_case_ter           = 24
	ChemsRULE_instr_default_ter        = 25
	ChemsRULE_instr_while              = 26
	ChemsRULE_instr_loop               = 27
	ChemsRULE_instr_loop_ternario      = 28
	ChemsRULE_instr_for_in             = 29
	ChemsRULE_instr_break              = 30
	ChemsRULE_instr_continue           = 31
	ChemsRULE_instr_return             = 32
	ChemsRULE_instr_func               = 33
	ChemsRULE_list_function_parameters = 34
	ChemsRULE_block_parameters_fn      = 35
	ChemsRULE_instr_llamada            = 36
	ChemsRULE_instr_llamada_expre      = 37
	ChemsRULE_instr_tipo               = 38
	ChemsRULE_list_expression          = 39
	ChemsRULE_block_list_expression    = 40
	ChemsRULE_expression               = 41
	ChemsRULE_exp_arit                 = 42
	ChemsRULE_primitivo                = 43
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
		p.SetState(88)

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
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_PRINT)|(1<<ChemsR_PRINTLN)|(1<<ChemsR_LET)|(1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_WHILE)|(1<<ChemsR_BREAK)|(1<<ChemsR_CONTINUE)|(1<<ChemsR_RETURN)|(1<<ChemsR_LOOP)|(1<<ChemsR_FOR)|(1<<ChemsR_FUNCTION)|(1<<ChemsID))) != 0) {
		{
			p.SetState(91)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).list = append(localctx.(*InstruccionesContext).list, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(94)
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

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_PUNTOCOMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*End_instrContext).v = 1

	case ChemsEOF, ChemsR_PRINT, ChemsR_PRINTLN, ChemsR_LET, ChemsR_IF, ChemsR_MATCH, ChemsR_WHILE, ChemsR_BREAK, ChemsR_CONTINUE, ChemsR_RETURN, ChemsR_LOOP, ChemsR_FOR, ChemsR_FUNCTION, ChemsID, ChemsTK_COMA, ChemsTK_LLAVEC:
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

	// Get_instr_return returns the _instr_return rule contexts.
	Get_instr_return() IInstr_returnContext

	// Get_instr_loop returns the _instr_loop rule contexts.
	Get_instr_loop() IInstr_loopContext

	// Get_instr_for_in returns the _instr_for_in rule contexts.
	Get_instr_for_in() IInstr_for_inContext

	// Get_instr_main returns the _instr_main rule contexts.
	Get_instr_main() IInstr_mainContext

	// Get_instr_func returns the _instr_func rule contexts.
	Get_instr_func() IInstr_funcContext

	// Get_instr_llamada returns the _instr_llamada rule contexts.
	Get_instr_llamada() IInstr_llamadaContext

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

	// Set_instr_return sets the _instr_return rule contexts.
	Set_instr_return(IInstr_returnContext)

	// Set_instr_loop sets the _instr_loop rule contexts.
	Set_instr_loop(IInstr_loopContext)

	// Set_instr_for_in sets the _instr_for_in rule contexts.
	Set_instr_for_in(IInstr_for_inContext)

	// Set_instr_main sets the _instr_main rule contexts.
	Set_instr_main(IInstr_mainContext)

	// Set_instr_func sets the _instr_func rule contexts.
	Set_instr_func(IInstr_funcContext)

	// Set_instr_llamada sets the _instr_llamada rule contexts.
	Set_instr_llamada(IInstr_llamadaContext)

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
	_instr_return      IInstr_returnContext
	_instr_loop        IInstr_loopContext
	_instr_for_in      IInstr_for_inContext
	_instr_main        IInstr_mainContext
	_instr_func        IInstr_funcContext
	_instr_llamada     IInstr_llamadaContext
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

func (s *InstruccionContext) Get_instr_return() IInstr_returnContext { return s._instr_return }

func (s *InstruccionContext) Get_instr_loop() IInstr_loopContext { return s._instr_loop }

func (s *InstruccionContext) Get_instr_for_in() IInstr_for_inContext { return s._instr_for_in }

func (s *InstruccionContext) Get_instr_main() IInstr_mainContext { return s._instr_main }

func (s *InstruccionContext) Get_instr_func() IInstr_funcContext { return s._instr_func }

func (s *InstruccionContext) Get_instr_llamada() IInstr_llamadaContext { return s._instr_llamada }

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

func (s *InstruccionContext) Set_instr_return(v IInstr_returnContext) { s._instr_return = v }

func (s *InstruccionContext) Set_instr_loop(v IInstr_loopContext) { s._instr_loop = v }

func (s *InstruccionContext) Set_instr_for_in(v IInstr_for_inContext) { s._instr_for_in = v }

func (s *InstruccionContext) Set_instr_main(v IInstr_mainContext) { s._instr_main = v }

func (s *InstruccionContext) Set_instr_func(v IInstr_funcContext) { s._instr_func = v }

func (s *InstruccionContext) Set_instr_llamada(v IInstr_llamadaContext) { s._instr_llamada = v }

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

func (s *InstruccionContext) Instr_return() IInstr_returnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_returnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_returnContext)
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

func (s *InstruccionContext) Instr_func() IInstr_funcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_funcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_funcContext)
}

func (s *InstruccionContext) Instr_llamada() IInstr_llamadaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_llamadaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_llamadaContext)
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

	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)

			var _x = p.Instr_println()

			localctx.(*InstruccionContext)._instr_println = _x
		}
		{
			p.SetState(104)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_println().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)

			var _x = p.Instr_declaracion()

			localctx.(*InstruccionContext)._instr_declaracion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_declaracion().GetInstr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)

			var _x = p.Instr_asignacion()

			localctx.(*InstruccionContext)._instr_asignacion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_asignacion().GetInstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)

			var _x = p.Instr_if()

			localctx.(*InstruccionContext)._instr_if = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_if().GetInstr()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(116)

			var _x = p.Instr_match()

			localctx.(*InstruccionContext)._instr_match = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_match().GetInstr()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(119)

			var _x = p.Instr_while()

			localctx.(*InstruccionContext)._instr_while = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_while().GetInstr()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(122)

			var _x = p.Instr_break()

			localctx.(*InstruccionContext)._instr_break = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_break().GetInstr()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(125)

			var _x = p.Instr_continue()

			localctx.(*InstruccionContext)._instr_continue = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_continue().GetInstr()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(128)

			var _x = p.Instr_return()

			localctx.(*InstruccionContext)._instr_return = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_return().GetInstr()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(131)

			var _x = p.Instr_loop()

			localctx.(*InstruccionContext)._instr_loop = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_loop().GetInstr()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(134)

			var _x = p.Instr_for_in()

			localctx.(*InstruccionContext)._instr_for_in = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_for_in().GetInstr()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(137)

			var _x = p.Instr_main()

			localctx.(*InstruccionContext)._instr_main = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_main().GetInstr()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(140)

			var _x = p.Instr_func()

			localctx.(*InstruccionContext)._instr_func = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_func().GetInstr()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(143)

			var _x = p.Instr_llamada()

			localctx.(*InstruccionContext)._instr_llamada = _x
		}
		{
			p.SetState(144)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_llamada().GetInstr()

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
		p.SetState(149)
		p.Match(ChemsR_FUNCTION)
	}
	{
		p.SetState(150)

		var _m = p.Match(ChemsR_MAIN)

		localctx.(*Instr_mainContext)._R_MAIN = _m
	}
	{
		p.SetState(151)
		p.Match(ChemsTK_PARA)
	}
	{
		p.SetState(152)
		p.Match(ChemsTK_PARC)
	}
	{
		p.SetState(153)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(154)

		var _x = p.Instrucciones()

		localctx.(*Instr_mainContext)._instrucciones = _x
	}
	{
		p.SetState(155)
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

	// GetSalto returns the salto token.
	GetSalto() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// SetSalto sets the salto token.
	SetSalto(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Get_list_expression returns the _list_expression rule contexts.
	Get_list_expression() IList_expressionContext

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// Set_list_expression sets the _list_expression rule contexts.
	Set_list_expression(IList_expressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_printlnContext differentiates from other interfaces.
	IsInstr_printlnContext()
}

type Instr_printlnContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	instr            interfaces.Instruction
	salto            antlr.Token
	_primitivo       IPrimitivoContext
	_STRING          antlr.Token
	_list_expression IList_expressionContext
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

func (s *Instr_printlnContext) GetSalto() antlr.Token { return s.salto }

func (s *Instr_printlnContext) Get_STRING() antlr.Token { return s._STRING }

func (s *Instr_printlnContext) SetSalto(v antlr.Token) { s.salto = v }

func (s *Instr_printlnContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *Instr_printlnContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Instr_printlnContext) Get_list_expression() IList_expressionContext {
	return s._list_expression
}

func (s *Instr_printlnContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Instr_printlnContext) Set_list_expression(v IList_expressionContext) { s._list_expression = v }

func (s *Instr_printlnContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_printlnContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_printlnContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_printlnContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Instr_printlnContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_printlnContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *Instr_printlnContext) R_PRINTLN() antlr.TerminalNode {
	return s.GetToken(ChemsR_PRINTLN, 0)
}

func (s *Instr_printlnContext) R_PRINT() antlr.TerminalNode {
	return s.GetToken(ChemsR_PRINT, 0)
}

func (s *Instr_printlnContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChemsSTRING, 0)
}

func (s *Instr_printlnContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_printlnContext) List_expression() IList_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expressionContext)
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

	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Instr_printlnContext).salto = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsR_PRINT || _la == ChemsR_PRINTLN) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Instr_printlnContext).salto = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(159)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(160)

			var _x = p.Primitivo()

			localctx.(*Instr_printlnContext)._primitivo = _x
		}
		{
			p.SetState(161)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(162)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_printlnContext).instr = instruction.PRINTLN(nil, "-1", localctx.(*Instr_printlnContext).Get_primitivo().GetP(), (func() string {
			if localctx.(*Instr_printlnContext).GetSalto() == nil {
				return ""
			} else {
				return localctx.(*Instr_printlnContext).GetSalto().GetText()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Instr_printlnContext).salto = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsR_PRINT || _la == ChemsR_PRINTLN) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Instr_printlnContext).salto = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(166)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(167)

			var _m = p.Match(ChemsSTRING)

			localctx.(*Instr_printlnContext)._STRING = _m
		}
		{
			p.SetState(168)
			p.Match(ChemsTK_COMA)
		}
		{
			p.SetState(169)

			var _x = p.List_expression()

			localctx.(*Instr_printlnContext)._list_expression = _x
		}
		{
			p.SetState(170)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(171)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_printlnContext).instr = instruction.PRINTLN(localctx.(*Instr_printlnContext).Get_list_expression().GetL(), (func() string {
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
		}()))-1], nil, (func() string {
			if localctx.(*Instr_printlnContext).GetSalto() == nil {
				return ""
			} else {
				return localctx.(*Instr_printlnContext).GetSalto().GetText()
			}
		}()))

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

	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(177)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(178)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(179)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(180)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(181)
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
			p.SetState(184)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(185)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(186)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(187)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(188)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(189)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(190)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(191)
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
			p.SetState(194)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(195)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(196)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(197)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(198)
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
			p.SetState(201)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(202)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(203)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(204)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(205)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(206)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(207)
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
		p.SetState(212)

		var _m = p.Match(ChemsID)

		localctx.(*Instr_asignacionContext)._ID = _m
	}
	{
		p.SetState(213)
		p.Match(ChemsTK_IGUAL)
	}
	{
		p.SetState(214)

		var _x = p.Expression()

		localctx.(*Instr_asignacionContext)._expression = _x
	}
	{
		p.SetState(215)
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

	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(219)

			var _x = p.Expression()

			localctx.(*Instr_ifContext)._expression = _x
		}
		{
			p.SetState(220)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(221)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(222)
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
			p.SetState(225)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(226)

			var _x = p.Expression()

			localctx.(*Instr_ifContext)._expression = _x
		}
		{
			p.SetState(227)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(228)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(229)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(230)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(231)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(232)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).right_intr = _x
		}
		{
			p.SetState(233)
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
			p.SetState(236)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(237)

			var _x = p.Expression()

			localctx.(*Instr_ifContext)._expression = _x
		}
		{
			p.SetState(238)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(239)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(240)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(241)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(242)

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
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(247)

				var _x = p.Instr_if()

				localctx.(*Instr_else_ifContext)._instr_if = _x
			}
			localctx.(*Instr_else_ifContext).e = append(localctx.(*Instr_else_ifContext).e, localctx.(*Instr_else_ifContext)._instr_if)

		}
		p.SetState(252)
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

	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(255)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(256)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).block = _x
		}
		{
			p.SetState(257)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(258)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(259)
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
			p.SetState(262)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(263)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).block = _x
		}
		{
			p.SetState(264)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(265)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(266)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(267)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(268)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(269)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).right_intr = _x
		}
		{
			p.SetState(270)
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
			p.SetState(273)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(274)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).block = _x
		}
		{
			p.SetState(275)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(276)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(277)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(278)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(279)

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
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(284)

				var _x = p.Instr_ternario()

				localctx.(*Instr_else_if_ternarioContext)._instr_ternario = _x
			}
			localctx.(*Instr_else_if_ternarioContext).e = append(localctx.(*Instr_else_if_ternarioContext).e, localctx.(*Instr_else_if_ternarioContext)._instr_ternario)

		}
		p.SetState(289)
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

	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(292)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_matchContext)._R_MATCH = _m
		}
		{
			p.SetState(293)

			var _x = p.Expression()

			localctx.(*Instr_matchContext)._expression = _x
		}
		{
			p.SetState(294)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(295)

			var _x = p.List_case()

			localctx.(*Instr_matchContext)._list_case = _x
		}
		{
			p.SetState(296)

			var _x = p.Block_default()

			localctx.(*Instr_matchContext)._block_default = _x
		}
		{
			p.SetState(297)
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
			p.SetState(300)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_matchContext)._R_MATCH = _m
		}
		{
			p.SetState(301)

			var _x = p.Expression()

			localctx.(*Instr_matchContext)._expression = _x
		}
		{
			p.SetState(302)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(303)

			var _x = p.Block_default()

			localctx.(*Instr_matchContext)._block_default = _x
		}
		{
			p.SetState(304)
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

	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)

			var _x = p.List_expre_case()

			localctx.(*Instr_caseContext)._list_expre_case = _x
		}
		{
			p.SetState(310)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(311)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(312)

			var _x = p.Instrucciones()

			localctx.(*Instr_caseContext)._instrucciones = _x
		}
		{
			p.SetState(313)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_caseContext).instr = control.NewCase(nil, localctx.(*Instr_caseContext).Get_list_expre_case().GetL(), localctx.(*Instr_caseContext).Get_instrucciones().GetL())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(316)

			var _x = p.List_expre_case()

			localctx.(*Instr_caseContext)._list_expre_case = _x
		}
		{
			p.SetState(317)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(318)

			var _x = p.Block_instr_match()

			localctx.(*Instr_caseContext)._block_instr_match = _x
		}
		{
			p.SetState(319)
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
		p.SetState(324)

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
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(327)

				var _x = p.Instr_case()

				localctx.(*List_caseContext)._instr_case = _x
			}
			localctx.(*List_caseContext).e = append(localctx.(*List_caseContext).e, localctx.(*List_caseContext)._instr_case)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(330)
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
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-49)&-(0x1f+1)) == 0 && ((1<<uint((_la-49)))&((1<<(ChemsTK_MENOS-49))|(1<<(ChemsTK_PARA-49))|(1<<(ChemsTK_NOT-49)))) != 0) {
		{
			p.SetState(334)

			var _x = p.Block_case()

			localctx.(*List_expre_caseContext)._block_case = _x
		}
		localctx.(*List_expre_caseContext).e = append(localctx.(*List_expre_caseContext).e, localctx.(*List_expre_caseContext)._block_case)

		p.SetState(337)
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

	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(341)

			var _x = p.Expression()

			localctx.(*Block_caseContext)._expression = _x
		}
		{
			p.SetState(342)
			p.Match(ChemsTK_BARRA)
		}
		localctx.(*Block_caseContext).instr = control.NewCase(localctx.(*Block_caseContext).Get_expression().GetP(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)

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

	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(350)
			p.Match(ChemsID)
		}
		{
			p.SetState(351)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(352)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(353)

			var _x = p.Instrucciones()

			localctx.(*Instr_defaultContext)._instrucciones = _x
		}
		{
			p.SetState(354)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_defaultContext).instr = control.NewDefault(localctx.(*Instr_defaultContext).Get_instrucciones().GetL())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
			p.Match(ChemsID)
		}
		{
			p.SetState(358)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(359)

			var _x = p.Block_instr_match()

			localctx.(*Instr_defaultContext)._block_instr_match = _x
		}
		{
			p.SetState(360)
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
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ChemsID {
		{
			p.SetState(365)

			var _x = p.Instr_default()

			localctx.(*Block_defaultContext)._instr_default = _x
		}
		localctx.(*Block_defaultContext).e = append(localctx.(*Block_defaultContext).e, localctx.(*Block_defaultContext)._instr_default)

		p.SetState(368)
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

	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(372)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_match_terContext)._R_MATCH = _m
		}
		{
			p.SetState(373)

			var _x = p.Expression()

			localctx.(*Instr_match_terContext)._expression = _x
		}
		{
			p.SetState(374)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(375)

			var _x = p.List_case_ternario()

			localctx.(*Instr_match_terContext)._list_case_ternario = _x
		}
		{
			p.SetState(376)

			var _x = p.Instr_default_ter()

			localctx.(*Instr_match_terContext)._instr_default_ter = _x
		}
		{
			p.SetState(377)
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
			p.SetState(380)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_match_terContext)._R_MATCH = _m
		}
		{
			p.SetState(381)

			var _x = p.Expression()

			localctx.(*Instr_match_terContext)._expression = _x
		}
		{
			p.SetState(382)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(383)

			var _x = p.Instr_default_ter()

			localctx.(*Instr_match_terContext)._instr_default_ter = _x
		}
		{
			p.SetState(384)
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
		p.SetState(389)

		var _x = p.List_expre_case_ter()

		localctx.(*Instr_case_terContext)._list_expre_case_ter = _x
	}
	{
		p.SetState(390)
		p.Match(ChemsTK_IGUALMAYOR)
	}
	{
		p.SetState(391)

		var _x = p.Expression()

		localctx.(*Instr_case_terContext)._expression = _x
	}
	{
		p.SetState(392)
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
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(395)

				var _x = p.Instr_case_ter()

				localctx.(*List_case_ternarioContext)._instr_case_ter = _x
			}
			localctx.(*List_case_ternarioContext).e = append(localctx.(*List_case_ternarioContext).e, localctx.(*List_case_ternarioContext)._instr_case_ter)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(398)
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
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-49)&-(0x1f+1)) == 0 && ((1<<uint((_la-49)))&((1<<(ChemsTK_MENOS-49))|(1<<(ChemsTK_PARA-49))|(1<<(ChemsTK_NOT-49)))) != 0) {
		{
			p.SetState(402)

			var _x = p.Block_case_ter()

			localctx.(*List_expre_case_terContext)._block_case_ter = _x
		}
		localctx.(*List_expre_case_terContext).e = append(localctx.(*List_expre_case_terContext).e, localctx.(*List_expre_case_terContext)._block_case_ter)

		p.SetState(405)
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

	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)

			var _x = p.Expression()

			localctx.(*Block_case_terContext)._expression = _x
		}
		{
			p.SetState(410)
			p.Match(ChemsTK_BARRA)
		}
		localctx.(*Block_case_terContext).instr = control.NewCaseTer(localctx.(*Block_case_terContext).Get_expression().GetP(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(413)

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
		p.SetState(418)
		p.Match(ChemsID)
	}
	{
		p.SetState(419)
		p.Match(ChemsTK_IGUALMAYOR)
	}
	{
		p.SetState(420)

		var _x = p.Expression()

		localctx.(*Instr_default_terContext)._expression = _x
	}
	{
		p.SetState(421)
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
		p.SetState(424)
		p.Match(ChemsR_WHILE)
	}
	{
		p.SetState(425)

		var _x = p.Expression()

		localctx.(*Instr_whileContext)._expression = _x
	}
	{
		p.SetState(426)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(427)

		var _x = p.Instrucciones()

		localctx.(*Instr_whileContext)._instrucciones = _x
	}
	{
		p.SetState(428)
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
		p.SetState(431)
		p.Match(ChemsR_LOOP)
	}
	{
		p.SetState(432)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(433)

		var _x = p.Instrucciones()

		localctx.(*Instr_loopContext)._instrucciones = _x
	}
	{
		p.SetState(434)
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
		p.SetState(437)
		p.Match(ChemsR_LOOP)
	}
	{
		p.SetState(438)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(439)

		var _x = p.Instrucciones()

		localctx.(*Instr_loop_ternarioContext)._instrucciones = _x
	}
	{
		p.SetState(440)
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

	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(443)

			var _m = p.Match(ChemsR_FOR)

			localctx.(*Instr_for_inContext)._R_FOR = _m
		}
		{
			p.SetState(444)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_for_inContext)._ID = _m
		}
		{
			p.SetState(445)
			p.Match(ChemsR_IN)
		}
		{
			p.SetState(446)

			var _x = p.Expression()

			localctx.(*Instr_for_inContext).left = _x
		}
		{
			p.SetState(447)
			p.Match(ChemsTK_DOBLEPUNTO)
		}
		{
			p.SetState(448)

			var _x = p.Expression()

			localctx.(*Instr_for_inContext).right = _x
		}
		{
			p.SetState(449)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(450)

			var _x = p.Instrucciones()

			localctx.(*Instr_for_inContext)._instrucciones = _x
		}
		{
			p.SetState(451)
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
			p.SetState(454)

			var _m = p.Match(ChemsR_FOR)

			localctx.(*Instr_for_inContext)._R_FOR = _m
		}
		{
			p.SetState(455)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_for_inContext)._ID = _m
		}
		{
			p.SetState(456)
			p.Match(ChemsR_IN)
		}
		{
			p.SetState(457)

			var _x = p.Expression()

			localctx.(*Instr_for_inContext).left = _x
		}
		{
			p.SetState(458)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(459)

			var _x = p.Instrucciones()

			localctx.(*Instr_for_inContext)._instrucciones = _x
		}
		{
			p.SetState(460)
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

	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(465)

			var _m = p.Match(ChemsR_BREAK)

			localctx.(*Instr_breakContext)._R_BREAK = _m
		}
		{
			p.SetState(466)
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
			p.SetState(469)

			var _m = p.Match(ChemsR_BREAK)

			localctx.(*Instr_breakContext)._R_BREAK = _m
		}
		{
			p.SetState(470)

			var _x = p.Expression()

			localctx.(*Instr_breakContext)._expression = _x
		}
		{
			p.SetState(471)
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
		p.SetState(476)

		var _m = p.Match(ChemsR_CONTINUE)

		localctx.(*Instr_continueContext)._R_CONTINUE = _m
	}
	{
		p.SetState(477)
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
		p.SetState(480)

		var _m = p.Match(ChemsR_RETURN)

		localctx.(*Instr_returnContext)._R_RETURN = _m
	}
	{
		p.SetState(481)

		var _x = p.Expression()

		localctx.(*Instr_returnContext)._expression = _x
	}
	{
		p.SetState(482)
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

// IInstr_funcContext is an interface to support dynamic dispatch.
type IInstr_funcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_FUNCTION returns the _R_FUNCTION token.
	Get_R_FUNCTION() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_R_FUNCTION sets the _R_FUNCTION token.
	Set_R_FUNCTION(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_list_function_parameters returns the _list_function_parameters rule contexts.
	Get_list_function_parameters() IList_function_parametersContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_instr_tipo returns the _instr_tipo rule contexts.
	Get_instr_tipo() IInstr_tipoContext

	// Set_list_function_parameters sets the _list_function_parameters rule contexts.
	Set_list_function_parameters(IList_function_parametersContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_instr_tipo sets the _instr_tipo rule contexts.
	Set_instr_tipo(IInstr_tipoContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_funcContext differentiates from other interfaces.
	IsInstr_funcContext()
}

type Instr_funcContext struct {
	*antlr.BaseParserRuleContext
	parser                    antlr.Parser
	instr                     interfaces.Instruction
	_R_FUNCTION               antlr.Token
	_ID                       antlr.Token
	_list_function_parameters IList_function_parametersContext
	_instrucciones            IInstruccionesContext
	_instr_tipo               IInstr_tipoContext
}

func NewEmptyInstr_funcContext() *Instr_funcContext {
	var p = new(Instr_funcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_func
	return p
}

func (*Instr_funcContext) IsInstr_funcContext() {}

func NewInstr_funcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_funcContext {
	var p = new(Instr_funcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_func

	return p
}

func (s *Instr_funcContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_funcContext) Get_R_FUNCTION() antlr.Token { return s._R_FUNCTION }

func (s *Instr_funcContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_funcContext) Set_R_FUNCTION(v antlr.Token) { s._R_FUNCTION = v }

func (s *Instr_funcContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_funcContext) Get_list_function_parameters() IList_function_parametersContext {
	return s._list_function_parameters
}

func (s *Instr_funcContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_funcContext) Get_instr_tipo() IInstr_tipoContext { return s._instr_tipo }

func (s *Instr_funcContext) Set_list_function_parameters(v IList_function_parametersContext) {
	s._list_function_parameters = v
}

func (s *Instr_funcContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_funcContext) Set_instr_tipo(v IInstr_tipoContext) { s._instr_tipo = v }

func (s *Instr_funcContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_funcContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_funcContext) R_FUNCTION() antlr.TerminalNode {
	return s.GetToken(ChemsR_FUNCTION, 0)
}

func (s *Instr_funcContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_funcContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_funcContext) List_function_parameters() IList_function_parametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_function_parametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_function_parametersContext)
}

func (s *Instr_funcContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_funcContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_funcContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_funcContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_funcContext) TK_MENOSMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENOSMAYOR, 0)
}

func (s *Instr_funcContext) Instr_tipo() IInstr_tipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_tipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_tipoContext)
}

func (s *Instr_funcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_funcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_funcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_func(s)
	}
}

func (s *Instr_funcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_func(s)
	}
}

func (p *Chems) Instr_func() (localctx IInstr_funcContext) {
	localctx = NewInstr_funcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ChemsRULE_instr_func)

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

	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(485)

			var _m = p.Match(ChemsR_FUNCTION)

			localctx.(*Instr_funcContext)._R_FUNCTION = _m
		}
		{
			p.SetState(486)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_funcContext)._ID = _m
		}
		{
			p.SetState(487)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(488)

			var _x = p.List_function_parameters()

			localctx.(*Instr_funcContext)._list_function_parameters = _x
		}
		{
			p.SetState(489)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(490)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(491)

			var _x = p.Instrucciones()

			localctx.(*Instr_funcContext)._instrucciones = _x
		}
		{
			p.SetState(492)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_funcContext).instr = function.NewFunction((func() string {
			if localctx.(*Instr_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_funcContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_funcContext).Get_list_function_parameters().GetL(), localctx.(*Instr_funcContext).Get_instrucciones().GetL(), interfaces.NULL, (func() int {
			if localctx.(*Instr_funcContext).Get_R_FUNCTION() == nil {
				return 0
			} else {
				return localctx.(*Instr_funcContext).Get_R_FUNCTION().GetLine()
			}
		}()), localctx.(*Instr_funcContext).Get_R_FUNCTION().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(495)

			var _m = p.Match(ChemsR_FUNCTION)

			localctx.(*Instr_funcContext)._R_FUNCTION = _m
		}
		{
			p.SetState(496)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_funcContext)._ID = _m
		}
		{
			p.SetState(497)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(498)

			var _x = p.List_function_parameters()

			localctx.(*Instr_funcContext)._list_function_parameters = _x
		}
		{
			p.SetState(499)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(500)
			p.Match(ChemsTK_MENOSMAYOR)
		}
		{
			p.SetState(501)

			var _x = p.Instr_tipo()

			localctx.(*Instr_funcContext)._instr_tipo = _x
		}
		{
			p.SetState(502)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(503)

			var _x = p.Instrucciones()

			localctx.(*Instr_funcContext)._instrucciones = _x
		}
		{
			p.SetState(504)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_funcContext).instr = function.NewFunction((func() string {
			if localctx.(*Instr_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_funcContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_funcContext).Get_list_function_parameters().GetL(), localctx.(*Instr_funcContext).Get_instrucciones().GetL(), localctx.(*Instr_funcContext).Get_instr_tipo().GetTipo_exp(), (func() int {
			if localctx.(*Instr_funcContext).Get_R_FUNCTION() == nil {
				return 0
			} else {
				return localctx.(*Instr_funcContext).Get_R_FUNCTION().GetLine()
			}
		}()), localctx.(*Instr_funcContext).Get_R_FUNCTION().GetColumn())

	}

	return localctx
}

// IList_function_parametersContext is an interface to support dynamic dispatch.
type IList_function_parametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_parameters_fn returns the _block_parameters_fn rule contexts.
	Get_block_parameters_fn() IBlock_parameters_fnContext

	// Set_block_parameters_fn sets the _block_parameters_fn rule contexts.
	Set_block_parameters_fn(IBlock_parameters_fnContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_parameters_fnContext

	// SetE sets the e rule context list.
	SetE([]IBlock_parameters_fnContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_function_parametersContext differentiates from other interfaces.
	IsList_function_parametersContext()
}

type List_function_parametersContext struct {
	*antlr.BaseParserRuleContext
	parser               antlr.Parser
	l                    *arrayList.List
	_block_parameters_fn IBlock_parameters_fnContext
	e                    []IBlock_parameters_fnContext
}

func NewEmptyList_function_parametersContext() *List_function_parametersContext {
	var p = new(List_function_parametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_function_parameters
	return p
}

func (*List_function_parametersContext) IsList_function_parametersContext() {}

func NewList_function_parametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_function_parametersContext {
	var p = new(List_function_parametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_function_parameters

	return p
}

func (s *List_function_parametersContext) GetParser() antlr.Parser { return s.parser }

func (s *List_function_parametersContext) Get_block_parameters_fn() IBlock_parameters_fnContext {
	return s._block_parameters_fn
}

func (s *List_function_parametersContext) Set_block_parameters_fn(v IBlock_parameters_fnContext) {
	s._block_parameters_fn = v
}

func (s *List_function_parametersContext) GetE() []IBlock_parameters_fnContext { return s.e }

func (s *List_function_parametersContext) SetE(v []IBlock_parameters_fnContext) { s.e = v }

func (s *List_function_parametersContext) GetL() *arrayList.List { return s.l }

func (s *List_function_parametersContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_function_parametersContext) AllBlock_parameters_fn() []IBlock_parameters_fnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlock_parameters_fnContext)(nil)).Elem())
	var tst = make([]IBlock_parameters_fnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlock_parameters_fnContext)
		}
	}

	return tst
}

func (s *List_function_parametersContext) Block_parameters_fn(i int) IBlock_parameters_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_parameters_fnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlock_parameters_fnContext)
}

func (s *List_function_parametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_function_parametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_function_parametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_function_parameters(s)
	}
}

func (s *List_function_parametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_function_parameters(s)
	}
}

func (p *Chems) List_function_parameters() (localctx IList_function_parametersContext) {
	localctx = NewList_function_parametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ChemsRULE_list_function_parameters)

	localctx.(*List_function_parametersContext).l = arrayList.New()

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
	p.SetState(510)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ChemsID {
		{
			p.SetState(509)

			var _x = p.Block_parameters_fn()

			localctx.(*List_function_parametersContext)._block_parameters_fn = _x
		}
		localctx.(*List_function_parametersContext).e = append(localctx.(*List_function_parametersContext).e, localctx.(*List_function_parametersContext)._block_parameters_fn)

		p.SetState(512)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_function_parametersContext).GetE()
	for _, e := range listInt {
		localctx.(*List_function_parametersContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IBlock_parameters_fnContext is an interface to support dynamic dispatch.
type IBlock_parameters_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_instr_tipo returns the _instr_tipo rule contexts.
	Get_instr_tipo() IInstr_tipoContext

	// Set_instr_tipo sets the _instr_tipo rule contexts.
	Set_instr_tipo(IInstr_tipoContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsBlock_parameters_fnContext differentiates from other interfaces.
	IsBlock_parameters_fnContext()
}

type Block_parameters_fnContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruction
	_ID         antlr.Token
	_instr_tipo IInstr_tipoContext
}

func NewEmptyBlock_parameters_fnContext() *Block_parameters_fnContext {
	var p = new(Block_parameters_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_parameters_fn
	return p
}

func (*Block_parameters_fnContext) IsBlock_parameters_fnContext() {}

func NewBlock_parameters_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_parameters_fnContext {
	var p = new(Block_parameters_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_parameters_fn

	return p
}

func (s *Block_parameters_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_parameters_fnContext) Get_ID() antlr.Token { return s._ID }

func (s *Block_parameters_fnContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Block_parameters_fnContext) Get_instr_tipo() IInstr_tipoContext { return s._instr_tipo }

func (s *Block_parameters_fnContext) Set_instr_tipo(v IInstr_tipoContext) { s._instr_tipo = v }

func (s *Block_parameters_fnContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Block_parameters_fnContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Block_parameters_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Block_parameters_fnContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DOSPUNTOS, 0)
}

func (s *Block_parameters_fnContext) Instr_tipo() IInstr_tipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_tipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_tipoContext)
}

func (s *Block_parameters_fnContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Block_parameters_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_parameters_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_parameters_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_parameters_fn(s)
	}
}

func (s *Block_parameters_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_parameters_fn(s)
	}
}

func (p *Chems) Block_parameters_fn() (localctx IBlock_parameters_fnContext) {
	localctx = NewBlock_parameters_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ChemsRULE_block_parameters_fn)

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

	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(516)

			var _m = p.Match(ChemsID)

			localctx.(*Block_parameters_fnContext)._ID = _m
		}
		{
			p.SetState(517)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(518)

			var _x = p.Instr_tipo()

			localctx.(*Block_parameters_fnContext)._instr_tipo = _x
		}
		{
			p.SetState(519)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Block_parameters_fnContext).instr = function.NewListParam((func() string {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetText()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_instr_tipo().GetTipo_exp())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(522)

			var _m = p.Match(ChemsID)

			localctx.(*Block_parameters_fnContext)._ID = _m
		}
		{
			p.SetState(523)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(524)

			var _x = p.Instr_tipo()

			localctx.(*Block_parameters_fnContext)._instr_tipo = _x
		}
		localctx.(*Block_parameters_fnContext).instr = function.NewListParam((func() string {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetText()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_instr_tipo().GetTipo_exp())

	}

	return localctx
}

// IInstr_llamadaContext is an interface to support dynamic dispatch.
type IInstr_llamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_list_expression returns the _list_expression rule contexts.
	Get_list_expression() IList_expressionContext

	// Set_list_expression sets the _list_expression rule contexts.
	Set_list_expression(IList_expressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_llamadaContext differentiates from other interfaces.
	IsInstr_llamadaContext()
}

type Instr_llamadaContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	instr            interfaces.Instruction
	_ID              antlr.Token
	_list_expression IList_expressionContext
}

func NewEmptyInstr_llamadaContext() *Instr_llamadaContext {
	var p = new(Instr_llamadaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_llamada
	return p
}

func (*Instr_llamadaContext) IsInstr_llamadaContext() {}

func NewInstr_llamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_llamadaContext {
	var p = new(Instr_llamadaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_llamada

	return p
}

func (s *Instr_llamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_llamadaContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_llamadaContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_llamadaContext) Get_list_expression() IList_expressionContext {
	return s._list_expression
}

func (s *Instr_llamadaContext) Set_list_expression(v IList_expressionContext) { s._list_expression = v }

func (s *Instr_llamadaContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_llamadaContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_llamadaContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_llamadaContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_llamadaContext) List_expression() IList_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expressionContext)
}

func (s *Instr_llamadaContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_llamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_llamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_llamadaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_llamada(s)
	}
}

func (s *Instr_llamadaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_llamada(s)
	}
}

func (p *Chems) Instr_llamada() (localctx IInstr_llamadaContext) {
	localctx = NewInstr_llamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ChemsRULE_instr_llamada)

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
		p.SetState(529)

		var _m = p.Match(ChemsID)

		localctx.(*Instr_llamadaContext)._ID = _m
	}
	{
		p.SetState(530)
		p.Match(ChemsTK_PARA)
	}
	{
		p.SetState(531)

		var _x = p.List_expression()

		localctx.(*Instr_llamadaContext)._list_expression = _x
	}
	{
		p.SetState(532)
		p.Match(ChemsTK_PARC)
	}
	localctx.(*Instr_llamadaContext).instr = function.NewLlamada((func() string {
		if localctx.(*Instr_llamadaContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Instr_llamadaContext).Get_ID().GetText()
		}
	}()), localctx.(*Instr_llamadaContext).Get_list_expression().GetL(), (func() int {
		if localctx.(*Instr_llamadaContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Instr_llamadaContext).Get_ID().GetLine()
		}
	}()), localctx.(*Instr_llamadaContext).Get_ID().GetColumn())

	return localctx
}

// IInstr_llamada_expreContext is an interface to support dynamic dispatch.
type IInstr_llamada_expreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_list_expression returns the _list_expression rule contexts.
	Get_list_expression() IList_expressionContext

	// Set_list_expression sets the _list_expression rule contexts.
	Set_list_expression(IList_expressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Expresion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Expresion)

	// IsInstr_llamada_expreContext differentiates from other interfaces.
	IsInstr_llamada_expreContext()
}

type Instr_llamada_expreContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	instr            interfaces.Expresion
	_ID              antlr.Token
	_list_expression IList_expressionContext
}

func NewEmptyInstr_llamada_expreContext() *Instr_llamada_expreContext {
	var p = new(Instr_llamada_expreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_llamada_expre
	return p
}

func (*Instr_llamada_expreContext) IsInstr_llamada_expreContext() {}

func NewInstr_llamada_expreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_llamada_expreContext {
	var p = new(Instr_llamada_expreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_llamada_expre

	return p
}

func (s *Instr_llamada_expreContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_llamada_expreContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_llamada_expreContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_llamada_expreContext) Get_list_expression() IList_expressionContext {
	return s._list_expression
}

func (s *Instr_llamada_expreContext) Set_list_expression(v IList_expressionContext) {
	s._list_expression = v
}

func (s *Instr_llamada_expreContext) GetInstr() interfaces.Expresion { return s.instr }

func (s *Instr_llamada_expreContext) SetInstr(v interfaces.Expresion) { s.instr = v }

func (s *Instr_llamada_expreContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_llamada_expreContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_llamada_expreContext) List_expression() IList_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expressionContext)
}

func (s *Instr_llamada_expreContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_llamada_expreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_llamada_expreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_llamada_expreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_llamada_expre(s)
	}
}

func (s *Instr_llamada_expreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_llamada_expre(s)
	}
}

func (p *Chems) Instr_llamada_expre() (localctx IInstr_llamada_expreContext) {
	localctx = NewInstr_llamada_expreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ChemsRULE_instr_llamada_expre)

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
		p.SetState(535)

		var _m = p.Match(ChemsID)

		localctx.(*Instr_llamada_expreContext)._ID = _m
	}
	{
		p.SetState(536)
		p.Match(ChemsTK_PARA)
	}
	{
		p.SetState(537)

		var _x = p.List_expression()

		localctx.(*Instr_llamada_expreContext)._list_expression = _x
	}
	{
		p.SetState(538)
		p.Match(ChemsTK_PARC)
	}
	localctx.(*Instr_llamada_expreContext).instr = function.NewLlamadaExpre((func() string {
		if localctx.(*Instr_llamada_expreContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Instr_llamada_expreContext).Get_ID().GetText()
		}
	}()), localctx.(*Instr_llamada_expreContext).Get_list_expression().GetL(), (func() int {
		if localctx.(*Instr_llamada_expreContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Instr_llamada_expreContext).Get_ID().GetLine()
		}
	}()), localctx.(*Instr_llamada_expreContext).Get_ID().GetColumn())

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
	p.EnterRule(localctx, 76, ChemsRULE_instr_tipo)

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

	p.SetState(551)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(541)
			p.Match(ChemsR_INT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.INTEGER

	case ChemsR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(543)
			p.Match(ChemsR_FLOAT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.FLOAT

	case ChemsR_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(545)
			p.Match(ChemsR_STRING)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STRING

	case ChemsR_STR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(547)
			p.Match(ChemsR_STR)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STRING

	case ChemsR_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(549)
			p.Match(ChemsR_BOOL)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.BOOLEAN

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IList_expressionContext is an interface to support dynamic dispatch.
type IList_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_list_expression returns the _block_list_expression rule contexts.
	Get_block_list_expression() IBlock_list_expressionContext

	// Set_block_list_expression sets the _block_list_expression rule contexts.
	Set_block_list_expression(IBlock_list_expressionContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_list_expressionContext

	// SetE sets the e rule context list.
	SetE([]IBlock_list_expressionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_expressionContext differentiates from other interfaces.
	IsList_expressionContext()
}

type List_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser                 antlr.Parser
	l                      *arrayList.List
	_block_list_expression IBlock_list_expressionContext
	e                      []IBlock_list_expressionContext
}

func NewEmptyList_expressionContext() *List_expressionContext {
	var p = new(List_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_expression
	return p
}

func (*List_expressionContext) IsList_expressionContext() {}

func NewList_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expressionContext {
	var p = new(List_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_expression

	return p
}

func (s *List_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expressionContext) Get_block_list_expression() IBlock_list_expressionContext {
	return s._block_list_expression
}

func (s *List_expressionContext) Set_block_list_expression(v IBlock_list_expressionContext) {
	s._block_list_expression = v
}

func (s *List_expressionContext) GetE() []IBlock_list_expressionContext { return s.e }

func (s *List_expressionContext) SetE(v []IBlock_list_expressionContext) { s.e = v }

func (s *List_expressionContext) GetL() *arrayList.List { return s.l }

func (s *List_expressionContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_expressionContext) AllBlock_list_expression() []IBlock_list_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlock_list_expressionContext)(nil)).Elem())
	var tst = make([]IBlock_list_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlock_list_expressionContext)
		}
	}

	return tst
}

func (s *List_expressionContext) Block_list_expression(i int) IBlock_list_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_list_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlock_list_expressionContext)
}

func (s *List_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_expression(s)
	}
}

func (s *List_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_expression(s)
	}
}

func (p *Chems) List_expression() (localctx IList_expressionContext) {
	localctx = NewList_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ChemsRULE_list_expression)

	localctx.(*List_expressionContext).l = arrayList.New()

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
	p.SetState(554)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-49)&-(0x1f+1)) == 0 && ((1<<uint((_la-49)))&((1<<(ChemsTK_MENOS-49))|(1<<(ChemsTK_PARA-49))|(1<<(ChemsTK_NOT-49)))) != 0) {
		{
			p.SetState(553)

			var _x = p.Block_list_expression()

			localctx.(*List_expressionContext)._block_list_expression = _x
		}
		localctx.(*List_expressionContext).e = append(localctx.(*List_expressionContext).e, localctx.(*List_expressionContext)._block_list_expression)

		p.SetState(556)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_expressionContext).GetE()
	for _, e := range listInt {
		localctx.(*List_expressionContext).l.Add(e.GetP())
	}

	return localctx
}

// IBlock_list_expressionContext is an interface to support dynamic dispatch.
type IBlock_list_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsBlock_list_expressionContext differentiates from other interfaces.
	IsBlock_list_expressionContext()
}

type Block_list_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Expresion
	_expression IExpressionContext
}

func NewEmptyBlock_list_expressionContext() *Block_list_expressionContext {
	var p = new(Block_list_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_list_expression
	return p
}

func (*Block_list_expressionContext) IsBlock_list_expressionContext() {}

func NewBlock_list_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_list_expressionContext {
	var p = new(Block_list_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_list_expression

	return p
}

func (s *Block_list_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_list_expressionContext) Get_expression() IExpressionContext { return s._expression }

func (s *Block_list_expressionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Block_list_expressionContext) GetP() interfaces.Expresion { return s.p }

func (s *Block_list_expressionContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *Block_list_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Block_list_expressionContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Block_list_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_list_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_list_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_list_expression(s)
	}
}

func (s *Block_list_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_list_expression(s)
	}
}

func (p *Chems) Block_list_expression() (localctx IBlock_list_expressionContext) {
	localctx = NewBlock_list_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ChemsRULE_block_list_expression)

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

	p.SetState(567)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(560)

			var _x = p.Expression()

			localctx.(*Block_list_expressionContext)._expression = _x
		}
		{
			p.SetState(561)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Block_list_expressionContext).p = instruction.NewListExpre(localctx.(*Block_list_expressionContext).Get_expression().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(564)

			var _x = p.Expression()

			localctx.(*Block_list_expressionContext)._expression = _x
		}
		localctx.(*Block_list_expressionContext).p = instruction.NewListExpre(localctx.(*Block_list_expressionContext).Get_expression().GetP())

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
	p.EnterRule(localctx, 82, ChemsRULE_expression)

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
		p.SetState(569)

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
	_startState := 84
	p.EnterRecursionRule(localctx, 84, ChemsRULE_exp_arit, _p)
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
	p.SetState(649)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(573)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(574)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(575)

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
			p.SetState(576)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(577)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(ChemsTK_MULT-45))|(1<<(ChemsTK_DIV-45))|(1<<(ChemsTK_MODULO-45)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(578)

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
			p.SetState(581)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(582)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(583)

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
			p.SetState(584)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(585)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(ChemsTK_MULT-45))|(1<<(ChemsTK_DIV-45))|(1<<(ChemsTK_MODULO-45)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(586)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(587)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(588)

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
			p.SetState(589)
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
			p.SetState(592)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(593)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(594)

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
			p.SetState(595)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(596)

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
			p.SetState(597)

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
			p.SetState(600)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(601)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(602)

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
			p.SetState(603)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(604)

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
			p.SetState(605)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(606)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(607)

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
			p.SetState(608)
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
			p.SetState(611)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(612)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(613)

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
			p.SetState(614)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(615)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ChemsTK_IGUALIGUAL-37))|(1<<(ChemsTK_MAYORIGUAL-37))|(1<<(ChemsTK_MENORIGUAL-37))|(1<<(ChemsTK_DIFIGUAL-37))|(1<<(ChemsTK_MAYOR-37))|(1<<(ChemsTK_MENOR-37)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(616)

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
			p.SetState(619)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(620)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(621)

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
			p.SetState(622)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(623)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ChemsTK_IGUALIGUAL-37))|(1<<(ChemsTK_MAYORIGUAL-37))|(1<<(ChemsTK_MENORIGUAL-37))|(1<<(ChemsTK_DIFIGUAL-37))|(1<<(ChemsTK_MAYOR-37))|(1<<(ChemsTK_MENOR-37)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(624)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(625)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(626)

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
			p.SetState(627)
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
			p.SetState(630)

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
			p.SetState(631)

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
			p.SetState(634)

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
			p.SetState(635)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(636)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(637)

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
			p.SetState(638)
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
			p.SetState(641)

			var _x = p.Primitivo()

			localctx.(*Exp_aritContext)._primitivo = _x
		}
		localctx.(*Exp_aritContext).p = localctx.(*Exp_aritContext).Get_primitivo().GetP()

	case 10:
		{
			p.SetState(644)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(645)

			var _x = p.Expression()

			localctx.(*Exp_aritContext)._expression = _x
		}
		{
			p.SetState(646)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Exp_aritContext).p = localctx.(*Exp_aritContext).Get_expression().GetP()

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(697)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(695)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(651)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(652)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(ChemsTK_MULT-45))|(1<<(ChemsTK_DIV-45))|(1<<(ChemsTK_MODULO-45)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(653)

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
				p.SetState(656)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(657)

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
					p.SetState(658)

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
				p.SetState(661)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(662)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ChemsTK_IGUALIGUAL-37))|(1<<(ChemsTK_MAYORIGUAL-37))|(1<<(ChemsTK_MENORIGUAL-37))|(1<<(ChemsTK_DIFIGUAL-37))|(1<<(ChemsTK_MAYOR-37))|(1<<(ChemsTK_MENOR-37)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(663)

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
				p.SetState(666)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(667)

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
					p.SetState(668)

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
				p.SetState(671)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(672)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(ChemsTK_MULT-45))|(1<<(ChemsTK_DIV-45))|(1<<(ChemsTK_MODULO-45)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(673)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(674)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(675)

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
					p.SetState(676)
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
				p.SetState(679)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(680)

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
					p.SetState(681)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(682)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(683)

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
					p.SetState(684)
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
				p.SetState(687)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(688)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ChemsTK_IGUALIGUAL-37))|(1<<(ChemsTK_MAYORIGUAL-37))|(1<<(ChemsTK_MENORIGUAL-37))|(1<<(ChemsTK_DIFIGUAL-37))|(1<<(ChemsTK_MAYOR-37))|(1<<(ChemsTK_MENOR-37)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(689)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(690)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(691)

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
					p.SetState(692)
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
		p.SetState(699)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
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

	// Get_instr_llamada_expre returns the _instr_llamada_expre rule contexts.
	Get_instr_llamada_expre() IInstr_llamada_expreContext

	// Set_instr_ternario sets the _instr_ternario rule contexts.
	Set_instr_ternario(IInstr_ternarioContext)

	// Set_instr_match_ter sets the _instr_match_ter rule contexts.
	Set_instr_match_ter(IInstr_match_terContext)

	// Set_instr_loop_ternario sets the _instr_loop_ternario rule contexts.
	Set_instr_loop_ternario(IInstr_loop_ternarioContext)

	// Set_instr_llamada_expre sets the _instr_llamada_expre rule contexts.
	Set_instr_llamada_expre(IInstr_llamada_expreContext)

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
	_instr_llamada_expre IInstr_llamada_expreContext
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

func (s *PrimitivoContext) Get_instr_llamada_expre() IInstr_llamada_expreContext {
	return s._instr_llamada_expre
}

func (s *PrimitivoContext) Set_instr_ternario(v IInstr_ternarioContext) { s._instr_ternario = v }

func (s *PrimitivoContext) Set_instr_match_ter(v IInstr_match_terContext) { s._instr_match_ter = v }

func (s *PrimitivoContext) Set_instr_loop_ternario(v IInstr_loop_ternarioContext) {
	s._instr_loop_ternario = v
}

func (s *PrimitivoContext) Set_instr_llamada_expre(v IInstr_llamada_expreContext) {
	s._instr_llamada_expre = v
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

func (s *PrimitivoContext) Instr_llamada_expre() IInstr_llamada_expreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_llamada_expreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_llamada_expreContext)
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
	p.EnterRule(localctx, 86, ChemsRULE_primitivo)

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

	p.SetState(722)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(700)

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

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(702)

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

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(704)

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

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(706)

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

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(708)

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

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(710)

			var _x = p.Instr_ternario()

			localctx.(*PrimitivoContext)._instr_ternario = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_ternario().GetP()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(713)

			var _x = p.Instr_match_ter()

			localctx.(*PrimitivoContext)._instr_match_ter = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_match_ter().GetInstr()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(716)

			var _x = p.Instr_loop_ternario()

			localctx.(*PrimitivoContext)._instr_loop_ternario = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_loop_ternario().GetInstr()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(719)

			var _x = p.Instr_llamada_expre()

			localctx.(*PrimitivoContext)._instr_llamada_expre = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_llamada_expre().GetInstr()

	}

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 42:
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

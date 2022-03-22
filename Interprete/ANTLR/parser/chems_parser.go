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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 58, 584,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 3, 2, 3, 2, 3, 3, 6, 3, 77, 10,
	3, 13, 3, 14, 3, 78, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 86, 10, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 5, 5, 116, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 134,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7,
	170, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9,
	205, 10, 9, 3, 10, 7, 10, 208, 10, 10, 12, 10, 14, 10, 211, 11, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 242, 10, 11,
	3, 12, 7, 12, 245, 10, 12, 12, 12, 14, 12, 248, 11, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 267, 10, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	282, 10, 14, 3, 15, 3, 15, 3, 15, 3, 16, 6, 16, 288, 10, 16, 13, 16, 14,
	16, 289, 3, 16, 3, 16, 3, 17, 6, 17, 295, 10, 17, 13, 17, 14, 17, 296,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 308,
	10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 5, 19, 321, 10, 19, 3, 20, 6, 20, 324, 10, 20, 13, 20, 14,
	20, 325, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 345, 10,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 6, 23, 354, 10, 23,
	13, 23, 14, 23, 355, 3, 23, 3, 23, 3, 24, 6, 24, 361, 10, 24, 13, 24, 14,
	24, 362, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	5, 25, 374, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 409, 10, 30, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 430, 10, 33, 3, 34, 3, 34,
	3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 512, 10, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 7, 35,
	558, 10, 35, 12, 35, 14, 35, 561, 11, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 582, 10, 36, 3, 36, 2, 3, 68, 37, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 2, 8, 3, 2,
	5, 6, 3, 2, 40, 42, 3, 2, 43, 44, 4, 2, 33, 34, 36, 39, 4, 2, 44, 44, 54,
	54, 3, 2, 51, 52, 2, 607, 2, 72, 3, 2, 2, 2, 4, 76, 3, 2, 2, 2, 6, 85,
	3, 2, 2, 2, 8, 115, 3, 2, 2, 2, 10, 133, 3, 2, 2, 2, 12, 169, 3, 2, 2,
	2, 14, 171, 3, 2, 2, 2, 16, 204, 3, 2, 2, 2, 18, 209, 3, 2, 2, 2, 20, 241,
	3, 2, 2, 2, 22, 246, 3, 2, 2, 2, 24, 266, 3, 2, 2, 2, 26, 281, 3, 2, 2,
	2, 28, 283, 3, 2, 2, 2, 30, 287, 3, 2, 2, 2, 32, 294, 3, 2, 2, 2, 34, 307,
	3, 2, 2, 2, 36, 320, 3, 2, 2, 2, 38, 323, 3, 2, 2, 2, 40, 344, 3, 2, 2,
	2, 42, 346, 3, 2, 2, 2, 44, 353, 3, 2, 2, 2, 46, 360, 3, 2, 2, 2, 48, 373,
	3, 2, 2, 2, 50, 375, 3, 2, 2, 2, 52, 380, 3, 2, 2, 2, 54, 387, 3, 2, 2,
	2, 56, 393, 3, 2, 2, 2, 58, 408, 3, 2, 2, 2, 60, 410, 3, 2, 2, 2, 62, 414,
	3, 2, 2, 2, 64, 429, 3, 2, 2, 2, 66, 431, 3, 2, 2, 2, 68, 511, 3, 2, 2,
	2, 70, 581, 3, 2, 2, 2, 72, 73, 5, 4, 3, 2, 73, 74, 8, 2, 1, 2, 74, 3,
	3, 2, 2, 2, 75, 77, 5, 8, 5, 2, 76, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2,
	78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81, 8,
	3, 1, 2, 81, 5, 3, 2, 2, 2, 82, 83, 7, 29, 2, 2, 83, 86, 8, 4, 1, 2, 84,
	86, 8, 4, 1, 2, 85, 82, 3, 2, 2, 2, 85, 84, 3, 2, 2, 2, 86, 7, 3, 2, 2,
	2, 87, 88, 5, 10, 6, 2, 88, 89, 5, 6, 4, 2, 89, 90, 8, 5, 1, 2, 90, 116,
	3, 2, 2, 2, 91, 92, 5, 12, 7, 2, 92, 93, 8, 5, 1, 2, 93, 116, 3, 2, 2,
	2, 94, 95, 5, 14, 8, 2, 95, 96, 8, 5, 1, 2, 96, 116, 3, 2, 2, 2, 97, 98,
	5, 16, 9, 2, 98, 99, 8, 5, 1, 2, 99, 116, 3, 2, 2, 2, 100, 101, 5, 24,
	13, 2, 101, 102, 8, 5, 1, 2, 102, 116, 3, 2, 2, 2, 103, 104, 5, 52, 27,
	2, 104, 105, 8, 5, 1, 2, 105, 116, 3, 2, 2, 2, 106, 107, 5, 58, 30, 2,
	107, 108, 8, 5, 1, 2, 108, 116, 3, 2, 2, 2, 109, 110, 5, 60, 31, 2, 110,
	111, 8, 5, 1, 2, 111, 116, 3, 2, 2, 2, 112, 113, 5, 54, 28, 2, 113, 114,
	8, 5, 1, 2, 114, 116, 3, 2, 2, 2, 115, 87, 3, 2, 2, 2, 115, 91, 3, 2, 2,
	2, 115, 94, 3, 2, 2, 2, 115, 97, 3, 2, 2, 2, 115, 100, 3, 2, 2, 2, 115,
	103, 3, 2, 2, 2, 115, 106, 3, 2, 2, 2, 115, 109, 3, 2, 2, 2, 115, 112,
	3, 2, 2, 2, 116, 9, 3, 2, 2, 2, 117, 118, 7, 3, 2, 2, 118, 119, 7, 45,
	2, 2, 119, 120, 5, 66, 34, 2, 120, 121, 7, 46, 2, 2, 121, 122, 7, 29, 2,
	2, 122, 123, 8, 6, 1, 2, 123, 134, 3, 2, 2, 2, 124, 125, 7, 3, 2, 2, 125,
	126, 7, 45, 2, 2, 126, 127, 7, 25, 2, 2, 127, 128, 7, 30, 2, 2, 128, 129,
	5, 66, 34, 2, 129, 130, 7, 46, 2, 2, 130, 131, 7, 29, 2, 2, 131, 132, 8,
	6, 1, 2, 132, 134, 3, 2, 2, 2, 133, 117, 3, 2, 2, 2, 133, 124, 3, 2, 2,
	2, 134, 11, 3, 2, 2, 2, 135, 136, 7, 7, 2, 2, 136, 137, 7, 8, 2, 2, 137,
	138, 7, 27, 2, 2, 138, 139, 7, 32, 2, 2, 139, 140, 5, 66, 34, 2, 140, 141,
	7, 29, 2, 2, 141, 142, 8, 7, 1, 2, 142, 170, 3, 2, 2, 2, 143, 144, 7, 7,
	2, 2, 144, 145, 7, 8, 2, 2, 145, 146, 7, 27, 2, 2, 146, 147, 7, 31, 2,
	2, 147, 148, 5, 64, 33, 2, 148, 149, 7, 32, 2, 2, 149, 150, 5, 66, 34,
	2, 150, 151, 7, 29, 2, 2, 151, 152, 8, 7, 1, 2, 152, 170, 3, 2, 2, 2, 153,
	154, 7, 7, 2, 2, 154, 155, 7, 27, 2, 2, 155, 156, 7, 32, 2, 2, 156, 157,
	5, 66, 34, 2, 157, 158, 7, 29, 2, 2, 158, 159, 8, 7, 1, 2, 159, 170, 3,
	2, 2, 2, 160, 161, 7, 7, 2, 2, 161, 162, 7, 27, 2, 2, 162, 163, 7, 31,
	2, 2, 163, 164, 5, 64, 33, 2, 164, 165, 7, 32, 2, 2, 165, 166, 5, 66, 34,
	2, 166, 167, 7, 29, 2, 2, 167, 168, 8, 7, 1, 2, 168, 170, 3, 2, 2, 2, 169,
	135, 3, 2, 2, 2, 169, 143, 3, 2, 2, 2, 169, 153, 3, 2, 2, 2, 169, 160,
	3, 2, 2, 2, 170, 13, 3, 2, 2, 2, 171, 172, 7, 27, 2, 2, 172, 173, 7, 32,
	2, 2, 173, 174, 5, 66, 34, 2, 174, 175, 7, 29, 2, 2, 175, 176, 8, 8, 1,
	2, 176, 15, 3, 2, 2, 2, 177, 178, 7, 9, 2, 2, 178, 179, 5, 66, 34, 2, 179,
	180, 7, 47, 2, 2, 180, 181, 5, 4, 3, 2, 181, 182, 7, 48, 2, 2, 182, 183,
	8, 9, 1, 2, 183, 205, 3, 2, 2, 2, 184, 185, 7, 9, 2, 2, 185, 186, 5, 66,
	34, 2, 186, 187, 7, 47, 2, 2, 187, 188, 5, 4, 3, 2, 188, 189, 7, 48, 2,
	2, 189, 190, 7, 10, 2, 2, 190, 191, 7, 47, 2, 2, 191, 192, 5, 4, 3, 2,
	192, 193, 7, 48, 2, 2, 193, 194, 8, 9, 1, 2, 194, 205, 3, 2, 2, 2, 195,
	196, 7, 9, 2, 2, 196, 197, 5, 66, 34, 2, 197, 198, 7, 47, 2, 2, 198, 199,
	5, 4, 3, 2, 199, 200, 7, 48, 2, 2, 200, 201, 7, 10, 2, 2, 201, 202, 5,
	18, 10, 2, 202, 203, 8, 9, 1, 2, 203, 205, 3, 2, 2, 2, 204, 177, 3, 2,
	2, 2, 204, 184, 3, 2, 2, 2, 204, 195, 3, 2, 2, 2, 205, 17, 3, 2, 2, 2,
	206, 208, 5, 16, 9, 2, 207, 206, 3, 2, 2, 2, 208, 211, 3, 2, 2, 2, 209,
	207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 212, 3, 2, 2, 2, 211, 209,
	3, 2, 2, 2, 212, 213, 8, 10, 1, 2, 213, 19, 3, 2, 2, 2, 214, 215, 7, 9,
	2, 2, 215, 216, 5, 66, 34, 2, 216, 217, 7, 47, 2, 2, 217, 218, 5, 66, 34,
	2, 218, 219, 7, 48, 2, 2, 219, 220, 8, 11, 1, 2, 220, 242, 3, 2, 2, 2,
	221, 222, 7, 9, 2, 2, 222, 223, 5, 66, 34, 2, 223, 224, 7, 47, 2, 2, 224,
	225, 5, 66, 34, 2, 225, 226, 7, 48, 2, 2, 226, 227, 7, 10, 2, 2, 227, 228,
	7, 47, 2, 2, 228, 229, 5, 66, 34, 2, 229, 230, 7, 48, 2, 2, 230, 231, 8,
	11, 1, 2, 231, 242, 3, 2, 2, 2, 232, 233, 7, 9, 2, 2, 233, 234, 5, 66,
	34, 2, 234, 235, 7, 47, 2, 2, 235, 236, 5, 66, 34, 2, 236, 237, 7, 48,
	2, 2, 237, 238, 7, 10, 2, 2, 238, 239, 5, 22, 12, 2, 239, 240, 8, 11, 1,
	2, 240, 242, 3, 2, 2, 2, 241, 214, 3, 2, 2, 2, 241, 221, 3, 2, 2, 2, 241,
	232, 3, 2, 2, 2, 242, 21, 3, 2, 2, 2, 243, 245, 5, 20, 11, 2, 244, 243,
	3, 2, 2, 2, 245, 248, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2, 246, 247, 3, 2,
	2, 2, 247, 249, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 249, 250, 8, 12, 1, 2,
	250, 23, 3, 2, 2, 2, 251, 252, 7, 11, 2, 2, 252, 253, 5, 66, 34, 2, 253,
	254, 7, 47, 2, 2, 254, 255, 5, 30, 16, 2, 255, 256, 5, 38, 20, 2, 256,
	257, 7, 48, 2, 2, 257, 258, 8, 13, 1, 2, 258, 267, 3, 2, 2, 2, 259, 260,
	7, 11, 2, 2, 260, 261, 5, 66, 34, 2, 261, 262, 7, 47, 2, 2, 262, 263, 5,
	38, 20, 2, 263, 264, 7, 48, 2, 2, 264, 265, 8, 13, 1, 2, 265, 267, 3, 2,
	2, 2, 266, 251, 3, 2, 2, 2, 266, 259, 3, 2, 2, 2, 267, 25, 3, 2, 2, 2,
	268, 269, 5, 32, 17, 2, 269, 270, 7, 35, 2, 2, 270, 271, 7, 47, 2, 2, 271,
	272, 5, 4, 3, 2, 272, 273, 7, 48, 2, 2, 273, 274, 8, 14, 1, 2, 274, 282,
	3, 2, 2, 2, 275, 276, 5, 32, 17, 2, 276, 277, 7, 35, 2, 2, 277, 278, 5,
	28, 15, 2, 278, 279, 7, 30, 2, 2, 279, 280, 8, 14, 1, 2, 280, 282, 3, 2,
	2, 2, 281, 268, 3, 2, 2, 2, 281, 275, 3, 2, 2, 2, 282, 27, 3, 2, 2, 2,
	283, 284, 5, 8, 5, 2, 284, 285, 8, 15, 1, 2, 285, 29, 3, 2, 2, 2, 286,
	288, 5, 26, 14, 2, 287, 286, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 287,
	3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 292, 8, 16,
	1, 2, 292, 31, 3, 2, 2, 2, 293, 295, 5, 34, 18, 2, 294, 293, 3, 2, 2, 2,
	295, 296, 3, 2, 2, 2, 296, 294, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297,
	298, 3, 2, 2, 2, 298, 299, 8, 17, 1, 2, 299, 33, 3, 2, 2, 2, 300, 301,
	5, 66, 34, 2, 301, 302, 7, 53, 2, 2, 302, 303, 8, 18, 1, 2, 303, 308, 3,
	2, 2, 2, 304, 305, 5, 66, 34, 2, 305, 306, 8, 18, 1, 2, 306, 308, 3, 2,
	2, 2, 307, 300, 3, 2, 2, 2, 307, 304, 3, 2, 2, 2, 308, 35, 3, 2, 2, 2,
	309, 310, 7, 55, 2, 2, 310, 311, 7, 47, 2, 2, 311, 312, 5, 4, 3, 2, 312,
	313, 7, 48, 2, 2, 313, 314, 8, 19, 1, 2, 314, 321, 3, 2, 2, 2, 315, 316,
	7, 55, 2, 2, 316, 317, 5, 28, 15, 2, 317, 318, 7, 30, 2, 2, 318, 319, 8,
	19, 1, 2, 319, 321, 3, 2, 2, 2, 320, 309, 3, 2, 2, 2, 320, 315, 3, 2, 2,
	2, 321, 37, 3, 2, 2, 2, 322, 324, 5, 36, 19, 2, 323, 322, 3, 2, 2, 2, 324,
	325, 3, 2, 2, 2, 325, 323, 3, 2, 2, 2, 325, 326, 3, 2, 2, 2, 326, 327,
	3, 2, 2, 2, 327, 328, 8, 20, 1, 2, 328, 39, 3, 2, 2, 2, 329, 330, 7, 11,
	2, 2, 330, 331, 5, 66, 34, 2, 331, 332, 7, 47, 2, 2, 332, 333, 5, 44, 23,
	2, 333, 334, 5, 50, 26, 2, 334, 335, 7, 48, 2, 2, 335, 336, 8, 21, 1, 2,
	336, 345, 3, 2, 2, 2, 337, 338, 7, 11, 2, 2, 338, 339, 5, 66, 34, 2, 339,
	340, 7, 47, 2, 2, 340, 341, 5, 50, 26, 2, 341, 342, 7, 48, 2, 2, 342, 343,
	8, 21, 1, 2, 343, 345, 3, 2, 2, 2, 344, 329, 3, 2, 2, 2, 344, 337, 3, 2,
	2, 2, 345, 41, 3, 2, 2, 2, 346, 347, 5, 46, 24, 2, 347, 348, 7, 35, 2,
	2, 348, 349, 5, 66, 34, 2, 349, 350, 7, 30, 2, 2, 350, 351, 8, 22, 1, 2,
	351, 43, 3, 2, 2, 2, 352, 354, 5, 42, 22, 2, 353, 352, 3, 2, 2, 2, 354,
	355, 3, 2, 2, 2, 355, 353, 3, 2, 2, 2, 355, 356, 3, 2, 2, 2, 356, 357,
	3, 2, 2, 2, 357, 358, 8, 23, 1, 2, 358, 45, 3, 2, 2, 2, 359, 361, 5, 48,
	25, 2, 360, 359, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2,
	362, 363, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 365, 8, 24, 1, 2, 365,
	47, 3, 2, 2, 2, 366, 367, 5, 66, 34, 2, 367, 368, 7, 53, 2, 2, 368, 369,
	8, 25, 1, 2, 369, 374, 3, 2, 2, 2, 370, 371, 5, 66, 34, 2, 371, 372, 8,
	25, 1, 2, 372, 374, 3, 2, 2, 2, 373, 366, 3, 2, 2, 2, 373, 370, 3, 2, 2,
	2, 374, 49, 3, 2, 2, 2, 375, 376, 7, 55, 2, 2, 376, 377, 5, 66, 34, 2,
	377, 378, 7, 30, 2, 2, 378, 379, 8, 26, 1, 2, 379, 51, 3, 2, 2, 2, 380,
	381, 7, 12, 2, 2, 381, 382, 5, 66, 34, 2, 382, 383, 7, 47, 2, 2, 383, 384,
	5, 4, 3, 2, 384, 385, 7, 48, 2, 2, 385, 386, 8, 27, 1, 2, 386, 53, 3, 2,
	2, 2, 387, 388, 7, 16, 2, 2, 388, 389, 7, 47, 2, 2, 389, 390, 5, 4, 3,
	2, 390, 391, 7, 48, 2, 2, 391, 392, 8, 28, 1, 2, 392, 55, 3, 2, 2, 2, 393,
	394, 7, 16, 2, 2, 394, 395, 7, 47, 2, 2, 395, 396, 5, 4, 3, 2, 396, 397,
	7, 48, 2, 2, 397, 398, 8, 29, 1, 2, 398, 57, 3, 2, 2, 2, 399, 400, 7, 13,
	2, 2, 400, 401, 5, 6, 4, 2, 401, 402, 8, 30, 1, 2, 402, 409, 3, 2, 2, 2,
	403, 404, 7, 13, 2, 2, 404, 405, 5, 66, 34, 2, 405, 406, 5, 6, 4, 2, 406,
	407, 8, 30, 1, 2, 407, 409, 3, 2, 2, 2, 408, 399, 3, 2, 2, 2, 408, 403,
	3, 2, 2, 2, 409, 59, 3, 2, 2, 2, 410, 411, 7, 14, 2, 2, 411, 412, 5, 6,
	4, 2, 412, 413, 8, 31, 1, 2, 413, 61, 3, 2, 2, 2, 414, 415, 7, 15, 2, 2,
	415, 416, 5, 66, 34, 2, 416, 417, 5, 6, 4, 2, 417, 418, 8, 32, 1, 2, 418,
	63, 3, 2, 2, 2, 419, 420, 7, 17, 2, 2, 420, 430, 8, 33, 1, 2, 421, 422,
	7, 18, 2, 2, 422, 430, 8, 33, 1, 2, 423, 424, 7, 19, 2, 2, 424, 430, 8,
	33, 1, 2, 425, 426, 7, 21, 2, 2, 426, 430, 8, 33, 1, 2, 427, 428, 7, 20,
	2, 2, 428, 430, 8, 33, 1, 2, 429, 419, 3, 2, 2, 2, 429, 421, 3, 2, 2, 2,
	429, 423, 3, 2, 2, 2, 429, 425, 3, 2, 2, 2, 429, 427, 3, 2, 2, 2, 430,
	65, 3, 2, 2, 2, 431, 432, 5, 68, 35, 2, 432, 433, 8, 34, 1, 2, 433, 67,
	3, 2, 2, 2, 434, 435, 8, 35, 1, 2, 435, 436, 7, 45, 2, 2, 436, 437, 5,
	68, 35, 2, 437, 438, 9, 2, 2, 2, 438, 439, 7, 46, 2, 2, 439, 440, 9, 3,
	2, 2, 440, 441, 5, 68, 35, 18, 441, 442, 8, 35, 1, 2, 442, 512, 3, 2, 2,
	2, 443, 444, 7, 45, 2, 2, 444, 445, 5, 68, 35, 2, 445, 446, 9, 2, 2, 2,
	446, 447, 7, 46, 2, 2, 447, 448, 9, 3, 2, 2, 448, 449, 7, 45, 2, 2, 449,
	450, 5, 68, 35, 2, 450, 451, 9, 2, 2, 2, 451, 452, 7, 46, 2, 2, 452, 453,
	8, 35, 1, 2, 453, 512, 3, 2, 2, 2, 454, 455, 7, 45, 2, 2, 455, 456, 5,
	68, 35, 2, 456, 457, 9, 2, 2, 2, 457, 458, 7, 46, 2, 2, 458, 459, 9, 4,
	2, 2, 459, 460, 5, 68, 35, 14, 460, 461, 8, 35, 1, 2, 461, 512, 3, 2, 2,
	2, 462, 463, 7, 45, 2, 2, 463, 464, 5, 68, 35, 2, 464, 465, 9, 2, 2, 2,
	465, 466, 7, 46, 2, 2, 466, 467, 9, 4, 2, 2, 467, 468, 7, 45, 2, 2, 468,
	469, 5, 68, 35, 2, 469, 470, 9, 2, 2, 2, 470, 471, 7, 46, 2, 2, 471, 472,
	8, 35, 1, 2, 472, 512, 3, 2, 2, 2, 473, 474, 7, 45, 2, 2, 474, 475, 5,
	68, 35, 2, 475, 476, 9, 2, 2, 2, 476, 477, 7, 46, 2, 2, 477, 478, 9, 5,
	2, 2, 478, 479, 5, 68, 35, 10, 479, 480, 8, 35, 1, 2, 480, 512, 3, 2, 2,
	2, 481, 482, 7, 45, 2, 2, 482, 483, 5, 68, 35, 2, 483, 484, 9, 2, 2, 2,
	484, 485, 7, 46, 2, 2, 485, 486, 9, 5, 2, 2, 486, 487, 7, 45, 2, 2, 487,
	488, 5, 68, 35, 2, 488, 489, 9, 2, 2, 2, 489, 490, 7, 46, 2, 2, 490, 491,
	8, 35, 1, 2, 491, 512, 3, 2, 2, 2, 492, 493, 9, 6, 2, 2, 493, 494, 5, 66,
	34, 2, 494, 495, 8, 35, 1, 2, 495, 512, 3, 2, 2, 2, 496, 497, 9, 6, 2,
	2, 497, 498, 7, 45, 2, 2, 498, 499, 5, 68, 35, 2, 499, 500, 9, 2, 2, 2,
	500, 501, 7, 46, 2, 2, 501, 502, 8, 35, 1, 2, 502, 512, 3, 2, 2, 2, 503,
	504, 5, 70, 36, 2, 504, 505, 8, 35, 1, 2, 505, 512, 3, 2, 2, 2, 506, 507,
	7, 45, 2, 2, 507, 508, 5, 66, 34, 2, 508, 509, 7, 46, 2, 2, 509, 510, 8,
	35, 1, 2, 510, 512, 3, 2, 2, 2, 511, 434, 3, 2, 2, 2, 511, 443, 3, 2, 2,
	2, 511, 454, 3, 2, 2, 2, 511, 462, 3, 2, 2, 2, 511, 473, 3, 2, 2, 2, 511,
	481, 3, 2, 2, 2, 511, 492, 3, 2, 2, 2, 511, 496, 3, 2, 2, 2, 511, 503,
	3, 2, 2, 2, 511, 506, 3, 2, 2, 2, 512, 559, 3, 2, 2, 2, 513, 514, 12, 19,
	2, 2, 514, 515, 9, 3, 2, 2, 515, 516, 5, 68, 35, 20, 516, 517, 8, 35, 1,
	2, 517, 558, 3, 2, 2, 2, 518, 519, 12, 15, 2, 2, 519, 520, 9, 4, 2, 2,
	520, 521, 5, 68, 35, 16, 521, 522, 8, 35, 1, 2, 522, 558, 3, 2, 2, 2, 523,
	524, 12, 11, 2, 2, 524, 525, 9, 5, 2, 2, 525, 526, 5, 68, 35, 12, 526,
	527, 8, 35, 1, 2, 527, 558, 3, 2, 2, 2, 528, 529, 12, 7, 2, 2, 529, 530,
	9, 7, 2, 2, 530, 531, 5, 68, 35, 8, 531, 532, 8, 35, 1, 2, 532, 558, 3,
	2, 2, 2, 533, 534, 12, 17, 2, 2, 534, 535, 9, 3, 2, 2, 535, 536, 7, 45,
	2, 2, 536, 537, 5, 68, 35, 2, 537, 538, 9, 2, 2, 2, 538, 539, 7, 46, 2,
	2, 539, 540, 8, 35, 1, 2, 540, 558, 3, 2, 2, 2, 541, 542, 12, 13, 2, 2,
	542, 543, 9, 4, 2, 2, 543, 544, 7, 45, 2, 2, 544, 545, 5, 68, 35, 2, 545,
	546, 9, 2, 2, 2, 546, 547, 7, 46, 2, 2, 547, 548, 8, 35, 1, 2, 548, 558,
	3, 2, 2, 2, 549, 550, 12, 9, 2, 2, 550, 551, 9, 5, 2, 2, 551, 552, 7, 45,
	2, 2, 552, 553, 5, 68, 35, 2, 553, 554, 9, 2, 2, 2, 554, 555, 7, 46, 2,
	2, 555, 556, 8, 35, 1, 2, 556, 558, 3, 2, 2, 2, 557, 513, 3, 2, 2, 2, 557,
	518, 3, 2, 2, 2, 557, 523, 3, 2, 2, 2, 557, 528, 3, 2, 2, 2, 557, 533,
	3, 2, 2, 2, 557, 541, 3, 2, 2, 2, 557, 549, 3, 2, 2, 2, 558, 561, 3, 2,
	2, 2, 559, 557, 3, 2, 2, 2, 559, 560, 3, 2, 2, 2, 560, 69, 3, 2, 2, 2,
	561, 559, 3, 2, 2, 2, 562, 563, 7, 22, 2, 2, 563, 582, 8, 36, 1, 2, 564,
	565, 7, 23, 2, 2, 565, 582, 8, 36, 1, 2, 566, 567, 7, 25, 2, 2, 567, 582,
	8, 36, 1, 2, 568, 569, 7, 26, 2, 2, 569, 582, 8, 36, 1, 2, 570, 571, 7,
	27, 2, 2, 571, 582, 8, 36, 1, 2, 572, 573, 5, 20, 11, 2, 573, 574, 8, 36,
	1, 2, 574, 582, 3, 2, 2, 2, 575, 576, 5, 40, 21, 2, 576, 577, 8, 36, 1,
	2, 577, 582, 3, 2, 2, 2, 578, 579, 5, 56, 29, 2, 579, 580, 8, 36, 1, 2,
	580, 582, 3, 2, 2, 2, 581, 562, 3, 2, 2, 2, 581, 564, 3, 2, 2, 2, 581,
	566, 3, 2, 2, 2, 581, 568, 3, 2, 2, 2, 581, 570, 3, 2, 2, 2, 581, 572,
	3, 2, 2, 2, 581, 575, 3, 2, 2, 2, 581, 578, 3, 2, 2, 2, 582, 71, 3, 2,
	2, 2, 28, 78, 85, 115, 133, 169, 204, 209, 241, 246, 266, 281, 289, 296,
	307, 320, 325, 344, 355, 362, 373, 408, 429, 511, 557, 559, 581,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'println!'", "'number'", "'as f64'", "'as i64'", "'let'", "'mut'",
	"'if'", "'else'", "'match'", "'while'", "'break'", "'continue'", "'return'",
	"'loop'", "'i64'", "'f64'", "'String'", "'bool'", "'&str'", "", "", "",
	"", "", "", "'.'", "';'", "','", "':'", "'='", "'=='", "'>='", "'=>'",
	"'<='", "'!='", "'>'", "'<'", "'*'", "'/'", "'%'", "'+'", "'-'", "'('",
	"')'", "'{'", "'}'", "'['", "']'", "'&&'", "'||'", "'|'", "'!'", "'_ =>'",
}
var symbolicNames = []string{
	"", "R_PRINTLN", "P_NUMBER", "R_AS_DOUBLE", "R_AS_INTEGER", "R_LET", "R_MUT",
	"R_IF", "R_ELSE", "R_MATCH", "R_WHILE", "R_BREAK", "R_CONTINUE", "R_RETURN",
	"R_LOOP", "R_INT", "R_FLOAT", "R_STRING", "R_BOOL", "R_STR", "NUMBER",
	"DOUBLE", "CHAR", "STRING", "BOOLEAN", "ID", "TK_PUNTO", "TK_PUNTOCOMA",
	"TK_COMA", "TK_DOSPUNTOS", "TK_IGUAL", "TK_IGUALIGUAL", "TK_MAYORIGUAL",
	"TK_IGUALMAYOR", "TK_MENORIGUAL", "TK_DIFIGUAL", "TK_MAYOR", "TK_MENOR",
	"TK_MULT", "TK_DIV", "TK_MODULO", "TK_MAS", "TK_MENOS", "TK_PARA", "TK_PARC",
	"TK_LLAVEA", "TK_LLAVEC", "TK_CORA", "TK_CORC", "TK_AND", "TK_OR", "TK_BARRA",
	"TK_NOT", "TK_GUIONBAJO", "WHITESPACE", "TK_MULTI", "TK_LINE",
}

var ruleNames = []string{
	"start", "instrucciones", "end_instr", "instruccion", "instr_println",
	"instr_declaracion", "instr_asignacion", "instr_if", "instr_else_if", "instr_ternario",
	"instr_else_if_ternario", "instr_match", "instr_case", "block_instr_match",
	"list_case", "list_expre_case", "block_case", "instr_default", "block_default",
	"instr_match_ter", "instr_case_ter", "list_case_ternario", "list_expre_case_ter",
	"block_case_ter", "instr_default_ter", "instr_while", "instr_loop", "instr_loop_ternario",
	"instr_break", "instr_continue", "instr_return", "instr_tipo", "expression",
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
	ChemsR_INT         = 15
	ChemsR_FLOAT       = 16
	ChemsR_STRING      = 17
	ChemsR_BOOL        = 18
	ChemsR_STR         = 19
	ChemsNUMBER        = 20
	ChemsDOUBLE        = 21
	ChemsCHAR          = 22
	ChemsSTRING        = 23
	ChemsBOOLEAN       = 24
	ChemsID            = 25
	ChemsTK_PUNTO      = 26
	ChemsTK_PUNTOCOMA  = 27
	ChemsTK_COMA       = 28
	ChemsTK_DOSPUNTOS  = 29
	ChemsTK_IGUAL      = 30
	ChemsTK_IGUALIGUAL = 31
	ChemsTK_MAYORIGUAL = 32
	ChemsTK_IGUALMAYOR = 33
	ChemsTK_MENORIGUAL = 34
	ChemsTK_DIFIGUAL   = 35
	ChemsTK_MAYOR      = 36
	ChemsTK_MENOR      = 37
	ChemsTK_MULT       = 38
	ChemsTK_DIV        = 39
	ChemsTK_MODULO     = 40
	ChemsTK_MAS        = 41
	ChemsTK_MENOS      = 42
	ChemsTK_PARA       = 43
	ChemsTK_PARC       = 44
	ChemsTK_LLAVEA     = 45
	ChemsTK_LLAVEC     = 46
	ChemsTK_CORA       = 47
	ChemsTK_CORC       = 48
	ChemsTK_AND        = 49
	ChemsTK_OR         = 50
	ChemsTK_BARRA      = 51
	ChemsTK_NOT        = 52
	ChemsTK_GUIONBAJO  = 53
	ChemsWHITESPACE    = 54
	ChemsTK_MULTI      = 55
	ChemsTK_LINE       = 56
)

// Chems rules.
const (
	ChemsRULE_start                  = 0
	ChemsRULE_instrucciones          = 1
	ChemsRULE_end_instr              = 2
	ChemsRULE_instruccion            = 3
	ChemsRULE_instr_println          = 4
	ChemsRULE_instr_declaracion      = 5
	ChemsRULE_instr_asignacion       = 6
	ChemsRULE_instr_if               = 7
	ChemsRULE_instr_else_if          = 8
	ChemsRULE_instr_ternario         = 9
	ChemsRULE_instr_else_if_ternario = 10
	ChemsRULE_instr_match            = 11
	ChemsRULE_instr_case             = 12
	ChemsRULE_block_instr_match      = 13
	ChemsRULE_list_case              = 14
	ChemsRULE_list_expre_case        = 15
	ChemsRULE_block_case             = 16
	ChemsRULE_instr_default          = 17
	ChemsRULE_block_default          = 18
	ChemsRULE_instr_match_ter        = 19
	ChemsRULE_instr_case_ter         = 20
	ChemsRULE_list_case_ternario     = 21
	ChemsRULE_list_expre_case_ter    = 22
	ChemsRULE_block_case_ter         = 23
	ChemsRULE_instr_default_ter      = 24
	ChemsRULE_instr_while            = 25
	ChemsRULE_instr_loop             = 26
	ChemsRULE_instr_loop_ternario    = 27
	ChemsRULE_instr_break            = 28
	ChemsRULE_instr_continue         = 29
	ChemsRULE_instr_return           = 30
	ChemsRULE_instr_tipo             = 31
	ChemsRULE_expression             = 32
	ChemsRULE_exp_arit               = 33
	ChemsRULE_primitivo              = 34
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
		p.SetState(70)

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
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_PRINTLN)|(1<<ChemsR_LET)|(1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_WHILE)|(1<<ChemsR_BREAK)|(1<<ChemsR_CONTINUE)|(1<<ChemsR_LOOP)|(1<<ChemsID))) != 0) {
		{
			p.SetState(73)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).list = append(localctx.(*InstruccionesContext).list, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(76)
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

	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_PUNTOCOMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*End_instrContext).v = 1

	case ChemsEOF, ChemsR_PRINTLN, ChemsR_LET, ChemsR_IF, ChemsR_MATCH, ChemsR_WHILE, ChemsR_BREAK, ChemsR_CONTINUE, ChemsR_LOOP, ChemsID, ChemsTK_COMA, ChemsTK_LLAVEC:
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

	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_PRINTLN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)

			var _x = p.Instr_println()

			localctx.(*InstruccionContext)._instr_println = _x
		}
		{
			p.SetState(86)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_println().GetInstr()

	case ChemsR_LET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)

			var _x = p.Instr_declaracion()

			localctx.(*InstruccionContext)._instr_declaracion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_declaracion().GetInstr()

	case ChemsID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)

			var _x = p.Instr_asignacion()

			localctx.(*InstruccionContext)._instr_asignacion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_asignacion().GetInstr()

	case ChemsR_IF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(95)

			var _x = p.Instr_if()

			localctx.(*InstruccionContext)._instr_if = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_if().GetInstr()

	case ChemsR_MATCH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(98)

			var _x = p.Instr_match()

			localctx.(*InstruccionContext)._instr_match = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_match().GetInstr()

	case ChemsR_WHILE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(101)

			var _x = p.Instr_while()

			localctx.(*InstruccionContext)._instr_while = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_while().GetInstr()

	case ChemsR_BREAK:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(104)

			var _x = p.Instr_break()

			localctx.(*InstruccionContext)._instr_break = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_break().GetInstr()

	case ChemsR_CONTINUE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(107)

			var _x = p.Instr_continue()

			localctx.(*InstruccionContext)._instr_continue = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_continue().GetInstr()

	case ChemsR_LOOP:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(110)

			var _x = p.Instr_loop()

			localctx.(*InstruccionContext)._instr_loop = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_loop().GetInstr()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

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
	p.EnterRule(localctx, 8, ChemsRULE_instr_println)

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

	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Match(ChemsR_PRINTLN)
		}
		{
			p.SetState(116)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(117)

			var _x = p.Expression()

			localctx.(*Instr_printlnContext)._expression = _x
		}
		{
			p.SetState(118)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(119)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_printlnContext).instr = instruction.PRINTLN(localctx.(*Instr_printlnContext).Get_expression().GetP(), "-1")

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Match(ChemsR_PRINTLN)
		}
		{
			p.SetState(123)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(124)

			var _m = p.Match(ChemsSTRING)

			localctx.(*Instr_printlnContext)._STRING = _m
		}
		{
			p.SetState(125)
			p.Match(ChemsTK_COMA)
		}
		{
			p.SetState(126)

			var _x = p.Expression()

			localctx.(*Instr_printlnContext)._expression = _x
		}
		{
			p.SetState(127)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(128)
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
	p.EnterRule(localctx, 10, ChemsRULE_instr_declaracion)

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

	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(134)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(135)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(136)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(137)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(138)
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
			p.SetState(141)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(142)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(143)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(144)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(145)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(146)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(147)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(148)
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
			p.SetState(151)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(152)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(153)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(154)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(155)
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
			p.SetState(158)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(159)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(160)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(161)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(162)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(163)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(164)
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
	p.EnterRule(localctx, 12, ChemsRULE_instr_asignacion)

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
		p.SetState(169)

		var _m = p.Match(ChemsID)

		localctx.(*Instr_asignacionContext)._ID = _m
	}
	{
		p.SetState(170)
		p.Match(ChemsTK_IGUAL)
	}
	{
		p.SetState(171)

		var _x = p.Expression()

		localctx.(*Instr_asignacionContext)._expression = _x
	}
	{
		p.SetState(172)
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
	p.EnterRule(localctx, 14, ChemsRULE_instr_if)

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

	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(176)

			var _x = p.Expression()

			localctx.(*Instr_ifContext)._expression = _x
		}
		{
			p.SetState(177)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(178)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(179)
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
			p.SetState(182)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(183)

			var _x = p.Expression()

			localctx.(*Instr_ifContext)._expression = _x
		}
		{
			p.SetState(184)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(185)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(186)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(187)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(188)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(189)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).right_intr = _x
		}
		{
			p.SetState(190)
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
			p.SetState(193)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(194)

			var _x = p.Expression()

			localctx.(*Instr_ifContext)._expression = _x
		}
		{
			p.SetState(195)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(196)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(197)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(198)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(199)

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
	p.EnterRule(localctx, 16, ChemsRULE_instr_else_if)

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
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(204)

				var _x = p.Instr_if()

				localctx.(*Instr_else_ifContext)._instr_if = _x
			}
			localctx.(*Instr_else_ifContext).e = append(localctx.(*Instr_else_ifContext).e, localctx.(*Instr_else_ifContext)._instr_if)

		}
		p.SetState(209)
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
	p.EnterRule(localctx, 18, ChemsRULE_instr_ternario)

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

	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(213)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).block = _x
		}
		{
			p.SetState(214)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(215)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(216)
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
			p.SetState(219)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(220)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).block = _x
		}
		{
			p.SetState(221)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(222)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(223)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(224)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(225)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(226)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).right_intr = _x
		}
		{
			p.SetState(227)
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
			p.SetState(230)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(231)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).block = _x
		}
		{
			p.SetState(232)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(233)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(234)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(235)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(236)

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
	p.EnterRule(localctx, 20, ChemsRULE_instr_else_if_ternario)

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
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(241)

				var _x = p.Instr_ternario()

				localctx.(*Instr_else_if_ternarioContext)._instr_ternario = _x
			}
			localctx.(*Instr_else_if_ternarioContext).e = append(localctx.(*Instr_else_if_ternarioContext).e, localctx.(*Instr_else_if_ternarioContext)._instr_ternario)

		}
		p.SetState(246)
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
	p.EnterRule(localctx, 22, ChemsRULE_instr_match)

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

	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_matchContext)._R_MATCH = _m
		}
		{
			p.SetState(250)

			var _x = p.Expression()

			localctx.(*Instr_matchContext)._expression = _x
		}
		{
			p.SetState(251)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(252)

			var _x = p.List_case()

			localctx.(*Instr_matchContext)._list_case = _x
		}
		{
			p.SetState(253)

			var _x = p.Block_default()

			localctx.(*Instr_matchContext)._block_default = _x
		}
		{
			p.SetState(254)
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
			p.SetState(257)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_matchContext)._R_MATCH = _m
		}
		{
			p.SetState(258)

			var _x = p.Expression()

			localctx.(*Instr_matchContext)._expression = _x
		}
		{
			p.SetState(259)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(260)

			var _x = p.Block_default()

			localctx.(*Instr_matchContext)._block_default = _x
		}
		{
			p.SetState(261)
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
	p.EnterRule(localctx, 24, ChemsRULE_instr_case)

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

	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)

			var _x = p.List_expre_case()

			localctx.(*Instr_caseContext)._list_expre_case = _x
		}
		{
			p.SetState(267)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(268)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(269)

			var _x = p.Instrucciones()

			localctx.(*Instr_caseContext)._instrucciones = _x
		}
		{
			p.SetState(270)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_caseContext).instr = control.NewCase(nil, localctx.(*Instr_caseContext).Get_list_expre_case().GetL(), localctx.(*Instr_caseContext).Get_instrucciones().GetL())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(273)

			var _x = p.List_expre_case()

			localctx.(*Instr_caseContext)._list_expre_case = _x
		}
		{
			p.SetState(274)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(275)

			var _x = p.Block_instr_match()

			localctx.(*Instr_caseContext)._block_instr_match = _x
		}
		{
			p.SetState(276)
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
	p.EnterRule(localctx, 26, ChemsRULE_block_instr_match)

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
		p.SetState(281)

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
	p.EnterRule(localctx, 28, ChemsRULE_list_case)

	localctx.(*List_caseContext).l = arrayList.New()

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
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(ChemsTK_MENOS-42))|(1<<(ChemsTK_PARA-42))|(1<<(ChemsTK_NOT-42)))) != 0) {
		{
			p.SetState(284)

			var _x = p.Instr_case()

			localctx.(*List_caseContext)._instr_case = _x
		}
		localctx.(*List_caseContext).e = append(localctx.(*List_caseContext).e, localctx.(*List_caseContext)._instr_case)

		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 30, ChemsRULE_list_expre_case)

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
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(ChemsTK_MENOS-42))|(1<<(ChemsTK_PARA-42))|(1<<(ChemsTK_NOT-42)))) != 0) {
		{
			p.SetState(291)

			var _x = p.Block_case()

			localctx.(*List_expre_caseContext)._block_case = _x
		}
		localctx.(*List_expre_caseContext).e = append(localctx.(*List_expre_caseContext).e, localctx.(*List_expre_caseContext)._block_case)

		p.SetState(294)
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
	p.EnterRule(localctx, 32, ChemsRULE_block_case)

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

	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)

			var _x = p.Expression()

			localctx.(*Block_caseContext)._expression = _x
		}
		{
			p.SetState(299)
			p.Match(ChemsTK_BARRA)
		}
		localctx.(*Block_caseContext).instr = control.NewCase(localctx.(*Block_caseContext).Get_expression().GetP(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(302)

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

func (s *Instr_defaultContext) TK_GUIONBAJO() antlr.TerminalNode {
	return s.GetToken(ChemsTK_GUIONBAJO, 0)
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
	p.EnterRule(localctx, 34, ChemsRULE_instr_default)

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

	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(307)
			p.Match(ChemsTK_GUIONBAJO)
		}
		{
			p.SetState(308)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(309)

			var _x = p.Instrucciones()

			localctx.(*Instr_defaultContext)._instrucciones = _x
		}
		{
			p.SetState(310)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_defaultContext).instr = control.NewDefault(localctx.(*Instr_defaultContext).Get_instrucciones().GetL())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(313)
			p.Match(ChemsTK_GUIONBAJO)
		}
		{
			p.SetState(314)

			var _x = p.Block_instr_match()

			localctx.(*Instr_defaultContext)._block_instr_match = _x
		}
		{
			p.SetState(315)
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
	p.EnterRule(localctx, 36, ChemsRULE_block_default)

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
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ChemsTK_GUIONBAJO {
		{
			p.SetState(320)

			var _x = p.Instr_default()

			localctx.(*Block_defaultContext)._instr_default = _x
		}
		localctx.(*Block_defaultContext).e = append(localctx.(*Block_defaultContext).e, localctx.(*Block_defaultContext)._instr_default)

		p.SetState(323)
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
	p.EnterRule(localctx, 38, ChemsRULE_instr_match_ter)

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

	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(327)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_match_terContext)._R_MATCH = _m
		}
		{
			p.SetState(328)

			var _x = p.Expression()

			localctx.(*Instr_match_terContext)._expression = _x
		}
		{
			p.SetState(329)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(330)

			var _x = p.List_case_ternario()

			localctx.(*Instr_match_terContext)._list_case_ternario = _x
		}
		{
			p.SetState(331)

			var _x = p.Instr_default_ter()

			localctx.(*Instr_match_terContext)._instr_default_ter = _x
		}
		{
			p.SetState(332)
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
			p.SetState(335)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_match_terContext)._R_MATCH = _m
		}
		{
			p.SetState(336)

			var _x = p.Expression()

			localctx.(*Instr_match_terContext)._expression = _x
		}
		{
			p.SetState(337)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(338)

			var _x = p.Instr_default_ter()

			localctx.(*Instr_match_terContext)._instr_default_ter = _x
		}
		{
			p.SetState(339)
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
	p.EnterRule(localctx, 40, ChemsRULE_instr_case_ter)

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
		p.SetState(344)

		var _x = p.List_expre_case_ter()

		localctx.(*Instr_case_terContext)._list_expre_case_ter = _x
	}
	{
		p.SetState(345)
		p.Match(ChemsTK_IGUALMAYOR)
	}
	{
		p.SetState(346)

		var _x = p.Expression()

		localctx.(*Instr_case_terContext)._expression = _x
	}
	{
		p.SetState(347)
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
	p.EnterRule(localctx, 42, ChemsRULE_list_case_ternario)

	localctx.(*List_case_ternarioContext).l = arrayList.New()

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
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(ChemsTK_MENOS-42))|(1<<(ChemsTK_PARA-42))|(1<<(ChemsTK_NOT-42)))) != 0) {
		{
			p.SetState(350)

			var _x = p.Instr_case_ter()

			localctx.(*List_case_ternarioContext)._instr_case_ter = _x
		}
		localctx.(*List_case_ternarioContext).e = append(localctx.(*List_case_ternarioContext).e, localctx.(*List_case_ternarioContext)._instr_case_ter)

		p.SetState(353)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 44, ChemsRULE_list_expre_case_ter)

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
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(ChemsTK_MENOS-42))|(1<<(ChemsTK_PARA-42))|(1<<(ChemsTK_NOT-42)))) != 0) {
		{
			p.SetState(357)

			var _x = p.Block_case_ter()

			localctx.(*List_expre_case_terContext)._block_case_ter = _x
		}
		localctx.(*List_expre_case_terContext).e = append(localctx.(*List_expre_case_terContext).e, localctx.(*List_expre_case_terContext)._block_case_ter)

		p.SetState(360)
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
	p.EnterRule(localctx, 46, ChemsRULE_block_case_ter)

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

	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(364)

			var _x = p.Expression()

			localctx.(*Block_case_terContext)._expression = _x
		}
		{
			p.SetState(365)
			p.Match(ChemsTK_BARRA)
		}
		localctx.(*Block_case_terContext).instr = control.NewCaseTer(localctx.(*Block_case_terContext).Get_expression().GetP(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(368)

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

func (s *Instr_default_terContext) TK_GUIONBAJO() antlr.TerminalNode {
	return s.GetToken(ChemsTK_GUIONBAJO, 0)
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
	p.EnterRule(localctx, 48, ChemsRULE_instr_default_ter)

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
		p.SetState(373)
		p.Match(ChemsTK_GUIONBAJO)
	}
	{
		p.SetState(374)

		var _x = p.Expression()

		localctx.(*Instr_default_terContext)._expression = _x
	}
	{
		p.SetState(375)
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
	p.EnterRule(localctx, 50, ChemsRULE_instr_while)

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
		p.SetState(378)
		p.Match(ChemsR_WHILE)
	}
	{
		p.SetState(379)

		var _x = p.Expression()

		localctx.(*Instr_whileContext)._expression = _x
	}
	{
		p.SetState(380)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(381)

		var _x = p.Instrucciones()

		localctx.(*Instr_whileContext)._instrucciones = _x
	}
	{
		p.SetState(382)
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
	p.EnterRule(localctx, 52, ChemsRULE_instr_loop)

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
		p.SetState(385)
		p.Match(ChemsR_LOOP)
	}
	{
		p.SetState(386)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(387)

		var _x = p.Instrucciones()

		localctx.(*Instr_loopContext)._instrucciones = _x
	}
	{
		p.SetState(388)
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
	p.EnterRule(localctx, 54, ChemsRULE_instr_loop_ternario)

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
		p.SetState(391)
		p.Match(ChemsR_LOOP)
	}
	{
		p.SetState(392)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(393)

		var _x = p.Instrucciones()

		localctx.(*Instr_loop_ternarioContext)._instrucciones = _x
	}
	{
		p.SetState(394)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_loop_ternarioContext).instr = loops.NewLoopTernario(localctx.(*Instr_loop_ternarioContext).Get_instrucciones().GetL())

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
	p.EnterRule(localctx, 56, ChemsRULE_instr_break)

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

	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(397)

			var _m = p.Match(ChemsR_BREAK)

			localctx.(*Instr_breakContext)._R_BREAK = _m
		}
		{
			p.SetState(398)
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
			p.SetState(401)

			var _m = p.Match(ChemsR_BREAK)

			localctx.(*Instr_breakContext)._R_BREAK = _m
		}
		{
			p.SetState(402)

			var _x = p.Expression()

			localctx.(*Instr_breakContext)._expression = _x
		}
		{
			p.SetState(403)
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
	p.EnterRule(localctx, 58, ChemsRULE_instr_continue)

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
		p.SetState(408)

		var _m = p.Match(ChemsR_CONTINUE)

		localctx.(*Instr_continueContext)._R_CONTINUE = _m
	}
	{
		p.SetState(409)
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
	p.EnterRule(localctx, 60, ChemsRULE_instr_return)

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
		p.SetState(412)

		var _m = p.Match(ChemsR_RETURN)

		localctx.(*Instr_returnContext)._R_RETURN = _m
	}
	{
		p.SetState(413)

		var _x = p.Expression()

		localctx.(*Instr_returnContext)._expression = _x
	}
	{
		p.SetState(414)
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
	p.EnterRule(localctx, 62, ChemsRULE_instr_tipo)

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

	p.SetState(427)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(417)
			p.Match(ChemsR_INT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.INTEGER

	case ChemsR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(419)
			p.Match(ChemsR_FLOAT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.FLOAT

	case ChemsR_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(421)
			p.Match(ChemsR_STRING)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STRING

	case ChemsR_STR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(423)
			p.Match(ChemsR_STR)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STRING

	case ChemsR_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(425)
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
	p.EnterRule(localctx, 64, ChemsRULE_expression)

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
		p.SetState(429)

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
	_startState := 66
	p.EnterRecursionRule(localctx, 66, ChemsRULE_exp_arit, _p)
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
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(433)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(434)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(435)

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
			p.SetState(436)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(437)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(ChemsTK_MULT-38))|(1<<(ChemsTK_DIV-38))|(1<<(ChemsTK_MODULO-38)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(438)

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
			p.SetState(441)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(442)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(443)

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
			p.SetState(444)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(445)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(ChemsTK_MULT-38))|(1<<(ChemsTK_DIV-38))|(1<<(ChemsTK_MODULO-38)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(446)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(447)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(448)

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
			p.SetState(449)
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
			p.SetState(452)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(453)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(454)

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
			p.SetState(455)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(456)

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
			p.SetState(457)

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
			p.SetState(460)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(461)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(462)

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
			p.SetState(463)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(464)

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
			p.SetState(465)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(466)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(467)

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
			p.SetState(468)
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
			p.SetState(471)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(472)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(473)

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
			p.SetState(474)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(475)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(ChemsTK_IGUALIGUAL-31))|(1<<(ChemsTK_MAYORIGUAL-31))|(1<<(ChemsTK_MENORIGUAL-31))|(1<<(ChemsTK_DIFIGUAL-31))|(1<<(ChemsTK_MAYOR-31))|(1<<(ChemsTK_MENOR-31)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(476)

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
			p.SetState(479)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(480)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(481)

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
			p.SetState(482)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(483)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(ChemsTK_IGUALIGUAL-31))|(1<<(ChemsTK_MAYORIGUAL-31))|(1<<(ChemsTK_MENORIGUAL-31))|(1<<(ChemsTK_DIFIGUAL-31))|(1<<(ChemsTK_MAYOR-31))|(1<<(ChemsTK_MENOR-31)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(484)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(485)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(486)

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
			p.SetState(487)
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
			p.SetState(490)

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
			p.SetState(491)

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
			p.SetState(494)

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
			p.SetState(495)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(496)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(497)

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
			p.SetState(498)
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
			p.SetState(501)

			var _x = p.Primitivo()

			localctx.(*Exp_aritContext)._primitivo = _x
		}
		localctx.(*Exp_aritContext).p = localctx.(*Exp_aritContext).Get_primitivo().GetP()

	case 10:
		{
			p.SetState(504)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(505)

			var _x = p.Expression()

			localctx.(*Exp_aritContext)._expression = _x
		}
		{
			p.SetState(506)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Exp_aritContext).p = localctx.(*Exp_aritContext).Get_expression().GetP()

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(555)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(511)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(512)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(ChemsTK_MULT-38))|(1<<(ChemsTK_DIV-38))|(1<<(ChemsTK_MODULO-38)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(513)

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
				p.SetState(516)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(517)

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
					p.SetState(518)

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
				p.SetState(521)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(522)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(ChemsTK_IGUALIGUAL-31))|(1<<(ChemsTK_MAYORIGUAL-31))|(1<<(ChemsTK_MENORIGUAL-31))|(1<<(ChemsTK_DIFIGUAL-31))|(1<<(ChemsTK_MAYOR-31))|(1<<(ChemsTK_MENOR-31)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(523)

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
				p.SetState(526)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(527)

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
					p.SetState(528)

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
				p.SetState(531)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(532)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(ChemsTK_MULT-38))|(1<<(ChemsTK_DIV-38))|(1<<(ChemsTK_MODULO-38)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(533)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(534)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(535)

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
					p.SetState(536)
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
				p.SetState(539)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(540)

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
					p.SetState(541)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(542)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(543)

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
					p.SetState(544)
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
				p.SetState(547)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(548)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(ChemsTK_IGUALIGUAL-31))|(1<<(ChemsTK_MAYORIGUAL-31))|(1<<(ChemsTK_MENORIGUAL-31))|(1<<(ChemsTK_DIFIGUAL-31))|(1<<(ChemsTK_MAYOR-31))|(1<<(ChemsTK_MENOR-31)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(549)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(550)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(551)

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
					p.SetState(552)
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
		p.SetState(559)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 68, ChemsRULE_primitivo)

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

	p.SetState(579)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(560)

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
			p.SetState(562)

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
			p.SetState(564)

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
			p.SetState(566)

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
			p.SetState(568)

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
			p.SetState(570)

			var _x = p.Instr_ternario()

			localctx.(*PrimitivoContext)._instr_ternario = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_ternario().GetP()

	case ChemsR_MATCH:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(573)

			var _x = p.Instr_match_ter()

			localctx.(*PrimitivoContext)._instr_match_ter = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_match_ter().GetInstr()

	case ChemsR_LOOP:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(576)

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
	case 33:
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

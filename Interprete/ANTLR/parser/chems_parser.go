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
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 53, 417,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2, 3, 3, 6,
	3, 47, 10, 3, 13, 3, 14, 3, 48, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 81, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 5, 5, 117, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 5, 7, 152, 10, 7, 3, 8, 7, 8, 155, 10, 8, 12, 8, 14, 8, 158, 11, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 189, 10, 9, 3, 10, 7, 10, 192, 10,
	10, 12, 10, 14, 10, 195, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 5, 11, 214, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 13, 6, 13, 224, 10, 13, 13, 13, 14, 13, 225, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 237, 10, 14, 3, 15, 6,
	15, 240, 10, 15, 13, 15, 14, 15, 241, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 17, 6, 17, 253, 10, 17, 13, 17, 14, 17, 254, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 5, 18, 269, 10, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	5, 20, 351, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 397, 10, 20, 12, 20, 14, 20,
	400, 11, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 415, 10, 21, 3, 21, 2, 3, 38, 22,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 2, 8, 3, 2, 5, 6, 3, 2, 35, 37, 3, 2, 38, 39, 4, 2, 28, 29, 31, 34,
	4, 2, 39, 39, 49, 49, 3, 2, 46, 47, 2, 441, 2, 42, 3, 2, 2, 2, 4, 46, 3,
	2, 2, 2, 6, 80, 3, 2, 2, 2, 8, 116, 3, 2, 2, 2, 10, 118, 3, 2, 2, 2, 12,
	151, 3, 2, 2, 2, 14, 156, 3, 2, 2, 2, 16, 188, 3, 2, 2, 2, 18, 193, 3,
	2, 2, 2, 20, 213, 3, 2, 2, 2, 22, 215, 3, 2, 2, 2, 24, 223, 3, 2, 2, 2,
	26, 236, 3, 2, 2, 2, 28, 239, 3, 2, 2, 2, 30, 245, 3, 2, 2, 2, 32, 252,
	3, 2, 2, 2, 34, 268, 3, 2, 2, 2, 36, 270, 3, 2, 2, 2, 38, 350, 3, 2, 2,
	2, 40, 414, 3, 2, 2, 2, 42, 43, 5, 4, 3, 2, 43, 44, 8, 2, 1, 2, 44, 3,
	3, 2, 2, 2, 45, 47, 5, 6, 4, 2, 46, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2,
	48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 51, 8,
	3, 1, 2, 51, 5, 3, 2, 2, 2, 52, 53, 7, 3, 2, 2, 53, 54, 7, 40, 2, 2, 54,
	55, 5, 36, 19, 2, 55, 56, 7, 41, 2, 2, 56, 57, 7, 24, 2, 2, 57, 58, 8,
	4, 1, 2, 58, 81, 3, 2, 2, 2, 59, 60, 7, 3, 2, 2, 60, 61, 7, 40, 2, 2, 61,
	62, 7, 20, 2, 2, 62, 63, 7, 25, 2, 2, 63, 64, 5, 36, 19, 2, 64, 65, 7,
	41, 2, 2, 65, 66, 7, 24, 2, 2, 66, 67, 8, 4, 1, 2, 67, 81, 3, 2, 2, 2,
	68, 69, 5, 8, 5, 2, 69, 70, 8, 4, 1, 2, 70, 81, 3, 2, 2, 2, 71, 72, 5,
	10, 6, 2, 72, 73, 8, 4, 1, 2, 73, 81, 3, 2, 2, 2, 74, 75, 5, 12, 7, 2,
	75, 76, 8, 4, 1, 2, 76, 81, 3, 2, 2, 2, 77, 78, 5, 20, 11, 2, 78, 79, 8,
	4, 1, 2, 79, 81, 3, 2, 2, 2, 80, 52, 3, 2, 2, 2, 80, 59, 3, 2, 2, 2, 80,
	68, 3, 2, 2, 2, 80, 71, 3, 2, 2, 2, 80, 74, 3, 2, 2, 2, 80, 77, 3, 2, 2,
	2, 81, 7, 3, 2, 2, 2, 82, 83, 7, 7, 2, 2, 83, 84, 7, 8, 2, 2, 84, 85, 7,
	22, 2, 2, 85, 86, 7, 27, 2, 2, 86, 87, 5, 36, 19, 2, 87, 88, 7, 24, 2,
	2, 88, 89, 8, 5, 1, 2, 89, 117, 3, 2, 2, 2, 90, 91, 7, 7, 2, 2, 91, 92,
	7, 8, 2, 2, 92, 93, 7, 22, 2, 2, 93, 94, 7, 26, 2, 2, 94, 95, 5, 34, 18,
	2, 95, 96, 7, 27, 2, 2, 96, 97, 5, 36, 19, 2, 97, 98, 7, 24, 2, 2, 98,
	99, 8, 5, 1, 2, 99, 117, 3, 2, 2, 2, 100, 101, 7, 7, 2, 2, 101, 102, 7,
	22, 2, 2, 102, 103, 7, 27, 2, 2, 103, 104, 5, 36, 19, 2, 104, 105, 7, 24,
	2, 2, 105, 106, 8, 5, 1, 2, 106, 117, 3, 2, 2, 2, 107, 108, 7, 7, 2, 2,
	108, 109, 7, 22, 2, 2, 109, 110, 7, 26, 2, 2, 110, 111, 5, 34, 18, 2, 111,
	112, 7, 27, 2, 2, 112, 113, 5, 36, 19, 2, 113, 114, 7, 24, 2, 2, 114, 115,
	8, 5, 1, 2, 115, 117, 3, 2, 2, 2, 116, 82, 3, 2, 2, 2, 116, 90, 3, 2, 2,
	2, 116, 100, 3, 2, 2, 2, 116, 107, 3, 2, 2, 2, 117, 9, 3, 2, 2, 2, 118,
	119, 7, 22, 2, 2, 119, 120, 7, 27, 2, 2, 120, 121, 5, 36, 19, 2, 121, 122,
	7, 24, 2, 2, 122, 123, 8, 6, 1, 2, 123, 11, 3, 2, 2, 2, 124, 125, 7, 9,
	2, 2, 125, 126, 5, 36, 19, 2, 126, 127, 7, 42, 2, 2, 127, 128, 5, 4, 3,
	2, 128, 129, 7, 43, 2, 2, 129, 130, 8, 7, 1, 2, 130, 152, 3, 2, 2, 2, 131,
	132, 7, 9, 2, 2, 132, 133, 5, 36, 19, 2, 133, 134, 7, 42, 2, 2, 134, 135,
	5, 4, 3, 2, 135, 136, 7, 43, 2, 2, 136, 137, 7, 10, 2, 2, 137, 138, 7,
	42, 2, 2, 138, 139, 5, 4, 3, 2, 139, 140, 7, 43, 2, 2, 140, 141, 8, 7,
	1, 2, 141, 152, 3, 2, 2, 2, 142, 143, 7, 9, 2, 2, 143, 144, 5, 36, 19,
	2, 144, 145, 7, 42, 2, 2, 145, 146, 5, 4, 3, 2, 146, 147, 7, 43, 2, 2,
	147, 148, 7, 10, 2, 2, 148, 149, 5, 14, 8, 2, 149, 150, 8, 7, 1, 2, 150,
	152, 3, 2, 2, 2, 151, 124, 3, 2, 2, 2, 151, 131, 3, 2, 2, 2, 151, 142,
	3, 2, 2, 2, 152, 13, 3, 2, 2, 2, 153, 155, 5, 12, 7, 2, 154, 153, 3, 2,
	2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2,
	157, 159, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 159, 160, 8, 8, 1, 2, 160,
	15, 3, 2, 2, 2, 161, 162, 7, 9, 2, 2, 162, 163, 5, 36, 19, 2, 163, 164,
	7, 42, 2, 2, 164, 165, 5, 36, 19, 2, 165, 166, 7, 43, 2, 2, 166, 167, 8,
	9, 1, 2, 167, 189, 3, 2, 2, 2, 168, 169, 7, 9, 2, 2, 169, 170, 5, 36, 19,
	2, 170, 171, 7, 42, 2, 2, 171, 172, 5, 36, 19, 2, 172, 173, 7, 43, 2, 2,
	173, 174, 7, 10, 2, 2, 174, 175, 7, 42, 2, 2, 175, 176, 5, 36, 19, 2, 176,
	177, 7, 43, 2, 2, 177, 178, 8, 9, 1, 2, 178, 189, 3, 2, 2, 2, 179, 180,
	7, 9, 2, 2, 180, 181, 5, 36, 19, 2, 181, 182, 7, 42, 2, 2, 182, 183, 5,
	36, 19, 2, 183, 184, 7, 43, 2, 2, 184, 185, 7, 10, 2, 2, 185, 186, 5, 18,
	10, 2, 186, 187, 8, 9, 1, 2, 187, 189, 3, 2, 2, 2, 188, 161, 3, 2, 2, 2,
	188, 168, 3, 2, 2, 2, 188, 179, 3, 2, 2, 2, 189, 17, 3, 2, 2, 2, 190, 192,
	5, 16, 9, 2, 191, 190, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193, 191, 3, 2,
	2, 2, 193, 194, 3, 2, 2, 2, 194, 196, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2,
	196, 197, 8, 10, 1, 2, 197, 19, 3, 2, 2, 2, 198, 199, 7, 11, 2, 2, 199,
	200, 5, 36, 19, 2, 200, 201, 7, 42, 2, 2, 201, 202, 5, 24, 13, 2, 202,
	203, 5, 32, 17, 2, 203, 204, 7, 43, 2, 2, 204, 205, 8, 11, 1, 2, 205, 214,
	3, 2, 2, 2, 206, 207, 7, 11, 2, 2, 207, 208, 5, 36, 19, 2, 208, 209, 7,
	42, 2, 2, 209, 210, 5, 32, 17, 2, 210, 211, 7, 43, 2, 2, 211, 212, 8, 11,
	1, 2, 212, 214, 3, 2, 2, 2, 213, 198, 3, 2, 2, 2, 213, 206, 3, 2, 2, 2,
	214, 21, 3, 2, 2, 2, 215, 216, 5, 28, 15, 2, 216, 217, 7, 30, 2, 2, 217,
	218, 7, 42, 2, 2, 218, 219, 5, 4, 3, 2, 219, 220, 7, 43, 2, 2, 220, 221,
	8, 12, 1, 2, 221, 23, 3, 2, 2, 2, 222, 224, 5, 22, 12, 2, 223, 222, 3,
	2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 225, 226, 3, 2, 2,
	2, 226, 227, 3, 2, 2, 2, 227, 228, 8, 13, 1, 2, 228, 25, 3, 2, 2, 2, 229,
	230, 5, 36, 19, 2, 230, 231, 7, 48, 2, 2, 231, 232, 8, 14, 1, 2, 232, 237,
	3, 2, 2, 2, 233, 234, 5, 36, 19, 2, 234, 235, 8, 14, 1, 2, 235, 237, 3,
	2, 2, 2, 236, 229, 3, 2, 2, 2, 236, 233, 3, 2, 2, 2, 237, 27, 3, 2, 2,
	2, 238, 240, 5, 26, 14, 2, 239, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2,
	241, 239, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243,
	244, 8, 15, 1, 2, 244, 29, 3, 2, 2, 2, 245, 246, 7, 50, 2, 2, 246, 247,
	7, 42, 2, 2, 247, 248, 5, 4, 3, 2, 248, 249, 7, 43, 2, 2, 249, 250, 8,
	16, 1, 2, 250, 31, 3, 2, 2, 2, 251, 253, 5, 30, 16, 2, 252, 251, 3, 2,
	2, 2, 253, 254, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2,
	255, 256, 3, 2, 2, 2, 256, 257, 8, 17, 1, 2, 257, 33, 3, 2, 2, 2, 258,
	259, 7, 12, 2, 2, 259, 269, 8, 18, 1, 2, 260, 261, 7, 13, 2, 2, 261, 269,
	8, 18, 1, 2, 262, 263, 7, 14, 2, 2, 263, 269, 8, 18, 1, 2, 264, 265, 7,
	16, 2, 2, 265, 269, 8, 18, 1, 2, 266, 267, 7, 15, 2, 2, 267, 269, 8, 18,
	1, 2, 268, 258, 3, 2, 2, 2, 268, 260, 3, 2, 2, 2, 268, 262, 3, 2, 2, 2,
	268, 264, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 269, 35, 3, 2, 2, 2, 270, 271,
	5, 38, 20, 2, 271, 272, 8, 19, 1, 2, 272, 37, 3, 2, 2, 2, 273, 274, 8,
	20, 1, 2, 274, 275, 7, 40, 2, 2, 275, 276, 5, 38, 20, 2, 276, 277, 9, 2,
	2, 2, 277, 278, 7, 41, 2, 2, 278, 279, 9, 3, 2, 2, 279, 280, 5, 38, 20,
	18, 280, 281, 8, 20, 1, 2, 281, 351, 3, 2, 2, 2, 282, 283, 7, 40, 2, 2,
	283, 284, 5, 38, 20, 2, 284, 285, 9, 2, 2, 2, 285, 286, 7, 41, 2, 2, 286,
	287, 9, 3, 2, 2, 287, 288, 7, 40, 2, 2, 288, 289, 5, 38, 20, 2, 289, 290,
	9, 2, 2, 2, 290, 291, 7, 41, 2, 2, 291, 292, 8, 20, 1, 2, 292, 351, 3,
	2, 2, 2, 293, 294, 7, 40, 2, 2, 294, 295, 5, 38, 20, 2, 295, 296, 9, 2,
	2, 2, 296, 297, 7, 41, 2, 2, 297, 298, 9, 4, 2, 2, 298, 299, 5, 38, 20,
	14, 299, 300, 8, 20, 1, 2, 300, 351, 3, 2, 2, 2, 301, 302, 7, 40, 2, 2,
	302, 303, 5, 38, 20, 2, 303, 304, 9, 2, 2, 2, 304, 305, 7, 41, 2, 2, 305,
	306, 9, 4, 2, 2, 306, 307, 7, 40, 2, 2, 307, 308, 5, 38, 20, 2, 308, 309,
	9, 2, 2, 2, 309, 310, 7, 41, 2, 2, 310, 311, 8, 20, 1, 2, 311, 351, 3,
	2, 2, 2, 312, 313, 7, 40, 2, 2, 313, 314, 5, 38, 20, 2, 314, 315, 9, 2,
	2, 2, 315, 316, 7, 41, 2, 2, 316, 317, 9, 5, 2, 2, 317, 318, 5, 38, 20,
	10, 318, 319, 8, 20, 1, 2, 319, 351, 3, 2, 2, 2, 320, 321, 7, 40, 2, 2,
	321, 322, 5, 38, 20, 2, 322, 323, 9, 2, 2, 2, 323, 324, 7, 41, 2, 2, 324,
	325, 9, 5, 2, 2, 325, 326, 7, 40, 2, 2, 326, 327, 5, 38, 20, 2, 327, 328,
	9, 2, 2, 2, 328, 329, 7, 41, 2, 2, 329, 330, 8, 20, 1, 2, 330, 351, 3,
	2, 2, 2, 331, 332, 9, 6, 2, 2, 332, 333, 5, 36, 19, 2, 333, 334, 8, 20,
	1, 2, 334, 351, 3, 2, 2, 2, 335, 336, 9, 6, 2, 2, 336, 337, 7, 40, 2, 2,
	337, 338, 5, 38, 20, 2, 338, 339, 9, 2, 2, 2, 339, 340, 7, 41, 2, 2, 340,
	341, 8, 20, 1, 2, 341, 351, 3, 2, 2, 2, 342, 343, 5, 40, 21, 2, 343, 344,
	8, 20, 1, 2, 344, 351, 3, 2, 2, 2, 345, 346, 7, 40, 2, 2, 346, 347, 5,
	36, 19, 2, 347, 348, 7, 41, 2, 2, 348, 349, 8, 20, 1, 2, 349, 351, 3, 2,
	2, 2, 350, 273, 3, 2, 2, 2, 350, 282, 3, 2, 2, 2, 350, 293, 3, 2, 2, 2,
	350, 301, 3, 2, 2, 2, 350, 312, 3, 2, 2, 2, 350, 320, 3, 2, 2, 2, 350,
	331, 3, 2, 2, 2, 350, 335, 3, 2, 2, 2, 350, 342, 3, 2, 2, 2, 350, 345,
	3, 2, 2, 2, 351, 398, 3, 2, 2, 2, 352, 353, 12, 19, 2, 2, 353, 354, 9,
	3, 2, 2, 354, 355, 5, 38, 20, 20, 355, 356, 8, 20, 1, 2, 356, 397, 3, 2,
	2, 2, 357, 358, 12, 15, 2, 2, 358, 359, 9, 4, 2, 2, 359, 360, 5, 38, 20,
	16, 360, 361, 8, 20, 1, 2, 361, 397, 3, 2, 2, 2, 362, 363, 12, 11, 2, 2,
	363, 364, 9, 5, 2, 2, 364, 365, 5, 38, 20, 12, 365, 366, 8, 20, 1, 2, 366,
	397, 3, 2, 2, 2, 367, 368, 12, 7, 2, 2, 368, 369, 9, 7, 2, 2, 369, 370,
	5, 38, 20, 8, 370, 371, 8, 20, 1, 2, 371, 397, 3, 2, 2, 2, 372, 373, 12,
	17, 2, 2, 373, 374, 9, 3, 2, 2, 374, 375, 7, 40, 2, 2, 375, 376, 5, 38,
	20, 2, 376, 377, 9, 2, 2, 2, 377, 378, 7, 41, 2, 2, 378, 379, 8, 20, 1,
	2, 379, 397, 3, 2, 2, 2, 380, 381, 12, 13, 2, 2, 381, 382, 9, 4, 2, 2,
	382, 383, 7, 40, 2, 2, 383, 384, 5, 38, 20, 2, 384, 385, 9, 2, 2, 2, 385,
	386, 7, 41, 2, 2, 386, 387, 8, 20, 1, 2, 387, 397, 3, 2, 2, 2, 388, 389,
	12, 9, 2, 2, 389, 390, 9, 5, 2, 2, 390, 391, 7, 40, 2, 2, 391, 392, 5,
	38, 20, 2, 392, 393, 9, 2, 2, 2, 393, 394, 7, 41, 2, 2, 394, 395, 8, 20,
	1, 2, 395, 397, 3, 2, 2, 2, 396, 352, 3, 2, 2, 2, 396, 357, 3, 2, 2, 2,
	396, 362, 3, 2, 2, 2, 396, 367, 3, 2, 2, 2, 396, 372, 3, 2, 2, 2, 396,
	380, 3, 2, 2, 2, 396, 388, 3, 2, 2, 2, 397, 400, 3, 2, 2, 2, 398, 396,
	3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399, 39, 3, 2, 2, 2, 400, 398, 3, 2,
	2, 2, 401, 402, 7, 17, 2, 2, 402, 415, 8, 21, 1, 2, 403, 404, 7, 18, 2,
	2, 404, 415, 8, 21, 1, 2, 405, 406, 7, 20, 2, 2, 406, 415, 8, 21, 1, 2,
	407, 408, 7, 21, 2, 2, 408, 415, 8, 21, 1, 2, 409, 410, 7, 22, 2, 2, 410,
	415, 8, 21, 1, 2, 411, 412, 5, 16, 9, 2, 412, 413, 8, 21, 1, 2, 413, 415,
	3, 2, 2, 2, 414, 401, 3, 2, 2, 2, 414, 403, 3, 2, 2, 2, 414, 405, 3, 2,
	2, 2, 414, 407, 3, 2, 2, 2, 414, 409, 3, 2, 2, 2, 414, 411, 3, 2, 2, 2,
	415, 41, 3, 2, 2, 2, 19, 48, 80, 116, 151, 156, 188, 193, 213, 225, 236,
	241, 254, 268, 350, 396, 398, 414,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'println!'", "'number'", "'as f64'", "'as i64'", "'let'", "'mut'",
	"'if'", "'else'", "'match'", "'i64'", "'f64'", "'String'", "'bool'", "'&str'",
	"", "", "", "", "", "", "'.'", "';'", "','", "':'", "'='", "'=='", "'>='",
	"'=>'", "'<='", "'!='", "'>'", "'<'", "'*'", "'/'", "'%'", "'+'", "'-'",
	"'('", "')'", "'{'", "'}'", "'['", "']'", "'&&'", "'||'", "'|'", "'!'",
	"'_ =>'",
}
var symbolicNames = []string{
	"", "R_PRINTLN", "P_NUMBER", "R_AS_DOUBLE", "R_AS_INTEGER", "R_LET", "R_MUT",
	"R_IF", "R_ELSE", "R_MATCH", "R_INT", "R_FLOAT", "R_STRING", "R_BOOL",
	"R_STR", "NUMBER", "DOUBLE", "CHAR", "STRING", "BOOLEAN", "ID", "TK_PUNTO",
	"TK_PUNTOCOMA", "TK_COMA", "TK_DOSPUNTOS", "TK_IGUAL", "TK_IGUALIGUAL",
	"TK_MAYORIGUAL", "TK_IGUALMAYOR", "TK_MENORIGUAL", "TK_DIFIGUAL", "TK_MAYOR",
	"TK_MENOR", "TK_MULT", "TK_DIV", "TK_MODULO", "TK_MAS", "TK_MENOS", "TK_PARA",
	"TK_PARC", "TK_LLAVEA", "TK_LLAVEC", "TK_CORA", "TK_CORC", "TK_AND", "TK_OR",
	"TK_BARRA", "TK_NOT", "TK_GUIONBAJO", "WHITESPACE", "TK_MULTI", "TK_LINE",
}

var ruleNames = []string{
	"start", "instrucciones", "instruccion", "instr_declaracion", "instr_asignacion",
	"instr_if", "instr_else_if", "instr_ternario", "instr_else_if_ternario",
	"instr_match", "instr_case", "list_case", "block_case", "list_expre_case",
	"instr_default", "block_default", "instr_tipo", "expression", "exp_arit",
	"primitivo",
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
	ChemsR_INT         = 10
	ChemsR_FLOAT       = 11
	ChemsR_STRING      = 12
	ChemsR_BOOL        = 13
	ChemsR_STR         = 14
	ChemsNUMBER        = 15
	ChemsDOUBLE        = 16
	ChemsCHAR          = 17
	ChemsSTRING        = 18
	ChemsBOOLEAN       = 19
	ChemsID            = 20
	ChemsTK_PUNTO      = 21
	ChemsTK_PUNTOCOMA  = 22
	ChemsTK_COMA       = 23
	ChemsTK_DOSPUNTOS  = 24
	ChemsTK_IGUAL      = 25
	ChemsTK_IGUALIGUAL = 26
	ChemsTK_MAYORIGUAL = 27
	ChemsTK_IGUALMAYOR = 28
	ChemsTK_MENORIGUAL = 29
	ChemsTK_DIFIGUAL   = 30
	ChemsTK_MAYOR      = 31
	ChemsTK_MENOR      = 32
	ChemsTK_MULT       = 33
	ChemsTK_DIV        = 34
	ChemsTK_MODULO     = 35
	ChemsTK_MAS        = 36
	ChemsTK_MENOS      = 37
	ChemsTK_PARA       = 38
	ChemsTK_PARC       = 39
	ChemsTK_LLAVEA     = 40
	ChemsTK_LLAVEC     = 41
	ChemsTK_CORA       = 42
	ChemsTK_CORC       = 43
	ChemsTK_AND        = 44
	ChemsTK_OR         = 45
	ChemsTK_BARRA      = 46
	ChemsTK_NOT        = 47
	ChemsTK_GUIONBAJO  = 48
	ChemsWHITESPACE    = 49
	ChemsTK_MULTI      = 50
	ChemsTK_LINE       = 51
)

// Chems rules.
const (
	ChemsRULE_start                  = 0
	ChemsRULE_instrucciones          = 1
	ChemsRULE_instruccion            = 2
	ChemsRULE_instr_declaracion      = 3
	ChemsRULE_instr_asignacion       = 4
	ChemsRULE_instr_if               = 5
	ChemsRULE_instr_else_if          = 6
	ChemsRULE_instr_ternario         = 7
	ChemsRULE_instr_else_if_ternario = 8
	ChemsRULE_instr_match            = 9
	ChemsRULE_instr_case             = 10
	ChemsRULE_list_case              = 11
	ChemsRULE_block_case             = 12
	ChemsRULE_list_expre_case        = 13
	ChemsRULE_instr_default          = 14
	ChemsRULE_block_default          = 15
	ChemsRULE_instr_tipo             = 16
	ChemsRULE_expression             = 17
	ChemsRULE_exp_arit               = 18
	ChemsRULE_primitivo              = 19
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
		p.SetState(40)

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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_PRINTLN)|(1<<ChemsR_LET)|(1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsID))) != 0) {
		{
			p.SetState(43)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).list = append(localctx.(*InstruccionesContext).list, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetList()
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

	// Get_instr_if returns the _instr_if rule contexts.
	Get_instr_if() IInstr_ifContext

	// Get_instr_match returns the _instr_match rule contexts.
	Get_instr_match() IInstr_matchContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_instr_declaracion sets the _instr_declaracion rule contexts.
	Set_instr_declaracion(IInstr_declaracionContext)

	// Set_instr_asignacion sets the _instr_asignacion rule contexts.
	Set_instr_asignacion(IInstr_asignacionContext)

	// Set_instr_if sets the _instr_if rule contexts.
	Set_instr_if(IInstr_ifContext)

	// Set_instr_match sets the _instr_match rule contexts.
	Set_instr_match(IInstr_matchContext)

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
	_instr_if          IInstr_ifContext
	_instr_match       IInstr_matchContext
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

func (s *InstruccionContext) Get_instr_if() IInstr_ifContext { return s._instr_if }

func (s *InstruccionContext) Get_instr_match() IInstr_matchContext { return s._instr_match }

func (s *InstruccionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *InstruccionContext) Set_instr_declaracion(v IInstr_declaracionContext) {
	s._instr_declaracion = v
}

func (s *InstruccionContext) Set_instr_asignacion(v IInstr_asignacionContext) { s._instr_asignacion = v }

func (s *InstruccionContext) Set_instr_if(v IInstr_ifContext) { s._instr_if = v }

func (s *InstruccionContext) Set_instr_match(v IInstr_matchContext) { s._instr_match = v }

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

	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Match(ChemsR_PRINTLN)
		}
		{
			p.SetState(51)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(52)

			var _x = p.Expression()

			localctx.(*InstruccionContext)._expression = _x
		}
		{
			p.SetState(53)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(54)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*InstruccionContext).instr = instruction.PRINTLN(localctx.(*InstruccionContext).Get_expression().GetP(), "-1")

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(ChemsR_PRINTLN)
		}
		{
			p.SetState(58)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(59)

			var _m = p.Match(ChemsSTRING)

			localctx.(*InstruccionContext)._STRING = _m
		}
		{
			p.SetState(60)
			p.Match(ChemsTK_COMA)
		}
		{
			p.SetState(61)

			var _x = p.Expression()

			localctx.(*InstruccionContext)._expression = _x
		}
		{
			p.SetState(62)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(63)
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
			p.SetState(66)

			var _x = p.Instr_declaracion()

			localctx.(*InstruccionContext)._instr_declaracion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_declaracion().GetInstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(69)

			var _x = p.Instr_asignacion()

			localctx.(*InstruccionContext)._instr_asignacion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_asignacion().GetInstr()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(72)

			var _x = p.Instr_if()

			localctx.(*InstruccionContext)._instr_if = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_if().GetInstr()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(75)

			var _x = p.Instr_match()

			localctx.(*InstruccionContext)._instr_match = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_match().GetInstr()

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

	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(81)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(82)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(83)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(84)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(85)
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
			p.SetState(88)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(89)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(90)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(91)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(92)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(93)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(94)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(95)
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
			p.SetState(98)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(99)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(100)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(101)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(102)
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
			p.SetState(105)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(106)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(107)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(108)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(109)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(110)

			var _x = p.Expression()

			localctx.(*Instr_declaracionContext)._expression = _x
		}
		{
			p.SetState(111)
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
		p.SetState(116)

		var _m = p.Match(ChemsID)

		localctx.(*Instr_asignacionContext)._ID = _m
	}
	{
		p.SetState(117)
		p.Match(ChemsTK_IGUAL)
	}
	{
		p.SetState(118)

		var _x = p.Expression()

		localctx.(*Instr_asignacionContext)._expression = _x
	}
	{
		p.SetState(119)
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
	p.EnterRule(localctx, 10, ChemsRULE_instr_if)

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

	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(123)

			var _x = p.Expression()

			localctx.(*Instr_ifContext)._expression = _x
		}
		{
			p.SetState(124)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(125)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(126)
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
			p.SetState(129)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(130)

			var _x = p.Expression()

			localctx.(*Instr_ifContext)._expression = _x
		}
		{
			p.SetState(131)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(132)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(133)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(134)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(135)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(136)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).right_intr = _x
		}
		{
			p.SetState(137)
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
			p.SetState(140)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(141)

			var _x = p.Expression()

			localctx.(*Instr_ifContext)._expression = _x
		}
		{
			p.SetState(142)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(143)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(144)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(145)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(146)

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
	p.EnterRule(localctx, 12, ChemsRULE_instr_else_if)

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
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(151)

				var _x = p.Instr_if()

				localctx.(*Instr_else_ifContext)._instr_if = _x
			}
			localctx.(*Instr_else_ifContext).e = append(localctx.(*Instr_else_ifContext).e, localctx.(*Instr_else_ifContext)._instr_if)

		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 14, ChemsRULE_instr_ternario)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(160)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).block = _x
		}
		{
			p.SetState(161)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(162)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(163)
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
			p.SetState(166)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(167)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).block = _x
		}
		{
			p.SetState(168)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(169)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(170)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(171)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(172)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(173)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).right_intr = _x
		}
		{
			p.SetState(174)
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
			p.SetState(177)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(178)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).block = _x
		}
		{
			p.SetState(179)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(180)

			var _x = p.Expression()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(181)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(182)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(183)

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
	p.EnterRule(localctx, 16, ChemsRULE_instr_else_if_ternario)

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
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(188)

				var _x = p.Instr_ternario()

				localctx.(*Instr_else_if_ternarioContext)._instr_ternario = _x
			}
			localctx.(*Instr_else_if_ternarioContext).e = append(localctx.(*Instr_else_if_ternarioContext).e, localctx.(*Instr_else_if_ternarioContext)._instr_ternario)

		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 18, ChemsRULE_instr_match)

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

	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(196)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_matchContext)._R_MATCH = _m
		}
		{
			p.SetState(197)

			var _x = p.Expression()

			localctx.(*Instr_matchContext)._expression = _x
		}
		{
			p.SetState(198)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(199)

			var _x = p.List_case()

			localctx.(*Instr_matchContext)._list_case = _x
		}
		{
			p.SetState(200)

			var _x = p.Block_default()

			localctx.(*Instr_matchContext)._block_default = _x
		}
		{
			p.SetState(201)
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
			p.SetState(204)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_matchContext)._R_MATCH = _m
		}
		{
			p.SetState(205)

			var _x = p.Expression()

			localctx.(*Instr_matchContext)._expression = _x
		}
		{
			p.SetState(206)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(207)

			var _x = p.Block_default()

			localctx.(*Instr_matchContext)._block_default = _x
		}
		{
			p.SetState(208)
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

	// Set_list_expre_case sets the _list_expre_case rule contexts.
	Set_list_expre_case(IList_expre_caseContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Expresion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Expresion)

	// IsInstr_caseContext differentiates from other interfaces.
	IsInstr_caseContext()
}

type Instr_caseContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	instr            interfaces.Expresion
	_list_expre_case IList_expre_caseContext
	_instrucciones   IInstruccionesContext
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

func (s *Instr_caseContext) Set_list_expre_case(v IList_expre_caseContext) { s._list_expre_case = v }

func (s *Instr_caseContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

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
	p.EnterRule(localctx, 20, ChemsRULE_instr_case)

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
		p.SetState(213)

		var _x = p.List_expre_case()

		localctx.(*Instr_caseContext)._list_expre_case = _x
	}
	{
		p.SetState(214)
		p.Match(ChemsTK_IGUALMAYOR)
	}
	{
		p.SetState(215)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(216)

		var _x = p.Instrucciones()

		localctx.(*Instr_caseContext)._instrucciones = _x
	}
	{
		p.SetState(217)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_caseContext).instr = control.NewCase(nil, localctx.(*Instr_caseContext).Get_list_expre_case().GetL(), localctx.(*Instr_caseContext).Get_instrucciones().GetL())

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
	p.EnterRule(localctx, 22, ChemsRULE_list_case)

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
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ChemsTK_MENOS-37))|(1<<(ChemsTK_PARA-37))|(1<<(ChemsTK_NOT-37)))) != 0) {
		{
			p.SetState(220)

			var _x = p.Instr_case()

			localctx.(*List_caseContext)._instr_case = _x
		}
		localctx.(*List_caseContext).e = append(localctx.(*List_caseContext).e, localctx.(*List_caseContext)._instr_case)

		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_caseContext).GetE()
	for _, e := range listInt {
		localctx.(*List_caseContext).l.Add(e.GetInstr())
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
	p.EnterRule(localctx, 24, ChemsRULE_block_case)

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

	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)

			var _x = p.Expression()

			localctx.(*Block_caseContext)._expression = _x
		}
		{
			p.SetState(228)
			p.Match(ChemsTK_BARRA)
		}
		localctx.(*Block_caseContext).instr = control.NewCase(localctx.(*Block_caseContext).Get_expression().GetP(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)

			var _x = p.Expression()

			localctx.(*Block_caseContext)._expression = _x
		}
		localctx.(*Block_caseContext).instr = control.NewCase(localctx.(*Block_caseContext).Get_expression().GetP(), nil, nil)

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
	p.EnterRule(localctx, 26, ChemsRULE_list_expre_case)

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
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ChemsTK_MENOS-37))|(1<<(ChemsTK_PARA-37))|(1<<(ChemsTK_NOT-37)))) != 0) {
		{
			p.SetState(236)

			var _x = p.Block_case()

			localctx.(*List_expre_caseContext)._block_case = _x
		}
		localctx.(*List_expre_caseContext).e = append(localctx.(*List_expre_caseContext).e, localctx.(*List_expre_caseContext)._block_case)

		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_expre_caseContext).GetE()
	for _, e := range listInt {
		localctx.(*List_expre_caseContext).l.Add(e.GetInstr())
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

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_defaultContext differentiates from other interfaces.
	IsInstr_defaultContext()
}

type Instr_defaultContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_instrucciones IInstruccionesContext
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

func (s *Instr_defaultContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

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
	p.EnterRule(localctx, 28, ChemsRULE_instr_default)

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
		p.SetState(243)
		p.Match(ChemsTK_GUIONBAJO)
	}
	{
		p.SetState(244)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(245)

		var _x = p.Instrucciones()

		localctx.(*Instr_defaultContext)._instrucciones = _x
	}
	{
		p.SetState(246)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_defaultContext).instr = control.NewDefault(localctx.(*Instr_defaultContext).Get_instrucciones().GetL())

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
	p.EnterRule(localctx, 30, ChemsRULE_block_default)

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
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ChemsTK_GUIONBAJO {
		{
			p.SetState(249)

			var _x = p.Instr_default()

			localctx.(*Block_defaultContext)._instr_default = _x
		}
		localctx.(*Block_defaultContext).e = append(localctx.(*Block_defaultContext).e, localctx.(*Block_defaultContext)._instr_default)

		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*Block_defaultContext).GetE()
	for _, e := range listInt {
		localctx.(*Block_defaultContext).l.Add(e.GetInstr())
	}

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
	p.EnterRule(localctx, 32, ChemsRULE_instr_tipo)

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

	p.SetState(266)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.Match(ChemsR_INT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.INTEGER

	case ChemsR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(258)
			p.Match(ChemsR_FLOAT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.FLOAT

	case ChemsR_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			p.Match(ChemsR_STRING)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STRING

	case ChemsR_STR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(262)
			p.Match(ChemsR_STR)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STRING

	case ChemsR_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(264)
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
	p.EnterRule(localctx, 34, ChemsRULE_expression)

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
		p.SetState(268)

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
	_startState := 36
	p.EnterRecursionRule(localctx, 36, ChemsRULE_exp_arit, _p)
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
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(272)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(273)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(274)

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
			p.SetState(275)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(276)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(ChemsTK_MULT-33))|(1<<(ChemsTK_DIV-33))|(1<<(ChemsTK_MODULO-33)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(277)

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
			p.SetState(280)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(281)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(282)

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
			p.SetState(283)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(284)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(ChemsTK_MULT-33))|(1<<(ChemsTK_DIV-33))|(1<<(ChemsTK_MODULO-33)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(285)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(286)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(287)

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
			p.SetState(288)
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
			p.SetState(291)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(292)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(293)

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
			p.SetState(294)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(295)

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
			p.SetState(296)

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
			p.SetState(299)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(300)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(301)

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
			p.SetState(302)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(303)

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
			p.SetState(304)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(305)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(306)

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
			p.SetState(307)
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
			p.SetState(310)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(311)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(312)

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
			p.SetState(313)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(314)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(ChemsTK_IGUALIGUAL-26))|(1<<(ChemsTK_MAYORIGUAL-26))|(1<<(ChemsTK_MENORIGUAL-26))|(1<<(ChemsTK_DIFIGUAL-26))|(1<<(ChemsTK_MAYOR-26))|(1<<(ChemsTK_MENOR-26)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(315)

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
			p.SetState(318)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(319)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(320)

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
			p.SetState(321)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(322)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Exp_aritContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(ChemsTK_IGUALIGUAL-26))|(1<<(ChemsTK_MAYORIGUAL-26))|(1<<(ChemsTK_MENORIGUAL-26))|(1<<(ChemsTK_DIFIGUAL-26))|(1<<(ChemsTK_MAYOR-26))|(1<<(ChemsTK_MENOR-26)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Exp_aritContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(323)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(324)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).right = _x
		}
		{
			p.SetState(325)

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
			p.SetState(326)
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
			p.SetState(329)

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
			p.SetState(330)

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
			p.SetState(333)

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
			p.SetState(334)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(335)

			var _x = p.exp_arit(0)

			localctx.(*Exp_aritContext).left = _x
		}
		{
			p.SetState(336)

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
			p.SetState(337)
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
			p.SetState(340)

			var _x = p.Primitivo()

			localctx.(*Exp_aritContext)._primitivo = _x
		}
		localctx.(*Exp_aritContext).p = localctx.(*Exp_aritContext).Get_primitivo().GetP()

	case 10:
		{
			p.SetState(343)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(344)

			var _x = p.Expression()

			localctx.(*Exp_aritContext)._expression = _x
		}
		{
			p.SetState(345)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Exp_aritContext).p = localctx.(*Exp_aritContext).Get_expression().GetP()

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(394)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExp_aritContext(p, _parentctx, _parentState)
				localctx.(*Exp_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_exp_arit)
				p.SetState(350)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(351)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(ChemsTK_MULT-33))|(1<<(ChemsTK_DIV-33))|(1<<(ChemsTK_MODULO-33)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(352)

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
				p.SetState(355)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(356)

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
					p.SetState(357)

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
				p.SetState(360)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(361)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(ChemsTK_IGUALIGUAL-26))|(1<<(ChemsTK_MAYORIGUAL-26))|(1<<(ChemsTK_MENORIGUAL-26))|(1<<(ChemsTK_DIFIGUAL-26))|(1<<(ChemsTK_MAYOR-26))|(1<<(ChemsTK_MENOR-26)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(362)

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
				p.SetState(365)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(366)

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
					p.SetState(367)

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
				p.SetState(370)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(371)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(ChemsTK_MULT-33))|(1<<(ChemsTK_DIV-33))|(1<<(ChemsTK_MODULO-33)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(372)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(373)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(374)

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
					p.SetState(375)
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
				p.SetState(378)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(379)

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
					p.SetState(380)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(381)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(382)

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
					p.SetState(383)
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
				p.SetState(386)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(387)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(ChemsTK_IGUALIGUAL-26))|(1<<(ChemsTK_MAYORIGUAL-26))|(1<<(ChemsTK_MENORIGUAL-26))|(1<<(ChemsTK_DIFIGUAL-26))|(1<<(ChemsTK_MAYOR-26))|(1<<(ChemsTK_MENOR-26)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(388)
					p.Match(ChemsTK_PARA)
				}
				{
					p.SetState(389)

					var _x = p.exp_arit(0)

					localctx.(*Exp_aritContext).right = _x
				}
				{
					p.SetState(390)

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
					p.SetState(391)
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
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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

	// Set_instr_ternario sets the _instr_ternario rule contexts.
	Set_instr_ternario(IInstr_ternarioContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	p               interfaces.Expresion
	_NUMBER         antlr.Token
	_DOUBLE         antlr.Token
	_STRING         antlr.Token
	_BOOLEAN        antlr.Token
	_ID             antlr.Token
	_instr_ternario IInstr_ternarioContext
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

func (s *PrimitivoContext) Set_instr_ternario(v IInstr_ternarioContext) { s._instr_ternario = v }

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
	p.EnterRule(localctx, 38, ChemsRULE_primitivo)

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

	p.SetState(412)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(399)

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
			p.SetState(401)

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
			p.SetState(403)

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
			p.SetState(405)

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
			p.SetState(407)

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
			p.SetState(409)

			var _x = p.Instr_ternario()

			localctx.(*PrimitivoContext)._instr_ternario = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_ternario().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
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

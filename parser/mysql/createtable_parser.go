// Code generated from CreateTable.g4 by ANTLR 4.7.2. DO NOT EDIT.

package mysql // CreateTable
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 98, 419,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 3, 2, 5, 2, 49, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 55, 10, 3, 3,
	3, 3, 3, 5, 3, 59, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 65, 10, 3, 3, 3,
	7, 3, 68, 10, 3, 12, 3, 14, 3, 71, 11, 3, 5, 3, 73, 10, 3, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 85, 10, 6, 12, 6,
	14, 6, 88, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 5, 7, 95, 10, 7, 3, 8,
	3, 8, 7, 8, 99, 10, 8, 12, 8, 14, 8, 102, 11, 8, 3, 9, 3, 9, 5, 9, 106,
	10, 9, 3, 9, 5, 9, 109, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 114, 10, 9, 3, 9,
	5, 9, 117, 10, 9, 3, 9, 3, 9, 5, 9, 121, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9,
	126, 10, 9, 3, 9, 5, 9, 129, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 134, 10, 9,
	3, 9, 5, 9, 137, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 143, 10, 9, 3, 9,
	5, 9, 146, 10, 9, 3, 9, 3, 9, 5, 9, 150, 10, 9, 3, 9, 5, 9, 153, 10, 9,
	3, 9, 5, 9, 156, 10, 9, 3, 9, 3, 9, 5, 9, 160, 10, 9, 3, 9, 5, 9, 163,
	10, 9, 3, 9, 5, 9, 166, 10, 9, 3, 9, 3, 9, 5, 9, 170, 10, 9, 3, 9, 5, 9,
	173, 10, 9, 3, 9, 5, 9, 176, 10, 9, 3, 9, 5, 9, 179, 10, 9, 3, 9, 3, 9,
	5, 9, 183, 10, 9, 3, 9, 5, 9, 186, 10, 9, 3, 9, 5, 9, 189, 10, 9, 3, 9,
	3, 9, 3, 9, 5, 9, 194, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 199, 10, 9, 3, 9,
	3, 9, 3, 9, 5, 9, 204, 10, 9, 3, 9, 5, 9, 207, 10, 9, 5, 9, 209, 10, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 5, 12, 225, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 7, 13, 233, 10, 13, 12, 13, 14, 13, 236, 11, 13, 3, 13, 3,
	13, 3, 14, 5, 14, 241, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	248, 10, 14, 3, 14, 3, 14, 3, 14, 5, 14, 253, 10, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 262, 10, 14, 3, 15, 3, 15, 5, 15,
	266, 10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 273, 10, 15, 5,
	15, 275, 10, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 5, 17, 286, 10, 17, 3, 17, 5, 17, 289, 10, 17, 3, 18, 3, 18, 3,
	18, 5, 18, 294, 10, 18, 3, 18, 5, 18, 297, 10, 18, 3, 18, 3, 18, 3, 18,
	5, 18, 302, 10, 18, 3, 18, 5, 18, 305, 10, 18, 3, 19, 3, 19, 5, 19, 309,
	10, 19, 5, 19, 311, 10, 19, 3, 19, 3, 19, 3, 19, 5, 19, 316, 10, 19, 3,
	19, 5, 19, 319, 10, 19, 3, 19, 3, 19, 7, 19, 323, 10, 19, 12, 19, 14, 19,
	326, 11, 19, 3, 19, 3, 19, 5, 19, 330, 10, 19, 5, 19, 332, 10, 19, 3, 19,
	5, 19, 335, 10, 19, 3, 19, 5, 19, 338, 10, 19, 3, 19, 5, 19, 341, 10, 19,
	3, 19, 5, 19, 344, 10, 19, 3, 19, 3, 19, 7, 19, 348, 10, 19, 12, 19, 14,
	19, 351, 11, 19, 5, 19, 353, 10, 19, 3, 20, 3, 20, 5, 20, 357, 10, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 365, 10, 20, 3, 21, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 375, 10, 22, 12, 22,
	14, 22, 378, 11, 22, 3, 22, 3, 22, 3, 23, 3, 23, 5, 23, 384, 10, 23, 3,
	23, 3, 23, 3, 23, 5, 23, 389, 10, 23, 3, 23, 3, 23, 5, 23, 393, 10, 23,
	3, 23, 3, 23, 3, 23, 5, 23, 398, 10, 23, 3, 23, 5, 23, 401, 10, 23, 3,
	23, 3, 23, 5, 23, 405, 10, 23, 3, 23, 3, 23, 5, 23, 409, 10, 23, 3, 23,
	3, 23, 3, 23, 5, 23, 414, 10, 23, 3, 23, 5, 23, 417, 10, 23, 3, 23, 2,
	2, 24, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 2, 19, 3, 2, 96, 97, 10, 2, 22, 22, 51, 51, 54, 54,
	56, 56, 61, 61, 73, 73, 78, 78, 86, 86, 4, 2, 14, 14, 96, 97, 4, 2, 20,
	20, 86, 86, 4, 2, 20, 20, 22, 22, 7, 2, 13, 13, 43, 44, 53, 53, 70, 70,
	77, 77, 4, 2, 69, 69, 81, 81, 5, 2, 30, 31, 38, 39, 60, 60, 10, 2, 16,
	16, 18, 19, 29, 29, 45, 45, 50, 50, 52, 52, 67, 67, 76, 76, 7, 2, 14, 15,
	28, 28, 74, 75, 85, 85, 89, 89, 4, 2, 35, 35, 68, 68, 4, 2, 7, 10, 57,
	57, 4, 2, 37, 37, 79, 79, 4, 2, 27, 27, 48, 49, 4, 2, 94, 94, 97, 97, 4,
	2, 41, 41, 47, 47, 4, 2, 17, 17, 40, 40, 2, 499, 2, 46, 3, 2, 2, 2, 4,
	52, 3, 2, 2, 2, 6, 74, 3, 2, 2, 2, 8, 76, 3, 2, 2, 2, 10, 80, 3, 2, 2,
	2, 12, 94, 3, 2, 2, 2, 14, 96, 3, 2, 2, 2, 16, 208, 3, 2, 2, 2, 18, 210,
	3, 2, 2, 2, 20, 214, 3, 2, 2, 2, 22, 220, 3, 2, 2, 2, 24, 228, 3, 2, 2,
	2, 26, 261, 3, 2, 2, 2, 28, 274, 3, 2, 2, 2, 30, 276, 3, 2, 2, 2, 32, 288,
	3, 2, 2, 2, 34, 304, 3, 2, 2, 2, 36, 352, 3, 2, 2, 2, 38, 364, 3, 2, 2,
	2, 40, 366, 3, 2, 2, 2, 42, 369, 3, 2, 2, 2, 44, 416, 3, 2, 2, 2, 46, 48,
	5, 4, 3, 2, 47, 49, 7, 3, 2, 2, 48, 47, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2,
	49, 50, 3, 2, 2, 2, 50, 51, 7, 2, 2, 3, 51, 3, 3, 2, 2, 2, 52, 54, 7, 26,
	2, 2, 53, 55, 7, 72, 2, 2, 54, 53, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55,
	56, 3, 2, 2, 2, 56, 58, 7, 71, 2, 2, 57, 59, 5, 8, 5, 2, 58, 57, 3, 2,
	2, 2, 58, 59, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 5, 6, 4, 2, 61, 72,
	5, 10, 6, 2, 62, 69, 5, 44, 23, 2, 63, 65, 7, 4, 2, 2, 64, 63, 3, 2, 2,
	2, 64, 65, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 68, 5, 44, 23, 2, 67, 64,
	3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2,
	70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 62, 3, 2, 2, 2, 72, 73, 3,
	2, 2, 2, 73, 5, 3, 2, 2, 2, 74, 75, 9, 2, 2, 2, 75, 7, 3, 2, 2, 2, 76,
	77, 7, 42, 2, 2, 77, 78, 7, 57, 2, 2, 78, 79, 7, 36, 2, 2, 79, 9, 3, 2,
	2, 2, 80, 81, 7, 5, 2, 2, 81, 86, 5, 12, 7, 2, 82, 83, 7, 4, 2, 2, 83,
	85, 5, 12, 7, 2, 84, 82, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86, 84, 3, 2,
	2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 89, 90,
	7, 6, 2, 2, 90, 11, 3, 2, 2, 2, 91, 92, 9, 2, 2, 2, 92, 95, 5, 14, 8, 2,
	93, 95, 5, 36, 19, 2, 94, 91, 3, 2, 2, 2, 94, 93, 3, 2, 2, 2, 95, 13, 3,
	2, 2, 2, 96, 100, 5, 16, 9, 2, 97, 99, 5, 26, 14, 2, 98, 97, 3, 2, 2, 2,
	99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 15,
	3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 105, 9, 3, 2, 2, 104, 106, 5, 18,
	10, 2, 105, 104, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 108, 3, 2, 2, 2,
	107, 109, 7, 14, 2, 2, 108, 107, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109,
	116, 3, 2, 2, 2, 110, 111, 7, 20, 2, 2, 111, 114, 7, 68, 2, 2, 112, 114,
	7, 21, 2, 2, 113, 110, 3, 2, 2, 2, 113, 112, 3, 2, 2, 2, 114, 115, 3, 2,
	2, 2, 115, 117, 9, 4, 2, 2, 116, 113, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2,
	117, 120, 3, 2, 2, 2, 118, 119, 7, 23, 2, 2, 119, 121, 7, 94, 2, 2, 120,
	118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 209, 3, 2, 2, 2, 122, 123,
	7, 55, 2, 2, 123, 125, 9, 5, 2, 2, 124, 126, 5, 18, 10, 2, 125, 124, 3,
	2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 128, 3, 2, 2, 2, 127, 129, 7, 14, 2,
	2, 128, 127, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 209, 3, 2, 2, 2, 130,
	131, 7, 56, 2, 2, 131, 133, 7, 86, 2, 2, 132, 134, 5, 18, 10, 2, 133, 132,
	3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 136, 3, 2, 2, 2, 135, 137, 7, 14,
	2, 2, 136, 135, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 209, 3, 2, 2, 2,
	138, 139, 7, 55, 2, 2, 139, 140, 9, 6, 2, 2, 140, 142, 7, 87, 2, 2, 141,
	143, 5, 18, 10, 2, 142, 141, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145,
	3, 2, 2, 2, 144, 146, 7, 14, 2, 2, 145, 144, 3, 2, 2, 2, 145, 146, 3, 2,
	2, 2, 146, 209, 3, 2, 2, 2, 147, 149, 9, 7, 2, 2, 148, 150, 5, 18, 10,
	2, 149, 148, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 152, 3, 2, 2, 2, 151,
	153, 9, 8, 2, 2, 152, 151, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 155,
	3, 2, 2, 2, 154, 156, 7, 90, 2, 2, 155, 154, 3, 2, 2, 2, 155, 156, 3, 2,
	2, 2, 156, 209, 3, 2, 2, 2, 157, 159, 7, 66, 2, 2, 158, 160, 5, 20, 11,
	2, 159, 158, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 162, 3, 2, 2, 2, 161,
	163, 9, 8, 2, 2, 162, 161, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 165,
	3, 2, 2, 2, 164, 166, 7, 90, 2, 2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2,
	2, 2, 166, 209, 3, 2, 2, 2, 167, 169, 7, 33, 2, 2, 168, 170, 7, 64, 2,
	2, 169, 168, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 172, 3, 2, 2, 2, 171,
	173, 5, 20, 11, 2, 172, 171, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 175,
	3, 2, 2, 2, 174, 176, 9, 8, 2, 2, 175, 174, 3, 2, 2, 2, 175, 176, 3, 2,
	2, 2, 176, 178, 3, 2, 2, 2, 177, 179, 7, 90, 2, 2, 178, 177, 3, 2, 2, 2,
	178, 179, 3, 2, 2, 2, 179, 209, 3, 2, 2, 2, 180, 182, 9, 9, 2, 2, 181,
	183, 5, 22, 12, 2, 182, 181, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 185,
	3, 2, 2, 2, 184, 186, 9, 8, 2, 2, 185, 184, 3, 2, 2, 2, 185, 186, 3, 2,
	2, 2, 186, 188, 3, 2, 2, 2, 187, 189, 7, 90, 2, 2, 188, 187, 3, 2, 2, 2,
	188, 189, 3, 2, 2, 2, 189, 209, 3, 2, 2, 2, 190, 209, 9, 10, 2, 2, 191,
	193, 9, 11, 2, 2, 192, 194, 5, 18, 10, 2, 193, 192, 3, 2, 2, 2, 193, 194,
	3, 2, 2, 2, 194, 209, 3, 2, 2, 2, 195, 196, 9, 12, 2, 2, 196, 198, 5, 24,
	13, 2, 197, 199, 7, 14, 2, 2, 198, 197, 3, 2, 2, 2, 198, 199, 3, 2, 2,
	2, 199, 206, 3, 2, 2, 2, 200, 201, 7, 20, 2, 2, 201, 204, 7, 68, 2, 2,
	202, 204, 7, 21, 2, 2, 203, 200, 3, 2, 2, 2, 203, 202, 3, 2, 2, 2, 204,
	205, 3, 2, 2, 2, 205, 207, 9, 4, 2, 2, 206, 203, 3, 2, 2, 2, 206, 207,
	3, 2, 2, 2, 207, 209, 3, 2, 2, 2, 208, 103, 3, 2, 2, 2, 208, 122, 3, 2,
	2, 2, 208, 130, 3, 2, 2, 2, 208, 138, 3, 2, 2, 2, 208, 147, 3, 2, 2, 2,
	208, 157, 3, 2, 2, 2, 208, 167, 3, 2, 2, 2, 208, 180, 3, 2, 2, 2, 208,
	190, 3, 2, 2, 2, 208, 191, 3, 2, 2, 2, 208, 195, 3, 2, 2, 2, 209, 17, 3,
	2, 2, 2, 210, 211, 7, 5, 2, 2, 211, 212, 7, 91, 2, 2, 212, 213, 7, 6, 2,
	2, 213, 19, 3, 2, 2, 2, 214, 215, 7, 5, 2, 2, 215, 216, 7, 91, 2, 2, 216,
	217, 7, 4, 2, 2, 217, 218, 7, 91, 2, 2, 218, 219, 7, 6, 2, 2, 219, 21,
	3, 2, 2, 2, 220, 221, 7, 5, 2, 2, 221, 224, 7, 91, 2, 2, 222, 223, 7, 4,
	2, 2, 223, 225, 7, 91, 2, 2, 224, 222, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2,
	225, 226, 3, 2, 2, 2, 226, 227, 7, 6, 2, 2, 227, 23, 3, 2, 2, 2, 228, 229,
	7, 5, 2, 2, 229, 234, 7, 94, 2, 2, 230, 231, 7, 4, 2, 2, 231, 233, 7, 94,
	2, 2, 232, 230, 3, 2, 2, 2, 233, 236, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2,
	234, 235, 3, 2, 2, 2, 235, 237, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 237,
	238, 7, 6, 2, 2, 238, 25, 3, 2, 2, 2, 239, 241, 7, 57, 2, 2, 240, 239,
	3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 262, 7, 59,
	2, 2, 243, 244, 7, 32, 2, 2, 244, 262, 5, 28, 15, 2, 245, 262, 7, 12, 2,
	2, 246, 248, 7, 65, 2, 2, 247, 246, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248,
	249, 3, 2, 2, 2, 249, 262, 7, 47, 2, 2, 250, 252, 7, 80, 2, 2, 251, 253,
	7, 47, 2, 2, 252, 251, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 262, 3, 2,
	2, 2, 254, 255, 7, 24, 2, 2, 255, 262, 7, 94, 2, 2, 256, 257, 7, 23, 2,
	2, 257, 262, 7, 96, 2, 2, 258, 259, 7, 67, 2, 2, 259, 260, 7, 32, 2, 2,
	260, 262, 7, 84, 2, 2, 261, 240, 3, 2, 2, 2, 261, 243, 3, 2, 2, 2, 261,
	245, 3, 2, 2, 2, 261, 247, 3, 2, 2, 2, 261, 250, 3, 2, 2, 2, 261, 254,
	3, 2, 2, 2, 261, 256, 3, 2, 2, 2, 261, 258, 3, 2, 2, 2, 262, 27, 3, 2,
	2, 2, 263, 275, 7, 59, 2, 2, 264, 266, 5, 30, 16, 2, 265, 264, 3, 2, 2,
	2, 265, 266, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 275, 5, 32, 17, 2,
	268, 272, 5, 34, 18, 2, 269, 270, 7, 62, 2, 2, 270, 271, 7, 82, 2, 2, 271,
	273, 5, 34, 18, 2, 272, 269, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 275,
	3, 2, 2, 2, 274, 263, 3, 2, 2, 2, 274, 265, 3, 2, 2, 2, 274, 268, 3, 2,
	2, 2, 275, 29, 3, 2, 2, 2, 276, 277, 9, 13, 2, 2, 277, 31, 3, 2, 2, 2,
	278, 289, 7, 94, 2, 2, 279, 289, 7, 91, 2, 2, 280, 281, 7, 10, 2, 2, 281,
	289, 7, 91, 2, 2, 282, 289, 9, 14, 2, 2, 283, 289, 7, 92, 2, 2, 284, 286,
	7, 57, 2, 2, 285, 284, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 287, 3, 2,
	2, 2, 287, 289, 7, 59, 2, 2, 288, 278, 3, 2, 2, 2, 288, 279, 3, 2, 2, 2,
	288, 280, 3, 2, 2, 2, 288, 282, 3, 2, 2, 2, 288, 283, 3, 2, 2, 2, 288,
	285, 3, 2, 2, 2, 289, 33, 3, 2, 2, 2, 290, 296, 9, 15, 2, 2, 291, 293,
	7, 5, 2, 2, 292, 294, 7, 91, 2, 2, 293, 292, 3, 2, 2, 2, 293, 294, 3, 2,
	2, 2, 294, 295, 3, 2, 2, 2, 295, 297, 7, 6, 2, 2, 296, 291, 3, 2, 2, 2,
	296, 297, 3, 2, 2, 2, 297, 305, 3, 2, 2, 2, 298, 299, 7, 58, 2, 2, 299,
	301, 7, 5, 2, 2, 300, 302, 7, 91, 2, 2, 301, 300, 3, 2, 2, 2, 301, 302,
	3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 305, 7, 6, 2, 2, 304, 290, 3, 2,
	2, 2, 304, 298, 3, 2, 2, 2, 305, 35, 3, 2, 2, 2, 306, 308, 7, 25, 2, 2,
	307, 309, 9, 16, 2, 2, 308, 307, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309,
	311, 3, 2, 2, 2, 310, 306, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 312,
	3, 2, 2, 2, 312, 313, 7, 65, 2, 2, 313, 315, 7, 47, 2, 2, 314, 316, 9,
	16, 2, 2, 315, 314, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 318, 3, 2, 2,
	2, 317, 319, 9, 2, 2, 2, 318, 317, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319,
	320, 3, 2, 2, 2, 320, 324, 5, 42, 22, 2, 321, 323, 5, 38, 20, 2, 322, 321,
	3, 2, 2, 2, 323, 326, 3, 2, 2, 2, 324, 322, 3, 2, 2, 2, 324, 325, 3, 2,
	2, 2, 325, 353, 3, 2, 2, 2, 326, 324, 3, 2, 2, 2, 327, 329, 7, 25, 2, 2,
	328, 330, 9, 16, 2, 2, 329, 328, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330,
	332, 3, 2, 2, 2, 331, 327, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 334,
	3, 2, 2, 2, 333, 335, 7, 80, 2, 2, 334, 333, 3, 2, 2, 2, 334, 335, 3, 2,
	2, 2, 335, 337, 3, 2, 2, 2, 336, 338, 9, 17, 2, 2, 337, 336, 3, 2, 2, 2,
	337, 338, 3, 2, 2, 2, 338, 340, 3, 2, 2, 2, 339, 341, 9, 2, 2, 2, 340,
	339, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2, 341, 343, 3, 2, 2, 2, 342, 344,
	5, 40, 21, 2, 343, 342, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 345, 3,
	2, 2, 2, 345, 349, 5, 42, 22, 2, 346, 348, 5, 38, 20, 2, 347, 346, 3, 2,
	2, 2, 348, 351, 3, 2, 2, 2, 349, 347, 3, 2, 2, 2, 349, 350, 3, 2, 2, 2,
	350, 353, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2, 352, 310, 3, 2, 2, 2, 352,
	331, 3, 2, 2, 2, 353, 37, 3, 2, 2, 2, 354, 356, 7, 46, 2, 2, 355, 357,
	7, 11, 2, 2, 356, 355, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357, 358, 3, 2,
	2, 2, 358, 365, 7, 93, 2, 2, 359, 360, 7, 88, 2, 2, 360, 361, 7, 63, 2,
	2, 361, 365, 9, 16, 2, 2, 362, 363, 7, 24, 2, 2, 363, 365, 7, 94, 2, 2,
	364, 354, 3, 2, 2, 2, 364, 359, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 365,
	39, 3, 2, 2, 2, 366, 367, 7, 83, 2, 2, 367, 368, 9, 18, 2, 2, 368, 41,
	3, 2, 2, 2, 369, 370, 7, 5, 2, 2, 370, 376, 9, 2, 2, 2, 371, 372, 7, 4,
	2, 2, 372, 375, 7, 96, 2, 2, 373, 375, 7, 97, 2, 2, 374, 371, 3, 2, 2,
	2, 374, 373, 3, 2, 2, 2, 375, 378, 3, 2, 2, 2, 376, 374, 3, 2, 2, 2, 376,
	377, 3, 2, 2, 2, 377, 379, 3, 2, 2, 2, 378, 376, 3, 2, 2, 2, 379, 380,
	7, 6, 2, 2, 380, 43, 3, 2, 2, 2, 381, 383, 7, 34, 2, 2, 382, 384, 7, 11,
	2, 2, 383, 382, 3, 2, 2, 2, 383, 384, 3, 2, 2, 2, 384, 385, 3, 2, 2, 2,
	385, 417, 9, 2, 2, 2, 386, 388, 7, 12, 2, 2, 387, 389, 7, 11, 2, 2, 388,
	387, 3, 2, 2, 2, 388, 389, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 417,
	7, 91, 2, 2, 391, 393, 7, 32, 2, 2, 392, 391, 3, 2, 2, 2, 392, 393, 3,
	2, 2, 2, 393, 397, 3, 2, 2, 2, 394, 395, 7, 20, 2, 2, 395, 398, 7, 68,
	2, 2, 396, 398, 7, 21, 2, 2, 397, 394, 3, 2, 2, 2, 397, 396, 3, 2, 2, 2,
	398, 400, 3, 2, 2, 2, 399, 401, 7, 11, 2, 2, 400, 399, 3, 2, 2, 2, 400,
	401, 3, 2, 2, 2, 401, 402, 3, 2, 2, 2, 402, 417, 9, 2, 2, 2, 403, 405,
	7, 32, 2, 2, 404, 403, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 406, 3, 2,
	2, 2, 406, 408, 7, 23, 2, 2, 407, 409, 7, 11, 2, 2, 408, 407, 3, 2, 2,
	2, 408, 409, 3, 2, 2, 2, 409, 410, 3, 2, 2, 2, 410, 417, 9, 2, 2, 2, 411,
	413, 7, 24, 2, 2, 412, 414, 7, 11, 2, 2, 413, 412, 3, 2, 2, 2, 413, 414,
	3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415, 417, 7, 94, 2, 2, 416, 381, 3, 2,
	2, 2, 416, 386, 3, 2, 2, 2, 416, 392, 3, 2, 2, 2, 416, 404, 3, 2, 2, 2,
	416, 411, 3, 2, 2, 2, 417, 45, 3, 2, 2, 2, 81, 48, 54, 58, 64, 69, 72,
	86, 94, 100, 105, 108, 113, 116, 120, 125, 128, 133, 136, 142, 145, 149,
	152, 155, 159, 162, 165, 169, 172, 175, 178, 182, 185, 188, 193, 198, 203,
	206, 208, 224, 234, 240, 247, 252, 261, 265, 272, 274, 285, 288, 293, 296,
	301, 304, 308, 310, 315, 318, 324, 329, 331, 334, 337, 340, 343, 349, 352,
	356, 364, 374, 376, 383, 388, 392, 397, 400, 404, 408, 413, 416,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "','", "'('", "')'", "'!'", "'~'", "'+'", "'-'", "'='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "AUTO_INCREMENT", "BIGINT", "BINARY",
	"BIT", "BLOB", "BTREE", "BOOLEAN", "BOOL", "CHARACTER", "CHARSET", "CHAR",
	"COLLATE", "COMMENT", "CONSTRAINT", "CREATE", "CURRENT_TIMESTAMP", "DATETIME",
	"DATE", "DECIMAL", "DEC", "DEFAULT", "DOUBLE", "ENGINE", "ENUM", "EXISTS",
	"FALSE", "FIXED", "FLOAT", "HASH", "INDEX", "IF", "INTEGER", "INT", "JSON",
	"KEY_BLOCK_SIZE", "KEY", "LOCALTIMESTAMP", "LOCALTIME", "LONGBLOB", "LONGTEXT",
	"MEDIUMBLOB", "MEDIUMINT", "MEDIUMTEXT", "NATIONAL", "NCHAR", "NOT", "NOW",
	"NULL", "NUMERIC", "NVARCHAR", "ON", "PARSER", "PRECISION", "PRIMARY",
	"REAL", "SERIAL", "SET", "SIGNED", "SMALLINT", "TABLE", "TEMPORARY", "TEXT",
	"TIMESTAMP", "TIME", "TINYBLOB", "TINYINT", "TINYTEXT", "TRUE", "UNIQUE",
	"UNSIGNED", "UPDATE", "USING", "VALUE", "VARBINARY", "VARCHAR", "VARYING",
	"WITH", "YEAR", "ZEROFILL", "Integer", "Number", "FilesizeLiteral", "Literal",
	"PlaceholderString", "BackQuotedString", "RawString", "WS",
}

var ruleNames = []string{
	"root", "createTableDDL", "tableName", "ifNotExists", "createDefinitions",
	"createDefinition", "columnDefinition", "dataType", "lengthOneDimension",
	"lengthTwoDimension", "lengthTwoOptionalDimension", "collectionOptions",
	"columnConstraint", "defaultValue", "unaryOperator", "constant", "currentTimestamp",
	"tableConstraint", "indexOption", "indexType", "indexColumnNames", "tableOption",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CreateTableParser struct {
	*antlr.BaseParser
}

func NewCreateTableParser(input antlr.TokenStream) *CreateTableParser {
	this := new(CreateTableParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CreateTable.g4"

	return this
}

// CreateTableParser tokens.
const (
	CreateTableParserEOF               = antlr.TokenEOF
	CreateTableParserT__0              = 1
	CreateTableParserT__1              = 2
	CreateTableParserT__2              = 3
	CreateTableParserT__3              = 4
	CreateTableParserT__4              = 5
	CreateTableParserT__5              = 6
	CreateTableParserT__6              = 7
	CreateTableParserT__7              = 8
	CreateTableParserT__8              = 9
	CreateTableParserAUTO_INCREMENT    = 10
	CreateTableParserBIGINT            = 11
	CreateTableParserBINARY            = 12
	CreateTableParserBIT               = 13
	CreateTableParserBLOB              = 14
	CreateTableParserBTREE             = 15
	CreateTableParserBOOLEAN           = 16
	CreateTableParserBOOL              = 17
	CreateTableParserCHARACTER         = 18
	CreateTableParserCHARSET           = 19
	CreateTableParserCHAR              = 20
	CreateTableParserCOLLATE           = 21
	CreateTableParserCOMMENT           = 22
	CreateTableParserCONSTRAINT        = 23
	CreateTableParserCREATE            = 24
	CreateTableParserCURRENT_TIMESTAMP = 25
	CreateTableParserDATETIME          = 26
	CreateTableParserDATE              = 27
	CreateTableParserDECIMAL           = 28
	CreateTableParserDEC               = 29
	CreateTableParserDEFAULT           = 30
	CreateTableParserDOUBLE            = 31
	CreateTableParserENGINE            = 32
	CreateTableParserENUM              = 33
	CreateTableParserEXISTS            = 34
	CreateTableParserFALSE             = 35
	CreateTableParserFIXED             = 36
	CreateTableParserFLOAT             = 37
	CreateTableParserHASH              = 38
	CreateTableParserINDEX             = 39
	CreateTableParserIF                = 40
	CreateTableParserINTEGER           = 41
	CreateTableParserINT               = 42
	CreateTableParserJSON              = 43
	CreateTableParserKEY_BLOCK_SIZE    = 44
	CreateTableParserKEY               = 45
	CreateTableParserLOCALTIMESTAMP    = 46
	CreateTableParserLOCALTIME         = 47
	CreateTableParserLONGBLOB          = 48
	CreateTableParserLONGTEXT          = 49
	CreateTableParserMEDIUMBLOB        = 50
	CreateTableParserMEDIUMINT         = 51
	CreateTableParserMEDIUMTEXT        = 52
	CreateTableParserNATIONAL          = 53
	CreateTableParserNCHAR             = 54
	CreateTableParserNOT               = 55
	CreateTableParserNOW               = 56
	CreateTableParserNULL              = 57
	CreateTableParserNUMERIC           = 58
	CreateTableParserNVARCHAR          = 59
	CreateTableParserON                = 60
	CreateTableParserPARSER            = 61
	CreateTableParserPRECISION         = 62
	CreateTableParserPRIMARY           = 63
	CreateTableParserREAL              = 64
	CreateTableParserSERIAL            = 65
	CreateTableParserSET               = 66
	CreateTableParserSIGNED            = 67
	CreateTableParserSMALLINT          = 68
	CreateTableParserTABLE             = 69
	CreateTableParserTEMPORARY         = 70
	CreateTableParserTEXT              = 71
	CreateTableParserTIMESTAMP         = 72
	CreateTableParserTIME              = 73
	CreateTableParserTINYBLOB          = 74
	CreateTableParserTINYINT           = 75
	CreateTableParserTINYTEXT          = 76
	CreateTableParserTRUE              = 77
	CreateTableParserUNIQUE            = 78
	CreateTableParserUNSIGNED          = 79
	CreateTableParserUPDATE            = 80
	CreateTableParserUSING             = 81
	CreateTableParserVALUE             = 82
	CreateTableParserVARBINARY         = 83
	CreateTableParserVARCHAR           = 84
	CreateTableParserVARYING           = 85
	CreateTableParserWITH              = 86
	CreateTableParserYEAR              = 87
	CreateTableParserZEROFILL          = 88
	CreateTableParserInteger           = 89
	CreateTableParserNumber            = 90
	CreateTableParserFilesizeLiteral   = 91
	CreateTableParserLiteral           = 92
	CreateTableParserPlaceholderString = 93
	CreateTableParserBackQuotedString  = 94
	CreateTableParserRawString         = 95
	CreateTableParserWS                = 96
)

// CreateTableParser rules.
const (
	CreateTableParserRULE_root                       = 0
	CreateTableParserRULE_createTableDDL             = 1
	CreateTableParserRULE_tableName                  = 2
	CreateTableParserRULE_ifNotExists                = 3
	CreateTableParserRULE_createDefinitions          = 4
	CreateTableParserRULE_createDefinition           = 5
	CreateTableParserRULE_columnDefinition           = 6
	CreateTableParserRULE_dataType                   = 7
	CreateTableParserRULE_lengthOneDimension         = 8
	CreateTableParserRULE_lengthTwoDimension         = 9
	CreateTableParserRULE_lengthTwoOptionalDimension = 10
	CreateTableParserRULE_collectionOptions          = 11
	CreateTableParserRULE_columnConstraint           = 12
	CreateTableParserRULE_defaultValue               = 13
	CreateTableParserRULE_unaryOperator              = 14
	CreateTableParserRULE_constant                   = 15
	CreateTableParserRULE_currentTimestamp           = 16
	CreateTableParserRULE_tableConstraint            = 17
	CreateTableParserRULE_indexOption                = 18
	CreateTableParserRULE_indexType                  = 19
	CreateTableParserRULE_indexColumnNames           = 20
	CreateTableParserRULE_tableOption                = 21
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) CreateTableDDL() ICreateTableDDLContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateTableDDLContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreateTableDDLContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(CreateTableParserEOF, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *CreateTableParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CreateTableParserRULE_root)
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
	{
		p.SetState(44)
		p.CreateTableDDL()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CreateTableParserT__0 {
		{
			p.SetState(45)
			p.Match(CreateTableParserT__0)
		}

	}
	{
		p.SetState(48)
		p.Match(CreateTableParserEOF)
	}

	return localctx
}

// ICreateTableDDLContext is an interface to support dynamic dispatch.
type ICreateTableDDLContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateTableDDLContext differentiates from other interfaces.
	IsCreateTableDDLContext()
}

type CreateTableDDLContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateTableDDLContext() *CreateTableDDLContext {
	var p = new(CreateTableDDLContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_createTableDDL
	return p
}

func (*CreateTableDDLContext) IsCreateTableDDLContext() {}

func NewCreateTableDDLContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTableDDLContext {
	var p = new(CreateTableDDLContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_createTableDDL

	return p
}

func (s *CreateTableDDLContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTableDDLContext) CREATE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCREATE, 0)
}

func (s *CreateTableDDLContext) TABLE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserTABLE, 0)
}

func (s *CreateTableDDLContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *CreateTableDDLContext) CreateDefinitions() ICreateDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreateDefinitionsContext)
}

func (s *CreateTableDDLContext) TEMPORARY() antlr.TerminalNode {
	return s.GetToken(CreateTableParserTEMPORARY, 0)
}

func (s *CreateTableDDLContext) IfNotExists() IIfNotExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfNotExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateTableDDLContext) AllTableOption() []ITableOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableOptionContext)(nil)).Elem())
	var tst = make([]ITableOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableOptionContext)
		}
	}

	return tst
}

func (s *CreateTableDDLContext) TableOption(i int) ITableOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableOptionContext)
}

func (s *CreateTableDDLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTableDDLContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateTableDDLContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterCreateTableDDL(s)
	}
}

func (s *CreateTableDDLContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitCreateTableDDL(s)
	}
}

func (p *CreateTableParser) CreateTableDDL() (localctx ICreateTableDDLContext) {
	localctx = NewCreateTableDDLContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CreateTableParserRULE_createTableDDL)
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
	{
		p.SetState(50)
		p.Match(CreateTableParserCREATE)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CreateTableParserTEMPORARY {
		{
			p.SetState(51)
			p.Match(CreateTableParserTEMPORARY)
		}

	}
	{
		p.SetState(54)
		p.Match(CreateTableParserTABLE)
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CreateTableParserIF {
		{
			p.SetState(55)
			p.IfNotExists()
		}

	}
	{
		p.SetState(58)
		p.TableName()
	}
	{
		p.SetState(59)
		p.CreateDefinitions()
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(CreateTableParserAUTO_INCREMENT-10))|(1<<(CreateTableParserCHARACTER-10))|(1<<(CreateTableParserCHARSET-10))|(1<<(CreateTableParserCOLLATE-10))|(1<<(CreateTableParserCOMMENT-10))|(1<<(CreateTableParserDEFAULT-10))|(1<<(CreateTableParserENGINE-10)))) != 0 {
		{
			p.SetState(60)
			p.TableOption()
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(CreateTableParserT__1-2))|(1<<(CreateTableParserAUTO_INCREMENT-2))|(1<<(CreateTableParserCHARACTER-2))|(1<<(CreateTableParserCHARSET-2))|(1<<(CreateTableParserCOLLATE-2))|(1<<(CreateTableParserCOMMENT-2))|(1<<(CreateTableParserDEFAULT-2))|(1<<(CreateTableParserENGINE-2)))) != 0 {
			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CreateTableParserT__1 {
				{
					p.SetState(61)
					p.Match(CreateTableParserT__1)
				}

			}
			{
				p.SetState(64)
				p.TableOption()
			}

			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(CreateTableParserBackQuotedString, 0)
}

func (s *TableNameContext) RawString() antlr.TerminalNode {
	return s.GetToken(CreateTableParserRawString, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (p *CreateTableParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CreateTableParserRULE_tableName)
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
	{
		p.SetState(72)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CreateTableParserBackQuotedString || _la == CreateTableParserRawString) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIfNotExistsContext is an interface to support dynamic dispatch.
type IIfNotExistsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfNotExistsContext differentiates from other interfaces.
	IsIfNotExistsContext()
}

type IfNotExistsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfNotExistsContext() *IfNotExistsContext {
	var p = new(IfNotExistsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_ifNotExists
	return p
}

func (*IfNotExistsContext) IsIfNotExistsContext() {}

func NewIfNotExistsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfNotExistsContext {
	var p = new(IfNotExistsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_ifNotExists

	return p
}

func (s *IfNotExistsContext) GetParser() antlr.Parser { return s.parser }

func (s *IfNotExistsContext) IF() antlr.TerminalNode {
	return s.GetToken(CreateTableParserIF, 0)
}

func (s *IfNotExistsContext) NOT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserNOT, 0)
}

func (s *IfNotExistsContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(CreateTableParserEXISTS, 0)
}

func (s *IfNotExistsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfNotExistsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfNotExistsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterIfNotExists(s)
	}
}

func (s *IfNotExistsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitIfNotExists(s)
	}
}

func (p *CreateTableParser) IfNotExists() (localctx IIfNotExistsContext) {
	localctx = NewIfNotExistsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CreateTableParserRULE_ifNotExists)

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
		p.Match(CreateTableParserIF)
	}
	{
		p.SetState(75)
		p.Match(CreateTableParserNOT)
	}
	{
		p.SetState(76)
		p.Match(CreateTableParserEXISTS)
	}

	return localctx
}

// ICreateDefinitionsContext is an interface to support dynamic dispatch.
type ICreateDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateDefinitionsContext differentiates from other interfaces.
	IsCreateDefinitionsContext()
}

type CreateDefinitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateDefinitionsContext() *CreateDefinitionsContext {
	var p = new(CreateDefinitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_createDefinitions
	return p
}

func (*CreateDefinitionsContext) IsCreateDefinitionsContext() {}

func NewCreateDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateDefinitionsContext {
	var p = new(CreateDefinitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_createDefinitions

	return p
}

func (s *CreateDefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateDefinitionsContext) AllCreateDefinition() []ICreateDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICreateDefinitionContext)(nil)).Elem())
	var tst = make([]ICreateDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICreateDefinitionContext)
		}
	}

	return tst
}

func (s *CreateDefinitionsContext) CreateDefinition(i int) ICreateDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICreateDefinitionContext)
}

func (s *CreateDefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateDefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateDefinitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterCreateDefinitions(s)
	}
}

func (s *CreateDefinitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitCreateDefinitions(s)
	}
}

func (p *CreateTableParser) CreateDefinitions() (localctx ICreateDefinitionsContext) {
	localctx = NewCreateDefinitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CreateTableParserRULE_createDefinitions)
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
	{
		p.SetState(78)
		p.Match(CreateTableParserT__2)
	}
	{
		p.SetState(79)
		p.CreateDefinition()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CreateTableParserT__1 {
		{
			p.SetState(80)
			p.Match(CreateTableParserT__1)
		}
		{
			p.SetState(81)
			p.CreateDefinition()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(CreateTableParserT__3)
	}

	return localctx
}

// ICreateDefinitionContext is an interface to support dynamic dispatch.
type ICreateDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetField returns the field token.
	GetField() antlr.Token

	// SetField sets the field token.
	SetField(antlr.Token)

	// IsCreateDefinitionContext differentiates from other interfaces.
	IsCreateDefinitionContext()
}

type CreateDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	field  antlr.Token
}

func NewEmptyCreateDefinitionContext() *CreateDefinitionContext {
	var p = new(CreateDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_createDefinition
	return p
}

func (*CreateDefinitionContext) IsCreateDefinitionContext() {}

func NewCreateDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateDefinitionContext {
	var p = new(CreateDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_createDefinition

	return p
}

func (s *CreateDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateDefinitionContext) GetField() antlr.Token { return s.field }

func (s *CreateDefinitionContext) SetField(v antlr.Token) { s.field = v }

func (s *CreateDefinitionContext) ColumnDefinition() IColumnDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionContext)
}

func (s *CreateDefinitionContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(CreateTableParserBackQuotedString, 0)
}

func (s *CreateDefinitionContext) RawString() antlr.TerminalNode {
	return s.GetToken(CreateTableParserRawString, 0)
}

func (s *CreateDefinitionContext) TableConstraint() ITableConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableConstraintContext)
}

func (s *CreateDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterCreateDefinition(s)
	}
}

func (s *CreateDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitCreateDefinition(s)
	}
}

func (p *CreateTableParser) CreateDefinition() (localctx ICreateDefinitionContext) {
	localctx = NewCreateDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CreateTableParserRULE_createDefinition)
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

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CreateDefinitionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == CreateTableParserBackQuotedString || _la == CreateTableParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CreateDefinitionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(90)
			p.ColumnDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.TableConstraint()
		}

	}

	return localctx
}

// IColumnDefinitionContext is an interface to support dynamic dispatch.
type IColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnDefinitionContext differentiates from other interfaces.
	IsColumnDefinitionContext()
}

type ColumnDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnDefinitionContext() *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_columnDefinition
	return p
}

func (*ColumnDefinitionContext) IsColumnDefinitionContext() {}

func NewColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_columnDefinition

	return p
}

func (s *ColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnDefinitionContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *ColumnDefinitionContext) AllColumnConstraint() []IColumnConstraintContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumnConstraintContext)(nil)).Elem())
	var tst = make([]IColumnConstraintContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumnConstraintContext)
		}
	}

	return tst
}

func (s *ColumnDefinitionContext) ColumnConstraint(i int) IColumnConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnConstraintContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumnConstraintContext)
}

func (s *ColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterColumnDefinition(s)
	}
}

func (s *ColumnDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitColumnDefinition(s)
	}
}

func (p *CreateTableParser) ColumnDefinition() (localctx IColumnDefinitionContext) {
	localctx = NewColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CreateTableParserRULE_columnDefinition)
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
	{
		p.SetState(94)
		p.DataType()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CreateTableParserAUTO_INCREMENT)|(1<<CreateTableParserCOLLATE)|(1<<CreateTableParserCOMMENT)|(1<<CreateTableParserDEFAULT))) != 0) || (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(CreateTableParserKEY-45))|(1<<(CreateTableParserNOT-45))|(1<<(CreateTableParserNULL-45))|(1<<(CreateTableParserPRIMARY-45))|(1<<(CreateTableParserSERIAL-45)))) != 0) || _la == CreateTableParserUNIQUE {
		{
			p.SetState(95)
			p.ColumnConstraint()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDataTypeContext is an interface to support dynamic dispatch.
type IDataTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTypeName returns the typeName token.
	GetTypeName() antlr.Token

	// SetTypeName sets the typeName token.
	SetTypeName(antlr.Token)

	// IsDataTypeContext differentiates from other interfaces.
	IsDataTypeContext()
}

type DataTypeContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	typeName antlr.Token
}

func NewEmptyDataTypeContext() *DataTypeContext {
	var p = new(DataTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_dataType
	return p
}

func (*DataTypeContext) IsDataTypeContext() {}

func NewDataTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeContext {
	var p = new(DataTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_dataType

	return p
}

func (s *DataTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypeContext) GetTypeName() antlr.Token { return s.typeName }

func (s *DataTypeContext) SetTypeName(v antlr.Token) { s.typeName = v }

func (s *DataTypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCHAR, 0)
}

func (s *DataTypeContext) VARCHAR() antlr.TerminalNode {
	return s.GetToken(CreateTableParserVARCHAR, 0)
}

func (s *DataTypeContext) TINYTEXT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserTINYTEXT, 0)
}

func (s *DataTypeContext) TEXT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserTEXT, 0)
}

func (s *DataTypeContext) MEDIUMTEXT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserMEDIUMTEXT, 0)
}

func (s *DataTypeContext) LONGTEXT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserLONGTEXT, 0)
}

func (s *DataTypeContext) NCHAR() antlr.TerminalNode {
	return s.GetToken(CreateTableParserNCHAR, 0)
}

func (s *DataTypeContext) NVARCHAR() antlr.TerminalNode {
	return s.GetToken(CreateTableParserNVARCHAR, 0)
}

func (s *DataTypeContext) LengthOneDimension() ILengthOneDimensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILengthOneDimensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILengthOneDimensionContext)
}

func (s *DataTypeContext) AllBINARY() []antlr.TerminalNode {
	return s.GetTokens(CreateTableParserBINARY)
}

func (s *DataTypeContext) BINARY(i int) antlr.TerminalNode {
	return s.GetToken(CreateTableParserBINARY, i)
}

func (s *DataTypeContext) COLLATE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCOLLATE, 0)
}

func (s *DataTypeContext) Literal() antlr.TerminalNode {
	return s.GetToken(CreateTableParserLiteral, 0)
}

func (s *DataTypeContext) RawString() antlr.TerminalNode {
	return s.GetToken(CreateTableParserRawString, 0)
}

func (s *DataTypeContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(CreateTableParserBackQuotedString, 0)
}

func (s *DataTypeContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCHARACTER, 0)
}

func (s *DataTypeContext) AllSET() []antlr.TerminalNode {
	return s.GetTokens(CreateTableParserSET)
}

func (s *DataTypeContext) SET(i int) antlr.TerminalNode {
	return s.GetToken(CreateTableParserSET, i)
}

func (s *DataTypeContext) CHARSET() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCHARSET, 0)
}

func (s *DataTypeContext) NATIONAL() antlr.TerminalNode {
	return s.GetToken(CreateTableParserNATIONAL, 0)
}

func (s *DataTypeContext) VARYING() antlr.TerminalNode {
	return s.GetToken(CreateTableParserVARYING, 0)
}

func (s *DataTypeContext) TINYINT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserTINYINT, 0)
}

func (s *DataTypeContext) SMALLINT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserSMALLINT, 0)
}

func (s *DataTypeContext) MEDIUMINT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserMEDIUMINT, 0)
}

func (s *DataTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserINT, 0)
}

func (s *DataTypeContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(CreateTableParserINTEGER, 0)
}

func (s *DataTypeContext) BIGINT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserBIGINT, 0)
}

func (s *DataTypeContext) ZEROFILL() antlr.TerminalNode {
	return s.GetToken(CreateTableParserZEROFILL, 0)
}

func (s *DataTypeContext) SIGNED() antlr.TerminalNode {
	return s.GetToken(CreateTableParserSIGNED, 0)
}

func (s *DataTypeContext) UNSIGNED() antlr.TerminalNode {
	return s.GetToken(CreateTableParserUNSIGNED, 0)
}

func (s *DataTypeContext) REAL() antlr.TerminalNode {
	return s.GetToken(CreateTableParserREAL, 0)
}

func (s *DataTypeContext) LengthTwoDimension() ILengthTwoDimensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILengthTwoDimensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILengthTwoDimensionContext)
}

func (s *DataTypeContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserDOUBLE, 0)
}

func (s *DataTypeContext) PRECISION() antlr.TerminalNode {
	return s.GetToken(CreateTableParserPRECISION, 0)
}

func (s *DataTypeContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(CreateTableParserDECIMAL, 0)
}

func (s *DataTypeContext) DEC() antlr.TerminalNode {
	return s.GetToken(CreateTableParserDEC, 0)
}

func (s *DataTypeContext) FIXED() antlr.TerminalNode {
	return s.GetToken(CreateTableParserFIXED, 0)
}

func (s *DataTypeContext) NUMERIC() antlr.TerminalNode {
	return s.GetToken(CreateTableParserNUMERIC, 0)
}

func (s *DataTypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserFLOAT, 0)
}

func (s *DataTypeContext) LengthTwoOptionalDimension() ILengthTwoOptionalDimensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILengthTwoOptionalDimensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILengthTwoOptionalDimensionContext)
}

func (s *DataTypeContext) DATE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserDATE, 0)
}

func (s *DataTypeContext) TINYBLOB() antlr.TerminalNode {
	return s.GetToken(CreateTableParserTINYBLOB, 0)
}

func (s *DataTypeContext) BLOB() antlr.TerminalNode {
	return s.GetToken(CreateTableParserBLOB, 0)
}

func (s *DataTypeContext) MEDIUMBLOB() antlr.TerminalNode {
	return s.GetToken(CreateTableParserMEDIUMBLOB, 0)
}

func (s *DataTypeContext) LONGBLOB() antlr.TerminalNode {
	return s.GetToken(CreateTableParserLONGBLOB, 0)
}

func (s *DataTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(CreateTableParserBOOL, 0)
}

func (s *DataTypeContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(CreateTableParserBOOLEAN, 0)
}

func (s *DataTypeContext) SERIAL() antlr.TerminalNode {
	return s.GetToken(CreateTableParserSERIAL, 0)
}

func (s *DataTypeContext) JSON() antlr.TerminalNode {
	return s.GetToken(CreateTableParserJSON, 0)
}

func (s *DataTypeContext) BIT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserBIT, 0)
}

func (s *DataTypeContext) TIME() antlr.TerminalNode {
	return s.GetToken(CreateTableParserTIME, 0)
}

func (s *DataTypeContext) TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(CreateTableParserTIMESTAMP, 0)
}

func (s *DataTypeContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(CreateTableParserDATETIME, 0)
}

func (s *DataTypeContext) VARBINARY() antlr.TerminalNode {
	return s.GetToken(CreateTableParserVARBINARY, 0)
}

func (s *DataTypeContext) YEAR() antlr.TerminalNode {
	return s.GetToken(CreateTableParserYEAR, 0)
}

func (s *DataTypeContext) CollectionOptions() ICollectionOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollectionOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollectionOptionsContext)
}

func (s *DataTypeContext) ENUM() antlr.TerminalNode {
	return s.GetToken(CreateTableParserENUM, 0)
}

func (s *DataTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterDataType(s)
	}
}

func (s *DataTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitDataType(s)
	}
}

func (p *CreateTableParser) DataType() (localctx IDataTypeContext) {
	localctx = NewDataTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CreateTableParserRULE_dataType)
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

	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == CreateTableParserCHAR || (((_la-49)&-(0x1f+1)) == 0 && ((1<<uint((_la-49)))&((1<<(CreateTableParserLONGTEXT-49))|(1<<(CreateTableParserMEDIUMTEXT-49))|(1<<(CreateTableParserNCHAR-49))|(1<<(CreateTableParserNVARCHAR-49))|(1<<(CreateTableParserTEXT-49))|(1<<(CreateTableParserTINYTEXT-49)))) != 0) || _la == CreateTableParserVARCHAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__2 {
			{
				p.SetState(102)
				p.LengthOneDimension()
			}

		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserBINARY {
			{
				p.SetState(105)
				p.Match(CreateTableParserBINARY)
			}

		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserCHARACTER || _la == CreateTableParserCHARSET {
			p.SetState(111)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case CreateTableParserCHARACTER:
				{
					p.SetState(108)
					p.Match(CreateTableParserCHARACTER)
				}
				{
					p.SetState(109)
					p.Match(CreateTableParserSET)
				}

			case CreateTableParserCHARSET:
				{
					p.SetState(110)
					p.Match(CreateTableParserCHARSET)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			{
				p.SetState(113)
				_la = p.GetTokenStream().LA(1)

				if !(_la == CreateTableParserBINARY || _la == CreateTableParserBackQuotedString || _la == CreateTableParserRawString) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(116)
				p.Match(CreateTableParserCOLLATE)
			}
			{
				p.SetState(117)
				p.Match(CreateTableParserLiteral)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Match(CreateTableParserNATIONAL)
		}
		{
			p.SetState(121)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == CreateTableParserCHARACTER || _la == CreateTableParserVARCHAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__2 {
			{
				p.SetState(122)
				p.LengthOneDimension()
			}

		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserBINARY {
			{
				p.SetState(125)
				p.Match(CreateTableParserBINARY)
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
			p.Match(CreateTableParserNCHAR)
		}
		{
			p.SetState(129)

			var _m = p.Match(CreateTableParserVARCHAR)

			localctx.(*DataTypeContext).typeName = _m
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__2 {
			{
				p.SetState(130)
				p.LengthOneDimension()
			}

		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserBINARY {
			{
				p.SetState(133)
				p.Match(CreateTableParserBINARY)
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			p.Match(CreateTableParserNATIONAL)
		}
		{
			p.SetState(137)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == CreateTableParserCHARACTER || _la == CreateTableParserCHAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(138)
			p.Match(CreateTableParserVARYING)
		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__2 {
			{
				p.SetState(139)
				p.LengthOneDimension()
			}

		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserBINARY {
			{
				p.SetState(142)
				p.Match(CreateTableParserBINARY)
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(145)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !((((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(CreateTableParserBIGINT-11))|(1<<(CreateTableParserINTEGER-11))|(1<<(CreateTableParserINT-11)))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(CreateTableParserMEDIUMINT-51))|(1<<(CreateTableParserSMALLINT-51))|(1<<(CreateTableParserTINYINT-51)))) != 0)) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__2 {
			{
				p.SetState(146)
				p.LengthOneDimension()
			}

		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserSIGNED || _la == CreateTableParserUNSIGNED {
			{
				p.SetState(149)
				_la = p.GetTokenStream().LA(1)

				if !(_la == CreateTableParserSIGNED || _la == CreateTableParserUNSIGNED) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserZEROFILL {
			{
				p.SetState(152)
				p.Match(CreateTableParserZEROFILL)
			}

		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(155)

			var _m = p.Match(CreateTableParserREAL)

			localctx.(*DataTypeContext).typeName = _m
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__2 {
			{
				p.SetState(156)
				p.LengthTwoDimension()
			}

		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserSIGNED || _la == CreateTableParserUNSIGNED {
			{
				p.SetState(159)
				_la = p.GetTokenStream().LA(1)

				if !(_la == CreateTableParserSIGNED || _la == CreateTableParserUNSIGNED) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserZEROFILL {
			{
				p.SetState(162)
				p.Match(CreateTableParserZEROFILL)
			}

		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(165)

			var _m = p.Match(CreateTableParserDOUBLE)

			localctx.(*DataTypeContext).typeName = _m
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserPRECISION {
			{
				p.SetState(166)
				p.Match(CreateTableParserPRECISION)
			}

		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__2 {
			{
				p.SetState(169)
				p.LengthTwoDimension()
			}

		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserSIGNED || _la == CreateTableParserUNSIGNED {
			{
				p.SetState(172)
				_la = p.GetTokenStream().LA(1)

				if !(_la == CreateTableParserSIGNED || _la == CreateTableParserUNSIGNED) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserZEROFILL {
			{
				p.SetState(175)
				p.Match(CreateTableParserZEROFILL)
			}

		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(178)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-28)&-(0x1f+1)) == 0 && ((1<<uint((_la-28)))&((1<<(CreateTableParserDECIMAL-28))|(1<<(CreateTableParserDEC-28))|(1<<(CreateTableParserFIXED-28))|(1<<(CreateTableParserFLOAT-28))|(1<<(CreateTableParserNUMERIC-28)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__2 {
			{
				p.SetState(179)
				p.LengthTwoOptionalDimension()
			}

		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserSIGNED || _la == CreateTableParserUNSIGNED {
			{
				p.SetState(182)
				_la = p.GetTokenStream().LA(1)

				if !(_la == CreateTableParserSIGNED || _la == CreateTableParserUNSIGNED) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserZEROFILL {
			{
				p.SetState(185)
				p.Match(CreateTableParserZEROFILL)
			}

		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(188)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CreateTableParserBLOB)|(1<<CreateTableParserBOOLEAN)|(1<<CreateTableParserBOOL)|(1<<CreateTableParserDATE))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(CreateTableParserJSON-43))|(1<<(CreateTableParserLONGBLOB-43))|(1<<(CreateTableParserMEDIUMBLOB-43))|(1<<(CreateTableParserSERIAL-43))|(1<<(CreateTableParserTINYBLOB-43)))) != 0)) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(189)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CreateTableParserBINARY)|(1<<CreateTableParserBIT)|(1<<CreateTableParserDATETIME))) != 0) || (((_la-72)&-(0x1f+1)) == 0 && ((1<<uint((_la-72)))&((1<<(CreateTableParserTIMESTAMP-72))|(1<<(CreateTableParserTIME-72))|(1<<(CreateTableParserVARBINARY-72))|(1<<(CreateTableParserYEAR-72)))) != 0)) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__2 {
			{
				p.SetState(190)
				p.LengthOneDimension()
			}

		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(193)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == CreateTableParserENUM || _la == CreateTableParserSET) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(194)
			p.CollectionOptions()
		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserBINARY {
			{
				p.SetState(195)
				p.Match(CreateTableParserBINARY)
			}

		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserCHARACTER || _la == CreateTableParserCHARSET {
			p.SetState(201)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case CreateTableParserCHARACTER:
				{
					p.SetState(198)
					p.Match(CreateTableParserCHARACTER)
				}
				{
					p.SetState(199)
					p.Match(CreateTableParserSET)
				}

			case CreateTableParserCHARSET:
				{
					p.SetState(200)
					p.Match(CreateTableParserCHARSET)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			{
				p.SetState(203)
				_la = p.GetTokenStream().LA(1)

				if !(_la == CreateTableParserBINARY || _la == CreateTableParserBackQuotedString || _la == CreateTableParserRawString) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	}

	return localctx
}

// ILengthOneDimensionContext is an interface to support dynamic dispatch.
type ILengthOneDimensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLengthOneDimensionContext differentiates from other interfaces.
	IsLengthOneDimensionContext()
}

type LengthOneDimensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthOneDimensionContext() *LengthOneDimensionContext {
	var p = new(LengthOneDimensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_lengthOneDimension
	return p
}

func (*LengthOneDimensionContext) IsLengthOneDimensionContext() {}

func NewLengthOneDimensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthOneDimensionContext {
	var p = new(LengthOneDimensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_lengthOneDimension

	return p
}

func (s *LengthOneDimensionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthOneDimensionContext) Integer() antlr.TerminalNode {
	return s.GetToken(CreateTableParserInteger, 0)
}

func (s *LengthOneDimensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthOneDimensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthOneDimensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterLengthOneDimension(s)
	}
}

func (s *LengthOneDimensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitLengthOneDimension(s)
	}
}

func (p *CreateTableParser) LengthOneDimension() (localctx ILengthOneDimensionContext) {
	localctx = NewLengthOneDimensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CreateTableParserRULE_lengthOneDimension)

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
		p.SetState(208)
		p.Match(CreateTableParserT__2)
	}
	{
		p.SetState(209)
		p.Match(CreateTableParserInteger)
	}
	{
		p.SetState(210)
		p.Match(CreateTableParserT__3)
	}

	return localctx
}

// ILengthTwoDimensionContext is an interface to support dynamic dispatch.
type ILengthTwoDimensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLengthTwoDimensionContext differentiates from other interfaces.
	IsLengthTwoDimensionContext()
}

type LengthTwoDimensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthTwoDimensionContext() *LengthTwoDimensionContext {
	var p = new(LengthTwoDimensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_lengthTwoDimension
	return p
}

func (*LengthTwoDimensionContext) IsLengthTwoDimensionContext() {}

func NewLengthTwoDimensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthTwoDimensionContext {
	var p = new(LengthTwoDimensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_lengthTwoDimension

	return p
}

func (s *LengthTwoDimensionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthTwoDimensionContext) AllInteger() []antlr.TerminalNode {
	return s.GetTokens(CreateTableParserInteger)
}

func (s *LengthTwoDimensionContext) Integer(i int) antlr.TerminalNode {
	return s.GetToken(CreateTableParserInteger, i)
}

func (s *LengthTwoDimensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthTwoDimensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthTwoDimensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterLengthTwoDimension(s)
	}
}

func (s *LengthTwoDimensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitLengthTwoDimension(s)
	}
}

func (p *CreateTableParser) LengthTwoDimension() (localctx ILengthTwoDimensionContext) {
	localctx = NewLengthTwoDimensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CreateTableParserRULE_lengthTwoDimension)

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
		p.Match(CreateTableParserT__2)
	}
	{
		p.SetState(213)
		p.Match(CreateTableParserInteger)
	}
	{
		p.SetState(214)
		p.Match(CreateTableParserT__1)
	}
	{
		p.SetState(215)
		p.Match(CreateTableParserInteger)
	}
	{
		p.SetState(216)
		p.Match(CreateTableParserT__3)
	}

	return localctx
}

// ILengthTwoOptionalDimensionContext is an interface to support dynamic dispatch.
type ILengthTwoOptionalDimensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLengthTwoOptionalDimensionContext differentiates from other interfaces.
	IsLengthTwoOptionalDimensionContext()
}

type LengthTwoOptionalDimensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthTwoOptionalDimensionContext() *LengthTwoOptionalDimensionContext {
	var p = new(LengthTwoOptionalDimensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_lengthTwoOptionalDimension
	return p
}

func (*LengthTwoOptionalDimensionContext) IsLengthTwoOptionalDimensionContext() {}

func NewLengthTwoOptionalDimensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthTwoOptionalDimensionContext {
	var p = new(LengthTwoOptionalDimensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_lengthTwoOptionalDimension

	return p
}

func (s *LengthTwoOptionalDimensionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthTwoOptionalDimensionContext) AllInteger() []antlr.TerminalNode {
	return s.GetTokens(CreateTableParserInteger)
}

func (s *LengthTwoOptionalDimensionContext) Integer(i int) antlr.TerminalNode {
	return s.GetToken(CreateTableParserInteger, i)
}

func (s *LengthTwoOptionalDimensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthTwoOptionalDimensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthTwoOptionalDimensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterLengthTwoOptionalDimension(s)
	}
}

func (s *LengthTwoOptionalDimensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitLengthTwoOptionalDimension(s)
	}
}

func (p *CreateTableParser) LengthTwoOptionalDimension() (localctx ILengthTwoOptionalDimensionContext) {
	localctx = NewLengthTwoOptionalDimensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CreateTableParserRULE_lengthTwoOptionalDimension)
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
	{
		p.SetState(218)
		p.Match(CreateTableParserT__2)
	}
	{
		p.SetState(219)
		p.Match(CreateTableParserInteger)
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CreateTableParserT__1 {
		{
			p.SetState(220)
			p.Match(CreateTableParserT__1)
		}
		{
			p.SetState(221)
			p.Match(CreateTableParserInteger)
		}

	}
	{
		p.SetState(224)
		p.Match(CreateTableParserT__3)
	}

	return localctx
}

// ICollectionOptionsContext is an interface to support dynamic dispatch.
type ICollectionOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCollectionOptionsContext differentiates from other interfaces.
	IsCollectionOptionsContext()
}

type CollectionOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionOptionsContext() *CollectionOptionsContext {
	var p = new(CollectionOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_collectionOptions
	return p
}

func (*CollectionOptionsContext) IsCollectionOptionsContext() {}

func NewCollectionOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionOptionsContext {
	var p = new(CollectionOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_collectionOptions

	return p
}

func (s *CollectionOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionOptionsContext) AllLiteral() []antlr.TerminalNode {
	return s.GetTokens(CreateTableParserLiteral)
}

func (s *CollectionOptionsContext) Literal(i int) antlr.TerminalNode {
	return s.GetToken(CreateTableParserLiteral, i)
}

func (s *CollectionOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterCollectionOptions(s)
	}
}

func (s *CollectionOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitCollectionOptions(s)
	}
}

func (p *CreateTableParser) CollectionOptions() (localctx ICollectionOptionsContext) {
	localctx = NewCollectionOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CreateTableParserRULE_collectionOptions)
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
	{
		p.SetState(226)
		p.Match(CreateTableParserT__2)
	}
	{
		p.SetState(227)
		p.Match(CreateTableParserLiteral)
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CreateTableParserT__1 {
		{
			p.SetState(228)
			p.Match(CreateTableParserT__1)
		}
		{
			p.SetState(229)
			p.Match(CreateTableParserLiteral)
		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(235)
		p.Match(CreateTableParserT__3)
	}

	return localctx
}

// IColumnConstraintContext is an interface to support dynamic dispatch.
type IColumnConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnConstraintContext differentiates from other interfaces.
	IsColumnConstraintContext()
}

type ColumnConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnConstraintContext() *ColumnConstraintContext {
	var p = new(ColumnConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_columnConstraint
	return p
}

func (*ColumnConstraintContext) IsColumnConstraintContext() {}

func NewColumnConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnConstraintContext {
	var p = new(ColumnConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_columnConstraint

	return p
}

func (s *ColumnConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnConstraintContext) NULL() antlr.TerminalNode {
	return s.GetToken(CreateTableParserNULL, 0)
}

func (s *ColumnConstraintContext) NOT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserNOT, 0)
}

func (s *ColumnConstraintContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserDEFAULT, 0)
}

func (s *ColumnConstraintContext) DefaultValue() IDefaultValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *ColumnConstraintContext) AUTO_INCREMENT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserAUTO_INCREMENT, 0)
}

func (s *ColumnConstraintContext) KEY() antlr.TerminalNode {
	return s.GetToken(CreateTableParserKEY, 0)
}

func (s *ColumnConstraintContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(CreateTableParserPRIMARY, 0)
}

func (s *ColumnConstraintContext) UNIQUE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserUNIQUE, 0)
}

func (s *ColumnConstraintContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCOMMENT, 0)
}

func (s *ColumnConstraintContext) Literal() antlr.TerminalNode {
	return s.GetToken(CreateTableParserLiteral, 0)
}

func (s *ColumnConstraintContext) COLLATE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCOLLATE, 0)
}

func (s *ColumnConstraintContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(CreateTableParserBackQuotedString, 0)
}

func (s *ColumnConstraintContext) SERIAL() antlr.TerminalNode {
	return s.GetToken(CreateTableParserSERIAL, 0)
}

func (s *ColumnConstraintContext) VALUE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserVALUE, 0)
}

func (s *ColumnConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterColumnConstraint(s)
	}
}

func (s *ColumnConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitColumnConstraint(s)
	}
}

func (p *CreateTableParser) ColumnConstraint() (localctx IColumnConstraintContext) {
	localctx = NewColumnConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CreateTableParserRULE_columnConstraint)
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

	p.SetState(259)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CreateTableParserNOT, CreateTableParserNULL:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserNOT {
			{
				p.SetState(237)
				p.Match(CreateTableParserNOT)
			}

		}
		{
			p.SetState(240)
			p.Match(CreateTableParserNULL)
		}

	case CreateTableParserDEFAULT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)
			p.Match(CreateTableParserDEFAULT)
		}
		{
			p.SetState(242)
			p.DefaultValue()
		}

	case CreateTableParserAUTO_INCREMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(243)
			p.Match(CreateTableParserAUTO_INCREMENT)
		}

	case CreateTableParserKEY, CreateTableParserPRIMARY:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserPRIMARY {
			{
				p.SetState(244)
				p.Match(CreateTableParserPRIMARY)
			}

		}
		{
			p.SetState(247)
			p.Match(CreateTableParserKEY)
		}

	case CreateTableParserUNIQUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(248)
			p.Match(CreateTableParserUNIQUE)
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(249)
				p.Match(CreateTableParserKEY)
			}

		}

	case CreateTableParserCOMMENT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(252)
			p.Match(CreateTableParserCOMMENT)
		}
		{
			p.SetState(253)
			p.Match(CreateTableParserLiteral)
		}

	case CreateTableParserCOLLATE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(254)
			p.Match(CreateTableParserCOLLATE)
		}
		{
			p.SetState(255)
			p.Match(CreateTableParserBackQuotedString)
		}

	case CreateTableParserSERIAL:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(256)
			p.Match(CreateTableParserSERIAL)
		}
		{
			p.SetState(257)
			p.Match(CreateTableParserDEFAULT)
		}
		{
			p.SetState(258)
			p.Match(CreateTableParserVALUE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefaultValueContext is an interface to support dynamic dispatch.
type IDefaultValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultValueContext differentiates from other interfaces.
	IsDefaultValueContext()
}

type DefaultValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultValueContext() *DefaultValueContext {
	var p = new(DefaultValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_defaultValue
	return p
}

func (*DefaultValueContext) IsDefaultValueContext() {}

func NewDefaultValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultValueContext {
	var p = new(DefaultValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_defaultValue

	return p
}

func (s *DefaultValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultValueContext) NULL() antlr.TerminalNode {
	return s.GetToken(CreateTableParserNULL, 0)
}

func (s *DefaultValueContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *DefaultValueContext) UnaryOperator() IUnaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorContext)
}

func (s *DefaultValueContext) AllCurrentTimestamp() []ICurrentTimestampContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICurrentTimestampContext)(nil)).Elem())
	var tst = make([]ICurrentTimestampContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICurrentTimestampContext)
		}
	}

	return tst
}

func (s *DefaultValueContext) CurrentTimestamp(i int) ICurrentTimestampContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurrentTimestampContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICurrentTimestampContext)
}

func (s *DefaultValueContext) ON() antlr.TerminalNode {
	return s.GetToken(CreateTableParserON, 0)
}

func (s *DefaultValueContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserUPDATE, 0)
}

func (s *DefaultValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterDefaultValue(s)
	}
}

func (s *DefaultValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitDefaultValue(s)
	}
}

func (p *CreateTableParser) DefaultValue() (localctx IDefaultValueContext) {
	localctx = NewDefaultValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CreateTableParserRULE_defaultValue)
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

	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(261)
			p.Match(CreateTableParserNULL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(263)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(262)
				p.UnaryOperator()
			}

		}
		{
			p.SetState(265)
			p.Constant()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(266)
			p.CurrentTimestamp()
		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserON {
			{
				p.SetState(267)
				p.Match(CreateTableParserON)
			}
			{
				p.SetState(268)
				p.Match(CreateTableParserUPDATE)
			}
			{
				p.SetState(269)
				p.CurrentTimestamp()
			}

		}

	}

	return localctx
}

// IUnaryOperatorContext is an interface to support dynamic dispatch.
type IUnaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperatorContext differentiates from other interfaces.
	IsUnaryOperatorContext()
}

type UnaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorContext() *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_unaryOperator
	return p
}

func (*UnaryOperatorContext) IsUnaryOperatorContext() {}

func NewUnaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_unaryOperator

	return p
}

func (s *UnaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserNOT, 0)
}

func (s *UnaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitUnaryOperator(s)
	}
}

func (p *CreateTableParser) UnaryOperator() (localctx IUnaryOperatorContext) {
	localctx = NewUnaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CreateTableParserRULE_unaryOperator)
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
	{
		p.SetState(274)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CreateTableParserT__4)|(1<<CreateTableParserT__5)|(1<<CreateTableParserT__6)|(1<<CreateTableParserT__7))) != 0) || _la == CreateTableParserNOT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) Literal() antlr.TerminalNode {
	return s.GetToken(CreateTableParserLiteral, 0)
}

func (s *ConstantContext) Integer() antlr.TerminalNode {
	return s.GetToken(CreateTableParserInteger, 0)
}

func (s *ConstantContext) TRUE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserTRUE, 0)
}

func (s *ConstantContext) FALSE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserFALSE, 0)
}

func (s *ConstantContext) Number() antlr.TerminalNode {
	return s.GetToken(CreateTableParserNumber, 0)
}

func (s *ConstantContext) NULL() antlr.TerminalNode {
	return s.GetToken(CreateTableParserNULL, 0)
}

func (s *ConstantContext) NOT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserNOT, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *CreateTableParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CreateTableParserRULE_constant)
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

	p.SetState(286)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CreateTableParserLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)
			p.Match(CreateTableParserLiteral)
		}

	case CreateTableParserInteger:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.Match(CreateTableParserInteger)
		}

	case CreateTableParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(278)
			p.Match(CreateTableParserT__7)
		}
		{
			p.SetState(279)
			p.Match(CreateTableParserInteger)
		}

	case CreateTableParserFALSE, CreateTableParserTRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(280)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CreateTableParserFALSE || _la == CreateTableParserTRUE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case CreateTableParserNumber:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(281)
			p.Match(CreateTableParserNumber)
		}

	case CreateTableParserNOT, CreateTableParserNULL:
		p.EnterOuterAlt(localctx, 6)
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserNOT {
			{
				p.SetState(282)
				p.Match(CreateTableParserNOT)
			}

		}
		{
			p.SetState(285)
			p.Match(CreateTableParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICurrentTimestampContext is an interface to support dynamic dispatch.
type ICurrentTimestampContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCurrentTimestampContext differentiates from other interfaces.
	IsCurrentTimestampContext()
}

type CurrentTimestampContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurrentTimestampContext() *CurrentTimestampContext {
	var p = new(CurrentTimestampContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_currentTimestamp
	return p
}

func (*CurrentTimestampContext) IsCurrentTimestampContext() {}

func NewCurrentTimestampContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurrentTimestampContext {
	var p = new(CurrentTimestampContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_currentTimestamp

	return p
}

func (s *CurrentTimestampContext) GetParser() antlr.Parser { return s.parser }

func (s *CurrentTimestampContext) NOW() antlr.TerminalNode {
	return s.GetToken(CreateTableParserNOW, 0)
}

func (s *CurrentTimestampContext) CURRENT_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCURRENT_TIMESTAMP, 0)
}

func (s *CurrentTimestampContext) LOCALTIME() antlr.TerminalNode {
	return s.GetToken(CreateTableParserLOCALTIME, 0)
}

func (s *CurrentTimestampContext) LOCALTIMESTAMP() antlr.TerminalNode {
	return s.GetToken(CreateTableParserLOCALTIMESTAMP, 0)
}

func (s *CurrentTimestampContext) Integer() antlr.TerminalNode {
	return s.GetToken(CreateTableParserInteger, 0)
}

func (s *CurrentTimestampContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurrentTimestampContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurrentTimestampContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterCurrentTimestamp(s)
	}
}

func (s *CurrentTimestampContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitCurrentTimestamp(s)
	}
}

func (p *CreateTableParser) CurrentTimestamp() (localctx ICurrentTimestampContext) {
	localctx = NewCurrentTimestampContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CreateTableParserRULE_currentTimestamp)
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
	p.SetState(302)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CreateTableParserCURRENT_TIMESTAMP, CreateTableParserLOCALTIMESTAMP, CreateTableParserLOCALTIME:
		{
			p.SetState(288)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-25)&-(0x1f+1)) == 0 && ((1<<uint((_la-25)))&((1<<(CreateTableParserCURRENT_TIMESTAMP-25))|(1<<(CreateTableParserLOCALTIMESTAMP-25))|(1<<(CreateTableParserLOCALTIME-25)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__2 {
			{
				p.SetState(289)
				p.Match(CreateTableParserT__2)
			}
			p.SetState(291)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CreateTableParserInteger {
				{
					p.SetState(290)
					p.Match(CreateTableParserInteger)
				}

			}
			{
				p.SetState(293)
				p.Match(CreateTableParserT__3)
			}

		}

	case CreateTableParserNOW:
		{
			p.SetState(296)
			p.Match(CreateTableParserNOW)
		}
		{
			p.SetState(297)
			p.Match(CreateTableParserT__2)
		}
		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserInteger {
			{
				p.SetState(298)
				p.Match(CreateTableParserInteger)
			}

		}
		{
			p.SetState(301)
			p.Match(CreateTableParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITableConstraintContext is an interface to support dynamic dispatch.
type ITableConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetIndex returns the index token.
	GetIndex() antlr.Token

	// GetIndexFormat returns the indexFormat token.
	GetIndexFormat() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetIndex sets the index token.
	SetIndex(antlr.Token)

	// SetIndexFormat sets the indexFormat token.
	SetIndexFormat(antlr.Token)

	// IsTableConstraintContext differentiates from other interfaces.
	IsTableConstraintContext()
}

type TableConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	name        antlr.Token
	index       antlr.Token
	indexFormat antlr.Token
}

func NewEmptyTableConstraintContext() *TableConstraintContext {
	var p = new(TableConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_tableConstraint
	return p
}

func (*TableConstraintContext) IsTableConstraintContext() {}

func NewTableConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableConstraintContext {
	var p = new(TableConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_tableConstraint

	return p
}

func (s *TableConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *TableConstraintContext) GetName() antlr.Token { return s.name }

func (s *TableConstraintContext) GetIndex() antlr.Token { return s.index }

func (s *TableConstraintContext) GetIndexFormat() antlr.Token { return s.indexFormat }

func (s *TableConstraintContext) SetName(v antlr.Token) { s.name = v }

func (s *TableConstraintContext) SetIndex(v antlr.Token) { s.index = v }

func (s *TableConstraintContext) SetIndexFormat(v antlr.Token) { s.indexFormat = v }

func (s *TableConstraintContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(CreateTableParserPRIMARY, 0)
}

func (s *TableConstraintContext) KEY() antlr.TerminalNode {
	return s.GetToken(CreateTableParserKEY, 0)
}

func (s *TableConstraintContext) IndexColumnNames() IIndexColumnNamesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexColumnNamesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexColumnNamesContext)
}

func (s *TableConstraintContext) CONSTRAINT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCONSTRAINT, 0)
}

func (s *TableConstraintContext) AllIndexOption() []IIndexOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIndexOptionContext)(nil)).Elem())
	var tst = make([]IIndexOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIndexOptionContext)
		}
	}

	return tst
}

func (s *TableConstraintContext) IndexOption(i int) IIndexOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIndexOptionContext)
}

func (s *TableConstraintContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(CreateTableParserBackQuotedString, 0)
}

func (s *TableConstraintContext) AllRawString() []antlr.TerminalNode {
	return s.GetTokens(CreateTableParserRawString)
}

func (s *TableConstraintContext) RawString(i int) antlr.TerminalNode {
	return s.GetToken(CreateTableParserRawString, i)
}

func (s *TableConstraintContext) AllLiteral() []antlr.TerminalNode {
	return s.GetTokens(CreateTableParserLiteral)
}

func (s *TableConstraintContext) Literal(i int) antlr.TerminalNode {
	return s.GetToken(CreateTableParserLiteral, i)
}

func (s *TableConstraintContext) UNIQUE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserUNIQUE, 0)
}

func (s *TableConstraintContext) IndexType() IIndexTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexTypeContext)
}

func (s *TableConstraintContext) INDEX() antlr.TerminalNode {
	return s.GetToken(CreateTableParserINDEX, 0)
}

func (s *TableConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterTableConstraint(s)
	}
}

func (s *TableConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitTableConstraint(s)
	}
}

func (p *CreateTableParser) TableConstraint() (localctx ITableConstraintContext) {
	localctx = NewTableConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CreateTableParserRULE_tableConstraint)
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

	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserCONSTRAINT {
			{
				p.SetState(304)
				p.Match(CreateTableParserCONSTRAINT)
			}
			p.SetState(306)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CreateTableParserLiteral || _la == CreateTableParserRawString {
				{
					p.SetState(305)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*TableConstraintContext).name = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CreateTableParserLiteral || _la == CreateTableParserRawString) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*TableConstraintContext).name = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			}

		}
		{
			p.SetState(310)
			p.Match(CreateTableParserPRIMARY)
		}
		{
			p.SetState(311)
			p.Match(CreateTableParserKEY)
		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(312)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*TableConstraintContext).index = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == CreateTableParserLiteral || _la == CreateTableParserRawString) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*TableConstraintContext).index = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserBackQuotedString || _la == CreateTableParserRawString {
			{
				p.SetState(315)
				_la = p.GetTokenStream().LA(1)

				if !(_la == CreateTableParserBackQuotedString || _la == CreateTableParserRawString) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(318)
			p.IndexColumnNames()
		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CreateTableParserCOMMENT || _la == CreateTableParserKEY_BLOCK_SIZE || _la == CreateTableParserWITH {
			{
				p.SetState(319)
				p.IndexOption()
			}

			p.SetState(324)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserCONSTRAINT {
			{
				p.SetState(325)
				p.Match(CreateTableParserCONSTRAINT)
			}
			p.SetState(327)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(326)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*TableConstraintContext).name = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CreateTableParserLiteral || _la == CreateTableParserRawString) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*TableConstraintContext).name = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			}

		}
		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserUNIQUE {
			{
				p.SetState(331)
				p.Match(CreateTableParserUNIQUE)
			}

		}
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserINDEX || _la == CreateTableParserKEY {
			{
				p.SetState(334)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*TableConstraintContext).indexFormat = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == CreateTableParserINDEX || _la == CreateTableParserKEY) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*TableConstraintContext).indexFormat = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserBackQuotedString || _la == CreateTableParserRawString {
			{
				p.SetState(337)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*TableConstraintContext).index = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == CreateTableParserBackQuotedString || _la == CreateTableParserRawString) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*TableConstraintContext).index = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserUSING {
			{
				p.SetState(340)
				p.IndexType()
			}

		}
		{
			p.SetState(343)
			p.IndexColumnNames()
		}
		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CreateTableParserCOMMENT || _la == CreateTableParserKEY_BLOCK_SIZE || _la == CreateTableParserWITH {
			{
				p.SetState(344)
				p.IndexOption()
			}

			p.SetState(349)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IIndexOptionContext is an interface to support dynamic dispatch.
type IIndexOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexOptionContext differentiates from other interfaces.
	IsIndexOptionContext()
}

type IndexOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexOptionContext() *IndexOptionContext {
	var p = new(IndexOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_indexOption
	return p
}

func (*IndexOptionContext) IsIndexOptionContext() {}

func NewIndexOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexOptionContext {
	var p = new(IndexOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_indexOption

	return p
}

func (s *IndexOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexOptionContext) KEY_BLOCK_SIZE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserKEY_BLOCK_SIZE, 0)
}

func (s *IndexOptionContext) FilesizeLiteral() antlr.TerminalNode {
	return s.GetToken(CreateTableParserFilesizeLiteral, 0)
}

func (s *IndexOptionContext) WITH() antlr.TerminalNode {
	return s.GetToken(CreateTableParserWITH, 0)
}

func (s *IndexOptionContext) PARSER() antlr.TerminalNode {
	return s.GetToken(CreateTableParserPARSER, 0)
}

func (s *IndexOptionContext) Literal() antlr.TerminalNode {
	return s.GetToken(CreateTableParserLiteral, 0)
}

func (s *IndexOptionContext) RawString() antlr.TerminalNode {
	return s.GetToken(CreateTableParserRawString, 0)
}

func (s *IndexOptionContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCOMMENT, 0)
}

func (s *IndexOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterIndexOption(s)
	}
}

func (s *IndexOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitIndexOption(s)
	}
}

func (p *CreateTableParser) IndexOption() (localctx IIndexOptionContext) {
	localctx = NewIndexOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CreateTableParserRULE_indexOption)
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

	p.SetState(362)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CreateTableParserKEY_BLOCK_SIZE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(352)
			p.Match(CreateTableParserKEY_BLOCK_SIZE)
		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__8 {
			{
				p.SetState(353)
				p.Match(CreateTableParserT__8)
			}

		}
		{
			p.SetState(356)
			p.Match(CreateTableParserFilesizeLiteral)
		}

	case CreateTableParserWITH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
			p.Match(CreateTableParserWITH)
		}
		{
			p.SetState(358)
			p.Match(CreateTableParserPARSER)
		}
		{
			p.SetState(359)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CreateTableParserLiteral || _la == CreateTableParserRawString) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case CreateTableParserCOMMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(360)
			p.Match(CreateTableParserCOMMENT)
		}
		{
			p.SetState(361)
			p.Match(CreateTableParserLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIndexTypeContext is an interface to support dynamic dispatch.
type IIndexTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexTypeContext differentiates from other interfaces.
	IsIndexTypeContext()
}

type IndexTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexTypeContext() *IndexTypeContext {
	var p = new(IndexTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_indexType
	return p
}

func (*IndexTypeContext) IsIndexTypeContext() {}

func NewIndexTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexTypeContext {
	var p = new(IndexTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_indexType

	return p
}

func (s *IndexTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexTypeContext) USING() antlr.TerminalNode {
	return s.GetToken(CreateTableParserUSING, 0)
}

func (s *IndexTypeContext) BTREE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserBTREE, 0)
}

func (s *IndexTypeContext) HASH() antlr.TerminalNode {
	return s.GetToken(CreateTableParserHASH, 0)
}

func (s *IndexTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterIndexType(s)
	}
}

func (s *IndexTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitIndexType(s)
	}
}

func (p *CreateTableParser) IndexType() (localctx IIndexTypeContext) {
	localctx = NewIndexTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CreateTableParserRULE_indexType)
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
	{
		p.SetState(364)
		p.Match(CreateTableParserUSING)
	}
	{
		p.SetState(365)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CreateTableParserBTREE || _la == CreateTableParserHASH) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIndexColumnNamesContext is an interface to support dynamic dispatch.
type IIndexColumnNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexColumnNamesContext differentiates from other interfaces.
	IsIndexColumnNamesContext()
}

type IndexColumnNamesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexColumnNamesContext() *IndexColumnNamesContext {
	var p = new(IndexColumnNamesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_indexColumnNames
	return p
}

func (*IndexColumnNamesContext) IsIndexColumnNamesContext() {}

func NewIndexColumnNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexColumnNamesContext {
	var p = new(IndexColumnNamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_indexColumnNames

	return p
}

func (s *IndexColumnNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexColumnNamesContext) AllBackQuotedString() []antlr.TerminalNode {
	return s.GetTokens(CreateTableParserBackQuotedString)
}

func (s *IndexColumnNamesContext) BackQuotedString(i int) antlr.TerminalNode {
	return s.GetToken(CreateTableParserBackQuotedString, i)
}

func (s *IndexColumnNamesContext) AllRawString() []antlr.TerminalNode {
	return s.GetTokens(CreateTableParserRawString)
}

func (s *IndexColumnNamesContext) RawString(i int) antlr.TerminalNode {
	return s.GetToken(CreateTableParserRawString, i)
}

func (s *IndexColumnNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexColumnNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexColumnNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterIndexColumnNames(s)
	}
}

func (s *IndexColumnNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitIndexColumnNames(s)
	}
}

func (p *CreateTableParser) IndexColumnNames() (localctx IIndexColumnNamesContext) {
	localctx = NewIndexColumnNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CreateTableParserRULE_indexColumnNames)
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
	{
		p.SetState(367)
		p.Match(CreateTableParserT__2)
	}
	{
		p.SetState(368)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CreateTableParserBackQuotedString || _la == CreateTableParserRawString) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CreateTableParserT__1 || _la == CreateTableParserRawString {
		p.SetState(372)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CreateTableParserT__1:
			{
				p.SetState(369)
				p.Match(CreateTableParserT__1)
			}
			{
				p.SetState(370)
				p.Match(CreateTableParserBackQuotedString)
			}

		case CreateTableParserRawString:
			{
				p.SetState(371)
				p.Match(CreateTableParserRawString)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(377)
		p.Match(CreateTableParserT__3)
	}

	return localctx
}

// ITableOptionContext is an interface to support dynamic dispatch.
type ITableOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableOptionContext differentiates from other interfaces.
	IsTableOptionContext()
}

type TableOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableOptionContext() *TableOptionContext {
	var p = new(TableOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CreateTableParserRULE_tableOption
	return p
}

func (*TableOptionContext) IsTableOptionContext() {}

func NewTableOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableOptionContext {
	var p = new(TableOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CreateTableParserRULE_tableOption

	return p
}

func (s *TableOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *TableOptionContext) ENGINE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserENGINE, 0)
}

func (s *TableOptionContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(CreateTableParserBackQuotedString, 0)
}

func (s *TableOptionContext) RawString() antlr.TerminalNode {
	return s.GetToken(CreateTableParserRawString, 0)
}

func (s *TableOptionContext) AUTO_INCREMENT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserAUTO_INCREMENT, 0)
}

func (s *TableOptionContext) Integer() antlr.TerminalNode {
	return s.GetToken(CreateTableParserInteger, 0)
}

func (s *TableOptionContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCHARACTER, 0)
}

func (s *TableOptionContext) SET() antlr.TerminalNode {
	return s.GetToken(CreateTableParserSET, 0)
}

func (s *TableOptionContext) CHARSET() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCHARSET, 0)
}

func (s *TableOptionContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserDEFAULT, 0)
}

func (s *TableOptionContext) COLLATE() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCOLLATE, 0)
}

func (s *TableOptionContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(CreateTableParserCOMMENT, 0)
}

func (s *TableOptionContext) Literal() antlr.TerminalNode {
	return s.GetToken(CreateTableParserLiteral, 0)
}

func (s *TableOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.EnterTableOption(s)
	}
}

func (s *TableOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CreateTableListener); ok {
		listenerT.ExitTableOption(s)
	}
}

func (p *CreateTableParser) TableOption() (localctx ITableOptionContext) {
	localctx = NewTableOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CreateTableParserRULE_tableOption)
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

	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 78, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(379)
			p.Match(CreateTableParserENGINE)
		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__8 {
			{
				p.SetState(380)
				p.Match(CreateTableParserT__8)
			}

		}
		{
			p.SetState(383)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CreateTableParserBackQuotedString || _la == CreateTableParserRawString) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(384)
			p.Match(CreateTableParserAUTO_INCREMENT)
		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__8 {
			{
				p.SetState(385)
				p.Match(CreateTableParserT__8)
			}

		}
		{
			p.SetState(388)
			p.Match(CreateTableParserInteger)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserDEFAULT {
			{
				p.SetState(389)
				p.Match(CreateTableParserDEFAULT)
			}

		}
		p.SetState(395)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CreateTableParserCHARACTER:
			{
				p.SetState(392)
				p.Match(CreateTableParserCHARACTER)
			}
			{
				p.SetState(393)
				p.Match(CreateTableParserSET)
			}

		case CreateTableParserCHARSET:
			{
				p.SetState(394)
				p.Match(CreateTableParserCHARSET)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__8 {
			{
				p.SetState(397)
				p.Match(CreateTableParserT__8)
			}

		}
		{
			p.SetState(400)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CreateTableParserBackQuotedString || _la == CreateTableParserRawString) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserDEFAULT {
			{
				p.SetState(401)
				p.Match(CreateTableParserDEFAULT)
			}

		}
		{
			p.SetState(404)
			p.Match(CreateTableParserCOLLATE)
		}
		p.SetState(406)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__8 {
			{
				p.SetState(405)
				p.Match(CreateTableParserT__8)
			}

		}
		{
			p.SetState(408)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CreateTableParserBackQuotedString || _la == CreateTableParserRawString) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(409)
			p.Match(CreateTableParserCOMMENT)
		}
		p.SetState(411)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CreateTableParserT__8 {
			{
				p.SetState(410)
				p.Match(CreateTableParserT__8)
			}

		}
		{
			p.SetState(413)
			p.Match(CreateTableParserLiteral)
		}

	}

	return localctx
}

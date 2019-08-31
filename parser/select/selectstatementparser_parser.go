// Code generated from SelectStatementParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package selectparser // SelectStatementParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 308,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 5, 2, 48, 10, 2, 3, 2, 5, 2, 51, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 5, 3, 58, 10, 3, 3, 3, 3, 3, 5, 3, 62, 10, 3, 3, 3, 5, 3, 65, 10, 3,
	3, 3, 5, 3, 68, 10, 3, 3, 3, 5, 3, 71, 10, 3, 3, 4, 5, 4, 74, 10, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 80, 10, 4, 3, 4, 7, 4, 83, 10, 4, 12, 4, 14,
	4, 86, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 105, 10, 8, 12, 8, 14,
	8, 108, 11, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 114, 10, 8, 12, 8, 14, 8,
	117, 11, 8, 5, 8, 119, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 7, 10, 128, 10, 10, 12, 10, 14, 10, 131, 11, 10, 3, 11, 3, 11, 5, 11,
	135, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 141, 10, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 5, 12, 148, 10, 12, 3, 13, 3, 13, 3, 13, 5, 13,
	153, 10, 13, 3, 14, 3, 14, 3, 14, 5, 14, 158, 10, 14, 3, 15, 3, 15, 5,
	15, 162, 10, 15, 3, 15, 3, 15, 3, 15, 5, 15, 167, 10, 15, 7, 15, 169, 10,
	15, 12, 15, 14, 15, 172, 11, 15, 3, 16, 3, 16, 5, 16, 176, 10, 16, 3, 16,
	3, 16, 3, 16, 5, 16, 181, 10, 16, 7, 16, 183, 10, 16, 12, 16, 14, 16, 186,
	11, 16, 3, 17, 3, 17, 3, 17, 5, 17, 191, 10, 17, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 201, 10, 18, 3, 18, 3, 18, 3, 18,
	5, 18, 206, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 5, 18, 219, 10, 18, 3, 19, 3, 19, 3, 19, 5, 19,
	224, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 230, 10, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 7, 21, 237, 10, 21, 12, 21, 14, 21, 240, 11, 21,
	3, 21, 3, 21, 3, 21, 7, 21, 245, 10, 21, 12, 21, 14, 21, 248, 11, 21, 5,
	21, 250, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 5, 22, 304, 10, 22, 3, 23, 3, 23, 3, 23, 2, 2, 24, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 2, 8, 3, 2, 43, 44, 3, 2, 26, 27, 3, 2, 31, 32, 3, 2, 13, 14, 3, 2,
	16, 19, 3, 2, 36, 40, 2, 354, 2, 47, 3, 2, 2, 2, 4, 54, 3, 2, 2, 2, 6,
	73, 3, 2, 2, 2, 8, 87, 3, 2, 2, 2, 10, 93, 3, 2, 2, 2, 12, 96, 3, 2, 2,
	2, 14, 99, 3, 2, 2, 2, 16, 120, 3, 2, 2, 2, 18, 122, 3, 2, 2, 2, 20, 132,
	3, 2, 2, 2, 22, 136, 3, 2, 2, 2, 24, 152, 3, 2, 2, 2, 26, 157, 3, 2, 2,
	2, 28, 161, 3, 2, 2, 2, 30, 175, 3, 2, 2, 2, 32, 187, 3, 2, 2, 2, 34, 218,
	3, 2, 2, 2, 36, 223, 3, 2, 2, 2, 38, 229, 3, 2, 2, 2, 40, 249, 3, 2, 2,
	2, 42, 303, 3, 2, 2, 2, 44, 305, 3, 2, 2, 2, 46, 48, 5, 4, 3, 2, 47, 46,
	3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 50, 3, 2, 2, 2, 49, 51, 7, 9, 2, 2,
	50, 49, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 53, 7,
	2, 2, 3, 53, 3, 3, 2, 2, 2, 54, 57, 7, 20, 2, 2, 55, 58, 7, 12, 2, 2, 56,
	58, 5, 6, 4, 2, 57, 55, 3, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2,
	2, 59, 61, 5, 10, 6, 2, 60, 62, 5, 12, 7, 2, 61, 60, 3, 2, 2, 2, 61, 62,
	3, 2, 2, 2, 62, 64, 3, 2, 2, 2, 63, 65, 5, 14, 8, 2, 64, 63, 3, 2, 2, 2,
	64, 65, 3, 2, 2, 2, 65, 67, 3, 2, 2, 2, 66, 68, 5, 18, 10, 2, 67, 66, 3,
	2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 71, 5, 22, 12, 2,
	70, 69, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 5, 3, 2, 2, 2, 72, 74, 5, 8,
	5, 2, 73, 72, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 76,
	9, 2, 2, 2, 76, 84, 3, 2, 2, 2, 77, 79, 7, 11, 2, 2, 78, 80, 5, 8, 5, 2,
	79, 78, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 83, 9,
	2, 2, 2, 82, 77, 3, 2, 2, 2, 83, 86, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84,
	85, 3, 2, 2, 2, 85, 7, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 87, 88, 5, 44, 23,
	2, 88, 89, 7, 7, 2, 2, 89, 90, 9, 2, 2, 2, 90, 91, 7, 8, 2, 2, 91, 92,
	7, 35, 2, 2, 92, 9, 3, 2, 2, 2, 93, 94, 7, 22, 2, 2, 94, 95, 9, 2, 2, 2,
	95, 11, 3, 2, 2, 2, 96, 97, 7, 23, 2, 2, 97, 98, 5, 28, 15, 2, 98, 13,
	3, 2, 2, 2, 99, 100, 7, 29, 2, 2, 100, 101, 7, 30, 2, 2, 101, 106, 5, 16,
	9, 2, 102, 103, 7, 11, 2, 2, 103, 105, 5, 16, 9, 2, 104, 102, 3, 2, 2,
	2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107,
	118, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 110, 7, 24, 2, 2, 110, 115,
	5, 34, 18, 2, 111, 112, 9, 3, 2, 2, 112, 114, 5, 34, 18, 2, 113, 111, 3,
	2, 2, 2, 114, 117, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2,
	2, 116, 119, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 118, 109, 3, 2, 2, 2, 118,
	119, 3, 2, 2, 2, 119, 15, 3, 2, 2, 2, 120, 121, 9, 2, 2, 2, 121, 17, 3,
	2, 2, 2, 122, 123, 7, 28, 2, 2, 123, 124, 7, 30, 2, 2, 124, 129, 5, 20,
	11, 2, 125, 126, 7, 11, 2, 2, 126, 128, 5, 20, 11, 2, 127, 125, 3, 2, 2,
	2, 128, 131, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130,
	19, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 132, 134, 9, 2, 2, 2, 133, 135, 9,
	4, 2, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 21, 3, 2, 2,
	2, 136, 147, 7, 34, 2, 2, 137, 138, 5, 24, 13, 2, 138, 139, 7, 11, 2, 2,
	139, 141, 3, 2, 2, 2, 140, 137, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141,
	142, 3, 2, 2, 2, 142, 148, 5, 26, 14, 2, 143, 144, 5, 26, 14, 2, 144, 145,
	7, 33, 2, 2, 145, 146, 5, 24, 13, 2, 146, 148, 3, 2, 2, 2, 147, 140, 3,
	2, 2, 2, 147, 143, 3, 2, 2, 2, 148, 23, 3, 2, 2, 2, 149, 153, 5, 42, 22,
	2, 150, 153, 7, 41, 2, 2, 151, 153, 7, 15, 2, 2, 152, 149, 3, 2, 2, 2,
	152, 150, 3, 2, 2, 2, 152, 151, 3, 2, 2, 2, 153, 25, 3, 2, 2, 2, 154, 158,
	5, 42, 22, 2, 155, 158, 7, 15, 2, 2, 156, 158, 7, 41, 2, 2, 157, 154, 3,
	2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 156, 3, 2, 2, 2, 158, 27, 3, 2, 2,
	2, 159, 162, 5, 32, 17, 2, 160, 162, 5, 30, 16, 2, 161, 159, 3, 2, 2, 2,
	161, 160, 3, 2, 2, 2, 162, 170, 3, 2, 2, 2, 163, 166, 7, 27, 2, 2, 164,
	167, 5, 32, 17, 2, 165, 167, 5, 30, 16, 2, 166, 164, 3, 2, 2, 2, 166, 165,
	3, 2, 2, 2, 167, 169, 3, 2, 2, 2, 168, 163, 3, 2, 2, 2, 169, 172, 3, 2,
	2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 29, 3, 2, 2, 2,
	172, 170, 3, 2, 2, 2, 173, 176, 5, 32, 17, 2, 174, 176, 5, 34, 18, 2, 175,
	173, 3, 2, 2, 2, 175, 174, 3, 2, 2, 2, 176, 184, 3, 2, 2, 2, 177, 180,
	7, 26, 2, 2, 178, 181, 5, 32, 17, 2, 179, 181, 5, 34, 18, 2, 180, 178,
	3, 2, 2, 2, 180, 179, 3, 2, 2, 2, 181, 183, 3, 2, 2, 2, 182, 177, 3, 2,
	2, 2, 183, 186, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2,
	185, 31, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 187, 190, 7, 7, 2, 2, 188, 191,
	5, 30, 16, 2, 189, 191, 5, 28, 15, 2, 190, 188, 3, 2, 2, 2, 190, 189, 3,
	2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 193, 7, 8, 2, 2, 193, 33, 3, 2, 2,
	2, 194, 195, 9, 2, 2, 2, 195, 196, 7, 25, 2, 2, 196, 219, 5, 36, 19, 2,
	197, 198, 9, 2, 2, 2, 198, 200, 7, 3, 2, 2, 199, 201, 7, 5, 2, 2, 200,
	199, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 219,
	7, 6, 2, 2, 203, 205, 9, 2, 2, 2, 204, 206, 7, 5, 2, 2, 205, 204, 3, 2,
	2, 2, 205, 206, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 208, 7, 4, 2, 2,
	208, 209, 7, 7, 2, 2, 209, 210, 5, 40, 21, 2, 210, 211, 7, 8, 2, 2, 211,
	219, 3, 2, 2, 2, 212, 213, 9, 2, 2, 2, 213, 214, 9, 5, 2, 2, 214, 219,
	5, 38, 20, 2, 215, 216, 9, 2, 2, 2, 216, 217, 9, 6, 2, 2, 217, 219, 5,
	38, 20, 2, 218, 194, 3, 2, 2, 2, 218, 197, 3, 2, 2, 2, 218, 203, 3, 2,
	2, 2, 218, 212, 3, 2, 2, 2, 218, 215, 3, 2, 2, 2, 219, 35, 3, 2, 2, 2,
	220, 224, 5, 42, 22, 2, 221, 224, 7, 42, 2, 2, 222, 224, 7, 15, 2, 2, 223,
	220, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223, 222, 3, 2, 2, 2, 224, 37, 3,
	2, 2, 2, 225, 230, 5, 42, 22, 2, 226, 230, 7, 41, 2, 2, 227, 230, 7, 42,
	2, 2, 228, 230, 7, 15, 2, 2, 229, 225, 3, 2, 2, 2, 229, 226, 3, 2, 2, 2,
	229, 227, 3, 2, 2, 2, 229, 228, 3, 2, 2, 2, 230, 39, 3, 2, 2, 2, 231, 250,
	5, 42, 22, 2, 232, 250, 7, 15, 2, 2, 233, 238, 7, 41, 2, 2, 234, 235, 7,
	11, 2, 2, 235, 237, 7, 41, 2, 2, 236, 234, 3, 2, 2, 2, 237, 240, 3, 2,
	2, 2, 238, 236, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 250, 3, 2, 2, 2,
	240, 238, 3, 2, 2, 2, 241, 246, 7, 42, 2, 2, 242, 243, 7, 11, 2, 2, 243,
	245, 7, 42, 2, 2, 244, 242, 3, 2, 2, 2, 245, 248, 3, 2, 2, 2, 246, 244,
	3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 250, 3, 2, 2, 2, 248, 246, 3, 2,
	2, 2, 249, 231, 3, 2, 2, 2, 249, 232, 3, 2, 2, 2, 249, 233, 3, 2, 2, 2,
	249, 241, 3, 2, 2, 2, 250, 41, 3, 2, 2, 2, 251, 252, 7, 10, 2, 2, 252,
	304, 7, 44, 2, 2, 253, 254, 7, 10, 2, 2, 254, 304, 7, 3, 2, 2, 255, 256,
	7, 10, 2, 2, 256, 304, 7, 4, 2, 2, 257, 258, 7, 10, 2, 2, 258, 304, 7,
	5, 2, 2, 259, 260, 7, 10, 2, 2, 260, 304, 7, 6, 2, 2, 261, 262, 7, 10,
	2, 2, 262, 304, 7, 20, 2, 2, 263, 264, 7, 10, 2, 2, 264, 304, 7, 21, 2,
	2, 265, 266, 7, 10, 2, 2, 266, 304, 7, 22, 2, 2, 267, 268, 7, 10, 2, 2,
	268, 304, 7, 23, 2, 2, 269, 270, 7, 10, 2, 2, 270, 304, 7, 24, 2, 2, 271,
	272, 7, 10, 2, 2, 272, 304, 7, 25, 2, 2, 273, 274, 7, 10, 2, 2, 274, 304,
	7, 26, 2, 2, 275, 276, 7, 10, 2, 2, 276, 304, 7, 27, 2, 2, 277, 278, 7,
	10, 2, 2, 278, 304, 7, 28, 2, 2, 279, 280, 7, 10, 2, 2, 280, 304, 7, 29,
	2, 2, 281, 282, 7, 10, 2, 2, 282, 304, 7, 30, 2, 2, 283, 284, 7, 10, 2,
	2, 284, 304, 7, 31, 2, 2, 285, 286, 7, 10, 2, 2, 286, 304, 7, 32, 2, 2,
	287, 288, 7, 10, 2, 2, 288, 304, 7, 33, 2, 2, 289, 290, 7, 10, 2, 2, 290,
	304, 7, 34, 2, 2, 291, 292, 7, 10, 2, 2, 292, 304, 7, 35, 2, 2, 293, 294,
	7, 10, 2, 2, 294, 304, 7, 36, 2, 2, 295, 296, 7, 10, 2, 2, 296, 304, 7,
	37, 2, 2, 297, 298, 7, 10, 2, 2, 298, 304, 7, 38, 2, 2, 299, 300, 7, 10,
	2, 2, 300, 304, 7, 39, 2, 2, 301, 302, 7, 10, 2, 2, 302, 304, 7, 40, 2,
	2, 303, 251, 3, 2, 2, 2, 303, 253, 3, 2, 2, 2, 303, 255, 3, 2, 2, 2, 303,
	257, 3, 2, 2, 2, 303, 259, 3, 2, 2, 2, 303, 261, 3, 2, 2, 2, 303, 263,
	3, 2, 2, 2, 303, 265, 3, 2, 2, 2, 303, 267, 3, 2, 2, 2, 303, 269, 3, 2,
	2, 2, 303, 271, 3, 2, 2, 2, 303, 273, 3, 2, 2, 2, 303, 275, 3, 2, 2, 2,
	303, 277, 3, 2, 2, 2, 303, 279, 3, 2, 2, 2, 303, 281, 3, 2, 2, 2, 303,
	283, 3, 2, 2, 2, 303, 285, 3, 2, 2, 2, 303, 287, 3, 2, 2, 2, 303, 289,
	3, 2, 2, 2, 303, 291, 3, 2, 2, 2, 303, 293, 3, 2, 2, 2, 303, 295, 3, 2,
	2, 2, 303, 297, 3, 2, 2, 2, 303, 299, 3, 2, 2, 2, 303, 301, 3, 2, 2, 2,
	304, 43, 3, 2, 2, 2, 305, 306, 9, 7, 2, 2, 306, 45, 3, 2, 2, 2, 37, 47,
	50, 57, 61, 64, 67, 70, 73, 79, 84, 106, 115, 118, 129, 134, 140, 147,
	152, 157, 161, 166, 170, 175, 180, 184, 190, 200, 205, 218, 223, 229, 238,
	246, 249, 303,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "'('", "')'", "';'", "':'", "','", "'*'", "'='", "'!='",
	"'?'", "'>'", "'>='", "'<'", "'<='",
}
var symbolicNames = []string{
	"", "Is", "In", "Not", "Null", "LeftParenthesis", "RightParenthesis", "Semicolon",
	"Colon", "Comma", "Asterisk", "Equal", "NotEqual", "QuestionMark", "GT",
	"GTE", "LT", "LTE", "Select", "Delete", "From", "Where", "Having", "Like",
	"And", "Or", "Order", "Group", "By", "Asc", "Desc", "Offset", "Limit",
	"As", "Count", "Sum", "Max", "Min", "Avg", "Number", "Literal", "BackQuotedString",
	"RawString", "WS",
}

var ruleNames = []string{
	"statement", "selectStatement", "selectField", "selectAsPrefix", "fromClause",
	"whereClause", "groupByClause", "groupByField", "orderByClause", "orderByField",
	"limitClause", "offsetValue", "limitValue", "orExpression", "andExpression",
	"subExpression", "atomExpression", "likeExpression", "compareExpression",
	"inExpression", "colonField", "groupFunction",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SelectStatementParserParser struct {
	*antlr.BaseParser
}

func NewSelectStatementParserParser(input antlr.TokenStream) *SelectStatementParserParser {
	this := new(SelectStatementParserParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SelectStatementParser.g4"

	return this
}

// SelectStatementParserParser tokens.
const (
	SelectStatementParserParserEOF              = antlr.TokenEOF
	SelectStatementParserParserIs               = 1
	SelectStatementParserParserIn               = 2
	SelectStatementParserParserNot              = 3
	SelectStatementParserParserNull             = 4
	SelectStatementParserParserLeftParenthesis  = 5
	SelectStatementParserParserRightParenthesis = 6
	SelectStatementParserParserSemicolon        = 7
	SelectStatementParserParserColon            = 8
	SelectStatementParserParserComma            = 9
	SelectStatementParserParserAsterisk         = 10
	SelectStatementParserParserEqual            = 11
	SelectStatementParserParserNotEqual         = 12
	SelectStatementParserParserQuestionMark     = 13
	SelectStatementParserParserGT               = 14
	SelectStatementParserParserGTE              = 15
	SelectStatementParserParserLT               = 16
	SelectStatementParserParserLTE              = 17
	SelectStatementParserParserSelect           = 18
	SelectStatementParserParserDelete           = 19
	SelectStatementParserParserFrom             = 20
	SelectStatementParserParserWhere            = 21
	SelectStatementParserParserHaving           = 22
	SelectStatementParserParserLike             = 23
	SelectStatementParserParserAnd              = 24
	SelectStatementParserParserOr               = 25
	SelectStatementParserParserOrder            = 26
	SelectStatementParserParserGroup            = 27
	SelectStatementParserParserBy               = 28
	SelectStatementParserParserAsc              = 29
	SelectStatementParserParserDesc             = 30
	SelectStatementParserParserOffset           = 31
	SelectStatementParserParserLimit            = 32
	SelectStatementParserParserAs               = 33
	SelectStatementParserParserCount            = 34
	SelectStatementParserParserSum              = 35
	SelectStatementParserParserMax              = 36
	SelectStatementParserParserMin              = 37
	SelectStatementParserParserAvg              = 38
	SelectStatementParserParserNumber           = 39
	SelectStatementParserParserLiteral          = 40
	SelectStatementParserParserBackQuotedString = 41
	SelectStatementParserParserRawString        = 42
	SelectStatementParserParserWS               = 43
)

// SelectStatementParserParser rules.
const (
	SelectStatementParserParserRULE_statement         = 0
	SelectStatementParserParserRULE_selectStatement   = 1
	SelectStatementParserParserRULE_selectField       = 2
	SelectStatementParserParserRULE_selectAsPrefix    = 3
	SelectStatementParserParserRULE_fromClause        = 4
	SelectStatementParserParserRULE_whereClause       = 5
	SelectStatementParserParserRULE_groupByClause     = 6
	SelectStatementParserParserRULE_groupByField      = 7
	SelectStatementParserParserRULE_orderByClause     = 8
	SelectStatementParserParserRULE_orderByField      = 9
	SelectStatementParserParserRULE_limitClause       = 10
	SelectStatementParserParserRULE_offsetValue       = 11
	SelectStatementParserParserRULE_limitValue        = 12
	SelectStatementParserParserRULE_orExpression      = 13
	SelectStatementParserParserRULE_andExpression     = 14
	SelectStatementParserParserRULE_subExpression     = 15
	SelectStatementParserParserRULE_atomExpression    = 16
	SelectStatementParserParserRULE_likeExpression    = 17
	SelectStatementParserParserRULE_compareExpression = 18
	SelectStatementParserParserRULE_inExpression      = 19
	SelectStatementParserParserRULE_colonField        = 20
	SelectStatementParserParserRULE_groupFunction     = 21
)

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) EOF() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserEOF, 0)
}

func (s *StatementContext) SelectStatement() ISelectStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *StatementContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserSemicolon, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *SelectStatementParserParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SelectStatementParserParserRULE_statement)
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserSelect {
		{
			p.SetState(44)
			p.SelectStatement()
		}

	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserSemicolon {
		{
			p.SetState(47)
			p.Match(SelectStatementParserParserSemicolon)
		}

	}
	{
		p.SetState(50)
		p.Match(SelectStatementParserParserEOF)
	}

	return localctx
}

// ISelectStatementContext is an interface to support dynamic dispatch.
type ISelectStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectStatementContext differentiates from other interfaces.
	IsSelectStatementContext()
}

type SelectStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStatementContext() *SelectStatementContext {
	var p = new(SelectStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_selectStatement
	return p
}

func (*SelectStatementContext) IsSelectStatementContext() {}

func NewSelectStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStatementContext {
	var p = new(SelectStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_selectStatement

	return p
}

func (s *SelectStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStatementContext) Select() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserSelect, 0)
}

func (s *SelectStatementContext) FromClause() IFromClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFromClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFromClauseContext)
}

func (s *SelectStatementContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAsterisk, 0)
}

func (s *SelectStatementContext) SelectField() ISelectFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectFieldContext)
}

func (s *SelectStatementContext) WhereClause() IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *SelectStatementContext) GroupByClause() IGroupByClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupByClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupByClauseContext)
}

func (s *SelectStatementContext) OrderByClause() IOrderByClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderByClauseContext)
}

func (s *SelectStatementContext) LimitClause() ILimitClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *SelectStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterSelectStatement(s)
	}
}

func (s *SelectStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitSelectStatement(s)
	}
}

func (p *SelectStatementParserParser) SelectStatement() (localctx ISelectStatementContext) {
	localctx = NewSelectStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SelectStatementParserParserRULE_selectStatement)
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
		p.SetState(52)
		p.Match(SelectStatementParserParserSelect)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SelectStatementParserParserAsterisk:
		{
			p.SetState(53)
			p.Match(SelectStatementParserParserAsterisk)
		}

	case SelectStatementParserParserCount, SelectStatementParserParserSum, SelectStatementParserParserMax, SelectStatementParserParserMin, SelectStatementParserParserAvg, SelectStatementParserParserBackQuotedString, SelectStatementParserParserRawString:
		{
			p.SetState(54)
			p.SelectField()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(57)
		p.FromClause()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserWhere {
		{
			p.SetState(58)
			p.WhereClause()
		}

	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserGroup {
		{
			p.SetState(61)
			p.GroupByClause()
		}

	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserOrder {
		{
			p.SetState(64)
			p.OrderByClause()
		}

	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserLimit {
		{
			p.SetState(67)
			p.LimitClause()
		}

	}

	return localctx
}

// ISelectFieldContext is an interface to support dynamic dispatch.
type ISelectFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectFieldContext differentiates from other interfaces.
	IsSelectFieldContext()
}

type SelectFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectFieldContext() *SelectFieldContext {
	var p = new(SelectFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_selectField
	return p
}

func (*SelectFieldContext) IsSelectFieldContext() {}

func NewSelectFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectFieldContext {
	var p = new(SelectFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_selectField

	return p
}

func (s *SelectFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectFieldContext) AllRawString() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserRawString)
}

func (s *SelectFieldContext) RawString(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRawString, i)
}

func (s *SelectFieldContext) AllBackQuotedString() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserBackQuotedString)
}

func (s *SelectFieldContext) BackQuotedString(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBackQuotedString, i)
}

func (s *SelectFieldContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserComma)
}

func (s *SelectFieldContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserComma, i)
}

func (s *SelectFieldContext) AllSelectAsPrefix() []ISelectAsPrefixContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectAsPrefixContext)(nil)).Elem())
	var tst = make([]ISelectAsPrefixContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectAsPrefixContext)
		}
	}

	return tst
}

func (s *SelectFieldContext) SelectAsPrefix(i int) ISelectAsPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectAsPrefixContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectAsPrefixContext)
}

func (s *SelectFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterSelectField(s)
	}
}

func (s *SelectFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitSelectField(s)
	}
}

func (p *SelectStatementParserParser) SelectField() (localctx ISelectFieldContext) {
	localctx = NewSelectFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SelectStatementParserParserRULE_selectField)
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
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(SelectStatementParserParserCount-34))|(1<<(SelectStatementParserParserSum-34))|(1<<(SelectStatementParserParserMax-34))|(1<<(SelectStatementParserParserMin-34))|(1<<(SelectStatementParserParserAvg-34)))) != 0 {
		{
			p.SetState(70)
			p.SelectAsPrefix()
		}

	}
	{
		p.SetState(73)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SelectStatementParserParserComma {
		{
			p.SetState(75)
			p.Match(SelectStatementParserParserComma)
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(SelectStatementParserParserCount-34))|(1<<(SelectStatementParserParserSum-34))|(1<<(SelectStatementParserParserMax-34))|(1<<(SelectStatementParserParserMin-34))|(1<<(SelectStatementParserParserAvg-34)))) != 0 {
			{
				p.SetState(76)
				p.SelectAsPrefix()
			}

		}
		{
			p.SetState(79)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelectAsPrefixContext is an interface to support dynamic dispatch.
type ISelectAsPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectAsPrefixContext differentiates from other interfaces.
	IsSelectAsPrefixContext()
}

type SelectAsPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectAsPrefixContext() *SelectAsPrefixContext {
	var p = new(SelectAsPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_selectAsPrefix
	return p
}

func (*SelectAsPrefixContext) IsSelectAsPrefixContext() {}

func NewSelectAsPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectAsPrefixContext {
	var p = new(SelectAsPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_selectAsPrefix

	return p
}

func (s *SelectAsPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectAsPrefixContext) GroupFunction() IGroupFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupFunctionContext)
}

func (s *SelectAsPrefixContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLeftParenthesis, 0)
}

func (s *SelectAsPrefixContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRightParenthesis, 0)
}

func (s *SelectAsPrefixContext) As() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAs, 0)
}

func (s *SelectAsPrefixContext) RawString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRawString, 0)
}

func (s *SelectAsPrefixContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBackQuotedString, 0)
}

func (s *SelectAsPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectAsPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectAsPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterSelectAsPrefix(s)
	}
}

func (s *SelectAsPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitSelectAsPrefix(s)
	}
}

func (p *SelectStatementParserParser) SelectAsPrefix() (localctx ISelectAsPrefixContext) {
	localctx = NewSelectAsPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SelectStatementParserParserRULE_selectAsPrefix)
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
		p.SetState(85)
		p.GroupFunction()
	}
	{
		p.SetState(86)
		p.Match(SelectStatementParserParserLeftParenthesis)
	}
	{
		p.SetState(87)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(88)
		p.Match(SelectStatementParserParserRightParenthesis)
	}
	{
		p.SetState(89)
		p.Match(SelectStatementParserParserAs)
	}

	return localctx
}

// IFromClauseContext is an interface to support dynamic dispatch.
type IFromClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTable returns the table token.
	GetTable() antlr.Token

	// SetTable sets the table token.
	SetTable(antlr.Token)

	// IsFromClauseContext differentiates from other interfaces.
	IsFromClauseContext()
}

type FromClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	table  antlr.Token
}

func NewEmptyFromClauseContext() *FromClauseContext {
	var p = new(FromClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_fromClause
	return p
}

func (*FromClauseContext) IsFromClauseContext() {}

func NewFromClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromClauseContext {
	var p = new(FromClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_fromClause

	return p
}

func (s *FromClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FromClauseContext) GetTable() antlr.Token { return s.table }

func (s *FromClauseContext) SetTable(v antlr.Token) { s.table = v }

func (s *FromClauseContext) From() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserFrom, 0)
}

func (s *FromClauseContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBackQuotedString, 0)
}

func (s *FromClauseContext) RawString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRawString, 0)
}

func (s *FromClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FromClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterFromClause(s)
	}
}

func (s *FromClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitFromClause(s)
	}
}

func (p *SelectStatementParserParser) FromClause() (localctx IFromClauseContext) {
	localctx = NewFromClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SelectStatementParserParserRULE_fromClause)
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
		p.SetState(91)
		p.Match(SelectStatementParserParserFrom)
	}
	{
		p.SetState(92)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*FromClauseContext).table = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*FromClauseContext).table = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_whereClause
	return p
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) Where() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserWhere, 0)
}

func (s *WhereClauseContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (p *SelectStatementParserParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SelectStatementParserParserRULE_whereClause)

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
		p.Match(SelectStatementParserParserWhere)
	}
	{
		p.SetState(95)
		p.OrExpression()
	}

	return localctx
}

// IGroupByClauseContext is an interface to support dynamic dispatch.
type IGroupByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupByClauseContext differentiates from other interfaces.
	IsGroupByClauseContext()
}

type GroupByClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByClauseContext() *GroupByClauseContext {
	var p = new(GroupByClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_groupByClause
	return p
}

func (*GroupByClauseContext) IsGroupByClauseContext() {}

func NewGroupByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByClauseContext {
	var p = new(GroupByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_groupByClause

	return p
}

func (s *GroupByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByClauseContext) Group() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserGroup, 0)
}

func (s *GroupByClauseContext) By() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBy, 0)
}

func (s *GroupByClauseContext) AllGroupByField() []IGroupByFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGroupByFieldContext)(nil)).Elem())
	var tst = make([]IGroupByFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGroupByFieldContext)
		}
	}

	return tst
}

func (s *GroupByClauseContext) GroupByField(i int) IGroupByFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupByFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGroupByFieldContext)
}

func (s *GroupByClauseContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserComma)
}

func (s *GroupByClauseContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserComma, i)
}

func (s *GroupByClauseContext) Having() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserHaving, 0)
}

func (s *GroupByClauseContext) AllAtomExpression() []IAtomExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomExpressionContext)(nil)).Elem())
	var tst = make([]IAtomExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomExpressionContext)
		}
	}

	return tst
}

func (s *GroupByClauseContext) AtomExpression(i int) IAtomExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomExpressionContext)
}

func (s *GroupByClauseContext) AllAnd() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserAnd)
}

func (s *GroupByClauseContext) And(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAnd, i)
}

func (s *GroupByClauseContext) AllOr() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserOr)
}

func (s *GroupByClauseContext) Or(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserOr, i)
}

func (s *GroupByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterGroupByClause(s)
	}
}

func (s *GroupByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitGroupByClause(s)
	}
}

func (p *SelectStatementParserParser) GroupByClause() (localctx IGroupByClauseContext) {
	localctx = NewGroupByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SelectStatementParserParserRULE_groupByClause)
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
		p.SetState(97)
		p.Match(SelectStatementParserParserGroup)
	}
	{
		p.SetState(98)
		p.Match(SelectStatementParserParserBy)
	}
	{
		p.SetState(99)
		p.GroupByField()
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SelectStatementParserParserComma {
		{
			p.SetState(100)
			p.Match(SelectStatementParserParserComma)
		}
		{
			p.SetState(101)
			p.GroupByField()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserHaving {
		{
			p.SetState(107)
			p.Match(SelectStatementParserParserHaving)
		}
		{
			p.SetState(108)
			p.AtomExpression()
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SelectStatementParserParserAnd || _la == SelectStatementParserParserOr {
			{
				p.SetState(109)
				_la = p.GetTokenStream().LA(1)

				if !(_la == SelectStatementParserParserAnd || _la == SelectStatementParserParserOr) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(110)
				p.AtomExpression()
			}

			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IGroupByFieldContext is an interface to support dynamic dispatch.
type IGroupByFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetField returns the field token.
	GetField() antlr.Token

	// SetField sets the field token.
	SetField(antlr.Token)

	// IsGroupByFieldContext differentiates from other interfaces.
	IsGroupByFieldContext()
}

type GroupByFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	field  antlr.Token
}

func NewEmptyGroupByFieldContext() *GroupByFieldContext {
	var p = new(GroupByFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_groupByField
	return p
}

func (*GroupByFieldContext) IsGroupByFieldContext() {}

func NewGroupByFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByFieldContext {
	var p = new(GroupByFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_groupByField

	return p
}

func (s *GroupByFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByFieldContext) GetField() antlr.Token { return s.field }

func (s *GroupByFieldContext) SetField(v antlr.Token) { s.field = v }

func (s *GroupByFieldContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBackQuotedString, 0)
}

func (s *GroupByFieldContext) RawString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRawString, 0)
}

func (s *GroupByFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterGroupByField(s)
	}
}

func (s *GroupByFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitGroupByField(s)
	}
}

func (p *SelectStatementParserParser) GroupByField() (localctx IGroupByFieldContext) {
	localctx = NewGroupByFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SelectStatementParserParserRULE_groupByField)
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
		p.SetState(118)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*GroupByFieldContext).field = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*GroupByFieldContext).field = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOrderByClauseContext is an interface to support dynamic dispatch.
type IOrderByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderByClauseContext differentiates from other interfaces.
	IsOrderByClauseContext()
}

type OrderByClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByClauseContext() *OrderByClauseContext {
	var p = new(OrderByClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_orderByClause
	return p
}

func (*OrderByClauseContext) IsOrderByClauseContext() {}

func NewOrderByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByClauseContext {
	var p = new(OrderByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_orderByClause

	return p
}

func (s *OrderByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByClauseContext) Order() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserOrder, 0)
}

func (s *OrderByClauseContext) By() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBy, 0)
}

func (s *OrderByClauseContext) AllOrderByField() []IOrderByFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOrderByFieldContext)(nil)).Elem())
	var tst = make([]IOrderByFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOrderByFieldContext)
		}
	}

	return tst
}

func (s *OrderByClauseContext) OrderByField(i int) IOrderByFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOrderByFieldContext)
}

func (s *OrderByClauseContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserComma)
}

func (s *OrderByClauseContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserComma, i)
}

func (s *OrderByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterOrderByClause(s)
	}
}

func (s *OrderByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitOrderByClause(s)
	}
}

func (p *SelectStatementParserParser) OrderByClause() (localctx IOrderByClauseContext) {
	localctx = NewOrderByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SelectStatementParserParserRULE_orderByClause)
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
		p.SetState(120)
		p.Match(SelectStatementParserParserOrder)
	}
	{
		p.SetState(121)
		p.Match(SelectStatementParserParserBy)
	}
	{
		p.SetState(122)
		p.OrderByField()
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SelectStatementParserParserComma {
		{
			p.SetState(123)
			p.Match(SelectStatementParserParserComma)
		}
		{
			p.SetState(124)
			p.OrderByField()
		}

		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrderByFieldContext is an interface to support dynamic dispatch.
type IOrderByFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetField returns the field token.
	GetField() antlr.Token

	// GetOrder returns the order token.
	GetOrder() antlr.Token

	// SetField sets the field token.
	SetField(antlr.Token)

	// SetOrder sets the order token.
	SetOrder(antlr.Token)

	// IsOrderByFieldContext differentiates from other interfaces.
	IsOrderByFieldContext()
}

type OrderByFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	field  antlr.Token
	order  antlr.Token
}

func NewEmptyOrderByFieldContext() *OrderByFieldContext {
	var p = new(OrderByFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_orderByField
	return p
}

func (*OrderByFieldContext) IsOrderByFieldContext() {}

func NewOrderByFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByFieldContext {
	var p = new(OrderByFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_orderByField

	return p
}

func (s *OrderByFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByFieldContext) GetField() antlr.Token { return s.field }

func (s *OrderByFieldContext) GetOrder() antlr.Token { return s.order }

func (s *OrderByFieldContext) SetField(v antlr.Token) { s.field = v }

func (s *OrderByFieldContext) SetOrder(v antlr.Token) { s.order = v }

func (s *OrderByFieldContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBackQuotedString, 0)
}

func (s *OrderByFieldContext) RawString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRawString, 0)
}

func (s *OrderByFieldContext) Asc() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAsc, 0)
}

func (s *OrderByFieldContext) Desc() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserDesc, 0)
}

func (s *OrderByFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterOrderByField(s)
	}
}

func (s *OrderByFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitOrderByField(s)
	}
}

func (p *SelectStatementParserParser) OrderByField() (localctx IOrderByFieldContext) {
	localctx = NewOrderByFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SelectStatementParserParserRULE_orderByField)
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
		p.SetState(130)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*OrderByFieldContext).field = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*OrderByFieldContext).field = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserAsc || _la == SelectStatementParserParserDesc {
		{
			p.SetState(131)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OrderByFieldContext).order = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserAsc || _la == SelectStatementParserParserDesc) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*OrderByFieldContext).order = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_limitClause
	return p
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) Limit() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLimit, 0)
}

func (s *LimitClauseContext) LimitValue() ILimitValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitValueContext)
}

func (s *LimitClauseContext) Offset() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserOffset, 0)
}

func (s *LimitClauseContext) OffsetValue() IOffsetValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOffsetValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOffsetValueContext)
}

func (s *LimitClauseContext) Comma() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserComma, 0)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (p *SelectStatementParserParser) LimitClause() (localctx ILimitClauseContext) {
	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SelectStatementParserParserRULE_limitClause)

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
		p.SetState(134)
		p.Match(SelectStatementParserParserLimit)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.SetState(138)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(135)
				p.OffsetValue()
			}
			{
				p.SetState(136)
				p.Match(SelectStatementParserParserComma)
			}

		}
		{
			p.SetState(140)
			p.LimitValue()
		}

	case 2:
		{
			p.SetState(141)
			p.LimitValue()
		}
		{
			p.SetState(142)
			p.Match(SelectStatementParserParserOffset)
		}
		{
			p.SetState(143)
			p.OffsetValue()
		}

	}

	return localctx
}

// IOffsetValueContext is an interface to support dynamic dispatch.
type IOffsetValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOffsetValueContext differentiates from other interfaces.
	IsOffsetValueContext()
}

type OffsetValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffsetValueContext() *OffsetValueContext {
	var p = new(OffsetValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_offsetValue
	return p
}

func (*OffsetValueContext) IsOffsetValueContext() {}

func NewOffsetValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetValueContext {
	var p = new(OffsetValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_offsetValue

	return p
}

func (s *OffsetValueContext) GetParser() antlr.Parser { return s.parser }

func (s *OffsetValueContext) ColonField() IColonFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColonFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColonFieldContext)
}

func (s *OffsetValueContext) Number() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNumber, 0)
}

func (s *OffsetValueContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserQuestionMark, 0)
}

func (s *OffsetValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OffsetValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OffsetValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterOffsetValue(s)
	}
}

func (s *OffsetValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitOffsetValue(s)
	}
}

func (p *SelectStatementParserParser) OffsetValue() (localctx IOffsetValueContext) {
	localctx = NewOffsetValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SelectStatementParserParserRULE_offsetValue)

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

	switch p.GetTokenStream().LA(1) {
	case SelectStatementParserParserColon:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.ColonField()
		}

	case SelectStatementParserParserNumber:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Match(SelectStatementParserParserNumber)
		}

	case SelectStatementParserParserQuestionMark:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(149)
			p.Match(SelectStatementParserParserQuestionMark)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILimitValueContext is an interface to support dynamic dispatch.
type ILimitValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitValueContext differentiates from other interfaces.
	IsLimitValueContext()
}

type LimitValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitValueContext() *LimitValueContext {
	var p = new(LimitValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_limitValue
	return p
}

func (*LimitValueContext) IsLimitValueContext() {}

func NewLimitValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitValueContext {
	var p = new(LimitValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_limitValue

	return p
}

func (s *LimitValueContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitValueContext) ColonField() IColonFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColonFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColonFieldContext)
}

func (s *LimitValueContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserQuestionMark, 0)
}

func (s *LimitValueContext) Number() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNumber, 0)
}

func (s *LimitValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterLimitValue(s)
	}
}

func (s *LimitValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitLimitValue(s)
	}
}

func (p *SelectStatementParserParser) LimitValue() (localctx ILimitValueContext) {
	localctx = NewLimitValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SelectStatementParserParserRULE_limitValue)

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

	p.SetState(155)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SelectStatementParserParserColon:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.ColonField()
		}

	case SelectStatementParserParserQuestionMark:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.Match(SelectStatementParserParserQuestionMark)
		}

	case SelectStatementParserParserNumber:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(154)
			p.Match(SelectStatementParserParserNumber)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOrExpressionContext is an interface to support dynamic dispatch.
type IOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrExpressionContext differentiates from other interfaces.
	IsOrExpressionContext()
}

type OrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExpressionContext() *OrExpressionContext {
	var p = new(OrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_orExpression
	return p
}

func (*OrExpressionContext) IsOrExpressionContext() {}

func NewOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_orExpression

	return p
}

func (s *OrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExpressionContext) AllSubExpression() []ISubExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubExpressionContext)(nil)).Elem())
	var tst = make([]ISubExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubExpressionContext)
		}
	}

	return tst
}

func (s *OrExpressionContext) SubExpression(i int) ISubExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubExpressionContext)
}

func (s *OrExpressionContext) AllAndExpression() []IAndExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem())
	var tst = make([]IAndExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAndExpressionContext)
		}
	}

	return tst
}

func (s *OrExpressionContext) AndExpression(i int) IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *OrExpressionContext) AllOr() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserOr)
}

func (s *OrExpressionContext) Or(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserOr, i)
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (p *SelectStatementParserParser) OrExpression() (localctx IOrExpressionContext) {
	localctx = NewOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SelectStatementParserParserRULE_orExpression)
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
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(157)
			p.SubExpression()
		}

	case 2:
		{
			p.SetState(158)
			p.AndExpression()
		}

	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SelectStatementParserParserOr {
		{
			p.SetState(161)
			p.Match(SelectStatementParserParserOr)
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(162)
				p.SubExpression()
			}

		case 2:
			{
				p.SetState(163)
				p.AndExpression()
			}

		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAndExpressionContext is an interface to support dynamic dispatch.
type IAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_andExpression
	return p
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExpressionContext) AllSubExpression() []ISubExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubExpressionContext)(nil)).Elem())
	var tst = make([]ISubExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubExpressionContext)
		}
	}

	return tst
}

func (s *AndExpressionContext) SubExpression(i int) ISubExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubExpressionContext)
}

func (s *AndExpressionContext) AllAtomExpression() []IAtomExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomExpressionContext)(nil)).Elem())
	var tst = make([]IAtomExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomExpressionContext)
		}
	}

	return tst
}

func (s *AndExpressionContext) AtomExpression(i int) IAtomExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomExpressionContext)
}

func (s *AndExpressionContext) AllAnd() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserAnd)
}

func (s *AndExpressionContext) And(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAnd, i)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (p *SelectStatementParserParser) AndExpression() (localctx IAndExpressionContext) {
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SelectStatementParserParserRULE_andExpression)
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
	p.SetState(173)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SelectStatementParserParserLeftParenthesis:
		{
			p.SetState(171)
			p.SubExpression()
		}

	case SelectStatementParserParserBackQuotedString, SelectStatementParserParserRawString:
		{
			p.SetState(172)
			p.AtomExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SelectStatementParserParserAnd {
		{
			p.SetState(175)
			p.Match(SelectStatementParserParserAnd)
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SelectStatementParserParserLeftParenthesis:
			{
				p.SetState(176)
				p.SubExpression()
			}

		case SelectStatementParserParserBackQuotedString, SelectStatementParserParserRawString:
			{
				p.SetState(177)
				p.AtomExpression()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISubExpressionContext is an interface to support dynamic dispatch.
type ISubExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubExpressionContext differentiates from other interfaces.
	IsSubExpressionContext()
}

type SubExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubExpressionContext() *SubExpressionContext {
	var p = new(SubExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_subExpression
	return p
}

func (*SubExpressionContext) IsSubExpressionContext() {}

func NewSubExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubExpressionContext {
	var p = new(SubExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_subExpression

	return p
}

func (s *SubExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SubExpressionContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLeftParenthesis, 0)
}

func (s *SubExpressionContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRightParenthesis, 0)
}

func (s *SubExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *SubExpressionContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *SubExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterSubExpression(s)
	}
}

func (s *SubExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitSubExpression(s)
	}
}

func (p *SelectStatementParserParser) SubExpression() (localctx ISubExpressionContext) {
	localctx = NewSubExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SelectStatementParserParserRULE_subExpression)

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
		p.SetState(185)
		p.Match(SelectStatementParserParserLeftParenthesis)
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(186)
			p.AndExpression()
		}

	case 2:
		{
			p.SetState(187)
			p.OrExpression()
		}

	}
	{
		p.SetState(190)
		p.Match(SelectStatementParserParserRightParenthesis)
	}

	return localctx
}

// IAtomExpressionContext is an interface to support dynamic dispatch.
type IAtomExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetField returns the field token.
	GetField() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetField sets the field token.
	SetField(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsAtomExpressionContext differentiates from other interfaces.
	IsAtomExpressionContext()
}

type AtomExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	field  antlr.Token
	op     antlr.Token
}

func NewEmptyAtomExpressionContext() *AtomExpressionContext {
	var p = new(AtomExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_atomExpression
	return p
}

func (*AtomExpressionContext) IsAtomExpressionContext() {}

func NewAtomExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomExpressionContext {
	var p = new(AtomExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_atomExpression

	return p
}

func (s *AtomExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomExpressionContext) GetField() antlr.Token { return s.field }

func (s *AtomExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AtomExpressionContext) SetField(v antlr.Token) { s.field = v }

func (s *AtomExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AtomExpressionContext) Like() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLike, 0)
}

func (s *AtomExpressionContext) LikeExpression() ILikeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILikeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILikeExpressionContext)
}

func (s *AtomExpressionContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBackQuotedString, 0)
}

func (s *AtomExpressionContext) RawString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRawString, 0)
}

func (s *AtomExpressionContext) Is() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserIs, 0)
}

func (s *AtomExpressionContext) Null() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNull, 0)
}

func (s *AtomExpressionContext) Not() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNot, 0)
}

func (s *AtomExpressionContext) In() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserIn, 0)
}

func (s *AtomExpressionContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLeftParenthesis, 0)
}

func (s *AtomExpressionContext) InExpression() IInExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInExpressionContext)
}

func (s *AtomExpressionContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRightParenthesis, 0)
}

func (s *AtomExpressionContext) CompareExpression() ICompareExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareExpressionContext)
}

func (s *AtomExpressionContext) Equal() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserEqual, 0)
}

func (s *AtomExpressionContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNotEqual, 0)
}

func (s *AtomExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserGT, 0)
}

func (s *AtomExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLT, 0)
}

func (s *AtomExpressionContext) GTE() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserGTE, 0)
}

func (s *AtomExpressionContext) LTE() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLTE, 0)
}

func (s *AtomExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterAtomExpression(s)
	}
}

func (s *AtomExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitAtomExpression(s)
	}
}

func (p *SelectStatementParserParser) AtomExpression() (localctx IAtomExpressionContext) {
	localctx = NewAtomExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SelectStatementParserParserRULE_atomExpression)
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

	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(193)
			p.Match(SelectStatementParserParserLike)
		}
		{
			p.SetState(194)
			p.LikeExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(196)
			p.Match(SelectStatementParserParserIs)
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SelectStatementParserParserNot {
			{
				p.SetState(197)
				p.Match(SelectStatementParserParserNot)
			}

		}
		{
			p.SetState(200)
			p.Match(SelectStatementParserParserNull)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(201)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SelectStatementParserParserNot {
			{
				p.SetState(202)
				p.Match(SelectStatementParserParserNot)
			}

		}
		{
			p.SetState(205)
			p.Match(SelectStatementParserParserIn)
		}
		{
			p.SetState(206)
			p.Match(SelectStatementParserParserLeftParenthesis)
		}
		{
			p.SetState(207)
			p.InExpression()
		}
		{
			p.SetState(208)
			p.Match(SelectStatementParserParserRightParenthesis)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(210)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(211)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserEqual || _la == SelectStatementParserParserNotEqual) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(212)
			p.CompareExpression()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(213)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(214)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SelectStatementParserParserGT)|(1<<SelectStatementParserParserGTE)|(1<<SelectStatementParserParserLT)|(1<<SelectStatementParserParserLTE))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(215)
			p.CompareExpression()
		}

	}

	return localctx
}

// ILikeExpressionContext is an interface to support dynamic dispatch.
type ILikeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLikeExpressionContext differentiates from other interfaces.
	IsLikeExpressionContext()
}

type LikeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLikeExpressionContext() *LikeExpressionContext {
	var p = new(LikeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_likeExpression
	return p
}

func (*LikeExpressionContext) IsLikeExpressionContext() {}

func NewLikeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LikeExpressionContext {
	var p = new(LikeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_likeExpression

	return p
}

func (s *LikeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LikeExpressionContext) ColonField() IColonFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColonFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColonFieldContext)
}

func (s *LikeExpressionContext) Literal() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLiteral, 0)
}

func (s *LikeExpressionContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserQuestionMark, 0)
}

func (s *LikeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LikeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterLikeExpression(s)
	}
}

func (s *LikeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitLikeExpression(s)
	}
}

func (p *SelectStatementParserParser) LikeExpression() (localctx ILikeExpressionContext) {
	localctx = NewLikeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SelectStatementParserParserRULE_likeExpression)

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

	switch p.GetTokenStream().LA(1) {
	case SelectStatementParserParserColon:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)
			p.ColonField()
		}

	case SelectStatementParserParserLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.Match(SelectStatementParserParserLiteral)
		}

	case SelectStatementParserParserQuestionMark:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)
			p.Match(SelectStatementParserParserQuestionMark)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICompareExpressionContext is an interface to support dynamic dispatch.
type ICompareExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareExpressionContext differentiates from other interfaces.
	IsCompareExpressionContext()
}

type CompareExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareExpressionContext() *CompareExpressionContext {
	var p = new(CompareExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_compareExpression
	return p
}

func (*CompareExpressionContext) IsCompareExpressionContext() {}

func NewCompareExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareExpressionContext {
	var p = new(CompareExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_compareExpression

	return p
}

func (s *CompareExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareExpressionContext) ColonField() IColonFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColonFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColonFieldContext)
}

func (s *CompareExpressionContext) Number() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNumber, 0)
}

func (s *CompareExpressionContext) Literal() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLiteral, 0)
}

func (s *CompareExpressionContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserQuestionMark, 0)
}

func (s *CompareExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterCompareExpression(s)
	}
}

func (s *CompareExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitCompareExpression(s)
	}
}

func (p *SelectStatementParserParser) CompareExpression() (localctx ICompareExpressionContext) {
	localctx = NewCompareExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SelectStatementParserParserRULE_compareExpression)

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

	p.SetState(227)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SelectStatementParserParserColon:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.ColonField()
		}

	case SelectStatementParserParserNumber:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.Match(SelectStatementParserParserNumber)
		}

	case SelectStatementParserParserLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(225)
			p.Match(SelectStatementParserParserLiteral)
		}

	case SelectStatementParserParserQuestionMark:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(226)
			p.Match(SelectStatementParserParserQuestionMark)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInExpressionContext is an interface to support dynamic dispatch.
type IInExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInExpressionContext differentiates from other interfaces.
	IsInExpressionContext()
}

type InExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInExpressionContext() *InExpressionContext {
	var p = new(InExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_inExpression
	return p
}

func (*InExpressionContext) IsInExpressionContext() {}

func NewInExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InExpressionContext {
	var p = new(InExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_inExpression

	return p
}

func (s *InExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *InExpressionContext) ColonField() IColonFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColonFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColonFieldContext)
}

func (s *InExpressionContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserQuestionMark, 0)
}

func (s *InExpressionContext) AllNumber() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserNumber)
}

func (s *InExpressionContext) Number(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNumber, i)
}

func (s *InExpressionContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserComma)
}

func (s *InExpressionContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserComma, i)
}

func (s *InExpressionContext) AllLiteral() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserLiteral)
}

func (s *InExpressionContext) Literal(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLiteral, i)
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterInExpression(s)
	}
}

func (s *InExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitInExpression(s)
	}
}

func (p *SelectStatementParserParser) InExpression() (localctx IInExpressionContext) {
	localctx = NewInExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SelectStatementParserParserRULE_inExpression)
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

	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SelectStatementParserParserColon:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.ColonField()
		}

	case SelectStatementParserParserQuestionMark:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.Match(SelectStatementParserParserQuestionMark)
		}

	case SelectStatementParserParserNumber:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(231)
			p.Match(SelectStatementParserParserNumber)
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SelectStatementParserParserComma {
			{
				p.SetState(232)
				p.Match(SelectStatementParserParserComma)
			}
			{
				p.SetState(233)
				p.Match(SelectStatementParserParserNumber)
			}

			p.SetState(238)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case SelectStatementParserParserLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(239)
			p.Match(SelectStatementParserParserLiteral)
		}
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SelectStatementParserParserComma {
			{
				p.SetState(240)
				p.Match(SelectStatementParserParserComma)
			}
			{
				p.SetState(241)
				p.Match(SelectStatementParserParserLiteral)
			}

			p.SetState(246)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IColonFieldContext is an interface to support dynamic dispatch.
type IColonFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColonFieldContext differentiates from other interfaces.
	IsColonFieldContext()
}

type ColonFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColonFieldContext() *ColonFieldContext {
	var p = new(ColonFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_colonField
	return p
}

func (*ColonFieldContext) IsColonFieldContext() {}

func NewColonFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColonFieldContext {
	var p = new(ColonFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_colonField

	return p
}

func (s *ColonFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ColonFieldContext) Colon() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserColon, 0)
}

func (s *ColonFieldContext) RawString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRawString, 0)
}

func (s *ColonFieldContext) Is() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserIs, 0)
}

func (s *ColonFieldContext) In() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserIn, 0)
}

func (s *ColonFieldContext) Not() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNot, 0)
}

func (s *ColonFieldContext) Null() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNull, 0)
}

func (s *ColonFieldContext) Select() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserSelect, 0)
}

func (s *ColonFieldContext) Delete() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserDelete, 0)
}

func (s *ColonFieldContext) From() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserFrom, 0)
}

func (s *ColonFieldContext) Where() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserWhere, 0)
}

func (s *ColonFieldContext) Having() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserHaving, 0)
}

func (s *ColonFieldContext) Like() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLike, 0)
}

func (s *ColonFieldContext) And() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAnd, 0)
}

func (s *ColonFieldContext) Or() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserOr, 0)
}

func (s *ColonFieldContext) Order() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserOrder, 0)
}

func (s *ColonFieldContext) Group() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserGroup, 0)
}

func (s *ColonFieldContext) By() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBy, 0)
}

func (s *ColonFieldContext) Asc() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAsc, 0)
}

func (s *ColonFieldContext) Desc() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserDesc, 0)
}

func (s *ColonFieldContext) Offset() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserOffset, 0)
}

func (s *ColonFieldContext) Limit() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLimit, 0)
}

func (s *ColonFieldContext) As() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAs, 0)
}

func (s *ColonFieldContext) Count() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserCount, 0)
}

func (s *ColonFieldContext) Sum() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserSum, 0)
}

func (s *ColonFieldContext) Max() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserMax, 0)
}

func (s *ColonFieldContext) Min() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserMin, 0)
}

func (s *ColonFieldContext) Avg() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAvg, 0)
}

func (s *ColonFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColonFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColonFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterColonField(s)
	}
}

func (s *ColonFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitColonField(s)
	}
}

func (p *SelectStatementParserParser) ColonField() (localctx IColonFieldContext) {
	localctx = NewColonFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SelectStatementParserParserRULE_colonField)

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

	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(250)
			p.Match(SelectStatementParserParserRawString)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(251)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(252)
			p.Match(SelectStatementParserParserIs)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(253)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(254)
			p.Match(SelectStatementParserParserIn)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(255)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(256)
			p.Match(SelectStatementParserParserNot)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(257)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(258)
			p.Match(SelectStatementParserParserNull)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(259)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(260)
			p.Match(SelectStatementParserParserSelect)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(261)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(262)
			p.Match(SelectStatementParserParserDelete)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(263)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(264)
			p.Match(SelectStatementParserParserFrom)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(265)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(266)
			p.Match(SelectStatementParserParserWhere)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(267)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(268)
			p.Match(SelectStatementParserParserHaving)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(269)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(270)
			p.Match(SelectStatementParserParserLike)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(271)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(272)
			p.Match(SelectStatementParserParserAnd)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(273)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(274)
			p.Match(SelectStatementParserParserOr)
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(275)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(276)
			p.Match(SelectStatementParserParserOrder)
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(277)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(278)
			p.Match(SelectStatementParserParserGroup)
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(279)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(280)
			p.Match(SelectStatementParserParserBy)
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(281)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(282)
			p.Match(SelectStatementParserParserAsc)
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(283)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(284)
			p.Match(SelectStatementParserParserDesc)
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(285)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(286)
			p.Match(SelectStatementParserParserOffset)
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(287)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(288)
			p.Match(SelectStatementParserParserLimit)
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(289)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(290)
			p.Match(SelectStatementParserParserAs)
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(291)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(292)
			p.Match(SelectStatementParserParserCount)
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(293)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(294)
			p.Match(SelectStatementParserParserSum)
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(295)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(296)
			p.Match(SelectStatementParserParserMax)
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(297)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(298)
			p.Match(SelectStatementParserParserMin)
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(299)
			p.Match(SelectStatementParserParserColon)
		}
		{
			p.SetState(300)
			p.Match(SelectStatementParserParserAvg)
		}

	}

	return localctx
}

// IGroupFunctionContext is an interface to support dynamic dispatch.
type IGroupFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupFunctionContext differentiates from other interfaces.
	IsGroupFunctionContext()
}

type GroupFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupFunctionContext() *GroupFunctionContext {
	var p = new(GroupFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_groupFunction
	return p
}

func (*GroupFunctionContext) IsGroupFunctionContext() {}

func NewGroupFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupFunctionContext {
	var p = new(GroupFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_groupFunction

	return p
}

func (s *GroupFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupFunctionContext) Count() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserCount, 0)
}

func (s *GroupFunctionContext) Sum() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserSum, 0)
}

func (s *GroupFunctionContext) Max() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserMax, 0)
}

func (s *GroupFunctionContext) Min() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserMin, 0)
}

func (s *GroupFunctionContext) Avg() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAvg, 0)
}

func (s *GroupFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterGroupFunction(s)
	}
}

func (s *GroupFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitGroupFunction(s)
	}
}

func (p *SelectStatementParserParser) GroupFunction() (localctx IGroupFunctionContext) {
	localctx = NewGroupFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SelectStatementParserParserRULE_groupFunction)
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
		p.SetState(303)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(SelectStatementParserParserCount-34))|(1<<(SelectStatementParserParserSum-34))|(1<<(SelectStatementParserParserMax-34))|(1<<(SelectStatementParserParserMin-34))|(1<<(SelectStatementParserParserAvg-34)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

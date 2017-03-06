//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
func setParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func setAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func incNesting(yylex interface{}) bool {
	yylex.(*Tokenizer).nesting++
	if yylex.(*Tokenizer).nesting == 200 {
		return true
	}
	return false
}

func decNesting(yylex interface{}) {
	yylex.(*Tokenizer).nesting--
}

func forceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

//line sql.y:34
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	boolVal     BoolVal
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	valTuple    ValTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
	colIdent    ColIdent
	colIdents   []ColIdent
	tableIdent  TableIdent
}

const LEX_ERROR = 57346
const UNION = 57347
const SELECT = 57348
const INSERT = 57349
const UPDATE = 57350
const DELETE = 57351
const FROM = 57352
const WHERE = 57353
const GROUP = 57354
const HAVING = 57355
const ORDER = 57356
const BY = 57357
const LIMIT = 57358
const FOR = 57359
const ALL = 57360
const DISTINCT = 57361
const AS = 57362
const EXISTS = 57363
const ASC = 57364
const DESC = 57365
const INTO = 57366
const DUPLICATE = 57367
const KEY = 57368
const DEFAULT = 57369
const SET = 57370
const LOCK = 57371
const VALUES = 57372
const LAST_INSERT_ID = 57373
const NEXT = 57374
const VALUE = 57375
const JOIN = 57376
const STRAIGHT_JOIN = 57377
const LEFT = 57378
const RIGHT = 57379
const INNER = 57380
const OUTER = 57381
const CROSS = 57382
const NATURAL = 57383
const USE = 57384
const FORCE = 57385
const ON = 57386
const ID = 57387
const HEX = 57388
const STRING = 57389
const INTEGRAL = 57390
const FLOAT = 57391
const HEXNUM = 57392
const VALUE_ARG = 57393
const LIST_ARG = 57394
const COMMENT = 57395
const NULL = 57396
const TRUE = 57397
const FALSE = 57398
const OR = 57399
const AND = 57400
const NOT = 57401
const BETWEEN = 57402
const CASE = 57403
const WHEN = 57404
const THEN = 57405
const ELSE = 57406
const END = 57407
const LE = 57408
const GE = 57409
const NE = 57410
const NULL_SAFE_EQUAL = 57411
const IS = 57412
const LIKE = 57413
const REGEXP = 57414
const IN = 57415
const SHIFT_LEFT = 57416
const SHIFT_RIGHT = 57417
const MOD = 57418
const UNARY = 57419
const INTERVAL = 57420
const JSON_EXTRACT_OP = 57421
const JSON_UNQUOTE_EXTRACT_OP = 57422
const CREATE = 57423
const ALTER = 57424
const DROP = 57425
const RENAME = 57426
const ANALYZE = 57427
const TABLE = 57428
const INDEX = 57429
const VIEW = 57430
const TO = 57431
const IGNORE = 57432
const IF = 57433
const UNIQUE = 57434
const USING = 57435
const SHOW = 57436
const DESCRIBE = 57437
const EXPLAIN = 57438
const XA = 57439
const PARTITION = 57440
const HASH = 57441
const CURRENT_TIMESTAMP = 57442
const DATABASE = 57443
const DATABASES = 57444
const TABLES = 57445
const UNUSED = 57446

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"UNION",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"ASC",
	"DESC",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"VALUES",
	"LAST_INSERT_ID",
	"NEXT",
	"VALUE",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"'('",
	"','",
	"')'",
	"ID",
	"HEX",
	"STRING",
	"INTEGRAL",
	"FLOAT",
	"HEXNUM",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"NULL",
	"TRUE",
	"FALSE",
	"OR",
	"AND",
	"NOT",
	"BETWEEN",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"'='",
	"'<'",
	"'>'",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"IS",
	"LIKE",
	"REGEXP",
	"IN",
	"'|'",
	"'&'",
	"SHIFT_LEFT",
	"SHIFT_RIGHT",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"MOD",
	"'^'",
	"'~'",
	"UNARY",
	"INTERVAL",
	"'.'",
	"JSON_EXTRACT_OP",
	"JSON_UNQUOTE_EXTRACT_OP",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
	"XA",
	"PARTITION",
	"HASH",
	"CURRENT_TIMESTAMP",
	"DATABASE",
	"DATABASES",
	"TABLES",
	"UNUSED",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 128,
	94, 253,
	-2, 252,
}

const yyNprod = 257
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1145

var yyAct = [...]int{

	140, 330, 261, 457, 126, 387, 341, 73, 250, 231,
	134, 53, 267, 351, 376, 285, 321, 265, 230, 3,
	121, 266, 278, 143, 194, 167, 122, 58, 132, 88,
	176, 49, 50, 407, 409, 42, 93, 74, 84, 78,
	269, 45, 75, 182, 59, 60, 51, 47, 56, 57,
	43, 260, 172, 61, 112, 440, 180, 439, 438, 80,
	81, 52, 48, 389, 135, 358, 97, 89, 90, 91,
	233, 234, 214, 215, 216, 217, 211, 184, 253, 198,
	69, 212, 213, 214, 215, 216, 217, 211, 103, 82,
	211, 463, 76, 87, 199, 128, 201, 408, 108, 106,
	292, 200, 199, 75, 428, 429, 75, 421, 419, 201,
	165, 322, 100, 104, 290, 291, 289, 201, 107, 322,
	118, 374, 111, 71, 113, 196, 262, 179, 181, 178,
	413, 188, 189, 273, 120, 227, 229, 76, 83, 200,
	199, 117, 164, 169, 118, 200, 199, 131, 71, 183,
	359, 360, 361, 185, 288, 201, 186, 71, 352, 354,
	173, 201, 210, 209, 218, 219, 212, 213, 214, 215,
	216, 217, 211, 252, 279, 281, 282, 257, 195, 280,
	76, 128, 197, 131, 131, 255, 71, 35, 354, 258,
	86, 249, 76, 239, 240, 98, 196, 242, 99, 275,
	245, 465, 262, 262, 276, 277, 76, 228, 197, 241,
	191, 262, 310, 262, 339, 262, 263, 248, 264, 187,
	272, 131, 383, 262, 168, 18, 283, 110, 85, 262,
	192, 296, 308, 309, 311, 55, 377, 105, 105, 314,
	425, 310, 315, 318, 131, 271, 426, 423, 75, 328,
	339, 118, 131, 131, 275, 326, 286, 270, 329, 105,
	316, 319, 312, 313, 118, 325, 191, 71, 287, 247,
	434, 377, 256, 336, 210, 209, 218, 219, 212, 213,
	214, 215, 216, 217, 211, 77, 118, 115, 95, 357,
	402, 437, 131, 131, 400, 403, 362, 334, 436, 401,
	335, 404, 127, 347, 348, 363, 399, 146, 145, 147,
	148, 149, 150, 46, 170, 151, 398, 163, 162, 369,
	66, 102, 371, 443, 271, 18, 101, 375, 455, 79,
	382, 384, 422, 65, 174, 379, 270, 373, 370, 232,
	456, 381, 286, 114, 235, 236, 237, 238, 356, 324,
	62, 63, 68, 395, 287, 397, 405, 331, 393, 414,
	412, 394, 161, 396, 244, 415, 410, 392, 332, 160,
	131, 54, 251, 418, 338, 131, 168, 380, 92, 72,
	254, 462, 453, 18, 424, 35, 37, 390, 1, 355,
	414, 350, 271, 271, 271, 271, 193, 432, 175, 127,
	431, 44, 430, 433, 270, 270, 270, 270, 36, 259,
	284, 177, 159, 293, 294, 295, 327, 297, 298, 299,
	300, 301, 302, 303, 304, 305, 306, 307, 38, 39,
	40, 41, 446, 447, 444, 246, 454, 427, 386, 448,
	449, 141, 391, 337, 372, 131, 243, 127, 127, 320,
	142, 458, 458, 458, 75, 459, 460, 133, 378, 96,
	464, 461, 466, 467, 468, 323, 469, 202, 171, 470,
	129, 406, 342, 340, 268, 190, 380, 209, 218, 219,
	212, 213, 214, 215, 216, 217, 211, 131, 131, 124,
	94, 450, 451, 452, 254, 64, 34, 67, 364, 365,
	366, 16, 15, 14, 13, 17, 12, 416, 70, 343,
	346, 347, 348, 344, 11, 345, 349, 70, 368, 435,
	10, 70, 9, 8, 7, 127, 210, 209, 218, 219,
	212, 213, 214, 215, 216, 217, 211, 6, 5, 385,
	388, 70, 4, 2, 0, 0, 70, 0, 109, 0,
	70, 0, 70, 0, 0, 116, 0, 0, 0, 0,
	0, 119, 70, 125, 0, 317, 0, 144, 0, 0,
	70, 0, 166, 0, 0, 417, 0, 0, 0, 0,
	0, 70, 420, 0, 70, 0, 0, 0, 254, 0,
	0, 118, 367, 262, 128, 146, 145, 147, 148, 149,
	150, 254, 0, 151, 157, 158, 0, 0, 130, 0,
	156, 210, 209, 218, 219, 212, 213, 214, 215, 216,
	217, 211, 0, 0, 441, 0, 0, 0, 70, 442,
	136, 137, 123, 445, 388, 155, 0, 138, 0, 139,
	0, 0, 343, 346, 347, 348, 344, 144, 345, 349,
	0, 0, 0, 152, 0, 0, 0, 0, 0, 76,
	125, 70, 153, 154, 0, 0, 0, 274, 0, 0,
	0, 118, 0, 262, 128, 146, 145, 147, 148, 149,
	150, 0, 0, 151, 157, 158, 0, 0, 130, 0,
	156, 210, 209, 218, 219, 212, 213, 214, 215, 216,
	217, 211, 0, 0, 0, 0, 0, 0, 125, 125,
	136, 137, 123, 0, 0, 155, 0, 138, 0, 139,
	144, 0, 333, 0, 0, 70, 0, 0, 70, 0,
	0, 0, 0, 152, 0, 0, 0, 0, 353, 0,
	70, 0, 153, 154, 118, 0, 0, 128, 146, 145,
	147, 148, 149, 150, 0, 0, 151, 157, 158, 0,
	0, 130, 0, 156, 210, 209, 218, 219, 212, 213,
	214, 215, 216, 217, 211, 0, 0, 0, 0, 0,
	0, 0, 0, 136, 137, 123, 125, 18, 155, 0,
	138, 0, 139, 218, 219, 212, 213, 214, 215, 216,
	217, 211, 144, 0, 0, 0, 152, 0, 70, 70,
	70, 70, 0, 0, 0, 153, 154, 0, 0, 0,
	0, 353, 0, 0, 411, 0, 118, 0, 0, 128,
	146, 145, 147, 148, 149, 150, 0, 0, 151, 157,
	158, 0, 0, 130, 0, 156, 0, 0, 0, 144,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 136, 137, 0, 0, 0,
	155, 0, 138, 118, 139, 0, 128, 146, 145, 147,
	148, 149, 150, 0, 0, 151, 157, 158, 152, 0,
	130, 0, 156, 0, 0, 0, 0, 153, 154, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 136, 137, 0, 0, 0, 155, 0, 138,
	118, 139, 0, 128, 146, 145, 147, 148, 149, 150,
	18, 0, 151, 157, 158, 152, 0, 0, 0, 156,
	0, 0, 0, 0, 153, 154, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 136,
	137, 0, 0, 0, 155, 0, 138, 0, 139, 118,
	0, 0, 128, 146, 145, 147, 148, 149, 150, 0,
	0, 151, 152, 0, 0, 0, 0, 0, 156, 0,
	0, 153, 154, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 136, 137,
	0, 0, 0, 155, 0, 138, 118, 139, 0, 128,
	146, 145, 147, 148, 149, 150, 0, 0, 151, 0,
	0, 152, 0, 0, 0, 156, 18, 19, 20, 21,
	153, 154, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 136, 137, 0, 22, 0,
	155, 0, 138, 0, 139, 0, 0, 0, 0, 0,
	0, 0, 30, 0, 0, 0, 0, 0, 152, 0,
	0, 0, 204, 207, 0, 0, 0, 153, 154, 220,
	221, 222, 223, 224, 225, 226, 208, 205, 206, 203,
	210, 209, 218, 219, 212, 213, 214, 215, 216, 217,
	211, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 23, 24, 26,
	25, 27, 0, 0, 0, 0, 0, 0, 0, 0,
	31, 32, 33, 28, 29,
}
var yyPact = [...]int{

	1030, -1000, -1000, 380, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -67, -59, -40, -71, -41, -1000, 356,
	187, -70, -1000, -1000, 377, 332, 301, -1000, -59, 100,
	369, 89, -68, -68, -44, -1000, -42, -1000, 100, -69,
	180, -69, 100, -1000, -86, -1000, -1000, 368, -66, -1000,
	-1000, -1000, -1000, -1000, 253, 144, -1000, 56, 302, 293,
	-6, -1000, 100, 191, -1000, 30, -1000, 100, 36, 100,
	179, 100, -51, 100, 322, 243, 100, -1000, 206, -1000,
	-1000, -1000, 100, 100, 699, -1000, 352, -1000, 288, 287,
	-1000, 100, 89, 100, 365, 89, 971, 206, 313, -1000,
	-79, 29, 100, -1000, -1000, 100, -1000, 171, -1000, -1000,
	-1000, 220, -1000, -1000, 158, -15, 85, 1020, -1000, -1000,
	828, 781, -1000, -25, -1000, -1000, 971, 971, 971, 971,
	206, 206, -1000, -1000, 206, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 971, -1000, -1000, 100,
	-1000, -1000, -1000, -1000, 241, 213, -1000, 358, 828, -1000,
	684, -16, 924, -1000, -1000, 228, 89, -1000, -54, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 156, -1000, -1000,
	365, 699, 75, -1000, -1000, 132, -1000, -1000, 47, 828,
	828, 117, 875, 99, 37, 971, 971, 971, 117, 971,
	971, 971, 971, 971, 971, 971, 971, 971, 971, 971,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 20, 1020, 79,
	182, 166, 1020, 258, 258, -1000, -1000, -1000, 611, 546,
	626, -1000, 377, 46, 684, -1000, 319, 89, 89, 358,
	341, 353, 85, 133, 684, -1000, 100, -1000, -1000, 100,
	-1000, -1000, -1000, 362, -1000, 204, 608, -1000, -1000, 138,
	328, 219, -1000, -1000, -29, -1000, 20, 33, -1000, -1000,
	93, -1000, -1000, -1000, 684, -1000, 924, -1000, -1000, 99,
	971, 971, 971, 684, 684, 531, -1000, 711, 396, -1000,
	-14, -14, 0, 0, 0, 0, -3, -3, -1000, -1000,
	971, -1000, -1000, -1000, -1000, -1000, 164, 699, -1000, 164,
	54, -1000, 828, 227, 206, 380, 192, 176, -1000, 341,
	-1000, 971, 971, -31, 206, -1000, -1000, 354, 343, 75,
	75, 75, 75, -1000, 282, 272, -1000, 260, 256, 267,
	-9, -1000, 109, -1000, -1000, 100, -1000, 168, 44, -1000,
	-1000, -1000, 166, -1000, 684, 684, 446, 971, 684, -1000,
	164, -1000, 40, -1000, 971, 41, -1000, 307, 201, -1000,
	971, -1000, -1000, 89, -1000, 194, 200, -1000, 82, 89,
	-1000, 358, 828, 971, 608, 226, 475, -1000, -1000, -1000,
	-1000, 264, -1000, 257, -1000, -1000, -1000, -45, -46, -48,
	-1000, -1000, -1000, -1000, -1000, -1000, 971, 684, -1000, -1000,
	684, 971, 297, 206, -1000, 971, 971, -1000, -1000, -1000,
	-1000, 341, 85, 195, 828, 828, -1000, -1000, 206, 206,
	206, 684, 684, 374, -1000, 684, -1000, 311, 85, 85,
	89, 89, 89, 89, -1000, 373, 12, 155, -1000, 155,
	155, 191, -1000, 89, -1000, 89, -1000, -1000, 89, -1000,
	-1000,
}
var yyPgo = [...]int{

	0, 543, 18, 542, 538, 537, 524, 523, 522, 520,
	514, 506, 505, 504, 503, 502, 501, 408, 497, 496,
	495, 490, 20, 26, 489, 475, 17, 21, 12, 474,
	473, 6, 472, 40, 471, 3, 25, 4, 470, 23,
	467, 465, 28, 207, 459, 22, 15, 9, 458, 10,
	64, 457, 450, 449, 16, 446, 444, 443, 442, 441,
	8, 438, 5, 437, 1, 436, 435, 416, 14, 7,
	37, 412, 313, 138, 411, 409, 401, 398, 0, 24,
	396, 468, 13, 391, 389, 11, 285, 388, 52, 2,
	386,
}
var yyR1 = [...]int{

	0, 87, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 2, 2, 2,
	3, 3, 4, 5, 6, 7, 7, 7, 8, 8,
	9, 10, 10, 10, 11, 13, 14, 15, 16, 16,
	16, 16, 12, 12, 90, 17, 18, 18, 19, 19,
	19, 20, 20, 21, 21, 22, 22, 23, 23, 23,
	23, 24, 24, 80, 80, 80, 79, 79, 25, 25,
	26, 26, 27, 27, 28, 28, 28, 29, 29, 29,
	29, 84, 84, 83, 83, 83, 82, 82, 30, 30,
	30, 30, 31, 31, 31, 31, 32, 32, 33, 33,
	34, 34, 34, 34, 35, 35, 36, 36, 37, 37,
	37, 37, 37, 37, 39, 39, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 45,
	45, 45, 45, 45, 45, 40, 40, 40, 40, 40,
	40, 40, 46, 46, 46, 50, 47, 47, 43, 43,
	43, 43, 43, 43, 43, 43, 43, 43, 43, 43,
	43, 43, 43, 43, 43, 43, 43, 43, 43, 43,
	43, 43, 43, 43, 43, 59, 59, 59, 59, 52,
	55, 55, 53, 53, 54, 56, 56, 51, 51, 51,
	42, 42, 42, 42, 42, 42, 42, 44, 44, 44,
	57, 57, 58, 58, 60, 60, 61, 61, 62, 63,
	63, 63, 64, 64, 64, 65, 65, 65, 66, 66,
	67, 67, 68, 68, 41, 41, 48, 48, 49, 69,
	69, 70, 71, 71, 73, 73, 86, 86, 72, 72,
	74, 74, 74, 74, 74, 74, 75, 75, 76, 76,
	77, 77, 78, 81, 88, 89, 85,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 12, 6, 3,
	8, 8, 8, 7, 3, 6, 4, 9, 6, 7,
	5, 4, 5, 4, 3, 2, 7, 3, 3, 3,
	5, 5, 2, 2, 0, 2, 0, 2, 1, 2,
	2, 0, 1, 0, 1, 1, 3, 1, 2, 3,
	5, 1, 1, 0, 1, 2, 1, 1, 0, 2,
	1, 3, 1, 1, 3, 3, 3, 3, 5, 5,
	3, 0, 1, 0, 1, 2, 1, 1, 1, 2,
	2, 1, 2, 3, 2, 3, 2, 2, 1, 3,
	0, 5, 5, 5, 1, 3, 0, 2, 1, 3,
	3, 2, 3, 3, 1, 1, 1, 3, 3, 3,
	4, 3, 4, 3, 4, 5, 6, 3, 2, 1,
	2, 1, 2, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 1, 3, 1, 3, 1, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 3, 3,
	4, 5, 3, 4, 1, 1, 1, 1, 1, 5,
	0, 1, 1, 2, 4, 0, 2, 1, 3, 5,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	0, 3, 0, 2, 0, 3, 1, 3, 2, 0,
	1, 1, 0, 2, 4, 0, 2, 4, 0, 3,
	1, 3, 0, 5, 2, 1, 1, 3, 3, 1,
	3, 3, 1, 1, 0, 2, 0, 3, 0, 1,
	1, 1, 1, 1, 1, 1, 0, 1, 0, 1,
	0, 2, 1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -87, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -13, -14, -15, -16, -12, 6, 7,
	8, 9, 28, 97, 98, 100, 99, 101, 113, 114,
	42, 110, 111, 112, -19, 5, -17, -90, -17, -17,
	-17, -17, 102, 117, -76, 108, -72, 106, 102, 102,
	103, 117, 102, -85, 15, 48, 118, 119, 97, -85,
	-85, -2, 18, 19, -20, 32, 19, -18, -72, -33,
	-81, 48, 10, -69, -70, -78, 48, -86, 107, -86,
	103, 102, -33, -73, 107, 48, -73, -33, 115, -85,
	-85, -85, 10, 102, -21, 35, -44, -78, 51, 54,
	56, 24, 28, 94, -33, 46, 69, -33, 62, -81,
	48, -33, 105, -33, 21, 44, -81, -88, 45, -81,
	-33, -22, -23, 86, -24, -81, -37, -43, 48, -38,
	62, -88, -42, -51, -49, -50, 84, 85, 91, 93,
	-78, -59, -52, -39, 21, 50, 49, 51, 52, 53,
	54, 57, 107, 116, 117, 89, 64, 58, 59, -71,
	17, 10, 30, 30, -33, -69, -81, -36, 11, -70,
	-43, -81, -88, -88, 21, -77, 109, -74, 100, 98,
	27, 99, 14, 120, 48, -33, -33, 48, -85, -85,
	-25, 46, 10, -80, -79, 20, -78, 50, 94, 61,
	60, 76, -40, 79, 62, 77, 78, 63, 76, 81,
	80, 90, 84, 85, 86, 87, 88, 89, 82, 83,
	69, 70, 71, 72, 73, 74, 75, -37, -43, -37,
	-2, -47, -43, 95, 96, -43, -43, -43, -43, -88,
	-88, -50, -88, -55, -43, -33, -66, 28, -88, -36,
	-60, 14, -37, 94, -43, -85, 44, -78, -85, -75,
	105, -89, 47, -36, -23, -26, -27, -28, -29, -33,
	-50, -88, -79, 86, -81, -78, -37, -37, -45, 57,
	62, 58, 59, -39, -43, -46, -88, -50, 55, 79,
	77, 78, 63, -43, -43, -43, -45, -43, -43, -43,
	-43, -43, -43, -43, -43, -43, -43, -43, -89, -89,
	46, -89, -42, -42, -78, -89, -22, 19, -89, -22,
	-53, -54, 65, -41, 30, -2, -69, -67, -78, -60,
	-64, 16, 15, -81, -33, -33, -85, -57, 12, 46,
	-30, -31, -32, 34, 38, 40, 35, 36, 37, 41,
	-83, -82, 20, -81, 50, -84, 20, -26, 94, 57,
	58, 59, -47, -46, -43, -43, -43, 61, -43, -89,
	-22, -89, -56, -54, 67, -37, -68, 44, -48, -49,
	-88, -68, -89, 46, -64, -43, -61, -62, -43, 94,
	-88, -58, 13, 15, -27, -28, -27, -28, 34, 34,
	34, 39, 34, 39, 34, -31, -34, 42, 106, 43,
	-82, -81, -89, 86, -78, -89, 61, -43, -89, 68,
	-43, 66, 25, 46, -78, 46, 46, -63, 22, 23,
	-85, -60, -37, -47, 44, 44, 34, 34, 103, 103,
	103, -43, -43, 26, -49, -43, -62, -64, -37, -37,
	-88, -88, -88, 8, -65, 17, 29, -35, -78, -35,
	-35, -69, 8, 79, -89, 46, -89, -89, -78, -78,
	-78,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 44, 44,
	44, 44, 44, 248, 238, 0, 0, 0, 256, 0,
	0, 0, 256, 256, 0, 48, 51, 46, 238, 0,
	0, 0, 236, 236, 0, 249, 0, 239, 0, 234,
	0, 234, 0, 35, 0, 256, 256, 256, 0, 42,
	43, 19, 49, 50, 53, 0, 52, 45, 0, 0,
	98, 253, 0, 24, 229, 0, 252, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 34, 0, 37,
	38, 39, 0, 0, 0, 54, 0, 197, 0, 0,
	47, 0, 0, 0, 106, 0, 0, 0, 0, 26,
	250, 0, 0, 31, 235, 0, 33, 0, 254, 256,
	256, 68, 55, 57, 63, 0, 61, 62, -2, 108,
	0, 0, 148, 149, 150, 151, 0, 0, 0, 0,
	187, 0, 174, 116, 0, 190, 191, 192, 193, 194,
	195, 196, 175, 176, 177, 178, 180, 114, 115, 0,
	232, 233, 198, 199, 218, 106, 99, 204, 0, 230,
	231, 0, 0, 256, 237, 0, 0, 256, 246, 240,
	241, 242, 243, 244, 245, 30, 32, 0, 40, 41,
	106, 0, 0, 58, 64, 0, 66, 67, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	135, 136, 137, 138, 139, 140, 141, 111, 0, 0,
	0, 0, 146, 0, 0, 165, 166, 167, 0, 0,
	0, 128, 0, 0, 181, 18, 0, 0, 0, 204,
	212, 0, 107, 0, 146, 25, 0, 251, 28, 0,
	247, 256, 255, 200, 56, 69, 70, 72, 73, 83,
	81, 0, 65, 59, 0, 188, 109, 110, 113, 129,
	0, 131, 133, 117, 118, 119, 0, 143, 144, 0,
	0, 0, 0, 121, 123, 0, 127, 152, 153, 154,
	155, 156, 157, 158, 159, 160, 161, 162, 112, 145,
	0, 228, 163, 164, 168, 169, 0, 0, 172, 0,
	185, 182, 0, 222, 0, 225, 222, 0, 220, 212,
	23, 0, 0, 0, 0, 29, 36, 202, 0, 0,
	0, 0, 0, 88, 0, 0, 91, 0, 0, 0,
	100, 84, 0, 86, 87, 0, 82, 0, 0, 130,
	132, 134, 0, 120, 122, 124, 0, 0, 147, 170,
	0, 173, 0, 183, 0, 0, 20, 0, 224, 226,
	0, 21, 219, 0, 22, 213, 205, 206, 209, 0,
	256, 204, 0, 0, 71, 77, 0, 80, 89, 90,
	92, 0, 94, 0, 96, 97, 74, 0, 0, 0,
	85, 75, 76, 60, 189, 142, 0, 125, 171, 179,
	186, 0, 0, 0, 221, 0, 0, 208, 210, 211,
	27, 212, 203, 201, 0, 0, 93, 95, 0, 0,
	0, 126, 184, 0, 227, 214, 207, 215, 78, 79,
	0, 0, 0, 0, 17, 0, 0, 0, 104, 0,
	0, 223, 216, 0, 101, 0, 102, 103, 0, 105,
	217,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 88, 81, 3,
	45, 47, 86, 84, 46, 85, 94, 87, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	70, 69, 71, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 90, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 80, 3, 91,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 72, 73, 74, 75, 76, 77,
	78, 79, 82, 83, 89, 92, 93, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 113, 114, 115, 116, 117,
	118, 119, 120,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:181
		{
			setParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:187
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 17:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:207
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(WhereStr, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(HavingStr, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 18:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:211
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{Expr: yyDollar[4].valExpr}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:215
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 20:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:221
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 21:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:225
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, updateList := range yyDollar[7].updateExprs {
				cols = append(cols, updateList.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:237
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 23:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:243
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:249
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 25:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:255
		{
			yyVAL.statement = &DDL{Action: CreateTableStr, IfNotExists: bool(yyDollar[3].boolVal), Table: yyDollar[4].tableName, NewName: yyDollar[4].tableName}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:260
		{
			yyVAL.statement = &DDL{Action: CreateDBStr, IfNotExists: bool(yyDollar[3].boolVal), Database: yyDollar[4].tableIdent}
		}
	case 27:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line sql.y:264
		{
			yyVAL.statement = &DDL{Action: CreateIndexStr, Table: yyDollar[7].tableName, NewName: yyDollar[7].tableName}
		}
	case 28:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:270
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableName, NewName: yyDollar[4].tableName}
		}
	case 29:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:274
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableName, NewName: yyDollar[7].tableName}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:280
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableName, NewName: yyDollar[5].tableName}
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:286
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropTableStr, Table: yyDollar[4].tableName, IfExists: exists}
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:294
		{
			yyVAL.statement = &DDL{Action: DropIndexStr, Table: yyDollar[5].tableName, NewName: yyDollar[5].tableName}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:298
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropDBStr, Database: yyDollar[4].tableIdent, IfExists: exists}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:309
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableName, NewName: yyDollar[3].tableName}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:315
		{
			yyVAL.statement = &Xa{}
		}
	case 36:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:321
		{
			yyVAL.statement = &Shard{ShardKey: string(yyDollar[5].bytes)}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:327
		{
			yyVAL.statement = &UseDB{Database: string(yyDollar[2].bytes)}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:333
		{
			yyVAL.statement = &ShowDatabases{}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:337
		{
			yyVAL.statement = &ShowTables{}
		}
	case 40:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:341
		{
			yyVAL.statement = &ShowTables{Database: yyDollar[4].tableIdent}
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:345
		{
			yyVAL.statement = &ShowCreateTable{Table: yyDollar[4].tableName}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:351
		{
			yyVAL.statement = &Other{}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:355
		{
			yyVAL.statement = &Other{}
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:360
		{
			setAllowComments(yylex, true)
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:364
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:370
		{
			yyVAL.bytes2 = nil
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:374
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:380
		{
			yyVAL.str = UnionStr
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:384
		{
			yyVAL.str = UnionAllStr
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:388
		{
			yyVAL.str = UnionDistinctStr
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:393
		{
			yyVAL.str = ""
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:397
		{
			yyVAL.str = DistinctStr
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:402
		{
			yyVAL.str = ""
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:406
		{
			yyVAL.str = StraightJoinHint
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:412
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:416
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:422
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:426
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:430
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Name: yyDollar[1].tableIdent}}
		}
	case 60:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:434
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}}
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:440
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:444
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:449
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:453
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:457
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:464
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 68:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:469
		{
			yyVAL.tableExprs = TableExprs{&AliasedTableExpr{Expr: &TableName{Name: NewTableIdent("dual")}}}
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:473
		{
			yyVAL.tableExprs = yyDollar[2].tableExprs
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:479
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:483
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:493
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:497
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:501
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:514
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:518
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 79:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:522
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:526
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:531
		{
			yyVAL.empty = struct{}{}
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:533
		{
			yyVAL.empty = struct{}{}
		}
	case 83:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:536
		{
			yyVAL.tableIdent = NewTableIdent("")
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:540
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:544
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:551
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:557
		{
			yyVAL.str = JoinStr
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:561
		{
			yyVAL.str = JoinStr
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:565
		{
			yyVAL.str = JoinStr
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:569
		{
			yyVAL.str = StraightJoinStr
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:575
		{
			yyVAL.str = LeftJoinStr
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:579
		{
			yyVAL.str = LeftJoinStr
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:583
		{
			yyVAL.str = RightJoinStr
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:587
		{
			yyVAL.str = RightJoinStr
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:593
		{
			yyVAL.str = NaturalJoinStr
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:597
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:607
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:611
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 100:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:616
		{
			yyVAL.indexHints = nil
		}
	case 101:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:620
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 102:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:624
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 103:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:628
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:634
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:638
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:643
		{
			yyVAL.boolExpr = nil
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:647
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:654
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:658
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:662
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:666
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:670
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:676
		{
			yyVAL.boolVal = BoolVal(true)
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:680
		{
			yyVAL.boolVal = BoolVal(false)
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:686
		{
			yyVAL.boolExpr = yyDollar[1].boolVal
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:690
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].boolVal}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:694
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:698
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:702
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:706
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:710
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:714
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:718
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 125:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:722
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 126:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:726
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:730
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:734
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:740
		{
			yyVAL.str = IsNullStr
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:744
		{
			yyVAL.str = IsNotNullStr
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:748
		{
			yyVAL.str = IsTrueStr
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:752
		{
			yyVAL.str = IsNotTrueStr
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:756
		{
			yyVAL.str = IsFalseStr
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:760
		{
			yyVAL.str = IsNotFalseStr
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:766
		{
			yyVAL.str = EqualStr
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:770
		{
			yyVAL.str = LessThanStr
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:774
		{
			yyVAL.str = GreaterThanStr
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:778
		{
			yyVAL.str = LessEqualStr
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:782
		{
			yyVAL.str = GreaterEqualStr
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:786
		{
			yyVAL.str = NotEqualStr
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:790
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:796
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:800
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:804
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:810
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:816
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:820
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:826
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:830
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:834
		{
			yyVAL.valExpr = yyDollar[1].valTuple
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:838
		{
			yyVAL.valExpr = yyDollar[1].subquery
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:842
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:846
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:850
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:854
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:858
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:862
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:866
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:870
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:874
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:878
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:882
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:886
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONExtractOp, Right: yyDollar[3].valExpr}
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:890
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONUnquoteExtractOp, Right: yyDollar[3].valExpr}
		}
	case 165:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:894
		{
			if num, ok := yyDollar[2].valExpr.(*SQLVal); ok && num.Type == IntVal {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:902
		{
			if num, ok := yyDollar[2].valExpr.(*SQLVal); ok && num.Type == IntVal {
				// Handle double negative
				if num.Val[0] == '-' {
					num.Val = num.Val[1:]
					yyVAL.valExpr = num
				} else {
					yyVAL.valExpr = NewIntVal(append([]byte("-"), num.Val...))
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UMinusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:916
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:920
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.valExpr = &IntervalExpr{Expr: yyDollar[2].valExpr, Unit: yyDollar[3].colIdent}
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:928
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 170:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:932
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 171:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:936
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:940
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 173:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:944
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:948
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:954
		{
			yyVAL.colIdent = NewColIdent("if")
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:958
		{
			yyVAL.colIdent = NewColIdent("current_timestamp")
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:962
		{
			yyVAL.colIdent = NewColIdent("database")
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:966
		{
			yyVAL.colIdent = NewColIdent("mod")
		}
	case 179:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:972
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:977
		{
			yyVAL.valExpr = nil
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:981
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:987
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:991
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 184:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:997
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1002
		{
			yyVAL.valExpr = nil
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1006
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1012
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 189:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1026
		{
			yyVAL.valExpr = NewStrVal(yyDollar[1].bytes)
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1030
		{
			yyVAL.valExpr = NewHexVal(yyDollar[1].bytes)
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1034
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1038
		{
			yyVAL.valExpr = NewFloatVal(yyDollar[1].bytes)
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1042
		{
			yyVAL.valExpr = NewHexNum(yyDollar[1].bytes)
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1046
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1050
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1056
		{
			// TODO(sougou): Deprecate this construct.
			if yyDollar[1].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.valExpr = NewIntVal([]byte("1"))
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1065
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1069
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1074
		{
			yyVAL.valExprs = nil
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1078
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.boolExpr = nil
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1087
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.orderBy = nil
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1096
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1102
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1106
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1112
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.str = AscScr
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1121
		{
			yyVAL.str = AscScr
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1125
		{
			yyVAL.str = DescScr
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1130
		{
			yyVAL.limit = nil
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1134
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 214:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1138
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1143
		{
			yyVAL.str = ""
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1147
		{
			yyVAL.str = ForUpdateStr
		}
	case 217:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1151
		{
			if yyDollar[3].colIdent.Lowered() != "share" {
				yylex.Error("expecting share")
				return 1
			}
			if yyDollar[4].colIdent.Lowered() != "mode" {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = ShareModeStr
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1164
		{
			yyVAL.columns = nil
		}
	case 219:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1168
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1174
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 221:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1178
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 222:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1183
		{
			yyVAL.updateExprs = nil
		}
	case 223:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1187
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 224:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1193
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1197
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1203
		{
			yyVAL.values = Values{yyDollar[1].valTuple}
		}
	case 227:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1207
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].valTuple)
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1213
		{
			yyVAL.valTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1219
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 230:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1223
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 231:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1229
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 234:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1238
		{
			yyVAL.byt = 0
		}
	case 235:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1240
		{
			yyVAL.byt = 1
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1243
		{
			yyVAL.boolVal = false
		}
	case 237:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1245
		{
			yyVAL.boolVal = true
		}
	case 238:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1248
		{
			yyVAL.str = ""
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1250
		{
			yyVAL.str = IgnoreStr
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1254
		{
			yyVAL.empty = struct{}{}
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1256
		{
			yyVAL.empty = struct{}{}
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1258
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1260
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1262
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1264
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1267
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1269
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1272
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1274
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1277
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1279
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1283
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1289
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1295
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1304
		{
			decNesting(yylex)
		}
	case 256:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1309
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}

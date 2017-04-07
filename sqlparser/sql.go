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
const OFFSET = 57359
const FOR = 57360
const ALL = 57361
const DISTINCT = 57362
const AS = 57363
const EXISTS = 57364
const ASC = 57365
const DESC = 57366
const INTO = 57367
const DUPLICATE = 57368
const KEY = 57369
const DEFAULT = 57370
const SET = 57371
const LOCK = 57372
const VALUES = 57373
const LAST_INSERT_ID = 57374
const NEXT = 57375
const VALUE = 57376
const JOIN = 57377
const STRAIGHT_JOIN = 57378
const LEFT = 57379
const RIGHT = 57380
const INNER = 57381
const OUTER = 57382
const CROSS = 57383
const NATURAL = 57384
const USE = 57385
const FORCE = 57386
const ON = 57387
const ID = 57388
const HEX = 57389
const STRING = 57390
const INTEGRAL = 57391
const FLOAT = 57392
const HEXNUM = 57393
const VALUE_ARG = 57394
const LIST_ARG = 57395
const COMMENT = 57396
const NULL = 57397
const TRUE = 57398
const FALSE = 57399
const OR = 57400
const AND = 57401
const NOT = 57402
const BETWEEN = 57403
const CASE = 57404
const WHEN = 57405
const THEN = 57406
const ELSE = 57407
const END = 57408
const LE = 57409
const GE = 57410
const NE = 57411
const NULL_SAFE_EQUAL = 57412
const IS = 57413
const LIKE = 57414
const REGEXP = 57415
const IN = 57416
const SHIFT_LEFT = 57417
const SHIFT_RIGHT = 57418
const MOD = 57419
const UNARY = 57420
const INTERVAL = 57421
const JSON_EXTRACT_OP = 57422
const JSON_UNQUOTE_EXTRACT_OP = 57423
const CREATE = 57424
const ALTER = 57425
const DROP = 57426
const RENAME = 57427
const ANALYZE = 57428
const TABLE = 57429
const INDEX = 57430
const VIEW = 57431
const TO = 57432
const IGNORE = 57433
const IF = 57434
const UNIQUE = 57435
const USING = 57436
const SHOW = 57437
const DESCRIBE = 57438
const EXPLAIN = 57439
const XA = 57440
const PARTITION = 57441
const HASH = 57442
const CURRENT_TIMESTAMP = 57443
const DATABASE = 57444
const DATABASES = 57445
const TABLES = 57446
const UNUSED = 57447

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
	"OFFSET",
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
	95, 254,
	-2, 253,
}

const yyNprod = 258
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1075

var yyAct = [...]int{

	140, 387, 261, 459, 73, 231, 134, 330, 250, 351,
	172, 53, 376, 321, 135, 285, 341, 265, 230, 3,
	121, 132, 278, 143, 194, 167, 267, 122, 266, 58,
	88, 49, 50, 42, 176, 84, 407, 409, 260, 45,
	126, 78, 75, 47, 59, 60, 51, 317, 43, 144,
	56, 57, 112, 61, 74, 441, 440, 439, 80, 93,
	81, 52, 48, 233, 234, 389, 97, 89, 90, 91,
	358, 253, 198, 118, 103, 262, 128, 146, 145, 147,
	148, 149, 150, 211, 76, 151, 157, 158, 128, 465,
	130, 201, 156, 214, 215, 216, 217, 211, 106, 117,
	408, 182, 419, 75, 322, 131, 75, 165, 200, 199,
	118, 108, 136, 137, 123, 180, 76, 155, 173, 138,
	288, 139, 413, 292, 201, 196, 273, 200, 199, 262,
	100, 188, 189, 421, 199, 152, 184, 290, 291, 289,
	71, 131, 131, 201, 153, 154, 322, 105, 374, 201,
	352, 239, 240, 195, 83, 242, 359, 360, 361, 241,
	169, 209, 218, 219, 212, 213, 214, 215, 216, 217,
	211, 227, 229, 128, 71, 248, 354, 257, 71, 131,
	354, 76, 35, 197, 262, 255, 179, 181, 178, 258,
	76, 249, 197, 467, 262, 18, 196, 200, 199, 275,
	191, 262, 131, 271, 228, 187, 86, 270, 183, 252,
	131, 131, 118, 201, 286, 71, 263, 110, 287, 264,
	272, 85, 279, 281, 282, 262, 283, 280, 55, 310,
	262, 296, 308, 309, 311, 118, 118, 192, 71, 314,
	276, 277, 315, 318, 339, 262, 383, 262, 75, 328,
	131, 131, 326, 310, 275, 312, 313, 377, 329, 105,
	316, 319, 168, 427, 423, 325, 212, 213, 214, 215,
	216, 217, 211, 336, 191, 146, 145, 147, 148, 149,
	150, 76, 271, 151, 98, 339, 270, 99, 435, 357,
	377, 256, 362, 115, 95, 438, 77, 247, 105, 127,
	286, 437, 399, 163, 287, 363, 343, 346, 347, 348,
	344, 170, 345, 349, 118, 402, 436, 398, 400, 369,
	403, 18, 371, 401, 404, 46, 347, 348, 131, 162,
	382, 379, 102, 131, 373, 380, 232, 384, 370, 381,
	79, 235, 236, 237, 238, 390, 324, 444, 422, 101,
	271, 271, 271, 271, 270, 270, 270, 270, 66, 414,
	412, 244, 410, 375, 68, 415, 405, 395, 394, 397,
	396, 65, 36, 418, 174, 457, 114, 254, 331, 343,
	346, 347, 348, 344, 424, 345, 349, 458, 356, 393,
	414, 332, 38, 39, 40, 41, 127, 62, 63, 434,
	432, 161, 431, 131, 54, 338, 251, 284, 392, 160,
	293, 294, 295, 168, 297, 298, 299, 300, 301, 302,
	303, 304, 305, 306, 307, 18, 92, 72, 464, 448,
	445, 455, 35, 433, 380, 37, 1, 355, 350, 269,
	449, 193, 175, 44, 127, 127, 131, 131, 259, 177,
	452, 453, 454, 460, 460, 460, 75, 461, 462, 159,
	463, 327, 466, 246, 468, 469, 470, 171, 471, 456,
	428, 472, 386, 141, 391, 337, 450, 451, 372, 69,
	243, 320, 142, 133, 378, 96, 323, 202, 82, 129,
	406, 254, 87, 342, 340, 364, 365, 366, 218, 219,
	212, 213, 214, 215, 216, 217, 211, 70, 268, 190,
	124, 94, 104, 64, 34, 368, 70, 107, 67, 16,
	70, 111, 127, 113, 15, 14, 13, 17, 429, 430,
	12, 11, 10, 120, 9, 8, 385, 388, 7, 6,
	70, 164, 5, 4, 2, 70, 0, 109, 0, 70,
	0, 70, 185, 0, 116, 186, 0, 0, 0, 0,
	119, 70, 125, 0, 0, 0, 0, 0, 0, 70,
	0, 166, 417, 0, 0, 0, 144, 0, 0, 420,
	70, 0, 0, 70, 0, 254, 210, 209, 218, 219,
	212, 213, 214, 215, 216, 217, 211, 0, 254, 245,
	118, 0, 262, 128, 146, 145, 147, 148, 149, 150,
	0, 0, 151, 157, 158, 0, 0, 130, 0, 156,
	0, 442, 0, 0, 0, 0, 443, 70, 0, 0,
	446, 447, 388, 0, 0, 0, 0, 0, 0, 136,
	137, 123, 0, 0, 155, 0, 138, 0, 139, 0,
	144, 0, 0, 0, 0, 0, 0, 0, 0, 125,
	70, 0, 152, 0, 0, 0, 274, 0, 0, 0,
	0, 153, 154, 0, 118, 0, 0, 128, 146, 145,
	147, 148, 149, 150, 0, 0, 151, 157, 158, 0,
	0, 130, 0, 156, 0, 0, 334, 0, 0, 335,
	0, 0, 0, 0, 0, 0, 0, 125, 125, 0,
	18, 0, 0, 136, 137, 123, 0, 0, 155, 0,
	138, 333, 139, 0, 70, 0, 144, 70, 0, 0,
	0, 0, 0, 0, 0, 0, 152, 353, 0, 70,
	0, 0, 0, 0, 0, 153, 154, 0, 0, 0,
	118, 416, 0, 128, 146, 145, 147, 148, 149, 150,
	0, 0, 151, 157, 158, 0, 0, 130, 0, 156,
	210, 209, 218, 219, 212, 213, 214, 215, 216, 217,
	211, 0, 0, 0, 426, 125, 0, 0, 0, 136,
	137, 0, 0, 0, 155, 0, 138, 0, 139, 0,
	144, 0, 0, 0, 0, 0, 0, 70, 70, 70,
	70, 0, 152, 0, 425, 0, 0, 0, 0, 0,
	353, 153, 154, 411, 118, 0, 0, 128, 146, 145,
	147, 148, 149, 150, 0, 0, 151, 157, 158, 0,
	0, 130, 0, 156, 18, 19, 20, 21, 210, 209,
	218, 219, 212, 213, 214, 215, 216, 217, 211, 0,
	0, 0, 0, 136, 137, 76, 0, 22, 155, 0,
	138, 118, 139, 0, 128, 146, 145, 147, 148, 149,
	150, 30, 18, 151, 157, 158, 152, 0, 0, 0,
	156, 0, 0, 0, 0, 153, 154, 210, 209, 218,
	219, 212, 213, 214, 215, 216, 217, 211, 0, 0,
	136, 137, 0, 0, 0, 155, 0, 138, 0, 139,
	0, 0, 118, 0, 0, 128, 146, 145, 147, 148,
	149, 150, 0, 152, 151, 0, 23, 24, 26, 25,
	27, 156, 153, 154, 0, 0, 0, 0, 0, 31,
	32, 33, 28, 29, 0, 0, 0, 0, 0, 0,
	0, 136, 137, 0, 0, 0, 155, 0, 138, 118,
	139, 0, 128, 146, 145, 147, 148, 149, 150, 0,
	0, 151, 0, 0, 152, 0, 0, 0, 156, 0,
	0, 0, 0, 153, 154, 0, 0, 0, 0, 367,
	0, 0, 0, 0, 0, 0, 0, 0, 136, 137,
	0, 0, 0, 155, 0, 138, 0, 139, 210, 209,
	218, 219, 212, 213, 214, 215, 216, 217, 211, 0,
	0, 152, 0, 0, 0, 204, 207, 0, 0, 0,
	153, 154, 220, 221, 222, 223, 224, 225, 226, 208,
	205, 206, 203, 210, 209, 218, 219, 212, 213, 214,
	215, 216, 217, 211, 210, 209, 218, 219, 212, 213,
	214, 215, 216, 217, 211,
}
var yyPact = [...]int{

	838, -1000, -1000, 427, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -70, -64, -41, -72, -42, -1000, 389,
	179, -69, -1000, -1000, 419, 378, 338, -1000, -64, 91,
	417, 67, -67, -67, -46, -1000, -43, -1000, 91, -73,
	172, -73, 91, -1000, -86, -1000, -1000, 416, -44, -1000,
	-1000, -1000, -1000, -1000, 258, 232, -1000, 73, 324, 303,
	-21, -1000, 91, 100, -1000, 28, -1000, 91, 48, 91,
	168, 91, -54, 91, 354, 248, 91, -1000, 190, -1000,
	-1000, -1000, 91, 91, 628, -1000, 391, -1000, 298, 272,
	-1000, 91, 67, 91, 402, 67, 923, 190, 352, -1000,
	-76, 87, 91, -1000, -1000, 91, -1000, 156, -1000, -1000,
	-1000, 227, -1000, -1000, 132, -23, 47, 972, -1000, -1000,
	778, 704, -1000, -33, -1000, -1000, 923, 923, 923, 923,
	190, 190, -1000, -1000, 190, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 923, -1000, -1000, 91,
	-1000, -1000, -1000, -1000, 268, 251, -1000, 392, 778, -1000,
	983, -24, 876, -1000, -1000, 246, 67, -1000, -68, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 81, -1000, -1000,
	402, 628, 166, -1000, -1000, 141, -1000, -1000, 39, 778,
	778, 164, 825, 64, 59, 923, 923, 923, 164, 923,
	923, 923, 923, 923, 923, 923, 923, 923, 923, 923,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 14, 972, 136,
	177, 182, 972, 225, 225, -1000, -1000, -1000, 816, 27,
	554, -1000, 419, 38, 983, -1000, 315, 67, 67, 392,
	362, 376, 47, 124, 983, -1000, 91, -1000, -1000, 91,
	-1000, -1000, -1000, 393, -1000, 238, 344, -1000, -1000, 129,
	367, 189, -1000, -1000, -25, -1000, 14, 72, -1000, -1000,
	98, -1000, -1000, -1000, 983, -1000, 876, -1000, -1000, 64,
	923, 923, 923, 983, 983, 937, -1000, 415, 79, -1000,
	6, 6, -8, -8, -8, -8, 181, 181, -1000, -1000,
	923, -1000, -1000, -1000, -1000, -1000, 153, 628, -1000, 153,
	80, -1000, 778, 245, 190, 427, 212, 199, -1000, 362,
	-1000, 923, 923, -30, 190, -1000, -1000, 395, 374, 166,
	166, 166, 166, -1000, 282, 267, -1000, 283, 280, 289,
	-7, -1000, 125, -1000, -1000, 91, -1000, 197, 35, -1000,
	-1000, -1000, 182, -1000, 983, 983, 689, 923, 983, -1000,
	153, -1000, 33, -1000, 923, 66, -1000, 322, 217, -1000,
	923, -1000, -1000, 67, -1000, 767, 216, -1000, 505, 67,
	-1000, 392, 778, 923, 344, 243, 271, -1000, -1000, -1000,
	-1000, 266, -1000, 260, -1000, -1000, -1000, -47, -48, -49,
	-1000, -1000, -1000, -1000, -1000, -1000, 923, 983, -1000, -1000,
	983, 923, 320, 190, -1000, 923, 923, 923, -1000, -1000,
	-1000, -1000, 362, 47, 206, 778, 778, -1000, -1000, 190,
	190, 190, 983, 983, 423, -1000, 983, 983, -1000, 357,
	47, 47, 67, 67, 67, 67, -1000, 420, 9, 146,
	-1000, 146, 146, 100, -1000, 67, -1000, 67, -1000, -1000,
	67, -1000, -1000,
}
var yyPgo = [...]int{

	0, 544, 18, 543, 542, 539, 538, 535, 534, 532,
	531, 530, 527, 526, 525, 524, 519, 372, 518, 514,
	513, 511, 20, 27, 510, 509, 17, 28, 26, 508,
	494, 16, 493, 439, 490, 3, 25, 40, 489, 23,
	487, 486, 21, 204, 485, 22, 15, 5, 484, 6,
	14, 483, 482, 481, 13, 480, 478, 475, 474, 473,
	8, 472, 1, 470, 7, 469, 463, 461, 12, 4,
	54, 459, 325, 154, 449, 448, 443, 442, 0, 24,
	441, 467, 9, 438, 437, 11, 296, 436, 10, 2,
	435,
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
	63, 63, 64, 64, 64, 64, 65, 65, 65, 66,
	66, 67, 67, 68, 68, 41, 41, 48, 48, 49,
	69, 69, 70, 71, 71, 73, 73, 86, 86, 72,
	72, 74, 74, 74, 74, 74, 74, 75, 75, 76,
	76, 77, 77, 78, 81, 88, 89, 85,
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
	1, 1, 0, 2, 4, 4, 0, 2, 4, 0,
	3, 1, 3, 0, 5, 2, 1, 1, 3, 3,
	1, 3, 3, 1, 1, 0, 2, 0, 3, 0,
	1, 1, 1, 1, 1, 1, 1, 0, 1, 0,
	1, 0, 2, 1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -87, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -13, -14, -15, -16, -12, 6, 7,
	8, 9, 29, 98, 99, 101, 100, 102, 114, 115,
	43, 111, 112, 113, -19, 5, -17, -90, -17, -17,
	-17, -17, 103, 118, -76, 109, -72, 107, 103, 103,
	104, 118, 103, -85, 15, 49, 119, 120, 98, -85,
	-85, -2, 19, 20, -20, 33, 20, -18, -72, -33,
	-81, 49, 10, -69, -70, -78, 49, -86, 108, -86,
	104, 103, -33, -73, 108, 49, -73, -33, 116, -85,
	-85, -85, 10, 103, -21, 36, -44, -78, 52, 55,
	57, 25, 29, 95, -33, 47, 70, -33, 63, -81,
	49, -33, 106, -33, 22, 45, -81, -88, 46, -81,
	-33, -22, -23, 87, -24, -81, -37, -43, 49, -38,
	63, -88, -42, -51, -49, -50, 85, 86, 92, 94,
	-78, -59, -52, -39, 22, 51, 50, 52, 53, 54,
	55, 58, 108, 117, 118, 90, 65, 59, 60, -71,
	18, 10, 31, 31, -33, -69, -81, -36, 11, -70,
	-43, -81, -88, -88, 22, -77, 110, -74, 101, 99,
	28, 100, 14, 121, 49, -33, -33, 49, -85, -85,
	-25, 47, 10, -80, -79, 21, -78, 51, 95, 62,
	61, 77, -40, 80, 63, 78, 79, 64, 77, 82,
	81, 91, 85, 86, 87, 88, 89, 90, 83, 84,
	70, 71, 72, 73, 74, 75, 76, -37, -43, -37,
	-2, -47, -43, 96, 97, -43, -43, -43, -43, -88,
	-88, -50, -88, -55, -43, -33, -66, 29, -88, -36,
	-60, 14, -37, 95, -43, -85, 45, -78, -85, -75,
	106, -89, 48, -36, -23, -26, -27, -28, -29, -33,
	-50, -88, -79, 87, -81, -78, -37, -37, -45, 58,
	63, 59, 60, -39, -43, -46, -88, -50, 56, 80,
	78, 79, 64, -43, -43, -43, -45, -43, -43, -43,
	-43, -43, -43, -43, -43, -43, -43, -43, -89, -89,
	47, -89, -42, -42, -78, -89, -22, 20, -89, -22,
	-53, -54, 66, -41, 31, -2, -69, -67, -78, -60,
	-64, 16, 15, -81, -33, -33, -85, -57, 12, 47,
	-30, -31, -32, 35, 39, 41, 36, 37, 38, 42,
	-83, -82, 21, -81, 51, -84, 21, -26, 95, 58,
	59, 60, -47, -46, -43, -43, -43, 62, -43, -89,
	-22, -89, -56, -54, 68, -37, -68, 45, -48, -49,
	-88, -68, -89, 47, -64, -43, -61, -62, -43, 95,
	-88, -58, 13, 15, -27, -28, -27, -28, 35, 35,
	35, 40, 35, 40, 35, -31, -34, 43, 107, 44,
	-82, -81, -89, 87, -78, -89, 62, -43, -89, 69,
	-43, 67, 26, 47, -78, 47, 17, 47, -63, 23,
	24, -85, -60, -37, -47, 45, 45, 35, 35, 104,
	104, 104, -43, -43, 27, -49, -43, -43, -62, -64,
	-37, -37, -88, -88, -88, 8, -65, 18, 30, -35,
	-78, -35, -35, -69, 8, 80, -89, 47, -89, -89,
	-78, -78, -78,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 44, 44,
	44, 44, 44, 249, 239, 0, 0, 0, 257, 0,
	0, 0, 257, 257, 0, 48, 51, 46, 239, 0,
	0, 0, 237, 237, 0, 250, 0, 240, 0, 235,
	0, 235, 0, 35, 0, 257, 257, 257, 0, 42,
	43, 19, 49, 50, 53, 0, 52, 45, 0, 0,
	98, 254, 0, 24, 230, 0, 253, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 34, 0, 37,
	38, 39, 0, 0, 0, 54, 0, 197, 0, 0,
	47, 0, 0, 0, 106, 0, 0, 0, 0, 26,
	251, 0, 0, 31, 236, 0, 33, 0, 255, 257,
	257, 68, 55, 57, 63, 0, 61, 62, -2, 108,
	0, 0, 148, 149, 150, 151, 0, 0, 0, 0,
	187, 0, 174, 116, 0, 190, 191, 192, 193, 194,
	195, 196, 175, 176, 177, 178, 180, 114, 115, 0,
	233, 234, 198, 199, 219, 106, 99, 204, 0, 231,
	232, 0, 0, 257, 238, 0, 0, 257, 247, 241,
	242, 243, 244, 245, 246, 30, 32, 0, 40, 41,
	106, 0, 0, 58, 64, 0, 66, 67, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	135, 136, 137, 138, 139, 140, 141, 111, 0, 0,
	0, 0, 146, 0, 0, 165, 166, 167, 0, 0,
	0, 128, 0, 0, 181, 18, 0, 0, 0, 204,
	212, 0, 107, 0, 146, 25, 0, 252, 28, 0,
	248, 257, 256, 200, 56, 69, 70, 72, 73, 83,
	81, 0, 65, 59, 0, 188, 109, 110, 113, 129,
	0, 131, 133, 117, 118, 119, 0, 143, 144, 0,
	0, 0, 0, 121, 123, 0, 127, 152, 153, 154,
	155, 156, 157, 158, 159, 160, 161, 162, 112, 145,
	0, 229, 163, 164, 168, 169, 0, 0, 172, 0,
	185, 182, 0, 223, 0, 226, 223, 0, 221, 212,
	23, 0, 0, 0, 0, 29, 36, 202, 0, 0,
	0, 0, 0, 88, 0, 0, 91, 0, 0, 0,
	100, 84, 0, 86, 87, 0, 82, 0, 0, 130,
	132, 134, 0, 120, 122, 124, 0, 0, 147, 170,
	0, 173, 0, 183, 0, 0, 20, 0, 225, 227,
	0, 21, 220, 0, 22, 213, 205, 206, 209, 0,
	257, 204, 0, 0, 71, 77, 0, 80, 89, 90,
	92, 0, 94, 0, 96, 97, 74, 0, 0, 0,
	85, 75, 76, 60, 189, 142, 0, 125, 171, 179,
	186, 0, 0, 0, 222, 0, 0, 0, 208, 210,
	211, 27, 212, 203, 201, 0, 0, 93, 95, 0,
	0, 0, 126, 184, 0, 228, 214, 215, 207, 216,
	78, 79, 0, 0, 0, 0, 17, 0, 0, 0,
	104, 0, 0, 224, 217, 0, 101, 0, 102, 103,
	0, 105, 218,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 89, 82, 3,
	46, 48, 87, 85, 47, 86, 95, 88, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	71, 70, 72, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 91, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 81, 3, 92,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 73, 74, 75, 76, 77,
	78, 79, 80, 83, 84, 90, 93, 94, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 113, 114, 115, 116, 117,
	118, 119, 120, 121,
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
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1142
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr, Offset: yyDollar[4].valExpr}
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1148
		{
			yyVAL.str = ""
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1152
		{
			yyVAL.str = ForUpdateStr
		}
	case 218:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1156
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
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1169
		{
			yyVAL.columns = nil
		}
	case 220:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1173
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1179
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 222:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1183
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 223:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1188
		{
			yyVAL.updateExprs = nil
		}
	case 224:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1192
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 225:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1198
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1202
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1208
		{
			yyVAL.values = Values{yyDollar[1].valTuple}
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1212
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].valTuple)
		}
	case 229:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1218
		{
			yyVAL.valTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1224
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 231:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1228
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1234
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 235:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1243
		{
			yyVAL.byt = 0
		}
	case 236:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1245
		{
			yyVAL.byt = 1
		}
	case 237:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1248
		{
			yyVAL.boolVal = false
		}
	case 238:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1250
		{
			yyVAL.boolVal = true
		}
	case 239:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1253
		{
			yyVAL.str = ""
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1255
		{
			yyVAL.str = IgnoreStr
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1259
		{
			yyVAL.empty = struct{}{}
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1261
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1263
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1265
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1267
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1269
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1272
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1274
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1277
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1279
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1282
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1284
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1288
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1294
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1300
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 256:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1309
		{
			decNesting(yylex)
		}
	case 257:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1314
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}

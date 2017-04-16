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
const PROCESSLIST = 57441
const PARTITION = 57442
const HASH = 57443
const KILL = 57444
const CURRENT_TIMESTAMP = 57445
const DATABASE = 57446
const DATABASES = 57447
const TABLES = 57448
const UNUSED = 57449

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
	"PROCESSLIST",
	"PARTITION",
	"HASH",
	"KILL",
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
	-1, 135,
	95, 258,
	-2, 257,
}

const yyNprod = 262
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1047

var yyAct = [...]int{

	147, 394, 268, 466, 78, 238, 56, 141, 257, 358,
	179, 337, 292, 383, 348, 273, 272, 237, 3, 285,
	150, 201, 328, 93, 174, 274, 128, 139, 129, 45,
	61, 52, 53, 79, 183, 48, 414, 416, 89, 83,
	133, 276, 64, 65, 50, 80, 46, 62, 54, 267,
	119, 448, 142, 59, 60, 66, 447, 446, 85, 98,
	86, 55, 51, 240, 241, 94, 95, 96, 189, 99,
	100, 104, 216, 225, 226, 219, 220, 221, 222, 223,
	224, 218, 187, 125, 74, 218, 135, 153, 152, 154,
	155, 156, 157, 87, 396, 158, 365, 92, 260, 205,
	415, 110, 163, 191, 124, 81, 135, 436, 437, 472,
	80, 206, 138, 80, 172, 208, 113, 107, 329, 111,
	426, 115, 143, 144, 114, 180, 208, 162, 118, 145,
	120, 146, 203, 195, 196, 221, 222, 223, 224, 218,
	127, 125, 63, 420, 280, 159, 176, 81, 138, 138,
	171, 295, 88, 186, 188, 185, 160, 161, 246, 247,
	125, 192, 249, 76, 193, 217, 216, 225, 226, 219,
	220, 221, 222, 223, 224, 218, 269, 190, 234, 236,
	299, 329, 255, 381, 264, 20, 138, 262, 76, 207,
	206, 265, 474, 269, 297, 298, 296, 256, 135, 235,
	366, 367, 368, 203, 248, 208, 282, 91, 252, 138,
	278, 38, 433, 207, 206, 359, 259, 138, 138, 428,
	81, 293, 270, 105, 279, 125, 106, 271, 76, 208,
	290, 286, 288, 289, 202, 303, 287, 207, 206, 315,
	316, 318, 432, 76, 269, 361, 321, 283, 284, 322,
	325, 194, 277, 208, 269, 80, 335, 138, 138, 333,
	117, 282, 81, 294, 204, 336, 198, 269, 319, 320,
	76, 332, 361, 323, 326, 343, 217, 216, 225, 226,
	219, 220, 221, 222, 223, 224, 218, 90, 81, 278,
	204, 317, 269, 346, 269, 364, 423, 390, 269, 369,
	58, 134, 384, 112, 112, 341, 317, 293, 342, 370,
	434, 430, 346, 177, 82, 217, 216, 225, 226, 219,
	220, 221, 222, 223, 224, 218, 376, 125, 199, 378,
	442, 277, 175, 384, 263, 138, 254, 389, 239, 386,
	138, 122, 387, 242, 243, 244, 245, 388, 391, 294,
	380, 377, 397, 125, 49, 102, 170, 278, 278, 278,
	278, 84, 401, 251, 403, 198, 421, 419, 112, 417,
	382, 412, 422, 402, 411, 404, 354, 355, 445, 261,
	425, 225, 226, 219, 220, 221, 222, 223, 224, 218,
	409, 431, 444, 406, 405, 410, 73, 421, 134, 277,
	277, 277, 277, 407, 438, 169, 441, 439, 408, 291,
	138, 109, 300, 301, 302, 451, 304, 305, 306, 307,
	308, 309, 310, 311, 312, 313, 314, 219, 220, 221,
	222, 223, 224, 218, 429, 71, 455, 464, 452, 108,
	440, 387, 363, 20, 178, 181, 134, 134, 70, 465,
	121, 456, 258, 138, 138, 67, 68, 459, 460, 461,
	467, 467, 467, 80, 468, 469, 168, 470, 331, 473,
	338, 475, 476, 477, 167, 478, 400, 339, 479, 57,
	399, 345, 374, 457, 458, 175, 97, 75, 350, 353,
	354, 355, 351, 261, 352, 356, 75, 371, 372, 373,
	75, 217, 216, 225, 226, 219, 220, 221, 222, 223,
	224, 218, 77, 471, 462, 20, 38, 375, 40, 1,
	362, 357, 75, 200, 134, 39, 182, 75, 47, 116,
	266, 75, 184, 75, 166, 334, 123, 253, 392, 395,
	463, 435, 126, 75, 393, 148, 132, 41, 42, 43,
	44, 398, 344, 75, 379, 173, 81, 250, 153, 152,
	154, 155, 156, 157, 75, 327, 158, 75, 149, 140,
	385, 103, 330, 209, 424, 324, 136, 151, 413, 349,
	347, 427, 275, 197, 131, 101, 69, 261, 217, 216,
	225, 226, 219, 220, 221, 222, 223, 224, 218, 37,
	261, 125, 72, 269, 135, 153, 152, 154, 155, 156,
	157, 75, 18, 158, 164, 165, 17, 16, 137, 15,
	163, 14, 13, 449, 19, 12, 11, 10, 450, 9,
	8, 7, 453, 454, 395, 6, 5, 4, 2, 0,
	143, 144, 130, 132, 75, 162, 0, 145, 0, 146,
	281, 0, 151, 0, 350, 353, 354, 355, 351, 0,
	352, 356, 0, 159, 443, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 160, 161, 125, 0, 269, 135,
	153, 152, 154, 155, 156, 157, 0, 0, 158, 164,
	165, 132, 132, 137, 0, 163, 0, 0, 0, 0,
	0, 0, 0, 151, 0, 340, 0, 0, 75, 0,
	0, 75, 0, 0, 0, 143, 144, 130, 0, 0,
	162, 360, 145, 75, 146, 0, 0, 125, 0, 0,
	135, 153, 152, 154, 155, 156, 157, 0, 159, 158,
	164, 165, 0, 0, 137, 0, 163, 0, 0, 160,
	161, 217, 216, 225, 226, 219, 220, 221, 222, 223,
	224, 218, 0, 0, 0, 20, 143, 144, 130, 132,
	0, 162, 0, 145, 0, 146, 0, 0, 0, 0,
	0, 151, 0, 0, 0, 0, 0, 0, 0, 159,
	0, 75, 75, 75, 75, 0, 0, 0, 0, 0,
	160, 161, 0, 0, 360, 125, 0, 418, 135, 153,
	152, 154, 155, 156, 157, 0, 0, 158, 164, 165,
	0, 0, 137, 0, 163, 0, 0, 0, 151, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 143, 144, 0, 0, 0, 162,
	0, 145, 125, 146, 0, 135, 153, 152, 154, 155,
	156, 157, 0, 0, 158, 164, 165, 159, 0, 137,
	0, 163, 20, 21, 22, 23, 0, 0, 160, 161,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 143, 144, 0, 0, 24, 162, 0, 145, 125,
	146, 0, 135, 153, 152, 154, 155, 156, 157, 32,
	20, 158, 164, 165, 159, 0, 0, 0, 163, 0,
	0, 0, 0, 0, 0, 160, 161, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 143, 144,
	0, 0, 0, 162, 0, 145, 0, 146, 0, 0,
	125, 0, 0, 135, 153, 152, 154, 155, 156, 157,
	0, 159, 158, 0, 25, 26, 28, 27, 29, 163,
	0, 0, 160, 161, 0, 0, 0, 33, 36, 35,
	30, 0, 31, 0, 34, 0, 0, 0, 0, 143,
	144, 0, 0, 0, 162, 0, 145, 0, 146, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 159, 0, 0, 0, 0, 0, 211, 214,
	0, 0, 0, 160, 161, 227, 228, 229, 230, 231,
	232, 233, 215, 212, 213, 210, 217, 216, 225, 226,
	219, 220, 221, 222, 223, 224, 218,
}
var yyPact = [...]int{

	866, -1000, -1000, 511, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -74, -63, -41, -72, -42,
	-1000, 464, 251, -68, 90, -1000, -1000, 509, 436, 415,
	-1000, -63, 139, 502, 98, -69, -69, -46, -1000, -43,
	-1000, 139, -70, 238, -70, 139, -1000, -94, -1000, -1000,
	476, -44, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 319,
	171, -1000, 60, 414, 382, 6, -1000, 139, 256, -1000,
	46, -1000, 139, 58, 139, 211, 139, -56, 139, 428,
	296, 139, -1000, 281, -1000, -1000, -1000, 139, 139, -1000,
	-1000, 681, -1000, 456, -1000, 374, 325, -1000, 139, 98,
	139, 474, 98, 37, 281, 423, -1000, -76, 54, 139,
	-1000, -1000, 139, -1000, 202, -1000, -1000, -1000, 318, -1000,
	-1000, 213, 4, 176, 955, -1000, -1000, 806, 759, -1000,
	-33, -1000, -1000, 37, 37, 37, 37, 281, 281, -1000,
	-1000, 281, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 37, -1000, -1000, 139, -1000, -1000, -1000,
	-1000, 307, 321, -1000, 438, 806, -1000, 670, 3, 904,
	-1000, -1000, 289, 98, -1000, -57, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 196, -1000, -1000, 474, 681, 114,
	-1000, -1000, 239, -1000, -1000, 57, 806, 806, 173, 853,
	95, 116, 37, 37, 37, 173, 37, 37, 37, 37,
	37, 37, 37, 37, 37, 37, 37, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 38, 955, 128, 206, 244, 955,
	508, 508, -1000, -1000, -1000, 507, 555, 630, -1000, 509,
	52, 670, -1000, 437, 98, 98, 438, 454, 462, 176,
	149, 670, -1000, 139, -1000, -1000, 139, -1000, -1000, -1000,
	469, -1000, 265, 453, -1000, -1000, 194, 421, 179, -1000,
	-1000, 1, -1000, 38, 49, -1000, -1000, 142, -1000, -1000,
	-1000, 670, -1000, 904, -1000, -1000, 95, 37, 37, 37,
	670, 670, 420, -1000, 298, -10, -1000, 48, 48, -6,
	-6, -6, -6, 342, 342, -1000, -1000, 37, -1000, -1000,
	-1000, -1000, -1000, 219, 681, -1000, 219, 115, -1000, 806,
	288, 281, 511, 257, 250, -1000, 454, -1000, 37, 37,
	-1, 281, -1000, -1000, 467, 461, 114, 114, 114, 114,
	-1000, 359, 358, -1000, 368, 355, 339, -7, -1000, 221,
	-1000, -1000, 139, -1000, 246, 56, -1000, -1000, -1000, 244,
	-1000, 670, 670, 234, 37, 670, -1000, 219, -1000, 51,
	-1000, 37, 152, -1000, 408, 264, -1000, 37, -1000, -1000,
	98, -1000, 195, 263, -1000, 84, 98, -1000, 438, 806,
	37, 453, 285, 619, -1000, -1000, -1000, -1000, 357, -1000,
	343, -1000, -1000, -1000, -47, -48, -53, -1000, -1000, -1000,
	-1000, -1000, -1000, 37, 670, -1000, -1000, 670, 37, 388,
	281, -1000, 37, 37, 37, -1000, -1000, -1000, -1000, 454,
	176, 259, 806, 806, -1000, -1000, 281, 281, 281, 670,
	670, 506, -1000, 670, 670, -1000, 419, 176, 176, 98,
	98, 98, 98, -1000, 505, 29, 145, -1000, 145, 145,
	256, -1000, 98, -1000, 98, -1000, -1000, 98, -1000, -1000,
}
var yyPgo = [...]int{

	0, 638, 17, 637, 636, 635, 631, 630, 629, 627,
	626, 625, 624, 622, 621, 619, 617, 616, 612, 525,
	602, 599, 586, 585, 26, 28, 584, 583, 16, 15,
	25, 582, 580, 14, 579, 41, 578, 3, 24, 40,
	576, 20, 573, 572, 27, 199, 571, 19, 12, 5,
	570, 7, 52, 569, 568, 565, 22, 557, 554, 552,
	551, 545, 8, 544, 1, 541, 11, 540, 537, 535,
	13, 4, 33, 534, 354, 152, 532, 530, 528, 526,
	0, 21, 523, 444, 9, 521, 520, 6, 314, 519,
	10, 2, 518,
}
var yyR1 = [...]int{

	0, 89, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 2, 3, 3, 4, 5, 6, 7, 7, 7,
	8, 8, 9, 10, 10, 10, 11, 13, 14, 15,
	16, 16, 16, 16, 16, 17, 18, 12, 92, 19,
	20, 20, 21, 21, 21, 22, 22, 23, 23, 24,
	24, 25, 25, 25, 25, 26, 26, 82, 82, 82,
	81, 81, 27, 27, 28, 28, 29, 29, 30, 30,
	30, 31, 31, 31, 31, 86, 86, 85, 85, 85,
	84, 84, 32, 32, 32, 32, 33, 33, 33, 33,
	34, 34, 35, 35, 36, 36, 36, 36, 37, 37,
	38, 38, 39, 39, 39, 39, 39, 39, 41, 41,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 47, 47, 47, 47, 47, 47, 42,
	42, 42, 42, 42, 42, 42, 48, 48, 48, 52,
	49, 49, 45, 45, 45, 45, 45, 45, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, 45, 61,
	61, 61, 61, 54, 57, 57, 55, 55, 56, 58,
	58, 53, 53, 53, 44, 44, 44, 44, 44, 44,
	44, 46, 46, 46, 59, 59, 60, 60, 62, 62,
	63, 63, 64, 65, 65, 65, 66, 66, 66, 66,
	67, 67, 67, 68, 68, 69, 69, 70, 70, 43,
	43, 50, 50, 51, 71, 71, 72, 73, 73, 75,
	75, 88, 88, 74, 74, 76, 76, 76, 76, 76,
	76, 77, 77, 78, 78, 79, 79, 80, 83, 90,
	91, 87,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 12,
	6, 3, 8, 8, 8, 7, 3, 6, 4, 9,
	6, 7, 5, 4, 5, 4, 3, 2, 7, 3,
	3, 3, 5, 5, 3, 3, 2, 2, 0, 2,
	0, 2, 1, 2, 2, 0, 1, 0, 1, 1,
	3, 1, 2, 3, 5, 1, 1, 0, 1, 2,
	1, 1, 0, 2, 1, 3, 1, 1, 3, 3,
	3, 3, 5, 5, 3, 0, 1, 0, 1, 2,
	1, 1, 1, 2, 2, 1, 2, 3, 2, 3,
	2, 2, 1, 3, 0, 5, 5, 5, 1, 3,
	0, 2, 1, 3, 3, 2, 3, 3, 1, 1,
	1, 3, 3, 3, 4, 3, 4, 3, 4, 5,
	6, 3, 2, 1, 2, 1, 2, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 1, 3,
	1, 3, 1, 1, 1, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 2, 3, 3, 4, 5, 3, 4, 1, 1,
	1, 1, 1, 5, 0, 1, 1, 2, 4, 0,
	2, 1, 3, 5, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 0, 3, 0, 2, 0, 3,
	1, 3, 2, 0, 1, 1, 0, 2, 4, 4,
	0, 2, 4, 0, 3, 1, 3, 0, 5, 2,
	1, 1, 3, 3, 1, 3, 3, 1, 1, 0,
	2, 0, 3, 0, 1, 1, 1, 1, 1, 1,
	1, 0, 1, 0, 1, 0, 2, 1, 1, 1,
	1, 0,
}
var yyChk = [...]int{

	-1000, -89, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -13, -14, -15, -16, -17, -18, -12,
	6, 7, 8, 9, 29, 98, 99, 101, 100, 102,
	114, 116, 43, 111, 118, 113, 112, -21, 5, -19,
	-92, -19, -19, -19, -19, 103, 120, -78, 109, -74,
	107, 103, 103, 104, 120, 103, -87, 15, 49, 121,
	122, 98, 115, 52, -87, -87, -2, 19, 20, -22,
	33, 20, -20, -74, -35, -83, 49, 10, -71, -72,
	-80, 49, -88, 108, -88, 104, 103, -35, -75, 108,
	49, -75, -35, 117, -87, -87, -87, 10, 103, -87,
	-87, -23, 36, -46, -80, 52, 55, 57, 25, 29,
	95, -35, 47, 70, -35, 63, -83, 49, -35, 106,
	-35, 22, 45, -83, -90, 46, -83, -35, -24, -25,
	87, -26, -83, -39, -45, 49, -40, 63, -90, -44,
	-53, -51, -52, 85, 86, 92, 94, -80, -61, -54,
	-41, 22, 51, 50, 52, 53, 54, 55, 58, 108,
	119, 120, 90, 65, 59, 60, -73, 18, 10, 31,
	31, -35, -71, -83, -38, 11, -72, -45, -83, -90,
	-90, 22, -79, 110, -76, 101, 99, 28, 100, 14,
	123, 49, -35, -35, 49, -87, -87, -27, 47, 10,
	-82, -81, 21, -80, 51, 95, 62, 61, 77, -42,
	80, 63, 78, 79, 64, 77, 82, 81, 91, 85,
	86, 87, 88, 89, 90, 83, 84, 70, 71, 72,
	73, 74, 75, 76, -39, -45, -39, -2, -49, -45,
	96, 97, -45, -45, -45, -45, -90, -90, -52, -90,
	-57, -45, -35, -68, 29, -90, -38, -62, 14, -39,
	95, -45, -87, 45, -80, -87, -77, 106, -91, 48,
	-38, -25, -28, -29, -30, -31, -35, -52, -90, -81,
	87, -83, -80, -39, -39, -47, 58, 63, 59, 60,
	-41, -45, -48, -90, -52, 56, 80, 78, 79, 64,
	-45, -45, -45, -47, -45, -45, -45, -45, -45, -45,
	-45, -45, -45, -45, -45, -91, -91, 47, -91, -44,
	-44, -80, -91, -24, 20, -91, -24, -55, -56, 66,
	-43, 31, -2, -71, -69, -80, -62, -66, 16, 15,
	-83, -35, -35, -87, -59, 12, 47, -32, -33, -34,
	35, 39, 41, 36, 37, 38, 42, -85, -84, 21,
	-83, 51, -86, 21, -28, 95, 58, 59, 60, -49,
	-48, -45, -45, -45, 62, -45, -91, -24, -91, -58,
	-56, 68, -39, -70, 45, -50, -51, -90, -70, -91,
	47, -66, -45, -63, -64, -45, 95, -90, -60, 13,
	15, -29, -30, -29, -30, 35, 35, 35, 40, 35,
	40, 35, -33, -36, 43, 107, 44, -84, -83, -91,
	87, -80, -91, 62, -45, -91, 69, -45, 67, 26,
	47, -80, 47, 17, 47, -65, 23, 24, -87, -62,
	-39, -49, 45, 45, 35, 35, 104, 104, 104, -45,
	-45, 27, -51, -45, -45, -64, -66, -39, -39, -90,
	-90, -90, 8, -67, 18, 30, -37, -80, -37, -37,
	-71, 8, 80, -91, 47, -91, -91, -80, -80, -80,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	48, 48, 48, 48, 48, 253, 243, 0, 0, 0,
	261, 0, 0, 0, 0, 261, 261, 0, 52, 55,
	50, 243, 0, 0, 0, 241, 241, 0, 254, 0,
	244, 0, 239, 0, 239, 0, 37, 0, 261, 261,
	261, 0, 261, 261, 46, 47, 21, 53, 54, 57,
	0, 56, 49, 0, 0, 102, 258, 0, 26, 234,
	0, 257, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 36, 0, 39, 40, 41, 0, 0, 44,
	45, 0, 58, 0, 201, 0, 0, 51, 0, 0,
	0, 110, 0, 0, 0, 0, 28, 255, 0, 0,
	33, 240, 0, 35, 0, 259, 261, 261, 72, 59,
	61, 67, 0, 65, 66, -2, 112, 0, 0, 152,
	153, 154, 155, 0, 0, 0, 0, 191, 0, 178,
	120, 0, 194, 195, 196, 197, 198, 199, 200, 179,
	180, 181, 182, 184, 118, 119, 0, 237, 238, 202,
	203, 223, 110, 103, 208, 0, 235, 236, 0, 0,
	261, 242, 0, 0, 261, 251, 245, 246, 247, 248,
	249, 250, 32, 34, 0, 42, 43, 110, 0, 0,
	62, 68, 0, 70, 71, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 139, 140, 141,
	142, 143, 144, 145, 115, 0, 0, 0, 0, 150,
	0, 0, 169, 170, 171, 0, 0, 0, 132, 0,
	0, 185, 20, 0, 0, 0, 208, 216, 0, 111,
	0, 150, 27, 0, 256, 30, 0, 252, 261, 260,
	204, 60, 73, 74, 76, 77, 87, 85, 0, 69,
	63, 0, 192, 113, 114, 117, 133, 0, 135, 137,
	121, 122, 123, 0, 147, 148, 0, 0, 0, 0,
	125, 127, 0, 131, 156, 157, 158, 159, 160, 161,
	162, 163, 164, 165, 166, 116, 149, 0, 233, 167,
	168, 172, 173, 0, 0, 176, 0, 189, 186, 0,
	227, 0, 230, 227, 0, 225, 216, 25, 0, 0,
	0, 0, 31, 38, 206, 0, 0, 0, 0, 0,
	92, 0, 0, 95, 0, 0, 0, 104, 88, 0,
	90, 91, 0, 86, 0, 0, 134, 136, 138, 0,
	124, 126, 128, 0, 0, 151, 174, 0, 177, 0,
	187, 0, 0, 22, 0, 229, 231, 0, 23, 224,
	0, 24, 217, 209, 210, 213, 0, 261, 208, 0,
	0, 75, 81, 0, 84, 93, 94, 96, 0, 98,
	0, 100, 101, 78, 0, 0, 0, 89, 79, 80,
	64, 193, 146, 0, 129, 175, 183, 190, 0, 0,
	0, 226, 0, 0, 0, 212, 214, 215, 29, 216,
	207, 205, 0, 0, 97, 99, 0, 0, 0, 130,
	188, 0, 232, 218, 219, 211, 220, 82, 83, 0,
	0, 0, 0, 19, 0, 0, 0, 108, 0, 0,
	228, 221, 0, 105, 0, 106, 107, 0, 109, 222,
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
	118, 119, 120, 121, 122, 123,
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
		//line sql.y:182
		{
			setParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:188
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 19:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:210
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(WhereStr, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(HavingStr, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 20:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:214
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{Expr: yyDollar[4].valExpr}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:218
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:224
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 23:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:228
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, updateList := range yyDollar[7].updateExprs {
				cols = append(cols, updateList.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:240
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:246
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:252
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:258
		{
			yyVAL.statement = &DDL{Action: CreateTableStr, IfNotExists: bool(yyDollar[3].boolVal), Table: yyDollar[4].tableName, NewName: yyDollar[4].tableName}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:263
		{
			yyVAL.statement = &DDL{Action: CreateDBStr, IfNotExists: bool(yyDollar[3].boolVal), Database: yyDollar[4].tableIdent}
		}
	case 29:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line sql.y:267
		{
			yyVAL.statement = &DDL{Action: CreateIndexStr, Table: yyDollar[7].tableName, NewName: yyDollar[7].tableName}
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:273
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableName, NewName: yyDollar[4].tableName}
		}
	case 31:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:277
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableName, NewName: yyDollar[7].tableName}
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:283
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableName, NewName: yyDollar[5].tableName}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:289
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropTableStr, Table: yyDollar[4].tableName, IfExists: exists}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:297
		{
			yyVAL.statement = &DDL{Action: DropIndexStr, Table: yyDollar[5].tableName, NewName: yyDollar[5].tableName}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:301
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropDBStr, Database: yyDollar[4].tableIdent, IfExists: exists}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:312
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableName, NewName: yyDollar[3].tableName}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:318
		{
			yyVAL.statement = &Xa{}
		}
	case 38:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:324
		{
			yyVAL.statement = &Shard{ShardKey: string(yyDollar[5].bytes)}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:330
		{
			yyVAL.statement = &UseDB{Database: string(yyDollar[2].bytes)}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:336
		{
			yyVAL.statement = &ShowDatabases{}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:340
		{
			yyVAL.statement = &ShowTables{}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:344
		{
			yyVAL.statement = &ShowTables{Database: yyDollar[4].tableIdent}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:348
		{
			yyVAL.statement = &ShowCreateTable{Table: yyDollar[4].tableName}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:352
		{
			yyVAL.statement = &ShowProcesslist{}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:358
		{
			yyVAL.statement = &Kill{QueryID: string(yyDollar[2].bytes)}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:364
		{
			yyVAL.statement = &Explain{}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:370
		{
			yyVAL.statement = &Other{}
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:375
		{
			setAllowComments(yylex, true)
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:379
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:385
		{
			yyVAL.bytes2 = nil
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:389
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:395
		{
			yyVAL.str = UnionStr
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:399
		{
			yyVAL.str = UnionAllStr
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:403
		{
			yyVAL.str = UnionDistinctStr
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:408
		{
			yyVAL.str = ""
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:412
		{
			yyVAL.str = DistinctStr
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:417
		{
			yyVAL.str = ""
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:421
		{
			yyVAL.str = StraightJoinHint
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:427
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:431
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:437
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:441
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:445
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Name: yyDollar[1].tableIdent}}
		}
	case 64:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:449
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}}
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:455
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:459
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:464
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:468
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:472
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:479
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 72:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:484
		{
			yyVAL.tableExprs = TableExprs{&AliasedTableExpr{Expr: &TableName{Name: NewTableIdent("dual")}}}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:488
		{
			yyVAL.tableExprs = yyDollar[2].tableExprs
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:494
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:498
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:508
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:512
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:516
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:529
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:533
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:537
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:541
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:546
		{
			yyVAL.empty = struct{}{}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:548
		{
			yyVAL.empty = struct{}{}
		}
	case 87:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:551
		{
			yyVAL.tableIdent = NewTableIdent("")
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:555
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:559
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:566
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:572
		{
			yyVAL.str = JoinStr
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:576
		{
			yyVAL.str = JoinStr
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:580
		{
			yyVAL.str = JoinStr
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:584
		{
			yyVAL.str = StraightJoinStr
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:590
		{
			yyVAL.str = LeftJoinStr
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:594
		{
			yyVAL.str = LeftJoinStr
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:598
		{
			yyVAL.str = RightJoinStr
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:602
		{
			yyVAL.str = RightJoinStr
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:608
		{
			yyVAL.str = NaturalJoinStr
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:612
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:622
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:626
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:631
		{
			yyVAL.indexHints = nil
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:635
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 106:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:639
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 107:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:643
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:649
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:653
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 110:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:658
		{
			yyVAL.boolExpr = nil
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:662
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:669
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:673
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:677
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:681
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:685
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:691
		{
			yyVAL.boolVal = BoolVal(true)
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:695
		{
			yyVAL.boolVal = BoolVal(false)
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:701
		{
			yyVAL.boolExpr = yyDollar[1].boolVal
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:705
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].boolVal}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:709
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:713
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 124:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:717
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:721
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 126:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:725
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:729
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 128:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:733
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 129:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:737
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 130:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:741
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:745
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:749
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:755
		{
			yyVAL.str = IsNullStr
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:759
		{
			yyVAL.str = IsNotNullStr
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:763
		{
			yyVAL.str = IsTrueStr
		}
	case 136:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:767
		{
			yyVAL.str = IsNotTrueStr
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:771
		{
			yyVAL.str = IsFalseStr
		}
	case 138:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:775
		{
			yyVAL.str = IsNotFalseStr
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:781
		{
			yyVAL.str = EqualStr
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:785
		{
			yyVAL.str = LessThanStr
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:789
		{
			yyVAL.str = GreaterThanStr
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:793
		{
			yyVAL.str = LessEqualStr
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:797
		{
			yyVAL.str = GreaterEqualStr
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:801
		{
			yyVAL.str = NotEqualStr
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:805
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:811
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:815
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:819
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:825
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:831
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:835
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:841
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:845
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:849
		{
			yyVAL.valExpr = yyDollar[1].valTuple
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:853
		{
			yyVAL.valExpr = yyDollar[1].subquery
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:857
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:861
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:865
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:869
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:873
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:877
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:881
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:885
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:889
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:893
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:897
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:901
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONExtractOp, Right: yyDollar[3].valExpr}
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:905
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONUnquoteExtractOp, Right: yyDollar[3].valExpr}
		}
	case 169:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:909
		{
			if num, ok := yyDollar[2].valExpr.(*SQLVal); ok && num.Type == IntVal {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:917
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
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:931
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:935
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.valExpr = &IntervalExpr{Expr: yyDollar[2].valExpr, Unit: yyDollar[3].colIdent}
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:943
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 174:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:947
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 175:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:951
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:955
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 177:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:959
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:963
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:969
		{
			yyVAL.colIdent = NewColIdent("if")
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:973
		{
			yyVAL.colIdent = NewColIdent("current_timestamp")
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:977
		{
			yyVAL.colIdent = NewColIdent("database")
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:981
		{
			yyVAL.colIdent = NewColIdent("mod")
		}
	case 183:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:987
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 184:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:992
		{
			yyVAL.valExpr = nil
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:996
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1002
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1006
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 188:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1012
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1017
		{
			yyVAL.valExpr = nil
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1021
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1027
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1031
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 193:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1041
		{
			yyVAL.valExpr = NewStrVal(yyDollar[1].bytes)
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1045
		{
			yyVAL.valExpr = NewHexVal(yyDollar[1].bytes)
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1049
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1053
		{
			yyVAL.valExpr = NewFloatVal(yyDollar[1].bytes)
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1057
		{
			yyVAL.valExpr = NewHexNum(yyDollar[1].bytes)
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1061
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1065
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1071
		{
			// TODO(sougou): Deprecate this construct.
			if yyDollar[1].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.valExpr = NewIntVal([]byte("1"))
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1080
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1084
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1089
		{
			yyVAL.valExprs = nil
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1093
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1098
		{
			yyVAL.boolExpr = nil
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1102
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.orderBy = nil
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1111
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1121
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 212:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1127
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1132
		{
			yyVAL.str = AscScr
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1136
		{
			yyVAL.str = AscScr
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1140
		{
			yyVAL.str = DescScr
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1145
		{
			yyVAL.limit = nil
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1149
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 218:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1153
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 219:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1157
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr, Offset: yyDollar[4].valExpr}
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1163
		{
			yyVAL.str = ""
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1167
		{
			yyVAL.str = ForUpdateStr
		}
	case 222:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1171
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
	case 223:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1184
		{
			yyVAL.columns = nil
		}
	case 224:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1188
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1194
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 226:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1198
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 227:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1203
		{
			yyVAL.updateExprs = nil
		}
	case 228:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1207
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1213
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1217
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1223
		{
			yyVAL.values = Values{yyDollar[1].valTuple}
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1227
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].valTuple)
		}
	case 233:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1233
		{
			yyVAL.valTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1239
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1243
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1249
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 239:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1258
		{
			yyVAL.byt = 0
		}
	case 240:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1260
		{
			yyVAL.byt = 1
		}
	case 241:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1263
		{
			yyVAL.boolVal = false
		}
	case 242:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1265
		{
			yyVAL.boolVal = true
		}
	case 243:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1268
		{
			yyVAL.str = ""
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1270
		{
			yyVAL.str = IgnoreStr
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1274
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1276
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1278
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1280
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1282
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1284
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1287
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1289
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1292
		{
			yyVAL.empty = struct{}{}
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1294
		{
			yyVAL.empty = struct{}{}
		}
	case 255:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1297
		{
			yyVAL.empty = struct{}{}
		}
	case 256:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1299
		{
			yyVAL.empty = struct{}{}
		}
	case 257:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1303
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1309
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 259:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1315
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1324
		{
			decNesting(yylex)
		}
	case 261:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1329
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}

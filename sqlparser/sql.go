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
	-1, 134,
	95, 257,
	-2, 256,
}

const yyNprod = 261
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1064

var yyAct = [...]int{

	146, 393, 267, 465, 77, 237, 55, 140, 256, 357,
	178, 336, 291, 382, 347, 272, 271, 236, 3, 284,
	149, 200, 327, 92, 173, 273, 127, 138, 128, 51,
	52, 44, 60, 78, 182, 413, 415, 47, 88, 82,
	132, 63, 64, 188, 79, 49, 53, 266, 45, 61,
	118, 323, 141, 150, 65, 58, 59, 186, 447, 446,
	445, 84, 97, 85, 93, 94, 95, 54, 98, 99,
	103, 50, 239, 240, 395, 364, 259, 124, 190, 268,
	134, 152, 151, 153, 154, 155, 156, 217, 204, 157,
	163, 164, 109, 471, 136, 205, 162, 207, 112, 414,
	328, 425, 80, 123, 435, 436, 134, 114, 106, 79,
	207, 137, 79, 171, 62, 80, 142, 143, 129, 75,
	328, 161, 380, 144, 179, 145, 87, 298, 185, 187,
	184, 202, 194, 195, 220, 221, 222, 223, 217, 158,
	419, 296, 297, 295, 279, 175, 268, 137, 137, 134,
	159, 160, 189, 365, 366, 367, 193, 245, 246, 206,
	205, 248, 216, 215, 224, 225, 218, 219, 220, 221,
	222, 223, 217, 206, 205, 207, 358, 233, 235, 427,
	90, 254, 201, 263, 19, 137, 261, 124, 116, 207,
	264, 206, 205, 285, 287, 288, 255, 294, 286, 234,
	473, 268, 202, 247, 75, 281, 360, 207, 137, 277,
	80, 432, 203, 197, 268, 258, 137, 137, 37, 80,
	292, 269, 104, 278, 124, 105, 270, 75, 75, 289,
	360, 80, 124, 203, 302, 75, 316, 268, 314, 315,
	317, 431, 345, 268, 268, 320, 282, 283, 321, 324,
	89, 276, 389, 268, 79, 334, 137, 137, 332, 57,
	281, 268, 293, 383, 335, 111, 174, 318, 319, 111,
	331, 316, 322, 325, 342, 216, 215, 224, 225, 218,
	219, 220, 221, 222, 223, 217, 433, 429, 277, 198,
	345, 124, 253, 441, 363, 422, 383, 262, 368, 121,
	133, 410, 111, 353, 354, 408, 292, 81, 369, 124,
	409, 48, 176, 101, 216, 215, 224, 225, 218, 219,
	220, 221, 222, 223, 217, 375, 197, 406, 377, 444,
	276, 443, 407, 169, 137, 405, 388, 238, 385, 137,
	404, 386, 241, 242, 243, 244, 387, 390, 293, 379,
	376, 396, 72, 83, 70, 168, 277, 277, 277, 277,
	108, 400, 250, 402, 463, 420, 418, 69, 416, 381,
	411, 421, 401, 428, 403, 450, 464, 19, 260, 424,
	215, 224, 225, 218, 219, 220, 221, 222, 223, 217,
	430, 107, 180, 337, 120, 362, 420, 133, 276, 276,
	276, 276, 330, 437, 399, 440, 438, 338, 290, 137,
	56, 299, 300, 301, 257, 303, 304, 305, 306, 307,
	308, 309, 310, 311, 312, 313, 224, 225, 218, 219,
	220, 221, 222, 223, 217, 454, 167, 451, 177, 439,
	386, 66, 67, 398, 166, 133, 133, 344, 174, 96,
	455, 76, 137, 137, 470, 461, 458, 459, 460, 466,
	466, 466, 79, 467, 468, 19, 469, 37, 472, 39,
	474, 475, 476, 1, 477, 361, 275, 478, 356, 199,
	74, 373, 456, 457, 349, 352, 353, 354, 350, 74,
	351, 355, 260, 74, 442, 181, 370, 371, 372, 46,
	216, 215, 224, 225, 218, 219, 220, 221, 222, 223,
	217, 265, 183, 165, 333, 74, 374, 252, 73, 38,
	74, 462, 115, 133, 74, 434, 74, 86, 392, 122,
	147, 91, 397, 343, 378, 125, 74, 391, 394, 131,
	40, 41, 42, 43, 249, 326, 74, 148, 172, 139,
	384, 102, 329, 110, 208, 135, 412, 74, 113, 348,
	74, 346, 117, 274, 119, 218, 219, 220, 221, 222,
	223, 217, 196, 423, 126, 130, 100, 68, 36, 71,
	426, 17, 150, 16, 170, 15, 260, 14, 152, 151,
	153, 154, 155, 156, 13, 191, 157, 18, 192, 260,
	12, 11, 10, 9, 74, 8, 124, 7, 268, 134,
	152, 151, 153, 154, 155, 156, 6, 5, 157, 163,
	164, 4, 448, 136, 2, 162, 0, 449, 0, 0,
	0, 452, 453, 394, 0, 0, 131, 74, 0, 0,
	0, 0, 251, 280, 0, 142, 143, 129, 0, 0,
	161, 0, 144, 0, 145, 0, 0, 0, 0, 150,
	349, 352, 353, 354, 350, 0, 351, 355, 158, 0,
	0, 19, 20, 21, 22, 0, 0, 0, 0, 159,
	160, 0, 0, 124, 131, 131, 134, 152, 151, 153,
	154, 155, 156, 0, 23, 157, 163, 164, 339, 0,
	136, 74, 162, 0, 74, 0, 0, 0, 31, 0,
	0, 0, 0, 0, 359, 0, 74, 0, 0, 0,
	0, 0, 142, 143, 129, 0, 0, 161, 0, 144,
	0, 145, 0, 0, 0, 0, 0, 0, 0, 340,
	0, 0, 341, 0, 0, 158, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 159, 160, 19, 0,
	0, 0, 131, 24, 25, 27, 26, 28, 0, 0,
	0, 0, 0, 0, 150, 0, 32, 34, 35, 29,
	0, 30, 0, 33, 74, 74, 74, 74, 0, 0,
	80, 0, 0, 0, 0, 0, 0, 359, 124, 0,
	417, 134, 152, 151, 153, 154, 155, 156, 0, 0,
	157, 163, 164, 0, 0, 136, 0, 162, 0, 0,
	0, 150, 216, 215, 224, 225, 218, 219, 220, 221,
	222, 223, 217, 0, 0, 0, 0, 142, 143, 0,
	0, 0, 161, 0, 144, 124, 145, 0, 134, 152,
	151, 153, 154, 155, 156, 0, 0, 157, 163, 164,
	158, 0, 136, 0, 162, 0, 0, 0, 0, 0,
	0, 159, 160, 216, 215, 224, 225, 218, 219, 220,
	221, 222, 223, 217, 142, 143, 0, 0, 0, 161,
	0, 144, 124, 145, 0, 134, 152, 151, 153, 154,
	155, 156, 0, 19, 157, 163, 164, 158, 0, 0,
	0, 162, 0, 0, 0, 0, 0, 0, 159, 160,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 142, 143, 0, 0, 0, 161, 0, 144, 0,
	145, 0, 0, 124, 0, 0, 134, 152, 151, 153,
	154, 155, 156, 0, 158, 157, 0, 0, 0, 0,
	0, 0, 162, 0, 0, 159, 160, 124, 0, 0,
	134, 152, 151, 153, 154, 155, 156, 0, 0, 157,
	0, 0, 142, 143, 0, 0, 162, 161, 0, 144,
	0, 145, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 158, 142, 143, 0, 0,
	0, 161, 0, 144, 0, 145, 159, 160, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 158,
	0, 0, 0, 0, 0, 210, 213, 0, 0, 0,
	159, 160, 226, 227, 228, 229, 230, 231, 232, 214,
	211, 212, 209, 216, 215, 224, 225, 218, 219, 220,
	221, 222, 223, 217,
}
var yyPact = [...]int{

	665, -1000, -1000, 462, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -72, -62, -32, -74, -36, -1000,
	395, 210, -66, 62, -1000, -1000, 459, 422, 334, -1000,
	-62, 70, 441, 66, -69, -69, -43, -1000, -40, -1000,
	70, -70, 201, -70, 70, -1000, -94, -1000, -1000, 439,
	-41, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 277, 170,
	-1000, 51, 366, 331, -3, -1000, 70, 222, -1000, 28,
	-1000, 70, 44, 70, 139, 70, -56, 70, 372, 254,
	70, -1000, 245, -1000, -1000, -1000, 70, 70, -1000, -1000,
	637, -1000, 426, -1000, 324, 302, -1000, 70, 66, 70,
	437, 66, 921, 245, 370, -1000, -76, 29, 70, -1000,
	-1000, 70, -1000, 107, -1000, -1000, -1000, 279, -1000, -1000,
	161, -7, 130, 972, -1000, -1000, 799, 752, -1000, -24,
	-1000, -1000, 921, 921, 921, 921, 245, 245, -1000, -1000,
	245, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 921, -1000, -1000, 70, -1000, -1000, -1000, -1000,
	263, 255, -1000, 400, 799, -1000, 792, -19, 897, -1000,
	-1000, 252, 66, -1000, -59, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 196, -1000, -1000, 437, 637, 186, -1000,
	-1000, 182, -1000, -1000, 57, 799, 799, 135, 846, 141,
	63, 921, 921, 921, 135, 921, 921, 921, 921, 921,
	921, 921, 921, 921, 921, 921, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 20, 972, 98, 213, 189, 972, 538,
	538, -1000, -1000, -1000, 741, 31, 560, -1000, 459, 34,
	792, -1000, 371, 66, 66, 400, 377, 392, 130, 100,
	792, -1000, 70, -1000, -1000, 70, -1000, -1000, -1000, 435,
	-1000, 243, 625, -1000, -1000, 155, 374, 178, -1000, -1000,
	-20, -1000, 20, 33, -1000, -1000, 95, -1000, -1000, -1000,
	792, -1000, 897, -1000, -1000, 141, 921, 921, 921, 792,
	792, 419, -1000, 343, 298, -1000, 47, 47, -4, -4,
	-4, -4, 480, 480, -1000, -1000, 921, -1000, -1000, -1000,
	-1000, -1000, 166, 637, -1000, 166, 54, -1000, 799, 251,
	245, 462, 218, 205, -1000, 377, -1000, 921, 921, -21,
	245, -1000, -1000, 430, 389, 186, 186, 186, 186, -1000,
	305, 300, -1000, 292, 270, 266, -8, -1000, 179, -1000,
	-1000, 70, -1000, 195, 53, -1000, -1000, -1000, 189, -1000,
	792, 792, 233, 921, 792, -1000, 166, -1000, 32, -1000,
	921, 112, -1000, 347, 240, -1000, 921, -1000, -1000, 66,
	-1000, 194, 239, -1000, 81, 66, -1000, 400, 799, 921,
	625, 248, 449, -1000, -1000, -1000, -1000, 296, -1000, 294,
	-1000, -1000, -1000, -44, -45, -46, -1000, -1000, -1000, -1000,
	-1000, -1000, 921, 792, -1000, -1000, 792, 921, 348, 245,
	-1000, 921, 921, 921, -1000, -1000, -1000, -1000, 377, 130,
	224, 799, 799, -1000, -1000, 245, 245, 245, 792, 792,
	447, -1000, 792, 792, -1000, 346, 130, 130, 66, 66,
	66, 66, -1000, 446, 13, 153, -1000, 153, 153, 222,
	-1000, 66, -1000, 66, -1000, -1000, 66, -1000, -1000,
}
var yyPgo = [...]int{

	0, 624, 17, 621, 617, 616, 607, 605, 603, 602,
	601, 600, 597, 594, 587, 585, 583, 581, 519, 579,
	578, 577, 576, 26, 28, 575, 572, 16, 15, 25,
	563, 561, 14, 559, 476, 556, 3, 24, 40, 555,
	20, 554, 552, 27, 199, 551, 19, 12, 5, 550,
	7, 52, 549, 547, 545, 22, 544, 534, 533, 532,
	530, 8, 528, 1, 525, 11, 521, 517, 514, 13,
	4, 33, 513, 311, 126, 512, 511, 499, 495, 0,
	21, 479, 438, 9, 478, 475, 6, 307, 473, 10,
	2, 469,
}
var yyR1 = [...]int{

	0, 88, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	2, 3, 3, 4, 5, 6, 7, 7, 7, 8,
	8, 9, 10, 10, 10, 11, 13, 14, 15, 16,
	16, 16, 16, 16, 17, 12, 12, 91, 18, 19,
	19, 20, 20, 20, 21, 21, 22, 22, 23, 23,
	24, 24, 24, 24, 25, 25, 81, 81, 81, 80,
	80, 26, 26, 27, 27, 28, 28, 29, 29, 29,
	30, 30, 30, 30, 85, 85, 84, 84, 84, 83,
	83, 31, 31, 31, 31, 32, 32, 32, 32, 33,
	33, 34, 34, 35, 35, 35, 35, 36, 36, 37,
	37, 38, 38, 38, 38, 38, 38, 40, 40, 39,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 46, 46, 46, 46, 46, 46, 41, 41,
	41, 41, 41, 41, 41, 47, 47, 47, 51, 48,
	48, 44, 44, 44, 44, 44, 44, 44, 44, 44,
	44, 44, 44, 44, 44, 44, 44, 44, 44, 44,
	44, 44, 44, 44, 44, 44, 44, 44, 60, 60,
	60, 60, 53, 56, 56, 54, 54, 55, 57, 57,
	52, 52, 52, 43, 43, 43, 43, 43, 43, 43,
	45, 45, 45, 58, 58, 59, 59, 61, 61, 62,
	62, 63, 64, 64, 64, 65, 65, 65, 65, 66,
	66, 66, 67, 67, 68, 68, 69, 69, 42, 42,
	49, 49, 50, 70, 70, 71, 72, 72, 74, 74,
	87, 87, 73, 73, 75, 75, 75, 75, 75, 75,
	76, 76, 77, 77, 78, 78, 79, 82, 89, 90,
	86,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 12, 6,
	3, 8, 8, 8, 7, 3, 6, 4, 9, 6,
	7, 5, 4, 5, 4, 3, 2, 7, 3, 3,
	3, 5, 5, 3, 3, 2, 2, 0, 2, 0,
	2, 1, 2, 2, 0, 1, 0, 1, 1, 3,
	1, 2, 3, 5, 1, 1, 0, 1, 2, 1,
	1, 0, 2, 1, 3, 1, 1, 3, 3, 3,
	3, 5, 5, 3, 0, 1, 0, 1, 2, 1,
	1, 1, 2, 2, 1, 2, 3, 2, 3, 2,
	2, 1, 3, 0, 5, 5, 5, 1, 3, 0,
	2, 1, 3, 3, 2, 3, 3, 1, 1, 1,
	3, 3, 3, 4, 3, 4, 3, 4, 5, 6,
	3, 2, 1, 2, 1, 2, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 1, 3, 1,
	3, 1, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 3, 3, 4, 5, 3, 4, 1, 1, 1,
	1, 1, 5, 0, 1, 1, 2, 4, 0, 2,
	1, 3, 5, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 0, 3, 0, 2, 0, 3, 1,
	3, 2, 0, 1, 1, 0, 2, 4, 4, 0,
	2, 4, 0, 3, 1, 3, 0, 5, 2, 1,
	1, 3, 3, 1, 3, 3, 1, 1, 0, 2,
	0, 3, 0, 1, 1, 1, 1, 1, 1, 1,
	0, 1, 0, 1, 0, 2, 1, 1, 1, 1,
	0,
}
var yyChk = [...]int{

	-1000, -88, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -13, -14, -15, -16, -17, -12, 6,
	7, 8, 9, 29, 98, 99, 101, 100, 102, 114,
	116, 43, 111, 118, 112, 113, -20, 5, -18, -91,
	-18, -18, -18, -18, 103, 120, -77, 109, -73, 107,
	103, 103, 104, 120, 103, -86, 15, 49, 121, 122,
	98, 115, 52, -86, -86, -2, 19, 20, -21, 33,
	20, -19, -73, -34, -82, 49, 10, -70, -71, -79,
	49, -87, 108, -87, 104, 103, -34, -74, 108, 49,
	-74, -34, 117, -86, -86, -86, 10, 103, -86, -86,
	-22, 36, -45, -79, 52, 55, 57, 25, 29, 95,
	-34, 47, 70, -34, 63, -82, 49, -34, 106, -34,
	22, 45, -82, -89, 46, -82, -34, -23, -24, 87,
	-25, -82, -38, -44, 49, -39, 63, -89, -43, -52,
	-50, -51, 85, 86, 92, 94, -79, -60, -53, -40,
	22, 51, 50, 52, 53, 54, 55, 58, 108, 119,
	120, 90, 65, 59, 60, -72, 18, 10, 31, 31,
	-34, -70, -82, -37, 11, -71, -44, -82, -89, -89,
	22, -78, 110, -75, 101, 99, 28, 100, 14, 123,
	49, -34, -34, 49, -86, -86, -26, 47, 10, -81,
	-80, 21, -79, 51, 95, 62, 61, 77, -41, 80,
	63, 78, 79, 64, 77, 82, 81, 91, 85, 86,
	87, 88, 89, 90, 83, 84, 70, 71, 72, 73,
	74, 75, 76, -38, -44, -38, -2, -48, -44, 96,
	97, -44, -44, -44, -44, -89, -89, -51, -89, -56,
	-44, -34, -67, 29, -89, -37, -61, 14, -38, 95,
	-44, -86, 45, -79, -86, -76, 106, -90, 48, -37,
	-24, -27, -28, -29, -30, -34, -51, -89, -80, 87,
	-82, -79, -38, -38, -46, 58, 63, 59, 60, -40,
	-44, -47, -89, -51, 56, 80, 78, 79, 64, -44,
	-44, -44, -46, -44, -44, -44, -44, -44, -44, -44,
	-44, -44, -44, -44, -90, -90, 47, -90, -43, -43,
	-79, -90, -23, 20, -90, -23, -54, -55, 66, -42,
	31, -2, -70, -68, -79, -61, -65, 16, 15, -82,
	-34, -34, -86, -58, 12, 47, -31, -32, -33, 35,
	39, 41, 36, 37, 38, 42, -84, -83, 21, -82,
	51, -85, 21, -27, 95, 58, 59, 60, -48, -47,
	-44, -44, -44, 62, -44, -90, -23, -90, -57, -55,
	68, -38, -69, 45, -49, -50, -89, -69, -90, 47,
	-65, -44, -62, -63, -44, 95, -89, -59, 13, 15,
	-28, -29, -28, -29, 35, 35, 35, 40, 35, 40,
	35, -32, -35, 43, 107, 44, -83, -82, -90, 87,
	-79, -90, 62, -44, -90, 69, -44, 67, 26, 47,
	-79, 47, 17, 47, -64, 23, 24, -86, -61, -38,
	-48, 45, 45, 35, 35, 104, 104, 104, -44, -44,
	27, -50, -44, -44, -63, -65, -38, -38, -89, -89,
	-89, 8, -66, 18, 30, -36, -79, -36, -36, -70,
	8, 80, -90, 47, -90, -90, -79, -79, -79,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 47,
	47, 47, 47, 47, 252, 242, 0, 0, 0, 260,
	0, 0, 0, 0, 260, 260, 0, 51, 54, 49,
	242, 0, 0, 0, 240, 240, 0, 253, 0, 243,
	0, 238, 0, 238, 0, 36, 0, 260, 260, 260,
	0, 260, 260, 45, 46, 20, 52, 53, 56, 0,
	55, 48, 0, 0, 101, 257, 0, 25, 233, 0,
	256, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 35, 0, 38, 39, 40, 0, 0, 43, 44,
	0, 57, 0, 200, 0, 0, 50, 0, 0, 0,
	109, 0, 0, 0, 0, 27, 254, 0, 0, 32,
	239, 0, 34, 0, 258, 260, 260, 71, 58, 60,
	66, 0, 64, 65, -2, 111, 0, 0, 151, 152,
	153, 154, 0, 0, 0, 0, 190, 0, 177, 119,
	0, 193, 194, 195, 196, 197, 198, 199, 178, 179,
	180, 181, 183, 117, 118, 0, 236, 237, 201, 202,
	222, 109, 102, 207, 0, 234, 235, 0, 0, 260,
	241, 0, 0, 260, 250, 244, 245, 246, 247, 248,
	249, 31, 33, 0, 41, 42, 109, 0, 0, 61,
	67, 0, 69, 70, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 138, 139, 140, 141,
	142, 143, 144, 114, 0, 0, 0, 0, 149, 0,
	0, 168, 169, 170, 0, 0, 0, 131, 0, 0,
	184, 19, 0, 0, 0, 207, 215, 0, 110, 0,
	149, 26, 0, 255, 29, 0, 251, 260, 259, 203,
	59, 72, 73, 75, 76, 86, 84, 0, 68, 62,
	0, 191, 112, 113, 116, 132, 0, 134, 136, 120,
	121, 122, 0, 146, 147, 0, 0, 0, 0, 124,
	126, 0, 130, 155, 156, 157, 158, 159, 160, 161,
	162, 163, 164, 165, 115, 148, 0, 232, 166, 167,
	171, 172, 0, 0, 175, 0, 188, 185, 0, 226,
	0, 229, 226, 0, 224, 215, 24, 0, 0, 0,
	0, 30, 37, 205, 0, 0, 0, 0, 0, 91,
	0, 0, 94, 0, 0, 0, 103, 87, 0, 89,
	90, 0, 85, 0, 0, 133, 135, 137, 0, 123,
	125, 127, 0, 0, 150, 173, 0, 176, 0, 186,
	0, 0, 21, 0, 228, 230, 0, 22, 223, 0,
	23, 216, 208, 209, 212, 0, 260, 207, 0, 0,
	74, 80, 0, 83, 92, 93, 95, 0, 97, 0,
	99, 100, 77, 0, 0, 0, 88, 78, 79, 63,
	192, 145, 0, 128, 174, 182, 189, 0, 0, 0,
	225, 0, 0, 0, 211, 213, 214, 28, 215, 206,
	204, 0, 0, 96, 98, 0, 0, 0, 129, 187,
	0, 231, 217, 218, 210, 219, 81, 82, 0, 0,
	0, 0, 18, 0, 0, 0, 107, 0, 0, 227,
	220, 0, 104, 0, 105, 106, 0, 108, 221,
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
	case 18:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:209
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(WhereStr, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(HavingStr, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 19:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:213
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{Expr: yyDollar[4].valExpr}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:217
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 21:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:223
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:227
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, updateList := range yyDollar[7].updateExprs {
				cols = append(cols, updateList.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 23:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:239
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 24:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:245
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:251
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:257
		{
			yyVAL.statement = &DDL{Action: CreateTableStr, IfNotExists: bool(yyDollar[3].boolVal), Table: yyDollar[4].tableName, NewName: yyDollar[4].tableName}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:262
		{
			yyVAL.statement = &DDL{Action: CreateDBStr, IfNotExists: bool(yyDollar[3].boolVal), Database: yyDollar[4].tableIdent}
		}
	case 28:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line sql.y:266
		{
			yyVAL.statement = &DDL{Action: CreateIndexStr, Table: yyDollar[7].tableName, NewName: yyDollar[7].tableName}
		}
	case 29:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:272
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableName, NewName: yyDollar[4].tableName}
		}
	case 30:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:276
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableName, NewName: yyDollar[7].tableName}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:282
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableName, NewName: yyDollar[5].tableName}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:288
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropTableStr, Table: yyDollar[4].tableName, IfExists: exists}
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:296
		{
			yyVAL.statement = &DDL{Action: DropIndexStr, Table: yyDollar[5].tableName, NewName: yyDollar[5].tableName}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:300
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropDBStr, Database: yyDollar[4].tableIdent, IfExists: exists}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:311
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableName, NewName: yyDollar[3].tableName}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:317
		{
			yyVAL.statement = &Xa{}
		}
	case 37:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:323
		{
			yyVAL.statement = &Shard{ShardKey: string(yyDollar[5].bytes)}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:329
		{
			yyVAL.statement = &UseDB{Database: string(yyDollar[2].bytes)}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:335
		{
			yyVAL.statement = &ShowDatabases{}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:339
		{
			yyVAL.statement = &ShowTables{}
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:343
		{
			yyVAL.statement = &ShowTables{Database: yyDollar[4].tableIdent}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:347
		{
			yyVAL.statement = &ShowCreateTable{Table: yyDollar[4].tableName}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:351
		{
			yyVAL.statement = &ShowProcesslist{}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:358
		{
			yyVAL.statement = &Kill{QueryID: string(yyDollar[2].bytes)}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:365
		{
			yyVAL.statement = &Other{}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:369
		{
			yyVAL.statement = &Other{}
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:374
		{
			setAllowComments(yylex, true)
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:378
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:384
		{
			yyVAL.bytes2 = nil
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:388
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:394
		{
			yyVAL.str = UnionStr
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:398
		{
			yyVAL.str = UnionAllStr
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:402
		{
			yyVAL.str = UnionDistinctStr
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:407
		{
			yyVAL.str = ""
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:411
		{
			yyVAL.str = DistinctStr
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:416
		{
			yyVAL.str = ""
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:420
		{
			yyVAL.str = StraightJoinHint
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:426
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:430
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:436
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:440
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:444
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Name: yyDollar[1].tableIdent}}
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:448
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}}
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:454
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:458
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:463
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:467
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:471
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:478
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 71:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:483
		{
			yyVAL.tableExprs = TableExprs{&AliasedTableExpr{Expr: &TableName{Name: NewTableIdent("dual")}}}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:487
		{
			yyVAL.tableExprs = yyDollar[2].tableExprs
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:493
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:497
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:507
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:511
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:515
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:528
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:532
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:536
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:540
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 84:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:545
		{
			yyVAL.empty = struct{}{}
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:547
		{
			yyVAL.empty = struct{}{}
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:550
		{
			yyVAL.tableIdent = NewTableIdent("")
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:554
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:558
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:565
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:571
		{
			yyVAL.str = JoinStr
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:575
		{
			yyVAL.str = JoinStr
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:579
		{
			yyVAL.str = JoinStr
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:583
		{
			yyVAL.str = StraightJoinStr
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:589
		{
			yyVAL.str = LeftJoinStr
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:593
		{
			yyVAL.str = LeftJoinStr
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:597
		{
			yyVAL.str = RightJoinStr
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:601
		{
			yyVAL.str = RightJoinStr
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:607
		{
			yyVAL.str = NaturalJoinStr
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:611
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:621
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:625
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 103:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:630
		{
			yyVAL.indexHints = nil
		}
	case 104:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:634
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:638
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 106:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:642
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:648
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:652
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:657
		{
			yyVAL.boolExpr = nil
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:661
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:668
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:672
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:676
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:680
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:684
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:690
		{
			yyVAL.boolVal = BoolVal(true)
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:694
		{
			yyVAL.boolVal = BoolVal(false)
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:700
		{
			yyVAL.boolExpr = yyDollar[1].boolVal
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:704
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].boolVal}
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:708
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:712
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:716
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:720
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:724
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:728
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:732
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 128:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:736
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 129:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:740
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:744
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:748
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:754
		{
			yyVAL.str = IsNullStr
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:758
		{
			yyVAL.str = IsNotNullStr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:762
		{
			yyVAL.str = IsTrueStr
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:766
		{
			yyVAL.str = IsNotTrueStr
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:770
		{
			yyVAL.str = IsFalseStr
		}
	case 137:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:774
		{
			yyVAL.str = IsNotFalseStr
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:780
		{
			yyVAL.str = EqualStr
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:784
		{
			yyVAL.str = LessThanStr
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:788
		{
			yyVAL.str = GreaterThanStr
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:792
		{
			yyVAL.str = LessEqualStr
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:796
		{
			yyVAL.str = GreaterEqualStr
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:800
		{
			yyVAL.str = NotEqualStr
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:804
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:810
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:814
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:818
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:824
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:830
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:834
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:840
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:844
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:848
		{
			yyVAL.valExpr = yyDollar[1].valTuple
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:852
		{
			yyVAL.valExpr = yyDollar[1].subquery
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:856
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:860
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:864
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:868
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:872
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:876
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:880
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:884
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:888
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:892
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:896
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:900
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONExtractOp, Right: yyDollar[3].valExpr}
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:904
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONUnquoteExtractOp, Right: yyDollar[3].valExpr}
		}
	case 168:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:908
		{
			if num, ok := yyDollar[2].valExpr.(*SQLVal); ok && num.Type == IntVal {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 169:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:916
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
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:930
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:934
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.valExpr = &IntervalExpr{Expr: yyDollar[2].valExpr, Unit: yyDollar[3].colIdent}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:942
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 173:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:946
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 174:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:950
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:954
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 176:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:958
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:962
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:968
		{
			yyVAL.colIdent = NewColIdent("if")
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:972
		{
			yyVAL.colIdent = NewColIdent("current_timestamp")
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:976
		{
			yyVAL.colIdent = NewColIdent("database")
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:980
		{
			yyVAL.colIdent = NewColIdent("mod")
		}
	case 182:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:986
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:991
		{
			yyVAL.valExpr = nil
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:995
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1001
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1005
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 187:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1011
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 188:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.valExpr = nil
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1026
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 191:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1030
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 192:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1034
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1040
		{
			yyVAL.valExpr = NewStrVal(yyDollar[1].bytes)
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1044
		{
			yyVAL.valExpr = NewHexVal(yyDollar[1].bytes)
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1048
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1052
		{
			yyVAL.valExpr = NewFloatVal(yyDollar[1].bytes)
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1056
		{
			yyVAL.valExpr = NewHexNum(yyDollar[1].bytes)
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1060
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1064
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1070
		{
			// TODO(sougou): Deprecate this construct.
			if yyDollar[1].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.valExpr = NewIntVal([]byte("1"))
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1079
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1088
		{
			yyVAL.valExprs = nil
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1097
		{
			yyVAL.boolExpr = nil
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1101
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1106
		{
			yyVAL.orderBy = nil
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1110
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1116
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1120
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 211:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1126
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1131
		{
			yyVAL.str = AscScr
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1135
		{
			yyVAL.str = AscScr
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1139
		{
			yyVAL.str = DescScr
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1144
		{
			yyVAL.limit = nil
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1148
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 217:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1152
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 218:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1156
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr, Offset: yyDollar[4].valExpr}
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1162
		{
			yyVAL.str = ""
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1166
		{
			yyVAL.str = ForUpdateStr
		}
	case 221:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1170
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
	case 222:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1183
		{
			yyVAL.columns = nil
		}
	case 223:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1187
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1193
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 225:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1197
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 226:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1202
		{
			yyVAL.updateExprs = nil
		}
	case 227:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1206
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 228:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1212
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1216
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1222
		{
			yyVAL.values = Values{yyDollar[1].valTuple}
		}
	case 231:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1226
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].valTuple)
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1232
		{
			yyVAL.valTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1238
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 234:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1242
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1248
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 238:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1257
		{
			yyVAL.byt = 0
		}
	case 239:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1259
		{
			yyVAL.byt = 1
		}
	case 240:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1262
		{
			yyVAL.boolVal = false
		}
	case 241:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1264
		{
			yyVAL.boolVal = true
		}
	case 242:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1267
		{
			yyVAL.str = ""
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1269
		{
			yyVAL.str = IgnoreStr
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1273
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1275
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1277
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1279
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1281
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1283
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1286
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1288
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1291
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1293
		{
			yyVAL.empty = struct{}{}
		}
	case 254:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1296
		{
			yyVAL.empty = struct{}{}
		}
	case 255:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1298
		{
			yyVAL.empty = struct{}{}
		}
	case 256:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1302
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 257:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1308
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1314
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 259:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1323
		{
			decNesting(yylex)
		}
	case 260:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1328
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}

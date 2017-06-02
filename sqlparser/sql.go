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
const REPLACE = 57350
const UPDATE = 57351
const DELETE = 57352
const FROM = 57353
const WHERE = 57354
const GROUP = 57355
const HAVING = 57356
const ORDER = 57357
const BY = 57358
const LIMIT = 57359
const OFFSET = 57360
const FOR = 57361
const ALL = 57362
const DISTINCT = 57363
const AS = 57364
const EXISTS = 57365
const ASC = 57366
const DESC = 57367
const INTO = 57368
const DUPLICATE = 57369
const KEY = 57370
const DEFAULT = 57371
const SET = 57372
const LOCK = 57373
const VALUES = 57374
const LAST_INSERT_ID = 57375
const NEXT = 57376
const VALUE = 57377
const JOIN = 57378
const STRAIGHT_JOIN = 57379
const LEFT = 57380
const RIGHT = 57381
const INNER = 57382
const OUTER = 57383
const CROSS = 57384
const NATURAL = 57385
const USE = 57386
const FORCE = 57387
const ON = 57388
const ID = 57389
const HEX = 57390
const STRING = 57391
const INTEGRAL = 57392
const FLOAT = 57393
const HEXNUM = 57394
const VALUE_ARG = 57395
const LIST_ARG = 57396
const COMMENT = 57397
const NULL = 57398
const TRUE = 57399
const FALSE = 57400
const OR = 57401
const AND = 57402
const NOT = 57403
const BETWEEN = 57404
const CASE = 57405
const WHEN = 57406
const THEN = 57407
const ELSE = 57408
const END = 57409
const LE = 57410
const GE = 57411
const NE = 57412
const NULL_SAFE_EQUAL = 57413
const IS = 57414
const LIKE = 57415
const REGEXP = 57416
const IN = 57417
const SHIFT_LEFT = 57418
const SHIFT_RIGHT = 57419
const MOD = 57420
const UNARY = 57421
const INTERVAL = 57422
const JSON_EXTRACT_OP = 57423
const JSON_UNQUOTE_EXTRACT_OP = 57424
const CREATE = 57425
const ALTER = 57426
const DROP = 57427
const RENAME = 57428
const ANALYZE = 57429
const TRUNCATE = 57430
const TABLE = 57431
const INDEX = 57432
const VIEW = 57433
const TO = 57434
const IGNORE = 57435
const IF = 57436
const UNIQUE = 57437
const USING = 57438
const SHOW = 57439
const DESCRIBE = 57440
const EXPLAIN = 57441
const XA = 57442
const PROCESSLIST = 57443
const STATUS = 57444
const PARTITION = 57445
const PARTITIONS = 57446
const HASH = 57447
const ENGINE = 57448
const ENGINES = 57449
const DATABASE = 57450
const DATABASES = 57451
const TABLES = 57452
const KILL = 57453
const CURRENT_TIMESTAMP = 57454
const UNUSED = 57455

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"UNION",
	"SELECT",
	"INSERT",
	"REPLACE",
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
	"TRUNCATE",
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
	"STATUS",
	"PARTITION",
	"PARTITIONS",
	"HASH",
	"ENGINE",
	"ENGINES",
	"DATABASE",
	"DATABASES",
	"TABLES",
	"KILL",
	"CURRENT_TIMESTAMP",
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
	-1, 151,
	96, 267,
	-2, 266,
}

const yyNprod = 271
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1193

var yyAct = [...]int{

	163, 420, 88, 365, 157, 198, 490, 387, 377, 158,
	412, 319, 295, 301, 312, 300, 355, 299, 166, 276,
	259, 144, 155, 283, 258, 3, 222, 303, 193, 188,
	69, 104, 62, 145, 202, 89, 99, 440, 442, 57,
	58, 293, 93, 111, 55, 134, 50, 209, 70, 66,
	90, 71, 53, 473, 68, 472, 65, 67, 59, 471,
	95, 207, 96, 61, 60, 51, 75, 149, 56, 261,
	262, 422, 73, 74, 394, 84, 286, 226, 125, 239,
	118, 91, 211, 496, 97, 151, 229, 294, 102, 103,
	240, 241, 242, 243, 244, 245, 239, 105, 106, 107,
	108, 110, 441, 112, 128, 114, 452, 356, 356, 410,
	139, 123, 130, 121, 458, 126, 86, 296, 390, 446,
	129, 154, 326, 307, 133, 90, 135, 191, 90, 190,
	228, 227, 206, 208, 204, 199, 324, 325, 323, 142,
	227, 143, 228, 227, 457, 91, 229, 225, 224, 98,
	187, 228, 227, 72, 205, 229, 140, 454, 229, 154,
	154, 210, 212, 195, 91, 213, 322, 229, 86, 267,
	268, 92, 371, 270, 215, 216, 217, 269, 238, 237,
	246, 247, 240, 241, 242, 243, 244, 245, 239, 151,
	90, 281, 279, 190, 313, 315, 316, 214, 256, 314,
	154, 388, 132, 290, 242, 243, 244, 245, 239, 101,
	273, 498, 296, 278, 395, 396, 397, 274, 223, 100,
	282, 255, 257, 94, 224, 154, 305, 309, 64, 86,
	304, 390, 288, 154, 154, 140, 291, 320, 86, 91,
	296, 321, 119, 219, 296, 120, 91, 297, 225, 317,
	306, 330, 194, 298, 42, 169, 168, 170, 171, 172,
	173, 22, 285, 174, 344, 296, 127, 348, 375, 296,
	342, 343, 345, 154, 154, 413, 90, 127, 358, 220,
	349, 352, 360, 361, 346, 347, 344, 309, 127, 350,
	353, 363, 296, 362, 357, 310, 311, 140, 296, 278,
	459, 415, 140, 467, 375, 86, 364, 275, 189, 413,
	91, 305, 289, 137, 150, 304, 219, 369, 113, 54,
	370, 435, 470, 393, 140, 140, 436, 196, 372, 320,
	116, 433, 469, 321, 432, 399, 434, 437, 431, 383,
	384, 398, 238, 237, 246, 247, 240, 241, 242, 243,
	244, 245, 239, 260, 80, 186, 185, 154, 263, 264,
	265, 266, 154, 405, 416, 82, 407, 79, 417, 414,
	124, 409, 476, 406, 455, 423, 122, 22, 272, 83,
	200, 305, 305, 305, 305, 304, 304, 304, 304, 43,
	428, 427, 430, 429, 438, 447, 443, 287, 379, 382,
	383, 384, 380, 277, 381, 385, 445, 136, 468, 392,
	488, 448, 366, 45, 46, 47, 48, 49, 150, 451,
	456, 361, 489, 447, 411, 76, 77, 184, 426, 318,
	367, 154, 327, 328, 329, 183, 331, 332, 333, 334,
	335, 336, 337, 338, 339, 340, 341, 466, 464, 63,
	284, 379, 382, 383, 384, 380, 463, 381, 385, 425,
	197, 479, 374, 194, 109, 87, 150, 150, 480, 495,
	486, 22, 42, 154, 154, 44, 1, 483, 484, 485,
	391, 386, 221, 201, 491, 491, 491, 90, 52, 494,
	292, 492, 493, 465, 203, 182, 280, 501, 487, 502,
	460, 419, 503, 497, 164, 499, 500, 424, 85, 373,
	408, 271, 354, 165, 156, 359, 117, 85, 230, 287,
	152, 85, 85, 400, 401, 402, 246, 247, 240, 241,
	242, 243, 244, 245, 239, 481, 482, 439, 378, 376,
	302, 218, 147, 404, 85, 115, 78, 41, 85, 81,
	150, 20, 19, 85, 18, 131, 17, 85, 16, 85,
	287, 15, 138, 21, 14, 418, 421, 13, 12, 11,
	141, 10, 85, 9, 85, 8, 148, 7, 6, 5,
	4, 2, 0, 85, 0, 0, 192, 0, 0, 0,
	0, 449, 0, 0, 0, 85, 0, 0, 85, 0,
	0, 351, 450, 167, 0, 0, 0, 0, 0, 453,
	238, 237, 246, 247, 240, 241, 242, 243, 244, 245,
	239, 0, 0, 0, 0, 287, 0, 140, 0, 296,
	151, 169, 168, 170, 171, 172, 173, 0, 0, 174,
	180, 181, 0, 85, 153, 0, 179, 0, 474, 0,
	0, 0, 0, 475, 0, 0, 477, 478, 421, 461,
	462, 0, 0, 0, 0, 0, 159, 160, 146, 167,
	0, 178, 0, 161, 0, 162, 0, 0, 0, 0,
	148, 85, 0, 0, 0, 0, 0, 308, 0, 0,
	175, 0, 0, 140, 0, 296, 151, 169, 168, 170,
	171, 172, 173, 0, 177, 174, 180, 181, 176, 0,
	153, 0, 179, 0, 0, 0, 0, 238, 237, 246,
	247, 240, 241, 242, 243, 244, 245, 239, 148, 148,
	0, 0, 159, 160, 146, 0, 0, 178, 0, 161,
	0, 162, 0, 0, 0, 0, 0, 368, 0, 0,
	85, 0, 0, 85, 167, 0, 175, 0, 0, 0,
	0, 0, 0, 0, 389, 0, 85, 0, 0, 0,
	177, 0, 0, 0, 176, 0, 0, 0, 140, 403,
	0, 151, 169, 168, 170, 171, 172, 173, 0, 0,
	174, 180, 181, 0, 0, 153, 0, 179, 238, 237,
	246, 247, 240, 241, 242, 243, 244, 245, 239, 0,
	0, 0, 148, 0, 0, 0, 22, 159, 160, 146,
	0, 0, 178, 0, 161, 0, 162, 0, 0, 0,
	0, 0, 0, 167, 0, 0, 85, 85, 85, 85,
	0, 175, 0, 0, 0, 0, 0, 0, 0, 389,
	0, 0, 444, 0, 0, 177, 0, 140, 0, 176,
	151, 169, 168, 170, 171, 172, 173, 0, 0, 174,
	180, 181, 0, 0, 153, 0, 179, 0, 0, 0,
	167, 238, 237, 246, 247, 240, 241, 242, 243, 244,
	245, 239, 0, 0, 0, 0, 159, 160, 0, 0,
	0, 178, 0, 161, 140, 162, 0, 151, 169, 168,
	170, 171, 172, 173, 0, 0, 174, 180, 181, 0,
	175, 153, 0, 179, 237, 246, 247, 240, 241, 242,
	243, 244, 245, 239, 177, 0, 0, 0, 176, 0,
	0, 0, 0, 159, 160, 0, 0, 0, 178, 0,
	161, 140, 162, 0, 151, 169, 168, 170, 171, 172,
	173, 0, 22, 174, 180, 181, 0, 175, 0, 0,
	179, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 177, 0, 0, 0, 176, 0, 0, 0, 0,
	159, 160, 0, 0, 0, 178, 0, 161, 0, 162,
	0, 0, 0, 140, 0, 0, 151, 169, 168, 170,
	171, 172, 173, 0, 175, 174, 0, 0, 0, 0,
	0, 0, 179, 0, 0, 0, 0, 0, 177, 0,
	0, 0, 176, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 159, 160, 0, 0, 0, 178, 0, 161,
	140, 162, 0, 151, 169, 168, 170, 171, 172, 173,
	0, 0, 174, 0, 0, 0, 175, 0, 0, 179,
	0, 22, 23, 24, 25, 26, 0, 0, 0, 0,
	177, 0, 0, 0, 176, 0, 0, 0, 0, 159,
	160, 0, 0, 0, 178, 27, 161, 0, 162, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 36,
	0, 0, 0, 175, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 232, 235, 177, 0, 0,
	0, 176, 248, 249, 250, 251, 252, 253, 254, 236,
	233, 234, 231, 238, 237, 246, 247, 240, 241, 242,
	243, 244, 245, 239, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 28, 29, 31, 30, 33, 32,
	0, 0, 0, 0, 0, 0, 0, 0, 37, 40,
	39, 34, 0, 0, 35, 0, 0, 0, 0, 0,
	0, 0, 38,
}
var yyPact = [...]int{

	1065, -1000, -1000, 467, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -59, -65,
	-37, -66, -41, -42, -1000, 433, 178, -69, 100, -1000,
	-1000, 465, 405, 333, -1000, -65, 353, 118, 454, 114,
	-68, -68, -46, -1000, -43, -1000, 118, -74, 169, -74,
	118, 118, -1000, -90, -1000, -1000, -1000, 453, -1000, -62,
	-1000, 272, -1000, -1000, -1000, -1000, -1000, -1000, 293, 189,
	-1000, 55, 350, 118, 340, -18, -1000, 118, 218, -1000,
	33, -1000, 118, 48, 118, 152, 118, -63, 118, 384,
	267, 118, -1000, -1000, 250, -1000, -1000, -1000, -1000, 118,
	-1000, 118, -1000, 118, -1000, 731, -1000, 416, -1000, 324,
	323, -1000, 118, 278, 114, 118, 451, 114, 1003, 250,
	357, -1000, -78, 32, 118, -1000, -1000, 118, -1000, 147,
	-1000, -1000, -1000, -1000, 268, -1000, -1000, 196, -19, 80,
	1061, -1000, -1000, 857, 810, -1000, -28, -1000, -1000, 1003,
	1003, 1003, 1003, 250, 250, -1000, -1000, 250, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 1003,
	-1000, -1000, 118, -1000, -1000, -1000, -1000, 277, 371, 114,
	114, 240, -1000, 435, 857, -1000, 799, -20, 956, -1000,
	-1000, 266, 114, -1000, -67, 16, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 191, -1000, -1000, -1000, 451, 731,
	188, -1000, -1000, 95, -1000, -1000, 35, 857, 857, 135,
	904, 109, 57, 1003, 1003, 1003, 135, 1003, 1003, 1003,
	1003, 1003, 1003, 1003, 1003, 1003, 1003, 1003, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 8, 1061, 68, 249, 216,
	1061, 204, 204, -1000, -1000, -1000, 260, 580, 646, -1000,
	465, 41, 799, -1000, 371, 114, -1000, 250, 467, 218,
	243, -1000, 435, 395, 414, 80, 139, 799, -1000, 118,
	-1000, -1000, 118, -1000, 122, -1000, -1000, 449, -1000, 256,
	415, -1000, -1000, 179, 387, 255, -1000, -1000, -22, -1000,
	8, 77, -1000, -1000, 155, -1000, -1000, -1000, 799, -1000,
	956, -1000, -1000, 109, 1003, 1003, 1003, 799, 799, 716,
	-1000, 442, 841, -1000, 116, 116, -13, -13, -13, -13,
	4, 4, -1000, -1000, 1003, -1000, -1000, -1000, -1000, -1000,
	195, 731, -1000, 195, 40, -1000, 857, 263, 229, 253,
	-1000, 1003, -1000, 114, 395, -1000, 1003, 1003, -25, 250,
	-1000, -1000, -1000, 445, 412, 188, 188, 188, 188, -1000,
	302, 298, -1000, 295, 285, 301, -7, -1000, 66, -1000,
	-1000, 118, -1000, 220, 31, -1000, -1000, -1000, 216, -1000,
	799, 799, 528, 1003, 799, -1000, 195, -1000, 36, -1000,
	1003, 89, -1000, 347, -1000, 250, -1000, -1000, 96, 252,
	-1000, 635, 114, -1000, 435, 857, 1003, 415, 257, 362,
	-1000, -1000, -1000, -1000, 296, -1000, 286, -1000, -1000, -1000,
	-47, -51, -53, -1000, -1000, -1000, -1000, -1000, -1000, 1003,
	799, -1000, -1000, 799, 1003, 344, -1000, 1003, 1003, 1003,
	-1000, -1000, -1000, -1000, 395, 80, 238, 857, 857, -1000,
	-1000, 250, 250, 250, 799, 799, 461, 799, 799, -1000,
	391, 80, 80, 114, 114, 114, 114, -1000, 460, 2,
	163, -1000, 163, 163, 218, -1000, 114, -1000, 114, -1000,
	-1000, 114, -1000, -1000,
}
var yyPgo = [...]int{

	0, 581, 24, 580, 579, 578, 577, 575, 573, 571,
	569, 568, 567, 564, 563, 561, 558, 556, 554, 552,
	551, 389, 549, 547, 546, 545, 21, 33, 542, 541,
	17, 15, 13, 540, 539, 8, 538, 27, 537, 6,
	28, 67, 520, 18, 518, 19, 22, 198, 516, 14,
	11, 20, 515, 4, 9, 514, 513, 512, 16, 511,
	510, 509, 507, 504, 23, 501, 1, 500, 3, 498,
	29, 496, 10, 2, 35, 495, 319, 149, 494, 490,
	488, 483, 0, 26, 482, 460, 7, 481, 480, 32,
	171, 476, 5, 12, 475,
}
var yyR1 = [...]int{

	0, 91, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 2, 3, 3, 4, 4, 5, 6,
	7, 8, 8, 8, 9, 9, 9, 10, 11, 11,
	11, 12, 13, 15, 16, 17, 18, 18, 18, 18,
	18, 18, 18, 18, 19, 20, 14, 94, 21, 22,
	22, 23, 23, 23, 24, 24, 25, 25, 26, 26,
	27, 27, 27, 27, 28, 28, 84, 84, 84, 83,
	83, 29, 29, 30, 30, 31, 31, 32, 32, 32,
	33, 33, 33, 33, 88, 88, 87, 87, 87, 86,
	86, 34, 34, 34, 34, 35, 35, 35, 35, 36,
	36, 37, 37, 38, 38, 38, 38, 39, 39, 40,
	40, 41, 41, 41, 41, 41, 41, 43, 43, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 49, 49, 49, 49, 49, 49, 44, 44,
	44, 44, 44, 44, 44, 50, 50, 50, 54, 51,
	51, 47, 47, 47, 47, 47, 47, 47, 47, 47,
	47, 47, 47, 47, 47, 47, 47, 47, 47, 47,
	47, 47, 47, 47, 47, 47, 47, 47, 63, 63,
	63, 63, 56, 59, 59, 57, 57, 58, 60, 60,
	55, 55, 55, 46, 46, 46, 46, 46, 46, 46,
	48, 48, 48, 61, 61, 62, 62, 64, 64, 65,
	65, 66, 67, 67, 67, 68, 68, 68, 68, 69,
	69, 69, 70, 70, 71, 71, 72, 72, 45, 45,
	52, 52, 53, 73, 73, 74, 75, 75, 77, 77,
	90, 90, 76, 76, 78, 78, 78, 78, 78, 78,
	79, 79, 80, 80, 81, 81, 82, 85, 92, 93,
	89,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 12, 6, 3, 8, 8, 6, 6, 8, 7,
	3, 6, 4, 9, 6, 7, 7, 5, 4, 5,
	4, 3, 3, 2, 7, 3, 3, 3, 3, 3,
	5, 5, 3, 5, 3, 2, 2, 0, 2, 0,
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

	-1000, -91, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -12, -13, -15, -16, -17, -18, -19,
	-20, -14, 6, 7, 8, 9, 10, 30, 99, 100,
	102, 101, 104, 103, 116, 119, 44, 113, 127, 115,
	114, -23, 5, -21, -94, -21, -21, -21, -21, -21,
	105, 124, -80, 111, -76, 109, 105, 105, 106, 124,
	105, 105, -89, 16, 50, 125, 118, 126, 123, 99,
	117, 120, 53, -89, -89, -2, 20, 21, -24, 34,
	21, -22, -76, 26, -37, -85, 50, 11, -73, -74,
	-82, 50, -90, 110, -90, 106, 105, -37, -77, 110,
	50, -77, -37, -37, 121, -89, -89, -89, -89, 11,
	-89, 105, -89, 46, -89, -25, 37, -48, -82, 53,
	56, 58, 26, -37, 30, 96, -37, 48, 71, -37,
	64, -85, 50, -37, 108, -37, 23, 46, -85, -92,
	47, -85, -37, -37, -26, -27, 88, -28, -85, -41,
	-47, 50, -42, 64, -92, -46, -55, -53, -54, 86,
	87, 93, 95, -82, -63, -56, -43, 23, 52, 51,
	53, 54, 55, 56, 59, 110, 128, 124, 91, 66,
	60, 61, -75, 19, 11, 32, 32, -37, -70, 30,
	-92, -73, -85, -40, 12, -74, -47, -85, -92, -92,
	23, -81, 112, -78, 102, 122, 100, 29, 101, 15,
	129, 50, -37, -37, 50, -89, -89, -89, -29, 48,
	11, -84, -83, 22, -82, 52, 96, 63, 62, 78,
	-44, 81, 64, 79, 80, 65, 78, 83, 82, 92,
	86, 87, 88, 89, 90, 91, 84, 85, 71, 72,
	73, 74, 75, 76, 77, -41, -47, -41, -2, -51,
	-47, 97, 98, -47, -47, -47, -47, -92, -92, -54,
	-92, -59, -47, -37, -70, 30, -45, 32, -2, -73,
	-71, -82, -40, -64, 15, -41, 96, -47, -89, 46,
	-82, -89, -79, 108, 71, -93, 49, -40, -27, -30,
	-31, -32, -33, -37, -54, -92, -83, 88, -85, -82,
	-41, -41, -49, 59, 64, 60, 61, -43, -47, -50,
	-92, -54, 57, 81, 79, 80, 65, -47, -47, -47,
	-49, -47, -47, -47, -47, -47, -47, -47, -47, -47,
	-47, -47, -93, -93, 48, -93, -46, -46, -82, -93,
	-26, 21, -93, -26, -57, -58, 67, -45, -73, -52,
	-53, -92, -93, 48, -64, -68, 17, 16, -85, -37,
	-37, 50, -89, -61, 13, 48, -34, -35, -36, 36,
	40, 42, 37, 38, 39, 43, -87, -86, 22, -85,
	52, -88, 22, -30, 96, 59, 60, 61, -51, -50,
	-47, -47, -47, 63, -47, -93, -26, -93, -60, -58,
	69, -41, -72, 46, -72, 48, -82, -68, -47, -65,
	-66, -47, 96, -92, -62, 14, 16, -31, -32, -31,
	-32, 36, 36, 36, 41, 36, 41, 36, -35, -38,
	44, 109, 45, -86, -85, -93, 88, -82, -93, 63,
	-47, -93, 70, -47, 68, 27, -53, 48, 18, 48,
	-67, 24, 25, -89, -64, -41, -51, 46, 46, 36,
	36, 106, 106, 106, -47, -47, 28, -47, -47, -66,
	-68, -41, -41, -92, -92, -92, 9, -69, 19, 31,
	-39, -82, -39, -39, -73, 9, 81, -93, 48, -93,
	-93, -82, -82, -82,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 20, 57, 57, 57, 57, 57, 57, 262, 252,
	0, 0, 0, 0, 270, 0, 0, 0, 0, 270,
	270, 0, 61, 64, 59, 252, 0, 0, 0, 0,
	250, 250, 0, 263, 0, 253, 0, 248, 0, 248,
	0, 0, 43, 0, 270, 270, 270, 270, 270, 0,
	270, 0, 270, 55, 56, 23, 62, 63, 66, 0,
	65, 58, 0, 0, 0, 111, 267, 0, 30, 243,
	0, 266, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 41, 42, 0, 45, 46, 47, 48, 0,
	49, 0, 52, 0, 54, 0, 67, 0, 210, 0,
	0, 60, 0, 232, 0, 0, 119, 0, 0, 0,
	0, 32, 264, 0, 0, 38, 249, 0, 40, 0,
	268, 270, 270, 270, 81, 68, 70, 76, 0, 74,
	75, -2, 121, 0, 0, 161, 162, 163, 164, 0,
	0, 0, 0, 200, 0, 187, 129, 0, 203, 204,
	205, 206, 207, 208, 209, 188, 189, 190, 191, 193,
	127, 128, 0, 246, 247, 211, 212, 232, 0, 0,
	0, 119, 112, 217, 0, 244, 245, 0, 0, 270,
	251, 0, 0, 270, 260, 0, 254, 255, 256, 257,
	258, 259, 37, 39, 0, 50, 51, 53, 119, 0,
	0, 71, 77, 0, 79, 80, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 148, 149,
	150, 151, 152, 153, 154, 124, 0, 0, 0, 0,
	159, 0, 0, 178, 179, 180, 0, 0, 0, 141,
	0, 0, 194, 22, 0, 0, 26, 0, 239, 27,
	0, 234, 217, 225, 0, 120, 0, 159, 31, 0,
	265, 34, 0, 261, 0, 270, 269, 213, 69, 82,
	83, 85, 86, 96, 94, 0, 78, 72, 0, 201,
	122, 123, 126, 142, 0, 144, 146, 130, 131, 132,
	0, 156, 157, 0, 0, 0, 0, 134, 136, 0,
	140, 165, 166, 167, 168, 169, 170, 171, 172, 173,
	174, 175, 125, 158, 0, 242, 176, 177, 181, 182,
	0, 0, 185, 0, 198, 195, 0, 236, 236, 238,
	240, 0, 233, 0, 225, 29, 0, 0, 0, 0,
	35, 36, 44, 215, 0, 0, 0, 0, 0, 101,
	0, 0, 104, 0, 0, 0, 113, 97, 0, 99,
	100, 0, 95, 0, 0, 143, 145, 147, 0, 133,
	135, 137, 0, 0, 160, 183, 0, 186, 0, 196,
	0, 0, 24, 0, 25, 0, 235, 28, 226, 218,
	219, 222, 0, 270, 217, 0, 0, 84, 90, 0,
	93, 102, 103, 105, 0, 107, 0, 109, 110, 87,
	0, 0, 0, 98, 88, 89, 73, 202, 155, 0,
	138, 184, 192, 199, 0, 0, 241, 0, 0, 0,
	221, 223, 224, 33, 225, 216, 214, 0, 0, 106,
	108, 0, 0, 0, 139, 197, 0, 227, 228, 220,
	229, 91, 92, 0, 0, 0, 0, 21, 0, 0,
	0, 117, 0, 0, 237, 230, 0, 114, 0, 115,
	116, 0, 118, 231,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 90, 83, 3,
	47, 49, 88, 86, 48, 87, 96, 89, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	72, 71, 73, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 92, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 82, 3, 93,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 74, 75, 76, 77,
	78, 79, 80, 81, 84, 85, 91, 94, 95, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 113, 114, 115, 116, 117,
	118, 119, 120, 121, 122, 123, 124, 125, 126, 127,
	128, 129,
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
	case 21:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:212
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(WhereStr, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(HavingStr, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:216
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{Expr: yyDollar[4].valExpr}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:220
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:226
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:230
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, updateList := range yyDollar[7].updateExprs {
				cols = append(cols, updateList.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:242
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:246
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, updateList := range yyDollar[6].updateExprs {
				cols = append(cols, updateList.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 28:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:258
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 29:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:264
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:270
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:276
		{
			yyVAL.statement = &DDL{Action: CreateTableStr, IfNotExists: bool(yyDollar[3].boolVal), Table: yyDollar[4].tableName, NewName: yyDollar[4].tableName}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:281
		{
			yyVAL.statement = &DDL{Action: CreateDBStr, IfNotExists: bool(yyDollar[3].boolVal), Database: yyDollar[4].tableIdent}
		}
	case 33:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line sql.y:285
		{
			yyVAL.statement = &DDL{Action: CreateIndexStr, Table: yyDollar[7].tableName, NewName: yyDollar[7].tableName}
		}
	case 34:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:291
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableName, NewName: yyDollar[4].tableName}
		}
	case 35:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:295
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableName, NewName: yyDollar[7].tableName}
		}
	case 36:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:299
		{
			yyVAL.statement = &DDL{Action: AlterEngineStr, Table: yyDollar[4].tableName, Engine: string(yyDollar[7].bytes)}
		}
	case 37:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:305
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableName, NewName: yyDollar[5].tableName}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:311
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropTableStr, Table: yyDollar[4].tableName, IfExists: exists}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:319
		{
			yyVAL.statement = &DDL{Action: DropIndexStr, Table: yyDollar[5].tableName, NewName: yyDollar[5].tableName}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:323
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropDBStr, Database: yyDollar[4].tableIdent, IfExists: exists}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:333
		{
			yyVAL.statement = &DDL{Action: TruncateTableStr, Table: yyDollar[3].tableName}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:339
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableName, NewName: yyDollar[3].tableName}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:345
		{
			yyVAL.statement = &Xa{}
		}
	case 44:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:351
		{
			yyVAL.statement = &Shard{ShardKey: string(yyDollar[5].bytes)}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:357
		{
			yyVAL.statement = &UseDB{Database: string(yyDollar[2].bytes)}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:363
		{
			yyVAL.statement = &ShowDatabases{}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:367
		{
			yyVAL.statement = &ShowStatus{}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:371
		{
			yyVAL.statement = &ShowTables{}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:375
		{
			yyVAL.statement = &ShowEngines{}
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:379
		{
			yyVAL.statement = &ShowTables{Database: yyDollar[4].tableIdent}
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:383
		{
			yyVAL.statement = &ShowCreateTable{Table: yyDollar[4].tableName}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:387
		{
			yyVAL.statement = &ShowProcesslist{}
		}
	case 53:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:391
		{
			yyVAL.statement = &ShowPartitions{Table: yyDollar[4].tableName}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:397
		{
			yyVAL.statement = &Kill{QueryID: string(yyDollar[2].bytes)}
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:403
		{
			yyVAL.statement = &Explain{}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:409
		{
			yyVAL.statement = &Other{}
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:414
		{
			setAllowComments(yylex, true)
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:418
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 59:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:424
		{
			yyVAL.bytes2 = nil
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:428
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:434
		{
			yyVAL.str = UnionStr
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:438
		{
			yyVAL.str = UnionAllStr
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:442
		{
			yyVAL.str = UnionDistinctStr
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:447
		{
			yyVAL.str = ""
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:451
		{
			yyVAL.str = DistinctStr
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:456
		{
			yyVAL.str = ""
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:460
		{
			yyVAL.str = StraightJoinHint
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:466
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:470
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:476
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:480
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:484
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Name: yyDollar[1].tableIdent}}
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:488
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:494
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:498
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 76:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:503
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:507
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:511
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:518
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:523
		{
			yyVAL.tableExprs = TableExprs{&AliasedTableExpr{Expr: &TableName{Name: NewTableIdent("dual")}}}
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:527
		{
			yyVAL.tableExprs = yyDollar[2].tableExprs
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:533
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:537
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:547
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:551
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:555
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:568
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:572
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:576
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:580
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 94:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:585
		{
			yyVAL.empty = struct{}{}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:587
		{
			yyVAL.empty = struct{}{}
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:590
		{
			yyVAL.tableIdent = NewTableIdent("")
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:594
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:598
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:605
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:611
		{
			yyVAL.str = JoinStr
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:615
		{
			yyVAL.str = JoinStr
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:619
		{
			yyVAL.str = JoinStr
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:623
		{
			yyVAL.str = StraightJoinStr
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:629
		{
			yyVAL.str = LeftJoinStr
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:633
		{
			yyVAL.str = LeftJoinStr
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:637
		{
			yyVAL.str = RightJoinStr
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:641
		{
			yyVAL.str = RightJoinStr
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:647
		{
			yyVAL.str = NaturalJoinStr
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:651
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:661
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:665
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 113:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:670
		{
			yyVAL.indexHints = nil
		}
	case 114:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:674
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 115:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:678
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 116:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:682
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:688
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:692
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 119:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:697
		{
			yyVAL.boolExpr = nil
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:701
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:708
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:712
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:716
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:720
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:724
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:730
		{
			yyVAL.boolVal = BoolVal(true)
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:734
		{
			yyVAL.boolVal = BoolVal(false)
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:740
		{
			yyVAL.boolExpr = yyDollar[1].boolVal
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:744
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].boolVal}
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:748
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:752
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 133:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:756
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:760
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 135:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:764
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:768
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:772
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:776
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:780
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:784
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:788
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:794
		{
			yyVAL.str = IsNullStr
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:798
		{
			yyVAL.str = IsNotNullStr
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:802
		{
			yyVAL.str = IsTrueStr
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:806
		{
			yyVAL.str = IsNotTrueStr
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:810
		{
			yyVAL.str = IsFalseStr
		}
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:814
		{
			yyVAL.str = IsNotFalseStr
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:820
		{
			yyVAL.str = EqualStr
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:824
		{
			yyVAL.str = LessThanStr
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:828
		{
			yyVAL.str = GreaterThanStr
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:832
		{
			yyVAL.str = LessEqualStr
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:836
		{
			yyVAL.str = GreaterEqualStr
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:840
		{
			yyVAL.str = NotEqualStr
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:844
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:850
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:854
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:858
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:864
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:870
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:874
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:880
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:884
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:888
		{
			yyVAL.valExpr = yyDollar[1].valTuple
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:892
		{
			yyVAL.valExpr = yyDollar[1].subquery
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:896
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:900
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:904
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:908
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:912
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:916
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:920
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:924
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:928
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:932
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:936
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:940
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONExtractOp, Right: yyDollar[3].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:944
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONUnquoteExtractOp, Right: yyDollar[3].valExpr}
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:948
		{
			if num, ok := yyDollar[2].valExpr.(*SQLVal); ok && num.Type == IntVal {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:956
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
	case 180:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:970
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:974
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.valExpr = &IntervalExpr{Expr: yyDollar[2].valExpr, Unit: yyDollar[3].colIdent}
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:982
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 183:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:986
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 184:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:990
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:994
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 186:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:998
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1002
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1008
		{
			yyVAL.colIdent = NewColIdent("if")
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1012
		{
			yyVAL.colIdent = NewColIdent("current_timestamp")
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.colIdent = NewColIdent("database")
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.colIdent = NewColIdent("mod")
		}
	case 192:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1026
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1031
		{
			yyVAL.valExpr = nil
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1041
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1045
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 197:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1051
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1056
		{
			yyVAL.valExpr = nil
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1060
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1066
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1070
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 202:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1074
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1080
		{
			yyVAL.valExpr = NewStrVal(yyDollar[1].bytes)
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1084
		{
			yyVAL.valExpr = NewHexVal(yyDollar[1].bytes)
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1088
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.valExpr = NewFloatVal(yyDollar[1].bytes)
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1096
		{
			yyVAL.valExpr = NewHexNum(yyDollar[1].bytes)
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1100
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1104
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1110
		{
			// TODO(sougou): Deprecate this construct.
			if yyDollar[1].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.valExpr = NewIntVal([]byte("1"))
		}
	case 211:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1119
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 212:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1123
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1128
		{
			yyVAL.valExprs = nil
		}
	case 214:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1132
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1137
		{
			yyVAL.boolExpr = nil
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1141
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1146
		{
			yyVAL.orderBy = nil
		}
	case 218:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1150
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1156
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 220:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1160
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1166
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 222:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1171
		{
			yyVAL.str = AscScr
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1175
		{
			yyVAL.str = AscScr
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1179
		{
			yyVAL.str = DescScr
		}
	case 225:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1184
		{
			yyVAL.limit = nil
		}
	case 226:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1188
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 227:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1192
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 228:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1196
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr, Offset: yyDollar[4].valExpr}
		}
	case 229:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1202
		{
			yyVAL.str = ""
		}
	case 230:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1206
		{
			yyVAL.str = ForUpdateStr
		}
	case 231:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1210
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
	case 232:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1223
		{
			yyVAL.columns = nil
		}
	case 233:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1227
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1233
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1237
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1242
		{
			yyVAL.updateExprs = nil
		}
	case 237:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1246
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 238:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1252
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1256
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1262
		{
			yyVAL.values = Values{yyDollar[1].valTuple}
		}
	case 241:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1266
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].valTuple)
		}
	case 242:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1272
		{
			yyVAL.valTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1278
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 244:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1282
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 245:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1288
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 248:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1297
		{
			yyVAL.byt = 0
		}
	case 249:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1299
		{
			yyVAL.byt = 1
		}
	case 250:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1302
		{
			yyVAL.boolVal = false
		}
	case 251:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1304
		{
			yyVAL.boolVal = true
		}
	case 252:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1307
		{
			yyVAL.str = ""
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1309
		{
			yyVAL.str = IgnoreStr
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1313
		{
			yyVAL.empty = struct{}{}
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1315
		{
			yyVAL.empty = struct{}{}
		}
	case 256:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1317
		{
			yyVAL.empty = struct{}{}
		}
	case 257:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1319
		{
			yyVAL.empty = struct{}{}
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1321
		{
			yyVAL.empty = struct{}{}
		}
	case 259:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1323
		{
			yyVAL.empty = struct{}{}
		}
	case 260:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1326
		{
			yyVAL.empty = struct{}{}
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1328
		{
			yyVAL.empty = struct{}{}
		}
	case 262:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1331
		{
			yyVAL.empty = struct{}{}
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1333
		{
			yyVAL.empty = struct{}{}
		}
	case 264:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1336
		{
			yyVAL.empty = struct{}{}
		}
	case 265:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1338
		{
			yyVAL.empty = struct{}{}
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1342
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 267:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1348
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1354
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 269:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1363
		{
			decNesting(yylex)
		}
	case 270:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1368
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}

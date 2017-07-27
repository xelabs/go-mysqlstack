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
const QUERYZ = 57445
const TXNZ = 57446
const VERSIONS = 57447
const PARTITION = 57448
const PARTITIONS = 57449
const HASH = 57450
const ENGINE = 57451
const ENGINES = 57452
const DATABASE = 57453
const DATABASES = 57454
const TABLES = 57455
const KILL = 57456
const CURRENT_TIMESTAMP = 57457
const UNUSED = 57458

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
	"QUERYZ",
	"TXNZ",
	"VERSIONS",
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
	-1, 173,
	96, 270,
	-2, 269,
}

const yyNprod = 274
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1195

var yyAct = [...]int{

	161, 185, 91, 118, 433, 172, 498, 307, 297, 409,
	399, 251, 155, 342, 359, 426, 340, 336, 290, 190,
	62, 341, 250, 3, 352, 259, 181, 180, 205, 200,
	153, 92, 57, 58, 50, 107, 212, 55, 457, 459,
	53, 102, 96, 305, 141, 484, 483, 482, 98, 115,
	93, 99, 61, 60, 59, 56, 51, 241, 242, 416,
	76, 77, 219, 389, 78, 229, 228, 237, 238, 231,
	232, 233, 234, 235, 236, 230, 217, 156, 120, 230,
	70, 263, 249, 125, 132, 108, 109, 110, 111, 112,
	114, 94, 116, 504, 266, 173, 121, 221, 71, 66,
	73, 74, 67, 458, 72, 306, 135, 69, 438, 65,
	68, 337, 147, 146, 137, 173, 165, 164, 166, 167,
	168, 169, 128, 75, 170, 192, 193, 410, 189, 463,
	264, 178, 93, 348, 203, 93, 202, 265, 264, 151,
	94, 179, 209, 440, 337, 266, 393, 216, 218, 214,
	101, 157, 158, 266, 95, 89, 177, 412, 159, 89,
	160, 233, 234, 235, 236, 230, 207, 247, 248, 225,
	226, 227, 215, 366, 384, 174, 147, 471, 472, 220,
	308, 260, 265, 264, 261, 173, 362, 364, 365, 363,
	281, 283, 176, 224, 189, 189, 175, 286, 266, 417,
	418, 419, 93, 295, 293, 202, 97, 42, 299, 94,
	104, 262, 189, 302, 237, 238, 231, 232, 233, 234,
	235, 236, 230, 292, 22, 139, 353, 355, 356, 288,
	300, 354, 296, 103, 303, 229, 228, 237, 238, 231,
	232, 233, 234, 235, 236, 230, 94, 324, 240, 126,
	330, 308, 127, 189, 189, 325, 328, 147, 332, 334,
	89, 261, 189, 346, 330, 147, 350, 351, 89, 285,
	189, 189, 322, 323, 360, 326, 329, 89, 239, 412,
	427, 308, 134, 339, 338, 64, 347, 357, 506, 308,
	93, 371, 373, 282, 265, 264, 206, 376, 370, 333,
	308, 380, 377, 257, 375, 379, 94, 372, 262, 134,
	266, 292, 229, 228, 237, 238, 231, 232, 233, 234,
	235, 236, 230, 397, 308, 256, 308, 333, 385, 378,
	308, 469, 134, 189, 386, 345, 429, 388, 397, 394,
	256, 147, 478, 189, 427, 301, 361, 144, 117, 452,
	54, 123, 346, 392, 453, 387, 231, 232, 233, 234,
	235, 236, 230, 415, 401, 404, 405, 406, 402, 360,
	403, 407, 420, 481, 479, 289, 201, 450, 421, 430,
	480, 83, 451, 431, 454, 449, 405, 406, 435, 428,
	437, 448, 147, 147, 82, 436, 85, 496, 198, 22,
	197, 131, 486, 346, 346, 346, 346, 467, 129, 497,
	86, 210, 445, 152, 447, 143, 186, 437, 455, 444,
	460, 446, 414, 462, 345, 291, 119, 443, 464, 208,
	401, 404, 405, 406, 402, 376, 403, 407, 79, 80,
	381, 361, 468, 43, 476, 63, 298, 442, 189, 396,
	475, 243, 244, 245, 246, 477, 473, 165, 164, 166,
	167, 168, 169, 206, 113, 170, 252, 45, 46, 47,
	48, 49, 254, 90, 487, 345, 345, 345, 345, 488,
	489, 490, 503, 284, 189, 189, 494, 22, 491, 492,
	493, 42, 499, 499, 499, 93, 171, 502, 44, 500,
	501, 1, 413, 408, 258, 509, 505, 510, 507, 508,
	511, 228, 237, 238, 231, 232, 233, 234, 235, 236,
	230, 211, 309, 310, 311, 312, 313, 314, 315, 316,
	317, 318, 319, 320, 321, 196, 52, 304, 213, 194,
	294, 186, 186, 195, 88, 495, 470, 432, 162, 441,
	186, 395, 391, 88, 253, 335, 163, 88, 88, 154,
	374, 358, 124, 267, 367, 368, 369, 187, 456, 147,
	400, 398, 173, 165, 164, 166, 167, 168, 169, 343,
	255, 170, 344, 88, 183, 122, 81, 88, 178, 41,
	84, 20, 88, 19, 138, 18, 88, 17, 88, 16,
	15, 145, 21, 14, 13, 12, 11, 10, 157, 158,
	148, 9, 88, 177, 88, 159, 8, 160, 7, 184,
	6, 186, 5, 4, 2, 0, 88, 390, 0, 204,
	87, 0, 174, 0, 0, 0, 0, 0, 88, 100,
	0, 88, 0, 105, 106, 0, 0, 0, 0, 176,
	0, 0, 0, 175, 252, 0, 0, 0, 422, 423,
	424, 0, 0, 0, 0, 0, 0, 0, 0, 130,
	252, 0, 0, 133, 0, 434, 0, 0, 136, 0,
	0, 0, 140, 0, 142, 0, 0, 439, 0, 0,
	0, 88, 0, 0, 0, 0, 0, 0, 149, 0,
	150, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 199, 0, 0, 0, 0, 0, 0, 466,
	0, 0, 0, 0, 222, 0, 0, 223, 0, 0,
	0, 0, 0, 0, 474, 0, 0, 252, 0, 0,
	327, 0, 191, 0, 184, 184, 331, 0, 0, 0,
	0, 0, 0, 184, 88, 0, 0, 0, 0, 485,
	349, 0, 0, 434, 0, 0, 147, 465, 308, 173,
	165, 164, 166, 167, 168, 169, 0, 287, 170, 192,
	193, 0, 0, 188, 0, 178, 229, 228, 237, 238,
	231, 232, 233, 234, 235, 236, 230, 0, 88, 0,
	0, 88, 0, 0, 0, 157, 158, 182, 0, 0,
	177, 0, 159, 0, 160, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 184, 0, 0, 0, 0, 174,
	0, 0, 0, 0, 0, 0, 191, 0, 0, 0,
	0, 411, 0, 88, 0, 0, 176, 0, 0, 0,
	175, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	147, 0, 308, 173, 165, 164, 166, 167, 168, 169,
	0, 0, 170, 192, 193, 0, 0, 188, 0, 178,
	0, 0, 0, 0, 382, 0, 0, 383, 191, 0,
	0, 0, 0, 0, 88, 88, 88, 88, 0, 157,
	158, 182, 0, 0, 177, 0, 159, 411, 160, 0,
	461, 0, 147, 0, 0, 173, 165, 164, 166, 167,
	168, 169, 0, 174, 170, 192, 193, 0, 0, 188,
	0, 178, 0, 22, 0, 0, 0, 0, 0, 0,
	176, 0, 425, 0, 175, 0, 0, 0, 0, 0,
	191, 157, 158, 182, 0, 0, 177, 0, 159, 0,
	160, 229, 228, 237, 238, 231, 232, 233, 234, 235,
	236, 230, 0, 0, 147, 174, 0, 173, 165, 164,
	166, 167, 168, 169, 0, 0, 170, 192, 193, 0,
	0, 188, 176, 178, 0, 0, 175, 191, 0, 0,
	0, 0, 22, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 157, 158, 0, 0, 0, 177, 0,
	159, 147, 160, 0, 173, 165, 164, 166, 167, 168,
	169, 0, 0, 170, 192, 193, 0, 174, 188, 0,
	178, 0, 0, 147, 0, 0, 173, 165, 164, 166,
	167, 168, 169, 0, 176, 170, 0, 0, 175, 0,
	157, 158, 178, 0, 0, 177, 0, 159, 0, 160,
	22, 23, 24, 25, 26, 0, 0, 94, 0, 0,
	0, 0, 157, 158, 174, 0, 0, 177, 0, 159,
	0, 160, 0, 0, 27, 0, 0, 0, 0, 0,
	0, 176, 0, 0, 0, 175, 174, 0, 36, 229,
	228, 237, 238, 231, 232, 233, 234, 235, 236, 230,
	0, 269, 272, 176, 0, 0, 0, 175, 274, 275,
	276, 277, 278, 279, 280, 273, 270, 271, 268, 229,
	228, 237, 238, 231, 232, 233, 234, 235, 236, 230,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 28, 29, 31, 30, 33, 32, 0,
	0, 0, 0, 0, 0, 0, 0, 37, 40, 39,
	34, 0, 0, 0, 0, 0, 35, 0, 0, 0,
	0, 0, 0, 0, 38,
}
var yyPact = [...]int{

	1064, -1000, -1000, 486, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -71, -72,
	-50, -73, -52, -53, -1000, 429, 235, -19, 70, -1000,
	-1000, 481, 418, 360, -1000, -72, 384, 109, 462, 90,
	-68, -68, -58, -1000, -54, -1000, 109, -69, 183, -69,
	109, 109, -1000, -89, -1000, -1000, -1000, -1000, 453, -1000,
	-56, -1000, 302, 409, 409, -1000, -1000, -1000, -1000, -1000,
	-1000, 314, 196, -1000, 64, 382, 109, 371, -12, -1000,
	109, 261, -1000, 35, -1000, 109, 50, 109, 175, 109,
	-64, 109, 392, 301, 109, -1000, -1000, 294, -1000, -1000,
	-1000, -1000, -1000, 109, -1000, 109, -1000, 109, -1000, 522,
	-1000, -1000, 865, -1000, 524, -1000, 368, 366, -1000, 109,
	346, 90, 109, 451, 90, 522, 294, 388, -1000, -76,
	47, 109, -1000, -1000, 109, -1000, 143, -1000, -1000, -1000,
	-1000, -1000, 230, -1000, -40, -1000, -1000, 522, 522, 522,
	522, 294, 294, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -14, 996, -1000, -1000, -1000, -1000, -1000, 522, -1000,
	292, -1000, -1000, 159, -15, 120, 1057, -1000, 974, 927,
	-1000, 294, -1000, -1000, 109, -1000, -1000, -1000, -1000, 345,
	393, 90, 90, 284, -1000, 431, 974, -1000, -17, -1000,
	-1000, 299, 90, -1000, -65, 34, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 131, -1000, -1000, -1000, 522, 522,
	522, 522, 522, 522, 522, 522, 522, 522, 522, 522,
	522, 406, 406, -1000, -1000, -1000, 1027, 719, 813, 135,
	202, 251, -17, 44, -17, 451, 865, 210, -1000, -1000,
	256, -1000, -1000, 45, 974, 974, 167, 65, 129, 108,
	522, 522, 522, 167, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 16, 1057, 232, 1057, -1000, 481, -1000, 393, 90,
	-1000, 294, 486, 261, 281, -1000, 431, 409, 424, 120,
	-1000, 109, -1000, -1000, 109, -1000, 124, -1000, -1000, 130,
	428, -1000, 73, 73, -13, -13, -13, -13, 270, 270,
	-17, -17, -1000, -1000, -1000, -1000, 277, 865, -1000, 277,
	-1000, -33, -1000, 522, -1000, 77, -1000, 974, 436, -1000,
	290, 394, -1000, -1000, 105, 400, 218, -1000, -1000, -37,
	16, 67, -1000, -1000, 140, -1000, -1000, -1000, -17, -1000,
	996, -1000, -1000, 129, 522, 522, 522, -17, -17, 879,
	-1000, -1000, 298, 234, 288, -1000, 522, -1000, 90, 409,
	-1000, 522, 294, -1000, -1000, -1000, -1000, 277, -1000, 90,
	-17, 38, -1000, 522, 75, 433, 411, 210, 210, 210,
	210, -1000, 355, 349, -1000, 341, 313, 348, -6, -1000,
	227, -1000, -1000, 109, -1000, 275, 41, -1000, -1000, -1000,
	251, -1000, -17, -17, 704, 522, -1000, 380, -1000, 294,
	-1000, -1000, 283, -1000, 153, -1000, -1000, -1000, -1000, -17,
	522, 431, 974, 522, 394, 296, 328, -1000, -1000, -1000,
	-1000, 344, -1000, 337, -1000, -1000, -1000, -59, -60, -61,
	-1000, -1000, -1000, -1000, -1000, 522, -17, 374, -1000, 522,
	-1000, -1000, -1000, -1000, -17, 409, 120, 279, 974, 974,
	-1000, -1000, 294, 294, 294, -17, 477, -1000, 378, 120,
	120, 90, 90, 90, 90, -1000, 473, 12, 240, -1000,
	240, 240, 261, -1000, 90, -1000, 90, -1000, -1000, 90,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 624, 22, 623, 622, 620, 618, 616, 611, 607,
	606, 605, 604, 603, 602, 600, 599, 597, 595, 593,
	591, 443, 590, 589, 586, 585, 27, 26, 584, 580,
	16, 21, 13, 579, 571, 10, 570, 582, 568, 6,
	28, 1, 567, 19, 563, 18, 30, 293, 562, 24,
	14, 11, 560, 12, 77, 559, 556, 555, 17, 554,
	552, 551, 549, 548, 8, 547, 4, 546, 3, 545,
	29, 540, 15, 2, 31, 539, 350, 150, 538, 537,
	536, 521, 0, 25, 504, 496, 9, 503, 502, 20,
	154, 501, 5, 7, 498,
}
var yyR1 = [...]int{

	0, 91, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 2, 3, 3, 4, 4, 5, 6,
	7, 8, 8, 8, 9, 9, 9, 10, 11, 11,
	11, 12, 13, 15, 16, 17, 18, 18, 18, 18,
	18, 18, 18, 18, 18, 18, 18, 19, 20, 14,
	94, 21, 22, 22, 23, 23, 23, 24, 24, 25,
	25, 26, 26, 27, 27, 27, 27, 28, 28, 84,
	84, 84, 83, 83, 29, 29, 30, 30, 31, 31,
	32, 32, 32, 33, 33, 33, 33, 88, 88, 87,
	87, 87, 86, 86, 34, 34, 34, 34, 35, 35,
	35, 35, 36, 36, 37, 37, 38, 38, 38, 38,
	39, 39, 40, 40, 41, 41, 41, 41, 41, 41,
	43, 43, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 49, 49, 49, 49, 49,
	49, 44, 44, 44, 44, 44, 44, 44, 50, 50,
	50, 54, 51, 51, 47, 47, 47, 47, 47, 47,
	47, 47, 47, 47, 47, 47, 47, 47, 47, 47,
	47, 47, 47, 47, 47, 47, 47, 47, 47, 47,
	47, 63, 63, 63, 63, 56, 59, 59, 57, 57,
	58, 60, 60, 55, 55, 55, 46, 46, 46, 46,
	46, 46, 46, 48, 48, 48, 61, 61, 62, 62,
	64, 64, 65, 65, 66, 67, 67, 67, 68, 68,
	68, 68, 69, 69, 69, 70, 70, 71, 71, 72,
	72, 45, 45, 52, 52, 53, 73, 73, 74, 75,
	75, 77, 77, 90, 90, 76, 76, 78, 78, 78,
	78, 78, 78, 79, 79, 80, 80, 81, 81, 82,
	85, 92, 93, 89,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 12, 6, 3, 8, 8, 6, 6, 8, 7,
	3, 6, 4, 9, 6, 7, 7, 5, 4, 5,
	4, 3, 3, 2, 7, 3, 3, 3, 3, 3,
	3, 5, 5, 3, 5, 4, 4, 3, 2, 2,
	0, 2, 0, 2, 1, 2, 2, 0, 1, 0,
	1, 1, 3, 1, 2, 3, 5, 1, 1, 0,
	1, 2, 1, 1, 0, 2, 1, 3, 1, 1,
	3, 3, 3, 3, 5, 5, 3, 0, 1, 0,
	1, 2, 1, 1, 1, 2, 2, 1, 2, 3,
	2, 3, 2, 2, 1, 3, 0, 5, 5, 5,
	1, 3, 0, 2, 1, 3, 3, 2, 3, 3,
	1, 1, 1, 3, 3, 3, 4, 3, 4, 3,
	4, 5, 6, 3, 2, 1, 2, 1, 2, 1,
	2, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	1, 3, 1, 3, 1, 1, 1, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 3, 3, 4, 5, 3, 4,
	1, 1, 1, 1, 1, 5, 0, 1, 1, 2,
	4, 0, 2, 1, 3, 5, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 0, 3, 0, 2,
	0, 3, 1, 3, 2, 0, 1, 1, 0, 2,
	4, 4, 0, 2, 4, 0, 3, 1, 3, 0,
	5, 2, 1, 1, 3, 3, 1, 3, 3, 1,
	1, 0, 2, 0, 3, 0, 1, 1, 1, 1,
	1, 1, 1, 0, 1, 0, 1, 0, 2, 1,
	1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -91, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -12, -13, -15, -16, -17, -18, -19,
	-20, -14, 6, 7, 8, 9, 10, 30, 99, 100,
	102, 101, 104, 103, 116, 122, 44, 113, 130, 115,
	114, -23, 5, -21, -94, -21, -21, -21, -21, -21,
	105, 127, -80, 111, -76, 109, 105, 105, 106, 127,
	105, 105, -89, 16, 50, 128, 118, 121, 129, 126,
	99, 117, 123, 119, 120, 53, -89, -89, -2, 20,
	21, -24, 34, 21, -22, -76, 26, -37, -85, 50,
	11, -73, -74, -82, 50, -90, 110, -90, 106, 105,
	-37, -77, 110, 50, -77, -37, -37, 124, -89, -89,
	-89, -89, -89, 11, -89, 105, -89, 46, -68, 17,
	-68, -89, -25, 37, -48, -82, 53, 56, 58, 26,
	-37, 30, 96, -37, 48, 71, -37, 64, -85, 50,
	-37, 108, -37, 23, 46, -85, -92, 47, -85, -37,
	-37, -89, -47, -46, -55, -53, -54, 86, 87, 93,
	95, -82, -63, -56, 52, 51, 53, 54, 55, 56,
	59, -85, -92, 50, 110, 131, 127, 91, 66, -89,
	-26, -27, 88, -28, -85, -41, -47, -42, 64, -92,
	-43, 23, 60, 61, -75, 19, 11, 32, 32, -37,
	-70, 30, -92, -73, -85, -40, 12, -74, -47, -92,
	23, -81, 112, -78, 102, 125, 100, 29, 101, 15,
	132, 50, -37, -37, 50, -89, -89, -89, 83, 82,
	92, 86, 87, 88, 89, 90, 91, 84, 85, 48,
	18, 97, 98, -47, -47, -47, -47, -92, -92, 96,
	-2, -51, -47, -59, -47, -29, 48, 11, -84, -83,
	22, -82, 52, 96, 63, 62, 78, -44, 81, 64,
	79, 80, 65, 78, 71, 72, 73, 74, 75, 76,
	77, -41, -47, -41, -47, -54, -92, -37, -70, 30,
	-45, 32, -2, -73, -71, -82, -40, -64, 15, -41,
	-89, 46, -82, -89, -79, 108, 71, -93, 49, -47,
	-47, -47, -47, -47, -47, -47, -47, -47, -47, -47,
	-47, -47, -46, -46, -82, -93, -26, 21, -93, -26,
	-82, -85, -93, 48, -93, -57, -58, 67, -40, -27,
	-30, -31, -32, -33, -37, -54, -92, -83, 88, -85,
	-41, -41, -49, 59, 64, 60, 61, -43, -47, -50,
	-92, -54, 57, 81, 79, 80, 65, -47, -47, -47,
	-49, -93, -45, -73, -52, -53, -92, -93, 48, -64,
	-68, 16, -37, -37, 50, -89, -93, -26, -93, 96,
	-47, -60, -58, 69, -41, -61, 13, 48, -34, -35,
	-36, 36, 40, 42, 37, 38, 39, 43, -87, -86,
	22, -85, 52, -88, 22, -30, 96, 59, 60, 61,
	-51, -50, -47, -47, -47, 63, -72, 46, -72, 48,
	-82, -68, -65, -66, -47, -92, -93, -82, 70, -47,
	68, -62, 14, 16, -31, -32, -31, -32, 36, 36,
	36, 41, 36, 41, 36, -35, -38, 44, 109, 45,
	-86, -85, -93, 88, -93, 63, -47, 27, -53, 48,
	-67, 24, 25, -89, -47, -64, -41, -51, 46, 46,
	36, 36, 106, 106, 106, -47, 28, -66, -68, -41,
	-41, -92, -92, -92, 9, -69, 19, 31, -39, -82,
	-39, -39, -73, 9, 81, -93, 48, -93, -93, -82,
	-82, -82,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 20, 60, 60, 60, 60, 60, 60, 265, 255,
	0, 0, 0, 0, 273, 0, 0, 0, 0, 273,
	273, 0, 64, 67, 62, 255, 0, 0, 0, 0,
	253, 253, 0, 266, 0, 256, 0, 251, 0, 251,
	0, 0, 43, 0, 273, 273, 273, 273, 273, 273,
	0, 273, 0, 228, 228, 273, 58, 59, 23, 65,
	66, 69, 0, 68, 61, 0, 0, 0, 114, 270,
	0, 30, 246, 0, 269, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 41, 42, 0, 45, 46,
	47, 48, 49, 0, 50, 0, 53, 0, 273, 0,
	273, 57, 0, 70, 0, 213, 0, 0, 63, 0,
	235, 0, 0, 122, 0, 0, 0, 0, 32, 267,
	0, 0, 38, 252, 0, 40, 0, 271, 273, 273,
	273, 55, 229, 164, 165, 166, 167, 0, 0, 0,
	0, 203, 0, 190, 206, 207, 208, 209, 210, 211,
	212, 0, 0, -2, 191, 192, 193, 194, 196, 56,
	84, 71, 73, 79, 0, 77, 78, 124, 0, 0,
	132, 0, 130, 131, 0, 249, 250, 214, 215, 235,
	0, 0, 0, 122, 115, 220, 0, 247, 248, 273,
	254, 0, 0, 273, 263, 0, 257, 258, 259, 260,
	261, 262, 37, 39, 0, 51, 52, 54, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 181, 182, 183, 0, 0, 0, 0,
	0, 0, 162, 0, 197, 122, 0, 0, 74, 80,
	0, 82, 83, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 151, 152, 153, 154, 155, 156,
	157, 127, 0, 0, 162, 144, 0, 22, 0, 0,
	26, 0, 242, 27, 0, 237, 220, 228, 0, 123,
	31, 0, 268, 34, 0, 264, 0, 273, 272, 168,
	169, 170, 171, 172, 173, 174, 175, 176, 177, 178,
	230, 231, 179, 180, 184, 185, 0, 0, 188, 0,
	204, 0, 161, 0, 245, 201, 198, 0, 216, 72,
	85, 86, 88, 89, 99, 97, 0, 81, 75, 0,
	125, 126, 129, 145, 0, 147, 149, 133, 134, 135,
	0, 159, 160, 0, 0, 0, 0, 137, 139, 0,
	143, 128, 239, 239, 241, 243, 0, 236, 0, 228,
	29, 0, 0, 35, 36, 44, 186, 0, 189, 0,
	163, 0, 199, 0, 0, 218, 0, 0, 0, 0,
	0, 104, 0, 0, 107, 0, 0, 0, 116, 100,
	0, 102, 103, 0, 98, 0, 0, 146, 148, 150,
	0, 136, 138, 140, 0, 0, 24, 0, 25, 0,
	238, 28, 221, 222, 225, 273, 187, 205, 195, 202,
	0, 220, 0, 0, 87, 93, 0, 96, 105, 106,
	108, 0, 110, 0, 112, 113, 90, 0, 0, 0,
	101, 91, 92, 76, 158, 0, 141, 0, 244, 0,
	224, 226, 227, 33, 200, 228, 219, 217, 0, 0,
	109, 111, 0, 0, 0, 142, 0, 223, 232, 94,
	95, 0, 0, 0, 0, 21, 0, 0, 0, 120,
	0, 0, 240, 233, 0, 117, 0, 118, 119, 0,
	121, 234,
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
	128, 129, 130, 131, 132,
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
			yyVAL.statement = &ShowVersions{}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:375
		{
			yyVAL.statement = &ShowTables{}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:379
		{
			yyVAL.statement = &ShowEngines{}
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:383
		{
			yyVAL.statement = &ShowTables{Database: yyDollar[4].tableIdent}
		}
	case 52:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:387
		{
			yyVAL.statement = &ShowCreateTable{Table: yyDollar[4].tableName}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:391
		{
			yyVAL.statement = &ShowProcesslist{}
		}
	case 54:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:395
		{
			yyVAL.statement = &ShowPartitions{Table: yyDollar[4].tableName}
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:399
		{
			yyVAL.statement = &ShowQueryz{Limit: yyDollar[3].limit}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:403
		{
			yyVAL.statement = &ShowTxnz{Limit: yyDollar[3].limit}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:409
		{
			yyVAL.statement = &Kill{QueryID: string(yyDollar[2].bytes)}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:415
		{
			yyVAL.statement = &Explain{}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:421
		{
			yyVAL.statement = &Other{}
		}
	case 60:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:426
		{
			setAllowComments(yylex, true)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:430
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:436
		{
			yyVAL.bytes2 = nil
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:440
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:446
		{
			yyVAL.str = UnionStr
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:450
		{
			yyVAL.str = UnionAllStr
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:454
		{
			yyVAL.str = UnionDistinctStr
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:459
		{
			yyVAL.str = ""
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:463
		{
			yyVAL.str = DistinctStr
		}
	case 69:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:468
		{
			yyVAL.str = ""
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:472
		{
			yyVAL.str = StraightJoinHint
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:478
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:482
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:488
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:492
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:496
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Name: yyDollar[1].tableIdent}}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:500
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:506
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:510
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 79:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:515
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:519
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:523
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:530
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 84:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:535
		{
			yyVAL.tableExprs = TableExprs{&AliasedTableExpr{Expr: &TableName{Name: NewTableIdent("dual")}}}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:539
		{
			yyVAL.tableExprs = yyDollar[2].tableExprs
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:545
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:549
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:559
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:563
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:567
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:580
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:584
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 95:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:588
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:592
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 97:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:597
		{
			yyVAL.empty = struct{}{}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:599
		{
			yyVAL.empty = struct{}{}
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:602
		{
			yyVAL.tableIdent = NewTableIdent("")
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:606
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:610
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:617
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:623
		{
			yyVAL.str = JoinStr
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:627
		{
			yyVAL.str = JoinStr
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:631
		{
			yyVAL.str = JoinStr
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:635
		{
			yyVAL.str = StraightJoinStr
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:641
		{
			yyVAL.str = LeftJoinStr
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:645
		{
			yyVAL.str = LeftJoinStr
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:649
		{
			yyVAL.str = RightJoinStr
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:653
		{
			yyVAL.str = RightJoinStr
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:659
		{
			yyVAL.str = NaturalJoinStr
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:663
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:673
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:677
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 116:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:682
		{
			yyVAL.indexHints = nil
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:686
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 118:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:690
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 119:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:694
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:700
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:704
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 122:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:709
		{
			yyVAL.boolExpr = nil
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:713
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:720
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:724
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:728
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:732
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:736
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:742
		{
			yyVAL.boolVal = BoolVal(true)
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:746
		{
			yyVAL.boolVal = BoolVal(false)
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:752
		{
			yyVAL.boolExpr = yyDollar[1].boolVal
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:756
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].boolVal}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:760
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:764
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 136:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:768
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:772
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:776
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:780
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:784
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:788
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:792
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:796
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:800
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:806
		{
			yyVAL.str = IsNullStr
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:810
		{
			yyVAL.str = IsNotNullStr
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:814
		{
			yyVAL.str = IsTrueStr
		}
	case 148:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:818
		{
			yyVAL.str = IsNotTrueStr
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:822
		{
			yyVAL.str = IsFalseStr
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:826
		{
			yyVAL.str = IsNotFalseStr
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:832
		{
			yyVAL.str = EqualStr
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:836
		{
			yyVAL.str = LessThanStr
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:840
		{
			yyVAL.str = GreaterThanStr
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:844
		{
			yyVAL.str = LessEqualStr
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:848
		{
			yyVAL.str = GreaterEqualStr
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:852
		{
			yyVAL.str = NotEqualStr
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:856
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:862
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:866
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:870
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:876
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:882
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:886
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:892
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:896
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:900
		{
			yyVAL.valExpr = yyDollar[1].valTuple
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:904
		{
			yyVAL.valExpr = yyDollar[1].subquery
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:908
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:912
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:916
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:920
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:924
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:928
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:932
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:936
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:940
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:944
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:948
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:952
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONExtractOp, Right: yyDollar[3].valExpr}
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:956
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONUnquoteExtractOp, Right: yyDollar[3].valExpr}
		}
	case 181:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:960
		{
			if num, ok := yyDollar[2].valExpr.(*SQLVal); ok && num.Type == IntVal {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:968
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
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:982
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:986
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.valExpr = &IntervalExpr{Expr: yyDollar[2].valExpr, Unit: yyDollar[3].colIdent}
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
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1002
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1006
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 189:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1010
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1014
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.colIdent = NewColIdent("if")
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.colIdent = NewColIdent("current_timestamp")
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1028
		{
			yyVAL.colIdent = NewColIdent("database")
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1032
		{
			yyVAL.colIdent = NewColIdent("mod")
		}
	case 195:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1038
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1043
		{
			yyVAL.valExpr = nil
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1047
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1053
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1057
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 200:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1063
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1068
		{
			yyVAL.valExpr = nil
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1072
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1078
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1082
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 205:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1086
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.valExpr = NewStrVal(yyDollar[1].bytes)
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1096
		{
			yyVAL.valExpr = NewHexVal(yyDollar[1].bytes)
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1100
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1104
		{
			yyVAL.valExpr = NewFloatVal(yyDollar[1].bytes)
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1108
		{
			yyVAL.valExpr = NewHexNum(yyDollar[1].bytes)
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1112
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1116
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1122
		{
			// TODO(sougou): Deprecate this construct.
			if yyDollar[1].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.valExpr = NewIntVal([]byte("1"))
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1131
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 215:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1135
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1140
		{
			yyVAL.valExprs = nil
		}
	case 217:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1144
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1149
		{
			yyVAL.boolExpr = nil
		}
	case 219:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1153
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1158
		{
			yyVAL.orderBy = nil
		}
	case 221:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1162
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1168
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 223:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1172
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 224:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1178
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 225:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1183
		{
			yyVAL.str = AscScr
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1187
		{
			yyVAL.str = AscScr
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1191
		{
			yyVAL.str = DescScr
		}
	case 228:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1196
		{
			yyVAL.limit = nil
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1200
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 230:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1204
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 231:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1208
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr, Offset: yyDollar[4].valExpr}
		}
	case 232:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1214
		{
			yyVAL.str = ""
		}
	case 233:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1218
		{
			yyVAL.str = ForUpdateStr
		}
	case 234:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1222
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
	case 235:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1235
		{
			yyVAL.columns = nil
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1239
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1245
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 238:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1249
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 239:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1254
		{
			yyVAL.updateExprs = nil
		}
	case 240:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1258
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 241:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1264
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1268
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1274
		{
			yyVAL.values = Values{yyDollar[1].valTuple}
		}
	case 244:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1278
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].valTuple)
		}
	case 245:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1284
		{
			yyVAL.valTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1290
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 247:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1294
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 248:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1300
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 251:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1309
		{
			yyVAL.byt = 0
		}
	case 252:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1311
		{
			yyVAL.byt = 1
		}
	case 253:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1314
		{
			yyVAL.boolVal = false
		}
	case 254:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1316
		{
			yyVAL.boolVal = true
		}
	case 255:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1319
		{
			yyVAL.str = ""
		}
	case 256:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1321
		{
			yyVAL.str = IgnoreStr
		}
	case 257:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1325
		{
			yyVAL.empty = struct{}{}
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1327
		{
			yyVAL.empty = struct{}{}
		}
	case 259:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1329
		{
			yyVAL.empty = struct{}{}
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1331
		{
			yyVAL.empty = struct{}{}
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1333
		{
			yyVAL.empty = struct{}{}
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1335
		{
			yyVAL.empty = struct{}{}
		}
	case 263:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1338
		{
			yyVAL.empty = struct{}{}
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1340
		{
			yyVAL.empty = struct{}{}
		}
	case 265:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1343
		{
			yyVAL.empty = struct{}{}
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1345
		{
			yyVAL.empty = struct{}{}
		}
	case 267:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1348
		{
			yyVAL.empty = struct{}{}
		}
	case 268:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1350
		{
			yyVAL.empty = struct{}{}
		}
	case 269:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1354
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1360
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 271:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1366
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 272:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1375
		{
			decNesting(yylex)
		}
	case 273:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1380
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}

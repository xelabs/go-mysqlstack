// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

import (
	"bytes"
	"fmt"

	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
)

const eofChar = 0x100

// Tokenizer is the struct used to generate SQL
// tokens for the parser.
type Tokenizer struct {
	AllowComments bool
	ForceEOF      bool
	lastToken     []byte
	LastError     error
	posVarIndex   int
	ParseTree     Statement
	nesting       int

	sql         []byte
	pos         int
	tokenStart  int
	tokenEnd    int
	altTokenBuf []byte
	lastChar    uint16
}

// NewStringTokenizer creates a new Tokenizer for the
// sql string.
func NewStringTokenizer(sql string) *Tokenizer {
	return &Tokenizer{sql: []byte(sql)}
}

var keywords = map[string]int{
	"accessible":          UNUSED,
	"add":                 UNUSED,
	"all":                 ALL,
	"alter":               ALTER,
	"analyze":             ANALYZE,
	"and":                 AND,
	"as":                  AS,
	"asc":                 ASC,
	"asensitive":          UNUSED,
	"before":              UNUSED,
	"between":             BETWEEN,
	"bigint":              UNUSED,
	"binary":              UNUSED,
	"blob":                UNUSED,
	"both":                UNUSED,
	"by":                  BY,
	"call":                UNUSED,
	"cascade":             UNUSED,
	"case":                CASE,
	"change":              UNUSED,
	"char":                UNUSED,
	"character":           UNUSED,
	"check":               UNUSED,
	"collate":             UNUSED,
	"column":              UNUSED,
	"condition":           UNUSED,
	"constraint":          UNUSED,
	"continue":            UNUSED,
	"convert":             UNUSED,
	"create":              CREATE,
	"cross":               CROSS,
	"current_date":        UNUSED,
	"current_time":        UNUSED,
	"current_timestamp":   CURRENT_TIMESTAMP,
	"current_user":        UNUSED,
	"cursor":              UNUSED,
	"database":            DATABASE,
	"databases":           DATABASES,
	"day_hour":            UNUSED,
	"day_microsecond":     UNUSED,
	"day_minute":          UNUSED,
	"day_second":          UNUSED,
	"dec":                 UNUSED,
	"decimal":             UNUSED,
	"declare":             UNUSED,
	"default":             DEFAULT,
	"delayed":             UNUSED,
	"delete":              DELETE,
	"desc":                DESC,
	"describe":            DESCRIBE,
	"deterministic":       UNUSED,
	"distinct":            DISTINCT,
	"distinctrow":         UNUSED,
	"div":                 UNUSED,
	"double":              UNUSED,
	"drop":                DROP,
	"duplicate":           DUPLICATE,
	"each":                UNUSED,
	"else":                ELSE,
	"elseif":              UNUSED,
	"enclosed":            UNUSED,
	"end":                 END,
	"engine":              ENGINE,
	"engines":             ENGINES,
	"escaped":             UNUSED,
	"exists":              EXISTS,
	"exit":                UNUSED,
	"explain":             EXPLAIN,
	"false":               FALSE,
	"fetch":               UNUSED,
	"float":               UNUSED,
	"float4":              UNUSED,
	"float8":              UNUSED,
	"for":                 FOR,
	"force":               FORCE,
	"foreign":             UNUSED,
	"from":                FROM,
	"fulltext":            UNUSED,
	"generated":           UNUSED,
	"get":                 UNUSED,
	"grant":               UNUSED,
	"group":               GROUP,
	"hash":                HASH,
	"having":              HAVING,
	"high_priority":       UNUSED,
	"hour_microsecond":    UNUSED,
	"hour_minute":         UNUSED,
	"hour_second":         UNUSED,
	"if":                  IF,
	"ignore":              IGNORE,
	"in":                  IN,
	"index":               INDEX,
	"infile":              UNUSED,
	"inout":               UNUSED,
	"inner":               INNER,
	"insensitive":         UNUSED,
	"insert":              INSERT,
	"int":                 UNUSED,
	"int1":                UNUSED,
	"int2":                UNUSED,
	"int3":                UNUSED,
	"int4":                UNUSED,
	"int8":                UNUSED,
	"integer":             UNUSED,
	"interval":            INTERVAL,
	"into":                INTO,
	"io_after_gtids":      UNUSED,
	"is":                  IS,
	"iterate":             UNUSED,
	"join":                JOIN,
	"key":                 KEY,
	"keys":                UNUSED,
	"kill":                KILL,
	"last_insert_id":      LAST_INSERT_ID,
	"leading":             UNUSED,
	"leave":               UNUSED,
	"left":                LEFT,
	"like":                LIKE,
	"limit":               LIMIT,
	"linear":              UNUSED,
	"lines":               UNUSED,
	"load":                UNUSED,
	"localtime":           UNUSED,
	"localtimestamp":      UNUSED,
	"lock":                LOCK,
	"long":                UNUSED,
	"longblob":            UNUSED,
	"longtext":            UNUSED,
	"loop":                UNUSED,
	"low_priority":        UNUSED,
	"master_bind":         UNUSED,
	"match":               UNUSED,
	"maxvalue":            UNUSED,
	"mediumblob":          UNUSED,
	"mediumint":           UNUSED,
	"mediumtext":          UNUSED,
	"middleint":           UNUSED,
	"minute_microsecond":  UNUSED,
	"minute_second":       UNUSED,
	"mod":                 MOD,
	"modifies":            UNUSED,
	"natural":             NATURAL,
	"next":                NEXT,
	"not":                 NOT,
	"no_write_to_binlog":  UNUSED,
	"null":                NULL,
	"numeric":             UNUSED,
	"offset":              OFFSET,
	"on":                  ON,
	"optimize":            UNUSED,
	"optimizer_costs":     UNUSED,
	"option":              UNUSED,
	"optionally":          UNUSED,
	"or":                  OR,
	"order":               ORDER,
	"out":                 UNUSED,
	"outer":               OUTER,
	"outfile":             UNUSED,
	"partition":           PARTITION,
	"partitions":          PARTITIONS,
	"precision":           UNUSED,
	"primary":             UNUSED,
	"procedure":           UNUSED,
	"processlist":         PROCESSLIST,
	"queryz":              QUERYZ,
	"range":               UNUSED,
	"read":                UNUSED,
	"reads":               UNUSED,
	"read_write":          UNUSED,
	"real":                UNUSED,
	"references":          UNUSED,
	"regexp":              REGEXP,
	"release":             UNUSED,
	"rename":              RENAME,
	"repeat":              UNUSED,
	"replace":             REPLACE,
	"require":             UNUSED,
	"resignal":            UNUSED,
	"restrict":            UNUSED,
	"return":              UNUSED,
	"revoke":              UNUSED,
	"right":               RIGHT,
	"rlike":               REGEXP,
	"schema":              UNUSED,
	"schemas":             UNUSED,
	"second_microsecond":  UNUSED,
	"select":              SELECT,
	"sensitive":           UNUSED,
	"separator":           UNUSED,
	"set":                 SET,
	"show":                SHOW,
	"signal":              UNUSED,
	"smallint":            UNUSED,
	"spatial":             UNUSED,
	"specific":            UNUSED,
	"sql":                 UNUSED,
	"sqlexception":        UNUSED,
	"sqlstate":            UNUSED,
	"sqlwarning":          UNUSED,
	"sql_big_result":      UNUSED,
	"sql_calc_found_rows": UNUSED,
	"sql_small_result":    UNUSED,
	"ssl":                 UNUSED,
	"status":              STATUS,
	"starting":            UNUSED,
	"stored":              UNUSED,
	"straight_join":       STRAIGHT_JOIN,
	"table":               TABLE,
	"tables":              TABLES,
	"terminated":          UNUSED,
	"then":                THEN,
	"tinyblob":            UNUSED,
	"tinyint":             UNUSED,
	"tinytext":            UNUSED,
	"to":                  TO,
	"trailing":            UNUSED,
	"trigger":             UNUSED,
	"true":                TRUE,
	"truncate":            TRUNCATE,
	"txnz":                TXNZ,
	"undo":                UNUSED,
	"union":               UNION,
	"unique":              UNIQUE,
	"unlock":              UNUSED,
	"unsigned":            UNUSED,
	"update":              UPDATE,
	"usage":               UNUSED,
	"use":                 USE,
	"using":               USING,
	"utc_date":            UNUSED,
	"utc_time":            UNUSED,
	"utc_timestamp":       UNUSED,
	"values":              VALUES,
	"varbinary":           UNUSED,
	"varchar":             UNUSED,
	"varcharacter":        UNUSED,
	"varying":             UNUSED,
	"versions":            VERSIONS,
	"virtual":             UNUSED,
	"view":                VIEW,
	"when":                WHEN,
	"where":               WHERE,
	"while":               UNUSED,
	"with":                UNUSED,
	"write":               UNUSED,
	"xa":                  XA,
	"xor":                 UNUSED,
	"year_month":          UNUSED,
	"zerofill":            UNUSED,
}

// Lex returns the next token form the Tokenizer.
// This function is used by go yacc.
func (tkn *Tokenizer) Lex(lval *yySymType) int {
	typ, val := tkn.Scan()
	for typ == COMMENT {
		if tkn.AllowComments {
			break
		}
		typ, val = tkn.Scan()
	}
	lval.bytes = val
	tkn.lastToken = val
	return typ
}

// Position returns the current position.
func (tkn *Tokenizer) Position() int {
	return tkn.pos
}

// Error is called by go yacc if there's a parsing error.
func (tkn *Tokenizer) Error(err string) {
	if tkn.lastToken != nil {
		tkn.LastError = fmt.Errorf("%s at position %v near '%s'", err, tkn.Position(), tkn.lastToken)
	} else {
		tkn.LastError = fmt.Errorf("%s at position %v", err, tkn.Position())
	}
}

// Scan scans the tokenizer for the next token and returns
// the token type and an optional value.
func (tkn *Tokenizer) Scan() (int, []byte) {
	if tkn.ForceEOF {
		return 0, nil
	}

	if tkn.lastChar == 0 {
		tkn.next()
	}
	tkn.skipBlank()
	tkn.startToken()
	switch ch := tkn.lastChar; {
	case isLetter(ch):
		tkn.next()
		if ch == 'X' || ch == 'x' {
			if tkn.lastChar == '\'' {
				tkn.writeByte(byte(ch))
				tkn.writeByte('\'')
				tkn.next()
				return tkn.scanHex()
			}
		}
		return tkn.scanIdentifier(byte(ch))
	case isDigit(ch):
		return tkn.scanNumber(false)
	case ch == ':':
		return tkn.scanBindVar()
	case ch == eofChar:
		return 0, nil
	default:
		tkn.next()
		switch ch {
		case eofChar:
			return 0, nil
		case '=', ',', ';', '(', ')', '+', '*', '%', '^', '~':
			return int(ch), nil
		case '&':
			if tkn.lastChar == '&' {
				tkn.next()
				return AND, nil
			}
			return int(ch), nil
		case '|':
			if tkn.lastChar == '|' {
				tkn.next()
				return OR, nil
			}
			return int(ch), nil
		case '?':
			tkn.posVarIndex++
			buf := new(bytes.Buffer)
			fmt.Fprintf(buf, ":v%d", tkn.posVarIndex)
			return VALUE_ARG, buf.Bytes()
		case '.':
			if isDigit(tkn.lastChar) {
				return tkn.scanNumber(true)
			}
			return int(ch), nil
		case '/':
			switch tkn.lastChar {
			case '/':
				tkn.next()
				return tkn.scanCommentType1("//")
			case '*':
				tkn.next()
				return tkn.scanCommentType2()
			default:
				return int(ch), nil
			}
		case '#':
			tkn.next()
			return tkn.scanCommentType1("#")
		case '-':
			switch tkn.lastChar {
			case '-':
				tkn.next()
				return tkn.scanCommentType1("--")
			case '>':
				tkn.next()
				if tkn.lastChar == '>' {
					tkn.next()
					return JSON_UNQUOTE_EXTRACT_OP, nil
				}
				return JSON_EXTRACT_OP, nil
			}
			return int(ch), nil
		case '<':
			switch tkn.lastChar {
			case '>':
				tkn.next()
				return NE, nil
			case '<':
				tkn.next()
				return SHIFT_LEFT, nil
			case '=':
				tkn.next()
				switch tkn.lastChar {
				case '>':
					tkn.next()
					return NULL_SAFE_EQUAL, nil
				default:
					return LE, nil
				}
			default:
				return int(ch), nil
			}
		case '>':
			switch tkn.lastChar {
			case '=':
				tkn.next()
				return GE, nil
			case '>':
				tkn.next()
				return SHIFT_RIGHT, nil
			default:
				return int(ch), nil
			}
		case '!':
			if tkn.lastChar == '=' {
				tkn.next()
				return NE, nil
			}
			return int(ch), nil
		case '\'', '"':
			return tkn.scanString(ch)
		case '`':
			return tkn.scanLiteralIdentifier()
		default:
			return LEX_ERROR, []byte{byte(ch)}
		}
	}
}

func (tkn *Tokenizer) skipBlank() {
	ch := tkn.lastChar
	for ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t' {
		tkn.next()
		ch = tkn.lastChar
	}
}

func (tkn *Tokenizer) scanIdentifier(firstByte byte) (int, []byte) {
	tkn.writeByte(firstByte)
	for isLetter(tkn.lastChar) || isDigit(tkn.lastChar) {
		tkn.writeByte(byte(tkn.lastChar))
		tkn.next()
	}
	lowered := bytes.ToLower(tkn.token())
	loweredStr := string(lowered)
	if keywordID, found := keywords[loweredStr]; found {
		return keywordID, lowered
	}
	// dual must always be case-insensitive
	if loweredStr == "dual" {
		return ID, lowered
	}
	return ID, tkn.token()
}

func (tkn *Tokenizer) scanHex() (int, []byte) {
	tkn.scanMantissa(16)
	if tkn.lastChar != '\'' {
		return LEX_ERROR, tkn.token()
	}
	tkn.writeByte('\'')
	tkn.next()
	return HEX, tkn.token()
}

func (tkn *Tokenizer) scanLiteralIdentifier() (int, []byte) {
	backTickSeen := false
	for {
		if backTickSeen {
			if tkn.lastChar != '`' {
				break
			}
			backTickSeen = false
			tkn.writeByte('`')
			tkn.next()
			continue
		}
		// The previous char was not a backtick.
		switch tkn.lastChar {
		case '`':
			backTickSeen = true
		case eofChar:
			// Premature EOF.
			return LEX_ERROR, tkn.token()
		default:
			tkn.writeByte(byte(tkn.lastChar))
		}
		tkn.next()
	}
	tok := tkn.token()
	if len(tok) == 0 {
		return LEX_ERROR, tok
	}
	return ID, tok
}

func (tkn *Tokenizer) scanBindVar() (int, []byte) {
	token := VALUE_ARG
	tkn.consumeNext()
	if tkn.lastChar == ':' {
		token = LIST_ARG
		tkn.consumeNext()
	}
	if !isLetter(tkn.lastChar) {
		return LEX_ERROR, tkn.token()
	}
	for isLetter(tkn.lastChar) || isDigit(tkn.lastChar) || tkn.lastChar == '.' {
		tkn.consumeNext()
	}
	return token, tkn.token()
}

func (tkn *Tokenizer) scanMantissa(base int) {
	for digitVal(tkn.lastChar) < base {
		tkn.consumeNext()
	}
}

func (tkn *Tokenizer) scanNumber(seenDecimalPoint bool) (int, []byte) {
	token := INTEGRAL
	if seenDecimalPoint {
		token = FLOAT
		tkn.writeByte('.')
		tkn.scanMantissa(10)
		goto exponent
	}

	// 0x construct.
	if tkn.lastChar == '0' {
		tkn.consumeNext()
		if tkn.lastChar == 'x' || tkn.lastChar == 'X' {
			token = HEXNUM
			tkn.consumeNext()
			tkn.scanMantissa(16)
			goto exit
		}
	}

	tkn.scanMantissa(10)

	if tkn.lastChar == '.' {
		token = FLOAT
		tkn.consumeNext()
		tkn.scanMantissa(10)
	}

exponent:
	if tkn.lastChar == 'e' || tkn.lastChar == 'E' {
		token = FLOAT
		tkn.consumeNext()
		if tkn.lastChar == '+' || tkn.lastChar == '-' {
			tkn.consumeNext()
		}
		tkn.scanMantissa(10)
	}

exit:
	// A letter cannot immediately follow a number.
	if isLetter(tkn.lastChar) {
		return LEX_ERROR, tkn.token()
	}

	return token, tkn.token()
}

func (tkn *Tokenizer) scanString(delim uint16) (int, []byte) {
	for {
		ch := tkn.lastChar
		tkn.next()
		if ch == delim {
			if tkn.lastChar == delim {
				tkn.next()
			} else {
				break
			}
		} else if ch == '\\' {
			if tkn.lastChar == eofChar {
				return LEX_ERROR, tkn.token()
			}
			if decodedChar := sqltypes.SQLDecodeMap[byte(tkn.lastChar)]; decodedChar == sqltypes.DontEscape {
				ch = tkn.lastChar
			} else {
				ch = uint16(decodedChar)
			}
			tkn.next()
		}
		if ch == eofChar || tkn.lastChar == eofChar {
			return LEX_ERROR, tkn.token()
		}
		tkn.writeByte(byte(ch))
	}
	return STRING, tkn.token()
}

func (tkn *Tokenizer) scanCommentType1(prefix string) (int, []byte) {
	tkn.writeString(prefix)
	for tkn.lastChar != eofChar {
		if tkn.lastChar == '\n' {
			tkn.consumeNext()
			break
		}
		tkn.consumeNext()
	}
	return COMMENT, tkn.token()
}

func (tkn *Tokenizer) scanCommentType2() (int, []byte) {
	tkn.writeString("/*")
	for {
		if tkn.lastChar == '*' {
			tkn.consumeNext()
			if tkn.lastChar == '/' {
				tkn.consumeNext()
				break
			}
			continue
		}
		if tkn.lastChar == eofChar {
			return LEX_ERROR, tkn.token()
		}
		tkn.consumeNext()
	}
	return COMMENT, tkn.token()
}

func (tkn *Tokenizer) startToken() {
	pos := tkn.pos - 1
	if pos < 0 {
		pos = 0
	}
	tkn.tokenStart = pos
	tkn.tokenEnd = pos
	tkn.altTokenBuf = nil
}

func (tkn *Tokenizer) next() {
	if tkn.pos >= len(tkn.sql) {
		tkn.lastChar = eofChar
		return
	}
	tkn.pos++
	tkn.lastChar = uint16(tkn.sql[tkn.pos-1])
}

func (tkn *Tokenizer) consumeNext() {
	tkn.writeByte(tkn.sql[tkn.pos-1])
	tkn.next()
}

// writeByte performs a copy on write if the
// byte written does not match the original content.
func (tkn *Tokenizer) writeByte(ch byte) {
	if tkn.altTokenBuf != nil {
		tkn.altTokenBuf = append(tkn.altTokenBuf, ch)
		return
	}
	if tkn.tokenEnd < len(tkn.sql) && ch == tkn.sql[tkn.tokenEnd] {
		tkn.tokenEnd++
		return
	}
	tkn.altTokenBuf = append(tkn.altTokenBuf, tkn.token()...)
	tkn.altTokenBuf = append(tkn.altTokenBuf, ch)
}

func (tkn *Tokenizer) writeString(str string) {
	for i := 0; i < len(str); i++ {
		tkn.writeByte(str[i])
	}
}

func (tkn *Tokenizer) token() []byte {
	if tkn.altTokenBuf != nil {
		return tkn.altTokenBuf
	}
	return tkn.sql[tkn.tokenStart:tkn.tokenEnd]
}

func isLetter(ch uint16) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '@'
}

func digitVal(ch uint16) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch) - '0'
	case 'a' <= ch && ch <= 'f':
		return int(ch) - 'a' + 10
	case 'A' <= ch && ch <= 'F':
		return int(ch) - 'A' + 10
	}
	return 16 // larger than any legal digit val
}

func isDigit(ch uint16) bool {
	return '0' <= ch && ch <= '9'
}

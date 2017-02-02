// Copyright 2015, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package driver

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

// Error codes for server-side errors.
// Originally found in include/mysql/mysqld_error.h
const (
	// ERAccessDeniedError is ER_ACCESS_DENIED_ERROR
	ERAccessDeniedError = 1045

	// ERUnknownComError is ER_UNKNOWN_COM_ERROR
	ERUnknownComError = 1047

	// ERUnknownError is ER_UNKNOWN_ERROR
	ERUnknownError = 1105

	// ERCantDoThisDuringAnTransaction is
	// ER_CANT_DO_THIS_DURING_AN_TRANSACTION
	ERCantDoThisDuringAnTransaction = 1179
)

// Sql states for errors.
// Originally found in include/mysql/sql_state.h
const (
	// SSSignalException is ER_SIGNAL_EXCEPTION
	SSSignalException = "HY000"

	// SSAccessDeniedError is ER_ACCESS_DENIED_ERROR
	SSAccessDeniedError = "28000"

	// SSUnknownComError is ER_UNKNOWN_COM_ERROR
	SSUnknownComError = "08S01"

	// SSHandshakeError is ER_HANDSHAKE_ERROR
	SSHandshakeError = "08S01"

	// SSCantDoThisDuringAnTransaction is
	// ER_CANT_DO_THIS_DURING_AN_TRANSACTION
	SSCantDoThisDuringAnTransaction = "25000"
)

const (
	// SQLStateGeneral is the SQLSTATE value for "general error".
	SQLStateGeneral = "HY000"
)

// SQLError is the error structure returned from calling a db library function
type SQLError struct {
	Num     int
	State   string
	Message string
	Query   string
}

// NewSQLError creates a new SQLError.
// If sqlState is left empty, it will default to "HY000" (general error).
func NewSQLError(number int, sqlState string, format string, args ...interface{}) *SQLError {
	if sqlState == "" {
		sqlState = SQLStateGeneral
	}
	return &SQLError{
		Num:     number,
		State:   sqlState,
		Message: fmt.Sprintf(format, args...),
	}
}

// Error implements the error interface
func (se *SQLError) Error() string {
	buf := &bytes.Buffer{}
	buf.WriteString(se.Message)

	// Add MySQL errno and SQLSTATE in a format that we can later parse.
	// There's no avoiding string parsing because all errors
	// are converted to strings anyway at RPC boundaries.
	// See NewSQLErrorFromError.
	fmt.Fprintf(buf, " (errno %v) (sqlstate %v)", se.Num, se.State)

	if se.Query != "" {
		fmt.Fprintf(buf, " during query: %s", se.Query)
	}
	return buf.String()
}

// Number returns the internal MySQL error code.
func (se *SQLError) Number() int {
	return se.Num
}

// SQLState returns the SQLSTATE value.
func (se *SQLError) SQLState() string {
	return se.State
}

var errExtract = regexp.MustCompile(`.*\(errno ([0-9]*)\) \(sqlstate ([0-9a-zA-Z]{5})\).*`)

// NewSQLErrorFromError returns a *SQLError from the provided error.
// If it's not the right type, it still tries to get it from a regexp.
func NewSQLErrorFromError(err error) error {
	if err == nil {
		return nil
	}

	if serr, ok := err.(*SQLError); ok {
		return serr
	}

	msg := err.Error()
	match := errExtract.FindStringSubmatch(msg)
	if len(match) < 2 {
		// Not found, build a generic SQLError.
		// TODO(alainjobart) maybe we can also check the canonical
		// error code, and translate that into the right error.

		// FIXME(alainjobart): 1105 is unknown error. Will
		// merge with sqlconn later.
		return &SQLError{
			Num:     1105,
			State:   SQLStateGeneral,
			Message: msg,
		}
	}

	num, err := strconv.Atoi(match[1])
	if err != nil {
		return &SQLError{
			Num:     1105,
			State:   SQLStateGeneral,
			Message: msg,
		}
	}

	serr := &SQLError{
		Num:     num,
		State:   match[2],
		Message: msg,
	}
	return serr
}

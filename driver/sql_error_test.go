/*
 * go-mysqlstack
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package driver

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestSqlError(t *testing.T) {
	sqlerr := NewSQLError(1, "HY000", "i.am.error.man")
	assert.Equal(t, 1, sqlerr.Number())
	assert.Equal(t, "i.am.error.man (errno 1) (sqlstate HY000)", sqlerr.Error())
	assert.Equal(t, "HY000", sqlerr.SQLState())
}

func TestSqlErrorFromErr(t *testing.T) {
	{
		err := errors.New("errorman")
		sqlerr := NewSQLErrorFromError(err)
		assert.NotNil(t, sqlerr)
	}

	{
		err := errors.New("i.am.error.man (errno 1) (sqlstate HY000)")
		sqlerr := NewSQLErrorFromError(err)
		assert.NotNil(t, sqlerr)
	}
}

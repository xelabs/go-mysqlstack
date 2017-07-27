// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

import ()

func (*ShowDatabases) iStatement() {}

// ShowDatabases represents a SHOW DATABASES statement.
type ShowDatabases struct {
}

// Format formats the node.
func (node *ShowDatabases) Format(buf *TrackedBuffer) {
	buf.WriteString("SHOW DATABASES")
}

// WalkSubtree walks the nodes of the subtree.
func (node *ShowDatabases) WalkSubtree(visit Visit) error {
	return nil
}

func (*ShowTables) iStatement() {}

// ShowTables represents a SHOW TABLES statement.
type ShowTables struct {
	Database TableIdent
}

// Format formats the node.
func (node *ShowTables) Format(buf *TrackedBuffer) {
	buf.WriteString("SHOW TABLES")
	if !node.Database.IsEmpty() {
		buf.Myprintf(" FROM %s", node.Database.String())
	}
}

// WalkSubtree walks the nodes of the subtree.
func (node *ShowTables) WalkSubtree(visit Visit) error {
	return nil
}

func (*ShowCreateTable) iStatement() {}

// ShowCreateTable represents a SHOW CREATE TABLE statement.
type ShowCreateTable struct {
	Table *TableName
}

// Format formats the node.
func (node *ShowCreateTable) Format(buf *TrackedBuffer) {
	buf.WriteString("SHOW CREATE TABLE ")
	if !node.Table.Qualifier.IsEmpty() {
		buf.Myprintf("%s.", node.Table.Qualifier.String())
	}
	buf.Myprintf("%s", node.Table.Name.String())
}

// WalkSubtree walks the nodes of the subtree.
func (node *ShowCreateTable) WalkSubtree(visit Visit) error {
	return nil
}

func (*ShowProcesslist) iStatement() {}

// ShowProcesslist represents a SHOW PROCESSLIST statement.
type ShowProcesslist struct {
}

// Format formats the node.
func (node *ShowProcesslist) Format(buf *TrackedBuffer) {
	buf.WriteString("SHOW PROCESSLIST")
}

// WalkSubtree walks the nodes of the subtree.
func (node *ShowProcesslist) WalkSubtree(visit Visit) error {
	return nil
}

func (*ShowEngines) iStatement() {}

// ShowEngines represents a SHOW ENGINES statement.
type ShowEngines struct {
}

// Format formats the node.
func (node *ShowEngines) Format(buf *TrackedBuffer) {
	buf.WriteString("SHOW ENGINES")
}

// WalkSubtree walks the nodes of the subtree.
func (node *ShowEngines) WalkSubtree(visit Visit) error {
	return nil
}

func (*ShowStatus) iStatement() {}

// ShowStatus represents a SHOW DATABASES statement.
type ShowStatus struct {
}

// Format formats the node.
func (node *ShowStatus) Format(buf *TrackedBuffer) {
	buf.WriteString("SHOW STATUS")
}

// WalkSubtree walks the nodes of the subtree.
func (node *ShowStatus) WalkSubtree(visit Visit) error {
	return nil
}

func (*ShowVersions) iStatement() {}

// ShowVersions represents a SHOW VERSIONS statement.
type ShowVersions struct {
}

// Format formats the node.
func (node *ShowVersions) Format(buf *TrackedBuffer) {
	buf.WriteString("SHOW VERSIONS")
}

// WalkSubtree walks the nodes of the subtree.
func (node *ShowVersions) WalkSubtree(visit Visit) error {
	return nil
}

func (*ShowPartitions) iStatement() {}

// ShowPartitions represents a SHOW PARTITIONS ON TABLE statement.
type ShowPartitions struct {
	Table *TableName
}

// Format formats the node.
func (node *ShowPartitions) Format(buf *TrackedBuffer) {
	buf.WriteString("SHOW PARTITIONS ON ")
	if !node.Table.Qualifier.IsEmpty() {
		buf.Myprintf("%s.", node.Table.Qualifier.String())
	}
	buf.Myprintf("%s", node.Table.Name.String())
}

// WalkSubtree walks the nodes of the subtree.
func (node *ShowPartitions) WalkSubtree(visit Visit) error {
	return nil
}

func (*ShowQueryz) iStatement() {}

// ShowQueryz represents a SHOW QUERYZ statement.
type ShowQueryz struct {
	Limit *Limit
}

// Format formats the node.
func (node *ShowQueryz) Format(buf *TrackedBuffer) {
	buf.Myprintf("SHOW QUERYZ%v", node.Limit)
}

// WalkSubtree walks the nodes of the subtree.
func (node *ShowQueryz) WalkSubtree(visit Visit) error {
	return nil
}

func (*ShowTxnz) iStatement() {}

// ShowTxnz represents a SHOW TXNZ statement.
type ShowTxnz struct {
	Limit *Limit
}

// Format formats the node.
func (node *ShowTxnz) Format(buf *TrackedBuffer) {
	buf.Myprintf("SHOW TXNZ%v", node.Limit)
}

// WalkSubtree walks the nodes of the subtree.
func (node *ShowTxnz) WalkSubtree(visit Visit) error {
	return nil
}

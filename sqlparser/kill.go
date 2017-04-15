// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

import ()

func (*Kill) iStatement() {}

// Kill represents a KILL statement.
type Kill struct {
	QueryID string
}

// Format formats the node.
func (node *Kill) Format(buf *TrackedBuffer) {
	buf.Myprintf("kill %s", node.QueryID)
}

// WalkSubtree walks the nodes of the subtree.
func (node *Kill) WalkSubtree(visit Visit) error {
	return nil
}

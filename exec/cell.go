// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exec

import (
	"context"

	cells "github.com/ipfn/go-ipfn-cells"
	"github.com/ipfn/go-ipfn-cells/chainops"
)

// Cell - Cell during execution.
type Cell interface {
	cells.Cell

	// Context - Execution context.
	Context() context.Context

	// WithContext - Cell with context.
	WithContext(context.Context) Cell

	// Parent - Parent cell.
	Parent() cells.Cell

	// ExecChild - Child cell by index.
	ExecChild(int) Cell
}

// NewRoot - Creates new root exec cell.
func NewRoot(ctx context.Context, cell cells.Cell) Cell {
	return &execCell{
		Cell: cell,
		ctx:  ctx,
	}
}

// NewCell - Creates new exec cell.
func NewCell(ctx context.Context, parent, cell cells.Cell) Cell {
	return &execCell{
		Cell:   cell,
		ctx:    ctx,
		parent: parent,
	}
}

type execCell struct {
	cells.Cell
	ctx    context.Context
	parent cells.Cell
}

func (c *execCell) Context() context.Context {
	return c.ctx
}

func (c *execCell) WithContext(ctx context.Context) Cell {
	return &execCell{
		Cell:   c.Cell,
		parent: c.parent,
		ctx:    ctx,
	}
}

func (c *execCell) ExecChild(n int) Cell {
	child := c.Cell.Child(n)
	switch v := child.(type) {
	case *execCell:
		return v
	default:
		return NewCell(c.ctx, c.parent, child)
	}
}

func (c *execCell) Parent() cells.Cell {
	if c.parent == nil {
		return chainops.Root(c.Cell)
	}
	return c.parent
}

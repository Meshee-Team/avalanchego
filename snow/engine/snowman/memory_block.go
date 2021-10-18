// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
)

// memoryBlock wraps a snowman Block to manage non-verified blocks
type memoryBlock struct {
	snowman.Block

	tree AncestorTree
}

// Accept accepts the underlying block & removes sibling subtrees
func (mb *memoryBlock) Accept() error {
	mb.tree.RemoveSiblings(mb.ID())
	return mb.Block.Accept()
}

// Reject rejects the underlying block & removes child subtrees
func (mb *memoryBlock) Reject() error {
	mb.tree.RemoveSubtree(mb.ID())
	return mb.Block.Reject()
}
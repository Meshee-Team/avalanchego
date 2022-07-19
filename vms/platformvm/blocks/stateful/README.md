This package deals with stateful blocks.

The *Block type implements the snowman.Block interface.
This is the type that the platformvm deals with when it uses snowman.Blocks.
*Block wraps a stateless.Block and a manager.
The *Block itself doesn't have any state.
The state is all held by the manager, and the *Block acts upon the manager to get/set the state.
Therefore, we don't need to worry about deduplicating *Block instances.

This package contains three implementations of stateless.Visitor: verifier, acceptor and rejector.
These implement the logic for verifying, accepting and rejecting blocks.
Each implementation has a reference to a shared *backend, which maintains state, etc.
(The manager has a reference to the shared *backend as well.)

The platformvm uses the manager to create blocks and query block state. 
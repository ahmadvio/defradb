package clock

// CRDTs are composed of two structures, the payload and a clock. The
// payload is the actual CRDT data which abides by the merge semantics.
// The clock is a mechanism to provide a casual ordering of events, so
// we can determine which event proceeded eachother and apply the
// various merge strategies.
//
// MerkleCRDTs are similar, they contain a CRDT payload, but instead
// of a logical or vector clock, it uses a MerkleClock.
//
// MerkleClock is a Merkle DAG, which provides casual ordering of nodes
// which link to previous nodes, with content addressable IDs creating a
// fully linked graph of content. The linked graph of nodes creates a
// natural history of events because a parent node contains a CID of a
// child node, which ensures parents occured AFTER a child.

//	  A			  	   B			  C
//	//////   link	//////	 link	//////
//	//--//--------->//--//--------->//	//
//	//////			//////			//////
//   head							 tail

// The above diagram shows the ordering of events A, B, and C.

// API
// mc = NewMerkleClock(blockstore)
// event = mc.NewEvent(delta)
// mc.AddEvent(event) cid
// mc.HasEvent(cid)
// mc.
// extractDelta(node) delta

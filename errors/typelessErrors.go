package errors

import "fmt"

var NoHead = fmt.Errorf("noHead: block header has no head")
var NoConsensusStateMatch = fmt.Errorf("noConsensusStateMatch: block dose'nt have relavent ConsensusState")
var InvalidSignatureHashPair = fmt.Errorf("invalidSignatureHashPair: can't verify signature for the hash")

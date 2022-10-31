// Copyright (C) 2019-2022, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package chains

import (
	"github.com/luxdefi/luxd/snow/engine/common"
)

// Registrant can register the existence of a chain
type Registrant interface {
	// Called when the chain described by [engine] is created
	// This function is called before the chain starts processing messages
	// [engine] should be an lux.Engine or snowman.Engine
	RegisterChain(name string, engine common.Engine)
}

// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/luxdefi/node/snow"
	"github.com/luxdefi/node/snow/consensus/snowball"
	"github.com/luxdefi/node/snow/consensus/snowman"
	"github.com/luxdefi/node/snow/engine/common"
	"github.com/luxdefi/node/snow/engine/common/tracker"
	"github.com/luxdefi/node/snow/engine/snowman/block"
	"github.com/luxdefi/node/snow/validators"
)

func DefaultConfig() Config {
	return Config{
		Ctx:                 snow.DefaultConsensusContextTest(),
		VM:                  &block.TestVM{},
		Sender:              &common.SenderTest{},
		Validators:          validators.NewManager(),
		ConnectedValidators: tracker.NewPeers(),
		Params: snowball.Parameters{
			K:                     1,
			AlphaPreference:       1,
			AlphaConfidence:       1,
			BetaVirtuous:          1,
			BetaRogue:             2,
			ConcurrentRepolls:     1,
			OptimalProcessing:     100,
			MaxOutstandingItems:   1,
			MaxItemProcessingTime: 1,
		},
		Consensus: &snowman.Topological{},
	}
}

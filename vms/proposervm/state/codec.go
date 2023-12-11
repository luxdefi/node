// Copyright (C) 2019-2023, Lux Partners Limited All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"math"

	"github.com/luxdefi/node/codec"
	"github.com/luxdefi/node/codec/linearcodec"
)

const version = 0

var c codec.Manager

func init() {
	lc := linearcodec.NewCustomMaxLength(math.MaxUint32)
	c = codec.NewManager(math.MaxInt32)

	err := c.RegisterCodec(version, lc)
	if err != nil {
		panic(err)
	}
}

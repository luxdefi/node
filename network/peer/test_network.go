// Copyright (C) 2022, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"crypto"
	"time"

	"github.com/luxdefi/luxd/ids"
	"github.com/luxdefi/luxd/message"
	"github.com/luxdefi/luxd/utils/ips"
	"github.com/luxdefi/luxd/version"
)

var _ Network = (*testNetwork)(nil)

// testNetwork is a network definition for a TestPeer
type testNetwork struct {
	mc message.Creator

	networkID uint32
	ip        ips.IPPort
	version   *version.Application
	signer    crypto.Signer
	subnets   ids.Set

	uptime uint8
}

// NewTestNetwork creates and returns a new TestNetwork
func NewTestNetwork(
	mc message.Creator,
	networkID uint32,
	ipPort ips.IPPort,
	version *version.Application,
	signer crypto.Signer,
	subnets ids.Set,
	uptime uint8,
) Network {
	return &testNetwork{
		mc:        mc,
		networkID: networkID,
		ip:        ipPort,
		version:   version,
		signer:    signer,
		subnets:   subnets,
		uptime:    uptime,
	}
}

func (*testNetwork) Connected(ids.NodeID) {}

func (*testNetwork) AllowConnection(ids.NodeID) bool {
	return true
}

func (*testNetwork) Track(ips.ClaimedIPPort) bool {
	return true
}

func (*testNetwork) Disconnected(ids.NodeID) {}

func (n *testNetwork) Version() (message.OutboundMessage, error) {
	now := uint64(time.Now().Unix())
	unsignedIP := UnsignedIP{
		IP:        n.ip,
		Timestamp: now,
	}
	signedIP, err := unsignedIP.Sign(n.signer)
	if err != nil {
		return nil, err
	}
	return n.mc.Version(
		n.networkID,
		now,
		n.ip,
		n.version.String(),
		now,
		signedIP.Signature,
		n.subnets.List(),
	)
}

func (*testNetwork) Peers(ids.NodeID) ([]ids.NodeID, []ips.ClaimedIPPort, error) {
	return nil, nil, nil
}

func (n *testNetwork) Pong(ids.NodeID) (message.OutboundMessage, error) {
	return n.mc.Pong(n.uptime)
}

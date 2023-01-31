// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"golang.org/x/exp/maps"

	"github.com/ava-labs/avalanchego/ids"
<<<<<<< HEAD
	"github.com/ava-labs/avalanchego/utils"
=======
>>>>>>> 4d169e12a (Add BLS keys to validator set (#2073))
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
)

var (
	_ Manager = (*manager)(nil)

	errMissingValidators = errors.New("missing validators")
)

// Manager holds the validator set of each subnet
type Manager interface {
	fmt.Stringer

	// Add a subnet's validator set to the manager.
	//
	// If the subnet had previously registered a validator set, false will be
	// returned and the manager will not be modified.
	Add(subnetID ids.ID, set Set) bool

<<<<<<< HEAD
<<<<<<< HEAD
	// Get returns the validator set for the given subnet
	// Returns false if the subnet doesn't exist
	Get(ids.ID) (Set, bool)
=======
	// AddWeight adds weight to a given validator on the given subnet
	AddWeight(ids.ID, ids.NodeID, uint64) error

	// RemoveWeight removes weight from a given validator on a given subnet
	RemoveWeight(ids.ID, ids.NodeID, uint64) error

	// Get returns the validator set for the given subnet
	// Returns false if the subnet doesn't exist
	Get(ids.ID) (Set, bool)

	// Contains returns true if there is a validator with the specified ID
	// currently in the set.
	Contains(ids.ID, ids.NodeID) bool
>>>>>>> f6ea8e56f (Rename validators.Manager#GetValidators to Get (#2279))
=======
	// Get returns the validator set for the given subnet
	// Returns false if the subnet doesn't exist
	Get(ids.ID) (Set, bool)
>>>>>>> f171d317d (Remove unnecessary functions from validators.Manager interface (#2277))
}

// NewManager returns a new, empty manager
func NewManager() Manager {
	return &manager{
		subnetToVdrs: make(map[ids.ID]Set),
	}
}

type manager struct {
	lock sync.RWMutex

	// Key: Subnet ID
	// Value: The validators that validate the subnet
	subnetToVdrs map[ids.ID]Set
}

func (m *manager) Add(subnetID ids.ID, set Set) bool {
	m.lock.Lock()
	defer m.lock.Unlock()

	if _, exists := m.subnetToVdrs[subnetID]; exists {
		return false
	}

	m.subnetToVdrs[subnetID] = set
	return true
}

<<<<<<< HEAD
<<<<<<< HEAD
=======
func (m *manager) AddWeight(subnetID ids.ID, vdrID ids.NodeID, weight uint64) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	vdrs, ok := m.subnetToVdrs[subnetID]
	if !ok {
		vdrs = NewSet()
		m.subnetToVdrs[subnetID] = vdrs
	}
	return vdrs.AddWeight(vdrID, weight)
}

func (m *manager) RemoveWeight(subnetID ids.ID, vdrID ids.NodeID, weight uint64) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if vdrs, ok := m.subnetToVdrs[subnetID]; ok {
		return vdrs.RemoveWeight(vdrID, weight)
	}
	return nil
}

>>>>>>> f6ea8e56f (Rename validators.Manager#GetValidators to Get (#2279))
=======
>>>>>>> f171d317d (Remove unnecessary functions from validators.Manager interface (#2277))
func (m *manager) Get(subnetID ids.ID) (Set, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	vdrs, ok := m.subnetToVdrs[subnetID]
	return vdrs, ok
}

func (m *manager) String() string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	subnets := maps.Keys(m.subnetToVdrs)
<<<<<<< HEAD
	utils.Sort(subnets)
=======
	ids.SortIDs(subnets)
>>>>>>> 78e44f3a8 (Use maps library where possible (#2280))

	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Validator Manager: (Size = %d)",
		len(subnets),
	))
	for _, subnetID := range subnets {
		vdrs := m.subnetToVdrs[subnetID]
		sb.WriteString(fmt.Sprintf(
			"\n    Subnet[%s]: %s",
			subnetID,
			vdrs.PrefixedString("    "),
		))
	}

	return sb.String()
}

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 749a0d8e9 (Add validators.Set#Add function and report errors (#2276))
// Add is a helper that fetches the validator set of [subnetID] from [m] and
// adds [nodeID] to the validator set.
// Returns an error if:
// - [subnetID] does not have a registered validator set in [m]
// - adding [nodeID] to the validator set returns an error
<<<<<<< HEAD
<<<<<<< HEAD
func Add(m Manager, subnetID ids.ID, nodeID ids.NodeID, pk *bls.PublicKey, txID ids.ID, weight uint64) error {
=======
func Add(m Manager, subnetID ids.ID, nodeID ids.NodeID, weight uint64) error {
>>>>>>> 749a0d8e9 (Add validators.Set#Add function and report errors (#2276))
=======
func Add(m Manager, subnetID ids.ID, nodeID ids.NodeID, pk *bls.PublicKey, weight uint64) error {
>>>>>>> 4d169e12a (Add BLS keys to validator set (#2073))
	vdrs, ok := m.Get(subnetID)
	if !ok {
		return fmt.Errorf("%w: %s", errMissingValidators, subnetID)
	}
<<<<<<< HEAD
<<<<<<< HEAD
	return vdrs.Add(nodeID, pk, txID, weight)
}

=======
>>>>>>> f171d317d (Remove unnecessary functions from validators.Manager interface (#2277))
=======
	return vdrs.Add(nodeID, weight)
=======
	return vdrs.Add(nodeID, pk, weight)
>>>>>>> 4d169e12a (Add BLS keys to validator set (#2073))
}

>>>>>>> 749a0d8e9 (Add validators.Set#Add function and report errors (#2276))
// AddWeight is a helper that fetches the validator set of [subnetID] from [m]
// and adds [weight] to [nodeID] in the validator set.
// Returns an error if:
// - [subnetID] does not have a registered validator set in [m]
// - adding [weight] to [nodeID] in the validator set returns an error
func AddWeight(m Manager, subnetID ids.ID, nodeID ids.NodeID, weight uint64) error {
	vdrs, ok := m.Get(subnetID)
	if !ok {
		return fmt.Errorf("%w: %s", errMissingValidators, subnetID)
	}
	return vdrs.AddWeight(nodeID, weight)
}

// RemoveWeight is a helper that fetches the validator set of [subnetID] from
// [m] and removes [weight] from [nodeID] in the validator set.
// Returns an error if:
// - [subnetID] does not have a registered validator set in [m]
// - removing [weight] from [nodeID] in the validator set returns an error
func RemoveWeight(m Manager, subnetID ids.ID, nodeID ids.NodeID, weight uint64) error {
	vdrs, ok := m.Get(subnetID)
	if !ok {
		return fmt.Errorf("%w: %s", errMissingValidators, subnetID)
	}
	return vdrs.RemoveWeight(nodeID, weight)
}

// AddWeight is a helper that fetches the validator set of [subnetID] from [m]
// and returns if the validator set contains [nodeID]. If [m] does not contain a
// validator set for [subnetID], false is returned.
func Contains(m Manager, subnetID ids.ID, nodeID ids.NodeID) bool {
	vdrs, ok := m.Get(subnetID)
	if !ok {
		return false
	}
	return vdrs.Contains(nodeID)
}

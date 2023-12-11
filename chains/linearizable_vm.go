// Copyright (C) 2019-2023, Lux Partners Limited All rights reserved.
// See the file LICENSE for licensing terms.

package chains

import (
	"context"

	"github.com/luxdefi/node/api/metrics"
	"github.com/luxdefi/node/database"
	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/node/snow"
	"github.com/luxdefi/node/snow/engine/lux/vertex"
	"github.com/luxdefi/node/snow/engine/common"
	"github.com/luxdefi/node/snow/engine/snowman/block"
)

var (
	_ vertex.LinearizableVM = (*initializeOnLinearizeVM)(nil)
	_ block.ChainVM         = (*linearizeOnInitializeVM)(nil)
)

// initializeOnLinearizeVM transforms the consensus engine's call to Linearize
// into a call to Initialize. This enables the proposervm to be initialized by
// the call to Linearize. This also provides the stopVertexID to the
// linearizeOnInitializeVM.
type initializeOnLinearizeVM struct {
	vertex.DAGVM
	vmToInitialize common.VM
	vmToLinearize  *linearizeOnInitializeVM

	registerer   metrics.OptionalGatherer
	ctx          *snow.Context
	db           database.Database
	genesisBytes []byte
	upgradeBytes []byte
	configBytes  []byte
	toEngine     chan<- common.Message
	fxs          []*common.Fx
	appSender    common.AppSender
}

func (vm *initializeOnLinearizeVM) Linearize(ctx context.Context, stopVertexID ids.ID) error {
	vm.vmToLinearize.stopVertexID = stopVertexID
	vm.ctx.Metrics = vm.registerer
	return vm.vmToInitialize.Initialize(
		ctx,
		vm.ctx,
		vm.db,
		vm.genesisBytes,
		vm.upgradeBytes,
		vm.configBytes,
		vm.toEngine,
		vm.fxs,
		vm.appSender,
	)
}

// linearizeOnInitializeVM transforms the proposervm's call to Initialize into a
// call to Linearize. This enables the proposervm to provide its toEngine
// channel to the VM that is being linearized.
type linearizeOnInitializeVM struct {
	vertex.LinearizableVMWithEngine
	stopVertexID ids.ID
}

func NewLinearizeOnInitializeVM(vm vertex.LinearizableVMWithEngine) *linearizeOnInitializeVM {
	return &linearizeOnInitializeVM{
		LinearizableVMWithEngine: vm,
	}
}

func (vm *linearizeOnInitializeVM) Initialize(
	ctx context.Context,
	_ *snow.Context,
	_ database.Database,
	_ []byte,
	_ []byte,
	_ []byte,
	toEngine chan<- common.Message,
	_ []*common.Fx,
	_ common.AppSender,
) error {
	return vm.Linearize(ctx, vm.stopVertexID, toEngine)
}

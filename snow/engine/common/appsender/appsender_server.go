// Copyright (C) 2019-2023, Lux Partners Limited All rights reserved.
// See the file LICENSE for licensing terms.

package appsender

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/node/snow/engine/common"
	"github.com/luxdefi/node/utils/set"

	appsenderpb "github.com/luxdefi/node/proto/pb/appsender"
)

var _ appsenderpb.AppSenderServer = (*Server)(nil)

type Server struct {
	appsenderpb.UnsafeAppSenderServer
	appSender common.AppSender
}

// NewServer returns a messenger connected to a remote channel
func NewServer(appSender common.AppSender) *Server {
	return &Server{appSender: appSender}
}

func (s *Server) SendCrossChainAppRequest(ctx context.Context, msg *appsenderpb.SendCrossChainAppRequestMsg) (*emptypb.Empty, error) {
	chainID, err := ids.ToID(msg.ChainId)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, s.appSender.SendCrossChainAppRequest(ctx, chainID, msg.RequestId, msg.Request)
}

func (s *Server) SendCrossChainAppResponse(ctx context.Context, msg *appsenderpb.SendCrossChainAppResponseMsg) (*emptypb.Empty, error) {
	chainID, err := ids.ToID(msg.ChainId)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, s.appSender.SendCrossChainAppResponse(ctx, chainID, msg.RequestId, msg.Response)
}

func (s *Server) SendAppRequest(ctx context.Context, req *appsenderpb.SendAppRequestMsg) (*emptypb.Empty, error) {
	nodeIDs := set.NewSet[ids.NodeID](len(req.NodeIds))
	for _, nodeIDBytes := range req.NodeIds {
		nodeID, err := ids.ToNodeID(nodeIDBytes)
		if err != nil {
			return nil, err
		}
		nodeIDs.Add(nodeID)
	}
	err := s.appSender.SendAppRequest(ctx, nodeIDs, req.RequestId, req.Request)
	return &emptypb.Empty{}, err
}

func (s *Server) SendAppResponse(ctx context.Context, req *appsenderpb.SendAppResponseMsg) (*emptypb.Empty, error) {
	nodeID, err := ids.ToNodeID(req.NodeId)
	if err != nil {
		return nil, err
	}
	err = s.appSender.SendAppResponse(ctx, nodeID, req.RequestId, req.Response)
	return &emptypb.Empty{}, err
}

func (s *Server) SendAppGossip(ctx context.Context, req *appsenderpb.SendAppGossipMsg) (*emptypb.Empty, error) {
	err := s.appSender.SendAppGossip(ctx, req.Msg)
	return &emptypb.Empty{}, err
}

func (s *Server) SendAppGossipSpecific(ctx context.Context, req *appsenderpb.SendAppGossipSpecificMsg) (*emptypb.Empty, error) {
	nodeIDs := set.NewSet[ids.NodeID](len(req.NodeIds))
	for _, nodeIDBytes := range req.NodeIds {
		nodeID, err := ids.ToNodeID(nodeIDBytes)
		if err != nil {
			return nil, err
		}
		nodeIDs.Add(nodeID)
	}
	err := s.appSender.SendAppGossipSpecific(ctx, nodeIDs, req.Msg)
	return &emptypb.Empty{}, err
}

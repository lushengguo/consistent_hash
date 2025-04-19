// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2
// Source: consistent_hash.proto

package server

import (
	"context"

	"consistent_hash/consistent_hash/proto"
	"consistent_hash/internal/logic/consistenthashserver"
	"consistent_hash/internal/svc"
)

type ConsistentHashServerServer struct {
	svcCtx *svc.ServiceContext
	proto_consistent_hash.UnimplementedConsistentHashServerServer
}

func NewConsistentHashServerServer(svcCtx *svc.ServiceContext) *ConsistentHashServerServer {
	return &ConsistentHashServerServer{
		svcCtx: svcCtx,
	}
}

func (s *ConsistentHashServerServer) GetValue(ctx context.Context, in *proto_consistent_hash.GetValueRequest) (*proto_consistent_hash.GetValueResponse, error) {
	l := consistenthashserverlogic.NewGetValueLogic(ctx, s.svcCtx)
	return l.GetValue(in)
}

func (s *ConsistentHashServerServer) SetKV(ctx context.Context, in *proto_consistent_hash.SetKVRequest) (*proto_consistent_hash.SetKVResponse, error) {
	l := consistenthashserverlogic.NewSetKVLogic(ctx, s.svcCtx)
	return l.SetKV(in)
}

func (s *ConsistentHashServerServer) DeleteKey(ctx context.Context, in *proto_consistent_hash.DeleteKeyRequest) (*proto_consistent_hash.DeleteKeyResponse, error) {
	l := consistenthashserverlogic.NewDeleteKeyLogic(ctx, s.svcCtx)
	return l.DeleteKey(in)
}

func (s *ConsistentHashServerServer) ListenGossip(ctx context.Context, in *proto_consistent_hash.Gossip) (*proto_consistent_hash.Gossip, error) {
	l := consistenthashserverlogic.NewListenGossipLogic(ctx, s.svcCtx)
	return l.ListenGossip(in)
}

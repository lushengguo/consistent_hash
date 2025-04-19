package consistenthashserverlogic

import (
	"context"

	"consistent_hash/consistent_hash/proto"
	"consistent_hash/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListenGossipLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListenGossipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListenGossipLogic {
	return &ListenGossipLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListenGossipLogic) ListenGossip(in *proto_consistent_hash.Gossip) (*proto_consistent_hash.Gossip, error) {
	// todo: add your logic here and delete this line

	return &proto_consistent_hash.Gossip{}, nil
}

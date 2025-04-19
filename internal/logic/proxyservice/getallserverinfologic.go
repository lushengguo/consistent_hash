package proxyservicelogic

import (
	"context"

	"consistent_hash/consistent_hash/proto"
	"consistent_hash/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllServerInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllServerInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllServerInfoLogic {
	return &GetAllServerInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllServerInfoLogic) GetAllServerInfo(in *proto_consistent_hash.Empty) (*proto_consistent_hash.Gossip, error) {
	// todo: add your logic here and delete this line

	return &proto_consistent_hash.Gossip{}, nil
}

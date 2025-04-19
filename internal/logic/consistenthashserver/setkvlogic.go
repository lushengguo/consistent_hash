package consistenthashserverlogic

import (
	"context"

	"consistent_hash/consistent_hash/proto"
	"consistent_hash/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetKVLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetKVLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetKVLogic {
	return &SetKVLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetKVLogic) SetKV(in *proto_consistent_hash.SetKVRequest) (*proto_consistent_hash.SetKVResponse, error) {
	// todo: add your logic here and delete this line

	return &proto_consistent_hash.SetKVResponse{}, nil
}

package consistenthashserverlogic

import (
	"context"

	"consistent_hash/consistent_hash/proto"
	"consistent_hash/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetValueLogic {
	return &GetValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetValueLogic) GetValue(in *proto_consistent_hash.GetValueRequest) (*proto_consistent_hash.GetValueResponse, error) {
	// todo: add your logic here and delete this line

	return &proto_consistent_hash.GetValueResponse{}, nil
}

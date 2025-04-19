package proxyservicelogic

import (
	"context"

	"consistent_hash/consistent_hash/proto"
	"consistent_hash/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteKeyLogic {
	return &DeleteKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteKeyLogic) DeleteKey(in *proto_consistent_hash.DeleteKeyRequest) (*proto_consistent_hash.DeleteKeyResponse, error) {
	// todo: add your logic here and delete this line

	return &proto_consistent_hash.DeleteKeyResponse{}, nil
}

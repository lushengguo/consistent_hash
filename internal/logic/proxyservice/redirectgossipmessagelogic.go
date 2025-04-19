package proxyservicelogic

import (
	"context"

	"consistent_hash/consistent_hash/proto"
	"consistent_hash/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedirectGossipMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRedirectGossipMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedirectGossipMessageLogic {
	return &RedirectGossipMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// redirect all kinds of message
func (l *RedirectGossipMessageLogic) RedirectGossipMessage(in *proto_consistent_hash.RedirectGossipMessageRequest) (*proto_consistent_hash.RedirectGossipMessageResponse, error) {
	// todo: add your logic here and delete this line

	return &proto_consistent_hash.RedirectGossipMessageResponse{}, nil
}

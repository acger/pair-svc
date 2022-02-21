package logic

import (
	"context"
	"github.com/acger/pair-svc/internal/svc"
	"github.com/acger/pair-svc/pair"

	"github.com/zeromicro/go-zero/core/logx"
)

type ElementPairLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewElementPairLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ElementPairLogic {
	return &ElementPairLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ElementPairLogic) ElementPair(in *pair.ElePairReq) (*pair.ELeListRsp, error) {
	//todo add pagination & mode pair

	return &pair.ELeListRsp{}, nil
}

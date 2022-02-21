package logic

import (
	"context"

	"github.com/acger/pair-svc/internal/svc"
	"github.com/acger/pair-svc/pair"

	"github.com/zeromicro/go-zero/core/logx"
)

type ElementListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewElementListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ElementListLogic {
	return &ElementListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ElementListLogic) ElementList(in *pair.EleListReq) (*pair.ELeListRsp, error) {
	// todo: add your logic here and delete this line

	return &pair.ELeListRsp{}, nil
}

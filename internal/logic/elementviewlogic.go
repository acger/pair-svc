package logic

import (
	"context"
	"errors"
	"github.com/acger/pair-svc/internal/svc"
	"github.com/acger/pair-svc/model"
	"github.com/acger/pair-svc/pair"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ElementViewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewElementViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ElementViewLogic {
	return &ElementViewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ElementViewLogic) ElementView(in *pair.EleViewReq) (*pair.EleViewRsp, error) {
	var ele model.Element
	var eleRsp pair.Element

	r := l.svcCtx.DB.Find(&ele, "uid = ?", in.Uid)

	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return &pair.EleViewRsp{Code: 0}, nil
	}

	if r.Error != nil {
		return &pair.EleViewRsp{Code: 20001, Message: r.Error.Error()}, nil
	}

	copier.Copy(&eleRsp, &ele)

	return &pair.EleViewRsp{
		Code:    0,
		Element: &eleRsp,
	}, nil
}

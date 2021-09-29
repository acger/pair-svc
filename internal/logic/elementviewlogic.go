package logic

import (
	"context"
	"errors"
	"github.com/acger/pair-svc/internal/svc"
	"github.com/acger/pair-svc/model"
	"github.com/acger/pair-svc/pair"
	"github.com/acger/pair-svc/pairclient"
	"gorm.io/gorm"

	"github.com/tal-tech/go-zero/core/logx"
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
	var ele []model.Element
	var eleRsp []*pairclient.Element
	r := l.svcCtx.DB.Find(&ele, "uid = ?", in.Uid)

	if errors.Is(r.Error, gorm.ErrRecordNotFound){
		return &pair.EleViewRsp{Code: 0}, nil
	}

	if r.Error != nil {
		return &pair.EleViewRsp{Code: 20001}, nil
	}

	for _, e := range ele {
		eleRsp = append(eleRsp, &pairclient.Element{
			Uid:  e.Uid,
			Name: e.Name,
			Mode: e.Mode,
			Star: e.Star,
			Sort: e.Sort,
		})
	}

	return &pair.EleViewRsp{
		Code:    0,
		Element: eleRsp,
	}, nil

}

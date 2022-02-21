package logic

import (
	"context"
	"encoding/json"
	"github.com/acger/pair-svc/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm/clause"
	"strconv"

	"github.com/acger/pair-svc/internal/svc"
	"github.com/acger/pair-svc/pair"

	"github.com/zeromicro/go-zero/core/logx"
)

type ElementSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewElementSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ElementSaveLogic {
	return &ElementSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ElementSaveLogic) ElementSave(in *pair.EleSaveReq) (*pair.Response, error) {
	cache := l.svcCtx.Cache
	cacheKey := "element"
	uidStr := strconv.FormatInt(int64(in.Uid), 10)
	eleCache, _ := cache.Hget(cacheKey, uidStr)
	inString, _ := json.Marshal(in)

	if eleCache != "" && string(inString) == eleCache {
		return &pair.Response{Code: 0}, nil
	}

	var ele *model.Element
	copier.Copy(ele, in)

	err := l.svcCtx.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "uid"}},
		UpdateAll: true,
	}).Create(ele)

	if err != nil {
		return &pair.Response{Code: 20001}, nil
	}

	cache.Hset(cacheKey, uidStr, string(inString))

	return &pair.Response{Code: 0}, nil
}

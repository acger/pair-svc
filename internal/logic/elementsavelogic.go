package logic

import (
	"context"
	"encoding/json"
	"github.com/acger/pair-svc/model"
	"gorm.io/gorm"
	"strconv"

	"github.com/acger/pair-svc/internal/svc"
	"github.com/acger/pair-svc/pair"

	"github.com/tal-tech/go-zero/core/logx"
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

	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		delResult := tx.Where("uid = ?", in.Uid).Delete(&model.Element{})

		if delResult.Error != nil {
			return delResult.Error
		}

		var ele []model.Element

		for _, e := range in.Element {
			ele = append(ele, model.Element{
				Uid:  in.Uid,
				Name: e.Name,
				Mode: e.Mode,
				Star: e.Star,
				Sort: e.Sort,
			})
		}

		tx.Create(ele)

		return nil
	})

	if err != nil {
		return &pair.Response{Code: 20001}, nil
	}

	cache.Hset(cacheKey, uidStr, string(inString))

	return &pair.Response{Code: 0}, nil
}

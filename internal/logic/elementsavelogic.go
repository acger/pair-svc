package logic

import (
	"context"
	"encoding/json"
	"github.com/acger/pair-svc/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
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

	var ele = model.Element{}
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		tx.Where("uid = ?", in.Uid).First(&ele)
		copier.Copy(&ele, in.Element)
		ele.Uid = in.Uid

		if ele.ID == 0 {
			if err := tx.Create(&ele).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Save(&ele).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return &pair.Response{Code: 20001, Message: err.Error()}, nil
	}

	cache.Hset(cacheKey, uidStr, string(inString))
	return &pair.Response{Code: 0}, nil
}

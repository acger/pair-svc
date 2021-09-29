package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/acger/pair-svc/common"
	"github.com/acger/pair-svc/model"
	"github.com/acger/pair-svc/pairclient"
	"github.com/acger/user-svc/userclient"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"strconv"

	"github.com/acger/pair-svc/internal/svc"
	"github.com/acger/pair-svc/pair"

	"github.com/tal-tech/go-zero/core/logx"
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

	cache := l.svcCtx.Cache
	uidStr := strconv.FormatInt(int64(in.Uid), 10)
	cacheKey := "pair:result:" + uidStr
	cacheResultString, err := cache.Get(cacheKey)
	if err == nil && cacheResultString != "" {
		var rsp pair.ELeListRsp
		json.Unmarshal([]byte(cacheResultString), &rsp)
		return &rsp, nil
	}

	//匹配出用户id
	db := l.svcCtx.DB
	type pairResult struct {
		Uid uint64
		Num int64
	}

	var pairList []pairResult
	subQuery := db.Table("elements").Select("name").Where("uid = ?", in.Uid).Where("deleted_at IS NULL")

	tx := db.Model(&model.Element{})
	tx = tx.Select("uid, count(uid) as num")
	tx = tx.Where("name in (?)", subQuery)
	tx = tx.Where("uid <> ? ", in.Uid)
	tx = tx.Group("uid")
	tx = tx.Order("num desc")
	tx = tx.Limit(30)
	pr := tx.Scan(&pairList)

	if errors.Is(pr.Error, gorm.ErrRecordNotFound) {
		return &pair.ELeListRsp{Code: 0}, nil
	}

	if pr.Error != nil {
		return &pair.ELeListRsp{Code: 20001, Message: common.ErrorCode[20001]}, nil
	}

	var uidList []uint64
	for _, e := range pairList {
		uidList = append(uidList, e.Uid)
	}

	//获取该批用户的所有元素
	var element []*model.Element
	er := db.Find(&element, "uid in ?", uidList)

	if er.Error != nil {
		return &pair.ELeListRsp{Code: 20001, Message: common.ErrorCode[20001]}, nil
	}

	//以用户id为键值，对元素进行分组
	userElementMap := make(map[uint64]*pairclient.UserElement)

	for _, e := range element {
		if _, ok := userElementMap[e.Uid]; !ok {
			userElementMap[e.Uid] = &pairclient.UserElement{Id: e.Uid}
		}

		userElementMap[e.Uid].Element = append(userElementMap[e.Uid].Element, &pairclient.Element{
			Uid:  e.Uid,
			Name: e.Name,
			Mode: e.Mode,
			Star: e.Star,
			Sort: e.Sort,
		})
	}

	//获取用户信息
	userListRsp, _ := l.svcCtx.UserSvc.UserList(l.ctx, &userclient.UserListReq{Id: uidList})
	userMap := make(map[uint64]*userclient.UserInfo)
	userElementList := make([]*pair.UserElement, len(uidList), len(uidList))

	for _, u := range userListRsp.User {
		userMap[u.Id] = u
	}

	//组装返回信息
	for i, item := range pairList {
		userId := item.Uid
		ue := userElementMap[userId]
		copier.Copy(ue, userMap[userId])
		userElementList[i] = ue
	}

	rsp := pair.ELeListRsp{
		Code:        0,
		UserElement: userElementList,
	}

	rspString, err := json.Marshal(rsp)
	cache.Set(cacheKey, string(rspString))
	cache.Expire(cacheKey, 180)

	return &rsp, nil
}

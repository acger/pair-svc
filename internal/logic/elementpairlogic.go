package logic

import (
	"bytes"
	"context"
	"fmt"
	"github.com/acger/pair-svc/database"
	"github.com/acger/pair-svc/internal/svc"
	"github.com/acger/pair-svc/pair"
	"github.com/acger/user-svc/user"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"time"

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

	//ES搜索
	es := l.svcCtx.ES
	ele := database.GetCacheElement(in.Uid, l.svcCtx.Cache)

	var boost string
	boostInt := int(ele.Element.Boost)
	boost = strconv.Itoa(boostInt)

	from := (in.Page - 1) * in.PageSize
	size := in.PageSize
	fromStr := strconv.FormatInt(int64(from), 10)
	sizeStr := strconv.FormatInt(int64(size), 10)
	uidStr := strconv.FormatInt(int64(in.Uid), 10)
	timeStr := fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"))

	body := &bytes.Buffer{}
	body.WriteString(`
		{
		  "from": ` + fromStr + `,
		  "size": ` + sizeStr + `,
		  "_source": ["uid"], 
		  "query": {
			"function_score": {
			  "query": {
				"bool": {
				  "must": [
					{
					  "multi_match": {
						"fields": [
						  "skill",
						  "skill_need"
						],
						"query": "` + ele.Element.SkillNeed + `"
					  }
					}
				  ],
				  "must_not": [
					{
					  "exists": {
						"field": "deleted_at"
					  }
					},
					{
					  "term": {
						"uid": {
						  "value": ` + uidStr + `
						}
					  }
					}
				  ]
				}
			  },
			  "functions": [
				{
				  "gauss": {
					"updated_at": {
					  "origin": "` + timeStr + `",
					  "scale": "10d",
					  "offset": "90d",
					  "decay": 0.5
					}
				  }
				},
				{
				  "filter": {
					"range": {
					  "boost": {
						"gt": 0
					  }
					}
				  },
				  "weight": ` + boost + `
				},
				{
				  "filter": {
					"term": {
					  "star": 1
					}
				  },
				  "weight": 5
				}
			  ],
			  "score_mode": "sum"
			}
		  },
		  "highlight": {
			"pre_tags": "<b>",
			"post_tags": "</b>", 
			"fields": {
			  "skill": {},
			  "skill_need": {}
			}
		  }
		}
    `)

	rsp, err := es.Search(es.Search.WithIndex(database.ES_ACGER_PAIR), es.Search.WithBody(body))
	if err != nil {
		return &pair.ELeListRsp{Code: 1}, nil
	}

	//读取搜索结果
	buff := bytes.Buffer{}
	buff.ReadFrom(rsp.Body)
	result := database.EsSearchPairResult{}
	jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(buff.Bytes(), &result)
	len := result.Hits.Total.Value

	if len == 0 {
		return &pair.ELeListRsp{}, nil
	}

	//RPC获取用户信息
	var uids []uint64
	uidHighlight := make(map[int]database.Highlight)

	for _, h := range result.Hits.Hits {
		uid := h.Source.UID
		uids = append(uids, uint64(uid))
		uidHighlight[uid] = h.Highlight
	}

	usersReq := user.UserListReq{Id: uids}
	usersRsp, userListErr := l.svcCtx.UserSvc.UserList(l.ctx, &usersReq)

	if userListErr != nil {
		return &pair.ELeListRsp{Code: usersRsp.Code, Message: userListErr.Error()}, nil
	}

	users := make(map[uint64]pair.UserElement)

	for _, u := range usersRsp.User {
		users[u.Id] = pair.UserElement{
			Id:      u.Id,
			Name:    u.Name,
			Account: u.Account,
			Avatar:  u.Avatar,
			Element: new(pair.Element),
		}
	}

	//给搜索结果填补上用户信息
	var ue []*pair.UserElement
	for uid, hl := range uidHighlight {
		id := uint64(uid)
		e := users[id]
		e.Element.Uid = id

		if hl.Skill != nil {
			e.Element.Skill = hl.Skill[0]
		}

		if hl.SkillNeed != nil {
			e.Element.SkillNeed = hl.SkillNeed[0]
		}

		ue = append(ue, &e)
	}

	return &pair.ELeListRsp{Code: 0, UserElement: ue}, nil
}

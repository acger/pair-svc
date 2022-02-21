// Code generated by goctl. DO NOT EDIT!
// Source: pair.proto

package server

import (
	"context"

	"github.com/acger/pair-svc/internal/logic"
	"github.com/acger/pair-svc/internal/svc"
	"github.com/acger/pair-svc/pb/pair"
)

type PairServer struct {
	svcCtx *svc.ServiceContext
	pair.UnimplementedPairServer
}

func NewPairServer(svcCtx *svc.ServiceContext) *PairServer {
	return &PairServer{
		svcCtx: svcCtx,
	}
}

func (s *PairServer) ElementSave(ctx context.Context, in *pair.EleSaveReq) (*pair.Response, error) {
	l := logic.NewElementSaveLogic(ctx, s.svcCtx)
	return l.ElementSave(in)
}

func (s *PairServer) ElementView(ctx context.Context, in *pair.EleViewReq) (*pair.EleViewRsp, error) {
	l := logic.NewElementViewLogic(ctx, s.svcCtx)
	return l.ElementView(in)
}

func (s *PairServer) ElementList(ctx context.Context, in *pair.EleListReq) (*pair.ELeListRsp, error) {
	l := logic.NewElementListLogic(ctx, s.svcCtx)
	return l.ElementList(in)
}

func (s *PairServer) ElementPair(ctx context.Context, in *pair.ElePairReq) (*pair.ELeListRsp, error) {
	l := logic.NewElementPairLogic(ctx, s.svcCtx)
	return l.ElementPair(in)
}

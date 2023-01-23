package logic

import (
	"context"

	"github.com/YOSVU/musical/test/internal/svc"
	"github.com/YOSVU/musical/test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FromLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFromLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FromLogic {
	return &FromLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FromLogic) From(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

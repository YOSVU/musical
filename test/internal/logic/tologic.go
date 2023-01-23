package logic

import (
	"context"

	"github.com/YOSVU/musical/test/internal/svc"
	"github.com/YOSVU/musical/test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ToLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewToLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ToLogic {
	return &ToLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ToLogic) To(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

package logic

import (
	"context"
	"strings"

	"github.com/YOSVU/musical/test/internal/svc"
	"github.com/YOSVU/musical/test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test(req *types.Request) (resp *types.Response, err error) {

	return &types.Response{
		Message: strings.Join([]string{"hello ", req.Name}, ""),
	}, nil
}

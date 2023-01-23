package handler

import (
	"net/http"

	"github.com/YOSVU/musical/test/internal/logic"
	"github.com/YOSVU/musical/test/internal/svc"
	"github.com/YOSVU/musical/test/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ToHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewToLogic(r.Context(), svcCtx)
		resp, err := l.To(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

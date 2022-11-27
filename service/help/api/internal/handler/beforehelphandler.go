package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"help/api/internal/logic"
	"help/api/internal/svc"
	"help/api/internal/types"
)

func beforeHelpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BeforeHelpReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewBeforeHelpLogic(r.Context(), svcCtx)
		resp, err := l.BeforeHelp(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

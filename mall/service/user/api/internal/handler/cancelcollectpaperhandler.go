package handler

import (
	"net/http"

	"Ingoland/backend/mall/service/user/api/internal/logic"
	"Ingoland/backend/mall/service/user/api/internal/svc"
	"Ingoland/backend/mall/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CancelCollectPaperHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CancelCollectPaperRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCancelCollectPaperLogic(r.Context(), svcCtx)
		resp, err := l.CancelCollectPaper(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"soft2_backend/service/user/api/internal/logic"
	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"
)

func GetSubscribeScholarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSubscribeScholarRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetSubscribeScholarLogic(r.Context(), svcCtx)
		resp, err := l.GetSubscribeScholar(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

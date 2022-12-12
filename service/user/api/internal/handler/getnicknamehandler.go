package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"soft2_backend/service/user/api/internal/logic"
	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"
)

func GetNickNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetNickNameRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetNickNameLogic(r.Context(), svcCtx)
		resp, err := l.GetNickName(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

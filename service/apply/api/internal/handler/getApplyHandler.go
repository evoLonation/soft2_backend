package handler

import (
	"net/http"
	"soft2_backend/service/apply/api/internal/logic"
	"soft2_backend/service/apply/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetApplyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetApplyLogic(r.Context(), svcCtx)
		resp, err := l.GetApply()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

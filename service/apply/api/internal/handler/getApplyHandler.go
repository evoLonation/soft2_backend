package handler

import (
	"go-zero-share/apply/api/internal/logic"
	"go-zero-share/apply/api/internal/svc"
	"net/http"

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

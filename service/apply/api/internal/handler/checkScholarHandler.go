package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-share/apply/api/internal/logic"
	"go-zero-share/apply/api/internal/svc"
)

func CheckScholarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCheckScholarLogic(r.Context(), svcCtx)
		resp, err := l.CheckScholar()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"soft2_backend/service/paper/api/internal/logic"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"
)

func FieldPaperHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FieldPaperRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFieldPaperLogic(r.Context(), svcCtx)
		resp, err := l.FieldPaper(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

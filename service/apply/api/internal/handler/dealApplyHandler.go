package handler

import (
	"net/http"
	"soft2_backend/service/apply/api/internal/logic"
	"soft2_backend/service/apply/api/internal/svc"
	"soft2_backend/service/apply/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DealApplyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DealApplyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDealApplyLogic(r.Context(), svcCtx)
		err := l.DealApply(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

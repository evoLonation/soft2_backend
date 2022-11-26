package handler

import (
	"net/http"
	"soft2_backend/service/file/api/internal/logic"

	"github.com/zeromicro/go-zero/rest/httpx"
	"soft2_backend/service/file/api/internal/svc"
	"soft2_backend/service/file/api/internal/types"
)

func UploadApplyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadApplyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUploadApplyLogic(r.Context(), svcCtx)
		err := l.UploadApply(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

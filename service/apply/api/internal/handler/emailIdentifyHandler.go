package handler

import (
	"go-zero-share/apply/api/internal/logic"
	"go-zero-share/apply/api/internal/svc"
	"go-zero-share/apply/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func EmailIdentifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailIdentifyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEmailIdentifyLogic(r.Context(), svcCtx)
		err := l.EmailIdentify(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

package handler

import (
	"net/http"
	"soft2_backend/service/message/api/internal/logic"
	"soft2_backend/service/message/api/internal/svc"
	"soft2_backend/service/message/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReadMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReadMessageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewReadMessageLogic(r.Context(), svcCtx)
		err := l.ReadMessage(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

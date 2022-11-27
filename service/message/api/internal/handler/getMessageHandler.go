package handler

import (
	"net/http"
	"soft2_backend/service/message/api/internal/logic"
	"soft2_backend/service/message/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetMessageLogic(r.Context(), svcCtx)
		resp, err := l.GetMessage()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

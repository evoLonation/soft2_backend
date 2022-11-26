package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"soft2_backend/service/jwttest/api/internal/logic"
	"soft2_backend/service/jwttest/api/internal/svc"
)

func checkJwtHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCheckJwtLogic(r.Context(), svcCtx)
		resp, err := l.CheckJwt()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

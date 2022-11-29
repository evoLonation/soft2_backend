package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"soft2_backend/service/file/rpctest/internal/logic"
	"soft2_backend/service/file/rpctest/internal/svc"
)

func rpcTestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewRpcTestLogic(r.Context(), svcCtx)
		err := l.RpcTest()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

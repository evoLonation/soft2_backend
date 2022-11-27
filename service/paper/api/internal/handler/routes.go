// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"soft2_backend/service/paper/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/search/paper",
				Handler: PaperHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/search/scholar",
				Handler: ScholarHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/search/auto-complete",
				Handler: AutoCompleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/paper",
				Handler: PaperDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/paper/cite",
				Handler: PaperCiteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/scholar/basic",
				Handler: ScholarBasicHandler(serverCtx),
			},
		},
	)
}
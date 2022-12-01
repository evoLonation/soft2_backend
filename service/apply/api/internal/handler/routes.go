// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"soft2_backend/service/apply/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/get-scholar-apply",
				Handler: GetApplyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/deal-scholar-apply",
				Handler: DealApplyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/scholar/check-user",
				Handler: CheckUserHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/scholar/email-verify-code",
				Handler: EmailVerifyCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/scholar/email-identify",
				Handler: EmailIdentifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/scholar/check-scholar",
				Handler: CheckScholarHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}

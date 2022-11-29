// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"soft2_backend/service/file/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user/upload-avatar",
				Handler: UploadAvatarHandler(serverCtx),

			},
			{
				Method:  http.MethodPost,
				Path:    "/api/help/upload",
				Handler: UploadHelpHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/scholar/file-identify",
				Handler: UploadApplyHandler(serverCtx),
			},
			

		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/get-file/:name/:suffix",
				Handler: GetFileHandler(serverCtx),
			},
		},
	)
}

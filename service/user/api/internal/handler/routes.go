// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"soft2_backend/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/paper/get-comment",
				Handler: GetCommentHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user/user-info",
				Handler: UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/paper/star",
				Handler: CollectPaperHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/paper/star/cancel",
				Handler: CancelCollectPaperHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/paper/comment",
				Handler: CommentPaperHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/paper/comment/delete",
				Handler: DeleteCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/paper/comment/like",
				Handler: LikeCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/paper/comment/cancel",
				Handler: CancelLikeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/scholar/subscribe",
				Handler: SubscribeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/scholar/delete-subscribe",
				Handler: DeleteSubscribeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/paper/grievance",
				Handler: LaunchGrievanceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/grievance/accept",
				Handler: GrievanceAcceptHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/grievance/refuse",
				Handler: GrievanceRefuseHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/paper/comment-liked",
				Handler: IfLikedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/paper/is-star",
				Handler: IfCollectHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}

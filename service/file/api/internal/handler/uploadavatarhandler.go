package handler

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"soft2_backend/service/file/api/internal/common"
	"soft2_backend/service/file/api/internal/logic"
	"soft2_backend/service/file/api/internal/svc"
)

func UploadAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(common.DefaultMultipartMemory); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUploadAvatarLogic(r.Context(), svcCtx)
		if r.MultipartForm.File["file"] == nil {
			httpx.Error(w, errors.New("请求的form-data请包含file字段"))
			return
		}
		l.File = r.MultipartForm.File["file"][0]
		if l.File == nil {
			httpx.Error(w, errors.New("请求的form-data请包含file字段"))
			return
		}

		err := l.UploadAvatar()

		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

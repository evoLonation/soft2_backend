package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"soft2_backend/service/file/api/internal/logic"
	"soft2_backend/service/file/api/internal/svc"
	"soft2_backend/service/file/filecommon"
)

func UploadAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(filecommon.DefaultMultipartMemory); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUploadAvatarLogic(r.Context(), svcCtx)

		var err error
		l.File.File, l.File.FileHeader, err = r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
		}

		err = l.UploadAvatar()

		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

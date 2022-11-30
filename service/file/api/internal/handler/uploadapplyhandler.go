package handler

import (
	"net/http"
	"soft2_backend/service/file/api/internal/logic"
	"soft2_backend/service/file/filecommon"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"soft2_backend/service/file/api/internal/svc"
)

func UploadApplyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(filecommon.DefaultMultipartMemory); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUploadApplyLogic(r.Context(), svcCtx)
		var err error
		l.File.File, l.File.FileHeader, err = r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
		}

		id := r.FormValue("scholar_id")
		l.ScholarId, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			httpx.Error(w, err)
		}

		err = l.UploadApply()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

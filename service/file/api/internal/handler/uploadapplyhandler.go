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

		l.File = filecommon.GetFormFile(w, r.MultipartForm)
		if l.File == nil {
			return
		}

		id, success := filecommon.GetFormValue(w, r.MultipartForm, "scholar_id")
		if !success {
			return
		}
		var err error
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

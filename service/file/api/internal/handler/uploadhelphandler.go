package handler

import (
	"mime/multipart"
	"net/http"
	"soft2_backend/service/file/api/internal/logic"
	"soft2_backend/service/file/filecommon"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"soft2_backend/service/file/api/internal/svc"
)

func UploadHelpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(filecommon.DefaultMultipartMemory); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUploadHelpLogic(r.Context(), svcCtx)
		fileNum, err := strconv.ParseInt(r.FormValue("file_num"), 10, 64)
		if err != nil {
			httpx.Error(w, err)
		}
		var index = 1
		for index <= int(fileNum) {
			file, header, err := r.FormFile("file" + strconv.Itoa(index))
			if err != nil {
				httpx.Error(w, err)
			}
			l.Files = append(l.Files, struct {
				*multipart.FileHeader
				multipart.File
			}{FileHeader: header, File: file})
		}

		id := r.FormValue("request_id")
		l.RequestId, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			httpx.Error(w, err)
		}

		err = l.UploadHelp()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

package handler

import (
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"mime/multipart"
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
		fileNum, err := strconv.ParseInt(r.FormValue("file_num"), 10, 64)
		if err != nil {
			httpx.Error(w, errors.New("无法得到file_num"))
			logx.Alert(err.Error())
			return
		}
		l.ScholarId = r.FormValue("scholar_id")
		if err != nil {
			httpx.Error(w, errors.New("无法得到request_id"))
			logx.Alert(err.Error())
			return
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
			index++
		}
		err = l.UploadApply()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

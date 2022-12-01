package handler

import (
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
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
		var err error
		l.File.File, l.File.FileHeader, err = r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		l.RequestId, err = strconv.ParseInt(r.FormValue("request_id"), 10, 64)
		if err != nil {
			httpx.Error(w, errors.New("无法得到request_id"))
			logx.Alert(err.Error())
			return
		}

		err = l.UploadHelp()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

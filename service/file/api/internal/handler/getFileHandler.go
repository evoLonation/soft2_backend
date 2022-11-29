package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"soft2_backend/service/file/api/internal/svc"
	"soft2_backend/service/file/api/internal/types"
	"soft2_backend/service/file/filecommon"
)

func GetFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		http.ServeFile(w, r, filecommon.FilePath+req.FileName+"."+req.FileSuffix)
	}
}

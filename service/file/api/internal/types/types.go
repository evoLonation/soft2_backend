// Code generated by goctl. DO NOT EDIT.
package types

type GetFileReq struct {
	FileName   string `path:"name"`
	FileSuffix string `path:"suffix"`
}

type UploadAvatarReq struct {
}

type UploadAvatarRes struct {
}

type UploadHelpReq struct {
	RequestId int64 `json:"request_id"`
}

type UploadHelpRes struct {
}

type UploadApplyReq struct {
	ApplyId int64 `json:"apply_id"`
}

type UploadApplyRes struct {
}

// Code generated by goctl. DO NOT EDIT.
package types

type ApplyInfo struct {
	ApplyId     int64  `json:"apply_id"`
	ScholarName string `json:"scholar_name"`
	Institution string `json:"institution"`
	ApplyType   int64  `json:"apply_type"`
	Email       string `json:"email,optional"`
	URL         string `json:"url,optional"`
}

type GetApplyRequest struct {
	Start int64 `json:"start,optional"`
	End   int64 `json:"end,optional"`
}

type GetApplyResponse struct {
	Count     int64       `json:"count"`
	ApplyList []ApplyInfo `json:"records"`
}

type DealApplyRequest struct {
	ApplyId int64 `json:"apply_id"`
	IsAgree bool  `json:"is_agree"`
}

type EmailIdentifyRequest struct {
	Email     string `json:"email"`
	ScholarId int64  `json:"scholar_id"`
}

type CheckScholarResponse struct {
	Code      int64  `json:"code"`
	ScholarId int64  `json:"scholar_id,optional"`
	Msg       string `json:"msg,optional"`
}

type CheckUserRequest struct {
	ScholarId int64 `json:"scholar_id"`
}

type CheckUserResponse struct {
	Code   int64  `json:"code"`
	UserId int64  `json:"user_id,optional"`
	Msg    string `json:"msg,optional"`
}

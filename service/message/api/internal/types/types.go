// Code generated by goctl. DO NOT EDIT.
package types

type MessageInfo struct {
	MessageId   int64  `json:"apply_id"`
	MessageType string `json:"scholar_name"`
	UId         int64  `json:"institution"`
	GId         int64  `json:"apply_type"`
	PId         string `json:"email"`
	RId         string `json:"url"`
	Result      string `json:"result"`
	Content     string `json:"content"`
	Read        string `json:"read"`
	MessageTime string `json:"msg_time"`
}

type GetMessageRequest struct {
	Start int64 `json:"start,optinal"`
	End   int64 `json:"end,optinal"`
}

type GetMessageResponse struct {
	MessageList []MessageInfo `json:"messages"`
}

type ReadMessageRequest struct {
	MessageId int64 `json:"apply_id"`
}

type DeleteMessageRequest struct {
	MessageId int64 `json:"apply_id"`
}

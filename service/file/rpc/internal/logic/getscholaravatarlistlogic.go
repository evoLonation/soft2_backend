package logic

import (
	"context"
	"soft2_backend/service/apply/rpc/types/apply"
	"soft2_backend/service/file/filecommon"

	"soft2_backend/service/file/rpc/internal/svc"
	"soft2_backend/service/file/rpc/types/file"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScholarAvatarListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetScholarAvatarListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScholarAvatarListLogic {
	return &GetScholarAvatarListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetScholarAvatarListLogic) GetScholarAvatarList(in *file.ListScholarIdReq) (*file.ListUrlReply, error) {
	var userId int64
	ch := make(chan error, len(in.Ids))
	mp := make(map[string]string)
	go func() {
		for _, id := range in.GetIds() {
			res, err := l.svcCtx.Apply.CheckUser(l.ctx, &apply.CheckUserReq{ScholarId: id})
			if err != nil {
				ch <- err
			}
			if res.IsVerified {
				userId = res.UserId
				userAvatar, err := l.svcCtx.UserAvatarModel.FindOne(l.ctx, userId)
				err = filecommon.SqlErrorCheck(err)
				if err != nil && err != filecommon.NoRowError {
					ch <- err
				}
				if err == filecommon.NoRowError {
					mp[id] = filecommon.GetDefaultAvatarUrl()
				} else {
					mp[id] = filecommon.GetUrl(userAvatar.FileName)
				}
			} else {
				mp[id] = filecommon.GetDefaultAvatarUrl()
			}
			ch <- nil
		}
	}()

	urls := make([]*file.UrlReply, len(in.GetIds()))
	for _, _ = range in.GetIds() {
		err := <-ch
		if err != nil {
			return nil, err
		}
	}
	for i, id := range in.GetIds() {
		urls[i] = &file.UrlReply{
			Url: mp[id],
		}
	}
	return &file.ListUrlReply{Urls: urls}, nil
}

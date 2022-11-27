package logic

import (
	"context"
	uuid "github.com/nu7hatch/gouuid"
	"mime/multipart"
	"os"
	"soft2_backend/service/file/api/internal/common"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"soft2_backend/service/file/api/internal/svc"
)

type UploadAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	File   *multipart.FileHeader
}

func NewUploadAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadAvatarLogic {
	return &UploadAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadAvatarLogic) UploadAvatar() error {
	file, err := l.File.Open()

	if err != nil {
		return err
	}
	var content [common.DefaultMultipartMemory]byte

	fileLen, err := file.Read(content[:])
	if err != nil {
		return err
	}
	newUuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	tmp := strings.Split(l.File.Filename, ".")
	suffix := tmp[len(tmp)-1]
	newFile, err := os.Create("./localfile/" + newUuid.String() + "." + suffix)
	if err != nil {
		panic(err)
	}
	if _, err := newFile.Write(content[:fileLen]); err != nil {
		panic(err)
	}
	err = newFile.Close()
	if err != nil {
		panic(err)
	}
	return nil
}

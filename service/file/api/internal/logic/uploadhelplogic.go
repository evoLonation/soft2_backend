package logic

import (
	"archive/zip"
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"soft2_backend/service/file/filecommon"
	"soft2_backend/service/file/model"
	"soft2_backend/service/help/rpc/types/help"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"soft2_backend/service/file/api/internal/svc"
)

type UploadHelpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	Files  []struct {
		*multipart.FileHeader
		multipart.File
	}
	RequestId int64
}

func NewUploadHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadHelpLogic {
	return &UploadHelpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadHelpLogic) UploadHelp() error {
	defer func() {
		if err := os.RemoveAll(filecommon.FilePath + "temp/"); err != nil {
			panic(err)
		}
	}()
	for i, file := range l.Files {
		if _, err := filecommon.CreateTempFile(file.File, file.FileHeader, "file"+strconv.Itoa(i)); err != nil {
			return err
		}
	}
	zipfilename := filecommon.NewUUid() + ".zip"
	if err := zipSource(filecommon.FilePath+"temp/", filecommon.FilePath+zipfilename); err != nil {
		return err
	}
	userId := l.ctx.Value("userId").(int64)
	err := l.svcCtx.HelpFileModel.Delete(l.ctx, l.RequestId)
	err = filecommon.SqlErrorCheck(err)
	if err != nil && err != filecommon.NoRowError {
		return err
	}
	_, err = l.svcCtx.HelpFileModel.Insert(l.ctx, &model.HelpFile{HelpId: l.RequestId, FileName: zipfilename})
	err = filecommon.SqlErrorCheck(err)
	if err != nil {
		return err
	}
	_, err = l.svcCtx.Help.UpDateStatus(l.ctx, &help.UpdateReq{Status: 1, UserId: userId, RequestId: l.RequestId})
	if err != nil {
		return err
	}
	return nil
}

func zipSource(source, target string) error {
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	// 2. Go through all the files of the source
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 3. Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// 4. Set relative path of a file as the header name
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		// 5. Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}

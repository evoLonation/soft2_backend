package filecommon

import (
	"context"
	"encoding/json"
	"errors"
	uuid "github.com/nu7hatch/gouuid"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

const (
	FilePath          = "./localfile/"
	GetFileBaseUrl    = "http://120.46.220.182:8894/api/get-file/"
	DefaultAvatarName = "_defaultavatar.webp"
	DefaultLogoName   = "logo.jpg"
)
const (
	DefaultMultipartMemory = 32 << 20 // 32 MB
)

var NoRowError = errors.New("no rows")

func GetUrl(fileName string) string {
	split := strings.Split(fileName, ".")
	return GetFileBaseUrl + split[0] + "/" + split[1]
}

func GetDefaultAvatarUrl() string {
	return GetUrl(DefaultAvatarName)
}

func SqlErrorCheck(err error) error {
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return NoRowError
		} else {
			panic(err.Error())
		}
	}
	return nil
}

func CreateFile(file multipart.File, fileHeader *multipart.FileHeader, dir string, prefix string) (string, error) {
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	var content [DefaultMultipartMemory]byte

	fileLen, err := file.Read(content[:])
	if err != nil {
		return "", err
	}
	tmp := strings.Split(fileHeader.Filename, ".")
	suffix := tmp[len(tmp)-1]
	filename := prefix + "." + suffix
	newFile, err := os.Create(dir + filename)
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

	return filename, nil
}
func CreateLocalFile(file multipart.File, fileHeader *multipart.FileHeader, prefix string) (string, error) {
	return CreateFile(file, fileHeader, FilePath, prefix)
}

func CreateUUidFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	return CreateLocalFile(file, fileHeader, NewUUid())
}
func CreateTempFile(file multipart.File, fileHeader *multipart.FileHeader, prefix string) (string, error) {
	return CreateFile(file, fileHeader, FilePath+"temp/", prefix)
}

func RemoveTempFile(filename string) {
	os.Remove(FilePath + "temp/" + filename)
}

func NewUUid() string {
	newUuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return newUuid.String()
}

//func GetFormFile(w http.ResponseWriter, form *multipart.Form) *multipart.FileHeader {
//	if form.File["file"] == nil {
//		httpx.Error(w, errors.New("请求的form-data请包含file字段"))
//		return nil
//	}
//	return form.File["file"][0]
//}
//
//func GetFormValue(w http.ResponseWriter, form *multipart.Form, key string) (string, bool) {
//	if form.Value[key] == nil {
//		httpx.Error(w, errors.New("请求的form-data请包含"+key+"字段"))
//		return "", false
//	}
//	return form.Value[key][0], true
//}

func InitFile() {
	avatarSrc, err := os.Open("etc/" + DefaultAvatarName)
	if err != nil {
		panic(err)
	}
	logoSrc, err := os.Open("etc/" + DefaultLogoName)
	if err != nil {
		panic(err)
	}
	os.MkdirAll(FilePath, 0x777)
	//os.MkdirAll(FilePath + "temp/", 0x777)
	avatarDst, err := os.Create(FilePath + DefaultAvatarName)
	if err != nil {
		panic(err)
	}
	logoDst, err := os.Create(FilePath + DefaultLogoName)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(avatarDst, avatarSrc)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(logoDst, logoSrc)
	if err != nil {
		panic(err)
	}
}

func GetUserId(ctx context.Context) int64 {
	userId, err := ctx.Value("UserId").(json.Number).Int64()
	if err != nil {
		panic(err)
	}
	return userId
}

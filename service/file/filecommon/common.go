package filecommon

import (
	"errors"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const (
	FilePath          = "./localfile/"
	GetFileBaseUrl    = "120.46.220.182:8894/api/get-file/"
	DefaultAvatarName = "_defaultavatar.webp"
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

func CreateFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	var content [DefaultMultipartMemory]byte

	fileLen, err := file.Read(content[:])
	if err != nil {
		return "", err
	}
	newUuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	tmp := strings.Split(fileHeader.Filename, ".")
	suffix := tmp[len(tmp)-1]
	filename := newUuid.String() + "." + suffix
	newFile, err := os.Create(FilePath + filename)
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

func GetFormFile(w http.ResponseWriter, form *multipart.Form) *multipart.FileHeader {
	if form.File["file"] == nil {
		httpx.Error(w, errors.New("请求的form-data请包含file字段"))
		return nil
	}
	return form.File["file"][0]
}

func GetFormValue(w http.ResponseWriter, form *multipart.Form, key string) (string, bool) {
	if form.Value[key] == nil {
		httpx.Error(w, errors.New("请求的form-data请包含"+key+"字段"))
		return "", false
	}
	return form.Value[key][0], true
}

func InitFile() {
	src, err := os.Open("etc/" + DefaultAvatarName)
	if err != nil {
		panic(err)
	}
	dst, err := os.Open(FilePath + DefaultAvatarName)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
}

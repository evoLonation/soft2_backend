package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"soft2_backend/service/file/rpc/types/file"
	"soft2_backend/service/paper/database"

	"soft2_backend/service/paper/rpc/internal/svc"
	"soft2_backend/service/paper/rpc/paper"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCheckScholarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCheckScholarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCheckScholarLogic {
	return &ListCheckScholarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCheckScholarLogic) ListCheckScholar(in *paper.ListCheckScholarReq) (*paper.ListCreateScholarReply, error) {
	// todo: add your logic here and delete this line
	var buf bytes.Buffer
	var idList []string
	for _, id := range in.ScholarId {
		idList = append(idList, id)
	}
	query := map[string]interface{}{
		"ids": idList,
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	res := database.MgetScholer(buf)

	scholars := NilHandler(res["docs"], "list").([]interface{})
	scholarList := make([]*paper.CreateScholarReply, 0)
	avatarUrls, err := l.svcCtx.FileRpc.GetScholarAvatarList(l.ctx, &file.ListScholarIdReq{
		Ids: in.ScholarId,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for i, scholar := range scholars {
		if scholar.(map[string]interface{})["found"].(bool) == false {
			continue
		}
		source := scholar.(map[string]interface{})["_source"].(map[string]interface{})
		scholarList = append(scholarList, &paper.CreateScholarReply{
			ScholarName: NilHandler(source["name"], "string").(string),
			Org:         NilHandler(source["orgs"].([]interface{})[0], "string").(string),
			PaperNum:    int64(NilHandler(source["n_pubs"], "int").(int)),
			Url:         avatarUrls.Urls[i].Url,
		})
	}
	resp := &paper.ListCreateScholarReply{
		Scholars: scholarList,
	}
	return resp, nil
}

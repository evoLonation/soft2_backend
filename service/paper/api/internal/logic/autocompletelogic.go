package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"
	"soft2_backend/service/paper/database"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAutoCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoCompleteLogic {
	return &AutoCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AutoCompleteLogic) AutoComplete(req *types.AutoCompleteRequest) (resp *types.AutoCompleteResponse, err error) {
	// todo: add your logic here and delete this line
	var buf bytes.Buffer
	query := map[string]interface{}{
		"suggest": map[string]interface{}{
			"my_suggest": map[string]interface{}{
				"text": req.Text,
				"completion": map[string]interface{}{
					"field":           TranslateSearchType(req.SearchType),
					"skip_duplicates": true,
					"size":            10,
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	var res map[string]interface{}
	if req.SearchType == 6 {
		res = database.SearchAuthor(buf)
	} else {
		res = database.SearchPaper(buf)
	}

	var autoCompletes []string
	for _, hit := range res["suggest"].(map[string]interface{})["my_suggest"].([]interface{}) {
		for _, option := range hit.(map[string]interface{})["options"].([]interface{}) {
			autoCompletes = append(autoCompletes, option.(map[string]interface{})["text"].(string))
		}
	}
	resp = &types.AutoCompleteResponse{
		AutoCompletes: autoCompletes,
	}
	return resp, nil
}

func TranslateSearchType(searchType int) string {
	switch searchType {
	case 0:
		return "title"
	case 1:
		return "authors.name"
	case 2:
		return "keywords"
	case 3:
		return "abstract"
	case 4:
		return "doi"
	case 5:
		return "venue.raw"
	case 6:
		return "orgs"
	default:
		return ""
	}
}

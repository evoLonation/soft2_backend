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
	var buf bytes.Buffer
	if req.Size == 0 {
		req.Size = 10
	}
	query := map[string]interface{}{
		"suggest": map[string]interface{}{
			"my_suggest": map[string]interface{}{
				"text": req.Text,
				"completion": map[string]interface{}{
					"field":           "hot_word",
					"skip_duplicates": true,
					"size":            req.Size,
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	var res map[string]interface{}
	res, err = database.SearchAutoComplete(buf)
	if err != nil {
		return nil, err
	}
	var autoCompletes []string
	for _, hit := range res["suggest"].(map[string]interface{})["my_suggest"].([]interface{}) {
		for _, option := range hit.(map[string]interface{})["options"].([]interface{}) {
			autoCompletes = append(autoCompletes, option.(map[string]interface{})["text"].(string))
		}
	}
	if len(autoCompletes) != 0 {
		resp = &types.AutoCompleteResponse{
			AutoCompletes: autoCompletes,
		}
	} else {
		resp = &types.AutoCompleteResponse{
			AutoCompletes: make([]string, 0),
		}
	}
	return resp, nil
}

func TranslateSearchType(searchType int) string {
	switch searchType {
	case 0:
		return "title.completion"
	case 1:
		return "authors.name.completion"
	case 2:
		return "keywords.completion"
	case 3:
		return "abstract.completion"
	case 4:
		return "doi.completion"
	case 5:
		return "venue.completion"
	case 6:
		return "orgs.completion"
	default:
		return ""
	}
}

package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"soft2_backend/service/paper/database"

	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScholarClaimLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScholarClaimLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScholarClaimLogic {
	return &ScholarClaimLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScholarClaimLogic) ScholarClaim(req *types.ScholarClaimRequest) (resp *types.ScholarClaimResponse, err error) {
	// todo: add your logic here and delete this line
	var paperBuf bytes.Buffer
	paperQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": req.PaperId,
			},
		},
	}
	if err := json.NewEncoder(&paperBuf).Encode(paperQuery); err != nil {
		log.Printf("Error encoding query: %s", err)
	}
	log.Println(paperBuf.String())
	paperRes := database.SearchPaper(paperBuf)

	var scholarBuf bytes.Buffer
	scholarQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": req.ScholarId,
			},
		},
	}
	if err := json.NewEncoder(&scholarBuf).Encode(scholarQuery); err != nil {
		log.Printf("Error encoding query: %s", err)
	}
	log.Println(scholarBuf.String())
	scholarRes := database.SearchAuthor(scholarBuf)

	paperSource := paperRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	scholarSource := scholarRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})

	minLevenshtein := 100.0
	minIter := 0
	authors := paperSource["authors"].([]interface{})
	for i, author := range authors {
		if NilHandler(author.(map[string]interface{})["id"], "string").(string) == req.ScholarId {
			return &types.ScholarClaimResponse{
				Code:      0,
				ScholarId: "",
			}, nil
		}
		authorName := NilHandler(author.(map[string]interface{})["name"], "string").(string)
		if Levenshtein(authorName, scholarSource["name"].(string), 1, 1, 1) < minLevenshtein {
			minLevenshtein = Levenshtein(authorName, scholarSource["name"].(string), 1, 1, 1)
			minIter = i
		}
	}

	if NilHandler(authors[minIter].(map[string]interface{})["id"], "string").(string) == "" {
		authors[minIter].(map[string]interface{})["id"] = req.ScholarId
		authors[minIter].(map[string]interface{})["name"] = scholarSource["name"]

		var updatePaperBuf bytes.Buffer
		updatePaper := map[string]interface{}{
			"doc": map[string]interface{}{
				"authors": authors,
			},
		}
		if err := json.NewEncoder(&updatePaperBuf).Encode(updatePaper); err != nil {
			log.Printf("Error encoding query: %s\n", err)
		}
		log.Println(updatePaperBuf.String())
		updatePaperRes := database.UpdatePaper(updatePaperBuf, req.PaperId)
		if updatePaperRes["_shards"].(map[string]interface{})["successful"] != 1 {
			return nil, errors.New("update paper error")
		}

		var updateScholarBuf bytes.Buffer
		scholarPapers := scholarSource["pubs"].([]interface{})
		scholarPapers = append(scholarPapers, map[string]interface{}{
			"i": req.PaperId,
			"r": minIter,
		})
		updateScholar := map[string]interface{}{
			"doc": map[string]interface{}{
				"pubs": scholarPapers,
			},
		}
		if err := json.NewEncoder(&updateScholarBuf).Encode(updateScholar); err != nil {
			log.Printf("Error encoding query: %s\n", err)
		}
		log.Println(updateScholarBuf.String())
		updateScholarRes := database.UpdateAuthor(updateScholarBuf, req.ScholarId)
		if updateScholarRes["_shards"].(map[string]interface{})["successful"] != 1 {
			return nil, errors.New("update scholar error")
		}

		return &types.ScholarClaimResponse{
			Code:      0,
			ScholarId: "",
		}, nil
	} else {
		return &types.ScholarClaimResponse{
			Code:      1,
			ScholarId: NilHandler(authors[minIter].(map[string]interface{})["id"], "string").(string),
		}, nil
	}
}

func Levenshtein(str1, str2 string, costIns, costRep, costDel int) float64 {
	var maxLen = 255
	l1 := len(str1)
	l2 := len(str2)
	if l1 == 0 {
		return float64(l2 * costIns)
	}
	if l2 == 0 {
		return float64(l1 * costDel)
	}
	if l1 > maxLen || l2 > maxLen {
		return -1
	}

	tmp := make([]int, l2+1)
	p1 := make([]int, l2+1)
	p2 := make([]int, l2+1)
	var c0, c1, c2 int
	var i1, i2 int
	for i2 := 0; i2 <= l2; i2++ {
		p1[i2] = i2 * costIns
	}
	for i1 = 0; i1 < l1; i1++ {
		p2[0] = p1[0] + costDel
		for i2 = 0; i2 < l2; i2++ {
			if str1[i1] == str2[i2] {
				c0 = p1[i2]
			} else {
				c0 = p1[i2] + costRep
			}
			c1 = p1[i2+1] + costDel
			if c1 < c0 {
				c0 = c1
			}
			c2 = p2[i2] + costIns
			if c2 < c0 {
				c0 = c2
			}
			p2[i2+1] = c0
		}
		tmp = p1
		p1 = p2
		p2 = tmp
	}
	c0 = p1[l2]

	return float64(c0) / float64(l1)
}

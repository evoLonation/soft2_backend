package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"soft2_backend/service/paper/database"

	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScholarRelationNetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScholarRelationNetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScholarRelationNetLogic {
	return &ScholarRelationNetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var coNodes map[string]types.CoNetNodeJSON
var coNodeList []types.CoNetNodeJSON
var coEdges []types.EdgeJSON
var maxCoNum = 0
var minCoNum = 1000000
var maxCoCitation = 0
var minCoCitation = 1000000

var ciNodes map[string]types.CiNetNodeJSON
var ciNodeList []types.CiNetNodeJSON
var ciEdges []types.EdgeJSON
var maxCiNum = 0
var minCiNum = 1000000
var maxCiCitation = 0
var minCiCitation = 1000000

func (l *ScholarRelationNetLogic) ScholarRelationNet(req *types.ScholarRelationNetRequest) (resp *types.ScholarRelationNetResponse, err error) {
	// todo: add your logic here and delete this line
	coNodes = make(map[string]types.CoNetNodeJSON)
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

	scholarSource := scholarRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})

	majorCoNode := types.CoNetNodeJSON{
		Id:    scholarSource["id"].(string),
		Label: scholarSource["name"].(string),
		Size:  0,
		CoNum: 0,
		Type:  "major",
		Style: types.StyleJSON{
			Fill: NilHandler(scholarSource["n_citation"], "int").(string),
		},
	}
	majorCiNode := types.CiNetNodeJSON{
		Id:    scholarSource["id"].(string),
		Label: scholarSource["name"].(string),
		Size:  0,
		CiNum: 0,
		Type:  "major",
		Style: types.StyleJSON{
			Fill: NilHandler(scholarSource["n_citation"], "int").(string),
		},
	}
	coNodes[scholarSource["id"].(string)] = majorCoNode
	ciNodes[scholarSource["id"].(string)] = majorCiNode

	pubs := scholarSource["pubs"].([]interface{})
	pubIds := make([]string, 0)
	for _, pub := range pubs {
		pubIds = append(pubIds, pub.(map[string]interface{})["i"].(string))
	}

	var pubBuf bytes.Buffer
	pubQuery := map[string]interface{}{
		"ids": pubIds,
	}
	if err := json.NewEncoder(&pubBuf).Encode(pubQuery); err != nil {
		log.Printf("Error encoding query: %s", err)
	}
	log.Println(pubBuf.String())
	pubRes := database.MgetPaper(pubBuf)

	pubs = pubRes["docs"].([]interface{})
	for _, pub := range pubs {
		authors := pub.(map[string]interface{})["_source"].(map[string]interface{})["authors"].([]interface{})
		for _, author := range authors {
			authorId := NilHandler(author.(map[string]interface{})["id"].(string), "string").(string)
			if authorId != req.ScholarId {
				if _, ok := coNodes[authorId]; ok && authorId != "" {
					thisNode := coNodes[authorId]
					thisNode.CoNum++
					coNodes[authorId] = thisNode
				} else {
					coNode := types.CoNetNodeJSON{
						Id:    authorId,
						Label: author.(map[string]interface{})["name"].(string),
						Size:  0,
						CoNum: 1,
						Style: types.StyleJSON{
							Fill: NilHandler(author.(map[string]interface{})["n_citation"], "int").(string),
						},
					}
					coNodes[authorId] = coNode
				}
			}
			coEdges = append(coEdges, types.EdgeJSON{
				Source: req.ScholarId,
				Target: authorId,
			})
		}
	}

	for _, coNode := range coNodes {
		if coNode.CoNum > maxCoNum {
			maxCoNum = coNode.CoNum
		}
		if coNode.CoNum < minCoNum {
			minCoNum = coNode.CoNum
		}
		nCitation := NilHandler(coNode.Style.Fill, "int").(int)
		if nCitation > maxCoCitation {
			maxCoCitation = nCitation
		}
		if nCitation < minCoCitation {
			minCoCitation = nCitation
		}
	}

	for _, coNode := range coNodes {
		coNode.Size = GetSize(coNode.CoNum, maxCoNum, minCoNum)
		coNode.Style.Fill = GetColor(GetD(NilHandler(coNode.Style.Fill, "int").(int), maxCoCitation, minCoCitation))
		coNodeList = append(coNodeList, coNode)
	}

	for _, pub := range pubs {
		references := pub.(map[string]interface{})["_source"].(map[string]interface{})["references"].([]interface{})
		referenceIds := make([]string, 0)
		for _, reference := range references {
			referenceIds = append(referenceIds, reference.(string))
		}

		var referenceBuf bytes.Buffer
		referenceQuery := map[string]interface{}{
			"ids": referenceIds,
		}
		if err := json.NewEncoder(&referenceBuf).Encode(referenceQuery); err != nil {
			log.Printf("encode query error: %v", err)
		}
		referenceRes := database.MgetPaper(referenceBuf)

		references = referenceRes["docs"].([]interface{})
		for _, reference := range references {
			firstAuthor := reference.(map[string]interface{})["_source"].(map[string]interface{})["authors"].([]interface{})[0].(map[string]interface{})
			authorId := NilHandler(firstAuthor["id"].(string), "string").(string)
			if _, ok := ciNodes[authorId]; ok && authorId != "" {
				thisNode := ciNodes[authorId]
				thisNode.CiNum++
				ciNodes[authorId] = thisNode
			} else {
				ciNode := types.CiNetNodeJSON{
					Id:    authorId,
					Label: firstAuthor["name"].(string),
					Size:  0,
					CiNum: 1,
					Style: types.StyleJSON{
						Fill: NilHandler(firstAuthor["n_citation"], "int").(string),
					},
				}
				ciNodes[authorId] = ciNode
			}
			ciEdges = append(ciEdges, types.EdgeJSON{
				Source: req.ScholarId,
				Target: authorId,
			})
		}
	}

	for _, ciNode := range ciNodes {
		if ciNode.CiNum > maxCiNum {
			maxCiNum = ciNode.CiNum
		}
		if ciNode.CiNum < minCiNum {
			minCiNum = ciNode.CiNum
		}
		nCitation := NilHandler(ciNode.Style.Fill, "int").(int)
		if nCitation > maxCiCitation {
			maxCiCitation = nCitation
		}
		if nCitation < minCiCitation {
			minCiCitation = nCitation
		}
	}

	for _, ciNode := range ciNodes {
		ciNode.Size = GetSize(ciNode.CiNum, maxCiNum, minCiNum)
		ciNode.Style.Fill = GetColor(GetD(NilHandler(ciNode.Style.Fill, "int").(int), maxCiCitation, minCiCitation))
		ciNodeList = append(ciNodeList, ciNode)
	}

	resp = &types.ScholarRelationNetResponse{
		CoNet: types.CoNetJSON{
			CoNetNodes: coNodeList,
			CoNetEdges: coEdges,
		},
		CiNet: types.CiNetJSON{
			CiNetNodes: ciNodeList,
			CiNetEdges: ciEdges,
		},
	}
	return resp, nil
}

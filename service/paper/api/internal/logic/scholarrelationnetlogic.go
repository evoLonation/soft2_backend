package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"soft2_backend/service/paper/database"
	"strconv"

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

func (l *ScholarRelationNetLogic) ScholarRelationNet(req *types.ScholarRelationNetRequest) (resp *types.ScholarRelationNetResponse, err error) {
	// todo: add your logic here and delete this line
	var coNodes = make(map[string]types.CoNetNodeJSON, 0)
	var coNodeList = make([]types.CoNetNodeJSON, 0)
	var coEdges = make([]types.EdgeJSON, 0)
	var maxCoNum = 0
	var minCoNum = 1000000
	var maxCoCitation = 0
	var minCoCitation = 1000000

	var ciNodes = make(map[string]types.CiNetNodeJSON, 0)
	var ciNodeList = make([]types.CiNetNodeJSON, 0)
	var ciEdges = make([]types.EdgeJSON, 0)
	var maxCiNum = 0
	var minCiNum = 1000000
	var maxCiCitation = 0
	var minCiCitation = 1000000

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

	hits := NilHandler(scholarRes["hits"].(map[string]interface{})["hits"], "list").([]interface{})
	scholarSource := hits[0].(map[string]interface{})["_source"].(map[string]interface{})

	majorCoNode := types.CoNetNodeJSON{
		Id:    NilHandler(scholarSource["id"], "string").(string),
		Label: NilHandler(scholarSource["name"], "string").(string),
		Size:  0,
		CoNum: 0,
		Type:  "major",
		Style: types.StyleJSON{
			Fill: strconv.Itoa(NilHandler(scholarSource["n_citation"], "int").(int)),
		},
	}
	majorCiNode := types.CiNetNodeJSON{
		Id:    NilHandler(scholarSource["id"], "string").(string),
		Label: NilHandler(scholarSource["name"], "string").(string),
		Size:  0,
		CiNum: 0,
		Type:  "major",
		Style: types.StyleJSON{
			Fill: strconv.Itoa(NilHandler(scholarSource["n_citation"], "int").(int)),
		},
	}
	coNodes[scholarSource["id"].(string)] = majorCoNode
	ciNodes[scholarSource["id"].(string)] = majorCiNode
	coNodeList = append(coNodeList, majorCoNode)
	ciNodeList = append(ciNodeList, majorCiNode)

	pubs := NilHandler(scholarSource["pubs"], "list").([]interface{})
	pubIds := make([]string, 0)
	for _, pub := range pubs {
		pubIds = append(pubIds, pub.(map[string]interface{})["i"].(string))
	}
	//log.Printf("pubIds: %v", pubIds)

	var pubBuf bytes.Buffer
	pubQuery := map[string]interface{}{
		"ids": pubIds,
	}
	if err := json.NewEncoder(&pubBuf).Encode(pubQuery); err != nil {
		log.Printf("Error encoding query: %s", err)
	}
	//log.Println(pubBuf.String())
	pubRes := database.MgetPaper(pubBuf)

	pubs = pubRes["docs"].([]interface{})
	pubCnt := 0
	for _, pub := range pubs {
		if pub.(map[string]interface{})["found"].(bool) == false {
			continue
		}
		if pubCnt > 8 {
			break
		}
		pubCnt++

		authors := NilHandler(pub.(map[string]interface{})["_source"].(map[string]interface{})["authors"], "list").([]interface{})
		for _, author := range authors {
			authorId := NilHandler(author.(map[string]interface{})["id"].(string), "string").(string)
			if authorId == req.ScholarId {
				continue
			}
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
						Fill: strconv.Itoa(NilHandler(author.(map[string]interface{})["n_citation"], "int").(int)),
					},
				}
				coNodes[authorId] = coNode
			}
		}
	}

	for _, coNode := range coNodes {
		if coNode.CoNum > maxCoNum {
			maxCoNum = coNode.CoNum
		}
		if coNode.CoNum < minCoNum {
			minCoNum = coNode.CoNum
		}
		nCitation, _ := strconv.Atoi(NilHandler(coNode.Style.Fill, "string").(string))
		if nCitation > maxCoCitation {
			maxCoCitation = nCitation
		}
		if nCitation < minCoCitation {
			minCoCitation = nCitation
		}
	}

	coNodeCnt := 0
	for _, coNode := range coNodes {
		if coNodeCnt > 20 {
			break
		}
		coNodeCnt++

		nCitation, _ := strconv.Atoi(NilHandler(coNode.Style.Fill, "string").(string))
		coNode.Size = GetSize(coNode.CoNum, maxCoNum, minCoNum)
		coNode.Style.Fill = GetColor(GetD(nCitation, maxCoCitation, minCoCitation))
		coNodeList = append(coNodeList, coNode)
		if req.ScholarId != coNode.Id {
			coEdges = append(coEdges, types.EdgeJSON{
				Source: req.ScholarId,
				Target: coNode.Id,
			})
		}
	}
	nCitation, _ := strconv.Atoi(NilHandler(coNodeList[0].Style.Fill, "string").(string))
	coNodeList[0].Size = GetSize(coNodeList[0].CoNum, maxCoNum, minCoNum)
	coNodeList[0].Style.Fill = GetColor(GetD(nCitation, maxCoCitation, minCoCitation))

	for _, pub := range pubs {
		references := NilHandler(pub.(map[string]interface{})["_source"].(map[string]interface{})["references"], "list").([]interface{})
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

		references = NilHandler(referenceRes["docs"], "list").([]interface{})
		referenceCnt := 0
		for _, reference := range references {
			if reference.(map[string]interface{})["found"].(bool) == false {
				continue
			}
			if referenceCnt > 8 {
				break
			}
			referenceCnt++

			firstAuthors := NilHandler(reference.(map[string]interface{})["_source"].(map[string]interface{})["authors"], "list").([]interface{})
			var firstAuthor map[string]interface{}
			if len(firstAuthors) == 0 {
				firstAuthor = map[string]interface{}{
					"id":   "",
					"name": "",
				}
			} else {
				firstAuthor = firstAuthors[0].(map[string]interface{})
			}
			authorId := NilHandler(firstAuthor["id"].(string), "string").(string)
			if authorId == req.ScholarId {
				continue
			}

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
						Fill: strconv.Itoa(NilHandler(firstAuthor["n_citation"], "int").(int)),
					},
				}
				ciNodes[authorId] = ciNode
			}
		}
	}

	for _, ciNode := range ciNodes {
		if ciNode.CiNum > maxCiNum {
			maxCiNum = ciNode.CiNum
		}
		if ciNode.CiNum < minCiNum {
			minCiNum = ciNode.CiNum
		}
		nCitation, _ := strconv.Atoi(NilHandler(ciNode.Style.Fill, "string").(string))
		if nCitation > maxCiCitation {
			maxCiCitation = nCitation
		}
		if nCitation < minCiCitation {
			minCiCitation = nCitation
		}
	}

	ciNodeCnt := 0
	for _, ciNode := range ciNodes {
		if ciNodeCnt > 20 {
			break
		}
		ciNodeCnt++

		nCitation, _ := strconv.Atoi(NilHandler(ciNode.Style.Fill, "string").(string))
		ciNode.Size = GetSize(ciNode.CiNum, maxCiNum, minCiNum)
		ciNode.Style.Fill = GetColor(GetD(nCitation, maxCiCitation, minCiCitation))
		ciNodeList = append(ciNodeList, ciNode)
		if req.ScholarId != ciNode.Id {
			ciEdges = append(ciEdges, types.EdgeJSON{
				Source: req.ScholarId,
				Target: ciNode.Id,
			})
		}
	}
	nCitation, _ = strconv.Atoi(NilHandler(ciNodeList[0].Style.Fill, "string").(string))
	ciNodeList[0].Size = GetSize(ciNodeList[0].CiNum, maxCiNum, minCiNum)
	ciNodeList[0].Style.Fill = GetColor(GetD(nCitation, maxCiCitation, minCiCitation))

	coNodeList[0].CoNum = 0
	ciNodeList[0].Size = 40
	ciNodeList[0].CiNum = 0

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

package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"
	"soft2_backend/service/paper/database"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaperCiteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaperCiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaperCiteLogic {
	return &PaperCiteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaperCiteLogic) PaperCite(req *types.PaperCiteRequest) (resp *types.PaperCiteResponse, err error) {
	// todo: add your logic here and delete this line
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": req.Id,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	res := database.SearchPaper(buf)

	if res["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64) == 0 {
		return nil, nil
	}
	source := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	// GB/T 7714
	var gbCite string
	gbCite += "[1] "
	gbCite += NilHandler(source["authors"].([]interface{})[0].(map[string]interface{})["name"], "string").(string) + "."
	gbCite += source["title"].(string)
	gbCite += "[" + ParseGBDocType(NilHandler(source["doc_type"], "string").(string)) + "]."
	gbCite += source["venue"].(map[string]interface{})["raw"].(string) + ", "
	gbCite += source["year"].(string) + ", "
	gbCite += source["volume"].(string) + "(" + source["issue"].(string) + "): "
	gbCite += source["page_start"].(string) + "-" + source["page_end"].(string) + "."
	// MLA
	var mlaCite string
	author := NilHandler(source["authors"].([]interface{})[0].(map[string]interface{})["name"], "string").(string)
	authorName := strings.Split(author, " ")
	if len(authorName) >= 2 {
		author = authorName[len(authorName)-1] + ", " + strings.Join(authorName[:len(authorName)-1], " ")
	} else {
		author = authorName[0]
	}
	mlaCite += author + ". "
	mlaCite += "\"" + source["title"].(string) + ".\" "
	mlaCite += source["venue"].(map[string]interface{})["raw"].(string) + ", "
	mlaCite += "vol. " + source["volume"].(string) + ", no. " + source["issue"].(string) + ", "
	mlaCite += source["year"].(string) + ", "
	mlaCite += "pp. " + source["page_start"].(string) + "-" + source["page_end"].(string) + "."
	// APA
	var apaCite string
	author = NilHandler(source["authors"].([]interface{})[0].(map[string]interface{})["name"], "string").(string)
	authorName = strings.Split(author, " ")
	if len(authorName) >= 2 {
		author = authorName[len(authorName)-1] + ", "
		for i := 0; i < len(authorName)-1; i++ {
			author += strings.ToUpper(authorName[i][:1]) + ". "
		}
	} else {
		author = authorName[0] + ", "
	}
	apaCite += author + "(" + source["year"].(string) + "). "
	apaCite += source["title"].(string) + ". "
	apaCite += source["venue"].(map[string]interface{})["raw"].(string) + ", "
	apaCite += "(" + source["issue"].(string) + "), "
	apaCite += source["page_start"].(string) + "-" + source["page_end"].(string) + "."
	// Bibtex
	var bibtex string
	bibtex += "@" + ParseBibtexDocType(NilHandler(source["doc_type"], "string").(string)) + "{CiteKey" + ParseBibtexDocType(NilHandler(source["doc_type"], "string").(string)) + ",\n"
	bibtex += "  title\t= " + source["title"].(string) + ",\n"
	bibtex += "  author\t= " + NilHandler(source["authors"].([]interface{})[0].(map[string]interface{})["name"], "string").(string) + ",\n"
	bibtex += "  journal\t= " + source["venue"].(map[string]interface{})["raw"].(string) + ",\n"
	bibtex += "  year\t= " + source["year"].(string) + ",\n"
	bibtex += "  volume\t= " + source["volume"].(string) + ",\n"
	bibtex += "  number\t= " + source["issue"].(string) + ",\n"
	bibtex += "  pages\t= " + source["page_start"].(string) + "-" + source["page_end"].(string) + ",\n"
	bibtex += "}"
	// generate caj_cd cite string
	var cajCdCite string
	cajCdCite += "[1]"
	cajCdCite += NilHandler(source["authors"].([]interface{})[0].(map[string]interface{})["name"], "string").(string) + ". "
	cajCdCite += source["title"].(string)
	cajCdCite += "[" + ParseGBDocType(NilHandler(source["doc_type"], "string").(string)) + "]."
	cajCdCite += source["venue"].(map[string]interface{})["raw"].(string) + ","
	cajCdCite += source["year"].(string) + "."
	cajCdCite += source["page_start"].(string) + "-" + source["page_end"].(string) + "."

	resp = &types.PaperCiteResponse{
		Gb:     gbCite,
		Mla:    mlaCite,
		Apa:    apaCite,
		Bibtex: bibtex,
		CajCd:  cajCdCite,
	}
	return resp, nil
}

func ParseGBDocType(docType string) string {
	docType = strings.ToLower(docType)
	switch docType {
	case "book":
		return "M"
	case "conference":
		return "C"
	case "journal":
		return "J"
	case "dissertation":
		return "D"
	case "patent":
		return "P"
	case "standard":
		return "S"
	case "newspaper":
		return "N"
	case "report":
		return "R"
	case "database":
		return "DB"
	default:
		return "J"
	}
}

func ParseBibtexDocType(docType string) string {
	docType = strings.ToLower(docType)
	switch docType {
	case "book":
		return "book"
	case "conference":
		return "coference"
	case "journal":
		return "article"
	case "patent":
		return "misc"
	case "standard":
		return "standard"
	case "newspaper":
		return "article"
	case "report":
		return "techreport"
	case "database":
		return "database"
	default:
		return "article"
	}
}

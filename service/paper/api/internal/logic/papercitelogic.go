package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"
	"soft2_backend/service/paper/database"
	"strconv"
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
	authors := NilHandler(source["authors"], "list").([]interface{})
	var author string
	if len(authors) == 0 {
		author = ""
	} else {
		author = NilHandler(authors[0].(map[string]interface{})["name"], "string").(string)
	}
	// GB/T 7714
	var gbCite string
	gbCite += "[1] "
	gbCite += author + "."
	gbCite += NilHandler(source["title"], "string").(string)
	gbCite += "[" + ParseGBDocType(NilHandler(source["doc_type"], "string").(string)) + "]."
	gbCite += NilHandler(source["venue"], "string").(string) + ", "
	gbCite += strconv.Itoa(NilHandler(source["year"], "int").(int)) + ", "
	gbCite += NilHandler(source["volume"], "string").(string) + "(" + NilHandler(source["issue"], "string").(string) + "): "
	gbCite += NilHandler(source["page_start"], "string").(string) +
		"-" + NilHandler(source["page_end"], "string").(string) + "."
	log.Println(gbCite)

	// MLA
	var mlaCite string
	authorName := strings.Split(author, " ")
	var mlaAuthor string
	if len(authorName) >= 2 {
		mlaAuthor = authorName[len(authorName)-1] + ", " + strings.Join(authorName[:len(authorName)-1], " ")
	} else {
		mlaAuthor = authorName[0]
	}
	mlaCite += mlaAuthor + ". "
	mlaCite += "\"" + NilHandler(source["title"], "string").(string) + ".\" "
	mlaCite += NilHandler(source["venue"], "string").(string) + ", "
	mlaCite += "vol. " + NilHandler(source["volume"], "string").(string)
	mlaCite += ", no. " + NilHandler(source["issue"], "string").(string) + ", "
	mlaCite += strconv.Itoa(NilHandler(source["year"], "int").(int)) + ", "
	mlaCite += "pp. " + NilHandler(source["page_start"], "string").(string) +
		"-" + NilHandler(source["page_end"], "string").(string) + "."
	log.Println(mlaCite)

	// APA
	var apaCite string
	authorName = strings.Split(author, " ")
	var apaAuthor string
	if len(authorName) >= 2 {
		apaAuthor = authorName[len(authorName)-1] + ", "
		for i := 0; i < len(authorName)-1; i++ {
			apaAuthor += strings.ToUpper(authorName[i][:1]) + ". "
		}
	} else {
		apaAuthor = authorName[0] + ", "
	}
	apaCite += apaAuthor + "(" + strconv.Itoa(NilHandler(source["year"], "int").(int)) + "). "
	apaCite += NilHandler(source["title"], "string").(string) + ". "
	apaCite += NilHandler(source["venue"], "string").(string) + ", "
	apaCite += "(" + NilHandler(source["issue"], "string").(string) + "), "
	apaCite += NilHandler(source["page_start"], "string").(string) +
		"-" + NilHandler(source["page_end"], "string").(string) + "."
	log.Println(apaCite)

	// Bibtex
	var bibtex string
	bibtex += "@" + ParseBibtexDocType(NilHandler(source["doc_type"], "string").(string))
	bibtex += "{CiteKey" + ParseBibtexDocType(NilHandler(source["doc_type"], "string").(string)) + ",\n"
	bibtex += "  title\t= " + NilHandler(source["title"], "string").(string) + ",\n"
	bibtex += "  author\t= " + author + ",\n"
	bibtex += "  journal\t= " + NilHandler(source["venue"], "string").(string) + ",\n"
	bibtex += "  year\t= " + strconv.Itoa(NilHandler(source["year"], "int").(int)) + ",\n"
	bibtex += "  volume\t= " + NilHandler(source["volume"], "string").(string) + ",\n"
	bibtex += "  number\t= " + NilHandler(source["issue"], "string").(string) + ",\n"
	bibtex += "  pages\t= " + NilHandler(source["page_start"], "string").(string) +
		"-" + NilHandler(source["page_end"], "string").(string) + ",\n"
	bibtex += "}"
	log.Println(bibtex)

	// generate caj_cd cite string
	var cajCdCite string
	cajCdCite += "[1]"
	cajCdCite += author + ". "
	cajCdCite += NilHandler(source["title"], "string").(string)
	cajCdCite += "[" + ParseGBDocType(NilHandler(source["doc_type"], "string").(string)) + "]."
	cajCdCite += NilHandler(source["venue"], "string").(string) + ","
	cajCdCite += strconv.Itoa(NilHandler(source["year"], "int").(int)) + "."
	cajCdCite += NilHandler(source["page_start"], "string").(string) +
		"-" + NilHandler(source["page_end"], "string").(string) + "."
	log.Println(cajCdCite)

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
	return "J"
	//docType = strings.ToLower(docType)
	//switch docType {
	//case "book":
	//	return "M"
	//case "conference":
	//	return "C"
	//case "journal":
	//	return "J"
	//case "dissertation":
	//	return "D"
	//case "patent":
	//	return "P"
	//case "standard":
	//	return "S"
	//case "newspaper":
	//	return "N"
	//case "report":
	//	return "R"
	//case "database":
	//	return "DB"
	//default:
	//	return "J"
	//}
}

func ParseBibtexDocType(docType string) string {
	return "article"
	//docType = strings.ToLower(docType)
	//switch docType {
	//case "book":
	//	return "book"
	//case "conference":
	//	return "coference"
	//case "journal":
	//	return "article"
	//case "patent":
	//	return "misc"
	//case "standard":
	//	return "standard"
	//case "newspaper":
	//	return "article"
	//case "report":
	//	return "techreport"
	//case "database":
	//	return "database"
	//default:
	//	return "article"
	//}
}

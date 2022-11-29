// Code generated by goctl. DO NOT EDIT.
package types

type PaperRequest struct {
	Content   []PaperRequestJSON `json:"search_content"`
	StartYear int                `json:"start_year"`
	EndYear   int                `json:"end_year"`
	SortType  int                `json:"sort_type"`
	Start     int                `json:"start"`
	End       int                `json:"end"`
}

type PaperRequestJSON struct {
	Type       int    `json:"type"`
	SearchType int    `json:"search_type"`
	IsExact    int    `json:"is_exact"`
	Content    string `json:"content"`
}

type PaperResponse struct {
	PaperNum int                 `json:"paper_num"`
	Papers   []PaperResponseJSON `json:"papers"`
	Themes   []string            `json:"themes"`
	Years    []int               `json:"years"`
}

type PaperResponseJSON struct {
	Title     string       `json:"title"`
	Abstract  string       `json:"abstract"`
	Authors   []AuthorJSON `json:"authors"`
	Year      int          `json:"year"`
	NCitation int          `json:"n_citation"`
	Publisher string       `json:"publisher"`
}

type AuthorJSON struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	HasId bool   `json:"has_id"`
}

type ScholarRequest struct {
	Name        string `json:"name, optional"`
	Institution string `json:"institution, optional"`
	Start       int    `json:"start"`
	End         int    `json:"end"`
}

type ScholarResponse struct {
	ScholarNum int                   `json:"scholar_num"`
	Scholar    []ScholarResponseJSON `json:"scholar"`
}

type ScholarResponseJSON struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Institution []string `json:"institution"`
	PaperNum    int      `json:"paper_num"`
}

type AutoCompleteRequest struct {
	SearchType int    `json:"search_type"`
	Text       string `json:"text"`
}

type AutoCompleteResponse struct {
	AutoCompletes []string `json:"auto_completes"`
}

type PaperDetailRequest struct {
	Id string `json:"id"`
}

type PaperDetailResponse struct {
	Title      string          `json:"title"`
	Abstract   string          `json:"abstract"`
	Authors    []AuthorJSON    `json:"authors"`
	Doi        string          `json:"doi"`
	ISBN       string          `json:"isbn"`
	Org        string          `json:"org"`
	Keywords   []string        `json:"keywords"`
	Year       int             `json:"year"`
	NCitation  int             `json:"n_citation"`
	Publisher  string          `json:"publisher"`
	References []ReferenceJSON `json:"references"`
	Urls       []string        `json:"urls"`
}

type ReferenceJSON struct {
	Id     string   `json:"id"`
	Title  string   `json:"title"`
	Author []string `json:"author"`
	Year   int      `json:"year"`
}

type PaperCiteRequest struct {
	Id string `json:"id"`
}

type PaperCiteResponse struct {
	Gb     string `json:"gb"`
	Mla    string `json:"mla"`
	Apa    string `json:"apa"`
	Bibtex string `json:"bibix"`
	CajCd  string `json:"caj_cd"`
}

type PaperRelationNetRequest struct {
	Id string `json:"id"`
}

type PaperRelationNetResponse struct {
	Nodes []PaperNodeJSON `json:"nodes"`
	Edges []PaperEdgeJSON `json:"edges"`
}

type PaperNodeJSON struct {
	Id    string    `json:"id"`
	Label string    `json:"label"`
	Size  int       `json:"size"`
	Type  string    `json:"type"`
	Style StyleJSON `json:"style"`
	Info  InfoJSON  `json:"info"`
}

type StyleJSON struct {
	Fill string `json:"fill"`
}

type InfoJSON struct {
	Id       string       `json:"id"`
	Title    string       `json:"title"`
	Abstract string       `json:"abstract"`
	Authors  []AuthorJSON `json:"authors"`
	Year     int          `json:"year"`
}

type PaperEdgeJSON struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type ScholarBasicRequest struct {
	ScholarId string `json:"scholar_id"`
}

type ScholarBasicResponse struct {
	ScholarId   string    `json:"scholar_id"`
	Name        string    `json:"name"`
	Institution []string  `json:"institution"`
	Position    string    `json:"position"`
	RefNum      int       `json:"ref_num"`
	AchNum      int       `json:"ach_num"`
	HIndex      int       `json:"h_index"`
	Years       []int     `json:"years"`
	Tags        []TagJSON `json:"tags"`
}

type TagJSON struct {
	T string `json:"t"`
	W int    `json:"w"`
}

type ScholarCooperationRequest struct {
	ScholarId string `json:"scholar_id"`
}

type ScholarCooperationResponse struct {
	CoopList []CoopJSON `json:"coop_list"`
}

type CoopJSON struct {
	ScholarId   string `json:"id"`
	Name        string `json:"name"`
	Institution string `json:"institution"`
	Time        int    `json:"time"`
}

type ScholarPapersRequest struct {
	ScholarId string `json:"scholar_id"`
	IsFirst   bool   `json:"is_first, optional"`
	Year      int    `json:"year, optional"`
	TimeOrder bool   `json:"time_order, optional"`
	Start     int    `json:"start"`
	End       int    `json:"end"`
}

type ScholarPapersResponse struct {
	PaperNum int                 `json:"paper_num"`
	Papers   []PaperResponseJSON `json:"papers"`
}

type ScholarBarchartRequest struct {
	ScholarId string `json:"scholar_id"`
}

type ScholarBarchartResponse struct {
	Achievements int `json:"achievements"`
	References   int `json:"references"`
}

type ScholarClaimRequest struct {
	PaperId   string `json:"paper_id"`
	ScholarId string `json:"scholar_id"`
}

type ScholarClaimResponse struct {
	Code      int    `json:"code"`
	ScholarId string `json:"scholar_id"`
}

type MovePaperRequest struct {
	PaperId  string `json:"paper_id"`
	OwnerId  string `json:"owner_id"`
	TargetId string `json:"target_id"`
}

type MovePaperResponse struct {
}

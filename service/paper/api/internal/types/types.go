// Code generated by goctl. DO NOT EDIT.
package types

type PaperRequest struct {
	NeedFilterStatistics bool     `json:"need_filter_statistics,optional"`
	Query                string   `json:"query"`
	Years                []int    `json:"years,optional"`
	Themes               []string `json:"themes,optional"`
	Venues               []string `json:"venues,optional"`
	Institutions         []string `json:"institutions,optional"`
	StartYear            int      `json:"start_year,optional"`
	EndYear              int      `json:"end_year,optional"`
	SortType             int      `json:"sort_type,optional"`
	Start                int      `json:"start"`
	End                  int      `json:"end"`
}

type PaperRequestJSON struct {
	Type       int    `json:"type"`
	SearchType int    `json:"search_type"`
	IsExact    int    `json:"is_exact"`
	Content    string `json:"content"`
}

type Statistic struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type StatisticNumber struct {
	Name  int `json:"name"`
	Count int `json:"count"`
}

type PaperResponse struct {
	PaperNum     int                 `json:"paper_num"`
	Papers       []PaperResponseJSON `json:"papers"`
	Themes       []Statistic         `json:"themes"`
	Years        []StatisticNumber   `json:"years"`
	Institutions []Statistic         `json:"institutions"`
	Venues       []Statistic         `json:"venues"`
}

type PaperResponseJSON struct {
	Id        string       `json:"id"`
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
	Id          string `json:"id"`
	Name        string `json:"name"`
	Institution string `json:"institution"`
	PaperNum    int    `json:"paper_num"`
}

type AutoCompleteRequest struct {
	Size int    `json:"size,optional"`
	Text string `json:"text"`
}

type AutoCompleteResponse struct {
	AutoCompletes []string `json:"auto_completes"`
}

type PaperDetailRequest struct {
	Id string `json:"id"`
}

type PaperDetailResponse struct {
	Id         string       `json:"id"`
	Title      string       `json:"title"`
	Abstract   string       `json:"abstract"`
	Authors    []AuthorJSON `json:"authors"`
	Doi        string       `json:"doi"`
	ISBN       string       `json:"isbn"`
	ISSN       string       `json:"issn"`
	Org        string       `json:"org"`
	Keywords   []string     `json:"keywords"`
	Year       int          `json:"year"`
	NCitation  int          `json:"n_citation"`
	Publisher  string       `json:"publisher"`
	References []PaperJSON  `json:"references"`
	Similars   []PaperJSON  `json:"similars"`
	Urls       []string     `json:"urls"`
}

type PaperJSON struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type PaperCiteRequest struct {
	Id string `json:"id"`
}

type PaperCiteResponse struct {
	Gb     string `json:"gb"`
	Mla    string `json:"mla"`
	Apa    string `json:"apa"`
	Bibtex string `json:"bibtex"`
	CajCd  string `json:"caj_cd"`
}

type PaperRelationNetRequest struct {
	Id string `json:"id"`
}

type PaperRelationNetResponse struct {
	Nodes []PaperNodeJSON `json:"nodes"`
	Edges []EdgeJSON      `json:"edges"`
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

type EdgeJSON struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type PaperSimilarNetRequest struct {
	Id string `json:"id"`
}

type PaperSimilarNetResponse struct {
	Nodes []PaperNodeJSON `json:"nodes"`
	Edges []EdgeJSON      `json:"edges"`
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
	ScholarId   string   `json:"id"`
	Name        string   `json:"name"`
	Institution []string `json:"institution"`
	Time        int      `json:"time"`
}

type ScholarPapersRequest struct {
	ScholarId string `json:"scholar_id"`
	IsFirst   int    `json:"is_first"`
	Year      int    `json:"year, optional"`
	TimeOrder bool   `json:"time_order, optional"`
	Start     int    `json:"start"`
	End       int    `json:"end"`
}

type ScholarPapersResponse struct {
	PaperNum  int                 `json:"paper_num"`
	StartYear int                 `json:"start_year"`
	EndYear   int                 `json:"end_year"`
	Papers    []PaperResponseJSON `json:"papers"`
}

type ScholarBarchartRequest struct {
	ScholarId string `json:"scholar_id"`
}

type ScholarBarchartResponse struct {
	Statistic []StatisticJSON `json:"statistic"`
}

type StatisticJSON struct {
	Year         int `json:"year"`
	Achievements int `json:"achievements"`
	References   int `json:"references"`
}

type ScholarRelationNetRequest struct {
	ScholarId string `json:"scholar_id"`
}

type ScholarRelationNetResponse struct {
	CoNet CoNetJSON `json:"co_net"`
	CiNet CiNetJSON `json:"ci_net"`
}

type CoNetJSON struct {
	CoNetNodes []CoNetNodeJSON `json:"nodes"`
	CoNetEdges []EdgeJSON      `json:"edges"`
}

type CoNetNodeJSON struct {
	Id    string    `json:"id"`
	Label string    `json:"label"`
	Size  int       `json:"size"`
	CoNum int       `json:"co_num"`
	Type  string    `json:"type"`
	Style StyleJSON `json:"style"`
}

type CiNetJSON struct {
	CiNetNodes []CiNetNodeJSON `json:"nodes"`
	CiNetEdges []EdgeJSON      `json:"edges"`
}

type CiNetNodeJSON struct {
	Id    string    `json:"id"`
	Label string    `json:"label"`
	Size  int       `json:"size"`
	CiNum int       `json:"ci_num"`
	Type  string    `json:"type"`
	Style StyleJSON `json:"style"`
}

type ScholarClaimRequest struct {
	PaperId   string `json:"paper_id"`
	ScholarId string `json:"scholar_id"`
}

type ScholarClaimResponse struct {
	Code      int    `json:"code"`
	ScholarId string `json:"scholar_id"`
}

type ScholarGetAvatarRequest struct {
	ScholarId string `json:"scholar_id"`
}

type ScholarGetAvatarResponse struct {
	Code int    `json:"code"`
	Url  string `json:"url"`
}

type MovePaperRequest struct {
	PaperId  string `json:"paper_id"`
	OwnerId  string `json:"owner_id"`
	TargetId string `json:"target_id"`
}

type MovePaperResponse struct {
}

type FieldPaperRequest struct {
	Field string `json:"field"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

type FieldPaperResponse struct {
	PaperNum int                 `json:"paper_num"`
	Papers   []PaperResponseJSON `json:"papers"`
}

type FieldScholarRequest struct {
	Field string `json:"field"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

type FieldScholarResponse struct {
	ScholarNum int                `json:"scholar_num"`
	Scholars   []FieldScholarJSON `json:"scholars"`
}

type FieldScholarJSON struct {
	ScholarId string  `json:"scholar_id"`
	Name      string  `json:"name"`
	NPaper    int     `json:"n_paper"`
	NCitation int     `json:"n_citation"`
	Weight    float64 `json:"weight"`
}

type HomeInfoRequest struct {
	AreasNum   int `json:"areas_num, optional"`
	PaperNum   int `json:"paper_num"`
	ScholarNum int `json:"scholar_num"`
	JournalNum int `json:"journal_num"`
}

type HomeInfoResponse struct {
	Areas []AreaJSON `json:"areas"`
}

type AreaJSON struct {
	Type     []string          `json:"type"`
	Papers   []PaperInfoJSON   `json:"papers"`
	Scholars []ScholarInfoJSON `json:"scholars"`
	Journals []string          `json:"journals"`
}

type PaperInfoJSON struct {
	Title     string   `json:"title"`
	Authors   []string `json:"authors"`
	NCitation int      `json:"n_citation"`
}

type ScholarInfoJSON struct {
	ScholarId string `json:"scholar_id"`
	Name      string `json:"name"`
	RefNum    int    `json:"ref_num"`
}

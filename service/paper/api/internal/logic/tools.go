package logic

import "soft2_backend/service/paper/api/internal/types"

func NilHandler(value interface{}, expectedType string) interface{} {
	if value == nil {
		switch expectedType {
		case "string":
			return ""
		case "int":
			return 0
		case "list":
			return []interface{}{}
		}
	} else {
		switch expectedType {
		case "string":
			return value.(string)
		case "int":
			return int(value.(float64))
		case "list":
			return value.([]interface{})
		}
	}
	return nil
}

func GetSize(s, max, min int) int {
	if max == min {
		return 20
	}
	return int((float64(s-min)/float64(max-min) + 1) * 20)
}

func GetD(d, max, min int) int {
	if max == min {
		return 0
	}
	return int((float64(d-min) / float64(max-min)) * 10)
}

func UpdateMaxMin(max *int, min *int, d int) {
	if d > *max {
		*max = d
	}
	if d < *min {
		*min = d
	}
}

func GetColor(d int) string {
	switch d {
	case 0:
		return "#ccffff"
	case 1:
		return "#b8f5fa"
	case 2:
		return "#7ad1db"
	case 3:
		return "#a3ebf5"
	case 4:
		return "#8fe0f0"
	case 5:
		return "#7ad6eb"
	case 6:
		return "#5cc7e3"
	case 7:
		return "#33b2d9"
	case 8:
		return "#0099cc"
	default:
		return "#ffffff"
	}
}

func GetPaperAuthors(paper map[string]interface{}) []types.AuthorJSON {
	authors := make([]types.AuthorJSON, 0)
	for _, author := range paper["authors"].([]interface{}) {
		hasId := false
		if author.(map[string]interface{})["id"] != nil {
			hasId = true
		}
		authors = append(authors, types.AuthorJSON{
			Name:  NilHandler(author.(map[string]interface{})["name"], "string").(string),
			Id:    NilHandler(author.(map[string]interface{})["id"], "string").(string),
			HasId: hasId,
		})
	}
	return authors
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

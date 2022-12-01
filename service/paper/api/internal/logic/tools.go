package logic

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

func GetColor(d int) string {
	switch d {
	case 0:
		return "#1C1C1C"
	case 1:
		return "#363636"
	case 2:
		return "#4F4F4F"
	case 3:
		return "#696969"
	case 4:
		return "#828282"
	case 5:
		return "#9C9C9C"
	case 6:
		return "#B5B5B5"
	case 7:
		return "#CFCFCF"
	case 8:
		return "#E8E8E8"
	default:
		return "#FFFFFF"
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

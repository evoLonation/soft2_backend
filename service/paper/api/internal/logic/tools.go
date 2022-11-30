package logic

func NilHandler(value interface{}, expectedType string) interface{} {
	if value == nil {
		switch expectedType {
		case "string":
			return ""
		case "int":
			return 0
		case "list":
			return []string{}
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
	return int((float64(s-min)/float64(max-min) + 1) * 20)
}

func GetD(d, max, min int) int {
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

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
			if int(value.(float64)) == -1 {
				return 0
			}
			return int(value.(float64))
		case "list":
			return value.([]interface{})
		}
	}
	return nil
}

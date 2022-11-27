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

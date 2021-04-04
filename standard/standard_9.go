package standard

import "strings"

// 常量
const (
	CaseLower = 1
	CaseUpper = 2
)

// ArrayChangeKeyCase .
func ArrayChangeKeyCase(array map[string]interface{}, scase int) map[string]interface{} {
	if len(array) == 0 {
		return array
	}
	var _array = make(map[string]interface{}, len(array))
	if scase == CaseLower {
		for index, val := range array {
			_array[strings.ToLower(index)] = val
		}
	} else if scase == CaseUpper {
		for index, val := range array {
			_array[strings.ToUpper(index)] = val
		}
	}
	return _array
}

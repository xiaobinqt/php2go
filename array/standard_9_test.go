package array

import "testing"

func TestArrayChangeKeyCase(t *testing.T) {
	array := map[string]interface{}{
		"1":   233,
		"abc": 233,
		"aBC": 233,
		"ABC": 233,
		"1Xb": 233,
	}
	t.Logf("原始值 %+v \n", array)
	t.Logf("全大写 %+v \n", ArrayChangeKeyCase(array, CaseUpper))
	t.Logf("全小写 %+v \n", ArrayChangeKeyCase(array, CaseLower))
}

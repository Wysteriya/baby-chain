package tools

import (
	"testing"
)

func TestData_Validate(t *testing.T) {
	var data1 Data

	data1 = Data{"test1": "test", "test2": Data{"test1": "1"}, "test3": "true"}
	TError(data1.Validate(), t)

	data1 = Data{"test1": "test", "test2": Data{"test1": 1}, "test3": true}
	TExpectedError(data1.Validate(), t)
}

func TestJson(t *testing.T) {
	data1 := Data{"test1": "test", "test2": Data{"test1": "0"}, "test3": "true"}
	TTestJson(data1, t)
	t.Log(data1.String())
}

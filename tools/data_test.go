package tools

import (
	"testing"
)

func GoodData() Data {
	return Data{
		"test1": "test",
		"test2": Data{"test1": "1"},
		"test3": Data{"test1": "1", "test2": Data{"test1": "true"}},
	}
}

func BadData() Data {
	return Data{
		"test1": "test",
		"test2": Data{"test1": 1},
		"test3": Data{"test1": 1, "test2": Data{"test1": true}},
	}
}

func TestData_Validate(t *testing.T) {
	var data Data

	data = GoodData()
	TError(data.Validate(), t)

	data = BadData()
	TExpectedError(data.Validate(), t)
}

func TestJson(t *testing.T) {
	data := GoodData()
	TTestJson(data, t)
	t.Log(data.String())
}

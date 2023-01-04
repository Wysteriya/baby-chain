package tools

import (
	"testing"
)

func TestData_Validate(t *testing.T) {
	var data Data

	data = GoodTestData()
	TError(data.Validate(), t)

	data = BadTestData()
	TExpectedError(data.Validate(), t)
}

func TestJson(t *testing.T) {
	data := GoodTestData()
	TTestJson(data, t)
	t.Log(data.String())
}

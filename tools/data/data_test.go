package data

import (
	"baby-chain/tools"
	"testing"
)

func TestData_Validate(t *testing.T) {
	var data Data

	data = GoodTestData()
	tools.TError(data.Validate(), t)

	data = BadTestData()
	tools.TExpectedError(data.Validate(), t)
}

func TestJson(t *testing.T) {
	data := GoodTestData()
	tools.TTestJson(data, t)
	t.Log(data.String())
}

package data

import (
	errors2 "baby-chain/errors"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// Data : all keys will be lower-cased, values can be any unicode
type Data map[string]interface{}
type Array []interface{}

func (d *Data) MarshalJSON() ([]byte, error) {
	type alias Data
	if err := d.Validate(); err != nil {
		return nil, err
	}
	return json.Marshal((*alias)(d))
}

func (d *Data) UnmarshalJSON(save []byte) error {
	type alias Data
	aux := alias{}
	if err := json.Unmarshal(save, &aux); err != nil {
		return err
	}

	*d = Data(aux)
	d.Dataify()
	if err := d.Validate(); err != nil {
		return err
	}

	return nil
}

func dataifyInterface(val interface{}) interface{} {
	switch dVal := val.(type) {
	case map[string]interface{}:
		var eval Data = dVal
		eval.Dataify()
		return eval
	case Data:
		dVal.Dataify()
		return dVal
	case []interface{}:
		var eVal = make(Array, len(dVal))
		for i, data := range dVal {
			eVal[i] = dataifyInterface(data)
		}
		return eVal
	case Array:
		for i, data := range dVal {
			dVal[i] = dataifyInterface(data)
		}
		return dVal
	}
	return val
}

func (d *Data) Dataify() {
	for key, val := range *d {
		(*d)[key] = dataifyInterface(val)
	}
}

func interfaceValidate(val interface{}, key string, errs *[]error) {
	switch dVal := val.(type) {
	case string:
	case Data:
		if err := dVal.Validate(); err != nil {
			*errs = append(*errs, errors.Unwrap(err))
		}
	case Array:
		for i, data := range dVal {
			interfaceValidate(data, string(rune(i)), errs)
		}
	default:
		*errs = append(*errs, fmt.Errorf("%s : %s", key, reflect.TypeOf(dVal)))
	}
}

func (d *Data) Validate() error {
	var errs []error
	for key, val := range *d {
		interfaceValidate(val, key, &errs)
	}
	return errors2.MultiError(errs, "panicDataType")
}

func (d *Data) String() string {
	s, err := d.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return string(s)
}

func GoodTestData() Data {
	return Data{
		"test1": "test",
		"test2": Data{"test1": "1"},
		"test3": Data{"test1": "1", "test2": Data{"test1": "true"}},
		"test4": Array{"1", "2", "3"},
		"test5": Array{
			Data{"test1": "test"},
			Data{"test1": "1"},
			Data{"test2": "false"},
		},
		"test6": Array{
			"test",
			"1",
			Array{"false", Data{"test1": "test"}},
			Data{"test1": "1"},
		},
	}
}

func BadTestData() Data {
	return Data{
		"test1": "test",
		"test2": Data{"test1": 1},
		"test3": Data{"test1": 1, "test2": Data{"test1": true}},
		"test4": Array{"1", "2", "3"},
		"test5": Array{
			Data{"test1": "test"},
			Data{"test1": 1},
			Data{"test2": false},
		},
		"test6": Array{
			"test",
			1,
			Array{false, Data{"test1": "test"}},
			Data{"test1": 1},
		},
	}
}

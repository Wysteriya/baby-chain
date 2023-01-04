package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type Data map[string]interface{}

func (d *Data) MarshalJSON() ([]byte, error) {
	type alias Data
	if err := d.Validate(); err != nil {
		return nil, err
	}
	return json.Marshal((*alias)(d))
}

func (d *Data) Dataify() {
	for key, val := range *d {
		if dVal, ok := val.(map[string]interface{}); ok {
			eVal := Data(dVal)
			eVal.Dataify()
			(*d)[key] = eVal
		} else if dVal, ok := val.(Data); ok {
			dVal.Dataify()
			(*d)[key] = dVal
		}
	}
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

func (d *Data) Validate() error {
	var errs []error
	for key, val := range *d {
		switch dVal := val.(type) {
		case string:
		case Data:
			if err := dVal.Validate(); err != nil {
				errs = append(errs, fmt.Errorf("%s : %w", key, errors.Unwrap(err)))
			}
		default:
			errs = append(errs, fmt.Errorf("%s : %s", key, reflect.TypeOf(dVal)))
		}
	}
	return MultiError(errs, "panicDataType")
}

func (d *Data) String() string {
	s, err := d.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return string(s)
}

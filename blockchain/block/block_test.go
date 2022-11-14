package block

import (
    "testing"
    "reflect"
)

func TestSaveLoad(t *testing.T) {
    b := Genesis(Data{"balances": map[string]interface {}{"amith": "1000", "yash": "500"}, "key": "10", "test": true})
    save, err := b.Save()
    if err != nil {
        t.Errorf("%s", err)
    }
    _b, err := Load(save)
    if err != nil {
        t.Errorf("%s", err)
    }
    if !reflect.DeepEqual(b, _b) {
        t.Fatalf("Saved and Loaded Data are not equal\n%#v\n%#v", b, _b)
    }
}

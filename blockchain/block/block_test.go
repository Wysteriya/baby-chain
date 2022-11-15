package block

import (
    "testing"
    "reflect"
)

func TestSaveLoad(t *testing.T) {
    b := Genesis(Data{"balances": Data{"amith": "1000", "yash": "500"}, "key": "10", "test": true})
    if save, err := b.Save(); err != nil {
        t.Errorf("%s", err)
    } else if _b, err := Load(save); err != nil {
        t.Errorf("%s", err)
    } else if !reflect.DeepEqual(b, _b) {
        t.Fatalf("Saved and Loaded Data are not equal\n%#v\n%#v", b, _b)
    }
}

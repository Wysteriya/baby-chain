package block

import (
    "fmt"
    "errors"
)

type Data map[string]interface {}

func (d *Data) String() string {
    dat := ""
    for _, value := range *d {dat += fmt.Sprint(value, "\n")}
    return dat
}

func (d *Data) Validate() error {
    for key, value := range *d {
        switch vt := value.(type) {
            case bool:
            case string:
            case []string:
            case Data:
                if value, ok := value.(Data); ok {
                    return value.Validate()
                }
            case []Data:
                for _, v := range value.([]Data) {
                    if err := v.Validate(); err != nil {
                        return err
                    }
                    return nil
                }
            case map[string]interface {}:
                if value, ok := value.(map[string]interface {}); ok {
                    (*d)[key] = Data(value)
                }
            case []map[string]interface {}:
                if value, ok := value.([]map[string]interface {}); ok {
                    for k, v := range value {
                        value[k] = Data(v)
                    }
                    (*d)[key] = value
                }
            default: return errors.New(fmt.Sprint(vt, " can't be treated as json object"))
        }
    }
    return nil
}

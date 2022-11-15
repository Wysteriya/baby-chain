package block

import (
    "time"
)

type Time int64

func (t Time) String() string {
    return time.Unix(int64(t), 0).String()
}

func CurrTime() Time {
    return Time(time.Now().Unix())
}

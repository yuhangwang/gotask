package encode

import (
    "encoding/json"
    "github.com/yuhangwang/gotask/err"
)

func Json(data interface{}) []byte {
    data, err_msg := json.Marshal(data)
    err.Check(err_msg)
    return data.([]byte)
}

package encode

import (
    "gopkg.in/yaml.v2"
    "github.com/yuhangwang/gotask/err"
)


func Yaml(data interface{}) []byte {
    output, err_msg := yaml.Marshal(data)
    err.Check(err_msg)
    return output
}

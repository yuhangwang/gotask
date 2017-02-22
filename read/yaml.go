package read

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "github.com/yuhangwang/gotask/err"
)



func Yaml(file_name string) interface{} {
    text_bytes, err_msg := ioutil.ReadFile(file_name)
    err.Check(err_msg)
    var data interface{}
    err.Check(yaml.Unmarshal(text_bytes, &data))
    return data
}

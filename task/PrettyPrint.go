package task

import(
    "encoding/json"
    "fmt"
    "github.com/yuhangwang/gotask/convert"
)

func PrettyPrint(v interface{}) {
    x, _ := json.MarshalIndent(convert.MapLike(v), "", " ")
    fmt.Println(string(x))
    fmt.Println("------------------------------------------")
}
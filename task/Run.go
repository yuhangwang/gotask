package task

import (
    "fmt"
    "os/exec"
    "github.com/yuhangwang/gotask/err"
)


func Run(cmd string, args ...string) {
    stdOutErr, err_msg := exec.Command(cmd,  args...).CombinedOutput()
    fmt.Printf("%s\n", stdOutErr)
    err.Check(err_msg)
}

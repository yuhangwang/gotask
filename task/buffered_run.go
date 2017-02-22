package task

import (
    "bufio"
    "log"
    "os/exec"
    "github.com/yuhangwang/gotask/err"
)


func BufRun(cmd_str string, args ...string) {
    cmd := exec.Command(cmd_str,  args...)
    stdout, err_msg := cmd.StdoutPipe()
    err.Check(err_msg)
    err.Check(cmd.Start())
    terminal := bufio.NewScanner(stdout)
    for terminal.Scan() {
        log.Printf(terminal.Text())
    }
    err.Check(terminal.Err())
}

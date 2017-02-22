package main

import (
    "os"
    "os/exec"
    "fmt"
    "sync"
    "strings"
    "github.com/urfave/cli"
    "github.com/yuhangwang/gotask/read"
    "github.com/yuhangwang/gotask/encode"
    "github.com/yuhangwang/gotask/err"
)

func run(cmd string, args ...string) {
    stdOutErr, err_msg := exec.Command(cmd, args...).CombinedOutput()
    fmt.Printf("%s\n", stdOutErr)
    err.Check(err_msg)
}


func main() {
    app := cli.NewApp()
    app.Name = "GoTask"
    app.Version = "0.1.0"
    app.Usage = "Command line parallel task manager"
    app.Action = func(context *cli.Context) error {
        data := read.Yaml(context.Args().First())
        cmd := context.Args().Get(1)
        primary_arg := strings.Join(cli.Args(context.Args().Tail()).Tail(), " ")
        var wg sync.WaitGroup
        for _, v := range data {
            wg.Add(1)
            go func(arg string) {
                defer wg.Done()
                run(cmd, primary_arg, arg)
            }(string(encode.Json(v)))
        }
        wg.Wait()
        return nil
    }

    app.Run(os.Args)
}
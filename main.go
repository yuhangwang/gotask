package main

import (
    "os"
    "os/exec"
    "fmt"
    "sync"
    "github.com/urfave/cli"
    "github.com/yuhangwang/gotask/read"
    "github.com/yuhangwang/gotask/encode"
    "github.com/yuhangwang/gotask/err"
    "github.com/yuhangwang/gotask/wrap"
)

func run(cmd string, args ...string) {
    stdOutErr, err_msg := exec.Command(cmd,  args...).CombinedOutput()
    fmt.Printf("%s\n", stdOutErr)
    err.Check(err_msg)
}


func main() {
    app := cli.NewApp()
    app.Name = "GoTask"
    app.Version = "0.1.0"
    app.Usage = "Command line parallel task manager"
    app.Action = func(context *cli.Context) error {
        argc := context.NArg()
        file_tasks := context.Args().Get(argc-1)
        data := read.Yaml(file_tasks)
        cmd := context.Args().First()
        primary_arg := context.Args()[1:argc-1]
        var wg sync.WaitGroup
        for _, v := range data {
            wg.Add(1)
            go func(arg string) {
                defer wg.Done()
                run(cmd, append(primary_arg, arg)...)
            }(wrap.Brackets("[",encode.Json(v),"]"))
        }
        wg.Wait()
        return nil
    }

    app.Run(os.Args)
}
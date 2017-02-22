package main

import (
    "os"
    "sync"
    "github.com/urfave/cli"
    "github.com/yuhangwang/gotask/read"
    "github.com/yuhangwang/gotask/encode"
    "github.com/yuhangwang/gotask/wrap"
    "github.com/yuhangwang/gotask/task"
    "github.com/yuhangwang/gotask/convert"
)


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
        for _, v := range data.([]interface{}) {
            wg.Add(1)
            go func(arg string) {
                defer wg.Done()
                task.Run(cmd, append(primary_arg, arg)...)
            }(wrap.Brackets("[",encode.Json(convert.MapLike(v)),"]"))
        }
        wg.Wait()
        return nil
    }

    app.Run(os.Args)
}
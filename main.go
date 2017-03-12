package main

import (
    "os"
    "sync"
    "regexp"
    "fmt"
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
        if len(context.Args()) == 0 {
            fmt.Println("Error hint: please provide arguments to gotask")
            os.Exit(1)
        }
        validYamlFileName := regexp.MustCompile(".*[.]yaml$")
        cmd := context.Args().First()
        argc := context.NArg()
        file_tasks := context.Args().Get(argc-1)
        if !validYamlFileName.MatchString(file_tasks) {
            fmt.Println("Error hint: please provide a YAML file")
            task.Run(cmd, context.Args()[1:]...)
            os.Exit(1)
        }

        data := read.Yaml(file_tasks)
        primary_arg := context.Args()[1:argc-1]
        
        var wg sync.WaitGroup
        for _, v := range data.([]interface{}) {
            wg.Add(1)
            go func(arg string) {
                defer wg.Done()
                tmp := make([]string, len(primary_arg))
                copy(tmp, primary_arg)
                task.Run(cmd, append(tmp, arg)...)
            }(wrap.Brackets("[",encode.Json(convert.MapLike(v)),"]"))
        }
        wg.Wait()
        return nil
    }

    app.Run(os.Args)
}
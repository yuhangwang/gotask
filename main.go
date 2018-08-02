package main

import (
    "strconv"
    "runtime"
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
    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "parallel",
            Value: "default",
            Usage: "number of parallel tasks to be executed at the same time",
            EnvVar: "GTPAR",
        },
    }

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
        
        fmt.Println("num of parallel", context.String("parallel"))
        max_parallel := runtime.NumCPU()
        if context.String("parallel") != "default" {
            max_parallel, _ = strconv.Atoi(context.String("parallel"))
        }
        var wg sync.WaitGroup
        // note: in go routine, don't include any mutable object
        // in any port of the go routine function body.
        // Use function arguments and make sure all go routines
        // are independent.
        // note: variables declared inside the loop
        //    are not shared between loops
        // I will rewrite this part using recursion instead of loop
        // Mutations in loops are VERY BAD!!!
        // Why can't Go be more functional?! -- Steven W.  03/15/17
        ccc := 0
        for _, v := range data.([]interface{}) {
            task.PrettyPrint(v)
            json_arg := wrap.Brackets("[",encode.Json(convert.MapLike(v)),"]")
            all_args := task.AppendArg(primary_arg, json_arg)
            wg.Add(1)
            ccc = ccc + 1
            go func(cmd string, all_args []string) {
                defer wg.Done()
                task.Run(cmd, all_args...)
            }(cmd, all_args)

            if ccc == max_parallel {
                wg.Wait()
                ccc = 0
            }
        }
        wg.Wait(); // wait for the residual go routines to finish

        fmt.Printf("=============== max parallel threads: %d\n", max_parallel)
        return nil
    }

    app.Run(os.Args)
}
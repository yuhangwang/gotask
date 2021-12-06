# GoTask
Parallel task automation using Go

## Installation
First, make sure you have access to GitHub. Then follow the following 
steps to install GoTask

1. install Google's Go https://go.dev/dl
2. `git clone https://github.com/yuhangwang/gotask`
3. `cd gotask`
4. `go install`

## Usage
The `gotask` command takes at least two arguments, i.e.,
the command-line command to be executed and a yaml file 
containing a list of command-line arguments to be given 
to the command-line command. `gotask` will execute the 
command in parallel and give each command one of the command-line
arguments in `args.yml`. The default number of executing threads
equals to the number of virtual CPU cores. It can be changed by 
setting the environmental variable `GTPAR`.
```bash
gotask [command] [args.yml]
```

## Example
```bash
gotask sleep args.yml
```
The content of the `args.yml` file is listed below.
```yaml
- 3
- 3
- 3
```
The effect is that `gotask` will execute `sleep 3` command three times in parallel. 


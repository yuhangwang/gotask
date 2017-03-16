package task

// Append a new arg to old arg array
func AppendArg(base_args []string, new_arg string) []string {
    tmp := make([]string, len(base_args))
    copy(tmp, base_args)
    return append(tmp, new_arg)
}

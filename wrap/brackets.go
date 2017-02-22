package wrap

import "fmt"

func Brackets(left string, s []byte, right string) string {
    return fmt.Sprintf("%s%s%s", left, string(s), right)
}
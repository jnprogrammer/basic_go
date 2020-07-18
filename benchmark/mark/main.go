package benchmark

import "fmt"

func Mark(s string) string {
	return fmt.Sprint("Target v% has been marked", s)
}

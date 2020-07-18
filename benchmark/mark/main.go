package benchmark

import "fmt"

func Mark(s string) string {
	return fmt.Sprintf("Target %v has been marked", s)
}

package saying

import "fmt"

func Greet(s string) string {
	return fmt.Sprint("Welcome,", s)
}
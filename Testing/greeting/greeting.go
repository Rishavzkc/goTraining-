package greeting

import "fmt"

func Hello(name string) string {
	if len(name) == 0 {
		return "Hello Dude!"
	}
	return fmt.Sprintf("Hello %v!", name)
}

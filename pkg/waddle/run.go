package waddle

import "fmt"

func Run(cwd string, args ...string) int {
	fmt.Printf(cwd)
	fmt.Printf(args[0])
	return 0
}

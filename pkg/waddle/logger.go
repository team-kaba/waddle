package waddle

import "fmt"

type logger struct {
	enabled bool
}

func (l *logger) print(a ...interface{}) {
	fmt.Print(a...)
}

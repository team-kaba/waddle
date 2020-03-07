package waddle

func Run(cwd string, args ...string) int {
	l := &logger{enabled: true}

	selectActions(cwd, l, args...)
	return 0
}

func selectActions(cwd string, l *logger, args ...string) error {
	l.print(cwd)

	return nil
}

package internal

type cliOptions struct {
	Name string
}

func NewCLIOptions(name string) cliOptions {
	return cliOptions{
		Name: name,
	}
}

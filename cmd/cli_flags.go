package cli_flags

//files
type files []string

func (f *files) String() string {
	return ""
}

func (f *files) Set(flag string) error {
	*f = append(*f, flag)
	return nil
}

var FileNamesFromCLI files

//numbers
type numbers []string

func (n *numbers) String() string {
	return ""
}

func (n *numbers) Set(flag string) error {
	*n = append(*n, flag)
	return nil
}

var NumbersFromCLI numbers

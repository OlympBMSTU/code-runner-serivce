package runner

// interface to run code
type IRunner interface {
	RunCode(codeFileName, answersFileName string) int
}

type CodeRunner struct {
	// channels
	// command
}

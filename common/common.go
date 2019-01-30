package common

import "errors"

type CompilerOptions struct {
	InFileName   string
	OutFileName  string
	Comiler      string
	CompileFlags []string
	Language     string
}

func NewCompilerOptions(language string) (CompilerOptions, error) {
	var options CompilerOptions
	switch language {
	case PASCAL:
		options.Comiler = FPCCOMPIL //compilers[]
	case CPP:
		options.Comiler = GPPCOMPIL
	case CLANG:
		options.Comiler = GCCCOMPIL
	case PY27:
		options.Comiler = PY27INTERP
	case PY3:
		options.Comiler = PY3INTERP
	default:
		return CompilerOptions{}, errors.New("Incorrect language type")
	}
	return options, nil
}

func (options *CompilerOptions) GetCompiler() string {
	return options.Comiler
}

func (options *CompilerOptions) AsArgs() []interface{} {
	var params []interface{}
	return params
}

type RunnerOptions struct {
	FileName      string
	Interpretator string
	InterpFlags   string
}

// func NewRunnerOptinos(fileName, language) (RunnerOptions, error) {

// }

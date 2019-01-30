package runner

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/OlympBMSTU/code-runner-serivce/common"
)

// interface to run code
type IRunner interface {
	RunCode(codeFileName, answersFileName string) int
}

type CodeRunner struct {
	// channels
	// command
}

func worker(data chan string) {

}

func (runner *CodeRunner) RunCode(codeFileName, answersFileName string) int {
	// get answers from file
	// os.ReadAll
	// if err := os.Open(answersFileName); err != nil {
	// return bad status => please contact to admin
	// }
	return 0
}

func CompileProgram(fName, lang string) (string, error) {
	compilerOptions, err := common.NewCompilerOptions(lang)

	if err != nil {
		return "", err
	}
	// 	// example
	// 	newFileName := fName - ".lang" + ".out"
	cmd := exec.Command(compilerOptions.GetCompiler(), compilerOptions.AsArgs()...) //fName, compileOptions.AsArgs())

	stdout, err := cmd.StderrPipe()

	if err != nil {
		fmt.Print(err)
		return "", err
	}

	err = cmd.Start()
	if err != nil {

	}
	b, err := ioutil.ReadAll(stdout)

	if err = cmd.Wait(); err != nil {
		// 		// todo check error, by default think that its user problem
		fmt.Print(err)
	}

	return "", compileOptions.GetOutFile()
}

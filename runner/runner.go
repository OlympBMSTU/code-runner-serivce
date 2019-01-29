package runner

import (
	"fmt"
	"io/ioutil"
	"os/exec"
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

func CompileProgram(fName, lang string) (string, error) {
	compileOptions := GetCompileOptions(lang)
	// example
	newFileName := fName - ".lang" + ".out"
	cmd := exec.Command(compleOptinos.Compiler, fName, compileOptions.AsArgs())

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
		// todo check error, by default think that its user problem
		fmt.Print(err)
	}

	return "", newFileName
}

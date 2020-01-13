package build

import (
	"fmt"
	"github.com/arata-nvm/Solitude/compiler"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func EmitLLVM(filename, outputPath string, isDebug, optimize bool) error {
	fmt.Printf("Compiling %s\n", filename)
	c := compiler.New(isDebug)
	errors := c.Compile(filename)
	printErrors(errors)
	if optimize {
		fmt.Println("Optimizing")
		c.Optimize()
	}
	compiled := c.GenIR()

	if outputPath == "" {
		outputPath = getFileNameWithoutExt(filename) + ".ll"
	}

	err := ioutil.WriteFile(outputPath, []byte(compiled), 0666)
	if err != nil {
		return err
	}

	fmt.Println("Finished")
	return nil
}

func Build(filename, outputPath string, isDebug, optimize bool) error {
	fmt.Printf("Compiling %s\n", filename)
	c := compiler.New(isDebug)
	errors := c.Compile(filename)
	printErrors(errors)
	if optimize {
		fmt.Println("Optimizing")
		c.Optimize()
	}
	compiled := c.GenIR()

	tmpDir, err := ioutil.TempDir("", "solitude")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(tmpDir+"/main.ll", []byte(compiled), 0666)
	if err != nil {
		return err
	}

	if outputPath == "" {
		outputPath = getFileNameWithoutExt(filename)
	}

	clangArgs := []string{
		"-Wno-override-module",
		tmpDir + "/main.ll",
		"-o", outputPath,
	}

	if optimize {
		clangArgs = append(clangArgs, "-O3")
	}

	cmd := exec.Command("clang", clangArgs...)
	err = cmd.Run()
	if err != nil {
		return err
	}

	os.RemoveAll(tmpDir)

	fmt.Println("Finished")
	return nil
}

func printErrors(errors []string) {
	if len(errors) != 0 {
		for _, e := range errors {
			_, _ = fmt.Fprintln(os.Stderr, e)
		}
		os.Exit(1)
	}
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

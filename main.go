package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/frsfahd/go-wc/util"
)

func main() {
	var (
		result   string
		filedata util.File
		txtFile  *os.File
		err      error
	)

	// validate number of arguments
	if flag.NArg() == 0 {
		if os.Stdin != nil {
			txtFile, filedata, err = util.GetData_stdin()
			if err != nil {
				panic(err)
			}
			result, err = util.ProcessData_stdin(txtFile, filedata)
			if err != nil {
				panic(err)
			}
		} else {
			panic(errors.New("a filepath argument is required"))
		}
	} else {
		filedata, err = util.GetData_file()
		if err != nil {
			panic(err)
		}
		result = util.ProcessData_file(filedata)

	}

	fmt.Println(result)

	// debug
	fmt.Println(filedata.Filepath)
	fmt.Println(filedata.Option)
	// fmt.Printf("num of flag: %d\n", flag.NFlag())
	// fmt.Printf("num of args: %d\n", flag.NArg())
	// fmt.Printf("all args (%d): %v\n", len(os.Args), os.Args)
}

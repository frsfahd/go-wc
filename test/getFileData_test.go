package test

import (
	"flag"
	"os"
	"reflect"
	"testing"

	"github.com/frsfahd/go-wc/util"
)

func TestGetFileData(t *testing.T) {
	tests := []struct {
		desc        string
		expected    util.File
		expectedErr bool
		osArgs      []string
	}{
		// Test 1
		{
			desc: "default parameters",
			expected: util.File{
				Filepath: "test.txt",
				Option: map[string]bool{
					"bytes": true,
					"lines": true,
					"words": true,
					"chars": true,
				},
			},
			expectedErr: false,
			osArgs:      []string{"cmd", "test.txt"},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {

			// clean flags before each test
			flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)

			actualOSArgs := os.Args

			// reset to original os.Args after each test
			defer func() {
				os.Args = actualOSArgs
			}()

			// set up variables
			os.Args = test.osArgs
			got, _, err := util.GetFileData()

			gotErr := (err != nil)

			// start of test
			if gotErr != test.expectedErr {
				t.Errorf("GetFileData() error = %v, want = %v", gotErr, test.expectedErr)
			}

			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("GetFileData() = %v, want = %v", got, test.expected)
			}

		})
	}

}

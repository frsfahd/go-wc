package test

import (
	"testing"

	"github.com/frsfahd/go-wc/util"
)

func TestProcessData(t *testing.T) {
	tests := []struct {
		desc     string
		expected string
		filedata util.File
	}{
		{
			desc:     "default parameters",
			expected: "7145 342190 58164  test.txt",
			filedata: util.NewFile("../text.txt"),
		},
		{
			desc:     "display only character count",
			expected: "339292  test.txt",
			filedata: util.File{Filepath: "../test.txt", Option: map[string]bool{"chars": true, "lines": false, "words": false, "bytes": false}},
		},
	}

	// start of test
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			// got := util.ProcessData(test.filedata)

			// if got != test.expected {
			// 	t.Errorf("ProcessData() = %s, want = %s", got, test.expected)
			// }
		})
	}
}

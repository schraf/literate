package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateCode(t *testing.T) {
	dir, err := os.MkdirTemp("", "literate")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	md1 := `
# Title

## Body

` + "```" + `go {name="body"}
fmt.Println("Hello, World!")
` + "```" + `

## Main

` + "```" + `go {name="block1" filename="file1.go"}
package main

import "fmt"

func main() {
{{ include "body" }}
}
` + "```" + `
`
	md2 := `
# Title 2

` + "```" + `go {name="block2" filename="file2.go"}
package main

import "fmt"

func main() {
{{ include "body" }}
}
` + "```" + `
`
	err = os.WriteFile(filepath.Join(dir, "test1.md"), []byte(md1), 0644)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(dir, "test2.md"), []byte(md2), 0644)
	require.NoError(t, err)

	inputs := []string{
		filepath.Join(dir, "test1.md"),
		filepath.Join(dir, "test2.md"),
	}

	err = GenerateCode(inputs, dir)
	require.NoError(t, err)

	_, err = os.Stat(filepath.Join(dir, "file1.go"))
	assert.NoError(t, err)

	_, err = os.Stat(filepath.Join(dir, "file2.go"))
	assert.NoError(t, err)
}

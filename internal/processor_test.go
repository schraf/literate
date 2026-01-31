package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessor_GenerateCodeFiles(t *testing.T) {
	storage := NewCodeBlockStorage()
	block1 := &CodeBlock{
		Attributes: CodeBlockAttributes{
			Name:     "block1",
			Filename: "file1.go",
		},
		Body: "package main",
	}
	block2 := &CodeBlock{
		Attributes: CodeBlockAttributes{
			Name: "block2",
		},
		Body: `import "fmt"`,
	}
	block3 := &CodeBlock{
		Attributes: CodeBlockAttributes{
			Name:     "block3",
			Filename: "file2.go",
		},
		Body: `{{include "block1"}}
{{include "block2"}}

func main() {
	fmt.Println("Hello")
}
`,
	}
	require.NoError(t, storage.AddCodeBlock(block1))
	require.NoError(t, storage.AddCodeBlock(block2))
	require.NoError(t, storage.AddCodeBlock(block3))

	dir, err := os.MkdirTemp("", "literate")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	processor := NewProcessor(dir)
	err = processor.GenerateCodeFiles(storage)
	require.NoError(t, err)

	content, err := os.ReadFile(filepath.Join(dir, "file1.go"))
	require.NoError(t, err)
	assert.Equal(t, "package main", string(content))

	content, err = os.ReadFile(filepath.Join(dir, "file2.go"))
	require.NoError(t, err)
	expected := `package main
import "fmt"

func main() {
	fmt.Println("Hello")
}`
	assert.Equal(t, expected, string(content))
}

package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseMarkdownFile(t *testing.T) {
	dir, err := os.MkdirTemp("", "literate")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	md := `
# Title

` + "```" + `go {name="block1"}
code block 1
` + "```" + `

` + "```" + `go {name="block2"}
code block 2
` + "```" + `
`
	err = os.WriteFile(filepath.Join(dir, "test.md"), []byte(md), 0644)
	require.NoError(t, err)

	blocks := make(map[string]string)
	for block, err := range ParseMarkdownFile(filepath.Join(dir, "test.md")) {
		require.NoError(t, err)
		blocks[block.Attributes.Name] = block.Body
	}

	assert.Equal(t, "code block 1\n", blocks["block1"])
	assert.Equal(t, "code block 2\n", blocks["block2"])
}

package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCodeBlockAttributes_Parse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected CodeBlockAttributes
	}{
		{
			name:  "name and filename",
			input: `name="test" filename="test.go"`,
			expected: CodeBlockAttributes{
				Name:     "test",
				Filename: "test.go",
			},
		},
		{
			name:  "only name",
			input: `name="test"`,
			expected: CodeBlockAttributes{
				Name: "test",
			},
		},
		{
			name:  "only filename",
			input: `filename="test.go"`,
			expected: CodeBlockAttributes{
				Filename: "test.go",
			},
		},
		{
			name:     "no attributes",
			input:    "",
			expected: CodeBlockAttributes{},
		},
		{
			name:     "unknown attributes",
			input:    `foo="bar"`,
			expected: CodeBlockAttributes{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var attrs CodeBlockAttributes
			attrs.Parse(tt.input)
			assert.Equal(t, tt.expected, attrs)
		})
	}
}

func TestCodeBlockStorage_AddCodeBlock(t *testing.T) {
	storage := NewCodeBlockStorage()
	block1 := &CodeBlock{
		Attributes: CodeBlockAttributes{
			Name:     "block1",
			Filename: "file1.go",
		},
	}
	err := storage.AddCodeBlock(block1)
	require.NoError(t, err)

	// test duplicate name
	block2 := &CodeBlock{
		Attributes: CodeBlockAttributes{
			Name: "block1",
		},
	}
	err = storage.AddCodeBlock(block2)
	assert.Error(t, err)

	// test duplicate filename
	block3 := &CodeBlock{
		Attributes: CodeBlockAttributes{
			Name:     "block3",
			Filename: "file1.go",
		},
	}
	err = storage.AddCodeBlock(block3)
	assert.Error(t, err)
}

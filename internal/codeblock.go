package internal

import (
	"fmt"
	"regexp"
)

//--=====================================================================--
//--== CODE BLOCK ATTRIBUTES
//--=====================================================================--

var AttributePattern = regexp.MustCompile(`(\w+)="([^"]*)"`)

type CodeBlockAttributes struct {
	Name     string
	Filename string
}

func (c *CodeBlockAttributes) Parse(input string) {
	var attributes CodeBlockAttributes

	matches := AttributePattern.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		switch match[1] {
		case "name":
			attributes.Name = match[2]
		case "filename":
			attributes.Filename = match[2]
		}
	}

	*c = attributes
}

//--=====================================================================--
//--== CODE BLOCK
//--=====================================================================--

type CodeBlock struct {
	Attributes CodeBlockAttributes
	Body       string
}

//--=====================================================================--
//--== CODE BLOCK STORAGE
//--=====================================================================--

type CodeBlockStorage struct {
	Blocks map[string]*CodeBlock
	Files  map[string]*CodeBlock
}

func NewCodeBlockStorage() *CodeBlockStorage {
	return &CodeBlockStorage{
		Blocks: make(map[string]*CodeBlock),
		Files:  make(map[string]*CodeBlock),
	}
}

func (s *CodeBlockStorage) AddCodeBlock(block *CodeBlock) error {
	if _, exists := s.Blocks[block.Attributes.Name]; exists {
		return fmt.Errorf("duplicate code block name found '%s'", block.Attributes.Name)
	}

	if block.Attributes.Filename != "" {
		if _, exists := s.Files[block.Attributes.Filename]; exists {
			return fmt.Errorf("duplicate output file found '%s'", block.Attributes.Filename)
		}

		s.Files[block.Attributes.Filename] = block
	}

	s.Blocks[block.Attributes.Name] = block

	return nil
}

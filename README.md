# Literate

A modern tool for **Literate Programming** that extracts and composes source code from Markdown files.

## What is Literate Programming?

Literate programming (LP) is a programming paradigm introduced in 1984 by **Donald Knuth**. In this approach, a computer program is treated as a piece of literature: an explanation of the program logic in a natural language (like English), interspersed with snippets of macros and traditional source code.

> "Literate programming is writing out the program logic in a human language with included code snippets and macros."

### Key Concepts

*   **Logic over Compiler:** Instead of writing code in the order imposed by the compiler, the programmer develops the program in the order demanded by the logic and flow of their thoughts.
*   **Web of Thoughts:** Programs are treated as an interconnected "web" of concepts. Natural language macros describe abstractions, hiding lower-level implementation details.
*   **Two Representations:**
    *   **Weaving:** Generating formatted documentation from the source (for humans).
    *   **Tangling:** Generating compilable machine code from the source (for computers).

### Advantages

*   **Higher Quality:** Forces explicit statement of thoughts, revealing design flaws early.
*   **Better Documentation:** Documentation is not an add-on but grows naturally with the code.
*   **Context:** Provides a "bird's eye view" of the code, aiding memory and processing of complex concepts.

## About This Project

This application is a streamlined tool designed to bring Literate Programming to modern workflows. It reads **Markdown** files—the standard for modern documentation—and extracts ("tangles") source code from code blocks defined within them.

It allows you to write your program as a Markdown document, explaining your logic step-by-step, while the tool handles the extraction of executable code.

### Features

*   **Markdown Support:** Uses standard Markdown files as the source of truth.
*   **Code Extraction:** Scans files for code blocks and composes them into source files.
*   **Simple CLI:** Easy-to-use command line interface.

## Getting Started

### Prerequisites

*   Go (Golang) installed on your system.

### Installation

Clone the repository and build the project using the provided `Makefile`:

```bash
make build
```

This will create the `literate` executable.

### Usage

Run the tool by providing input Markdown files and an optional output directory.

```bash
./literate [flags] <input-files>
```

#### Flags

*   `-output <path>`: Specify the output path for generated code.
*   `-verbose`: Enable verbose logging to see the extraction process in detail.

#### Example

```bash
./literate -output ./src -verbose design_doc.md
```

### CMake Integration

If you are working with C/C++ projects, you can use the provided CMake module to integrate literate programming into your build process.

1.  Copy `cmake/literate.cmake` to your project's module path (e.g., `cmake/literate.cmake`).
2.  Include and initialize it in your `CMakeLists.txt`:

```cmake
# Include the module
include(cmake/literate.cmake)

# Initialize 'literate' (fetches the tool from GitHub and builds it)
# Pass the git tag/branch you want to use (e.g., "main" or "v1.0.0")
literate_init(main)
```

3.  Use `literate_project` to generate sources and add them to your target:

```cmake
# Define your executable or library target first
add_executable(my_app main.c)

# Generate code from markdown and attach to the target
literate_project(my_app
    INPUT_FILES
        src/logic.md
        src/algorithms.md
    OUTPUT_DIR
        ${CMAKE_CURRENT_BINARY_DIR}/generated
    OUTPUTS
        ${CMAKE_CURRENT_BINARY_DIR}/generated/logic.c
        ${CMAKE_CURRENT_BINARY_DIR}/generated/logic.h
        ${CMAKE_CURRENT_BINARY_DIR}/generated/algorithms.c
)
```

This will automatically build the `literate` tool, generate the specified source files from your Markdown inputs, and compile them as part of your application.

### Development

To run tests or vet the code:

```bash
make test
make vet
```

## Source Code

This project is self-hosting. The following sections contain the source code used to build the `literate` tool.

### Entry Point

#### Argument Parsing

We first define the arguments to the `literate` application. Outside of the
list of input markdown files, we also support verbose logging and specifying a
output directory. If no output directory is provided, we just default to the
current working directory.

```go {name="argument_parsing"}
//--=====================================================================--
//--== ARGUMENT PARSING
//--=====================================================================--

var output string
var verbose bool

flag.StringVar(&output, "output", "", "output path for generated code")
flag.BoolVar(&verbose, "verbose", false, "verbose logging")
flag.Parse()

inputs := flag.Args()
```

#### Logging Setup

We will use the `slog` package for structured logging to standard output. If
the `verbose` argument is set we will include debug level log messages to the
output.

```go {name="logging_setup"}
//--=====================================================================--
//--== SETUP LOGGING
//--=====================================================================--

level := slog.LevelInfo
if verbose {
	level = slog.LevelDebug
}

handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: level,
})

slog.SetDefault(slog.New(handler))
```

#### Main

The `main` will parse the command line arguments, setup logging, and then
invoke the literate code generation process by calling `GenerateCode`.

```go name="main" filename="main.go"
package main

import (
	"flag"
	"log/slog"
	"os"
)

func main() {
    {{include "argument_parsing"}}

    {{include "logging_setup"}}

	//--=====================================================================--
	//--== GENERATE SOURCE CODE FILES
	//--=====================================================================--

	if err := GenerateCode(inputs, output); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
```

### Code Generation

#### Code Block Attributes

In order for the literate tool to stitch together the code blocks that
reference other code blocks, we need to add parse some attributes from the
Markdown. 

We will start by defining a regular expression that will parse the key value
attribute pairs.

```go {name="attribute_pattern"}
var AttributePattern = regexp.MustCompile(`(\w+)="([^"]*)"`)
```

The only two attributes the tool needs is a `name` and optional `filename`. The
`name` attribute is used to declare a unique name for the code block that
others code blocks can include. If a code block has a `filename` attribute,
then the code block output will be stored in that file.

We will define a `CodeBlockAttributes` struct for holding these values.

```go {name="code_block_attributes"}
type CodeBlockAttributes struct {
	Name     string
	Filename string
}
```

On the `CodeBlockAttributes` type we will add a function to parse the
attributes out of an attributes string using the `AttributePattern` we defined
above.

```go {name="code_block_attributes_parse"}
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
```

#### Code Block

For representing a single code block, we will define a type that contains the
attributes associated with the block along with the contents of the block.

```go {name="code_block"}
type CodeBlock struct {
	Attributes CodeBlockAttributes
	Body       string
}
```

#### Code Block Storage

For storing all of the `CodeBlock` instances that are found while parsing the
Markdown files, we will create a `CodeBlockStorage` type that will be the used
to lookup blocks and files by name.

```go {name="code_block_storage"}
type CodeBlockStorage struct {
	Blocks map[string]*CodeBlock
	Files  map[string]*CodeBlock
}
```

For instantiating a instance of a `CodeBlockStorage` we will add a new function for it.

```go {name="new_code_block_storage"}
func NewCodeBlockStorage() *CodeBlockStorage {
	return &CodeBlockStorage{
		Blocks: make(map[string]*CodeBlock),
		Files:  make(map[string]*CodeBlock),
	}
}

```

Finally, we will add an `AddCodeBlock` function to the type so that we can validate its uniqueness
before adding it to the storage.

```go {name="add_code_block"}
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
```

With these now in place, we will group this code together inside a single
`codeblock.go` file.

```go name="codeblock" filename="codeblock.go"
package main

import (
	"fmt"
	"regexp"
)

//--=====================================================================--
//--== CODE BLOCK ATTRIBUTES
//--=====================================================================--

{{include "attribute_pattern"}}

{{include "code_block_attributes"}}

{{include "code_block_attributes_parse"}}

//--=====================================================================--
//--== CODE BLOCK
//--=====================================================================--

{{include "code_block"}}

//--=====================================================================--
//--== CODE BLOCK STORAGE
//--=====================================================================--

{{include "code_block_storage"}}

{{include "new_code_block_storage"}}

{{include "add_code_block"}}

```

#### Markdown Parser

The Markdown parser will be responsible for extracting all of the code blocks
from input files. We will do this by creating a function `ParseMarkdownFile`
which takes a filename as an argument and returns an iterator over each code
block.

```go {name="parse_markdown_function"}
func ParseMarkdownFile(filename string) iter.Seq2[*CodeBlock, error] {
	return func(yield func(*CodeBlock, error) bool) {
		slog.Debug("parsing markdown file", slog.String("filename", filename))

        {{include "parse_file"}}

        {{include "extract_code_blocks"}}
	}
}
```

To parse a Markdown file we will use the `goldmark` Go package. 

```go {name="parse_file"}
//--=====================================================================--
//--== PARSE SOURCE FILE DOCUMENT
//--=====================================================================--

parser := goldmark.DefaultParser()

content, err := ioutil.ReadFile(filename)
if err != nil {
	yield(nil, fmt.Errorf("failed to read file '%s': %w", filename, err))
	return
}

reader := text.NewReader(content)
doc := parser.Parse(reader)
```

Once the document has been parsed, we can now walk the Abstract Syntax Tree
(AST) to find each of the code blocks.

```go {name="extract_code_blocks"}
//--=====================================================================--
//--== EXTRACT CODE BLOCKS
//--=====================================================================--

err = ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	if codeBlockNode, ok := node.(*ast.FencedCodeBlock); ok {
		var codeBlock CodeBlock

        {{include "parse_code_block_attributes"}}

        {{include "extract_code_block_body"}}

		//--=====================================================================--
		//--== RETURN NEXT CODE BLOCK
		//--=====================================================================--

		slog.Debug("extracted code block", slog.String("name", codeBlock.Attributes.Name))

		if !yield(&codeBlock, nil) {
			return ast.WalkStop, nil
		}
	}

	return ast.WalkContinue, nil
})
if err != nil {
	yield(nil, fmt.Errorf("failed to parse file '%s': %w", filename, err))
	return
}
```

For each code block we encounter in the AST, we will first need to parse the attributes. If there is
no name attribute we will skip the block.

```go {name="parse_code_block_attributes"}
//--=====================================================================--
//--== PARSE CODE BLOCK ATTRIBUTES
//--=====================================================================--

info := string(codeBlockNode.Info.Value(content))
codeBlock.Attributes.Parse(info)

if codeBlock.Attributes.Name == "" {
	return ast.WalkContinue, nil
}
```

Now that we know this is a code block we will need to reference. We can extract
the body of the block from the AST node.

```go {name="extract_code_block_body"}
//--=====================================================================--
//--== EXTRACT CODE BLOCK BODY
//--=====================================================================--

var codeBody bytes.Buffer

lines := codeBlockNode.Lines()

for i := 0; i < lines.Len(); i++ {
	line := lines.At(i)
	codeBody.Write(line.Value(content))
}

codeBlock.Body = codeBody.String()
```

And here is the output file for the parser function.

```go name="parser" filename="parser.go"
package main 

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"iter"
	"log/slog"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

{{include "parse_markdown_function"}}
```

#### Code Processor

At this point we now have code for parsing and extracting code blocks from
Markdown files and storing them into easy to reference maps. We will now need
some code that will detangle all the code block references and generate the
output source code files.

We will start by defining the `Processor` struct which holds the configuration
for the generation process.

```go {name="processor_struct"}
type Processor struct {
	OutputPath string
}

func NewProcessor(output string) *Processor {
	return &Processor{
		OutputPath: output,
	}
}
```

The `GenerateCodeFiles` method iterates over all files defined in the storage,
processes their content, and writes them to disk.

```go {name="generate_code_files"}
func (p Processor) GenerateCodeFiles(blocks *CodeBlockStorage) error {
	processor := NewCodeBlockProcessor(blocks)

	for _, block := range blocks.Files {
		path := filepath.Join(p.OutputPath, block.Attributes.Filename)

		code, err := processor.ProcessCodeBlock(block.Attributes.Name)
		if err != nil {
			return err
		}

		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return fmt.Errorf("failed to create directory for '%s': %w", path, err)
		}

		if err := ioutil.WriteFile(path, []byte(code), 0644); err != nil {
			return fmt.Errorf("failed to write to file '%s': %w", path, err)
		}

		slog.Debug("saved code file", slog.String("filename", block.Attributes.Filename))
	}

	return nil
}
```

To handle the actual expansion of code blocks and the `include` function, we
define a `CodeBlockProcessor`. It maintains the state to detect recursion loops.

```go {name="code_block_processor_struct"}
type CodeBlockProcessor struct {
	Funcs   template.FuncMap
	Storage *CodeBlockStorage
	State   map[string]bool
}

func NewCodeBlockProcessor(storage *CodeBlockStorage) *CodeBlockProcessor {
	processor := &CodeBlockProcessor{
		Storage: storage,
		State:   make(map[string]bool),
	}

	processor.Funcs = template.FuncMap{
		"include": processor.ProcessCodeBlock,
	}

	return processor
}
```

The `ProcessCodeBlock` method resolves a block by name, parses it as a Go
template, and executes it. This allows for nested inclusions.

```go {name="process_code_block"}
func (p *CodeBlockProcessor) ProcessCodeBlock(name string) (string, error) {
	if p.State[name] {
		return "", fmt.Errorf("recursive block inclusion: '%s'", name)
	}

	p.State[name] = true
	defer func() { p.State[name] = false }()

	block, ok := p.Storage.Blocks[name]
	if !ok {
		return "", fmt.Errorf("code block '%s' not found", name)
	}

	tmpl, err := template.New(name).Funcs(p.Funcs).Parse(block.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse code block '%s': %w", name, err)
	}

	var code bytes.Buffer

	if err := tmpl.Execute(&code, nil); err != nil {
		return "", fmt.Errorf("failed to generate code block '%s': %w", name, err)
	}

	return strings.TrimSpace(code.String()), nil
}
```

Finally, we combine these parts into `processor.go`.

```go name="processor" filename="processor.go"
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//--=====================================================================--
//--== PROCESSOR
//--=====================================================================--

{{include "processor_struct"}}

{{include "generate_code_files"}}

//--=====================================================================--
//--== CODE BLOCK PROCESSOR
//--=====================================================================--

{{include "code_block_processor_struct"}}

{{include "process_code_block"}}
```

#### Code Generator

With the code block storage and parser now defined, we only need to create the `GenerateCode` function that
our `main` calls.

This function will handle:
- Parsing each input Markdown file
- Constructing a code block storage from each code block
- Processing each code block output file

```go name="generate" filename="generate.go"
package main

func GenerateCode(inputs []string, output string) error {
	storage := NewCodeBlockStorage()

	for _, input := range inputs {
		for block, err := range ParseMarkdownFile(input) {
			if err != nil {
				return err
			}

			storage.AddCodeBlock(block)
		}
	}

	processor := NewProcessor(output)

	if err := processor.GenerateCodeFiles(storage); err != nil {
		return err
	}

	return nil
}
```

### Testing

To test the project we will have the generated code generate itself again by
parsing this `README.md` file.

We will use the `testify` package for performing the tests. Here is the
skeleton for the test file.

```go {name="main_test" filename="main_test.go"}
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
    "golang.org/x/tools/go/packages"
)

func TestLiterate(t *testing.T) {
	//--========================================--
	//--== CREATE TEMPORARY OUTPUT DIRECTORY
	//--========================================--
    {{include "test_create_output_directory"}}

	//--========================================--
	//--== GENERATE CODE
	//--========================================--
    {{include "test_generate_code"}}

	//--========================================--
	//--== VALIDATE GENERATED CODE
	//--========================================--
    {{include "test_validate_code"}}
}
```

We will have the test create a temporary output directory to store the
generated source code files. This is because we don't want to dirty the
existing source code files for the project.

```go {name="test_create_output_directory"}
    outputDirectory := t.TempDir()
```

We can now call the `GenerateCode` function with the `README.md` file as the
only input and this output directory as the destination.

```go {name="test_generate_code"}
    err := GenerateCode([]string{`README.md`}, outputDirectory)
    require.NoError(t, err)
```

To validate the generated code we will use the builtin `go/parser` package.

```go {name="test_validate_code"}
cfg := &packages.Config{
    Mode:  packages.NeedTypes | packages.NeedSyntax | packages.NeedImports,
    Dir:   outputDirectory, 
    Tests: false,
}

pkgs, err := packages.Load(cfg, ".")
require.NoError(t, err)

for _, pkg := range pkgs {
	assert.Empty(t, pkg.Errors)
}
```



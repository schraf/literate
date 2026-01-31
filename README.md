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

# Build System

To build this project, we use a `Makefile` to orchestrate the build steps. The
build process relies on the [literate](https://github.com/schraf/literate) tool
to generate the source code from these Markdown files.

## Dependencies

We provide a convenient way to install the `literate` tool using a shell script.

```makefile {name="makefile" filename="Makefile"}
.PHONY: all build clean deps test

all: test build

deps:
	curl -fsSL https://raw.githubusercontent.com/schraf/literate/main/install.sh | bash

build:
	$$(go env GOPATH)/bin/literate [YOUR MARKDOWN FILES]
	[COMMAND TO BUILD PROJECT]

test:
    [COMMEND TO RUN PROJECT TESTS]

clean:
	rm -f [BUILD ARTIFACTS]
```


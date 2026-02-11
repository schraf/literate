# Orchestration

This is the conclusion of the document. It ties everything together.

## Main Execution

The `main()` function is the high-level script that calls the loaders, runs the
algorithms, and saves the output.

```{name="main_execution"}
[SETUP ENVIRONMENT]

[INITIALIZE INTERFACES]

[INVOKE CORE LOCATION]

[EXPORT RESULTS]
```

## Source File

```{name="main" filename="main.EXT"}
{{include "file_header"}}

// ╔════════════════════════════════════════════════════════════════════╗
// ║ MAIN                                                               ║
// ╚════════════════════════════════════════════════════════════════════╝

{{include "main_execution"}}
```

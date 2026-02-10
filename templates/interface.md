# 4. The Interface

Now that we have defined the environment and common tools, we define how data
enters and leaves the system. In a narrative sense, I/O is often "dirty" work.
Keeping it separate from your pure algorithms ensures your core logic remains
testable and clean.

## Inputs

Functions to read CSVs, APIs, databases, or user input.

```go {name="interface_inputs"}
[INPUT FUNCTIONS]
```

## Outputs

Functions to save results or generate reports.

```go {name="interface_outputs"}
[OUTPUT FUNCTIONS]
```

## Source File

```{name="interface" filename="interface.EXT"}
{{include "file_header"}}

//--=====================================================================--
//--== INPUTS
//--=====================================================================--

{{include "interface_inputs"}}

//--=====================================================================--
//--== OUTPUTS
//--=====================================================================--

{{include "interface_outputs"}}
```


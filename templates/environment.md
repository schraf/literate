# 2. The Environment

This section bridges the gap between the prose and the machine's requirements.
It handles global imports, configuration, and logging.

## Imports

We gather external libraries here.

*   *Explain why you chose specific libraries here.*

```{name="imports"}
[IMPORT CODE]
```

## Global Configuration

Define constants, file paths, and hyperparameters here.

```{name="configuration"}
[CONFIGS]
[CONSTANTS]
```

## Logging Strategy

We define our logging "vocabulary" (log levels, formats) immediately so it is
available for every subsequent function.

```{name="logging"}
[LOGGING INTERFACE]
```

## Source File

We combine these into the environment source file.

```{name="environment" filename="environment.EXT"}
{{include "file_header"}}

{{include "imports"}}

//--=====================================================================--
//--== CONFIGURATION
//--=====================================================================--

{{include "configuration"}}

//--=====================================================================--
//--== LOGGING
//--=====================================================================--

{{include "logging"}}
```

# Module: [Module Name]

This module represents a core component of the system.

## The Vocabulary (Data Models)

This corresponds to the **"Nouns"** of your system. We define our data
structures early so the reader understands the objects being manipulated.

### Schema Definitions

```{name="module_datamodels"}
[DATA MODELS]
```

### Invariants

Explain the rules of your data (e.g., *"The 'ID' must never be empty"*).

```{name="module_invariants"}
[INVARIANT FUNCTIONS]
```

## The Core (Algorithms)

This is the **"Climax"** of the module—the **"Verbs"**. This section should be
text-heavy, explaining the logic deeply.

### Logic Chunk A: Pre-processing

First, we prepare the data for the operation.

```{name="module_preprocessing"}
[PREPROCESSING FUNCTIONS]
```

### Logic Chunk B: Core Transformation

This is the heart of the algorithm.

```{name="module_processing"}
[PROCESSING FUNCTIONS]
```

### Logic Chunk C: Post-processing

Cleanup and return.

```{name="module_postprocessing"}
[POSTPROCESSING FUNCTIONS]
```

### Full Algorithm Implementation

We assemble the chunks into the final function.

```go {name="module_algorithm"}
//--========================================--
//--== PREPROCESSING
//--========================================--

{{include "module_preprocessing"}}

//--========================================--
//--== PROCESSING
//--========================================--

{{include "module_processing"}}

//--========================================--
//--== POSTPROCESSING
//--========================================--

{{include "module_postprocessing"}}

//--========================================--
//--== MODULE ENTRY POINT
//--========================================--

[ENTRY POINTS]
```

## Source File

```{name="module_source" filename="module_name.EXT"}
{{include "file_header"}}

//--=====================================================================--
//--== DATA MODELS
//--=====================================================================--

{{include "module_datamodels"}}

//--=====================================================================--
//--== DATA VALIDATION
//--=====================================================================--

{{include "module_invariants"}}

//--=====================================================================--
//--== ALGORITHMS
//--=====================================================================--

{{include "module_algorithm"}}
```


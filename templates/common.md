# 3. The Common Toolkit

These are the small helpers that don't deserve their own chapter but are
necessary for the heavy lifting. Placing them here prevents interrupting the
flow of the "good stuff" later.

## Conversion Helpers

Boring but necessary code for converting between data formats.

``` {name="common_conversions"}
[CONVERSION FUNCTIONS]
```

## Generic Utilities

Simple calculations or wrappers used in multiple places.

```{name="common_utilities"}
[UTILITY FUNCTIONS]
```

## Source File

```{name="common" filename="common.EXT"}
{{include "file_header"}}

//--=====================================================================--
//--== COMMON CONVERSION FUNCTIONS
//--=====================================================================--

{{include "common_conversions"}}

//--=====================================================================--
//--== COMMON UTILITY FUNCTIONS
//--=====================================================================--

{{include "common_utilities"}}
```

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

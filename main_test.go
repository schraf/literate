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
	outputDirectory := t.TempDir()

	//--========================================--
	//--== GENERATE CODE
	//--========================================--
	err := GenerateCode([]string{`README.md`}, outputDirectory)
	require.NoError(t, err)

	//--========================================--
	//--== VALIDATE GENERATED CODE
	//--========================================--
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
}

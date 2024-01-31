package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yourselfhosted/gomark/ast"
	"github.com/yourselfhosted/gomark/parser/tokenizer"
	"github.com/yourselfhosted/gomark/restore"
)

func TestEscapingCharacterParser(t *testing.T) {
	tests := []struct {
		text string
		node ast.Node
	}{
		{
			text: `\# 123`,
			node: &ast.EscapingCharacter{
				Symbol: "#",
			},
		},
	}

	for _, test := range tests {
		tokens := tokenizer.Tokenize(test.text)
		node, _ := NewEscapingCharacterParser().Match(tokens)
		require.Equal(t, restore.Restore([]ast.Node{test.node}), restore.Restore([]ast.Node{node}))
	}
}

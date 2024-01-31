package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yourselfhosted/gomark/ast"
	"github.com/yourselfhosted/gomark/parser/tokenizer"
	"github.com/yourselfhosted/gomark/restore"
)

func TestAutoLinkParser(t *testing.T) {
	tests := []struct {
		text string
		link ast.Node
	}{
		{
			text: "<https://example.com)",
			link: nil,
		},
		{
			text: "<https://example.com>",
			link: &ast.AutoLink{
				URL: "https://example.com",
			},
		},
		{
			text: "https://example.com",
			link: &ast.AutoLink{
				URL:       "https://example.com",
				IsRawText: true,
			},
		},
	}

	for _, test := range tests {
		tokens := tokenizer.Tokenize(test.text)
		node, _ := NewAutoLinkParser().Match(tokens)
		require.Equal(t, restore.Restore([]ast.Node{test.link}), restore.Restore([]ast.Node{node}))
	}
}

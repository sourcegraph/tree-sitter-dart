package tree_sitter_dart_test

import (
	"testing"

	tree_sitter_dart "github.com/UserNobody14/tree-sitter-dart/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_dart.Language())
	parser := tree_sitter.NewParser()
	defer parser.Close()
	if err := parser.SetLanguage(language); err != nil {
		t.Fatal(err)
	}
	tree := parser.Parse([]byte("void main() { print('Hello, Dart!'); }"), nil)
	defer tree.Close()
	if root := tree.RootNode(); root.Kind() != "program" || root.HasError() {
		t.Fatalf("Unexpected parse: %s", root.ToSexp())
	}
}

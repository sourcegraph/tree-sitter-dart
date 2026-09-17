from unittest import TestCase

import tree_sitter, tree_sitter_dart


class TestLanguage(TestCase):
    def test_can_parse_dart(self):
        language = tree_sitter.Language(tree_sitter_dart.language())
        parser = tree_sitter.Parser(language)
        tree = parser.parse(b"void main() { print('Hello, Dart!'); }")
        self.assertEqual(tree.root_node.type, "program")
        self.assertFalse(tree.root_node.has_error)

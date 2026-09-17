/// <reference types="node" />

const assert = require("node:assert");
const { test } = require("node:test");

test("can parse Dart", () => {
  const parser = new (require("tree-sitter"))();
  parser.setLanguage(require("."));
  const tree = parser.parse("void main() { print('Hello, Dart!'); }");
  assert.strictEqual(tree.rootNode.type, "program");
  assert.strictEqual(tree.rootNode.hasError, false);
});

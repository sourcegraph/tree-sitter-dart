# tree-sitter-dart
This is a tree-sitter grammar written for the dart programming language. We attempt to adhere as closely as possible to the dart language spec. Initially it was started with a copy of the tree-sitter Java grammar which is why there may be a few relics included in here. For the sake of simplifying the syntax tree, many items were made inlined with tree-sitter's "underscore" method of writing rules.

# Getting Started
- Go to the project directory
- run `npm install` (first time)
- run `npm run build_init` (first time) `npm run build` (subsequent times)
- run `npm run test`

# Tree-sitter 0.23 compatibility branch

This branch uses the grammar from upstream commit
[`0038280`](https://github.com/UserNobody14/tree-sitter-dart/commit/0038280),
with `grammar.js`, `src/scanner.c`, and queries unchanged. The parser and WASM
module are regenerated with Tree-sitter CLI 0.23.2 for language ABI 14.

Rust consumers now use `tree_sitter_dart::LANGUAGE.into()` instead of
`tree_sitter_dart::language()`. The binding depends on `tree-sitter-language`;
its tests use the Tree-sitter 0.23 runtime. Go and Python bindings also target
0.23 (requiring Go 1.23+ and Python 3.9+). There is no Node runtime 0.23 release,
so Node tests use 0.22.4, which supports ABI 14.

Tree-sitter 0.23 error recovery returns `(ERROR (ERROR (identifier)))` for the
incomplete file `d`, whereas 0.26 returns `(program (ERROR (identifier)))` with
the same parser. Only that corpus expectation has been adjusted.

Verification commands:

```sh
npm run build_init
npm test
cargo test --locked
go test ./bindings/go
python -m pip install '.[core]'
python -m unittest discover -s bindings/python/tests -v
```

To rebuild the WASM module, install Emscripten and run
`npx tree-sitter build --wasm`. Use `npx tree-sitter generate --no-bindings`
when regenerating the parser; the 0.23 binding generator would otherwise
replace the customized package scripts.

# To test a single highlight file
- run `tree-sitter highlight test/highlight/types.dart`

# To test a single test file
- run `tree-sitter test -f 'testcasefilter'`
- for example `tree-sitter test -f 'dart string literals'`

# To show the output of a parse for a sample file (for example while debugging highlight issues)
- run `tree-sitter parse path/to/file.dart`

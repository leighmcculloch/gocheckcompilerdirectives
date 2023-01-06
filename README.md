# gocheckdirectives

Check that go directories (`//go:` comments) are valid and catch easy mistakes.

For example, directives like `//go:generate`, `//go:embed`, `//go:build`, etc.

## Why

Go directives are comments in the form of `//go:` that provide an instruction
to the compiler.

Go directives are easy to make mistakes with. The linter will detect the
following mistakes.

1. Adding a space in between the comment bars and the first character, e.g. `//
go:`, will cause the compiler to silently ignore the comment.

2. Mistyping a directives name, e.g. `//go:embod`, will cause the compiler to silently ignore the comment.

## Install

```
go install github.com/leighmcculloch/gocheckdirectives@latest
```

## Usage

```
gocheckdirectives [package]
```

```
gocheckdirectives ./...
```

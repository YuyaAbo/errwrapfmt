# errwrapfmt

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/YuyaAbo/errwrapfmt)
[![Go Reference](https://pkg.go.dev/badge/github.com/YuyaAbo/errwrapfmt.svg)](https://pkg.go.dev/github.com/YuyaAbo/errwrapfmt)

errwrapfmt finds wrong error wrap format.

Using the wrong format makes the stack trace redundant. An example is https://go.dev/play/p/pPTUFmBvToa
```
": %w" // OK
":%w"  // NG
"%w"   // NG
```

## Install

```bash
go install github.com/YuyaAbo/errwrapfmt/cmd/errwrapfmt@latest
```

## Useage

```bash
errwrapfmt ./...
```

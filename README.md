# errwrapfmt

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/YuyaAbo/errwrapfmt)

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

# errwrapfmt

errwrapfmt finds wrong error wrap format.

Using the wrong format makes the stack trace redundant. An example is https://go.dev/play/p/pPTUFmBvToa
```
"error: %w" // OK
"error :%w" // NG
"error %w"  // NG
```

## Install

```bash
go install github.com/YuyaAbo/errwrapfmt/cmd/errwrapfmt@latest
```

## Useage

```bash
errwrapfmt ./...
```

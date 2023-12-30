## Dump IR log
Use a go of version 1.19.
```shell
go tool compile -W 2 main.go > main.ir 
```

## Compile to WASM.

### JS/Browser target

>Pass the '-w' flag to the linker to omit the debug information (for example, go build -ldflags=-w prog.go).

```shell
GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o main.wasm main.go
```

### Other target
[Go reference](https://go.dev/blog/wasi)

```shell
GOOS=wasip1 GOARCH=wasm go build -o main.wasm main.go
```

## Go SSA tool
Dump the ssa of a package.

```shell
ssadump -build=F main.go
```

Interpret a trivial go package.

```shell
ssadump -run hi.go
```

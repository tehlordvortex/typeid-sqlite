# typeid-sqlite

An SQLite extension that wraps the [Go TypeID](https://github.com/jetify-com/typeid-go) implementation, exposing two methods: `typeid_generate_text` and `typeid_check_text`.

## Compiling

You should compile this extension with the same toolchain you use for SQLite (or LibSQL) itself.
Building this as a shared library is straightforward with Go 1.24.4+ and a build toolchain installed:

```sh
go build -buildmode=c-shared -o dist/typeid.so
```

## API

### typeid_generate_text(prefix)

Generates a TypeID string with the given `prefix`. This will fail if `prefix` isn't a valid TypeID prefix

### typeid_check_text(prefix, value)

Returns `1` if the given value is a valid TypeID with the given `prefix`. Returns `0` in all other cases.

## Caveats

- 3.3MB binary to expose 2 functions. ¯\_(ツ)_/¯

fftoml - A TOML config file parser
----------------------------------

This is a super basic alternative toml parser to the existing: https://github.com/peterbourgon/ff/blob/master/fftoml/fftoml.go

The purpose of this version is to support TOML tables. When presented with a table this parser concatenates the nested
keys with a `-` before settings them via the `set()` function.

For example:

```toml
[config]
key = "value"
[config.property]
key = "other-value"
```

Can be parsed into the following flags:

```go
flag.String("config-key", "", "some config flag")
flag.String("config-property-key", "", "some other config flag")
```

# Go Sane

Go Sane is a set of linters that forbids some Go features I think should never have been part of/added to Go:
- generics
- iterators / range-over-func
- inline if syntax
- init() functions
- global variables (with a few exceptions)
- global-state functions like http.ListenAndServe() or viper.Set/Get*()

It:
1. Defines two new linters: `nogenerics` and `noiter`.
2. Provide custom patterns for the existing [forbidigo](https://github.com/ashanbrown/forbidigo), to match global-state functions.
3. Uses the existing [noinlineerr](https://github.com/AlwxSin/noinlineerr), [gochecknoinits](https://github.com/leighmcculloch/gochecknoinits), [gochecknoglobals](https://github.com/leighmcculloch/gochecknoglobals).

### Why?

A great read on this topic: [Go evolves in the wrong direction](https://itnext.io/go-evolves-in-the-wrong-direction-7dfda8a1a620), by Aliaksandr Valialkin.

### Usage

- `golangci-lint`: Just copy all or part of `.golangci.yaml` into your own config.
- standalone: Run `nogenerics` or `noiter` binaries themselves or with `go vet`.

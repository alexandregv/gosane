# Go Sane

Go Sane is a set of linters that forbids some Go features I think should never have been part of/added to Go:
- generics
- iterators / range-over-func
- inline if syntax
- init() functions
- global variables (with a few exceptions)
- global-state functions like http.ListenAndServe() or viper.Set/Get*()

The repository contains a `.golangci.yaml` template that contains both external linters and new linters defined in this repository.

External linters:
- [noinlineerr](https://github.com/AlwxSin/noinlineerr)
- [gochecknoinits](https://github.com/leighmcculloch/gochecknoinits)
- [gochecknoglobals](https://github.com/leighmcculloch/gochecknoglobals)
- [forbidigo](https://github.com/ashanbrown/forbidigo) with curated patterns

New linters:
- nogenerics
- noiter

`nogenerics` and `noiter` can also be used independentely, like any other linter.

### Why?

A great read on this topic: [Go evolves in the wrong direction](https://itnext.io/go-evolves-in-the-wrong-direction-7dfda8a1a620), by Aliaksandr Valialkin.

### Usage

- `golangci-lint`: Just copy all or part of `.golangci.yaml` into your own config.
- standalone `go vet`: Run `nogenerics` or `noiter` binaries themselves or with `go vet`.

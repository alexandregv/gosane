# Go Sane

Go Sane is a set of linters that forbids some Go features I think should never have been part of/added to Go:
- generics
- iterators / range-over-func
- short-if syntax
- init() functions
- global variables (with a few exceptions)
- global-state functions like http.ListenAndServe(), log.Print() or viper.Set/Get*()

This repository contains three distinct parts:
1. New linters: `nogenerics`, `noiter`, `noshortif`. Usable on their own or with other tools.
2. The `gosane` command itself: an all-in-one linter combining these new linters and some existing ones.
3. A `.golangci.yml` template, combining these new linters and some existing ones.

> [!NOTE]
> `nogenerics`, `noiter` and `noshortif` are in the process of being added to `golangci-lint` list of linters.
> `gosane` will not, as it's a multi-linter and uses some existing ones.

## Linters

### New linters defined in this repository

<details>
<summary>nogenerics</summary>

Forbids use of Go 1.18's generics.  
Yes, sometimes they are useful. Often not really, they just clutter and complexify things a lot.  
See [Go evolves in the wrong direction](https://itnext.io/go-evolves-in-the-wrong-direction-7dfda8a1a620), by Aliaksandr Valialkin.

</details>

<details>
<summary>noiter</summary>

Forbids use of Go 1.23's iterators.  
They complexify for not so much gain.  
See [Go evolves in the wrong direction](https://itnext.io/go-evolves-in-the-wrong-direction-7dfda8a1a620), by Aliaksandr Valialkin.

</details>

<details>
<summary>noshortif</summary>

Forbids short-if syntax.  
In Go, you should have only one way to do something.

</details>

### Existing linters used by gosane and `.golangci.yml` template

<details>
<summary><a href="https://github.com/ashanbrown/forbidigo">forbidigo</a> with custom patterns</summary>

Forbids global-state functions like `http.ListenAndServe()`, `log.Print()` or `viper.Set/Get*()`.  
They make testing more complex. `spf13/viper` now regrets providing them: https://github.com/spf13/viper/#should-viper-be-a-global-singleton-or-passed-around and https://github.com/spf13/viper/issues/1855.

</details>

<details>
<summary><a href="https://github.com/leighmcculloch/gochecknoinits">gochecknoinits</a></summary>

Forbids `init()` functions.  
They hide behavior and are never the only choice. `spf13/cobra` made them famous but now regrets it: https://github.com/spf13/cobra/issues/1862 and https://github.com/spf13/cobra/issues/1335.

Note: I actually had to reimplement it for `gosane`, because `gochecknoinits` doesn't use the `go/analysis` standard. But it's exactly the same thing.

</details>

<details>
<summary><a href="https://github.com/leighmcculloch/gochecknoglobals">gochecknoglobals</a></summary>

Forbids global variables, except for some legitimate cases: the `version` variables, `errors.New()`, `regex.MustCompile()`, etc.

</details>

## Usage

### golangci-lint

Just copy all or part of `.golangci.yaml` into your own config. This allows you to pick the rules you want, or modify them.

### Standalone

`nogenerics`, `noiter`, `shortif` or the all-in-one `gosane` are binaries that can be used by themselves, or with `go vet`.  

```shell
# go install + manual run
go install github.com/alexandregv/gosane/cmd/gosane@latest
gosane ./...

# go install + go vet
go install github.com/alexandregv/gosane/cmd/gosane@latest
go vet -vettool=$(which gosane) ./...

# go run + go vet
go vet -vettool=$(go run -n github.com/alexandregv/gosane/cmd/gosane@latest 2>&1) ./...

# go tool + go vet
go get -tool github.com/alexandregv/gosane/cmd/gosane@latest
go vet -vettool=$(go tool -n github.com/alexandregv/gosane/cmd/gosane) ./...
```
Replace `gosane` with any of `nogenerics`, `noiter` or `noshortif` if you want just a specific linter.

# GoClassic

GoClassic is a set of golangci-lint rules that forbids every recent Go addition that I think should not have been added to Go.  
These features are:
- generics (can be re-enabled if really needed)
- yield / generators / iterators
- inline iferr
- init() 
- Go 1.27 keyed fields
- use of global-state function like http.ListenAndServe() in place of srv := http.Server{} + srv.ListenAndServe(), or same with viper.*

A great reading on this topic: [Go evolves in the wrong direction](https://itnext.io/go-evolves-in-the-wrong-direction-7dfda8a1a620), by Aliaksandr Valialkin.

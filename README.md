# classical

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/classical) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/classical?status.svg)](https://godoc.org/github.com/meirizarrygelpi/classical)

Package `classical` consist of tools for modeling classical bits and gates.

There is one type:

* `classical.Bit`

This type corresponds to a classical bit, which can have two possible states.

There are many functions. Some correspond to **irreversible** logic gates:

* `classical.And`
* `classical.NAnd`
* `classical.Or`
* `classical.NOr`
* `classical.XOr`
* `classical.NXOr`
* `classical.FanOut`

The other functions correspond to **reversible** gates:

* `classical.Not`
* `classical.Swap`
* `classical.CNot`
* `classical.Fredkin`
* `classical.Toffoli`
* `classical.Swap243`
* `classical.CCSwap`
* `classical.DeMultiplexer`

This package is for having fun and is not optimized for high-performance.
# dp [![GoDoc](https://godoc.org/github.com/armen/dp?status.png)](https://godoc.org/github.com/armen/dp) [![Build Status](https://travis-ci.org/armen/dp.svg?branch=master)](https://travis-ci.org/armen/dp) [![codecov](https://codecov.io/gh/armen/dp/branch/master/graph/badge.svg)](https://codecov.io/gh/armen/dp)

A pure Go implementation of [*Introduction to Reliable and Secure Distributed Programming*][dp] abstractions.

## Example Abstractions

- Job handler ([interface, properties](https://raw.githubusercontent.com/armen/dp/master/job/handler.go))
	- Synchronous job handler ([implementation](https://raw.githubusercontent.com/armen/dp/master/job/sync/sync.go), [algorithm](https://raw.githubusercontent.com/armen/dp/master/job/sync/sync.txt))
	- Asynchronous job handler ([implementation](https://raw.githubusercontent.com/armen/dp/master/job/async/async.go), [algorithm](https://raw.githubusercontent.com/armen/dp/master/job/async/async.txt))
- Job transformation and processing abstraction ([interface, properties](https://raw.githubusercontent.com/armen/dp/master/job/transformation_handler.go))
	- Job-Transformation by buffering ([implementation](https://raw.githubusercontent.com/armen/dp/master/job/transformation/transformation.go), [algorithm](https://raw.githubusercontent.com/armen/dp/master/job/transformation/transformation.txt))

## List of Abstractions

- Perfect point-to-point link ([interface, properties](https://raw.githubusercontent.com/armen/dp/master/link/perfect.go))
	- TCP based perfect peer-to-peer link ([implementation](https://raw.githubusercontent.com/armen/dp/master/link/tcp/ppp/ppp.go))
- Perfect failure detector ([interface, properties](https://raw.githubusercontent.com/armen/dp/master/fd/perfect.go))
- Eventually perfect failure detector ([interface, properties](https://raw.githubusercontent.com/armen/dp/master/fd/eventually_perfect.go))

[dp]: http://distributedprogramming.net

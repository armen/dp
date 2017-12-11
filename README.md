# dp [![GoDoc](https://godoc.org/github.com/armen/dp?status.png)](https://godoc.org/github.com/armen/dp) [![Build Status](https://travis-ci.org/armen/dp.svg?branch=master)](https://travis-ci.org/armen/dp)

<img width="165" src="http://www.distributedprogramming.net/images/cover-5.png" align="right">

An implementation of [*Introduction to Reliable and Secure Distributed Programming*][dp] algorithms.

## List of Algorithms

### Chapter 1: Introduction
- **Module 1.1** Interface and properties of a job handler ([interface](https://raw.githubusercontent.com/armen/dp/master/job/handler.go))
	- **Algorithm 1.1** Synchronous Job Handler ([implementation](https://raw.githubusercontent.com/armen/dp/master/job/sync/sync.go), [algorithm](https://raw.githubusercontent.com/armen/dp/master/job/sync/sync.txt))
	- **Algorithm 1.2** Asynchronous Job Handler ([implementation](https://raw.githubusercontent.com/armen/dp/master/job/async/async.go), [algorithm](https://raw.githubusercontent.com/armen/dp/master/job/async/async.txt))
- **Module 1.2** Interface and properties of a job transformation and processing abstraction ([interface](https://raw.githubusercontent.com/armen/dp/master/job/transformation_handler.go))
	- **Algorithm 1.3** Job-Transformation by Buffering ([implementation](https://raw.githubusercontent.com/armen/dp/master/job/transformation/transformation.go), [algorithm](https://raw.githubusercontent.com/armen/dp/master/job/transformation/transformation.txt))

[dp]: http://distributedprogramming.net

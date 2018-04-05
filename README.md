# goClustering

[![Build Status](https://travis-ci.org/cipepser/goClustering.svg?branch=master)](https://travis-ci.org/cipepser/goClustering)
[![Coverage Status](https://coveralls.io/repos/github/cipepser/goClustering/badge.svg?branch=master)](https://coveralls.io/github/cipepser/goClustering?branch=master)

goClustering is an implementation of following clustering algorithm in Golang.

## Ward's method

Ward's method is a distance criterion in hierarchical cluster algorithm.
This method put two groups into one to minimize the total within-cluster variance. For detail, see [here](https://en.wikipedia.org/wiki/Ward%27s_method).

## How to Install

You can get goClustering by using go get:

```sh
go get github.com/cipepser/goClustering/...
```

## Example

```go
package main

import (
	"fmt"

	"github.com/cipepser/goClustering/ward"
)

func main() {
	X := [][]float64{
		{0, 0},
		{1, 0},
		{5, 0},
	}

	T := ward.Ward(X)

	fmt.Println(T)
}
```

## References
* [Ward's method](https://en.wikipedia.org/wiki/Ward%27s_method)
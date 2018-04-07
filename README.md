# goClustering

[![Build Status](https://travis-ci.org/cipepser/goClustering.svg?branch=master)](https://travis-ci.org/cipepser/goClustering)
[![Coverage Status](https://coveralls.io/repos/github/cipepser/goClustering/badge.svg?branch=master)](https://coveralls.io/github/cipepser/goClustering?branch=master)

goClustering is an implementation of following clustering algorithm in Golang.

## Ward's method

Ward's method is a distance criterion in hierarchical cluster algorithm.
This method puts two groups into one to minimize the total within-cluster variance. For detail, see [here](https://en.wikipedia.org/wiki/Ward%27s_method).

## How to Install

You can get goClustering by using go get:

```sh
go get github.com/cipepser/goClustering/...
```

To visualize the dendrogram, you also install `gonum/plot` and `cipepser/plot`.

```sh
go get gonum.org/v1/plot/...
go get github.com/cipepser/plot/...
```

## How to Use

This hierarchical cluster algorithm have two parts:
1. Execute Ward's method
1. Visualize

In execute Ward's method step, your input data must be `n * d` matrix, where n is an observation and d is a dimension of feature vectors.
As a result of this step, you can get the `ward.Tree` which is binary-tree of the result of hierarchical clustering with distance of each group.
You can execute Ward's method like this:

```go
T := ward.Ward(X)
```

Second stap is Visualize step. You can visualize the dendrogram of `ward.Tree` by `gonum/plot` like:

```go
d, _ := plotter.NewDendrogram(T)
```

Where, `T` is a result of first step. And you create a plot to add above `d`.

```go
p, err := plot.New()
p.Add(d)
```

You can also customize title and labels of x-axis and y-axis.

```go
p.Title.Text = "Dendrogram"
p.X.Label.Text = "data"
p.Y.Label.Text = "distance"
```

Name the leaf nodes and rotate names.

```go
p.NominalX("aaa", "bbb", "ccc", "ddd", "eee", "fff")
p.X.Tick.Label.Rotation = math.Pi / 3
p.X.Tick.Label.YAlign = draw.YCenter
p.X.Tick.Label.XAlign = draw.XRight
```

### Example

```go
package main

import (
	"math"

	"github.com/cipepser/goClustering/ward"
	"github.com/cipepser/plot/plotter"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	// input data set
	X := [][]float64{
		{0, 0},
		{2, 2},
		{1, 1},
		{2, -1.2},
		{3, 2.2},
		{3.5, 0.5},
	}

	// Ward's method
	T := ward.Ward(X)

	// draw the dendrogram
	d, err := plotter.NewDendrogram(T)
	if err != nil {
		panic(err)
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Add(d)

	p.Title.Text = "Dendrogram"
	p.X.Label.Text = "data"
	p.NominalX("aaa", "bbb", "ccc", "ddd", "eee", "fff")

	p.X.Tick.Label.Rotation = math.Pi / 3
	p.X.Tick.Label.YAlign = draw.YCenter
	p.X.Tick.Label.XAlign = draw.XRight

	p.Y.Label.Text = "distance"

	// save as a png file
	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}
}
```

### Result

![Result](https://github.com/cipepser/goClustering/blob/master/img/img.png)

## References
* [Ward's method](https://en.wikipedia.org/wiki/Ward%27s_method)
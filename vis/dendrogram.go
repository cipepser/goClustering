package vis

import (
	"github.com/cipepser/goClustering/ward"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// Dendrogram represents a dendrogram of given `ward.Tree`.
type Dendrogram struct {
	tree ward.Tree
	draw.LineStyle
}

// NewDendrogram creates a new dendrogram plotter for
// the given Tree.
func NewDendrogram(T ward.Tree) (*Dendrogram, error) {
	d := &Dendrogram{
		tree:      T,
		LineStyle: plotter.DefaultLineStyle,
	}

	return d, nil
}

// Plot implements the Plot method of the plot.Plotter interface.
func (d *Dendrogram) Plot(c draw.Canvas, plt *plot.Plot) {
	if len(d.tree) == 1 {
		return
	}

	trX, trY := plt.Transforms(&c)
	x := (1 / 1.05) * (plt.X.Max) / 2
	l := float64(len(d.tree)/2 + 1)
	width := x * l / (l - 1)
	d.strokeLine(c, d.tree, len(d.tree)-1, x, width, trX, trY)
}

// strokeLine draws the dendrogram.
func (d *Dendrogram) strokeLine(c draw.Canvas, T ward.Tree, idx int, x, w float64, trX, trY func(float64) vg.Length) {
	if T[idx].N == 1 {
		return
	}

	x0 := x - w*float64(T[T[idx].Right].N)/float64(T[idx].N)
	x1 := x0 + w
	y0 := T[idx].GetDist()
	y1 := y0

	// stroke horizontal line
	c.StrokeLine2(d.LineStyle, trX(x0), trY(y0), trX(x1), trY(y1))

	// stroke vertical line and call strokeLine recursively
	c.StrokeLine2(d.LineStyle, trX(x0), trY(y0), trX(x0), trY(T[T[idx].Left].GetDist()))
	d.strokeLine(c, d.tree, T[idx].Left, x0, w*float64(T[T[idx].Left].N)/float64(T[idx].N), trX, trY)

	c.StrokeLine2(d.LineStyle, trX(x1), trY(y1), trX(x1), trY(T[T[idx].Right].GetDist()))
	d.strokeLine(c, d.tree, T[idx].Right, x1, w*float64(T[T[idx].Right].N)/float64(T[idx].N), trX, trY)

}

// DataRange implements the DataRange method
// of the plot.DataRanger interface with 5% padding.
func (d *Dendrogram) DataRange() (xmin, xmax, ymin, ymax float64) {
	// return 0, float64(len(d.tree)/2 + 1), 0, d.tree[len(d.tree)-1].GetDist()
	return -0.05 * float64(len(d.tree)/2), 1.05 * float64(len(d.tree)/2), 0, 1.05 * d.tree[len(d.tree)-1].GetDist()
}

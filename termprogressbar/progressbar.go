package termuiprogressbar

import (
	"github.com/cheggaaa/pb"
)

type termuiprogressbar struct {
	progressbar *pb.ProgressBar
	Visible     bool
}

func NewProgressBar(total int, visible bool) *termuiprogressbar {
	pbar := termuiprogressbar{}
	pbar.progressbar = pb.New64(int64(total))
	pbar.Visible = visible
	return &pbar
}

func (pb *termuiprogressbar) Start() {
	if pb.Visible {
		pb.progressbar.Start()
	}
}

func (pb *termuiprogressbar) Increment() {
	if pb.Visible {
		pb.progressbar.Increment()
	}
}

func (pb *termuiprogressbar) FinishPrint(str string) {
	if pb.Visible {
		pb.progressbar.FinishPrint(str)
	}
}

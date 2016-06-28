package termuiprogressbar

import (
	"github.com/cheggaaa/pb"
)

type termuiprogressbar struct {
	progressbar *pb.ProgressBar
	visible     bool
}

// creates a new progressbar
func NewProgressBar(total int, visible bool) *termuiprogressbar {
	return &termuiprogressbar{
		progressbar: pb.New64(int64(total)),
		visible:     visible,
	}
}

// starts the progressbar
func (pb *termuiprogressbar) Start() {
	if pb.visible {
		pb.progressbar.Start()
	}
}

// increments the progressbar
func (pb *termuiprogressbar) Increment() {
	if pb.visible {
		pb.progressbar.Increment()
	}
}

// gets the progressbar to 100% and prints out a message
func (pb *termuiprogressbar) FinishPrint(str string) {
	if pb.visible {
		pb.progressbar.FinishPrint(str)
	}
}

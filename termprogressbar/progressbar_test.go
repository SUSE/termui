package termuiprogressbar_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/hpcloud/termui/termprogressbar"
	"github.com/stretchr/testify/assert"
)

// tests if the progressbar is visible if it is initialized with visible=true
func TestVisible(t *testing.T) {
	assert := assert.New(t)

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	bar := termuiprogressbar.NewProgressBar(10, true)
	bar.Start()
	bar.FinishPrint("")

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC

	assert.NotEmpty(out)
}

// tests if the progressbar is invisible if it is initialized with visible=false
func TestInvisible(t *testing.T) {
	assert := assert.New(t)

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	bar := termuiprogressbar.NewProgressBar(10, false)
	bar.Start()
	bar.FinishPrint("")

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC

	assert.Empty(out)
}

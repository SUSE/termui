package termui

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
)

const boldMagentaCode = "\x1b[1;35m"
const backToNormalCode = "\x1b[0m"

func TestColor_PrintColor(t *testing.T) {
	t.Parallel()

	ui, state := setup()
	one, three := "one", "three"
	two := 5
	colors := []color.Attribute{color.Bold, color.FgMagenta}

	ui.PrintColor(colors, one, two, three)
	expected := fmt.Sprint(one, two, three)
	expected = boldMagentaCode + expected + backToNormalCode
	if got := state.out.String(); got != expected {
		t.Errorf("output was wrong: %q, expected: %q", got, expected)
	}
}

func TestColor_PrintfColor(t *testing.T) {
	t.Parallel()

	ui, state := setup()
	format, value := "test %s printf", "the"
	colors := []color.Attribute{color.Bold, color.FgMagenta}

	ui.PrintfColor(colors, format, value)
	expected := fmt.Sprintf(format, value)
	expected = boldMagentaCode + expected + backToNormalCode
	if got := state.out.String(); got != expected {
		t.Errorf("output was wrong: %q, expected: %q", got, expected)
	}
}

func TestColor_PrintlnColor(t *testing.T) {
	t.Parallel()

	ui, state := setup()
	one, three := "one", "three"
	two := 5
	colors := []color.Attribute{color.Bold, color.FgMagenta}

	ui.PrintlnColor(colors, one, two, three)
	expected := fmt.Sprintln(one, two, three)
	expected = boldMagentaCode + expected + backToNormalCode
	if got := state.out.String(); got != expected {
		t.Errorf("output was wrong: %q, expected: %q", got, expected)
	}
}

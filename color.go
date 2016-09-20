package termui

import "github.com/fatih/color"

// These mirror their equivalent functions in ui.go
// Unfortunately since Print is already variadic,
// we can't have an elegant "...color.Attribute" argument and have to use a slice

func (u *UI) PrintColor(colors []color.Attribute, args ...interface{}) (int, error) {
	colorString := color.New(colors...).SprintFunc()(args)
	return u.Print(colorString)
}

func (u *UI) PrintfColor(colors []color.Attribute, format string, args ...interface{}) (int, error) {
	colorString := color.New(colors...).SprintFunc()(args)
	return u.Printf(format, colorString)
}

func (u *UI) PrintlnColor(colors []color.Attribute, args ...interface{}) (int, error) {
	colorString := color.New(colors...).SprintFunc()(args)
	return u.Println(colorString)
}

// Helper functions for simple color printing

func (u *UI) RedPrintln(args ...interface{}) (int, error) {
	colors := []color.Attribute{color.FgRed}
	return u.PrintlnColor(colors, args)
}

func (u *UI) RedBoldPrintln(args ...interface{}) (int, error) {
	colors := []color.Attribute{color.FgRed, color.Bold}
	return u.PrintlnColor(colors, args)
}

func (u *UI) YellowPrintln(args ...interface{}) (int, error) {
	colors := []color.Attribute{color.FgYellow}
	return u.PrintlnColor(colors, args)
}

func (u *UI) YellowBoldPrintln(args ...interface{}) (int, error) {
	colors := []color.Attribute{color.FgYellow, color.Bold}
	return u.PrintlnColor(colors, args)
}

func (u *UI) GreenPrintln(args ...interface{}) (int, error) {
	colors := []color.Attribute{color.FgGreen}
	return u.PrintlnColor(colors, args)
}

func (u *UI) GreenBoldPrintln(args ...interface{}) (int, error) {
	colors := []color.Attribute{color.FgGreen, color.Bold}
	return u.PrintlnColor(colors, args)
}

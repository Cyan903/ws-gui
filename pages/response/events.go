package response

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

var resText = widget.NewMultiLineEntry()
var ClearOnConnect = false
var logFormat = map[string]string{
	"hr":    "---------------------------------------------------------------------------",
	"log":   "[MSG]",
	"sent":  "[YOU]",
	"error": "[FATAL]",
}

func fmtText(s string) {
	resText.Text += s + "\n"
	resText.Refresh()
}

func Clear() {
	resText.Text = ""
	resText.Refresh()
}

func Connection(c string, u string) {
	if ClearOnConnect {
		Clear()
	} else {
		fmtText(logFormat["hr"])
	}

	fmtText(fmt.Sprintf("> %s: %s", c, u))
}

func Print(t, m string) {
	fmtText(fmt.Sprintf("%s | %s", logFormat[t], m))
}

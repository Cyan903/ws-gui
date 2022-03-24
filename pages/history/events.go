package history

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type tblHistory struct {
	data [][4]string
	tbl  *widget.Table
}

// Add syntax:
// {"", "12:45am", "wss://", "example.com"}
func (t *tblHistory) Add(d [4]string) {
	t.data = append(t.data, d)

	if t.tbl != nil {
		t.tbl.Refresh()
	}
}

func (t *tblHistory) setDefault() {
	t.Add([4]string{"#", "Time", "ws/wss", "Host"})
}

func (t *tblHistory) toTable() {
	t.tbl = widget.NewTable(
		func() (int, int) {
			return len(t.data), len(t.data[0])
		},
		func() fyne.CanvasObject {
			// Dummy text for width
			return widget.NewLabel("example.com")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Row == 0 {
				o.(*widget.Label).SetText(t.data[i.Row][i.Col])
				o.(*widget.Label).TextStyle.Bold = true
			}

			if i.Col == 0 && i.Row != 0 {
				o.(*widget.Label).SetText(fmt.Sprintf("#%d", i.Row))
				return
			}

			o.(*widget.Label).SetText(t.data[i.Row][i.Col])
		},
	)
}

func (t *tblHistory) Clear() {
	t.data = [][4]string{}
	t.setDefault()
}

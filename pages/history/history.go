package history

import (
	"github.com/Cyan903/ws-gui/gui"
)

var HistoryData = tblHistory{}

func History() gui.Page {
	history := gui.CreatePage("Connection History")

	HistoryData.setDefault()
	HistoryData.toTable()
	history.AddItem(HistoryData.tbl)

	return history
}

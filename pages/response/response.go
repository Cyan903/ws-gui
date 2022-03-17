package response

import (
	"github.com/Cyan903/ws-gui/gui"
)

func Response() gui.Page {
	response := gui.CreatePage("Response")
	response.AddItem(resText)

	return response
}

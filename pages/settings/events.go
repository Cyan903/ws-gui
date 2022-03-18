package settings

import (
	"github.com/Cyan903/ws-gui/pages/history"
	"github.com/Cyan903/ws-gui/pages/response"
)

func clearDisconnectFn(n bool) {
	response.ClearOnConnect = true
}

func clearConsole() {
	response.Clear()
}

func clearConnectionHistory() {
	history.HistoryData.Clear()
}

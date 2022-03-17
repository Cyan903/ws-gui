package settings

import "github.com/Cyan903/ws-gui/pages/response"

func clearDisconnectFn(n bool) {
	response.ClearOnConnect = true
}

func clearConsole() {
	response.Clear()
}

func clearConnectionHistory() {
	panic("unimplemented")
}

func checkUpdates() {
	panic("unimplemented")
}

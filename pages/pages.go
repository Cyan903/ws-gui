package pages

import (
	"github.com/Cyan903/ws-gui/gui"
	"github.com/Cyan903/ws-gui/pages/history"
	"github.com/Cyan903/ws-gui/pages/request"
	"github.com/Cyan903/ws-gui/pages/response"
	"github.com/Cyan903/ws-gui/pages/settings"
)

type pageFunc func() gui.Page

var Pages map[string]pageFunc = map[string]pageFunc{
	"history":  history.History,
	"request":  request.Request,
	"response": response.Response,
	"settings": settings.Settings,
}

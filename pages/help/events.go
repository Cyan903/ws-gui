package help

import (
	"io"
	"log"
	"net/http"
	"os"
)

var helpURL = "https://github.com/Cyan903/ws-gui/blob/main/static/help.md"
var helpRaw = "https://github.com/Cyan903/ws-gui/raw/main/static/help.md"

func readHelp() string {
	md, err := os.ReadFile("./static/help.md")

	if err != nil {
		log.Println("[ERROR] could not find help.md, loading default.")
		return ""
	}

	return string(md)
}

func downloadHelp() {
	url := helpRaw
	res, err := http.Get(url)

	if err != nil {
		log.Println("[ERROR] could not download help markdown.")
		return
	}

	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		os.Mkdir("static", os.ModePerm)
	}

	defer res.Body.Close()
	out, err := os.Create("static/help.md")

	if err != nil {
		log.Println("[ERROR] could not create file.")
		return
	}

	defer out.Close()
	io.Copy(out, res.Body)
}

package help

import (
	"log"
	"os"
)

func readHelp() string {
	md, err := os.ReadFile("./static/help.md")

	if err != nil {
		log.Println("[ERROR] could not find help.md, loading default.")
		return ""
	}

	return string(md)
}

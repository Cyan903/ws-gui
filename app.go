package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
)

func setIcon(w *fyne.Window) {
	icon, err := os.ReadFile("./static/logo.png")

	if err != nil {
		log.Println("[ERROR] could not find icon, using default icon.")
		return
	}

	(*w).SetIcon(fyne.NewStaticResource("logo", icon))
}

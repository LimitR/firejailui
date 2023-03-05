package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	flag := false
	_, err := exec.Command("firejail", "--version").Output()
	if err != nil {
		log.Fatal(err)
	}
	gtk.Init(nil)
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("Firejail UI")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)
	filesBox, _ := gtk.FileChooserButtonNew("File open", gtk.FILE_CHOOSER_ACTION_OPEN)
	filesBox.Connect("selection-changed", func() {
		go func() {
			name := filesBox.FileChooser.GetFilename()
			parse := strings.Split(name, ".")
			if len(parse) > 1 {
				if parse[1] == "png" || parse[1] == "jpg" {
					cmd := exec.Command("firejail", "xdg-open", filesBox.FileChooser.GetFilename())
					if err := cmd.Run(); err != nil {
						log.Fatal(err)
					}
				}
			} else {
				if flag {
					cmd := exec.Command("firejail", filesBox.FileChooser.GetFilename())
					if err := cmd.Run(); err != nil {
						log.Fatal(err)
					}
				} else {
					cmd := exec.Command("firejail", "--net=none", filesBox.FileChooser.GetFilename())
					if err := cmd.Run(); err != nil {
						log.Fatal(err)
					}
				}

			}
		}()
	})
	btn, _ := gtk.CheckButtonNewWithLabel("Internet ON/OFF")
	btn.Connect("toggled", func() {
		flag = !flag
	})
	// win.Add(filesBox)
	box.Add(filesBox)
	box.SetChildPacking(&filesBox.Widget, true, true, 0, 0)
	btns, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 3)
	btns.Add(btn)
	btns.SetChildPacking(&btn.Widget, false, false, 0, 0)
	box.Add(btns)
	win.Add(box)
	win.SetDefaultSize(800, 600)
	win.SetIconFromFile("./logo.png")
	win.ShowAll()

	gtk.Main()
}

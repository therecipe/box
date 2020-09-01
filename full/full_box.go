package main

import (
	"os"

	_ "github.com/therecipe/qt/interop"

	"github.com/therecipe/qt/core"
	_ "github.com/therecipe/qt/gui"
	_ "github.com/therecipe/qt/multimedia"
	_ "github.com/therecipe/qt/quick"
	_ "github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
)

func main() {

	// enable high dpi scaling
	// useful for devices with high pixel density displays
	// such as smartphones, retina displays, ...
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	widgets.NewQApplication(len(os.Args), os.Args)

	widgets.QApplication_Exec()
}

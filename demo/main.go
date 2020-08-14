package main

import (
	"os"
	"path/filepath"

	"github.com/therecipe/qt/interop/gow" //TODO:

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
)

func main() {

	p, r := gow.InitProcess() //TODO: app := widgets.NewQApplication(len(os.Args), os.Args)

	//
	//https://github.com/therecipe/examples/tree/master/basic/widgets
	//

	// create a window
	// with a minimum size of 250*200
	// and sets the title to "Hello Widgets Example"
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("Hello Widgets Example")

	// create a regular widget
	// give it a QVBoxLayout
	// and make it the central widget of the window
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	// create a line edit
	// with a custom placeholder text
	// and add it to the central widgets layout
	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Write something ...")
	widget.Layout().AddWidget(input)

	// create a button
	// connect the clicked signal
	// and add it to the central widgets layout
	button := widgets.NewQPushButton2("and click me!", nil)
	button.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	widget.Layout().AddWidget(button)

	// make the window visible
	window.Show()

	//
	//https://github.com/therecipe/examples/tree/master/basic/quick
	//

	// use the material style
	// the other inbuild styles are:
	// Default, Fusion, Imagine, Universal
	quickcontrols2.QQuickStyle_SetStyle("Material")

	// create the quick view
	// with a minimum size of 250*200
	// set the window title to "Hello QML/Quick Example"
	// and let the root item of the view resize itself to the size of the view automatically
	view := quick.NewQQuickView(nil)
	view.SetMinimumSize(core.NewQSize2(250, 200))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetTitle("Hello QML/Quick Example")

	// load the local qml file
	pwd, _ := os.Getwd()
	arg, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	for _, path := range []string{pwd, arg} {
		path = filepath.Join(path, "qml")
		if _, err := os.Stat(path); err == nil {
			view.SetSource(core.NewQUrl3(filepath.Join(path, "main.qml"), 0))
			break
		}
	}

	// make the view visible
	view.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	gow.Exec(p, r) //TODO: app.Exec()
}

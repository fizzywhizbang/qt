package main

import (
	"os"

	"github.com/fizzywhizbang/qt/core"
	"github.com/fizzywhizbang/qt/gui"
	"github.com/fizzywhizbang/qt/quick"

	_ "github.com/fizzywhizbang/qt/internal/examples/qml/extending/components/test_module/component"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.Engine().AddImportPath("qrc:/")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/app.qml", 0))

	view.Show()

	gui.QGuiApplication_Exec()
}

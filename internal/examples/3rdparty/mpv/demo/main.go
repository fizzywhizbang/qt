//source: https://github.com/mpv-player/mpv-examples/tree/master/libmpv/qml

package main

import (
	"os"

	"github.com/fizzywhizbang/qt/core"
	"github.com/fizzywhizbang/qt/gui"
	"github.com/fizzywhizbang/qt/quick"

	"github.com/fizzywhizbang/qt/internal/examples/3rdparty/mpv"
)

func main() {
	gui.NewQGuiApplication(len(os.Args), os.Args)

	mpv.InitMpv()

	view := quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0)) //qrc:/mpvtest/main.qml should work as well, since init.go is not patching out the embedding
	view.Show()

	gui.QGuiApplication_Exec()
}

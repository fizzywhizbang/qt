package main

import (
	"os"

	"github.com/fizzywhizbang/qt/widgets"

	"github.com/fizzywhizbang/qt/internal/examples/uitools/calculator/ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	ui.NewCalculatorForm(nil).Show()

	widgets.QApplication_Exec()
}

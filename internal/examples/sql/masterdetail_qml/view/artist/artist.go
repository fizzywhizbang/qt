// +build !qml

package artist

import (
	"github.com/fizzywhizbang/qt/core"
	"github.com/fizzywhizbang/qt/widgets"

	"github.com/fizzywhizbang/qt/internal/examples/sql/masterdetail_qml/controller"
)

type artistController struct {
	widgets.QGroupBox

	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"listModel"`

	//->controller
	_ func(row int) `signal:"changeArtist"`

	artistView *widgets.QComboBox
}

func (a *artistController) init() {

	a.artistView = widgets.NewQComboBox(nil)
	a.artistView.SetModel(controller.Instance.ArtistModel())

	layout := widgets.NewQGridLayout2()
	layout.AddWidget2(a.artistView, 0, 0, 0)
	a.SetLayout(layout)

	//

	//->controller
	a.artistView.ConnectCurrentIndexChanged(controller.Instance.ChangeArtist)
}

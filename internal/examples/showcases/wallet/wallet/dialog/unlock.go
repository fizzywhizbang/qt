package dialog

import (
	"github.com/fizzywhizbang/qt/core"

	_ "github.com/fizzywhizbang/qt/internal/examples/showcases/wallet/wallet/dialog/controller"
)

func init() { unlockTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "UnlockTemplate") }

type unlockTemplate struct {
	dialogTemplate

	_ func(string) *core.QVariant `slot:"unlock,->(controller.Controller)"`
}

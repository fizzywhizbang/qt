package runtime

import (
	"github.com/fizzywhizbang/qt/core"
	"github.com/fizzywhizbang/qt/qml"
)

func init() {
	qml.Z_newEngineHelper = func(p core.QObject_ITF) *core.QObject { return NewHelper(p).QObject_PTR() }
}

type helper struct {
	core.QObject

	_ func()                            `constructor:"init"`
	_ func(*qml.QJSValue) *qml.QJSValue `slot:"wrapperFunction"`
}

func (h *helper) init() { h.ConnectWrapperFunction(qml.Z_wrapperFunction) }

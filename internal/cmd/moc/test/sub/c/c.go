package c

import (
	"github.com/fizzywhizbang/qt/core"

	_ "github.com/fizzywhizbang/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoC struct{}
type StructSubMocC struct{ core.QObject }

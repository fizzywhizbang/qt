package b

import (
	"github.com/fizzywhizbang/qt/core"

	_ "github.com/fizzywhizbang/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoB struct{}
type StructSubMocB struct{ core.QObject }

package d

import (
	"github.com/fizzywhizbang/qt/core"

	_ "github.com/fizzywhizbang/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoD struct{}
type StructSubMocD struct{ core.QObject }

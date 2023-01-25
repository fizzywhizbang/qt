package subsubcustom

import (
	"github.com/fizzywhizbang/qt/gui"

	_ "github.com/fizzywhizbang/qt/internal/cmd/moc/test/sub/conf"
)

var SubSubTestStructInstance *SubSubTestStruct

type SubSubTestStruct struct {
	gui.QWindow

	_ func() `constructor:"Init"`

	_ func(string) `signal:"subPropertySignal"`
	_ func(string) `slot:"subPropertySlot"`

	_ string `property:"subsubProperty"`

	SubSubConstructorProperty int
}

func (s *SubSubTestStruct) Init() {
	SubSubTestStructInstance = s
	s.SubSubConstructorProperty++
}

package vesclient

import (
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"testing"
)

func testInit() {
	wrapper.SetCodeDescriptor(func(code int) string {
		return types.CodeType(code).String()
	})
}

func TestMain(m *testing.M) {
	testInit()
	m.Run()
}


package control

import "github.com/Myriad-Dreamin/go-ves/types"

var _ NSBClient = types.NSBClient(nil)
var _ types.NSBClient = NSBClient(nil)

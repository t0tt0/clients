package main

import (
	vesclient "github.com/Myriad-Dreamin/go-ves/lib/net/ves-client"
)

func init() {
	vesclient.Init()
}

func main() {
	vesclient.Main()
}

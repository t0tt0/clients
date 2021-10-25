package jsonrpcclient

import bytespool "github.com/HyperService-Consortium/go-ves/lib/basic/bytes-pool"

const (
	httpPrefix   = "http://"
	httpsPrefix  = "https://"
	maxBytesSize = 64 * 1024
)

var staticJsonRpcClient = &JsonRpcClient{bufferPool: bytespool.NewBytesPool(maxBytesSize)}

var fixedJsonHeader = map[string]string{
	"Content-Type": "application/json",
	"charset":      "UTF-8",
}

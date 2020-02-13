package wsrpc

import (
	"bytes"
	"encoding/binary"
	"github.com/Myriad-Dreamin/go-ves/types"
	"sync"

	"github.com/gogo/protobuf/proto"
)

const maxSize = 65536

// BufferPool Object
// accept temporary buffer object whose size is larger than maxBufferSize
// provide the steady source of temporary buffer.
type BufferPool struct {
	*sync.Pool
	maxBufferSize int
}

// NewBufferPool return a pointer of BufferPool
func NewBufferPool(maxBufferSize int) *BufferPool {
	return &BufferPool{Pool: &sync.Pool{
		New: func() interface{} {
			// buf :=
			// fmt.Printf("gee %p", buf)
			return bytes.NewBuffer(make([]byte, 0, maxBufferSize))
		},
	},
		maxBufferSize: maxBufferSize,
	}
}

// Put a buffer into this BufferPool. the buffer whose size is smaller than the
// bufferPool.maxBufferSize will be ignored
func (bp *BufferPool) Put(buf interface{}) {
	if buf, ok := buf.(*bytes.Buffer); ok {
		// fmt.Printf("puu %p", buf)
		if buf.Cap() >= bp.maxBufferSize {
			buf.Reset()
			bp.Pool.Put(buf)
		}
	}
}

// Serializer works for provide the environment of serializing protoBuf
// messages
type Serializer struct {
	bufferPool *BufferPool
}

var serial *Serializer

// NewSerializer return a pointer of Serializer
func NewSerializer(maxBufferSize int) *Serializer {
	return &Serializer{bufferPool: NewBufferPool(maxBufferSize)}
}

// Serialize concat the msg id and serialized msg
func (ser *Serializer) SerializeRaw(msgID MessageType, msg []byte) (*bytes.Buffer, error) {
	var buf = ser.bufferPool.Get().(*bytes.Buffer)
	buf.Reset()

	err := binary.Write(buf, types.WebSocketEndian, msgID)
	if err != nil {
		return nil, err
	}

	buf.Write(msg)
	return buf, nil
}

// Serialize concat the msg id and serialized msg
func (ser *Serializer) Serialize(msgID MessageType, msg proto.Message) (*bytes.Buffer, error) {

	b, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return ser.SerializeRaw(msgID, b)
}

func Deserialize(rawMessage []byte) (message []byte, messageID MessageType, err error) {
	var buf = bytes.NewBuffer(rawMessage)
	err = binary.Read(buf, types.WebSocketEndian, &messageID)
	message = buf.Bytes()
	return
}

// Put a buffer into its buffer pool
func (ser *Serializer) Put(buf *bytes.Buffer) {
	ser.bufferPool.Put(buf)
}

// GetDefaultSerializer gets the default serializer singleton
func GetDefaultSerializer() *Serializer {
	return serial
}

func init() {
	serial = NewSerializer(maxSize)
}

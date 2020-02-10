package uniquer

import (
	"crypto/md5"
	"unsafe"
)

// must be used in single-thread env
type Uniquer struct {
	s      map[string]bool
	hacker [8]byte
}

func MakeUniquer() *Uniquer {
	return &Uniquer{s: make(map[string]bool)}
}

func (c *Uniquer) Insert(a uint64, b []byte) bool {
	h := md5.New()
	*(*uint64)(unsafe.Pointer(&c.hacker[0])) = a
	h.Write(c.hacker[:])
	h.Write(b)
	var nb = h.Sum(nil)
	if _, ok := c.s[string(nb)]; ok {
		return false
	} else {
		c.s[string(nb)] = true
		return true
	}
}

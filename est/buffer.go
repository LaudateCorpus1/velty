package est

import "strconv"

type Buffer struct {
	buf      []byte
	index    int
	poolSize int
}

func (b *Buffer) AppendBytes(bs []byte) {
	bsLen := len(bs)
	if bsLen == 0 {
		return
	}

	b.growIfNeeded(bsLen)
	copy(b.buf[b.index:], bs)
	b.index += bsLen
}
func (b *Buffer) AppendByte(bs byte) {
	if b.index+1 >= len(b.buf) {
		newBuffer := make([]byte, len(b.buf)+b.poolSize)
		b.buf = append(b.buf, newBuffer...)
	}
	b.buf[b.index] = bs
	b.index++
}

func (b *Buffer) AppendInt(v int) {
	s := strconv.Itoa(v)
	b.AppendString(s)
}

func (b *Buffer) AppendBool(v bool) {
	s := strconv.FormatBool(v)
	b.AppendString(s)
}

func (b *Buffer) AppendFloat(v float64) {
	s := strconv.FormatFloat(v, 'f', -1, 64)
	b.AppendString(s)
}

func (b *Buffer) AppendString(s string) {
	sLen := len(s)
	if sLen == 0 {
		return
	}
	b.growIfNeeded(sLen)
	copy(b.buf[b.index:], s)
	b.index += sLen
}

func (b *Buffer) growIfNeeded(sLen int) {
	if sLen+b.index >= len(b.buf) {
		size := len(b.buf) + b.poolSize
		if size < sLen {
			size = sLen
		}
		newBuffer := make([]byte, size)
		b.buf = append(b.buf, newBuffer...)
	}
}

func (b *Buffer) Reset() {
	b.index = 0
}

func (b *Buffer) Bytes() []byte {
	return b.buf[:b.index]
}

func (b *Buffer) String() string {
	return string(b.buf[:b.index])
}

func NewBuffer(size int) *Buffer {
	return &Buffer{buf: make([]byte, size)}
}

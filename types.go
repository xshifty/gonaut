package gonaut

import "bytes"

type Map map[any]any

type String string

func (s String) Bytes() []byte {
	return bytes.NewBufferString(string(s)).Bytes()
}

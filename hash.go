package hash

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
)

type Type struct {
	data interface{}
}

func NewType(data interface{}) *Type {
	return &Type{data: data}
}

func (t Type) Bytes() []byte {
	switch t.data.(type) {
	case []byte:
		return t.data.([]byte)
	case string:
		return []byte(t.data.(string))
	}

	panic("not allowed type")
}

func (t Type) String() string {
	switch t.data.(type) {
	case []byte:
		return string(t.data.([]byte))
	case string:
		return t.data.(string)
	}

	panic("not allowed type")
}

func Md5(data interface{}) *Type {
	return doHash(crypto.MD5, data)
}

func Sha1(data interface{}) *Type {
	return doHash(crypto.SHA1, data)
}

func Sha256(data interface{}) *Type {
	return doHash(crypto.SHA256, data)
}

func doHash(hash crypto.Hash, data interface{}) *Type {
	t := NewType(data)
	object := hash.New()
	object.Write(t.Bytes())
	return NewType(object.Sum(nil))
}

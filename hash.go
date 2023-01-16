package hash

import "crypto"

func Md5(data []byte) []byte {
	return doHash(crypto.MD5, data)
}

func Sha1(data []byte) []byte {
	return doHash(crypto.SHA1, data)
}

func Sha256(data []byte) []byte {
	return doHash(crypto.SHA256, data)
}

func doHash(hash crypto.Hash, data []byte) []byte {
	object := hash.New()
	object.Write(data)
	return object.Sum(nil)
}

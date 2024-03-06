package base

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"github.com/zohu/zutils/zcpt"
	"io"
)

const cbckey = "QKcartfnQIopSymo"

func Decrypt(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	d, err := zcpt.AesDecryptCBC(_unGzip(b), []byte(cbckey))
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func Encrypt(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	d, err := zcpt.AesEncryptCBC([]byte(str), []byte(cbckey))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(_gzip(d)), nil
}
func _gzip(in []byte) []byte {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	_, _ = gz.Write(in)
	_ = gz.Flush()
	_ = gz.Close()
	return b.Bytes()
}

func _unGzip(out []byte) []byte {
	gz, _ := gzip.NewReader(bytes.NewBuffer(out))
	d, _ := io.ReadAll(gz)
	_ = gz.Close()
	return d
}

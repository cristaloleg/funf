package jsonblob

import (
	"encoding/base64"
	"encoding/json"
)

// define for the clarity
var b64 = base64.URLEncoding

type Blob []byte

func (b Blob) String() string {
	return string(b)
}

func (b Blob) MarshalJSON() ([]byte, error) {
	dst := make([]byte, b64.EncodedLen(len(b))+2) // +2 for quotes (JSON string)
	dst[0], dst[len(dst)-1] = '"', '"'
	b64.Encode(dst[1:], []byte(b))
	return dst, nil
}

func (b *Blob) UnmarshalJSON(p []byte) error {
	if len(p) < 3 {
		return nil
	}
	dst := make([]byte, b64.DecodedLen(len(p)-2))
	_, err := b64.Decode(dst, p[1:len(p)-1])
	if err != nil {
		return err
	}
	*b = Blob(dst)
	return nil
}

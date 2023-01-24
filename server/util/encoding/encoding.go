package encoding

import (
	"errors"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func AutoDecode(s []byte) ([]byte, error) {
	r, err := chardet.NewTextDetector().DetectBest(s)
	if err != nil {
		return nil, err
	}
	switch r.Charset {
	case "GB-18030":
		out, _ := simplifiedchinese.GBK.NewDecoder().Bytes(s)
		return out, nil
	case "UTF8":
		return s, nil
	default:
		return s, errors.New("unknown charset")
	}
}

func ShouldAutoDecode(s []byte) []byte {
	out, err := AutoDecode(s)
	if err != nil {
		//panic(err)
		return s
	}
	return out
}

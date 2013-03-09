package iconv

import "errors"

// Decoder convert go []byte to string
type Decoder struct {
	converter *Converter
}


func NewDecoder(encoding string) (*Decoder, error) {
	converter, err := NewConverter(encoding, "UTF-8")
	if err != nil {
		return nil, err
	}
	return &Decoder{converter}, nil
}


func (decoder *Decoder) Decode(source []byte) (string, error) {
	inputLength := len(source)
	output := make([]byte, inputLength*3)
	read, _, err := decoder.converter.Convert(source, output)
	if err != nil {
		return string(output), err
	}
	if read == inputLength {
		return string(output), nil
	}
	return string(output), errors.New("source string too large")
}

func (decoder *Decoder) MustDecode(source []byte) string {
	output, err := decoder.Decode(source)
	if err != nil {
		panic(err)
	}
	return output
}

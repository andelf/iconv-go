package iconv

import "errors"

// Encoder convert go string to []byte
type Encoder struct {
	converter *Converter
}


func NewEncoder(encoding string) (*Encoder, error) {
	converter, err := NewConverter("UTF-8", encoding)
	if err != nil {
		return nil, err
	}
	return &Encoder{converter}, nil
}


func (encoder *Encoder) Encode(source string) ([]byte, error) {
	input := []byte(source)
	inputLength := len(input)
	output := make([]byte, inputLength*3)
	read, _, err := encoder.converter.Convert(input, output)
	if err != nil {
		return output, err
	}
	if read == inputLength {
		return output, nil
	}
	return output, errors.New("source string too large")
}

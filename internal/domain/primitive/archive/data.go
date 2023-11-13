package archive

import "github.com/CityBear3/satellite/internal/pkg/apperrs"

// max data size: 10MB
const MaxDataSize = 10e6

type DataSize struct {
	value int
}

func NewDataSize(value int) (DataSize, error) {
	if value <= 0 || value >= MaxDataSize {
		return DataSize{}, apperrs.InvalidFileSizeError
	}

	return DataSize{
		value: value,
	}, nil
}

func (d DataSize) Value() int {
	return d.value
}

type Data struct {
	Chunks []byte
	Size   DataSize
}

func NewData(chunks []byte) (Data, error) {
	size, err := NewDataSize(len(chunks))
	if err != nil {
		return Data{}, err
	}

	return Data{
		Chunks: chunks,
		Size:   size,
	}, nil
}

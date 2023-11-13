package primitive

import "github.com/oklog/ulid/v2"

type ID struct {
	value ulid.ULID
}

func NewID() ID {
	return ID{value: ulid.Make()}
}

func ParseID(value string) (ID, error) {
	id, err := ulid.Parse(value)
	if err != nil {
		return ID{}, err
	}

	return ID{value: id}, nil
}

func (i ID) Value() ulid.ULID {
	return i.value
}

func (i ID) String() string {
	return i.value.String()
}

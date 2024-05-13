package rpc

import (
	"github.com/gofrs/uuid/v5"
)

func PackUuid(id *uuid.UUID) *Uuid {
	return &Uuid{Data: id.String()}
}

func UnpackUuid(from *Uuid) (uuid.UUID, error) {
	id, err := uuid.FromString(from.Data)
	return id, err
}

package format

import (
	"encoding/base64"
	"github.com/google/uuid"
)

func ConvertUUIDToKey(id uuid.UUID) string {
	return base64.RawURLEncoding.EncodeToString(id[:])
}

func ConvertKeyToUUID(key string) (uuid.UUID, error) {
	data, err := base64.RawURLEncoding.DecodeString(key)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.FromBytes(data)
}

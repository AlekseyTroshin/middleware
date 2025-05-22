package storage

import (
	storageMap "hexlet-auth/internal/storage"
)

type Storage struct {
	Data map[string]storageMap.User
}

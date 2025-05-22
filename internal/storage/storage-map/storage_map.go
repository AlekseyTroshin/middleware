package storage_map

import (
	"fmt"
	storageInternal "hexlet-auth/internal/storage"
	"hexlet-auth/storage"
)

type Storage struct {
	storage.Storage
}

type Store interface {
	Load(email string) (storageInternal.User, error)
	Save(name string, email string) (storageInternal.User, error)
	Delete(email string) error
}

func New() *Storage {
	return &Storage{
		Storage: storage.Storage{
			Data: make(map[string]storageInternal.User),
		},
	}
}

func (s *Storage) Load(email string) (storageInternal.User, error) {
	user, ok := s.Data[email]
	if !ok {
		return storageInternal.User{}, fmt.Errorf("not find user")
	}

	return user, nil
}

func (s *Storage) Save(name string, email string) (storageInternal.User, error) {
	_, ok := s.Data[email]
	if ok {
		return storageInternal.User{}, fmt.Errorf("this user already exists")
	}

	user := storageInternal.User{
		Name:  name,
		Email: email,
	}

	s.Data[email] = user

	return user, nil
}

func (s *Storage) Delete(email string) (storageInternal.User, error) {
	user, err := s.Load(email)
	if err != nil {
		return storageInternal.User{}, err
	}

	delete(s.Data, email)

	return user, nil
}

package main

import (
	"errors"
	"sync"
)

// TODO #1: implement in-memory user store

var userStore = NewUserStore()

func NewUserStore() *UserStore {
	return &UserStore{data: make(map[string]UserInfo)}
}

type UserStore struct {
	mu   sync.Mutex
	data map[string]UserInfo
}

func (u *UserStore) Save(info UserInfo) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	// Save the user information in the data map
	u.data[info.Username] = info

	return nil
}

func (u *UserStore) Get(username string) (UserInfo, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	// Retrieve the user information from the data map
	user, found := u.data[username]
	if !found {
		return UserInfo{}, ErrUserNotFound
	}

	return user, nil
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

var ErrUserNotFound = errors.New("user not found")
var ErrUserExisted = errors.New("user existed")

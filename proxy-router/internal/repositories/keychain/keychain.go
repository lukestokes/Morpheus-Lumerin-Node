package keychain

import (
	"github.com/zalando/go-keyring"
)

const SERVICE_NAME = "morpheus-proxy-router"

type Keychain struct {
	service string
}

func NewKeychain() *Keychain {
	return &Keychain{}
}

func (k *Keychain) Get(key string) (string, error) {
	return keyring.Get(SERVICE_NAME, key)
}

func (k *Keychain) Insert(key string, value string) error {
	return keyring.Set(SERVICE_NAME, key, value)
}

func (k *Keychain) Upsert(key string, value string) error {
	return keyring.Set(SERVICE_NAME, key, value)
}

func (k *Keychain) Delete(key string) error {
	return keyring.Delete(SERVICE_NAME, key)
}

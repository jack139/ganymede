package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UsersKeyPrefix is the prefix to retrieve all Users
	UsersKeyPrefix = "Users/value/"
	UsersLinkPrefix = "Users/link/"
)

// UsersKey returns the store key to retrieve a Users from the index fields
func UsersKey(
	chainAddr string,
) []byte {
	var key []byte

	chainAddrBytes := []byte(chainAddr)
	key = append(key, chainAddrBytes...)
	key = append(key, []byte("/")...)

	return key
}

// 生成 链表头 key
func UsersStatusLinkKey(status string) []byte {
	var key []byte

	key = append(key, []byte("@@LINK:status:")...)

	statusBytes := []byte(status)
	key = append(key, statusBytes...)

	return key
}

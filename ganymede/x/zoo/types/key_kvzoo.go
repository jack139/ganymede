package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// KvzooKeyPrefix is the prefix to retrieve all Kvzoo
	KvzooKeyPrefix = "Kvzoo/value/"
	KvzooLinkPrefix = "Kvzoo/link/"
)

// KvzooKey returns the store key to retrieve a Kvzoo from the index fields
func KvzooKey(
	owner string,
	zooKey string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	zooKeyBytes := []byte(zooKey)
	key = append(key, zooKeyBytes...)
	key = append(key, []byte("/")...)

	return key
}

// 生成 链表头 key
func KvzooOwnerLinkKey(owner string) []byte {
	var key []byte

	key = append(key, []byte("@@LINK:owner:")...)

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)

	return key
}

package types

const (
	// ModuleName defines the module name
	ModuleName = "postoffice"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_postoffice"

	// Version defines the current version the IBC module supports
	Version = "postoffice-1"

	// PortID is the default port id that module binds to
	PortID = "postoffice"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("postoffice-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PostKey      = "Post/value/"
	PostCountKey = "Post/count/"
)

const (
	SentPostKey      = "SentPost/value/"
	SentPostCountKey = "SentPost/count/"
)

const (
	TimedoutPostKey      = "TimedoutPost/value/"
	TimedoutPostCountKey = "TimedoutPost/count/"
)

const (
	PostSenderLinkPrefix   = "Post/link/sender/"
	PostReceiverLinkPrefix = "Post/link/receiver/"

	SentSenderLinkPrefix   = "Sent/link/sender/"
	SentReceiverLinkPrefix = "Sent/link/receiver/"

	TimeoutSenderLinkPrefix   = "Timeout/link/sender/"
	TimeoutReceiverLinkPrefix = "Timeout/link/receiver/"
)

// 生成 链表头 key
func PostofficeLinkKey(cate string, link string) []byte {
	var key []byte

	key = append(key, []byte("@@LINK:")...)

	cateBytes := []byte(cate)
	key = append(key, cateBytes...)
	key = append(key, []byte(":")...)

	linkBytes := []byte(link)
	key = append(key, linkBytes...)

	return key
}

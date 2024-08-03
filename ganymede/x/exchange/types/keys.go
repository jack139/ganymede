package types

const (
	// ModuleName defines the module name
	ModuleName = "exchange"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_exchange"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	AskKey      = "Ask/value/"
	AskCountKey = "Ask/count/"
)

const (
	ReplyKey      = "Reply/value/"
	ReplyCountKey = "Reply/count/"
)

const (
	AskSenderLinkPrefix  = "Ask/link/sender/"
	AskReplierLinkPrefix = "Ask/link/replier/"
	AskStatusLinkPrefix  = "Ask/link/status/"

	ReplySenderLinkPrefix  = "Reply/link/sender/"
	ReplyReplierLinkPrefix = "Reply/link/replier/"
)


// 生成 链表头 key
func ExchangeLinkKey(cate string, link string) []byte {
	var key []byte

	key = append(key, []byte("@@LINK:")...)

	cateBytes := []byte(cate)
	key = append(key, cateBytes...)
	key = append(key, []byte(":")...)

	linkBytes := []byte(link)
	key = append(key, linkBytes...)

	return key
}

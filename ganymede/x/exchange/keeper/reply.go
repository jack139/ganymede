package keeper

import (
	"encoding/binary"
	"log"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

// GetReplyCount get the total number of reply
func (k Keeper) GetReplyCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ReplyCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetReplyCount set the total number of reply
func (k Keeper) SetReplyCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ReplyCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendReply appends a reply in the store with a new id and update the count
func (k Keeper) AppendReply(
	ctx sdk.Context,
	reply types.Reply,
) uint64 {
	// Create the reply
	count := k.GetReplyCount(ctx)

	// Set the ID of the appended value
	reply.Id = count
	idByteKey := GetReplyIDBytes(reply.Id)

	// 加入链

	// 检索链表头信息, 2个链表： sender, replier
	senderKey := types.ExchangeLinkKey("replySender", reply.Sender) // 链表头 key
	storeSenderLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReplySenderLinkPrefix))
	b := storeSenderLink.Get(senderKey)
	if b == nil { // 未找到，说明是新的值
		reply.LinkSender = "@@LINK:$" 
	} else {
		reply.LinkSender = string(b)
	}
	storeSenderLink.Set(senderKey, idByteKey)  // 保存表头数据

	replierKey := types.ExchangeLinkKey("replyReplier", reply.Replier) // 链表头 key
	storeReplierLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReplyReplierLinkPrefix))
	b = storeReplierLink.Get(replierKey)
	if b == nil { // 未找到，说明是新的值
		reply.LinkReplier = "@@LINK:$" 
	} else {
		reply.LinkReplier = string(b)
	}
	storeReplierLink.Set(replierKey, idByteKey)  // 保存表头数据

	// 保存数据
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReplyKey))
	appendedValue := k.cdc.MustMarshal(&reply)
	store.Set(idByteKey, appendedValue)

	// Update reply count
	k.SetReplyCount(ctx, count+1)

	return count
}

// SetReply set a specific reply in the store
func (k Keeper) SetReply(ctx sdk.Context, reply types.Reply) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReplyKey))
	b := k.cdc.MustMarshal(&reply)
	store.Set(GetReplyIDBytes(reply.Id), b)
}

// GetReply returns a reply from its id
func (k Keeper) GetReply(ctx sdk.Context, id uint64) (val types.Reply, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReplyKey))
	b := store.Get(GetReplyIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetReply returns a reply from its id
func (k Keeper) GetReplyByKey(ctx sdk.Context, key []byte) (val types.Reply, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReplyKey))
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveReply removes a reply from the store
func (k Keeper) RemoveReply(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReplyKey))
	store.Delete(GetReplyIDBytes(id))
}

// GetAllReply returns all reply
func (k Keeper) GetAllReply(ctx sdk.Context) (list []types.Reply) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReplyKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Reply
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetReplyIDBytes returns the byte representation of the ID
func GetReplyIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetReplyIDFromBytes returns ID in uint64 format from a byte array
func GetReplyIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

// GetAllUsers returns all users
func (k Keeper) GetAllReplyLinks(ctx sdk.Context) () {
	storeLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReplySenderLinkPrefix))
	iterator := sdk.KVStorePrefixIterator(storeLink, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		log.Println("Reply Sender Link:", string(iterator.Key()), "-->", iterator.Value())
	}

	storeLink = prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReplyReplierLinkPrefix))
	iterator = sdk.KVStorePrefixIterator(storeLink, []byte{})

	for ; iterator.Valid(); iterator.Next() {
		log.Println("Reply Replier Link:", string(iterator.Key()), "-->", iterator.Value())
	}

	return
}

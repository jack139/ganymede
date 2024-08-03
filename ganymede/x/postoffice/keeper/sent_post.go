package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/ganymede/ganymede/x/postoffice/types"
)

// GetSentPostCount get the total number of sentPost
func (k Keeper) GetSentPostCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SentPostCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSentPostCount set the total number of sentPost
func (k Keeper) SetSentPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SentPostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSentPost appends a sentPost in the store with a new id and update the count
func (k Keeper) AppendSentPost(
	ctx sdk.Context,
	sentPost types.SentPost,
) uint64 {
	// Create the sentPost
	count := k.GetSentPostCount(ctx)

	// Set the ID of the appended value
	sentPost.Id = count
	idByteKey := GetSentPostIDBytes(sentPost.Id)

	// 检索链表头信息, 2个链表： sender, receiver
	senderKey := types.PostofficeLinkKey("sentSender", sentPost.Sender) // 链表头 key
	storeSenderLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentSenderLinkPrefix))
	b := storeSenderLink.Get(senderKey)
	if b == nil { // 未找到，说明是新的值
		sentPost.LinkSender = "@@LINK:$" 
	} else {
		sentPost.LinkSender = string(b)
	}
	storeSenderLink.Set(senderKey, idByteKey)  // 保存表头数据

	receiverKey := types.PostofficeLinkKey("sentReceiver", sentPost.Receiver) // 链表头 key
	storeReceiverLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentReceiverLinkPrefix))
	b = storeReceiverLink.Get(receiverKey)
	if b == nil { // 未找到，说明是新的值
		sentPost.LinkReceiver = "@@LINK:$" 
	} else {
		sentPost.LinkReceiver = string(b)
	}
	storeReceiverLink.Set(receiverKey, idByteKey)  // 保存表头数据

	// 保存数据
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentPostKey))
	appendedValue := k.cdc.MustMarshal(&sentPost)
	store.Set(idByteKey, appendedValue)

	// Update sentPost count
	k.SetSentPostCount(ctx, count+1)

	return count
}

// SetSentPost set a specific sentPost in the store
func (k Keeper) SetSentPost(ctx sdk.Context, sentPost types.SentPost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentPostKey))
	b := k.cdc.MustMarshal(&sentPost)
	store.Set(GetSentPostIDBytes(sentPost.Id), b)
}

// GetSentPost returns a sentPost from its id
func (k Keeper) GetSentPost(ctx sdk.Context, id uint64) (val types.SentPost, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentPostKey))
	b := store.Get(GetSentPostIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetSentPost returns a sentPost from its id
func (k Keeper) GetSentPostByKey(ctx sdk.Context, key []byte) (val types.SentPost, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentPostKey))
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSentPost removes a sentPost from the store
func (k Keeper) RemoveSentPost(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentPostKey))
	store.Delete(GetSentPostIDBytes(id))
}

// GetAllSentPost returns all sentPost
func (k Keeper) GetAllSentPost(ctx sdk.Context) (list []types.SentPost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentPostKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SentPost
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSentPostIDBytes returns the byte representation of the ID
func GetSentPostIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSentPostIDFromBytes returns ID in uint64 format from a byte array
func GetSentPostIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

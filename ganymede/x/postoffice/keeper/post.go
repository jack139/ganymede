package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/ganymede/ganymede/x/postoffice/types"
)

// GetPostCount get the total number of post
func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPostCount set the total number of post
func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPost appends a post in the store with a new id and update the count
func (k Keeper) AppendPost(
	ctx sdk.Context,
	post types.Post,
) uint64 {
	// Create the post
	count := k.GetPostCount(ctx)

	// Set the ID of the appended value
	post.Id = count
	idByteKey := GetPostIDBytes(post.Id)

	// 检索链表头信息, 2个链表： sender, receiver
	senderKey := types.PostofficeLinkKey("postSender", post.Sender) // 链表头 key
	storeSenderLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostSenderLinkPrefix))
	b := storeSenderLink.Get(senderKey)
	if b == nil { // 未找到，说明是新的值
		post.LinkSender = "@@LINK:$" 
	} else {
		post.LinkSender = string(b)
	}
	storeSenderLink.Set(senderKey, idByteKey)  // 保存表头数据

	receiverKey := types.PostofficeLinkKey("postReceiver", post.Receiver) // 链表头 key
	storeReceiverLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostReceiverLinkPrefix))
	b = storeReceiverLink.Get(receiverKey)
	if b == nil { // 未找到，说明是新的值
		post.LinkReceiver = "@@LINK:$" 
	} else {
		post.LinkReceiver = string(b)
	}
	storeReceiverLink.Set(receiverKey, idByteKey)  // 保存表头数据

	// 保存数据
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	appendedValue := k.cdc.MustMarshal(&post)
	store.Set(idByteKey, appendedValue)

	// Update post count
	k.SetPostCount(ctx, count+1)

	return count
}

// SetPost set a specific post in the store
func (k Keeper) SetPost(ctx sdk.Context, post types.Post) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	b := k.cdc.MustMarshal(&post)
	store.Set(GetPostIDBytes(post.Id), b)
}

// GetPost returns a post from its id
func (k Keeper) GetPost(ctx sdk.Context, id uint64) (val types.Post, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	b := store.Get(GetPostIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetPost returns a post from its id
func (k Keeper) GetPostByKey(ctx sdk.Context, key []byte) (val types.Post, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePost removes a post from the store
func (k Keeper) RemovePost(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	store.Delete(GetPostIDBytes(id))
}

// GetAllPost returns all post
func (k Keeper) GetAllPost(ctx sdk.Context) (list []types.Post) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Post
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPostIDBytes returns the byte representation of the ID
func GetPostIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPostIDFromBytes returns ID in uint64 format from a byte array
func GetPostIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

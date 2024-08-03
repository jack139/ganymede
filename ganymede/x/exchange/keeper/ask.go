package keeper

import (
	"log"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

// GetAskCount get the total number of ask
func (k Keeper) GetAskCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AskCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAskCount set the total number of ask
func (k Keeper) SetAskCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AskCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAsk appends a ask in the store with a new id and update the count
func (k Keeper) AppendAsk(
	ctx sdk.Context,
	ask types.Ask,
) uint64 {
	// Create the ask
	count := k.GetAskCount(ctx)

	// Set the ID of the appended value
	ask.Id = count
	idByteKey := GetAskIDBytes(ask.Id)

	// 加入链

	// 检索链表头信息, 3个链表： sender, replier
	senderKey := types.ExchangeLinkKey("sender", ask.Sender) // 链表头 key
	storeSenderLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AskSenderLinkPrefix))
	b := storeSenderLink.Get(senderKey)
	if b == nil { // 未找到，说明是新的值
		ask.LinkSender = "@@LINK:$" 
	} else {
		ask.LinkSender = string(b)
	}
	storeSenderLink.Set(senderKey, idByteKey)  // 保存表头数据

	replierKey := types.ExchangeLinkKey("replier", ask.Replier) // 链表头 key
	storeReplierLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AskReplierLinkPrefix))
	b = storeReplierLink.Get(replierKey)
	if b == nil { // 未找到，说明是新的值
		ask.LinkReplier = "@@LINK:$" 
	} else {
		ask.LinkReplier = string(b)
	}
	storeReplierLink.Set(replierKey, idByteKey)  // 保存表头数据

	// 保存数据
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AskKey))
	appendedValue := k.cdc.MustMarshal(&ask)
	store.Set(idByteKey, appendedValue)

	// Update ask count
	k.SetAskCount(ctx, count+1)

	return count
}

// SetAsk set a specific ask in the store
func (k Keeper) SetAsk(ctx sdk.Context, ask types.Ask) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AskKey))
	b := k.cdc.MustMarshal(&ask)
	store.Set(GetAskIDBytes(ask.Id), b)
}

// GetAsk returns a ask from its id
func (k Keeper) GetAsk(ctx sdk.Context, id uint64) (val types.Ask, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AskKey))
	b := store.Get(GetAskIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAsk returns a ask from its id
func (k Keeper) GetAskByKey(ctx sdk.Context, key []byte) (val types.Ask, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AskKey))
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAsk removes a ask from the store
func (k Keeper) RemoveAsk(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AskKey))
	store.Delete(GetAskIDBytes(id))
}

// GetAllAsk returns all ask
func (k Keeper) GetAllAsk(ctx sdk.Context) (list []types.Ask) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AskKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Ask
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAskIDBytes returns the byte representation of the ID
func GetAskIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAskIDFromBytes returns ID in uint64 format from a byte array
func GetAskIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

// GetAllUsers returns all users
func (k Keeper) GetAllAskLinks(ctx sdk.Context) () {
	storeLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AskSenderLinkPrefix))
	iterator := sdk.KVStorePrefixIterator(storeLink, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		log.Println("Ask Sender Link:", string(iterator.Key()), "-->", iterator.Value())
	}

	storeLink = prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AskReplierLinkPrefix))
	iterator = sdk.KVStorePrefixIterator(storeLink, []byte{})

	for ; iterator.Valid(); iterator.Next() {
		log.Println("Ask Replier Link:", string(iterator.Key()), "-->", iterator.Value())
	}

	return
}


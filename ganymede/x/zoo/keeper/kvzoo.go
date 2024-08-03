package keeper

import (
	"log"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/ganymede/ganymede/x/zoo/types"
)

// SetKvzoo set a specific kvzoo in the store from its index
func (k Keeper) SetKvzoo(ctx sdk.Context, kvzoo types.Kvzoo) {
	kvKey := types.KvzooKey(kvzoo.Owner, kvzoo.ZooKey)
	linkKey := types.KvzooOwnerLinkKey(kvzoo.Owner) // 链表头 key

	// 检索链表头信息
	storeLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KvzooLinkPrefix))
	b := storeLink.Get(linkKey)
	if b == nil { // 未找到，说明是新的owner值
		kvzoo.LinkOwner = "@@LINK:$" 
	} else {
		kvzoo.LinkOwner = string(b)
	}
	storeLink.Set(linkKey, kvKey)  // 保存表头数据

	// 保存数据
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KvzooKeyPrefix))
	b = k.cdc.MustMarshal(&kvzoo)
	store.Set(kvKey, b)
}

// SetKvzoo set a specific kvzoo in the store from its index
func (k Keeper) SetKvzooUpdate(ctx sdk.Context, kvzoo types.Kvzoo, oldOwner string, oldLinkOwner string) {
	// 修改了owner
	if kvzoo.Owner != oldOwner {
		// API 不支持修改owner，因此不进行修改
		kvzoo.Owner = oldOwner
	}

	// 设置 
	kvzoo.LinkOwner = oldLinkOwner

	// 保存数据
	kvKey := types.KvzooKey(kvzoo.Owner, kvzoo.ZooKey)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KvzooKeyPrefix))
	b := k.cdc.MustMarshal(&kvzoo)
	store.Set(kvKey, b)
}

// GetKvzoo returns a kvzoo from its index
func (k Keeper) GetKvzoo(
	ctx sdk.Context,
	owner string,
	zooKey string,

) (val types.Kvzoo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KvzooKeyPrefix))

	b := store.Get(types.KvzooKey(
		owner,
		zooKey,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetKvzoo returns a kvzoo from its index, Key is given
func (k Keeper) GetKvzooByKey(
	ctx sdk.Context,
	kvKey []byte,
) (val types.Kvzoo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KvzooKeyPrefix))

	b := store.Get(kvKey)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKvzoo removes a kvzoo from the store
func (k Keeper) RemoveKvzoo(
	ctx sdk.Context,
	owner string,
	zooKey string,
	linkOwner string,
) {
	kvKey := types.KvzooKey(owner, zooKey)
	linkKey := types.KvzooOwnerLinkKey(owner)

	storeLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KvzooLinkPrefix))
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KvzooKeyPrefix))

	// 从owner链中删除	
	old := storeLink.Get(linkKey)
	if string(old) == string(kvKey) { // 在表头
		storeLink.Set(linkKey, []byte(linkOwner))  // 保存表头数据
	} else {
		// 不在表头，要找到它的父节点，修改link信息
		thisKvKey := old
		for (string(thisKvKey) != "@@LINK:$") {
			u, found := k.GetKvzooByKey(ctx, thisKvKey)
			if !found {
				log.Println("!ERROR: RemoveKvzoo() not found thisKvKey: ", string(thisKvKey))
				break
			}
			if u.LinkOwner == string(kvKey) { // 找到父节点， 修改链接信息
				u.LinkOwner = linkOwner
				b := k.cdc.MustMarshal(&u)
				store.Set(thisKvKey, b)
				break
			}
			thisKvKey = []byte(u.LinkOwner)
		}
	}

	// 删除kv数据
	store.Delete(kvKey)
}

// GetAllKvzoo returns all kvzoo
func (k Keeper) GetAllKvzoo(ctx sdk.Context) (list []types.Kvzoo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KvzooKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Kvzoo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllUsers returns all users
func (k Keeper) GetAllLinks(ctx sdk.Context) () {
	storeLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KvzooLinkPrefix))
	iterator := sdk.KVStorePrefixIterator(storeLink, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		log.Println("KV Link:", string(iterator.Key()), "-->", string(iterator.Value()))
	}

	return
}

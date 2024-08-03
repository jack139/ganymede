package keeper

import (
	"log"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/ganymede/ganymede/x/ganymede/types"
)

// SetUsers set a specific users in the store from its index
func (k Keeper) SetUsers(ctx sdk.Context, users types.Users) {
	kvKey := types.UsersKey(users.ChainAddr)
	linkKey := types.UsersStatusLinkKey(users.Status) // 链表头 key

	// 检索链表头信息
	storeLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsersLinkPrefix))
	b := storeLink.Get(linkKey)
	if b == nil { // 未找到，说明是新的owner值
		users.LinkStatus = "@@LINK:$"
	} else {
		users.LinkStatus = string(b)
	}
	storeLink.Set(linkKey, kvKey)  // 保存表头数据

	// 保存数据
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsersKeyPrefix))
	b = k.cdc.MustMarshal(&users)
	store.Set(kvKey, b)
}

// SetUsers set a specific users in the store from its index  --- 修改
func (k Keeper) SetUsersUpdate(ctx sdk.Context, users types.Users, oldStatus string, oldLinkStatus string) {	
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsersKeyPrefix))
	kvKey := types.UsersKey(users.ChainAddr)

	// 检查 status 是否有修改
	if oldStatus != users.Status { // 需要修改链表
		storeLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsersLinkPrefix))
		linkKey := types.UsersStatusLinkKey(users.Status) // 链表头 key

		// 从旧status链中删除
		oldLinkKey := types.UsersStatusLinkKey(oldStatus)
		old := storeLink.Get(oldLinkKey)
		if string(old) == string(kvKey) { // 在表头
			storeLink.Set(oldLinkKey, []byte(oldLinkStatus))  // 保存表头数据
		} else {
			// 不在表头，要找到它的父节点，修改link信息
			thisKvKey := old
			for (string(thisKvKey) != "@@LINK:$") {
				u, found := k.GetUsersByKey(ctx, thisKvKey)
				if !found {
					log.Println("!ERROR: thisKvKey SetUsersUpdate()", string(thisKvKey))
					break
				}
				if u.LinkStatus == string(kvKey) { // 找到父节点， 修改链接信息
					u.LinkStatus = oldLinkStatus
					b := k.cdc.MustMarshal(&u)
					store.Set(thisKvKey, b)
					break
				}
				thisKvKey = []byte(u.LinkStatus)
			}
		}

		// 加入新status的链表
		new := storeLink.Get(linkKey)
		if new == nil { // 未找到，说明是新的owner值
			users.LinkStatus = "@@LINK:$"
		} else {
			users.LinkStatus = string(new)
		}
		storeLink.Set(linkKey, kvKey)  // 保存表头数据
	}

	// 保存数据
	b := k.cdc.MustMarshal(&users)
	store.Set(kvKey, b)
}

// GetUsers returns a users from its index
func (k Keeper) GetUsers(
	ctx sdk.Context,
	chainAddr string,

) (val types.Users, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsersKeyPrefix))

	b := store.Get(types.UsersKey(
		chainAddr,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetUsers returns a users from its index by KVkey
func (k Keeper) GetUsersByKey(
	ctx sdk.Context,
	kvKey []byte,
) (val types.Users, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsersKeyPrefix))

	b := store.Get(kvKey)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUsers removes a users from the store
func (k Keeper) RemoveUsers(
	ctx sdk.Context,
	chainAddr string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsersKeyPrefix))
	store.Delete(types.UsersKey(
		chainAddr,
	))
}

// GetAllUsers returns all users
func (k Keeper) GetAllUsers(ctx sdk.Context) (list []types.Users) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsersKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Users
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllUsers returns all users
func (k Keeper) GetAllLinks(ctx sdk.Context) () {
	storeLink := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsersLinkPrefix))
	iterator := sdk.KVStorePrefixIterator(storeLink, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		log.Println("Users Link:", string(iterator.Key()), "-->", string(iterator.Value()))
	}

	return
}

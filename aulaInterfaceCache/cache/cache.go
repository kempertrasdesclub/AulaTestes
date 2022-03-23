package cache

import (
	"time"
)

type InterfaceCache interface {
	//bigcache
	Get(key string) ([]byte, error)
	Set(key string, entry []byte) error
	Append(key string, entry []byte) error
	Delete(key string) error
	Reset() error
	Len() int
	Capacity() int
	Iterator() *EntryInfoIterator

	//gomemcache
	FlushAll() error
	Get(key string) (item *Item, err error)
	Touch(key string, seconds int32)
	GetMulti(keys []string) (map[string]*Item, error)
	Set(item *Item) error
	Add(item *Item) error
	Replace(item *Item) error
	CompareAndSwap(item *Item) error
	Delete(key string) error
	DeleteAll() error

	//redis
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
	Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd
	Get(ctx context.Context, key string) *StringCmd
	Del(ctx context.Context, keys ...string) *IntCmd
	Append(ctx context.Context, key, value string) *IntCmd
}

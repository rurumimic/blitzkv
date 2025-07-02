package kvstore

type KVStore interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
}

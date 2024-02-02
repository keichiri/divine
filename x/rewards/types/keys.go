package types

const (
	// ModuleName defines the module name
	ModuleName = "rewards"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rewards"
)

var (
	ParamsKey = []byte("p_rewards")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

package types

const (
	// ModuleName defines the module name
	ModuleName = "nft"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_nft"
)

var (
	ParamsKey = []byte("p_nft")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

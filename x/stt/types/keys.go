package types

const (
	// ModuleName defines the module name
	ModuleName = "stt"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stt"
)

var (
	ParamsKey = []byte("p_stt")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

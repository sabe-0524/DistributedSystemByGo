package log

type Config struct {
	Segment struct {
		MaxStoreBytes uint64 `json:"max_store_bytes"`
		MaxIndexBytes uint64 `json:"max_index_bytes"`
		InitialOffset  uint64 `json:"initial_offset"`
	} `json:"segment"`
}
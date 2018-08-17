package metadata

type IMetadata interface {
	// 存储metadata
	Set(dna []byte, data []byte) error

	// 根据dna查询数据
	GetWithDna(dna ... []byte) [][]byte

	// 根据ID查询数据
	GetWithID(id ... []byte) [][]byte

	// 根据类别查询
	GetWithCategory(path ... string) [][]byte

	// 根据tag查询
	GetWithTag(path ... string) [][]byte
}

package index

type IIndex interface {
	// 添加索引
	Add(data map[string][]string) error

	// 获得索引的并集
	Get(ks ... string) []string

	// 删除索引
	Delete(k string, id string) error
}

package tag

type ITag interface {
	// 添加tag
	Add(data map[string][]string) error

	// 获得tag对应的metadata
	Get(ks ... string) []string

	// 根据所有的分页
	GetAll() []string

	// 删除tag
	Delete(k string, id string) error
}

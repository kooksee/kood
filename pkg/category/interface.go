package category

// 类别,分组操作
type ICategory interface {
	//	创建一个新文件
	Add(parent string, name string, isDir bool) error

	// 删除一个新文件
	Delete(parent, name string) error

	// 获得一个新文件
	Get(name string, isDir bool)

	// 修改名称
	ReName(name, newName string) error

	// 移动文件
	Move(name, fromParent, toParent string) error
}

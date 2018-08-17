package object

import "github.com/kooksee/kdb"

const objectName = "obj"

func NewObject(name []byte) IObject {
	return &object{db: db.KHash([]byte(objectName))}
}

type object struct {
	IObject
	db kdb.IKHash
}

func (o *object) Set(k []byte, data []byte) error {
	return o.db.Set(k, data)
}

func (o *object) Get(k []byte) []byte {
	data, err := o.db.Get(k)
	if err != nil {
		log.Error("object get error", "err", err.Error())
		return nil
	}
	return data
}

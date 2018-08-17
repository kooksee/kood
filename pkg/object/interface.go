package object

type IObject interface {
	Set(k []byte, data []byte) error
	Delete(k []byte) error
	Get(k []byte) []byte
}

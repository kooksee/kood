package oss

type IOSS interface {
	Set(k []byte, data []byte) error
	Delete(k []byte) error
	Get(k []byte) []byte
}

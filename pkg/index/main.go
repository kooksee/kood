package index

import (
	"github.com/kooksee/kdb"
	"math/big"
	"time"
)

const indexName = "idx"

func NewIndex() IIndex {
	return &index{name: indexName}
}

type index struct {
	IIndex
	name string
}

func (i *index) getDb(name string) kdb.IKHash {
	return db.KHash(append([]byte(i.name), name...))
}

func (i *index) Add(data map[string][]string) error {
	for k, v := range data {
		idb := i.getDb(k)
		for _, h := range v {
			if err := idb.Set([]byte(h), big.NewInt(time.Now().Unix()).Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (i *index) Get(ks ... string) (ret []string) {
	dataSet := make(map[string]bool)
	// 获得第一次过滤的值
	if err := i.getDb(ks[0]).Range(func(key, _ []byte) error {
		dataSet[string(key)] = true
		return nil
	}); err != nil {
		log.Error("index Get error", "err", err.Error(), "key", ks)
		return nil
	}

	// 过滤以后的时
	for _, k := range ks[1:] {
		dataFilter := make(map[string]bool)
		if err := i.getDb(k).Range(func(key, _ []byte) error {
			if dataSet[string(key)] {
				dataFilter[string(key)] = true
			}
			return nil
		}); err != nil {
			log.Error("index Get error", "err", err.Error(), "key", ks)
			return nil
		}
		dataSet = dataFilter
	}

	for k := range dataSet {
		ret = append(ret, k)
	}
	return ret
}

func (i *index) Delete(k string, id string) error {
	return i.getDb(k).Del([]byte(id))
}

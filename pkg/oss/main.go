package oss

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

type ossImp struct {
	IOSS
	bucket *oss.Bucket
}

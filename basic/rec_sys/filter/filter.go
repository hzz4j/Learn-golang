package filter

import "rec_sys/common"

type Filter interface {
	Filter([]*common.Product) []*common.Product
	Name() string
}

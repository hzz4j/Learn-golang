package sort

import "rec_sys/common"

type Sorter interface {
	Sort([]*common.Product) []*common.Product
	Name() string
}

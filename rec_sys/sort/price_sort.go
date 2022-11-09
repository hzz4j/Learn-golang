package sort

import (
	"rec_sys/common"
	"sort"
)

type PriceSorter struct {
	Tag string
}

func (s *PriceSorter) Sort(products []*common.Product) []*common.Product {

	sort.Slice(products, func(i, j int) bool {
		return products[i].Price > products[j].Price
	})
	return products
}

func (s *PriceSorter) Name() string {
	return s.Tag
}

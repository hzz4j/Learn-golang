package filter

import "rec_sys/common"

type SaleFilter struct {
	Tag string
}

func (filter *SaleFilter) Filter(products []*common.Product) []*common.Product {
	result := make([]*common.Product, 0, len(products))
	for _, p := range products {
		if p.Sale > 500 {
			result = append(result, p)
		}
	}
	return result
}

func (filter *SaleFilter) Name() string {
	return filter.Tag
}

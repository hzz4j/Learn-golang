package recall

import (
	"rec_sys/common"
	"sort"
)

var allProducts = []*common.Product{
	{Id: 1, Sale: 610, Name: "p1", Price: 1000},
	{Id: 2, Sale: 322, Name: "p2", Price: 199},
	{Id: 3, Sale: 811, Name: "p3", Price: 230},
	{Id: 4, Sale: 567, Name: "p4", Price: 870},
	{Id: 5, Sale: 99, Name: "p5", Price: 789},
}

type HotRecaller struct {
	Tag string
}

func (h *HotRecaller) Recall(n int) []*common.Product {
	if n > len(allProducts) {
		n = len(allProducts)
	}

	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].Sale > allProducts[j].Sale
	})

	tmp := allProducts[:n]
	result := make([]*common.Product, 0, len(tmp))
	result = append(result, tmp...)
	return result
}

func (h *HotRecaller) Name() string {
	return h.Tag
}

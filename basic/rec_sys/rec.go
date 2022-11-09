package main

import (
	"log"
	"rec_sys/common"
	"rec_sys/filter"
	"rec_sys/recall"
	"rec_sys/sort"
	"time"
)

type Recommender struct {
	Recallers []recall.Recaller
	Sort      sort.Sorter
	Filter    []filter.Filter
}

func (rec *Recommender) Rec() []*common.Product {
	// 召回
	RecallMap := make(map[int]*common.Product, 100)
	for _, recaller := range rec.Recallers {
		begin := time.Now()
		products := recaller.Recall(10)
		log.Printf("召回%s耗时%dns,召回了%d个商品\n", recaller.Name(), time.Since(begin).Nanoseconds(), len(products))
		for _, product := range products {
			RecallMap[product.Id] = product
		}
	}
	log.Printf("一共召回了%d个商品\n", len(RecallMap))
	RecallSlice := make([]*common.Product, len(RecallMap))

	for _, product := range RecallMap {
		RecallSlice = append(RecallSlice, product)
	}

	// 排序
	begin := time.Now()
	SortedResult := rec.Sort.Sort(RecallSlice)
	log.Printf("排序耗时%dns\n", time.Since(begin).Nanoseconds())

	// 过滤
	FilteredResult := SortedResult
	for _, filter := range rec.Filter {
		begin := time.Now()
		FilteredResult = filter.Filter(FilteredResult)
		log.Printf("过滤%s规则耗时%d\n", filter.Name(), time.Since(begin).Nanoseconds())
	}
	return FilteredResult
}

func main() {
}

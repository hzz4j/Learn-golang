package recall

import "rec_sys/common"

// 召回
type Recaller interface {
	Recall(int) []*common.Product
	Name() string
}

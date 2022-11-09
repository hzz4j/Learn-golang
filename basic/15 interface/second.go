package main

type Transporter interface {
	whilstle(int) int
}

type Steamer interface {
	Transporter // 接口嵌入
	displacement() int
}

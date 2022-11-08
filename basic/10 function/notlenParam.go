package main

// 不定参数要放在最后面
func func1(name string, books ...string) {
	println(name + " learning")
	for _, book := range books {
		println(book)
	}
}

func main() {
	books := []string{"Javascript", "Java", "Golang"}
	// 解构
	func1("静默", books...)
}

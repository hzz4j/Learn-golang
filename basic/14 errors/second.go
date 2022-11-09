package main

import "fmt"

type PathError struct {
	Path   string
	Op     string
	OpTime string
	Msg    string
}

// 实现接口函数
func (e *PathError) Error() string {
	return fmt.Sprintf("%s: %s %s at %s", e.Op, e.Path, e.Msg, e.OpTime)
}

func NewPathError(path string, msg string) *PathError {
	return &PathError{
		Path:   path,
		Msg:    msg,
		Op:     "create",
		OpTime: "Wed Nov 09 2022 15:33:05 GMT+0800",
	}
}

func deletePath(path string) error {
	if 2 > 1 {
		return NewPathError(path, "Not found")
	}
	return nil
}

func main() {
	if err := deletePath("/usr/loadbal/1.jpg"); err != nil {
		fmt.Println(err.Error())
	}
}

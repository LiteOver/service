package main

import (
	"Service/FileMask"
	"fmt"
	"strings"
)

func main() {
	mProducer := FileMask.NewProducer("/Users/bocman/GolandProjects/Service/file.txt")
	mPresenter := FileMask.NewPresenter("/Users/bocman/GolandProjects/test/fi.txt")
	workNewService := FileMask.NewService(mProducer, mPresenter)

	st, _ := workNewService.Prod.Produce()
	result := make([]string, 0)

	result = append(result, workNewService.Mask([]byte(strings.Join(st, "\n"))))
	fmt.Println(workNewService.Pres.Present(result))
}

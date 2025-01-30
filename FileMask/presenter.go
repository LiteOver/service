package FileMask

import (
	"fmt"
	"log"
	"os"
)

type Presenter interface {
	Present([]string) error
}

type Present struct {
	Adress string
}

func NewPresenter(adress string) *Present {
	if adress == "" {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		return &Present{Adress: dir + "/newfile.txt"}
	}
	return &Present{Adress: adress}
}
func (p Present) Present(s []string) error {
	fmt.Println(p.Adress)
	file, err := os.Create(p.Adress)

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	for _, line := range s {
		file.Write([]byte(line))
	}
	return nil

}

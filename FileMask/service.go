package FileMask

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

type Producer interface {
	Produce() ([]string, error)
}

type Prod struct {
	Adress string
}

func NewProducer(adress string) *Prod {
	return &Prod{Adress: adress}
}

func (p Prod) Produce() ([]string, error) {
	file, err := os.Open(p.Adress)
	res := make([]string, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
	defer file.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}

	res = append(res, wr.String())
	return res, nil
}

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

type Service struct {
	Prod Producer
	Pres Presenter
}

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{Prod: prod, Pres: pres}
}

func (s *Service) Mask(data []byte) string {
	var prefix string = "http://"
	for i := range data {
		if i < len(data)-len(prefix) && string(data[i:i+len(prefix)]) == prefix {
			for i < len(data)-len(prefix) && (data[i+len(prefix)] != 32) {
				data[i+len(prefix)] = 42
				i++
			}
		}
	}
	return string(data)
}

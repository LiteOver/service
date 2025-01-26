package Service_test

import (
	"Service"
	"github.com/google/go-cmp/cmp"
	"os"
	"strings"
	"testing"
)

func TestNewProducer(t *testing.T) {
	expected := &Service.Prod{
		Adress: "/Users/bocman/GolandProjects/Service/file.txt",
	}
	result := Service.NewProducer("/Users/bocman/GolandProjects/Service/file.txt")
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
	expected2 := &Service.Prod{
		Adress: "",
	}
	result2 := Service.NewProducer("")
	if diff := cmp.Diff(expected2, result2); diff != "" {
		t.Error(diff)
	}
}

func TestProduce(t *testing.T) {
	data := []struct {
		name     string
		prod     Service.Prod
		expected []string
	}{
		{"first", Service.Prod{Adress: "/Users/bocman/GolandProjects/Service/file.txt"}, []string{"http:// it is test http://dsadsa and i won http://*"}},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := d.prod.Produce()
			if err != nil {
				t.Error(err)
			}
			if strings.Join(result, "\n") != strings.Join(d.expected, "\n") {
				t.Errorf("got %v, want %v", result, d.expected)
			}
		})
	}
}

func TestNewPresentor(t *testing.T) {
	expected := &Service.Present{
		Adress: "/Users/bocman/GolandProjects/Service/file.txt",
	}
	result := Service.NewPresenter("/Users/bocman/GolandProjects/Service/file.txt")
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
	expected2 := &Service.Present{
		Adress: "/Users/bocman/GolandProjects/Service/newfile.txt",
	}
	result2 := Service.NewPresenter("")
	if diff := cmp.Diff(expected2, result2); diff != "" {
		t.Error(diff)
	}
}
func TestPresent(t *testing.T) {
	p1 := Service.Present{Adress: "/Users/bocman/GolandProjects/Service/f1.txt"}
	err := p1.Present([]string{"some text http://hi and bye"})
	if err != nil {
		t.Error(err)
	}
	os.Remove(p1.Adress)

	p2 := Service.NewPresenter("")
	err = p2.Present([]string{"some text http://hi and bye"})
	if err != nil {
		t.Error(err)
	}
	os.Remove(p2.Adress)
}

func TestNewService(t *testing.T) {
	expected := &Service.Service{&Service.Prod{Adress: "/Users/bocman/GolandProjects/Service/file.txt"}, &Service.Present{"/Users/bocman/GolandProjects/Service/find.txt"}}
	result := Service.NewService(Service.NewProducer("/Users/bocman/GolandProjects/Service/file.txt"), Service.NewPresenter("/Users/bocman/GolandProjects/Service/find.txt"))
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestMask(t *testing.T) {
	mProducer := Service.NewProducer("/Users/bocman/GolandProjects/Service/file.txt")
	s := Service.Service{mProducer, Service.Present{}}
	expected := []string{"http:// it is test http://****** and i won http://*"}
	result := make([]string, 0)
	data, _ := s.Prod.Produce()
	result = append(result, s.Mask([]byte(strings.Join(data, "\n"))))
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}

}

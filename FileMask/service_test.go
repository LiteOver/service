package FileMask_test

import (
	"Service/FileMask"
	"github.com/google/go-cmp/cmp"
	"os"
	"strings"
	"testing"
)

func TestNewProducer(t *testing.T) {
	expected := &FileMask.Prod{
		Adress: "/Users/bocman/GolandProjects/Service/file.txt",
	}
	result := FileMask.NewProducer("/Users/bocman/GolandProjects/Service/file.txt")
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
	expected2 := &FileMask.Prod{
		Adress: "",
	}
	result2 := FileMask.NewProducer("")
	if diff := cmp.Diff(expected2, result2); diff != "" {
		t.Error(diff)
	}
}

func TestProduce(t *testing.T) {
	data := []struct {
		name     string
		prod     FileMask.Prod
		expected []string
	}{
		{"first", FileMask.Prod{Adress: "/Users/bocman/GolandProjects/Service/file.txt"}, []string{"http:// it is test http://dsadsa and i won http://*"}},
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
	expected := &FileMask.Present{
		Adress: "/Users/bocman/GolandProjects/Service/file.txt",
	}
	result := FileMask.NewPresenter("/Users/bocman/GolandProjects/Service/file.txt")
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
	expected2 := &FileMask.Present{
		Adress: "/Users/bocman/GolandProjects/Service/FileMask/newfile.txt",
	}
	result2 := FileMask.NewPresenter("")
	if diff := cmp.Diff(expected2, result2); diff != "" {
		t.Error(diff)
	}
}
func TestPresent(t *testing.T) {
	p1 := FileMask.Present{Adress: "/Users/bocman/GolandProjects/Service/f1.txt"}
	err := p1.Present([]string{"some text http://hi and bye"})
	if err != nil {
		t.Error(err)
	}
	os.Remove(p1.Adress)

	p2 := FileMask.NewPresenter("")
	err = p2.Present([]string{"some text http://hi and bye"})
	if err != nil {
		t.Error(err)
	}
	os.Remove(p2.Adress)
}

func TestNewService(t *testing.T) {
	expected := &FileMask.Service{&FileMask.Prod{Adress: "/Users/bocman/GolandProjects/Service/file.txt"}, &FileMask.Present{"/Users/bocman/GolandProjects/Service/find.txt"}}
	result := FileMask.NewService(FileMask.NewProducer("/Users/bocman/GolandProjects/Service/file.txt"), FileMask.NewPresenter("/Users/bocman/GolandProjects/Service/find.txt"))
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestMask(t *testing.T) {
	mProducer := FileMask.NewProducer("/Users/bocman/GolandProjects/Service/file.txt")
	s := FileMask.Service{mProducer, FileMask.Present{}}
	expected := []string{"http:// it is test http://****** and i won http://*"}
	result := make([]string, 0)
	data, _ := s.Prod.Produce()
	result = append(result, s.Mask([]byte(strings.Join(data, "\n"))))
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}

}

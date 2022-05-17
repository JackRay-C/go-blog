package transform

import (
	"testing"
)

type Target struct {
	ID int
	Username string
	File *File
}

type File struct {
	Name string
	Url string
}

type Ref struct {
	ID int
	Username string
	Age int
	Password string
	File *File
}

func Test01(t *testing.T) {
	t.Parallel()
	ref := Ref{
		ID:       1,
		Username: "renhj",
		Age:      10,
		Password: "123",
		File: &File{
			Name: "a",
			Url: "a",
		},
	}

	target := &Target{}

	if err := Transition(ref, target); err != nil {
		t.Errorf("tranform failed: %s", err)
	}

	if target.ID!=ref.ID && target.Username != ref.Username && target.File.Name != ref.File.Name {
		t.Errorf("tranform failed! ")
	}
}

func Benchmark_01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ref := Ref{
			ID:       1,
			Username: "renhj",
			Age:      10,
			Password: "123",
			File: &File{
				Name: "a",
				Url: "a",
			},
		}

		target := &Target{}
		if err := Transition(ref, target); err != nil {
			return
		}
	}
}

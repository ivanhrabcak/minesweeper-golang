package field

import (
	"fmt"
	"math"
	"math/rand"
	"minesweeper/field/box"
	"os"
	"time"
)

type Field struct {
	Size  int
	Boxes []box.Box
}

func NewField(size int) Field {
	if size < 5 {
		fmt.Printf("Size cannot be smaller than 5\n")
		os.Exit(1)
	}
	if size > 20 {
		fmt.Printf("Size cannot be bigger than 20\n")
		os.Exit(1)
	}

	f := Field{
		Size: size,
	}

	f.Boxes = []box.Box{}

	for i := 0; i < size*size; i++ {
		f.Boxes = append(f.Boxes,
			box.NewBox(i, i == (f.Size*int(math.Floor(float64(f.Size/2))))-1-int(math.Floor(float64(f.Size/2)))+f.Size),
		)
	}

	return f
}

func (f *Field) IsFull() bool {
	for i := range f.Boxes {
		b := f.Boxes[i]

		if b.State == box.HIDDEN {
			return false
		}
		if b.State == box.FLAGGED && b.Value != box.BOMB {
			return false
		}
	}

	return true
}

func (f *Field) IsEmpty() bool {
	for i := range f.Boxes {
		if f.Boxes[i].Value == box.BOMB {
			return false
		}
	}

	return true
}

func (f *Field) Init() {
	// TODO
	bombCountLeft := f.Size

	for bombCountLeft != 0 {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(63)

		if f.Boxes[index].Selected {
			continue
		}

		if f.Boxes[index].Value == box.EMPTY {
			f.Boxes[index].Value = box.BOMB

			bombCountLeft--
		}
	}

	f.Uncover()
}

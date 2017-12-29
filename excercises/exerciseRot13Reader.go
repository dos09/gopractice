package excercises

import (
	"io"
	"os"
	"strings"
	"fmt"
)

func rot13(b byte) byte {
	var min, max, offset byte
	if b >= 65 && b <= 90 {
		min = 64
		max = 90
	} else if b >= 97 && b <= 122 {
		min = 96
		max = 122
	} else {
		return b
	}
	
	b = b + 13
	if b > max {
		offset = b - max
		b = min + offset
	}
	return b
}

type rot13Reader struct {
	r io.Reader
}

func (m rot13Reader) Read(b []byte) (int, error) {
	n, err := m.r.Read(b)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, nil
}

func RunExcerciseRot13Reader() {
	fmt.Println(" * RunExcerciseRot13Reader")
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

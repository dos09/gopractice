package basics

import (
	"fmt"
	"io"
	"strings"
)

func RunReaders() {
	fmt.Println(" * RunReaders")
	practiceReader1()
}

func practiceReader1() {
	fmt.Println("practiceReader1")
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

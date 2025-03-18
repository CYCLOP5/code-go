package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type FastReader struct {
	r   *bufio.Reader
	buf []byte
	pos int
}

func NewFastReader() *FastReader {
	return &FastReader{
		r: bufio.NewReaderSize(os.Stdin, 1<<20),
	}
}

func (fr *FastReader) readByte() (byte, error) {
	if fr.pos >= len(fr.buf) {
		var err error
		fr.buf, err = fr.r.ReadBytes('\n')
		if err != nil {
			return 0, err
		}
		fr.pos = 0
	}
	b := fr.buf[fr.pos]
	fr.pos++
	return b, nil
}

func (fr *FastReader) readInt() (int64, error) {
	sign := int64(1)
	var x int64 = 0

	for {
		b, err := fr.readByte()
		if err != nil {
			return 0, err
		}
		if b == '-' {
			sign = -1
			break
		}
		if b >= '0' && b <= '9' {
			x = int64(b - '0')
			break
		}
	}

	for {
		b, err := fr.readByte()
		if err != nil {
			break
		}
		if b < '0' || b > '9' {
			break
		}
		x = x*10 + int64(b-'0')
	}
	return sign * x, nil
}

func (fr *FastReader) readString(n int) (string, error) {
	res := make([]byte, n)
	i := 0
	for i < n {
		b, err := fr.readByte()
		if err != nil {
			return "", err
		}
		if b == '\n' || b == ' ' || b == '\r' {
			continue
		}
		res[i] = b
		i++
	}
	return string(res), nil
}

func main() {
	scanner := NewFastReader()
	out := new(bytes.Buffer)

	t, err := scanner.readInt()
	if err != nil {
		return
	}

	for i := int64(0); i < t; i++ {

		a, err := scanner.readInt()
		if err != nil {
			return
		}
		b, err := scanner.readInt()
		if err != nil {
			return
		}

		res := a + b

		out.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(out.String())
}

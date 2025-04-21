package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"sort"
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
func readLine(fr *FastReader) (string, error) {
	var buf bytes.Buffer
	for {
		b, err := fr.readByte()
		if err != nil {
			return "", err
		}
		if b == '\n' {
			break
		}
		if b != '\r' {
			buf.WriteByte(b)
		}
	}
	return buf.String(), nil
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

func readUint64(fr *FastReader) (uint64, error) {
	x, err := fr.readInt()
	if err != nil {
		return 0, err
	}
	return uint64(x), nil
}


func (fr *FastReader) readLine() (string, error) {
	line, err := fr.r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimRight(line, "\r\n"), nil
}


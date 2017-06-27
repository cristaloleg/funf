package iochan

import (
	"bufio"
	"io"
)

type Mode int

const (
	Delimiter Mode = iota
	Timer
	Size
)

func NewReader(r io.Reader, m Mode, arg interface{}) <-chan string {
	ch := make(chan string)

	delim := func() {
		buf := bufio.NewReader(r)
		delimiter, ok := arg.(byte)
		if !ok {
			return
		}

		for {
			line, err := buf.ReadString(delimiter)
			if err != nil {
				break
			}

			ch <- line
		}

		close(ch)
	}

	timer := func() {
		buf := bufio.NewReader(r)

		for {
			line, err := buf.ReadString(byte(13))
			if err != nil {
				break
			}

			ch <- line
		}

		close(ch)
	}

	sized := func() {
		buf := bufio.NewReader(r)
		size, ok := arg.(int)
		if !ok {
			return
		}

		for {
			line := make([]byte, size)
			if _, err := buf.Read(line); err != nil {
				break
			}

			ch <- string(line)
		}

		close(ch)
	}

	switch m {
	case Delimiter:
		go delim()

	case Timer:
		go timer()

	case Size:
		go sized()
	}

	return ch
}

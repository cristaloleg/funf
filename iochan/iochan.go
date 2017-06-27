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

// ReaderDelimiter reads from io.Reader till delimiter, send readed bytes to channel
// works while io.Reader
func ReaderDelimiter(r io.Reader, delimiter byte) <-chan string {
	ch := make(chan string)

	go func() {
		buf := bufio.NewReader(r)

		for {
			line, err := buf.ReadString(delimiter)
			if err != nil {
				break
			}

			ch <- line
		}

		close(ch)
	}()

	return ch
}

// ReaderSize reads from io.Reader upto size, send readed bytes to channel
func ReaderSize(r io.Reader, size int) <-chan string {
	ch := make(chan string)

	go func() {
		buf := bufio.NewReader(r)

		for {
			line := make([]byte, size)
			if _, err := buf.Read(line); err != nil {
				break
			}

			ch <- string(line)
		}

		close(ch)
	}()

	return ch
}

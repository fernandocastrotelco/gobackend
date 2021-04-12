package main

import (
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/fernandocastrotelco/gobackend/functions/simplemath"
)

type MathExpr = string

const (
	AddExpr = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func main() {
	if err := ReadFullFile(); err != nil {
		fmt.Printf("something bad occurred: %s", err)
	}
}

func ReadFullFile() (err error) {
	var r io.ReadCloser = &SimpleReader{}
	defer func() {
		_ = r.Close()
		if p := recover(); p == errCatastrophicReader {
			println(p)
			err = errors.New("a panic occurred but it is ok")
		} else if p != nil {
			panic("an unexpected error occurred and we do not want to recover")
		}
	}()

	defer func() {
		println("before for-loop")
	}()

	for {
		value, readerErr := r.Read([]byte("text that does nothing"))
		if readerErr == io.EOF {
			println("finished reading file, breaking out of loop")
			break
		} else if readerErr != nil {
			err = readerErr
			return
		}

		println(value)
	}

	defer func() {
		println("after for-loop")
	}()

	return nil
}

var errNumberTwo = errors.New("hit the second number")

func ReadSomethingBad() error {
	var r io.Reader = BadReader{errors.New("my nonsense reader")}
	value, err := r.Read([]byte("test something"))
	if err != nil {
		fmt.Printf("an error occurred %s", err)
		return err
	}

	println(value)

	return nil
}

type BadReader struct {
	err error
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

type SimpleReader struct {
	count int
}

var errCatastrophicReader = errors.New("something catastrophic occurred in the reader")

func (br *SimpleReader) Read(p []byte) (n int, err error) {
	if br.count == 2 {
		panic(errors.New("another error"))
	}
	if br.count > 3 {
		return 0, io.EOF
	}
	br.count += 1
	return br.count, nil
}

func (br *SimpleReader) Close() error {
	println("closing reader")
	return nil
}

func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}

func mathExpression(expr MathExpr) func(float64, float64) float64 {
	switch expr {
	case AddExpr:
		return simplemath.Add
	case SubtractExpr:
		return simplemath.Subtract
	case MultiplyExpr:
		return simplemath.Multiply
	default:
		panic("an invalid math expression was used")
	}
}

func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
}


package main

import (
	"errors"
	"fmt"
	"go-basic/functions/simplemath"
	"io"
	"net/http"
	"strings"
)

type BadReader struct {
	err error
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

func ReadSomething() error {

	var r io.Reader = BadReader{errors.New("My reader")}
	value, error := r.Read([]byte("Test something"))
	if error != nil {
		fmt.Printf("An error occurred %s", error)
		return error
	}
	println(value)
	return nil
}

type SimpleReader struct {
	count int
}

func (s SimpleReader) Read(p []byte) (n int, err error) {
	println(s.count)
	if s.count > 3 {
		return 0, io.EOF
	}
	s.count += 1
	return s.count, nil
}

func main() {

	ReadSomething()
	func() {
		println("My first anonumous fuction")
	}()

	a := func() {
		println("My first anonumous fuction")
	}
	a()

	sv := simplemath.NewSemanticVersion(1, 2, 3)
	sv.IncrementMajor()
	println(sv.String())
	answer, err := testFunc("A", "B")
	println(answer)
	println(err.Error())

	answer2, _ := testFunc("A", "B")
	println(answer2)

	numbers := []float64{12.2, 14, 16, 22.4}
	total := sum(numbers...)
	println(total)

	var tripper http.RoundTripper = &RoundTripCounter{}
	r, _ := http.NewRequest(http.MethodGet, "http://pluralsight.com", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)
}

func mathExpression() func(float64, float64) float64 {
	return func(f float64, f2 float64) float64 {
		return f + f2
	}
}

func double(f1, f2 float64, mathExp func(float64, float64) float64) float64 {
	return 0.0
}

type RoundTripCounter struct {
	count int
}

func (rt *RoundTripCounter) RoundTrip(tripper *http.Request) (*http.Response, error) {
	rt.count += 1
	return nil, nil
}

func testFunc(p1, p2 string) (string, error) {
	return "test func", errors.New("test error")
}
func testFuncNamed(p1, p2 string) (answer string, err error) {
	answer = "test func"
	err = errors.New("test error")
	return
}

func sum(values ...float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total
}

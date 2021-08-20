package blog


import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"io"
	"os"
)

type MyInt int

func TestOne(t *testing.T) {
	var i int
	var j MyInt

	iType := reflect.TypeOf(i)
	jType := reflect.TypeOf(j)

	t.Log("i type:", iType)
	t.Log("j type:", jType)

	t.Log("i kind:", iType.Kind())
	t.Log("j kind:", jType.Kind())
}

func TestTwo(t *testing.T) {
	var r io.Reader
	r = os.Stdin
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)
	t.Log(r)
}

func TestThree(t *testing.T) {
	var r io.Reader
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		t.Log(err)
	}
	r = tty
	t.Log(r)

	var w io.Writer
	w = r.(io.Writer)
	t.Log(w)

	var empty interface{}
	empty = w
	t.Log(empty)
}

func TestFirstLaw(t *testing.T) {
	var x float64 = 3.4
	t.Log("a", "b")
	t.Log("type:", reflect.TypeOf(x))

	t.Log("value:", reflect.ValueOf(x).String())
}

func TestFirstLaw1(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	t.Log("type:", v.Type())
	t.Log("kind is float64:", v.Kind() == reflect.Float64)
	t.Log("value:", v.Float())
}

func TestFirstLaw2(t *testing.T) {
	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	t.Log("type:", v.Type())                            // uint8.
	t.Log("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	x = uint8(v.Uint())                                       // v.Uint returns a uint64.
}

func TestSecondLaw(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	y := v.Interface().(float64)
	t.Log("y:", y)

	// t.Log(v.Interface())
	// t.Log(v)
	fmt.Printf("value is %7.1e\n", v.Interface())
}

func TestThirdLaw(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// v.SetFloat(7.1)
	t.Log("settability of v:", v.CanSet())
}

func TestThirdLaw1(t *testing.T) {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}

type T struct {
	A int
	B string
}

func TestStruct(test *testing.T) {
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	fmt.Println()

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset strip")
	fmt.Println("t is now", t)
}
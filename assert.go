package go_test_utils

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type any = interface{}

func AssertCallResult(t *testing.T, callf string, callv, expected, actual []any, expectedPositive bool) {
	t.Helper()

	if reflect.DeepEqual(expected, actual) != expectedPositive {
		buf := &bytes.Buffer{}

		buf.Write([]byte("Got unexpected result from "))
		fmt.Fprintf(buf, callf, callv...)
		buf.Write([]byte(":"))

		for _, e := range expected {
			buf.Write([]byte("\n- "))
			fprintHumanReadable(buf, e)
		}

		for _, a := range actual {
			buf.Write([]byte("\n+ "))
			fprintHumanReadable(buf, a)
		}

		t.Error(buf.String())
	}
}

func fprintHumanReadable(w *bytes.Buffer, v any) {
	switch x := v.(type) {
	case nil:
		w.Write([]byte("nil"))
	case bool:
		fmt.Fprintf(w, "bool(%v)", x)
	case uint:
		fmt.Fprintf(w, "uint(%d)", x)
	case uint8:
		fmt.Fprintf(w, "uint8(%d)", x)
	case uint16:
		fmt.Fprintf(w, "uint16(%d)", x)
	case uint32:
		fmt.Fprintf(w, "uint32(%d)", x)
	case uint64:
		fmt.Fprintf(w, "uint64(%d)", x)
	case int:
		fmt.Fprintf(w, "int(%d)", x)
	case int8:
		fmt.Fprintf(w, "int8(%d)", x)
	case int16:
		fmt.Fprintf(w, "int16(%d)", x)
	case int32:
		fmt.Fprintf(w, "int32(%d)", x)
	case int64:
		fmt.Fprintf(w, "int64(%d)", x)
	case float32:
		fmt.Fprintf(w, "float32(%s)", strconv.FormatFloat(float64(x), 'g', -1, 64))
	case float64:
		fmt.Fprintf(w, "float64(%s)", strconv.FormatFloat(x, 'g', -1, 64))
	case string:
		fmt.Fprintf(w, "string(%#v)", x)
	case []byte:
		fmt.Fprintf(w, "[]byte(%#v)", string(x))
	default:
		fmt.Fprintf(w, "%#v", v)
	}
}

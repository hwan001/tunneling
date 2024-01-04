package func1_test

import (
	"example/code_20231227/func1"
	"testing"
)

func func1_test() (t *testing.T) {
	case1 := "johnny"

	if func1.Hello(case1) != "Hi, johnny. Welcome!" {
		t.Error("case1 error")
	}

}

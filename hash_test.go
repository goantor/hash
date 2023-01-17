package hash

import (
	"testing"
)

func TestMd5(t *testing.T) {

	n := Md5("123123")
	if n.String() == "4297f44b13955235245b2497399d7a93" {
		t.Logf("succeeded")
	}

}

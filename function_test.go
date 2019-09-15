package b

import (
	"log"
	"strings"
	"testing"
)

func TestF(t *testing.T) {
	s := "foo"
	x := F("foo")
	log.Println(x)
	if !strings.HasPrefix(x, "A v") {
		t.Errorf("errornous prefix")
	}
	if !strings.HasSuffix(x, s) {
		t.Errorf("errornous suffix")
	}
}

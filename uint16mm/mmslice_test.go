package uint16mm_test

import (
	"os"
	"testing"

	"github.com/brentp/mmslice/uint16mm"
	mmap "github.com/edsrzf/mmap-go"
)

func TestCreate(t *testing.T) {

	f, err := uint16mm.Create("/tmp/uint16s.bin", 2000)
	if err != nil {
		t.Errorf("error creating file")
	}

	if len(f.A) != 2000 {
		t.Errorf("expected length 2000. found %d\n", len(f.A))
	}

	for i := 0; i < 1000; i++ {
		f.A[i] = uint16(22)
	}
	for i := 0; i < 1000; i++ {
		if f.A[i] != uint16(22) {
			t.Errorf("expected to find 22. found %d\n", f.A[i])
		}
	}
	for i := 1000; i < 2000; i++ {
		if f.A[i] != uint16(0) {
			t.Errorf("expected to find 0 at index %d. found %d\n", i, f.A[i])
		}
	}

	if err := f.Flush(); err != nil {
		t.Errorf("error on flushing: %s", err)
	}
	if err := f.Close(); err != nil {
		t.Errorf("error on flushing: %s", err)
	}

	fh, err := os.Open("/tmp/uint16s.bin")
	if err != nil {
		t.Errorf("error on opening: %s", err)
	}

	r, err := uint16mm.Open(fh, mmap.RDONLY)
	if err != nil {
		t.Errorf("error on opening: %s", err)
	}
	if len(r.A) != 2000 {
		t.Errorf("expected length 2000. found %d\n", len(r.A))
	}

	for i := 0; i < 1000; i++ {
		if r.A[i] != uint16(22) {
			t.Errorf("expected to find 22. found %d\n", r.A[i])
		}
	}
	for i := 1000; i < 2000; i++ {
		if r.A[i] != uint16(0) {
			t.Errorf("expected to find 0 at index %d. found %d\n", i, r.A[i])
		}
	}

	os.Remove("/tmp/uint16s.bin")

}
